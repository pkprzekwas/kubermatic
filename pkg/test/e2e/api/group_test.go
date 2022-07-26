//go:build e2e

/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package api

import (
	"context"
	"testing"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils"

	"k8s.io/apimachinery/pkg/util/rand"
)

func TestOidcGroupSupport(t *testing.T) {
	ctx := context.Background()

	masterToken, err := utils.RetrieveMasterToken(ctx)
	if err != nil {
		t.Fatalf("failed to get master token: %v", err)
	}
	testClient := utils.NewTestClient(masterToken, t)
	project, err := testClient.CreateProject(rand.String(10))
	if err != nil {
		t.Fatalf("failed to create project: %v", err)
	}
	defer cleanupProject(t, project.ID)

	t.Logf("ID: %s, Name: %s", project.ID, project.Name)

	janeToken, err := utils.RetrieveLDAPToken(ctx)
	if err != nil {
		t.Fatalf("failed to get jane's token: %v", err)
	}
	t.Logf("oidc: %s", janeToken)
	janeClient := utils.NewTestClient(janeToken, t)
	p, err := janeClient.GetProject(project.ID)
	if err == nil {
		t.Logf("Project: %v", p)
		t.Fatalf("expected error")
	}
	t.Logf("error: %s", err.Error())
}
