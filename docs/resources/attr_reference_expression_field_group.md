---
page_title: "field_group_expr"
subcategory: ""
description: |-
  This represents a reference to a field group in filter expressions, which contains a set of log fields. It is a system object and is read-only.
---

# field_group_expr (Nested-Attribute)

This represents a reference to a field group in filter expressions, which contains a set of log fields. It is a system object and is read-only.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `field_group_id` (Number) The ID of the field group being referenced.
- `log_field_refs` (List of String) A list of references to log fields contained in the referenced field group. This is read-only and derived from the field IDs.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
