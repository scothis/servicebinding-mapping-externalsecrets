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

package v1beta1

import (
	diecorev1 "dies.dev/apis/core/v1"
	diemetav1 "dies.dev/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mappingv1alpha1 "github.com/servicebinding/mapping-externalsecrets/apis/v1alpha1"
)

// +die:object=true
type _ = mappingv1alpha1.ExternalSecretMapping

// +die
type _ = mappingv1alpha1.ExternalSecretMappingSpec

// +die
type _ = mappingv1alpha1.ExternalSecretMappingStatus

func (d *ExternalSecretMappingStatusDie) ConditionsDie(conditions ...*diemetav1.ConditionDie) *ExternalSecretMappingStatusDie {
	return d.DieStamp(func(r *mappingv1alpha1.ExternalSecretMappingStatus) {
		r.Conditions = make([]metav1.Condition, len(conditions))
		for i := range conditions {
			r.Conditions[i] = conditions[i].DieRelease()
		}
	})
}

var ExternalSecretMappingConditionReady = diemetav1.ConditionBlank.Type(mappingv1alpha1.ExternalSecretMappingConditionReady).Unknown().Reason("Initializing")
var ExternalSecretMappingConditionServiceAvailable = diemetav1.ConditionBlank.Type(mappingv1alpha1.ExternalSecretMappingConditionServiceAvailable).Unknown().Reason("Initializing")

func (d *ExternalSecretMappingStatusDie) BindingDie(fn func(d *diecorev1.LocalObjectReferenceDie)) *ExternalSecretMappingStatusDie {
	return d.DieStamp(func(r *mappingv1alpha1.ExternalSecretMappingStatus) {
		d := diecorev1.LocalObjectReferenceBlank.DieImmutable(false).DieFeedPtr(r.Binding)
		fn(d)
		r.Binding = d.DieReleasePtr()
	})
}
