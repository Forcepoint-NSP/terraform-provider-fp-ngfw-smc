---
page_title: "filter_expr"
subcategory: ""
description: |-
  This represents a reference to a filter expression in the system, allowing for complex filtering logic by referencing existing filters.
---

# filter_expr (Nested-Attribute)

This represents a reference to a filter expression in the system, allowing for complex filtering logic by referencing existing filters.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `filter_ref` (String) This represents a Filter Expression Container, which is a top-level re-usable filter that contains filter expression nodes.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
