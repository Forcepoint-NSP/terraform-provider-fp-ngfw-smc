---
page_title: "default_expr"
subcategory: ""
description: |-
  This represents a default comparison expression that can be used in filter expressions, allowing for flexible filtering logic based on default comparison criteria.
---

# default_expr (Nested-Attribute)

This represents a default comparison expression that can be used in filter expressions, allowing for flexible filtering logic based on default comparison criteria.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
