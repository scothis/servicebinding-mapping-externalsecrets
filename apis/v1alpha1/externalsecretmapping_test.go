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
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/pointer"
)

func TestExternalSecretMappingDefault(t *testing.T) {
	tests := []struct {
		name     string
		seed     *ExternalSecretMapping
		expected *ExternalSecretMapping
	}{
		{
			name:     "none",
			seed:     &ExternalSecretMapping{},
			expected: &ExternalSecretMapping{},
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			actual := c.seed.DeepCopy()
			actual.Default()
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("(-expected, +actual): %s", diff)
			}
		})
	}
}

func TestExternalSecretMappingValidate(t *testing.T) {
	tests := []struct {
		name     string
		seed     *ExternalSecretMapping
		expected field.ErrorList
	}{
		{
			name: "empty is not valid",
			seed: &ExternalSecretMapping{},
			expected: field.ErrorList{
				field.Required(field.NewPath("metadata", "ownerReferences"), "must be owned by an ExternalSecret"),
			},
		},
		{
			name: "valid",
			seed: &ExternalSecretMapping{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-secret",
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion:         "external-secrets.io/v1beta1",
							Kind:               "ExternalSecret",
							Name:               "my-secret",
							UID:                types.UID("37e9ef98-6946-4252-a345-7efdf387ea3c"),
							Controller:         pointer.Bool(true),
							BlockOwnerDeletion: pointer.Bool(true),
						},
					},
				},
			},
			expected: field.ErrorList{},
		},
		{
			name: "wrong owner kind",
			seed: &ExternalSecretMapping{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-secret",
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion:         "external-secrets.io/v1beta1",
							Kind:               "SecretStore",
							Name:               "my-secret",
							UID:                types.UID("37e9ef98-6946-4252-a345-7efdf387ea3c"),
							Controller:         pointer.Bool(true),
							BlockOwnerDeletion: pointer.Bool(true),
						},
					},
				},
			},
			expected: field.ErrorList{
				field.Invalid(field.NewPath("metadata", "ownerReferences").Index(0), []metav1.OwnerReference{
					{
						APIVersion:         "external-secrets.io/v1beta1",
						Kind:               "SecretStore",
						Name:               "my-secret",
						UID:                types.UID("37e9ef98-6946-4252-a345-7efdf387ea3c"),
						Controller:         pointer.Bool(true),
						BlockOwnerDeletion: pointer.Bool(true),
					},
				}, "must be owned by an ExternalSecret"),
			},
		},
		{
			name: "too many owner refs",
			seed: &ExternalSecretMapping{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-secret",
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion:         "external-secrets.io/v1beta1",
							Kind:               "ExternalSecret",
							Name:               "my-secret",
							UID:                types.UID("37e9ef98-6946-4252-a345-7efdf387ea3c"),
							Controller:         pointer.Bool(true),
							BlockOwnerDeletion: pointer.Bool(true),
						},
						{
							APIVersion:         "external-secrets.io/v1beta1",
							Kind:               "ExternalSecret",
							Name:               "my-secret",
							UID:                types.UID("37e9ef98-6946-4252-a345-7efdf387ea3c"),
							Controller:         pointer.Bool(true),
							BlockOwnerDeletion: pointer.Bool(true),
						},
					},
				},
			},
			expected: field.ErrorList{
				field.Invalid(field.NewPath("metadata", "ownerReferences"), []metav1.OwnerReference{
					{
						APIVersion:         "external-secrets.io/v1beta1",
						Kind:               "ExternalSecret",
						Name:               "my-secret",
						UID:                types.UID("37e9ef98-6946-4252-a345-7efdf387ea3c"),
						Controller:         pointer.Bool(true),
						BlockOwnerDeletion: pointer.Bool(true),
					},
					{
						APIVersion:         "external-secrets.io/v1beta1",
						Kind:               "ExternalSecret",
						Name:               "my-secret",
						UID:                types.UID("37e9ef98-6946-4252-a345-7efdf387ea3c"),
						Controller:         pointer.Bool(true),
						BlockOwnerDeletion: pointer.Bool(true),
					},
				}, "must be owned by exactly one ExternalSecret"),
			},
		},
		{
			name: "owner by resource with different name",
			seed: &ExternalSecretMapping{
				ObjectMeta: metav1.ObjectMeta{
					Name: "my-secret",
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion:         "external-secrets.io/v1beta1",
							Kind:               "ExternalSecret",
							Name:               "some-other-secret",
							UID:                types.UID("37e9ef98-6946-4252-a345-7efdf387ea3c"),
							Controller:         pointer.Bool(true),
							BlockOwnerDeletion: pointer.Bool(true),
						},
					},
				},
			},
			expected: field.ErrorList{
				field.Invalid(field.NewPath("metadata", "ownerReferences").Index(0).Child("name"), "some-other-secret", "owner name must match resource name"),
			},
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if diff := cmp.Diff(c.expected, c.seed.validate()); diff != "" {
				t.Errorf("validate (-expected, +actual): %s", diff)
			}

			expectedErr := c.expected.ToAggregate()
			if diff := cmp.Diff(expectedErr, c.seed.ValidateCreate()); diff != "" {
				t.Errorf("ValidateCreate (-expected, +actual): %s", diff)
			}
			if diff := cmp.Diff(expectedErr, c.seed.ValidateUpdate(c.seed.DeepCopy())); diff != "" {
				t.Errorf("ValidateCreate (-expected, +actual): %s", diff)
			}
			if diff := cmp.Diff(nil, c.seed.ValidateDelete()); diff != "" {
				t.Errorf("ValidateDelete (-expected, +actual): %s", diff)
			}
		})
	}
}
