---
page_title: "action"
subcategory: "policy"
description: |-
  This represents a File Filtering Rule Action, which includes options for rematch archive content, file reputation, antivirus, sandboxing, and ICAP DLP scanning.
---

# action (Nested-Attribute)

This represents a File Filtering Rule Action, which includes options for rematch archive content, file reputation, antivirus, sandboxing, and ICAP DLP scanning.




## Simple Attributes
- `action` (List of String) 
- `antivirus` (Boolean) Indicates whether antivirus scanning is enabled.
- `decrypting` (Boolean) Flag to indicate whether decryption is enabled for the traffic matching this rule.
- `deep_inspection` (Boolean) Flag to enable deep inspection of traffic that matches this rule. This will inspect the traffic against the Inspection Policy referenced by this policy.
- `default_behavior` (String) The default behavior for file filtering actions.
- `dirty_log_level` (String) The log level for dirty files, which determines how they are logged.
- `dos_protection` (Boolean) Flag to enable or disable DoS protection for matching traffic. This will apply the DoS protection settings defined in the policy.
- `file_filtering` (Boolean) Flag to enable file filtering for matching traffic. This should also activate the Deep Inspection option. You can further adjust virus scanning in the Inspection Policy.
- `file_reputation` (Boolean) Indicates whether file reputation checking is enabled.
- `file_reputation_allow_level` (String) The level of file reputation that is allowed.
- `file_reputation_discard_level` (String) The level of file reputation that results in discarding the file.
- `forward_to` (String) This represents an abstract node element in the network, which can be extended to represent specific types of nodes.
- `icap_dlp_file_size_exceeded_action` (String) The action to take if the ICAP DLP file size exceeds the maximum limit.
- `icap_dlp_max_file_size` (Number) The maximum file size for ICAP DLP scanning in MB.
- `icap_dlp_scan_enabled` (Boolean) Indicates whether ICAP DLP scanning is enabled.
- `icap_dlp_service_fail_action` (String) The action to take if the ICAP DLP service fails.
- `mobile_vpn` (Boolean) If 'apply_vpn', 'forward_vpn', or 'enforce_vpn' actions are selected, this indicates if it is an IPsec VPN client.
- `network_application_latency_monitoring` (String) Flag to enable Network Application Latency Monitoring. This will enable the Application Health Monitoring.
- `rematch_archive_content` (Boolean) Indicates whether rematching of archive content is enabled.
- `reset_icmp` (Boolean) If 'terminate' action is selected, this indicates whether to send an ICMP notification for non-TCP traffic termination.
- `sandbox` (Boolean) Indicates whether sandboxing is enabled for file analysis.
- `sandbox_allow_level` (String) The level of sandboxing that is allowed.
- `sandbox_delay_file_transfer` (Boolean) Indicates whether file transfer is delayed until sandbox analysis results are received.
- `scan_detection` (String) Enable or disable Scan Detection for traffic that matches the rule. This overrides the option set in the NGFW properties.
- `snort` (Boolean) Flag to indicate whether Snort intrusion detection is enabled for the traffic matching this rule.
- `spooling_level` (String) The spooling level for file transfers, which determines how files are handled during processing.
- `sub_policy` (String) This represents a sub-policy that can be applied to various elements in the system, such as network elements, inspection rules, etc.
- `unavailable_method` (String) The method to handle unavailable files.
- `user_response` (String) This represents a User Response, which defines additional notification actions for rule matches, such as redirecting access to a forbidden URL to a page on an internal web server instead.
- `valid_block_lister` (List of String) URI of the network element used as blocklister.
- `vpn` (String) This represents a Policy Based Virtual Private Network (VPN), which is used to establish secure connections over unsecured networks. It includes various configurations such as NAT rules, mobile VPN topology modes, and associated profiles.

## Nested Attributes
- `connection_tracking_options` (Single Block, see [here](attr_connection_tracking_options.md)) 

