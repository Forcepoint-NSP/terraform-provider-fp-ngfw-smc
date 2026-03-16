# SMC Terraform Provider - fp-ngfw-smc

- Website: https://registry.terraform.io/providers/forcepoint-nsp/fp-ngfw-smc
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

## Which provider version to use ?

Please refer to the official documentation:
https://registry.terraform.io/providers/forcepoint-nsp/fp-ngfw-smc/latest/docs

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.14.x +
- [Go](https://golang.org/doc/install) 1.25.x

## Building the Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using "make":

## Using the Provider

## Provider management and versioning
This provider is internally maintained by Forcepoint R&D.
The public repository is updated by Forcepoint when:

- a new SMC version is released (only if required), or
- a new provider version is published.

You can still build the provider locally using the following command:
```shell
make build
```

## OpenAPI model and version mapping
The provider implementation is generated from the SMC API OpenAPI specification.
The exact SMC API model version used is reflected in the provider version number:

For provider version M.XYZ.P, the OpenAPI model comes from SMC version X.Y.Z.

Example:
- Provider version 1.741.0
- Uses the OpenAPI model from SMC version 7.4.1

> Starting with SMC 7.4.1, the SMC API OpenAPI model is included in the SMC package.
