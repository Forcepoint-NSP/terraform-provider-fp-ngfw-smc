---
page_title: "logical_true_expr"
subcategory: ""
description: |-
  This represents a logical expression that evaluates to true, allowing for filtering logic that affirms conditions or serves as a base case in logical operations.
---

# logical_true_expr (Nested-Attribute)

This represents a logical expression that evaluates to true, allowing for filtering logic that affirms conditions or serves as a base case in logical operations.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
