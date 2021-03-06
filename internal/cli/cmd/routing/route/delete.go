// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package route

import (
	"emperror.dev/errors"
	"github.com/AlecAivazis/survey/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/types"

	"github.com/banzaicloud/backyards-cli/internal/cli/cmd/routing/common"

	cmdCommon "github.com/banzaicloud/backyards-cli/internal/cli/cmd/common"
	"github.com/banzaicloud/backyards-cli/internal/cli/cmd/util"
	"github.com/banzaicloud/backyards-cli/pkg/cli"
	"github.com/banzaicloud/backyards-cli/pkg/graphql"
	"github.com/banzaicloud/istio-client-go/pkg/networking/v1alpha3"
)

type deleteCommand struct{}

type deleteOptions struct {
	serviceID string

	matches []string

	parsedMatches []*v1alpha3.HTTPMatchRequest
	serviceName   types.NamespacedName
}

func newDeleteOptions() *deleteOptions {
	return &deleteOptions{}
}

func newDeleteCommand(cli cli.CLI) *cobra.Command {
	c := &deleteCommand{}
	options := newDeleteOptions()

	cmd := &cobra.Command{
		Use:           "delete [[--service=]namespace/servicename] [[--match=]field:kind=value] ...",
		Short:         "Delete http route of a service",
		Args:          cobra.MaximumNArgs(1),
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error

			if len(args) > 0 {
				options.serviceID = args[0]
			}

			if options.serviceID == "" {
				return errors.New("service must be specified")
			}

			options.serviceName, err = util.ParseK8sResourceID(options.serviceID)
			if err != nil {
				return errors.WrapIf(err, "could not parse service ID")
			}

			if len(options.matches) == 0 {
				return errors.New("at least one match must be specified")
			}

			options.parsedMatches, err = common.ParseHTTPRequestMatches(options.matches)
			if err != nil {
				return errors.WrapIf(err, "could not parse matches")
			}

			return c.run(cli, options)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&options.serviceID, "service", "", "Service name")
	flags.StringArrayVarP(&options.matches, "match", "m", options.matches, "HTTP request match")

	return cmd
}

func (c *deleteCommand) run(cli cli.CLI, options *deleteOptions) error {
	var err error

	client, err := cmdCommon.GetGraphQLClient(cli)
	if err != nil {
		return errors.WrapIf(err, "could not get initialized graphql client")
	}
	defer client.Close()

	service, err := client.GetService(options.serviceName.Namespace, options.serviceName.Name)
	if err != nil {
		return errors.WrapIf(err, "could not get service")
	}

	if len(service.VirtualServices) == 0 {
		return errors.Errorf("http route not found for %s", common.HTTPMatchRequests(common.ConvertHTTPMatchRequestsPointers(options.parsedMatches)))
	}

	matchedRoute := common.HTTPRoutes(service.VirtualServices[0].Spec.HTTP).GetMatchedRoute(options.parsedMatches)
	if matchedRoute == nil {
		return errors.Errorf("http route not found for %s", common.HTTPMatchRequests(common.ConvertHTTPMatchRequestsPointers(options.parsedMatches)))
	}

	if matchedRoute.Route == nil && matchedRoute.Redirect == nil {
		log.Infof("no route found for %s", options.serviceName)
		return nil
	}

	if cli.Interactive() {
		err = Output(cli, options.serviceName, *matchedRoute)
		if err != nil {
			return err
		}

		confirmed := false
		err = survey.AskOne(&survey.Confirm{Message: "Do you want to DELETE the http route?"}, &confirmed)
		if err != nil {
			return errors.WrapIf(err, "could not ask for confirmation")
		}
		if !confirmed {
			return errors.New("deletion cancelled")
		}
	}

	req := graphql.DisableHTTPRouteRequest{
		Selector: graphql.HTTPRouteSelector{
			Name:      service.Name,
			Namespace: service.Namespace,
			Matches:   options.parsedMatches,
		},
		Rules: []string{"Route"},
	}

	r2, err := client.DisableHTTPRoute(req)
	if err != nil {
		return err
	}

	if !r2 {
		return errors.New("unknown error: cannot delete http route")
	}

	log.Infof("http route with match %s set to %s successfully deleted", common.HTTPMatchRequests(common.ConvertHTTPMatchRequestsPointers(options.parsedMatches)), options.serviceName)

	return nil
}
