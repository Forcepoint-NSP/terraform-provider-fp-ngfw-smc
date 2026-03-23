---
page_title: "smc_ipv6_access_list"
subcategory: "policy"
description: |-
  This represents an IPv6 Access List, which is used to define a list of IPv6 addresses and prefixes for dynamic routing configurations.
---

# smc_ipv6_access_list (Resource)

This represents an IPv6 Access List, which is used to define a list of IPv6 addresses and prefixes for dynamic routing configurations.

## Examples

- [ipv6_access_list/main.tf](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/1.741.0/examples/engines/dynamic_routing/ipv6_access_list/main.tf): Lists IPv6 permit/deny entries for access lists.

This snippet demonstrates how to configure IPv6 access lists, using both permit and deny entries.

```hcl
resource "smc_ipv6_access_list" "ipv6_access_element" {
  comment = var.resource_comment
  entries {
    ipv6_access_list_entry {
      action = "deny"
      subnet = "2001:db8:85a3::8a2e:370:7334/128"
    }
  }
  entries {
    ipv6_access_list_entry {
      action = "permit"
      subnet = "2606:2800:220:1:248:1893:25c8:1946/64"
    }
  }
  name = "ipv6_access_element"
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
