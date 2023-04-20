<img src="https://avatars.githubusercontent.com/u/64096231" align="right" />

# Service Binding Mapping for External Secrets <!-- omit in toc -->

[![CI](https://github.com/servicebinding/mapping-externalsecrets/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/servicebinding/mapping-externalsecrets/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/servicebinding/mapping-externalsecrets)](https://goreportcard.com/report/github.com/servicebinding/mapping-externalsecrets)
[![Go Reference](https://pkg.go.dev/badge/github.com/servicebinding/mapping-externalsecrets.svg)](https://pkg.go.dev/github.com/servicebinding/mapping-externalsecrets)
[![codecov](https://codecov.io/gh/servicebinding/mapping-externalsecrets/branch/main/graph/badge.svg?token=D2Hs4MIXBZ)](https://codecov.io/gh/servicebinding/mapping-externalsecrets)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Reference implementation of the [ServiceBinding.io](https://servicebinding.io) [1.0 spec](https://servicebinding.io/spec/core/1.0.0/). The full specification is implemented, please open an issue for any discrepancies.

- [Getting Started](#getting-started)
  - [Deploy a released build](#deploy-a-released-build)
  - [Build from source](#build-from-source)
    - [Undeploy controller](#undeploy-controller)
- [Architecture](#architecture)
- [Contributing](#contributing)
  - [Test It Out](#test-it-out)
  - [Modifying the API definitions](#modifying-the-api-definitions)
- [Community, discussion, contribution, and support](#community-discussion-contribution-and-support)
  - [Code of conduct](#code-of-conduct)

## Getting Started

You’ll need a Kubernetes cluster to run against. You can use [kind](https://kind.sigs.k8s.io) to get a local cluster for testing, or run against a remote cluster.

After the controller is deployed, try out the [samples](#samples).

### Deploy a released build

The easiest way to get started is by deploying the [latest release](https://github.com/servicebinding/mapping-externalsecrets/releases). Alternatively, you can [build the runtime from source](#build-from-source).

### Build from source

1. Define where to publish images:

   ```sh
   export KO_DOCKER_REPO=<a-repository-you-can-write-to>
   ```

   For kind, a registry is not required:

   ```sh
   export KO_DOCKER_REPO=kind.local
   ```
	
1. Build and deploy the controller to the cluster:

   Note: The cluster must have the [cert-manager](https://cert-manager.io) and [external-secrets](https://external-secrets.io) deployed.  There is a `make deploy-cert-manager` and `make deploy-external-secrets` target to deploy the cert-manager and external-secrets respectively.

   ```sh
   make deploy
   ```

#### Undeploy controller
Undeploy the controller to the cluster:

```sh
make undeploy
```

## Architecture

A `ExternalSecretMapping` mirrors every `ExternalSecret` resource in the cluster. The `ExternalSecretMapping` is a [Service Binding Provisioned Service][provisioned-service] compatible resource that reflects the name of the secret defined by the `ExternalSecret` on to the `ExternalSecretMapping` status. Users should not create the `ExternalSecretMapping` resource directly, as new `ExternalSecret` are created/updated/deleted, the `ExternalSecretMapping` with the same namespace/name is created/updated/deleted. The mapping does not alter the `Secret` or `ExternalSecret` resources in anyway.

Once installed, a `ServiceBinding` can target an `ExternalSecretMapping` of the same name as the `ExternalSecret` as a service.

In the `ServiceBinding` replace:

```yaml
spec:
  service:
    apiVersion: external-secrets.io/v1beta1
    kind: ExternalSecret
    name: my-secret
```

with:

```yaml
spec:
  service:
    apiVersion: x-mapping.servicebinding.io/v1alpha1
    kind: ExternalSecretMapping
    name: my-secret
```

## Contributing

### Test It Out

Run the unit tests:

```sh
make test
```

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## Community, discussion, contribution, and support

The Service Binding Mapping for External Secrets project is a community lead effort.
A bi-weekly [working group call][working-group] is open to the public.
Discussions occur here on GitHub and on the [#bindings-discuss channel in the Kubernetes Slack][slack].

If you catch an error in the implementation, please let us know by opening an issue at our
[GitHub repository][repo].

### Code of conduct

Participation in the Service Binding community is governed by the [Contributor Covenant][code-of-conduct].

[working-group]: https://docs.google.com/document/d/1rR0qLpsjU38nRXxeich7F5QUy73RHJ90hnZiFIQ-JJ8/edit#heading=h.ar8ibc31ux6f
[slack]: https://kubernetes.slack.com/archives/C012F2GPMTQ
[repo]: https://github.com/servicebinding/mapping-externalsecrets
[code-of-conduct]: ./CODE_OF_CONDUCT.md
[provisioned-service]: https://servicebinding.io/spec/core/1.0.0/#provisioned-service
