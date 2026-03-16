---
page_title: "literal_service_expr"
subcategory: ""
description: |-
  This represents a literal expression that matches service fields defined by the ServiceKeyFactory, allowing for complex filtering logic based on service attributes.
---

# literal_service_expr (Nested-Attribute)

This represents a literal expression that matches service fields defined by the ServiceKeyFactory, allowing for complex filtering logic based on service attributes.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `string_value` (String) The string value that represents the service field to match against.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
