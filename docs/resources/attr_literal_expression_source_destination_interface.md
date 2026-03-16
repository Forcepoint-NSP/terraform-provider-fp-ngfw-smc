---
page_title: "literal_src_dst_interface_expr"
subcategory: ""
description: |-
  This represents a literal expression for source and destination interfaces, allowing for filtering based on specific interface identifiers in log entries.
---

# literal_src_dst_interface_expr (Nested-Attribute)

This represents a literal expression for source and destination interfaces, allowing for filtering based on specific interface identifiers in log entries.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `string_value` (String) The string value representing the source and destination interface, which can be in a specific format to match log entries.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
