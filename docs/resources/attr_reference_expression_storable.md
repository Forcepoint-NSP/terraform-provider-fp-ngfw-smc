---
page_title: "literal_ref_storable_expr"
subcategory: ""
description: |-
  This represents a reference expression that holds a storable reference, allowing for filtering logic based on storable entities.
---

# literal_ref_storable_expr (Nested-Attribute)

This represents a reference expression that holds a storable reference, allowing for filtering logic based on storable entities.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `storable_ref` (String) This is a base class for objects that can be stored in the database. It includes properties such as key, comment, and links.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
