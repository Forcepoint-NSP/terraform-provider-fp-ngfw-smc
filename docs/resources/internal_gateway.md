---
page_title: "smc_internal_gateway"
subcategory: "vpn"
description: |-
  This represents the Internal Gateway, which is used for managing VPN connections and related settings.
---

# smc_internal_gateway (Sub-resource)

This represents the Internal Gateway, which is used for managing VPN connections and related settings.

## Examples

<!-- - [two_gateway_in_one_engine_disabled/main.tf](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/sdwan/policy_based_vpn/two_gateway_in_one_engine_disabled/main.tf): Configures an internal gateway as part of a firewall cluster's SD-WAN setup. -->

This example demonstrates how to configure an `smc_internal_gateway` resource for enabling internal gateway features in a firewall engine within an SD-WAN topology. It shows typical attributes including VPN client mode and SSL/TLS settings for remote access.

```hcl
resource "smc_internal_gateway" "tf_single_fw1_2" {
  from_ref                 = smc_single_fw.tf_single_fw1.link.internal_gateway
  antivirus                = false
  auto_certificate         = true
  auto_site_content        = true
  cluster_ref              = smc_single_fw.tf_single_fw1.id
  dhcp_relay {
    dhcp_add_info              = 0
    dhcp_client_mode           = 0
    proxy_arp_address_list     = ""
    restricted_address_enabled = false
    use_arp_proxy_enabled      = false
  }
  dtls                    = false
  firewall                = false
  name                    = "tf_single_fw1-bis"
  ssl_vpn_proxy {
    renegociation_timeout = 7200
    ssl_3_0              = false
    tls_1_0              = false
    tls_1_1              = false
    tls_1_2              = true
    tls_1_3              = false
  }
  ssl_vpn_tunneling {
    renegociation_timeout = 7200
    ssl_3_0              = false
    tls_1_0              = false
    tls_1_1              = false
    tls_1_2              = true
    tls_1_3              = false
  }
  trust_all_cas              = true
  vpn_client_mode            = "ipsec"
  windows_update             = false
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `antivirus` (Boolean) Indicates whether the AntiVirus is enabled for this Internal Gateway.
- `auto_certificate` (Boolean) Indicates whether the Internal Gateway automatically generates and manages RSA certificates.
- `auto_site_content` (Boolean) Indicates whether the site content is automatically generated from the routing view. This is applicable only for Internal Gateways.
- `cluster_ref` (String) This represents a group of devices, or nodes, that share a given work load. You can cluster Firewalls to share the load and provide redundancy, allowing, for example, scheduled maintenance that takes one node out of service without interrupting services to the users.
- `comment` (String) An optional comment for the element. This field is not required.
- `dtls` (Boolean) 
- `end_point` (List of String) URI of the internal end-point.
- `firewall` (Boolean) Indicates whether the Firewall is enabled for this Internal Gateway.
- `gateway_profile` (String) Gateway Profiles describe the capabilities of a Gateway, i.e. supported cipher, hash, etc. Gateway Profiles of Internal Gateways are read-only and computed from Firewall version and FIPS mode. Gateway Profiles of External Gateways are user-defined.
- `name` (String) Name of the object.
- `site` (List of String) URI of the site.
- `trust_all_cas` (Boolean) Indicates whether the EndPoint trusts all VPN Certificate Authorities. If true, it trusts all CAs; if false, it requires specific trusted CAs.
- `trusted_certificate_authorities` (List of String) URI of the trusted VPN Certificate Authority.
- `vpn_client_mode` (String) The VPN Client Mode for this Internal Gateway, which can be 'no', 'ipsec', 'ssl', or 'both'.
- `windows_update` (Boolean) Indicates whether Windows Update is enabled for this Internal Gateway.

## Nested Attributes
- `dhcp_relay` (Single Block, see [here](attr_dhcp_client_configuration.md)) 
- `ssl_vpn_portal_setting` (List of Blocks, see [here](attr_ssl_vpn_portal_setting.md)) The Application Access Portal Settings for this Internal Gateway.
- `ssl_vpn_proxy` (Single Block, see [here](attr_ssl_vpn_setting.md)) 
- `ssl_vpn_tunneling` (Single Block, see [here](attr_ssl_vpn_setting.md)) 

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
