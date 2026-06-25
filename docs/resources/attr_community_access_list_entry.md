---
page_title: "community_access_list_entry"
subcategory: ""
description: |-
  This represents an entry in a Community Access List, which is used to filter BGP routes based on BGP communities.
---

# community_access_list_entry (Nested-Attribute)

This represents an entry in a Community Access List, which is used to filter BGP routes based on BGP communities.




## Simple Attributes
- `action` (String) The action to be taken for this access list entry.
- `community` (String) The expression that defines the community access list entry, typically a BGP community.
- `rank` (Number) The rank of the access list entry.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
