---
page_title: "data_tag_expr"
subcategory: ""
description: |-
  This represents a reference expression for data tags, used in filter expressions to match log entries based on specific data tag values.
---

# data_tag_expr (Nested-Attribute)

This represents a reference expression for data tags, used in filter expressions to match log entries based on specific data tag values.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `data_tag_values` (String) A comma-delimited list of data tag values used in the reference expression for filtering log fields.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
