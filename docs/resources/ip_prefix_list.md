---
page_title: "smc_ip_prefix_list"
subcategory: "routing"
description: |-
  This represents an IP Prefix List, which is used to define a list of IP prefixes for dynamic routing configurations.
---

# smc_ip_prefix_list (Resource)

This represents an IP Prefix List, which is used to define a list of IP prefixes for dynamic routing configurations.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/1.741.0/examples/engines/dynamic_routing/ip_prefix_list) for a minimal example

This example creates an IP prefix list for BGP/OSPF filtering.

```hcl
resource "smc_ip_prefix_list" "ip_prefixl_list" {
  comment = var.resource_comment
  entries {
    ip_prefix_list_entry {
      action            = "permit"
      max_prefix_length = 28
      min_prefix_length = 25
      subnet            = "192.168.1.0/24"
    }
  }
  entries {
    ip_prefix_list_entry {
      action            = "deny"
      max_prefix_length = 26
      min_prefix_length = 25
      subnet            = "10.100.100.0/24"
    }
  }
  name = "tf_ip_prefixl_list"
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
