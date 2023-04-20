/*
Copyright 2021 the original author or authors.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *ExternalSecretMapping) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

var _ webhook.Defaulter = &ExternalSecretMapping{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *ExternalSecretMapping) Default() {
}

//+kubebuilder:webhook:path=/validate-x-mapping-servicebinding-io-v1alpha1-externalsecretmapping,mutating=false,failurePolicy=fail,sideEffects=None,groups=x-mapping.servicebinding.io,resources=externalsecretmappings,verbs=create;update,versions=v1alpha1,name=veso.xmapping.servicebinding.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &ExternalSecretMapping{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ExternalSecretMapping) ValidateCreate() error {
	r.Default()
	return r.validate().ToAggregate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ExternalSecretMapping) ValidateUpdate(old runtime.Object) error {
	// TODO(user): check for immutable fields, if any
	r.Default()
	return r.validate().ToAggregate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ExternalSecretMapping) ValidateDelete() error {
	return nil
}

func (r *ExternalSecretMapping) validate() field.ErrorList {
	errs := field.ErrorList{}

	// TODO require an owner reference to an ExternalSecret

	errs = append(errs, r.Spec.validate(field.NewPath("spec"))...)

	return errs
}

func (r *ExternalSecretMappingSpec) validate(fldPath *field.Path) field.ErrorList {
	errs := field.ErrorList{}

	return errs
}
