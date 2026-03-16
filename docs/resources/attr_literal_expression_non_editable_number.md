---
page_title: "literal_non_editable_number_expr"
subcategory: ""
description: |-
  This represents a literal expression that contains a non-editable number, used in filtering logic where a standard number literal is not applicable, such as when referencing deleted storable items.
---

# literal_non_editable_number_expr (Nested-Attribute)

This represents a literal expression that contains a non-editable number, used in filtering logic where a standard number literal is not applicable, such as when referencing deleted storable items.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `long_value` (Number) The long value representing the non-editable number used in filtering logic.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
