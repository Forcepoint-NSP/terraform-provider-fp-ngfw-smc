---
page_title: "extended_community_access_list_entry"
subcategory: ""
description: |-
  This represents an entry in an Extended Community Access List, which is used to filter BGP routes based on extended communities.
---

# extended_community_access_list_entry (Nested-Attribute)

This represents an entry in an Extended Community Access List, which is used to filter BGP routes based on extended communities.




## Simple Attributes
- `action` (String) The action to be taken for this access list entry.
- `community` (String) The expression that defines the extended community access list entry.
- `rank` (Number) The rank of the access list entry.
- `type` (String) The type of the extended community access list entry, e.g., 'rt' for Route Target or 'soo' for Site of Origin.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
