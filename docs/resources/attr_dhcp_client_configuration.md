---
page_title: "dhcp_client_config"
subcategory: ""
description: |-
  This represents the DHCP Client Configuration element for VPN Clients, which can specify DHCP mode, servers, interfaces, and address pools.
---

# dhcp_client_config (Nested-Attribute)

This represents the DHCP Client Configuration element for VPN Clients, which can specify DHCP mode, servers, interfaces, and address pools.




## Simple Attributes
- `dhcp_add_info` (Number) Indicates the mode for adding information to DHCP in the VPN Client configuration. 0 for None, 1 for User, 2 for Group.
- `dhcp_client_interface` (List of String) The interface IP addresses used for DHCP Relay in the VPN Client configuration. This is applicable when DHCP Relay mode is selected.
- `dhcp_client_mode` (Number) The DHCP mode used for the VPN Client configuration. 0 for Disabled, 1 for Direct, 2 for Relay.
- `dhcp_server_ref` (List of String) URI of the DHCP Server.
- `proxy_arp_address_list` (String) If the useProxyARP flag is false, don't use this field, there is no valuable information in it. Take in account that if the flag is set to true, the value can come from a value entered in the SMC (or API) or an automatically generated IP Address range.
- `restricted_address_enabled` (Boolean) Indicates if the virtual address ranges are restricted. If false, the address pool is not used.
- `restricted_address_list` (String) If the 'restricted_address_enabled' flag is false, don't use this field, there is no valuable information in it. Take in account that if the flag is set to true, the value can come from a value entered in the SMC (or API) or an automatically generated IP Address range.
- `use_arp_proxy_enabled` (Boolean) Indicates if Proxy ARP is used in the VPN Client configuration.


