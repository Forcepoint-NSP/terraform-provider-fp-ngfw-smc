---
page_title: "action"
subcategory: "policy"
description: |-
  This represents a Firewall Rule Action, which includes options for forced next hop, explicit proxy jump, and other actions.
---

# action (Nested-Attribute)

This represents a Firewall Rule Action, which includes options for forced next hop, explicit proxy jump, and other actions.




## Simple Attributes
- `action` (List of String) 
- `decrypting` (Boolean) Flag to indicate whether decryption is enabled for the traffic matching this rule.
- `deep_inspection` (Boolean) Flag to enable deep inspection of traffic that matches this rule. This will inspect the traffic against the Inspection Policy referenced by this policy.
- `dos_protection` (Boolean) Flag to enable or disable DoS protection for matching traffic. This will apply the DoS protection settings defined in the policy.
- `explicit_proxy_sub_policy` (String) This represents a sub-policy that can be applied to various elements in the system, such as network elements, inspection rules, etc.
- `file_filtering` (Boolean) Flag to enable file filtering for matching traffic. This should also activate the Deep Inspection option. You can further adjust virus scanning in the Inspection Policy.
- `forced_next_hop_element` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `forced_next_hop_ip` (String) If 'forced_next_hop' action is selected, specify a valid IPv4 or IPv6 address for the forced next hop.
- `forward_to` (String) This represents an abstract node element in the network, which can be extended to represent specific types of nodes.
- `mobile_vpn` (Boolean) If 'apply_vpn', 'forward_vpn', or 'enforce_vpn' actions are selected, this indicates if it is an IPsec VPN client.
- `network_application_latency_monitoring` (String) Flag to enable Network Application Latency Monitoring. This will enable the Application Health Monitoring.
- `reset_icmp` (Boolean) If 'terminate' action is selected, this indicates whether to send an ICMP notification for non-TCP traffic termination.
- `scan_detection` (String) Enable or disable Scan Detection for traffic that matches the rule. This overrides the option set in the NGFW properties.
- `snort` (Boolean) Flag to indicate whether Snort intrusion detection is enabled for the traffic matching this rule.
- `sub_policy` (String) This represents a sub-policy that can be applied to various elements in the system, such as network elements, inspection rules, etc.
- `user_response` (String) This represents a User Response, which defines additional notification actions for rule matches, such as redirecting access to a forbidden URL to a page on an internal web server instead.
- `valid_block_lister` (List of String) URI of the network element used as blocklister.
- `vpn` (String) This represents a Policy Based Virtual Private Network (VPN), which is used to establish secure connections over unsecured networks. It includes various configurations such as NAT rules, mobile VPN topology modes, and associated profiles.

## Nested Attributes
- `connection_tracking_options` (Single Block, see [here](attr_connection_tracking_options.md)) 

