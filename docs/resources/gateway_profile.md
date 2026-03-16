---
page_title: "smc_gateway_profile"
subcategory: "vpn"
description: |-
  Gateway Profiles describe the capabilities of a Gateway, i.e. supported cipher, hash, etc. Gateway Profiles of Internal Gateways are read-only and computed from Firewall version and FIPS mode. Gateway Profiles of External Gateways are user-defined.
---

# smc_gateway_profile (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
Gateway Profiles describe the capabilities of a Gateway, i.e. supported cipher, hash, etc. Gateway Profiles of Internal Gateways are read-only and computed from Firewall version and FIPS mode. Gateway Profiles of External Gateways are user-defined.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.

## Nested Attributes
- `capabilities` (Single Block, see [here](attr_capabilities.md)) 

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
