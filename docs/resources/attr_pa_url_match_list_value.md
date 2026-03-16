---
page_title: "pa_urlmatchlist"
subcategory: ""
description: |-
  This represents the Protocol Agent URL Match List parameter value, which contains a list of URL match rules.
---

# pa_urlmatchlist (Nested-Attribute)

This represents the Protocol Agent URL Match List parameter value, which contains a list of URL match rules.




## Simple Attributes
- `parameter_ref` (String) This represents a parameter for the Protocol Agent, allowing for detailed configuration of agent settings.

## Nested Attributes
- `url_entry` (List of Blocks, see [here](attr_pa_url_match_value.md)) The list of URL match rules used in the Protocol Agent Parameter.

