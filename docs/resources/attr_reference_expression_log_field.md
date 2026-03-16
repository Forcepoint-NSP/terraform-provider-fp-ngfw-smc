---
page_title: "log_field_expr"
subcategory: ""
description: |-
  This represents a reference expression for a log field, allowing for filtering based on specific log field identifiers.
---

# log_field_expr (Nested-Attribute)

This represents a reference expression for a log field, allowing for filtering based on specific log field identifiers.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `log_field_ref` (String) The reference to the log field, represented as a URI that links to the specific log field identifier.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
