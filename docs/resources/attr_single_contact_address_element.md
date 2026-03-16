---
page_title: "contact_addresses"
subcategory: ""
description: |-
  This represents a single contact address element, which links a location to an IP address. It includes fields for the location reference, address, and dynamic status.
---

# contact_addresses (Nested-Attribute)

This represents a single contact address element, which links a location to an IP address. It includes fields for the location reference, address, and dynamic status.




## Simple Attributes
- `address` (String) The address associated with this contact address element, typically an IP address.
- `dynamic` (Boolean) Indicates whether the contact address element is dynamic, meaning it can change over time.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.


