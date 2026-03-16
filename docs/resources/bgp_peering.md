---
page_title: "smc_bgp_peering"
subcategory: "routing"
description: |-
  This represents the BGP Peering for Dynamic Routing Firewall functionality.
---

# smc_bgp_peering (Resource)

This represents the BGP Peering for Dynamic Routing Firewall functionality.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/engines/dynamic_routing/BGP/bgp_peering) for an example

This example configures a BGP peering with filters and connection profile.

```hcl
resource "smc_bgp_peering" "bgp_peering" {
  bfd_enabled               = false
  connected_check           = "disabled"
  connection_profile        = smc_bgp_connection_profile.bgp_connection_profile.id
  default_originate         = false
  dont_capability_negotiate = false
  inbound_aspath_filter     = smc_as_path_access_list.as_path_access_list.id
  inbound_ip_filter         = smc_ip_access_list.ip_access_list.id
  local_as_option           = "not_set"
  max_prefix_option         = "not_enabled"
  name                      = "tf_bgp_peering"
  next_hop_self             = true
  orf_option                = "disabled"
  outbound_aspath_filter    = smc_as_path_access_list.as_path_access_list.id
  outbound_ip_filter        = smc_ip_access_list.ip_access_list.id
  override_capability       = false
  remove_private_as         = false
  route_reflector_client    = true
  send_community            = "no"
  soft_reconfiguration      = true
  ttl_option                = "disabled"
  comment                   = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `bfd_enabled` (Boolean) Indicates whether Bidirectional Forwarding Detection (BFD) is enabled for the BGP Peering. If true, BFD will be used to monitor the BGP peering.
- `bfd_interval` (Number) The Bidirectional Forwarding Detection (BFD) interval in milliseconds for the BGP Peering. This is used to set the interval for BFD control packets.
- `bfd_min_rx` (Number) The Bidirectional Forwarding Detection (BFD) minimum receive interval in milliseconds for the BGP Peering. This is used to set the minimum interval for receiving BFD control packets.
- `bfd_multiplier` (Number) The Bidirectional Forwarding Detection (BFD) multiplier for the BGP Peering. This is used to set the multiplier for BFD control packets.
- `bfd_passive_mode` (Boolean) Indicates whether Bidirectional Forwarding Detection (BFD) is in passive mode for the BGP Peering. If true, BFD will operate in passive mode, meaning it will not initiate BFD sessions.
- `comment` (String) An optional comment for the element. This field is not required.
- `connected_check` (String) The Connected Check mode for the BGP Peering. This determines how connected checks are performed in BGP operations.
- `connection_profile` (String) This represents the BGP Connection Profile for Dynamic Routing Firewall functionality. It contains settings for BGP connections, such as session timers and password.
- `default_originate` (Boolean) Indicates whether the BGP Peering should originate default routes. If true, the peering will originate default routes in the BGP network.
- `dont_capability_negotiate` (Boolean) Indicates whether capabilities should not be sent in BGP updates. If true, capabilities will not be sent in the BGP updates.
- `inbound_aspath_filter` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `inbound_ip_filter` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `inbound_ipprefix_filter` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `inbound_ipv6_filter` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `inbound_ipv6prefix_filter` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `inbound_rm_filter` (String) This represents a Route Map Policy for the Dynamic Routing Firewall settings, which is used to control the routing behavior based on specific rules.
- `local_as_option` (String) The Local AS mode for the BGP Peering. This determines how the Local AS number is handled in BGP operations.
- `local_as_value` (String) The Local AS value for the BGP Peering. This is used when the Local AS mode is set to prepend or replace_as.
- `max_prefix_option` (String) The Max Prefix mode for the BGP Peering. This determines how the maximum number of prefixes is handled in BGP operations.
- `max_prefix_value` (Number) The Max Prefix value for the BGP Peering. This is used when the Max Prefix mode is set to enabled or warning_only.
- `md5_password` (String) The MD5 password for the BGP Peering. This is used for authentication of the BGP connection.
- `name` (String) Name of the object.
- `next_hop_self` (Boolean) Indicates whether the Next Hop Self feature is enabled for the BGP Peering. If true, the peering will set itself as the next hop for routes.
- `orf_option` (String) The Outbound Route Filtering (ORF) mode for the BGP Peering. This determines how outbound route filtering is handled in BGP operations.
- `outbound_aspath_filter` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `outbound_ip_filter` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `outbound_ipprefix_filter` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `outbound_ipv6_filter` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `outbound_ipv6prefix_filter` (String) This represents an abstract access list used in dynamic routing. It contains a list of access list entries that define the rules for routing decisions.
- `outbound_rm_filter` (String) This represents a Route Map Policy for the Dynamic Routing Firewall settings, which is used to control the routing behavior based on specific rules.
- `override_capability` (Boolean) Indicates whether received capabilities should be overridden. If true, received capabilities will be overridden in the BGP peering.
- `remove_private_as` (Boolean) Indicates whether private AS numbers should be removed from BGP updates. If true, private AS numbers will be removed from the BGP updates.
- `route_reflector_client` (Boolean) Indicates whether the BGP Peering is a Route Reflector Client. If true, the peering will act as a Route Reflector Client in the BGP network.
- `send_community` (String) The Send Community mode for the BGP Peering. This determines how community attributes are sent in BGP operations.
- `soft_reconfiguration` (Boolean) Indicates whether soft reconfiguration is enabled for inbound BGP updates. If true, soft reconfiguration will be used to handle inbound updates.
- `ttl_option` (String) The TTL Check mode for the BGP Peering. This determines how TTL checks are performed in BGP operations.
- `ttl_value` (Number) The TTL value for the BGP Peering. This is used when the TTL Check mode is set to ttl-security or ebgp-multihop.


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
