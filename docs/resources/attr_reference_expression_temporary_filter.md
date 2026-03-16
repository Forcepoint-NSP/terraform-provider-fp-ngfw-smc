---
page_title: "ref_temporary_filter_expr"
subcategory: ""
description: |-
  This represents a reference to a temporary filter expression, used in local filtering logic to apply specific conditions based on temporary filter definitions.
---

# ref_temporary_filter_expr (Nested-Attribute)

This represents a reference to a temporary filter expression, used in local filtering logic to apply specific conditions based on temporary filter definitions.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
