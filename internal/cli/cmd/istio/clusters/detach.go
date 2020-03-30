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

package clusters

import (
	"context"
	"fmt"

	"emperror.dev/errors"
	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"istio.io/operator/pkg/object"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	cmdCommon "github.com/banzaicloud/backyards-cli/internal/cli/cmd/common"
	"github.com/banzaicloud/backyards-cli/pkg/cli"
	k8sclient "github.com/banzaicloud/backyards-cli/pkg/k8s/client"
	"github.com/banzaicloud/backyards-cli/pkg/k8s/resourcemanager"
)

type detachCommand struct {
	cli cli.CLI
}

type DetachOptions struct {
	name           string
	kubeconfigPath string
}

func NewDetachOptions() *DetachOptions {
	return &DetachOptions{}
}

var namespacesOnPeerCluster = []string{"backyards-system", "istio-system"}

func NewDetachCommand(cli cli.CLI, options *DetachOptions) *cobra.Command {
	c := &detachCommand{
		cli: cli,
	}

	cmd := &cobra.Command{
		Use:   "detach [kubeconfig]",
		Args:  cobra.ExactArgs(1),
		Short: "Detach peer cluster from the mesh",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceErrors = true
			cmd.SilenceUsage = true

			options.kubeconfigPath = args[0]

			return c.run(options)
		},
	}

	cmd.Flags().StringVar(&options.name, "name", options.name, "Name override for the peer cluster")

	return cmd
}

func (c *detachCommand) run(options *DetachOptions) error {
	client, err := cmdCommon.GetGraphQLClient(c.cli)
	if err != nil {
		return errors.WrapIf(err, "could not get initialized graphql client")
	}
	defer client.Close()

	clusters, err := client.Clusters()
	if err != nil {
		return errors.WrapIf(err, "could not get clusters")
	}

	err = ConfirmKubeconfig(c.cli, options.kubeconfigPath)
	if err != nil {
		return err
	}

	k8sconfig, err := k8sclient.GetConfigWithContext(options.kubeconfigPath, "")
	if err != nil {
		return errors.WrapIf(err, "could not get k8s config")
	}

	k8sclient, err := k8sclient.NewClient(k8sconfig, k8sclient.Options{})
	if err != nil {
		return errors.WrapIf(err, "could not get k8s client")
	}

	if options.name == "" {
		options.name, err = GetClusterNameFromKubeconfig(options.kubeconfigPath)
		if err != nil {
			return err
		}
	}

	ok, peerCluster := clusters.GetClusterByName(options.name)
	if !ok || peerCluster.Type != "peer" {
		return errors.Errorf("peer cluster '%s' not found", options.name)
	}

	return c.cli.IfConfirmed(fmt.Sprintf("Detach peer cluster '%s'. Are you sure to proceed?", peerCluster.Name), func() error {
		ok, err = client.DetachPeerCluster(peerCluster.Name)
		if err != nil {
			return errors.WrapIf(err, "could not detach peer cluster")
		}

		if ok && c.cli.Interactive() {
			log.Infof("detaching peer cluster '%s' started successfully\n", options.name)
		}

		err = c.removePrometheusFederationService(peerCluster.Name)

		return c.removeNamespaces(k8sclient, namespacesOnPeerCluster)
	})
}

func (c *detachCommand) removePrometheusFederationService(clusterName string) error {
	k8sclient, err := c.cli.GetK8sClient()
	if err != nil {
		return err
	}

	var services corev1.ServiceList
	err = k8sclient.List(context.Background(), &services, client.MatchingLabels(map[string]string{
		"backyards.banzaicloud.io/cluster-name": clusterName,
	}))
	if err != nil {
		return err
	}

	for _, svc := range services.Items {
		svc := svc
		err = k8sclient.Delete(context.Background(), &svc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *detachCommand) removeNamespaces(client k8sclient.Client, namespace []string) error {
	labelManager := c.cli.LabelManager()
	m := resourcemanager.New(client, labelManager)

	objs := make(object.K8sObjects, 0)
	for _, ns := range namespace {
		obj, err := getNamespaceObject(ns)
		if err != nil {
			return err
		}
		objs = append(objs, obj)
	}

	m.SetObjects(objs)

	return m.Uninstall().Do()
}

func getNamespaceObject(namespace string) (*object.K8sObject, error) {
	manifest := fmt.Sprintf(heredoc.Doc(`
		apiVersion: v1
		kind: Namespace
		metadata:
		  name: %s
	`), namespace)

	objs, err := object.ParseK8sObjectsFromYAMLManifest(manifest)
	if err != nil {
		return nil, err
	}

	return objs[0], nil
}
