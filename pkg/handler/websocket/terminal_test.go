/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package websocket_test

import (
	"fmt"
	apiv1 "k8c.io/kubermatic/v2/pkg/api/v1"
	"k8c.io/kubermatic/v2/pkg/handler/test"
	"k8c.io/kubermatic/v2/pkg/handler/test/hack"
	"net/http/httptest"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
	"testing"
	"time"
)

func TestTerminalEndpoint(t *testing.T) {
	testcases := []struct {
		name            string
		existingAPIUser *apiv1.User
	}{
		{
			name:            "simple test",
			existingAPIUser: test.GenDefaultAPIUser(),
		},
	}

	//ctx := context.Background()

	// TODO: Create user project binding
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			seed := test.GenTestSeed()
			project := test.GenDefaultProject()
			cluster := test.GenDefaultCluster()
			binding := test.GenBinding(project.Name, tc.existingAPIUser.Email, "owners")

			runtimeObjects := []ctrlruntimeclient.Object{
				test.APIUserToKubermaticUser(*tc.existingAPIUser),
				seed,
				project,
				cluster,
				binding,
			}

			ep, _, err := test.CreateTestEndpointAndGetClients(*tc.existingAPIUser, nil, []ctrlruntimeclient.Object{},
				nil, runtimeObjects, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}
			server := httptest.NewServer(ep)
			defer server.Close()

			wsURL := "ws" + strings.TrimPrefix(server.URL, "http") +
				fmt.Sprintf("/api/v1/ws/projects/%s/clusters/%s/terminal", project.Name, cluster.Name)
			ch, err := createWSClient(wsURL)
			if err != nil {
				t.Fatalf("failed to initialize websocket client: %v", err)
			}

			var wsMsg wsMessage
			select {
			case <-time.After(5 * time.Second):
				t.Fatalf("timeout waiting for ws message")
			case wsMsg = <-ch:
			}

			fmt.Println(wsMsg)
		})
	}

}
