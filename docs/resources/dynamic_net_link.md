---
page_title: "smc_dynamic_netlink"
subcategory: "routing"
description: |-
  This represents a Dynamic NetLink, which is a type of NetLink used for routing in Multi-Link features. It includes attributes for automatically learning DNS servers.
---

# smc_dynamic_netlink (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Dynamic NetLink, which is a type of NetLink used for routing in Multi-Link features. It includes attributes for automatically learning DNS servers.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `active_mode_period` (Number) The active period in milliseconds for the NetLink, which defines how often the link is probed when it is in Active or Standby mode. Leave the setting for Standby Mode as 0 if you prefer not to test this link when it is on standby.
- `active_mode_timeout` (Number) The active timeout in milliseconds for the NetLink, which defines how long the firewall waits before it considers the probe failed. Change the setting for Standby Mode to 0 if you prefer not to test this link when it is on standby.
- `comment` (String) An optional comment for the element. This field is not required.
- `connection_type_ref` (String) This represents a connection type used in endpoints, which defines the connectivity group, mode, and link type for VPN connections.
- `input_speed` (Number) The Input Speed in bits per second for the NetLink, which defines the real-life bandwidth this network connection provides. It is used to calculate how much traffic each link receives in relation to the other links.
- `learn_dns_server_automatically` (Boolean) Indicates whether the Dynamic NetLink automatically learns DNS servers. If true, the NetLink will learn DNS servers dynamically.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `output_speed` (Number) The Output Speed in bits per second for the NetLink, which defines the real-life bandwidth this network connection provides. It is used to calculate how much traffic each link receives in relation to the other links.
- `probe_address` (List of String) The IP address that is probed with ICMP echo requests (ping) to determine if the link is up. Repeat this for each IP address you want to add. We recommend entering more than one address to avoid excluding the link in case the host that is probed goes down.
- `standby_mode_period` (Number) The standby period in seconds for the NetLink, which defines how often the link is probed when it is in Standby mode. Leave the setting for Standby Mode as 0 if you prefer not to test this link when it is on standby.
- `standby_mode_timeout` (Number) The standby timeout in seconds for the NetLink, which defines how long the firewall waits before it considers the probe failed in Standby mode. Change the setting for Standby Mode to 0 if you prefer not to test this link when it is on standby.


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
