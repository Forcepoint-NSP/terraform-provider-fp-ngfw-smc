---
page_title: "in_expr"
subcategory: ""
description: |-
  This represents a comparison expression that checks if a value is contained within a specified set of values, allowing for flexible filtering logic based on inclusion in a list.
---

# in_expr (Nested-Attribute)

This represents a comparison expression that checks if a value is contained within a specified set of values, allowing for flexible filtering logic based on inclusion in a list.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
