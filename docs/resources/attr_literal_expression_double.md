---
page_title: "literal_double_expr"
subcategory: ""
description: |-
  This represents a literal expression for a double value, used in filter expressions to specify conditions based on numerical data.
---

# literal_double_expr (Nested-Attribute)

This represents a literal expression for a double value, used in filter expressions to specify conditions based on numerical data.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `double_value` (String) The double value used in the literal expression for filtering log fields.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
