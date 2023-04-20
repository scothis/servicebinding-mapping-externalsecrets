/*
Copyright 2023 the original author or authors.

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

package controllers

import (
	"context"

	"github.com/vmware-labs/reconciler-runtime/reconcilers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	externalsecretsv1beta1 "github.com/servicebinding/mapping-externalsecrets/apis/thirdparty/externalsecrets/v1beta1"
	mappingv1alpha1 "github.com/servicebinding/mapping-externalsecrets/apis/v1alpha1"
)

//+kubebuilder:rbac:groups=external-secrets.io,resources=externalsecrets,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete

// ExternalSecretReconciler reconciles a ExternalSecret object
func ExternalSecretReconciler(c reconcilers.Config) *reconcilers.ResourceReconciler[*externalsecretsv1beta1.ExternalSecret] {
	return &reconcilers.ResourceReconciler[*externalsecretsv1beta1.ExternalSecret]{
		SkipStatusUpdate: true,
		Reconciler: reconcilers.Sequence[*externalsecretsv1beta1.ExternalSecret]{
			ExternalSecretExternalSecretMappingChildReconciler(),
		},

		Config: c,
	}
}

//+kubebuilder:rbac:groups=x-mapping.servicebinding.io,resources=externalsecretmappings,verbs=get;list;watch;create;update;patch;delete

func ExternalSecretExternalSecretMappingChildReconciler() *reconcilers.ChildReconciler[*externalsecretsv1beta1.ExternalSecret, *mappingv1alpha1.ExternalSecretMapping, *mappingv1alpha1.ExternalSecretMappingList] {
	return &reconcilers.ChildReconciler[*externalsecretsv1beta1.ExternalSecret, *mappingv1alpha1.ExternalSecretMapping, *mappingv1alpha1.ExternalSecretMappingList]{
		DesiredChild: func(ctx context.Context, resource *externalsecretsv1beta1.ExternalSecret) (*mappingv1alpha1.ExternalSecretMapping, error) {
			child := &mappingv1alpha1.ExternalSecretMapping{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: resource.Namespace,
					Name:      resource.Name,
					Labels:    resource.Labels,
				},
			}

			return child, nil
		},
		MergeBeforeUpdate: func(current, desired *mappingv1alpha1.ExternalSecretMapping) {
			current.Labels = desired.Labels
			current.Spec = desired.Spec
		},
		ReflectChildStatusOnParent: func(parent *externalsecretsv1beta1.ExternalSecret, child *mappingv1alpha1.ExternalSecretMapping, err error) {
			// status updates are skipped, nothing to do
		},
	}
}
