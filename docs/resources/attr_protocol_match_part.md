---
page_title: "protocols"
subcategory: ""
description: |-
  This represents the Inspection Protocol matching criteria for an IPS Policy Rule, which can specify whether any or none of the protocols match, or list specific elements that match.
---

# protocols (Nested-Attribute)

This represents the Inspection Protocol matching criteria for an IPS Policy Rule, which can specify whether any or none of the protocols match, or list specific elements that match.




## Simple Attributes
- `any` (Boolean) Indicates if any inspection protocol matches the criteria. If true, no elements are specified.
- `protocol` (List of String) URI of the protocol or match expression.


