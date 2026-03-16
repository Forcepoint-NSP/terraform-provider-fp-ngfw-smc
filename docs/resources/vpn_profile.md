---
page_title: "smc_vpn_profile"
subcategory: "vpn"
description: |-
  This represents a VPN Profile. It contains settings for IKE and IPsec lifetimes, keep-alive options, certificate authorities, and authentication methods.
---

# smc_vpn_profile (Resource)

This represents a VPN Profile. It contains settings for IKE and IPsec lifetimes, keep-alive options, certificate authorities, and authentication methods.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `cn_authentication_for_mobile_vpn` (Boolean) Indicates whether CN authentication is allowed or not, for a client using a certificate authentication.
- `comment` (String) An optional comment for the element. This field is not required.
- `disable_anti_replay` (Boolean) Indicates whether anti-replay protection is disabled for IPsec.
- `disable_path_discovery` (Boolean) Indicates whether Path MTU Discovery is disabled for IPsec.
- `hybrid_authentication_for_mobile_vpn` (Boolean) Indicates whether Hybrid authentication is allowed or not, for a client using a certificate authentication.
- `ike_v2_ppk` (String) IKEv2 Post-quantum Preshared Key option for encryption. Possible values: not_in_use | enabled | mandatory.
- `keep_alive` (Boolean) Whether to enable keep-alive for the VPN profile.
- `name` (String) Name of the object.
- `preshared_key_authentication_for_mobile_vpn` (Boolean) Indicates whether preshared key authentication is allowed or not, together with a client certificate authentication.
- `sa_life_time` (Number) IKE Lifetime in seconds.
- `sa_to_any_network_allowed` (Boolean) Indicates whether IPsec Security Associations (SAs) to any network are allowed.
- `trust_all_cas` (Boolean) Indicates whether to trust all certificate authorities (CAs) or not. If set to true, all CAs are trusted without validation.
- `trusted_certificate_authority` (List of String) URI of the VPN Certificate Authority.
- `tunnel_life_time_kbytes` (Number) IPsec Lifetime in KBytes. This is the maximum amount of data that can be sent before the IPsec SA is rekeyed.
- `tunnel_life_time_seconds` (Number) IPsec Lifetime in seconds.

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
