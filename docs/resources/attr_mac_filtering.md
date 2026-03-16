---
page_title: "mac_filtering"
subcategory: ""
description: |-
  This represents the definition of Mac Filtering for SSID interface. MAC filtering allows you to restrict which clients can connect to the SSID Interface based on the clients' MAC address.
---

# mac_filtering (Nested-Attribute)

This represents the definition of Mac Filtering for SSID interface. MAC filtering allows you to restrict which clients can connect to the SSID Interface based on the clients' MAC address.




## Simple Attributes
- `mac_address_ref` (List of String) The list of MAC addresses that is either denied or allowed based on the selected mode. This is required when the mode is 'accept_unless_in_deny_list' or 'deny_unless_in_accept_list'.
- `mode` (String) The mode of MAC filtering, which can be 'disabled', 'accept_unless_in_deny_list', or 'deny_unless_in_accept_list'. 'disabled' means no filtering is applied, 'accept_unless_in_deny_list' means clients are allowed unless they are in the deny list, and 'deny_unless_in_accept_list' means only clients in the accept list are allowed.


