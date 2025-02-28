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

package kubermatic

import (
	"fmt"
	"strings"

	"go.uber.org/zap"

	"k8c.io/kubermatic/v2/pkg/controller/operator/defaults"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	operatorv1alpha1 "k8c.io/kubermatic/v2/pkg/crd/operator/v1alpha1"
	"k8c.io/kubermatic/v2/pkg/resources/reconciling"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

const (
	serviceAccountName             = "kubermatic-seed"
	backupContainersConfigMapName  = "backup-containers"
	restoreS3SettingsConfigMapName = "s3-settings"
	storeContainerKey              = "store-container.yaml"
	cleanupContainerKey            = "cleanup-container.yaml"
	deleteContainerKey             = "delete-container.yaml"
	s3EndpointKey                  = "ENDPOINT"
	s3BucketNameKey                = "BUCKET_NAME"
	caBundleConfigMapName          = "ca-bundle"
)

func ClusterRoleBindingName(cfg *operatorv1alpha1.KubermaticConfiguration) string {
	return fmt.Sprintf("%s:%s-seed:cluster-admin", cfg.Namespace, cfg.Name)
}

func ServiceAccountCreator(cfg *operatorv1alpha1.KubermaticConfiguration, seed *kubermaticv1.Seed) reconciling.NamedServiceAccountCreatorGetter {
	return func() (string, reconciling.ServiceAccountCreator) {
		return serviceAccountName, func(sa *corev1.ServiceAccount) (*corev1.ServiceAccount, error) {
			return sa, nil
		}
	}
}

func ClusterRoleBindingCreator(cfg *operatorv1alpha1.KubermaticConfiguration, seed *kubermaticv1.Seed) reconciling.NamedClusterRoleBindingCreatorGetter {
	name := ClusterRoleBindingName(cfg)

	return func() (string, reconciling.ClusterRoleBindingCreator) {
		return name, func(crb *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
			crb.RoleRef = rbacv1.RoleRef{
				APIGroup: rbacv1.GroupName,
				Kind:     "ClusterRole",
				Name:     "cluster-admin",
			}

			crb.Subjects = []rbacv1.Subject{
				{
					Kind:      rbacv1.ServiceAccountKind,
					Name:      serviceAccountName,
					Namespace: cfg.Namespace,
				},
			}

			return crb, nil
		}
	}
}

func BackupContainersConfigMapCreator(cfg *operatorv1alpha1.KubermaticConfiguration, seed *kubermaticv1.Seed, log *zap.SugaredLogger) reconciling.NamedConfigMapCreatorGetter {
	return func() (string, reconciling.ConfigMapCreator) {
		return backupContainersConfigMapName, func(c *corev1.ConfigMap) (*corev1.ConfigMap, error) {
			if c.Data == nil {
				c.Data = make(map[string]string)
			}

			c.Data[storeContainerKey] = cfg.Spec.SeedController.BackupStoreContainer
			c.Data[cleanupContainerKey] = cfg.Spec.SeedController.BackupCleanupContainer

			if cfg.Spec.SeedController.BackupStoreContainer == strings.TrimSpace(defaults.DefaultBackupStoreContainer) &&
				(seed.Spec.BackupRestore != nil || seed.Spec.EtcdBackupRestore != nil) {
				c.Data[storeContainerKey] = strings.TrimSpace(defaults.DefaultNewBackupStoreContainer)
				log.Debugw("Defaulting field", "field", "seedController.backupRestoreContainer")
			}

			if seed.Spec.BackupRestore != nil || seed.Spec.EtcdBackupRestore != nil {
				if cfg.Spec.SeedController.BackupDeleteContainer == "" {
					cfg.Spec.SeedController.BackupDeleteContainer = strings.TrimSpace(defaults.DefaultNewBackupDeleteContainer)
					log.Debugw("Defaulting field", "field", "seedController.backupDeleteContainer")
				}

				c.Data[deleteContainerKey] = cfg.Spec.SeedController.BackupDeleteContainer
			}

			return c, nil
		}
	}
}

func RestoreS3SettingsConfigMapCreator(backupRestore *kubermaticv1.SeedBackupRestoreConfiguration) reconciling.NamedConfigMapCreatorGetter {
	if backupRestore == nil {
		return nil
	}

	return func() (string, reconciling.ConfigMapCreator) {
		return restoreS3SettingsConfigMapName, func(c *corev1.ConfigMap) (*corev1.ConfigMap, error) {
			if c.Data == nil {
				c.Data = make(map[string]string)
			}

			c.Data[s3EndpointKey] = backupRestore.S3Endpoint
			c.Data[s3BucketNameKey] = backupRestore.S3BucketName

			return c, nil
		}
	}
}

func CABundleConfigMapCreator(caBundle *corev1.ConfigMap) reconciling.NamedConfigMapCreatorGetter {
	return func() (string, reconciling.ConfigMapCreator) {
		return caBundleConfigMapName, func(c *corev1.ConfigMap) (*corev1.ConfigMap, error) {
			c.Data = caBundle.Data

			return c, nil
		}
	}
}
