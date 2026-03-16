---
page_title: "Forcepoint NGFW SMC - Terraform Provider"
description: Automate Forcepoint NGFW configuration deployment though SMC using terraform
---

# Forcepoint NGFW SMC - Terraform Provider

Automate Forcepoint NGFW configuration deployment though SMC using
terraform. 

## Table of content

- [Quick Start](guides/10-quick-start.md)
- [Creating Resources](guides/20-creating-resources.md)
- [Relationship Between Resources](guides/30-relationship-between-resources.md)
- [Representation of Nested Attributes](guides/40-nested-attributes.md)
- [Tips for Writing Terraform Configuration](guides/50-tips-writing-terraform-smc.md)
- [Importing Resources](guides/60-importing-resources.md)
- [Deleting Resources](guides/70-deleting-resources.md)
- [Troubleshooting](guides/80-troubleshooting.md)
- [Misc](guides/90-misc.md)


## Provider Configuration

### Example Usage

```hcl
terraform {
  required_providers {
    smc = {
      source  = "forcepoint-nsp/fp-ngfw-smc"
      version = "1.741.1"
    }
  }

  provider "smc" {
    url          = "https://mysmc:8082"
    api_key      = "xxxxxxxxxxxxxxxxx"
    api_version  = "7.4"
    trusted_cert = "./mycert.pem"
    verify_ssl   = true
  }
}
```

### Required Attributes

- `api_key` (String) The API key of the SMC API
- `url` (String) The URL of the SMC API

### Optional Attributes

- `api_version` (String) The SMC API version to use for request
- `trusted_cert` (String) PEM-encoded certificate content to trust for HTTPS connections or path to a file containing the PEM-encoded certificate.
- `verify_ssl` (Boolean) Whether to verify SSL certificates
- `domain` (String) The SMC domain to use for operations


## Provider version for SMC 7.3

SMC 7.3 is not officially supported, but terraform provider for SMC
7.4 works reasonbly well with it. 

Use latest fp-ngfw-smc terraform version 1.74x.y **setup with SMC API 7.3**

```
api_version = "7.3"
```

## Provider version for SMC 7.4

Use latest fp-ngfw-smc terraform version 1.74x.y

> default SMC API version used by latest provider 1.74x.y is SMC API 7.4.

