---
page_title: "link_usage_exception"
subcategory: ""
description: |-
  This represents a Link Usage Exception, which associates a link type with preferences for using or avoiding certain links.
---

# link_usage_exception (Nested-Attribute)

This represents a Link Usage Exception, which associates a link type with preferences for using or avoiding certain links.




## Simple Attributes
- `link_type` (String) This represents the Link Type, which defines the characteristics and behavior of a link in the VPN configuration.

## Nested Attributes
- `link_usage_preference` (List of Blocks, see [here](attr_link_usage_preference.md)) The link usage preferences associated with this entry, which can include preferences for using or avoiding certain links.

