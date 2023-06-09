//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by diegen. DO NOT EDIT.

package v1beta1

import (
	testingx "testing"

	testing "dies.dev/testing"
)

func TestExternalSecretDie_MissingMethods(t *testingx.T) {
	die := ExternalSecretBlank
	ignore := []string{"TypeMeta", "ObjectMeta"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ExternalSecretDie: %s", diff.List())
	}
}

func TestExternalSecretSpecDie_MissingMethods(t *testingx.T) {
	die := ExternalSecretSpecBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ExternalSecretSpecDie: %s", diff.List())
	}
}

func TestExternalSecretStatusDie_MissingMethods(t *testingx.T) {
	die := ExternalSecretStatusBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ExternalSecretStatusDie: %s", diff.List())
	}
}
