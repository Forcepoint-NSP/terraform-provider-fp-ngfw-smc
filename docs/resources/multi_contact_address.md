---
page_title: "smc_multi_contact_mvia"
subcategory: "engines"
description: |-
  This represents a list of multi contact address elements, which are used to manage multiple contact addresses in the system.
---

# smc_multi_contact_mvia (Sub-resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a list of multi contact address elements, which are used to manage multiple contact addresses in the system.

## Examples

No usage found in the current example set.

A stub configuration:

```hcl
# resource "smc_multi_contact_mvia" "example" {
#   # Example attributes here
# }
```

> This resource type does not appear in any official example yet.


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource

## Nested Attributes
- `multi_contact_addresses` (List of Blocks, see [here](attr_multi_contact_address_element.md)) A list of multi contact address elements, which are used to manage multiple contact addresses in the system.

