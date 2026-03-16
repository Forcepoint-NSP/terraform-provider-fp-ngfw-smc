---
page_title: "as_path_access_list_entry"
subcategory: ""
description: |-
  This represents an entry in an AS Path Access List, which is used to filter BGP routes based on their AS path.
---

# as_path_access_list_entry (Nested-Attribute)

This represents an entry in an AS Path Access List, which is used to filter BGP routes based on their AS path.




## Simple Attributes
- `action` (String) The action to be taken for this access list entry.
- `comment` (String) An optional comment for the element. This field is not required.
- `expression` (String) The expression that defines the AS path access list entry.
- `name` (String) Name of the object.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
