---
page_title: "smc_contact_address"
subcategory: "engines"
description: |-
  This represents a List of single contact address elements, which can be either dynamic or static IP addresses.
---

# smc_contact_address (Sub-resource)

This represents a List of single contact address elements, which can be either dynamic or static IP addresses.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/engines/single_fw/contact_addr) for a minimal example

This example assigns a contact address to an engine interface for external communications.

```hcl
resource "smc_contact_address" "tf_sample_contact_address" {
  id = data.smc_sub_href.contact_address_intf1.id
  contact_addresses {
    address      = "12.12.12.21"
    dynamic      = false
    location_ref = data.smc_href.default_location.id
  }
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource

## Nested Attributes
- `contact_addresses` (List of Blocks, see [here](attr_single_contact_address_element.md)) A list of single contact address elements, which can be either dynamic or static IP addresses.

