---
page_title: "smc_virtual_ips"
subcategory: "engines"
description: |-
  This represents a Virtual IPS, which is a logically-separate IPS engine that runs as a virtual IPS instance on a Master NGFW Engine. It includes attributes for DNS elements, log buffer size, and various protection modes.
---

# smc_virtual_ips (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Virtual IPS, which is a logically-separate IPS engine that runs as a virtual IPS instance on a Master NGFW Engine. It includes attributes for DNS elements, log buffer size, and various protection modes.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `active_server_certificate_probing` (Boolean) Flag for enabling/disabling TLS probe connections.
- `active_wait_time` (String) The Active Wait Time. 'short' Short, 'medium' Medium, 'long' Long.
- `antispoofing_node_ref` (String) This represents an Antispoofing Node in the Security Management Client, which is used to protect against malicious packages with altered IP header information.
- `auto_reboot_timeout` (Number) The length of time after which an error situation is considered non-recoverable and the engine automatically reboots. The default value is 10 seconds.
- `comment` (String) An optional comment for the element. This field is not required.
- `contact_timeout` (Number) Used when opening any kind of communication between the Management Server and the engines. A consistently slow network connection may require increasing this value. The default value is 60 seconds.
- `custom_properties_profile` (List of String) URI of the custom properties profile.
- `destination_server_certificate_cache_timeout` (Number) The timeout value for certificate entries in the cache in minutes.
- `discard_quic_if_cant_inspect` (Boolean) Indicates whether QUIC traffic should be discarded when inspection is not possible.
- `dos_protection` (String) The DOS Protection Mode for the Virtual IPS. It defines how the engine handles Denial of Service (DOS) attacks, with options for always off, default off, and default on.
- `engine_version` (String) The version of the engine that this cluster is running.
- `fuid_network_filter` (List of String) The network filtering address ranges for the engine. It is used by the firewall to restrict the range of events to be received from the FUID server.
- `geolocation_ref` (String) This represents the definition of a geolocation. A Geolocation object keeps a list of Network Elements belonging to the same geolocation.
- `granted_policy_ref` (List of String) URI of the granted policies.
- `half_open_connections` (Number) The number of Half Open Connections for the Virtual IPS. It defines the maximum number of half-open TCP connections that the engine can handle simultaneously.
- `inspection_cpu_balancing_mode` (String) The Inspection CPU Balancing Mode. 'default' Default value, 'round_robin' Round Robin, 'numa' NUMA local Round Robin.
- `integrated_uis_ref` (String) This represents an Integrated User Identification Service, which is used to identify users based on their network activity. It allows for integration with authentication domains and provides configuration options for initial query time, polling interval, and ignore values.
- `is_cert_auto_renewal` (Boolean) Indicates whether automated certificate renewal is enabled for this cluster.
- `is_config_encrypted` (Boolean) Is the configuration file encrypted?
- `lldp_profile_ref` (String) This represents a LLDP Profile, which is used to configure the LLDP (Link Layer Discovery Protocol) settings for devices.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `log_buffer_size` (Number) The size of the log buffer in bytes. This is the amount of memory allocated for storing logs before they are written to disk.
- `log_server_ref` (String) This represents a Log Server, which is a component of the Management Center responsible for storing and managing log (and alert) data, and analyzing and correlating events detected by multiple NGFW Engines.
- `max_spooled_conf_count` (Number) The maximum number of configurations that can be spooled in memory. This is used to manage the number of configurations stored temporarily before being applied.
- `name` (String) Name of the object.
- `nondecrypted_ca_certificate_ref` (List of String) URI of the non-decrypted CA certificates.
- `nondecrypted_tls_server_credentials_ref` (List of String) URI of the non-decrypted TLS server credentials.
- `passive_discard_access_mode` (Boolean) Indicates whether passive discard mode for access is enabled. When enabled, it allows connections that match rules with the Terminate action in Access Rules to be logged without stopping traffic, useful for testing purposes.
- `passive_discard_mode` (Boolean) Indicates whether passive discard mode is enabled. When enabled, it allows connections that match rules with the Terminate action to be logged without stopping traffic, useful for testing purposes.
- `reporting_email_addresses` (String) The contact email address for reporting per sender.
- `rollback_timeout` (Number) The time the engine waits for a management connection before it rolls back to the previously installed policy when the Policy Handshake option is active. The default value is 60 seconds.
- `routing_node_ref` (String) This represents a Routing Node in the Security Management Client, which is used to configure routing for network traffic on firewalls.
- `sandbox_type` (String) The Sandbox service selection.
- `send_reset_for_out_of_state_tcp_packets` (String) The engine refuses TCP connections if the TCP connection does not start with a SYN packet, even if the TCP connection matches an Access rule with the Allow action.
- `server_credential` (List of String) URI of the TLS Server Credentials.
- `slow_request_block_list_timeout` (Number) The DOS Protection Slow Request Block List timeout in seconds. It defines how long an IP address is blocked after a slow request attack is detected.
- `slow_request_sensitivity` (String) The DOS Protection Slow Request Sensitivity for the Virtual IPS. It defines how sensitive the engine is to slow request attacks, with options for off, low, medium, and high.
- `snmp_agent_ref` (String) This represents an SNMP Agent. It contains the configuration details for SNMP (Simple Network Management Protocol) on an engine, including SNMP version, users, monitoring settings, and trap destinations.
- `snmp_location` (String) The location for the current SNMP client.
- `strict_tcp_mode` (Boolean) Indicates whether Strict TCP mode is enabled. When enabled, it provides enhanced protection against TCP evasion attempts by controlling the progress of a TCP connection and enforcing strict adherence to the TCP protocol.
- `syn_flood_sensitivity` (String) The DOS Protection Syn Flood Sensitivity for the Virtual IPS. It defines how sensitive the engine is to SYN flood attacks, with options for off, low, medium, and high.
- `syn_max_bursts` (Number) The SYN Burst Size value (>= 1). The number of allowed SYNs before the engine starts limiting the SYN rate.
- `syn_mode` (String) The SYN Rate Limits Mode. 'off' SYN Rate Limits are disabled. This is the default setting. 'auto' The engine automatically calculates the number of Allowed SYNs per Second and the Burst Size for the interface based on the engine's capacity and memory size. 'custom' Enter the desired values for Allowed SYNs per Second and Burst Size. We recommend that the Burst Size be at least one tenth of the Allowed SYNs per Second value. If the Burst Size is too small, SYN Rate Limits do not work.
- `syn_per_second` (Number) The Allowed SYNs per Second value (>= 1). The number of allowed SYN packets per second.
- `tcp_reset_sensitivity` (String) The DOS TCP Reset Sensitivity for the Virtual IPS. It defines how sensitive the engine is to TCP reset attacks, with options for off, low, medium, and high.
- `tls_crl_checks` (Boolean) Indicates whether the engine performs CRL checks for TLS certificates.
- `tls_cryptography_suite_set_ref` (String) This represents a TLS Cryptography Suite Set Element, which contains a set of cryptographic suites used in SSL VPN configurations.
- `tls_deny_decrypting` (Boolean) Indicates whether the engine denies decrypting TLS traffic.
- `tools_profile_ref` (String) This represents a Tools Profile. Tools Profiles add commands to the right-click menus of elements, allowing dynamic information inclusion from the element definition. Only one Tools Profile can be selected for each element, but each can include several commands. Commands are launched on the workstation running the Management Client and are operating-system-specific.
- `tracking_mode` (String) The Connection Tracking Mode. 'normal' When selected, the engine drops ICMP error messages related to connections that are not currently active in connection tracking. A valid, complete TCP handshake is required for TCP traffic. The engine checks the traffic direction and the port parameters of UDP traffic. 'strict' When selected, the engine does not permit TCP traffic to pass through before a complete, valid TCP handshake is performed. 'loose' When selected, the engine allows some connection patterns and address translation operations that are not allowed in the Normal mode. Note! You can override this general setting and configure connection tracking for TCP, UDP, and ICMP traffic in Access rules.
- `user_id_service_ref` (String) This represents a User ID Service element, which is used to manage user identification services. It includes attributes for port, IP addresses, monitored user domains, cache expiration, connect timeout, TLS profile, and TLS identity.
- `virtual_resource` (String) This represents a set of resources on the Master NGFW Engine that are allocated to each Virtual NGFW Engine. It includes properties such as ID, domain allocation, and resource limits.

