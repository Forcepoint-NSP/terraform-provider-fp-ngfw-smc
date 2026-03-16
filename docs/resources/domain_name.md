---
page_title: "smc_domain_name"
subcategory: "network_elements"
description: |-
  This represents a Domain Name, which is used to manage domain names and their associated IP addresses. It includes additional domain names that can be resolved to IP addresses through DNS queries.
---

# smc_domain_name (Resource)

This represents a Domain Name, which is used to manage domain names and their associated IP addresses. It includes additional domain names that can be resolved to IP addresses through DNS queries.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/network_elements/domain_name) for a complete minimal example

This example creates a Domain Name object recognized within SMC.

```hcl
resource "smc_domain_name" "tfdomain_name" {
  domain_name_entry = ["test.example.com", "prod.example.com", "example.com"]
  name              = "tfdomain_name"
  comment           = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `domain_name_entry` (List of String) A list of additional domain names that can be resolved to IP addresses through DNS queries.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.


## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
