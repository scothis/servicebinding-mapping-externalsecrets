/*
Copyright 2022 the original author or authors.

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
	externalsecretsv1beta1 "github.com/servicebinding/mapping-externalsecrets/apis/thirdparty/externalsecrets/v1beta1"
)

// +die:object=true
type _ = externalsecretsv1beta1.ExternalSecret

// +die
type _ = externalsecretsv1beta1.ExternalSecretSpec

// +die
type _ = externalsecretsv1beta1.ExternalSecretStatus
