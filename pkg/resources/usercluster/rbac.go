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

package usercluster

import (
	"fmt"

	"k8c.io/kubermatic/v2/pkg/resources"
	"k8c.io/kubermatic/v2/pkg/resources/reconciling"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	serviceAccountName = "kubermatic-usercluster-controller-manager"
	roleName           = "kubermatic:usercluster-controller-manager"
	roleBindingName    = "kubermatic:usercluster-controller-manager"
)

func ServiceAccountCreator() (string, reconciling.ServiceAccountCreator) {
	return serviceAccountName, func(sa *corev1.ServiceAccount) (*corev1.ServiceAccount, error) {
		return sa, nil
	}
}

func RoleCreator() (string, reconciling.RoleCreator) {
	return roleName, func(r *rbacv1.Role) (*rbacv1.Role, error) {
		r.Rules = []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"secrets"},
				Verbs: []string{
					"get",
					"list",
					"watch",
					"create",
				},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"configmaps"},
				Verbs: []string{
					"get",
					"list",
					"watch",
				},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"secrets"},
				ResourceNames: []string{
					resources.AdminKubeconfigSecretName,
				},
				Verbs: []string{"update"},
			},
			{
				APIGroups: []string{"kubermatic.k8s.io"},
				Resources: []string{"constraints"},
				Verbs: []string{
					"get",
					"list",
					"watch",
					"patch",
					"update",
				},
			},
		}
		return r, nil
	}
}

func RoleBindingCreator() (string, reconciling.RoleBindingCreator) {
	return roleBindingName, func(rb *rbacv1.RoleBinding) (*rbacv1.RoleBinding, error) {
		rb.RoleRef = rbacv1.RoleRef{
			Name:     roleName,
			Kind:     "Role",
			APIGroup: rbacv1.GroupName,
		}
		rb.Subjects = []rbacv1.Subject{
			{
				Kind: rbacv1.ServiceAccountKind,
				Name: serviceAccountName,
			},
		}
		return rb, nil
	}
}

func ClusterRole() reconciling.NamedClusterRoleCreatorGetter {
	return func() (string, reconciling.ClusterRoleCreator) {
		return roleName, func(r *rbacv1.ClusterRole) (*rbacv1.ClusterRole, error) {
			r.Rules = []rbacv1.PolicyRule{
				{
					APIGroups: []string{"kubermatic.k8s.io"},
					Resources: []string{"clusters"},
					Verbs: []string{
						"get",
						"list",
						"watch",
						"patch",
						"update",
					},
				},
			}
			return r, nil
		}
	}
}

func ClusterRoleBinding(namespace *corev1.Namespace) reconciling.NamedClusterRoleBindingCreatorGetter {
	return func() (string, reconciling.ClusterRoleBindingCreator) {
		return GenClusterRoleBindingName(namespace), func(rb *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
			rb.OwnerReferences = []metav1.OwnerReference{genOwnerReference(namespace)}
			rb.RoleRef = rbacv1.RoleRef{
				Name:     roleName,
				Kind:     "ClusterRole",
				APIGroup: rbacv1.GroupName,
			}
			rb.Subjects = []rbacv1.Subject{
				{
					Kind:      rbacv1.ServiceAccountKind,
					Name:      serviceAccountName,
					Namespace: namespace.Name,
				},
			}
			return rb, nil
		}
	}
}

// genOwnerReference returns an owner ref pointing to the cluster namespace. This
// ensures that when a cluster is deleted, the ClusterRole/Binding are deleted automatically
// *after* the cluster namespace is gone. Previously, we manually deleted them, but
// this could lead to cases where the usercluster-ctrl-mgr is still running and
// producing errors because it cannot access Cluster objects anymore.
func genOwnerReference(namespace *corev1.Namespace) metav1.OwnerReference {
	return metav1.OwnerReference{
		APIVersion: namespace.APIVersion,
		Kind:       namespace.Kind,
		Name:       namespace.Name,
		UID:        namespace.UID,
	}
}

func GenClusterRoleBindingName(targetNamespace *corev1.Namespace) string {
	return fmt.Sprintf("%s-%s", roleBindingName, targetNamespace.Name)
}
