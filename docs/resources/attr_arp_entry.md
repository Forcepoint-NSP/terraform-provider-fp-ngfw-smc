---
page_title: "arp_entry"
subcategory: ""
description: |-
  This represents the definition of an Address Resolution Protocol (ARP) entry, which associates IP addresses with MAC addresses on a local area network (LAN).
---

# arp_entry (Nested-Attribute)

This represents the definition of an Address Resolution Protocol (ARP) entry, which associates IP addresses with MAC addresses on a local area network (LAN).




## Simple Attributes
- `ipaddress` (String) The IP address associated with the ARP entry, which can be either IPv4 or IPv6.
- `macaddress` (String) The MAC address associated with the ARP entry. It is not required but can be provided for static entries.
- `netmask` (Number) The netmask for the ARP entry, which defines the network portion of the IP address. It can be a value between 0 and 128 for IPv6.
- `type` (String) The type of ARP entry, which can be 'static' for permanent references or 'dynamic' for proxy ARP. 'static': it gives the engine a permanent reference to an IP address/MAC address pair. All entries are of this type on IPS engines, Layer 2 Firewalls, Master Engines, and Virtual NGFW Engines. 'dynamic': it gives a Firewall engine a reference to an IP address/MAC address pair that the Firewall should perform proxy ARP for. Proxy ARP is possible only for hosts located in networks directly connected to the Firewall.


