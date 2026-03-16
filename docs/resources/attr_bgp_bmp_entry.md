---
page_title: "bmp_entry"
subcategory: "routing"
description: |-
  This represents the BGP BMP entry, which includes the BMP port, BMP address, and whether the connection is through the master.
---

# bmp_entry (Nested-Attribute)

This represents the BGP BMP entry, which includes the BMP port, BMP address, and whether the connection is through the master.




## Simple Attributes
- `bmp_address` (String) The BMP address value, which can be a FQDN, IPv4, or IPv6 address.
- `bmp_connect_through_master` (Boolean) Indicates whether the BMP connection is connected through the master.
- `bmp_port` (Number) The BMP port value.
- `subnet` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.


