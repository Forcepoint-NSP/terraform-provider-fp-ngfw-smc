---
page_title: "smc_extended_community_access_list"
subcategory: ""
description: |-
  This represents an Extended Community Access List, which is used to define a list of extended communities for dynamic routing configurations.
---

# smc_extended_community_access_list (Resource)

This represents an Extended Community Access List, which is used to define a list of extended communities for dynamic routing configurations.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/dynamic_routing/BGP/extended_community_access_list) for a minimal example

This example defines an extended BGP community access list.

```hcl
resource "smc_extended_community_access_list" "extended_community_access_list" {
  entries {
    extended_community_access_list_entry {
      action = "permit"
      community = ".*1545"
    }
  }
  name = "tf_extended_community_access_list"
  type = "expanded"
  comment = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `type` (String) The type of the extended community access list.

## Nested Attributes
- `entries` (List of Blocks, see [here](attr_abstract_access_list_entry_wrapper.md)) A list of entries in the Extended Community Access List.

## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
