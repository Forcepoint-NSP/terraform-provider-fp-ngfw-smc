---
page_title: "port_group_interface"
subcategory: "interfaces"
description: |-
  This represents a Port-Group Interface, which is used to group ports together on appliances that support Switch functionality. It is particularly relevant for N120 type switches ('Soft-Switch'), where only one Port-Group interface should be defined under the Switch.
---

# port_group_interface (Nested-Attribute)

This represents a Port-Group Interface, which is used to group ports together on appliances that support Switch functionality. It is particularly relevant for N120 type switches ('Soft-Switch'), where only one Port-Group interface should be defined under the Switch.




## Simple Attributes
- `aggregate_mode` (String) The aggregation type for the interface, which can be 'ha' for High Availability and represents two interfaces on the Firewall engine. Only the first interface in the aggregated link is actively used. The second interface becomes active only if the first interface fails. If you configure an Aggregated Link in High-Availability mode, connect the first interface to one switch and the second interface to another switch OR 'lb' for Load Balancing and Represents two interfaces on the Firewall engine. Both interfaces in the aggregated link are actively used and connections are automatically balanced between the two interfaces. Link aggregation in the load-balancing mode is implemented based on the IEEE 802.3ad Link Aggregation standard. If you configure an Aggregated Link in Load-Balancing Mode, connect both interfaces to a single switch. Make sure that the switch supports the Link Aggregation Control Protocol (LACP) and that LACP is configured on the switch.
- `comment` (String) An optional comment for the element. This field is not required.
- `custom_configuration` (String) Custom configuration for the physical interface, used for specific test or configuration purposes.
- `duplicate_address_detection` (Boolean) Indicates whether Duplicate Address Detection is enabled for IPv6, which helps prevent address conflicts on the network.
- `include_prefix_info_option_flag` (Boolean) Indicates whether the Include Prefix Information option is set in IPv6 Router Advertisements, which allows devices to automatically configure their IPv6 addresses based on the advertised prefixes.
- `interface_id` (String) The Interface ID automatically maps to a physical network port of the same number during the initial configuration of the engine, but the mapping can be changed as necessary through the engine's command line interface. In case of VLAN Physical Interface, enter the VLAN ID (1-4094). The VLAN IDs you add must be the same as the VLAN IDs that are used in the switch at the other end of the VLAN trunk. Each VLAN Interface is identified as Interface-ID.VLAN-ID, for example 2.100 for Interface ID 2 and VLAN ID 100.
- `mac_prefix` (String) The MAC Address prefix used for a shared interface, which is used to generate unique MAC addresses for interfaces in a shared configuration.
- `managed_address_flag` (Boolean) Indicates whether Managed Address Configuration is enabled in IPv6 Router Advertisements, which allows the Firewall to offer IPv6 addresses over DHCPv6.
- `mtu` (Number) The MTU (maximum transmission unit) size on the connected link. Either enter a value between 400-65535. The default value is 1500.
- `name` (String) Name of the object.
- `other_configuration_flag` (Boolean) Indicates whether Other Configuration is enabled in IPv6 Router Advertisements, which allows the Firewall to offer additional configuration information over DHCPv6.
- `override_engine_settings` (Boolean) Indicates whether the Engine's Default Settings are overridden for this interface, allowing for custom configurations.
- `override_log_moderation_settings` (Boolean) Indicates whether the Log Moderation settings are overridden for this interface, allowing for custom log moderation configurations.
- `qos_limit` (Number) The throughput limit for the link on this interface in kilobits per second (kbps). The same throughput is automatically applied to any VLANs created under this Physical Interface.
- `qos_mode` (String) Defines how QoS is applied to the link on this interface, such as 'none', 'dscp_handling', 'full_qos', or 'throttling'.
- `qos_policy_ref` (String) This represents a QoS Policy, which is used for Bandwidth Management and Traffic Prioritization based on QoS Classes or DSCP Matches.
- `router_advertisement` (Boolean) Indicates whether IPv6 Router Advertisements are sent, which provides configuration information to devices on the network.
- `second_interface_id` (String) The second interface ID in the aggregated link, which is used in conjunction with the first interface for high availability or load balancing.
- `set_autonomous_address_flag` (Boolean) Indicates whether the Set Autonomous Address-Configuration option is set in IPv6 Router Advertisements, which allows devices to automatically configure their IPv6 addresses based on the advertised prefixes.
- `shared_interface` (Boolean) Indicates whether this interface is shared across multiple engines, allowing for shared configurations and resources.
- `syn_max_bursts` (Number) The SYN Burst Size value, which is the number of allowed SYNs before the engine starts limiting the SYN rate. Must be at least 1.
- `syn_mode` (String) The SYN Rate Limits Mode, which can be 'off' for disabled, 'auto' for automatic calculation, or 'custom' for manual settings.
- `syn_per_second` (Number) The Allowed SYNs per Second value, which is the number of allowed SYN packets per second. Must be at least 1.
- `zone_ref` (String) This represents a Zone, which is used to group together network interfaces of Firewall, IPS, and Layer 2 Firewall engines. Zones can be used to specify receiving or sending interfaces in policies and automatically apply to new interfaces associated with the same Zone.

## Nested Attributes
- `arp_entry` (List of Blocks, see [here](attr_arp_entry.md)) the ARP entries associated with this physical interface.
- `dhcp_relay` (Single Block, see [here](attr_dhcp_relay.md)) 
- `dhcp_server_on_interface` (Single Block, see [here](attr_dhcp_server_settings.md)) 
- `dhcpv6_relay` (Single Block, see [here](attr_dhcp_relay.md)) 
- `interfaces` (List of Blocks, see [here](attr_engine_interface_wrapper.md)) the interfaces associated with this physical interface.
- `log_moderation` (List of Blocks, see [here](attr_log_moderation.md)) the log moderation settings associated with this physical interface.
- `switch_interface_port` (List of Blocks, see [here](attr_switch_interface_port.md)) A list of ports attached to this Port Group Interface. Each port is represented by a SwitchInterfacePortDTO.
- `virtual_resource_settings` (List of Blocks, see [here](attr_virtual_resource_settings.md)) the virtual resource settings associated with this physical interface.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
