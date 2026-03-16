---
page_title: "smc_community_access_list"
subcategory: "routing"
description: |-
  This represents a Community Access List, which is used to define a list of communities for dynamic routing configurations.
---

# smc_community_access_list (Resource)

This represents a Community Access List, which is used to define a list of communities for dynamic routing configurations.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/engines/dynamic_routing/BGP/community_access_list) for an example

This example creates a Community Access List to specify accepted or denied BGP communities.

```hcl
resource "smc_community_access_list" "community_access_list" {
  comment = var.resource_comment
  entries {
    community_access_list_entry {
      action    = "deny"
      community = "178"
    }
  }
  entries {
    community_access_list_entry {
      action    = "permit"
      community = "1022"
    }
  }
  name = "tf_community_access_list"
  type = "standard"
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `type` (String) The type of the community access list.

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
