---
page_title: "smc_ip_access_list"
subcategory: "routing"
description: |-
  This represents an IP Access List, which is used to define a list of IP addresses and prefixes for dynamic routing configurations.
---

# smc_ip_access_list (Resource)

This represents an IP Access List, which is used to define a list of IP addresses and prefixes for dynamic routing configurations.

## Examples

- [IP Access List Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/dynamic_routing/routing_node_bgp/main.tf)

An `smc_ip_access_list` allows you to define rules for matching and filtering IP subnets for dynamic routing filtering purposes.

```hcl
resource "smc_ip_access_list" "ip_access_list" {
  comment = var.resource_comment
  entries {
    ip_access_list_entry {
      action = "deny"
      subnet = "172.16.16.0/21"
    }
  }
  entries {
    ip_access_list_entry {
      action = "permit"
      subnet = "192.168.100.0/24"
    }
  }
  name = "tf_ip_access_list"
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
