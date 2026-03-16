---
page_title: "smc_alias"
subcategory: "network_elements"
description: |-
  This represents an Alias, which is used to represent other network elements in configurations. It allows for flexible configuration by defining default values and alias values that can be used in policies.
---

# smc_alias (Resource)

This represents an Alias, which is used to represent other network elements in configurations. It allows for flexible configuration by defining default values and alias values that can be used in policies.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/network_elements/alias) for a complete minimal example

This example creates an Alias object in SMC to reference other network elements by a different name.

```hcl
resource "smc_alias" "tf_alias_test" {
  name    = "$ tf test alias"
  comment = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `default_alias_value` (List of String) URI of the default value Storable.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.

## Nested Attributes
- `alias_value` (List of Blocks, see [here](attr_alias_value.md)) A list of alias values that define specific configurations for the Alias. Each alias value can specify different configurations for the Alias in policies.

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
