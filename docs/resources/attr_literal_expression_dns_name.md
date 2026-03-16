---
page_title: "literal_dns_name_expr"
subcategory: ""
description: |-
  This represents a literal expression for a DNS name, used in filter expressions to specify conditions based on domain names.
---

# literal_dns_name_expr (Nested-Attribute)

This represents a literal expression for a DNS name, used in filter expressions to specify conditions based on domain names.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `string_value` (String) The DNS name string value used in the literal expression for filtering log fields.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
