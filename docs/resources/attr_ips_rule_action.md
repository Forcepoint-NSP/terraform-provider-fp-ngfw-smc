---
page_title: "action"
subcategory: "engines"
description: |-
  This represents an IPS Rule Action, which is a sub part of IPS Ethernet/Access Rule used for the Rule action and for the Rule action options (block list settings/...).
---

# action (Nested-Attribute)

This represents an IPS Rule Action, which is a sub part of IPS Ethernet/Access Rule used for the Rule action and for the Rule action options (block list settings/...).




## Simple Attributes
- `action` (List of String) 
- `block_list_executor` (List of String) URI of the Block List Executor.
- `decrypting` (Boolean) Flag to indicate whether decryption is enabled for the traffic matching this rule.
- `deep_inspection` (Boolean) Flag to enable deep inspection of traffic that matches this rule. This will inspect the traffic against the Inspection Policy referenced by this policy.
- `discard_active` (Boolean) Indicates whether the discard action is currently active.
- `discard_override` (Boolean) Indicates whether the discard action overrides other actions.
- `discard_silent` (Boolean) Indicates whether the discard action is silent, meaning no response message is shown to the end-user.
- `dos_protection` (Boolean) Flag to enable or disable DoS protection for matching traffic. This will apply the DoS protection settings defined in the policy.
- `file_filtering` (Boolean) Flag to enable file filtering for matching traffic. This should also activate the Deep Inspection option. You can further adjust virus scanning in the Inspection Policy.
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
- `block_list_scope` (List of Blocks, see [here](attr_block_list_scope.md)) A block list scope, which define the entry that stops matching connections and block traffic between the matching IP addresses for the set duration.
- `connection_tracking_options` (Single Block, see [here](attr_connection_tracking_options.md)) 

