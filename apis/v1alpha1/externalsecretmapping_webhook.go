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
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
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
func (r *ExternalSecretMapping) ValidateCreate() (admission.Warnings, error) {
	r.Default()
	return nil, r.validate().ToAggregate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ExternalSecretMapping) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	// TODO(user): check for immutable fields, if any
	r.Default()
	return nil, r.validate().ToAggregate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ExternalSecretMapping) ValidateDelete() (admission.Warnings, error) {
	return nil, nil
}

func (r *ExternalSecretMapping) validate() field.ErrorList {
	errs := field.ErrorList{}

	// require a single owner reference to an ExternalSecret
	if len(r.OwnerReferences) == 0 {
		errs = append(errs, field.Required(field.NewPath("metadata", "ownerReferences"), "must be owned by an ExternalSecret"))
	} else if len(r.OwnerReferences) != 1 {
		errs = append(errs, field.Invalid(field.NewPath("metadata", "ownerReferences"), r.OwnerReferences, "must be owned by exactly one ExternalSecret"))
	} else if o := r.OwnerReferences[0]; o.Kind != "ExternalSecret" || !strings.HasPrefix(o.APIVersion, "external-secrets.io/") {
		errs = append(errs, field.Invalid(field.NewPath("metadata", "ownerReferences").Index(0), r.OwnerReferences, "must be owned by an ExternalSecret"))
	} else if o.Name != r.Name {
		errs = append(errs, field.Invalid(field.NewPath("metadata", "ownerReferences").Index(0).Child("name"), o.Name, "owner name must match resource name"))
	}

	errs = append(errs, r.Spec.validate(field.NewPath("spec"))...)

	return errs
}

func (r *ExternalSecretMappingSpec) validate(fldPath *field.Path) field.ErrorList {
	errs := field.ErrorList{}

	return errs
}
