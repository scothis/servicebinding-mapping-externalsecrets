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

// Package v1alpha1 contains API Schema definitions for the x-mapping.servicebinding.io v1alpha1 API group
// +kubebuilder:object:generate=true
// +groupName=x-mapping.servicebinding.io
package v1alpha1

import (
	"github.com/vmware-labs/reconciler-runtime/apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These are valid conditions of ExternalSecretMapping.
const (
	ExternalSecretMappingConditionReady            = apis.ConditionReady
	ExternalSecretMappingConditionServiceAvailable = "ServiceAvailable"
)

var servicebindingCondSet = apis.NewLivingConditionSetWithHappyReason(
	"ServiceBound",
	ExternalSecretMappingConditionServiceAvailable,
)

func (s *ExternalSecretMapping) GetConditionsAccessor() apis.ConditionsAccessor {
	return &s.Status
}

func (s *ExternalSecretMapping) GetConditionSet() apis.ConditionSet {
	return servicebindingCondSet
}

func (s *ExternalSecretMapping) GetConditionManager() apis.ConditionManager {
	return servicebindingCondSet.Manage(&s.Status)
}

func (s *ExternalSecretMappingStatus) InitializeConditions() {
	conditionManager := servicebindingCondSet.Manage(s)
	conditionManager.InitializeConditions()
	// reset existing managed conditions
	conditionManager.MarkUnknown(ExternalSecretMappingConditionServiceAvailable, "Initializing", "")
}

var _ apis.ConditionsAccessor = (*ExternalSecretMappingStatus)(nil)

// GetConditions implements ConditionsAccessor
func (s *ExternalSecretMappingStatus) GetConditions() []metav1.Condition {
	return s.Conditions
}

// SetConditions implements ConditionsAccessor
func (s *ExternalSecretMappingStatus) SetConditions(c []metav1.Condition) {
	s.Conditions = c
}

// GetCondition fetches the condition of the specified type.
func (s *ExternalSecretMappingStatus) GetCondition(t string) *metav1.Condition {
	for _, cond := range s.Conditions {
		if cond.Type == t {
			return &cond
		}
	}
	return nil
}
