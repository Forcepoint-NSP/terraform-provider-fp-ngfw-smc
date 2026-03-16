---
page_title: "literal_ip_network_expr"
subcategory: ""
description: |-
  This represents a literal IP network used in filter expressions, which includes an IPv4 or IPv6 address with a netmask.
---

# literal_ip_network_expr (Nested-Attribute)

This represents a literal IP network used in filter expressions, which includes an IPv4 or IPv6 address with a netmask.




## Simple Attributes
- `comment` (String) An optional comment for the element. This field is not required.
- `ip_address_with_netmask` (String) The IPv4 or IPv6 address with netmask in the format 'a.b.c.d/n', where 'n' is the netmask as an integer.
- `name` (String) Name of the object.

## Nested Attributes
- `children` (List of Blocks, see [here](attr_filter_expression_wrapper.md)) The list of child filter expressions that define the filtering rules.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
