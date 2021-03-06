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

package graphql

import (
	"context"

	"github.com/MakeNowJust/heredoc"

	"github.com/banzaicloud/istio-client-go/pkg/networking/v1alpha3"
)

type ApplySidecarEgressInput struct {
	Selector SidecarEgressSelector `json:"selector"`
	Egress   Egress                `json:"egress"`
}

type SidecarEgressSelector struct {
	Namespace      string             `json:"namespace"`
	WorkloadLabels *map[string]string `json:"workloadLabels,omitempty"`
	Port           *v1alpha3.Port     `json:"port,omitempty"`
	Bind           *string            `json:"bind,omitempty"`
}

type Egress struct {
	Port  *v1alpha3.Port `json:"port,omitempty"`
	Bind  *string        `json:"bind,omitempty"`
	Hosts []string       `json:"hosts,omitempty"`
}

type ApplySidecarEgressResponse bool

func (c *client) ApplySidecarEgress(input ApplySidecarEgressInput) (ApplySidecarEgressResponse, error) {
	request := heredoc.Doc(`
	  mutation applySidecarEgress(
        $input: ApplySidecarEgressInput!
      ) {
        applySidecarEgress(
          input: $input
        )
      }
`)

	r := c.NewRequest(request)
	r.Var("input", input)

	// run it and capture the response
	var respData map[string]ApplySidecarEgressResponse
	if err := c.client.Run(context.Background(), r, &respData); err != nil {
		return false, err
	}

	return respData["applySidecarEgress"], nil
}
