---
page_title: "literal_mac_address_expr"
subcategory: ""
description: |-
  This represents a literal expression that holds a MAC address value, allowing for filtering logic based on network addresses.
---

# literal_mac_address_expr (Nested-Attribute)

This represents a literal expression that holds a MAC address value, allowing for filtering logic based on network addresses.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `string_value` (String) The MAC address value represented by this literal expression, formatted as x:x:x:x:x:x.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
