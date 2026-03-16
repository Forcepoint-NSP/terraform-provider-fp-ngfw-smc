---
page_title: "literal_long_expr"
subcategory: ""
description: |-
  This represents a literal expression that holds a long value, allowing for filtering logic based on numerical values.
---

# literal_long_expr (Nested-Attribute)

This represents a literal expression that holds a long value, allowing for filtering logic based on numerical values.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `long_value` (String) The long value represented by this literal expression.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
