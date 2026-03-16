---
page_title: "smc_vpn_broker_domain"
subcategory: "vpn"
description: |-
  This represents a VPN Broker Domain. It is used to configure the broker domain settings, including the configuration file and MAC address prefix.
---

# smc_vpn_broker_domain (Resource)

This represents a VPN Broker Domain. It is used to configure the broker domain settings, including the configuration file and MAC address prefix.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `file_ref` (String) This is a configuration file for the VPN Broker Domain. It contains settings and configurations specific to the broker domain, used for managing VPN connections and policies.
- `link_usage_profile` (String) This represents a Link Usage Profile, which is used to manage link usage settings in a multi-link environment. It includes configurations for link balancing, packet duplication, and forward erasure correction.
- `mac_address_prefix` (String) First 3 octets of the Broker domain Mac Address.
- `name` (String) Name of the object.


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
