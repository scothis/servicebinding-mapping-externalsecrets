/*
 * Copyright 2020 Original Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ExternalSecretMappingSpec defines the desired state of ExternalSecretMapping
type ExternalSecretMappingSpec struct {
}

// ExternalSecretMappingStatus defines the observed state of ExternalSecretMapping
type ExternalSecretMappingStatus struct {
	// ObservedGeneration is the 'Generation' of the mapping that
	// was last processed by the controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions are the conditions of this mapping
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Binding exposes the projected secret for this mapping
	Binding *corev1.LocalObjectReference `json:"binding,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Secret",type=string,JSONPath=`.status.binding.name`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Reason",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].reason`
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// ExternalSecretMapping is the Schema for the mapping API
type ExternalSecretMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalSecretMappingSpec   `json:"spec,omitempty"`
	Status ExternalSecretMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalSecretMappingList contains a list of ExternalSecretMapping
type ExternalSecretMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ExternalSecretMapping `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExternalSecretMapping{}, &ExternalSecretMappingList{})
}
