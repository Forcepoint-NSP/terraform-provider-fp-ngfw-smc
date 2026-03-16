---
page_title: "loopback_cluster_virtual_interface"
subcategory: "interfaces"
description: |-
  This represents the CVI Loopback IP address, which is used for loopback traffic sent to the whole cluster and shared by all nodes in the cluster.
---

# loopback_cluster_virtual_interface (Nested-Attribute)

This represents the CVI Loopback IP address, which is used for loopback traffic sent to the whole cluster and shared by all nodes in the cluster.




## Simple Attributes
- `address` (String) The IP Address (IPv4 or IPv6) of the interface. For dynamic interfaces, this will be null.
- `comment` (String) An optional comment for the element. This field is not required.
- `igmp_mode` (String) The IGMP mode for this interface, which can be 'upstream' or 'downstream'. In upstream mode, the firewall acts as an IGMP querier for multicast servers and hosts in local networks. In downstream mode, the firewall queries downstream networks for hosts that want to join or leave multicast host groups.
- `igmp_version` (String) The IGMP version used by this interface, which can be 1, 2, or 3. The default version is 3, but you may need to select another version for compatibility with certain hosts.
- `name` (String) Name of the object.
- `network_value` (String) The IP Network (IPv4 or IPv6) of the interface. For dynamic interfaces, this will be null.
- `nicid` (String) The Interface ID of the interface.
- `ospfv2_area_ref` (String) This represents an abstract OSPF Area used as Dynamic Routing element.
- `rank` (Number) The ordering rank of the Loopback Cluster Interface, which determines its priority in the cluster.
- `relayed_by_dhcp` (Boolean) Indicates whether this interface is relayed by DHCP. This option is used to specify if the interface is relayed by DHCP, which can be useful in certain network configurations.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
