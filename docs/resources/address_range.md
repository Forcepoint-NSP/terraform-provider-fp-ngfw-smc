---
page_title: "smc_address_range"
subcategory: "network_elements"
description: |-
  This represents an Address Range, which is a Network Element that defines a range of IP addresses. It includes attributes for the range of IP addresses, which must be valid IPv4 or IPv6 addresses.
---

# smc_address_range (Resource)

This represents an Address Range, which is a Network Element that defines a range of IP addresses. It includes attributes for the range of IP addresses, which must be valid IPv4 or IPv6 addresses.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/network_elements/address_range) for a complete minimal example

This example creates an Address Range object for defining an IP range in SMC.

```hcl
resource "smc_address_range" "tf_address_range_example" {
  ip_range = "10.20.30.40-10.20.30.50"
  name     = "tf_address_range_example"
  comment  = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `ip_range` (String) The range of IP addresses for the Address Range. Must be a valid IPv4 or IPv6 address range.
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
