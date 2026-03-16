---
page_title: "destinations"
subcategory: ""
description: |-
  This represents the Destination matching criteria for a Policy Rule, allowing you to specify whether any or none of the destination criteria match, or to list specific elements that match.
---

# destinations (Nested-Attribute)

This represents the Destination matching criteria for a Policy Rule, allowing you to specify whether any or none of the destination criteria match, or to list specific elements that match.




## Simple Attributes
- `any` (Boolean) Indicates if any destination matches the criteria. If true, no specific elements are required.
- `dst` (List of String) URI of the destination entry.
- `network_details` (List of String) 
- `none` (Boolean) Indicates if no destination matches the criteria. If true, no specific elements are required.


