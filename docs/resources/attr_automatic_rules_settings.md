---
page_title: "automatic_rules_settings"
subcategory: "policy"
description: |-
  This represents the Automatic Rules Settings for the Engine Cluster, including VPN, QoS Class, Alert, and various traffic control settings.
---

# automatic_rules_settings (Nested-Attribute)

This represents the Automatic Rules Settings for the Engine Cluster, including VPN, QoS Class, Alert, and various traffic control settings.




## Simple Attributes
- `alert_ref` (String) This represents an abstract Alert, which is used to display messages when certain conditions are met.
- `allow_auth_traffic` (Boolean) Indicates whether authentication traffic is allowed for this Automatic Rules Settings.
- `allow_connections_to_dns_resolvers` (Boolean) Indicates whether connections to DNS resolvers are allowed for this Automatic Rules Settings.
- `allow_connections_to_remote_dhcp_server` (Boolean) Indicates whether connections to a remote DHCP server are allowed for this Automatic Rules Settings.
- `allow_icmp_traffic_for_route_probing` (Boolean) Indicates whether ICMP traffic is allowed for route probing in this Automatic Rules Settings.
- `allow_listening_interfaces_to_dns_relay_port` (Boolean) Indicates whether listening interfaces are allowed to relay DNS traffic for this Automatic Rules Settings.
- `allow_no_nat` (Boolean) Indicates whether NAT is allowed to be bypassed for this Automatic Rules Settings.
- `log_level` (String) The log level for the Automatic Rules Settings, which defines the verbosity of logging.
- `qos_class_ref` (String) This represents a QoS Class, which is an element that links a rule in a QoS Policy to one or more Firewall Actions. The traffic allowed in the access rule is assigned the QoS Class defined for the rule, and the QoS class is used as the matching criteria for applying QoS Policy rules.
- `vpn_ref` (String) This represents a Policy Based Virtual Private Network (VPN), which is used to establish secure connections over unsecured networks. It includes various configurations such as NAT rules, mobile VPN topology modes, and associated profiles.


