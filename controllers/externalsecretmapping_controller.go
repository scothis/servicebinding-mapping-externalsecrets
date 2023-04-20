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
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/source"

	externalsecretsv1beta1 "github.com/servicebinding/mapping-externalsecrets/apis/thirdparty/externalsecrets/v1beta1"
	mappingv1alpha1 "github.com/servicebinding/mapping-externalsecrets/apis/v1alpha1"
)

//+kubebuilder:rbac:groups=x-mapping.servicebinding.io,resources=externalsecretmappings,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=x-mapping.servicebinding.io,resources=externalsecretmappings/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=x-mapping.servicebinding.io,resources=externalsecretmappings/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete

// ExternalSecretMappingReconciler reconciles a ExternalSecretMapping object
func ExternalSecretMappingReconciler(c reconcilers.Config) *reconcilers.ResourceReconciler[*mappingv1alpha1.ExternalSecretMapping] {
	return &reconcilers.ResourceReconciler[*mappingv1alpha1.ExternalSecretMapping]{
		Reconciler: reconcilers.Sequence[*mappingv1alpha1.ExternalSecretMapping]{
			ExternalSecretMappingSyncReconciler(),
		},

		Config: c,
	}
}

func ExternalSecretMappingSyncReconciler() *reconcilers.SyncReconciler[*mappingv1alpha1.ExternalSecretMapping] {
	return &reconcilers.SyncReconciler[*mappingv1alpha1.ExternalSecretMapping]{
		Sync: func(ctx context.Context, resource *mappingv1alpha1.ExternalSecretMapping) error {
			c := reconcilers.RetrieveConfigOrDie(ctx)

			controllerRef := metav1.GetControllerOf(resource)
			if controllerRef == nil || controllerRef.Kind != externalsecretsv1beta1.ExtSecretKind {
				// should never get here
				resource.GetConditionManager().MarkFalse(mappingv1alpha1.ExternalSecretMappingConditionServiceAvailable, "MissingOwner", "no ExternalSecret owns this mapping")
				return nil
			}
			nsn := types.NamespacedName{Namespace: resource.GetNamespace(), Name: controllerRef.Name}

			owner := &externalsecretsv1beta1.ExternalSecret{}
			if err := c.TrackAndGet(ctx, nsn, owner); err != nil {
				if apierrs.IsNotFound(err) {
					resource.GetConditionManager().MarkFalse(mappingv1alpha1.ExternalSecretMappingConditionServiceAvailable, "NotFound", "ExternalSecret %q was not found", nsn)
					return nil
				}
				return err
			}

			secretName := owner.Name
			if name := owner.Spec.Target.Name; name != "" {
				secretName = name
			}

			// TODO consider condition of ExternalSecret before propagating secret
			resource.Status.Binding = &corev1.LocalObjectReference{
				Name: secretName,
			}
			resource.GetConditionManager().MarkTrue(mappingv1alpha1.ExternalSecretMappingConditionServiceAvailable, "Available", "")

			return nil
		},

		Setup: func(ctx context.Context, mgr controllerruntime.Manager, bldr *builder.Builder) error {
			bldr.Watches(&source.Kind{Type: &externalsecretsv1beta1.ExternalSecret{}}, reconcilers.EnqueueTracked(ctx, &externalsecretsv1beta1.ExternalSecret{}))

			return nil
		},
	}
}
