---
page_title: "smc_gateway_settings"
subcategory: "vpn"
description: |-
  This represents various gateway and VPN parameters to be set at Firewall level. These settings are used for the negotiation of VPN connections.
---

# smc_gateway_settings (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents various gateway and VPN parameters to be set at Firewall level. These settings are used for the negotiation of VPN connections.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `certificate_cache_crl_validity` (Number) Time in seconds for which the certificate cache is valid.
- `comment` (String) An optional comment for the element. This field is not required.
- `mobike_after_sa_update` (Boolean) Whether the After SA flag is set for Mobike Policy.
- `mobike_before_sa_update` (Boolean) Whether the Before SA flag is set for Mobike Policy.
- `mobike_no_rrc` (Boolean) Whether the No RRC flag is set for Mobike Policy.
- `name` (String) Name of the object.
- `negotiation_expiration` (Number) Time in seconds after which the negotiation of VPN connections expires.
- `negotiation_retry_max_number` (Number) Maximum number of retries for the negotiation of VPN connections.
- `negotiation_retry_timer` (Number) Time in seconds to wait before retrying the negotiation of VPN connections.
- `negotiation_retry_timer_max` (Number) Maximum time in seconds to wait before retrying the negotiation of VPN connections.


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
