---
page_title: "LiteralExprSearch"
subcategory: ""
description: |-
  This represents a literal expression search used in filter expressions, which contains a string value for searching within log fields. It is intended for use in filtering logic based on string values.
---

# LiteralExprSearch (Nested-Attribute)

This represents a literal expression search used in filter expressions, which contains a string value for searching within log fields. It is intended for use in filtering logic based on string values.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `string_value` (String) The string value used in the literal expression search for filtering log fields.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
