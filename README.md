# Provider AzureExtraUpjet

https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md
    
```shell
export PATH="$PATH:$HOME/go/bin" && make generate

up xpkg push xpkg.upbound.io/chelala/provider-azureextraupjet:v0.0.2 -f linux_arm64/provider-azureextraupjet-v0.0.0-3.g73e7ffe.xpkg,linux_amd64/provider-azureextraupjet-v0.0.0-3.g73e7ffe.xpkg

kubectl -n crossplane-system logs -f provider-azureextraupjet-40512a196139-dbc9dccd5-96dpc
```

`provider-azureextraupjet` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
AzureExtraUpjet API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/chelala/provider-azureextraupjet):
```
up ctp provider install chelala/provider-azureextraupjet:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-azureextraupjet
spec:
  package: chelala/provider-azureextraupjet:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/chelala/provider-azureextraupjet).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/chelala/provider-azureextraupjet/issues).
