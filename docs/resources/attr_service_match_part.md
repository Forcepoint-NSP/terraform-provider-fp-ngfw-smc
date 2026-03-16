---
page_title: "services"
subcategory: ""
description: |-
  This represents the Service matching criteria for a Policy Rule, which can be 'any', 'none', or a list of specific service entries.
---

# services (Nested-Attribute)

This represents the Service matching criteria for a Policy Rule, which can be 'any', 'none', or a list of specific service entries.




## Simple Attributes
- `any` (Boolean) Indicates if any service matches the criteria. If true, no specific elements are required.
- `details` (List of String) 
- `none` (Boolean) Indicates if no service matches the criteria. If true, no specific elements are required.
- `service` (List of String) URI of the service or match expression.


