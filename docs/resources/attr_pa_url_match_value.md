---
page_title: "url_entry"
subcategory: ""
description: |-
  This represents the Protocol Agent URL Match parameter value, which defines how URLs are matched in the Protocol Agent configuration.
---

# url_entry (Nested-Attribute)

This represents the Protocol Agent URL Match parameter value, which defines how URLs are matched in the Protocol Agent configuration.




## Simple Attributes
- `match_paramter` (String) The Match parameter, which can be 'host', 'path', or 'all', defining the specific part of the URL to match.
- `match_type` (String) The Match type, which can be 'contains', 'begins_with', or 'ends_with', defining how the URL is matched.
- `url_string` (String) The URL path to match against, specified as a string.


