---
page_title: "smc_lldp_profile"
subcategory: ""
description: |-
  This represents a LLDP Profile, which is used to configure the LLDP (Link Layer Discovery Protocol) settings for devices.
---

# smc_lldp_profile (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a LLDP Profile, which is used to configure the LLDP (Link Layer Discovery Protocol) settings for devices.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `hold_time_multiplier` (Number) The multiplier for the hold time in LLDP advertisements. This value is used to calculate the hold time based on the transmit delay. Default is 4.
- `name` (String) Name of the object.
- `transmit_delay` (Number) The delay in milliseconds before transmitting LLDP advertisement frames. Default is 30 ms.


## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
