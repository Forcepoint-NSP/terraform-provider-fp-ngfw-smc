---
page_title: "smc_ipv6_prefix_list"
subcategory: "routing"
description: |-
  This represents an IPv6 Prefix List, which is used to define a list of IPv6 prefixes for dynamic routing configurations.
---

# smc_ipv6_prefix_list (Resource)

This represents an IPv6 Prefix List, which is used to define a list of IPv6 prefixes for dynamic routing configurations.

## Examples

- [ipv6_prefix_list/main.tf](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/dynamic_routing/ipv6_prefix_list/main.tf): Example of an IPv6 prefix list for BGP or route filtering.

This snippet demonstrates using `smc_ipv6_prefix_list` with multiple entry constraints for BGP prefix-lists.

```hcl
resource "smc_ipv6_prefix_list" "ipv6_prefix_list" {
  comment = var.resource_comment
  entries {
    ipv6_prefix_list_entry {
      action            = "deny"
      max_prefix_length = 78
      min_prefix_length = 65
      subnet            = "2001:4860:4860::/64"
    }
  }
  name = "tf_ipv6_prefix_list"
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.

## Nested Attributes
- `entries` (List of Blocks, see [here](attr_abstract_access_list_entry_wrapper.md)) The access list entries that define the rules for routing decisions.

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
