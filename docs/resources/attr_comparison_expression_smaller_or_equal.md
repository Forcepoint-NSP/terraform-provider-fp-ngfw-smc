---
page_title: "smaller_or_equal_expr"
subcategory: ""
description: |-
  This represents a comparison expression that checks if a value is smaller than or equal to another value, allowing for precise filtering logic based on numerical comparisons.
---

# smaller_or_equal_expr (Nested-Attribute)

This represents a comparison expression that checks if a value is smaller than or equal to another value, allowing for precise filtering logic based on numerical comparisons.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
