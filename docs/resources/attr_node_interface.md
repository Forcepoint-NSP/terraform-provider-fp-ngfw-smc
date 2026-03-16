---
page_title: "node_interface"
subcategory: "interfaces"
description: |-
  This represents a Node Dedicated Interface, which is a unique IP address for each machine, not used for operative traffic in Firewall Clusters, IPS engines, and Layer 2 Firewalls.
---

# node_interface (Nested-Attribute)

This represents a Node Dedicated Interface, which is a unique IP address for each machine, not used for operative traffic in Firewall Clusters, IPS engines, and Layer 2 Firewalls.




## Simple Attributes
- `address` (String) The IP Address (IPv4 or IPv6) of the interface. For dynamic interfaces, this will be null.
- `apn` (String) The Access Point Name (APN) used for the modem connection. This is not required and can be used for modem configurations.
- `auth_request` (Boolean) Indicates whether this interface is used as the identity for authentication requests.
- `auth_request_source` (Boolean) Indicates whether this interface is used as the source for authentication requests. If not specified, the source IP address is selected automatically based on routing.
- `automatic_default_route` (Boolean) Indicates whether the dynamic default route will be automatically created for this dynamic interface.
- `backup_heartbeat` (Boolean) Indicates whether this interface is the backup heartbeat interface that is used if the primary heartbeat interface is unavailable.
- `backup_mgt` (Boolean) Indicates whether this interface is the backup control interface that is used if the primary control interface is not available.
- `comment` (String) An optional comment for the element. This field is not required.
- `domain_specific_dns_queries_source` (Boolean) Indicates whether this interface is used as the source for domain specific DNS queries. By default, the source for domain specific DNS queries is selected according to routing.
- `dynamic` (Boolean) Flag to mark or not this interface as dynamic, allowing the system to compute the correct dynamic index.
- `dynamic_index` (Number) The dynamic index of the IPv4 dynamic interface, which is a value between 1 and 16. For example, for 'First DHCP Interface ip', the index value is 1.
- `dynamic_ipv6_index` (Number) The dynamic index of the IPv6 dynamic interface, which is a value between 1 and 16. For example, for 'First DHCP Interface ip', the index value is 1.
- `igmp_mode` (String) The IGMP mode for this interface, which can be 'upstream' or 'downstream'. In upstream mode, the firewall acts as an IGMP querier for multicast servers and hosts in local networks. In downstream mode, the firewall queries downstream networks for hosts that want to join or leave multicast host groups.
- `igmp_querier_settings_ref` (String) This represents the IGMP Querier Settings for Multicast Routing and especially the PIM dynamic routing feature.
- `igmp_version` (String) The IGMP version used by this interface, which can be 1, 2, or 3. The default version is 3, but you may need to select another version for compatibility with certain hosts.
- `name` (String) Name of the object.
- `network_value` (String) The IP Network (IPv4 or IPv6) of the interface. For dynamic interfaces, this will be null.
- `nicid` (String) The Interface ID of the interface.
- `nodeid` (Number) The unique identifier of the engine node in the cluster, which is a value between 1 and 16.
- `outgoing` (Boolean) Indicates whether this interface is the default IP address for outgoing traffic. There must be one and only one NDI defined for outgoing traffic.
- `phone_number` (String) The phone number used for the modem connection. This is not required and can be used for modem configurations.
- `pin_code` (String) The PIN code used for the modem's SIM card. This is not required and can be used for modem configurations.
- `pppoa` (Boolean) Indicates whether this interface is configured for PPPoA (Point-to-Point Protocol over ATM).
- `pppoe` (Boolean) Indicates whether this interface is configured for PPPoE (Point-to-Point Protocol over Ethernet).
- `pppoe_password` (String) The PPPoE password used for authentication when the interface is configured for PPPoE.
- `pppoe_servicename` (String) The PPPoE service name used for the interface when configured for PPPoE.
- `pppoe_username` (String) The PPPoE username used for authentication when the interface is configured for PPPoE.
- `primary_heartbeat` (Boolean) Indicates whether this interface is the primary heartbeat interface for communications between the nodes. There must be one and only one primary heartbeat interface.
- `primary_mgt` (Boolean) Indicates whether this interface is the primary control interface for management server contact. There must be one and only one primary control interface.
- `reverse_connection` (Boolean) Indicates whether reverse connection is enabled for this interface. This is not required and can be used to allow reverse connections.
- `vrrp` (Boolean) Indicates whether this interface is configured for VRRP (Virtual Router Redundancy Protocol). This can be used with Physical Interfaces only.
- `vrrp_address` (String) The VRRP IP address used for the interface when configured for VRRP. This is required only for VRRP mode and must be a valid IPv4 address.
- `vrrp_id` (Number) The VRRP ID used for the interface when configured for VRRP. This is required only for VRRP mode and must be a valid integer.
- `vrrp_priority` (Number) The VRRP priority used for the interface when configured for VRRP. This is required only for VRRP mode and must be a valid integer.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