## Nested Attributes
- `advanced_multilink_settings` (Single Block, see [here](attr_advanced_multi_link_settings.md)) 
- `alias_value` (List of Blocks, see [here](attr_alias_value.md)) The alias values for the engine.
- `antivirus` (Single Block, see [here](attr_antivirus_settings.md)) 
- `automatic_rules_settings` (Single Block, see [here](attr_automatic_rules_settings.md)) 
- `certificate_validation_settings` (Single Block, see [here](attr_certificate_validation_common_settings.md)) 
- `connection_timeout` (List of Blocks, see [here](attr_idle_timeout.md)) The definitions of timeout by protocol or by TCP connection state.
- `domain_server_address` (List of Blocks, see [here](attr_dns_element.md)) The DNS elements that define the IP addresses of the DNS server used by the Virtual IPS. You can enter a single IP address manually or define an IP address using a network element.
- `eca_settings` (Single Block, see [here](attr_eca_settings.md)) 
- `file_reputation_settings` (Single Block, see [here](attr_file_reputation_settings.md)) 
- `local_log_storage` (Single Block, see [here](attr_local_log_storage_settings.md)) 
- `physical_interfaces` (List of Blocks, see [here](attr_abstract_physical_interface_wrapper.md)) The specific physical interfaces for the engine.
- `sandbox_settings` (Single Block, see [here](attr_sandbox_settings.md)) 
- `scan_detection` (Single Block, see [here](attr_scan_detection_settings.md)) 
- `snmp_interface` (List of Blocks, see [here](attr_snmp_interface_entry.md)) The SNMP listening interfaces for the engine.
- `tester_parameters` (Single Block, see [here](attr_tester_parameters.md)) 
- `tests` (List of Blocks, see [here](attr_test_entry_wrapper.md)) The test entries for the engine.
- `tls_client_protection` (Single Block, see [here](attr_tls_client_protection_wrapper.md)) 
- `ts_settings` (Single Block, see [here](attr_threat_seeker_settings.md)) 

## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `master_ref` (String) This represents a physical engine device that provides resources for Virtual Security Engines. One physical Master NGFW Engine can support multiple Virtual NGFW Engines.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
