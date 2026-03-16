---
page_title: "literal_constant_expr"
subcategory: ""
description: |-
  This represents a literal expression constant used in filter expressions, which contains a long value and a name for the constant. It is used to match against log field constants in filtering logic.
---

# literal_constant_expr (Nested-Attribute)

This represents a literal expression constant used in filter expressions, which contains a long value and a name for the constant. It is used to match against log field constants in filtering logic.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `constant_name` (String) The name of the log field constant, which is read-only and used for identification in filtering logic.
- `constant_value` (Number) The long value that represents the key of the log field constant used in filtering logic.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
