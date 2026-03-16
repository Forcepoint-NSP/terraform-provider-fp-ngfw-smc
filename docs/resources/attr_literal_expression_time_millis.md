---
page_title: "literal_time_millis_expr"
subcategory: ""
description: |-
  This represents a literal expression for time in milliseconds, used in filter expressions to specify time-based conditions.
---

# literal_time_millis_expr (Nested-Attribute)

This represents a literal expression for time in milliseconds, used in filter expressions to specify time-based conditions.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `long_value` (String) The time value in milliseconds used in the literal expression for filtering log fields.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
