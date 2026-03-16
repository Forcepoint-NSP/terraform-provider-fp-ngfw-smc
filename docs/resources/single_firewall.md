---
page_title: "smc_single_fw"
subcategory: "engines"
description: |-
  This represents a single Firewall device in the Stonesoft Management Center. It includes attributes for Snort configuration, High-Availability mode, and connection synchronization settings.
---

# smc_single_fw (Resource)

This represents a single Firewall device in the Stonesoft Management Center. It includes attributes for Snort configuration, High-Availability mode, and connection synchronization settings.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/getting_started) for a complete minimal example
- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/engines/single_fw) for the list of examples related to single firewall

This example creates a single firewall with two interfaces.  The first
interface will be used to contact the smc (there must be exactly one
such interface).

```hcl
resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw"
  log_server_ref = data.smc_href.log_server.href
  comment        = var.resource_comment
  nodes {
    firewall_node {
      name   = "myfwnode"
      nodeid = 1
    }
  }

  physical_interfaces {
    physical_interface {
      interface_id = 0
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 0
          address       = "192.168.100.14"
          network_value = "192.168.100.00/24"
          primary_mgt   = true
        }
      }
    }
  }
  physical_interfaces {
    physical_interface {
      interface_id = 1
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 1
          address       = "10.0.0.10"
          network_value = "10.0.0.0/24"
        }
      }
    }
  }
}

```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `active_server_certificate_probing` (Boolean) Flag for enabling/disabling TLS probe connections.
- `active_wait_time` (String) The Active Wait Time. 'short' Short, 'medium' Medium, 'long' Long.
- `admin_auth_method` (String) This represents an external authentication method, which can be used for user authentication in the system. It supports various types of authentication methods such as User Password, IAS, IPSec, Pre-Shared Key, RADIUS, TACACS+, and LDAP.
- `allow_email_upn_lookup` (Boolean) Indicates whether email UPN lookup is allowed. If true, the system allows lookup from known User Domain matching to client certificate email domain or UPN suffix.
- `allow_long_userid_lookup` (Boolean) Indicates whether long UserID lookup is allowed. If true, the system allows username lookup using Long UserID attribute.
- `allow_root_login` (String) Indicates whether root login is allowed for administrator authentication. Options include 'allow' for allowing root login, 'disable_pwd_ssh' for disabling root password login through SSH, and 'disable' for disabling all root password logins, including console access.
- `antispoofing_node_ref` (String) This represents an Antispoofing Node in the Security Management Client, which is used to protect against malicious packages with altered IP header information.
- `auto_reboot_timeout` (Number) The length of time after which an error situation is considered non-recoverable and the engine automatically reboots. The default value is 10 seconds.
- `client_cert_identity_field` (String) The Client Certificate Identity Field that defines how the client certificate identity is determined.
- `comment` (String) An optional comment for the element. This field is not required.
- `connection_limit` (Number) Sets a limit for connections from a single source and/or destination IP address. When the set number of connections is reached, the next connection attempts are blocked by the engine until a previously open connection is closed.
- `connection_sync_group` (String) This represents a Connection Sync Group, which is used for external high availability in single Engines. It includes attributes for monitoring and a list of elements that are part of the group.
- `connection_sync_mode` (String) The Connection Synchronization mode for External High Availability, which can be either disabled or enabled. When enabled, it allows synchronization of connection states between firewalls.
- `contact_timeout` (Number) Used when opening any kind of communication between the Management Server and the engines. A consistently slow network connection may require increasing this value. The default value is 60 seconds.
- `control_plane_mode` (String) The control plane mode of the Master Engine, which determines the allocation of CPU resources for control plane operations. Options include 'not_reserved' for no dedicated CPU allocation, 'one_cpu' for 1 CPU allocated, 'two_cpus' for 2 CPUs allocated, 'three_cpus' for 3 CPUs allocated, and 'four_cpus' for 4 CPUs allocated.
- `custom_properties_profile` (List of String) URI of the custom properties profile.
- `default_nat` (String) Indicates whether the Default NAT Address is used for Traffic from Internal Networks. It can be set to 'true', 'false', or 'automatic'.
- `default_user_domain` (String) This represents a User Domain, which is used to define the authentication domain for users. It can be either an authentication domain or the internal domain.
- `destination_server_certificate_cache_timeout` (Number) The timeout value for certificate entries in the cache in minutes.
- `disable_ahm` (Boolean) Indicates whether Application Health Monitoring (AHM) is disabled. If true, AHM is not performed for the Firewall Cluster.
- `discard_quic_if_cant_inspect` (Boolean) Indicates whether QUIC traffic should be discarded when inspection is not possible.
- `dns_relay_profile_ref` (String) This represents a DNS Relay profile element, which contains settings for hostname mappings, domain-specific DNS servers, fixed domain answers, and DNS answer translations.
- `dos_protection` (String) The DOS Protection Mode that defines the level of protection against Denial of Service attacks. It can be set to 'always_off', 'default_off', or 'default_on'.
- `enable_saml_for_application_access` (Boolean) Indicates whether the SAML authentication method is enabled for Application Access Portal. If true, SAML authentication is used for Application Access Portal.
- `enable_saml_for_bba` (Boolean) Indicates whether the SAML authentication method is enabled for user authentication (BBA). If true, SAML authentication is used for browser-based authentication.
- `engine_version` (String) The version of the engine that this cluster is running.
- `excluded_interface` (Number) The ID of the interface that is excluded from IP counting. By default, not defined.
- `fuid_network_filter` (List of String) The network filtering address ranges for the engine. It is used by the firewall to restrict the range of events to be received from the FUID server.
- `gateway_settings_ref` (String) This represents various gateway and VPN parameters to be set at Firewall level. These settings are used for the negotiation of VPN connections.
- `geolocation_ref` (String) This represents the definition of a geolocation. A Geolocation object keeps a list of Network Elements belonging to the same geolocation.
- `granted_policy_ref` (List of String) URI of the granted policies.
- `ha_backup_unit_mode` (String) The High-Availability mode for a single IP, which can be either disabled or enabled. When enabled, it may require a proof of serial for pairing.
- `ha_connection_sync_interface` (String) The interface for Unicast State Sync, specified by its ID, when the Connection Synchronization for External High Availability mode is enabled. This interface is used for synchronizing connection states between firewalls.
- `ha_pos_for_backup_unit` (String) The proof of serial for pairing when the High-Availability mode is enabled. This is required to establish a connection between the primary and backup units.
- `half_open_connections` (Number) The number of Half Open Connections allowed. If not specified, it defaults to 125.
- `icap_dlp_server_ref` (List of String) URI of the ICAP Server.
- `include_interfaces_for_control_plane` (Boolean) Indicates whether interfaces used for management and log connections are included in the control plane. If true, these interfaces are included in the control plane operations.
- `inspection_cpu_balancing_mode` (String) The Inspection CPU Balancing Mode. 'default' Default value, 'round_robin' Round Robin, 'numa' NUMA local Round Robin.
- `integrated_uis_ref` (String) This represents an Integrated User Identification Service, which is used to identify users based on their network activity. It allows for integration with authentication domains and provides configuration options for initial query time, polling interval, and ignore values.
- `internal_gateway_ref` (List of String) URI of the Internal Gateway.
- `is_cert_auto_renewal` (Boolean) Indicates whether automated certificate renewal is enabled for this cluster.
- `is_config_encrypted` (Boolean) Is the configuration file encrypted?
- `is_fips_compatible_operating_mode` (Boolean) Indicates whether the FIPS mode is enabled. FIPS standard specifies the security requirements that will be satisfied by a cryptographic module utilized within a security system protecting sensitive but unclassified information.
- `is_fips_disable_engine_sginfo` (Boolean) Indicates whether the creation of sginfo for engines in FIPS mode is disabled.
- `is_fips_disable_engine_upgrades` (Boolean) Indicates whether the ability to remotely upgrade engines in FIPS mode is disabled.
- `is_icap_dlp_enabled` (Boolean) Indicates whether ICAP DLP (Data Loss Prevention) is enabled for the firewall. If true, ICAP DLP servers are used to inspect and filter data for compliance and security purposes.
- `is_loopback_tunnel_ip_address_enforced` (Boolean) Indicates whether the Loopback Tunnel IP is enforced. If true, the Loopback Tunnel IP is used instead of the default outgoing IP if no IP Address is provided on Tunnel Interfaces.
- `is_snort_enabled` (Boolean) Indicates whether Snort is enabled for this Firewall.
- `is_snort_file_defined` (Boolean) Indicates whether a Snort configuration file is defined for this Firewall.
- `is_virtual_defrag` (Boolean) When the engine receives fragmented packets, it defragments the packets for inspection. This setting defines how the fragmented packets are sent onwards after inspection if the packets are allowed. If this option is selected, the fragmented packets are sent onwards using the same fragmentation as when they arrived at the engine. If the option is not selected, the packets are sent onwards as if they had arrived unfragmented.
- `known_host_lists_ref` (List of String) URI of the Known Host List.
- `link_usage_profile_ref` (String) This represents a Link Usage Profile, which is used to manage link usage settings in a multi-link environment. It includes configurations for link balancing, packet duplication, and forward erasure correction.
- `lldp_profile_ref` (String) This represents a LLDP Profile, which is used to configure the LLDP (Link Layer Discovery Protocol) settings for devices.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `log_server_ref` (String) This represents a Log Server, which is a component of the Management Center responsible for storing and managing log (and alert) data, and analyzing and correlating events detected by multiple NGFW Engines.
- `log_spooling_policy` (String) The Log Spooling Policy that defines what happens when the engine's log spool becomes full. It can be set to 'stop' (Stop Traffic) or 'discard' (Discard Log).
- `multicast_routing_mode` (String) The Multicast Routing Mode that defines how multicast traffic is handled. 'none' disables the multicast routing. 'static' enables the Static multicast routing. 'igmp_proxy' enables the IGMP Proxy multicast routing. 'pim_ipv4' enables the PIM multicast routing.
- `name` (String) Name of the object.
- `nondecrypted_ca_certificate_ref` (List of String) URI of the non-decrypted CA certificates.
- `nondecrypted_tls_server_credentials_ref` (List of String) URI of the non-decrypted TLS server credentials.
- `passive_discard_mode` (Boolean) Indicates whether the Passive Discard Mode is enabled. If true, it does not stop matching connections but creates a special log entry Terminate (passive) for testing purposes.
- `quic_enabled` (Boolean) Indicates whether QUIC ports are enabled for Web Traffic. If true, QUIC protocol is allowed for web traffic.
- `reporting_email_addresses` (String) The contact email address for reporting per sender.
- `rollback_timeout` (Number) The time the engine waits for a management connection before it rolls back to the previously installed policy when the Policy Handshake option is active. The default value is 60 seconds.
- `routing_node_ref` (String) This represents a Routing Node in the Security Management Client, which is used to configure routing for network traffic on firewalls.
- `saml_bba_clock_skew` (Number) The clock skew for SAML BBA authentication in seconds. It defines the allowed time difference between the SAML assertion and the system time.
- `saml_ssl_vpn_clock_skew` (Number) The clock skew for SAML Application Access authentication in seconds. It defines the allowed time difference between the SAML assertion and the system time.
- `sandbox_type` (String) The Sandbox service selection.
- `send_reset_for_out_of_state_tcp_packets` (String) The engine refuses TCP connections if the TCP connection does not start with a SYN packet, even if the TCP connection matches an Access rule with the Allow action.
- `server_credential` (List of String) URI of the TLS Server Credentials.
- `sidewinder_logging_profile_ref` (String) This represents a Sidewinder Logging Profile, which contains settings for logging in Sidewinder.
- `sidewinder_proxy_enabled` (Boolean) Indicates whether the Sidewinder Proxy is enabled. If true, the Sidewinder Proxy is used for enhanced security and logging.
- `slow_request_block_list_timeout` (Number) The DOS Protection Slow Request Block List timeout in seconds. It defines how long a slow request is blocked before it can be retried.
- `slow_request_sensitivity` (String) The DOS Protection Slow Request Sensitivity that defines the sensitivity level for slow requests. It can be set to 'off', 'low', 'medium', or 'high'.
- `snmp_agent_ref` (String) This represents an SNMP Agent. It contains the configuration details for SNMP (Simple Network Management Protocol) on an engine, including SNMP version, users, monitoring settings, and trap destinations.
- `snmp_location` (String) The location for the current SNMP client.
- `ssh_passwordless_login` (String) Indicates whether SSH passwordless logins are allowed for administrator authentication. Options include 'allow' for allowing passwordless logins and 'deny' for denying them.
- `strict_tcp_mode` (Boolean) Indicates whether Strict TCP mode is enabled. When enabled, it provides enhanced protection against TCP evasion attempts by controlling the progress of a TCP connection and enforcing the order of packets.
- `syn_flood_sensitivity` (String) The DOS Protection Syn Flood Sensitivity that defines the sensitivity level for SYN flood attacks. It can be set to 'off', 'low', 'medium', or 'high'.
- `syn_max_bursts` (Number) The SYN Burst Size value (>= 1). The number of allowed SYNs before the engine starts limiting the SYN rate.
- `syn_mode` (String) The SYN Rate Limits Mode. 'off' SYN Rate Limits are disabled. This is the default setting. 'auto' The engine automatically calculates the number of Allowed SYNs per Second and the Burst Size for the interface based on the engine's capacity and memory size. 'custom' Enter the desired values for Allowed SYNs per Second and Burst Size. We recommend that the Burst Size be at least one tenth of the Allowed SYNs per Second value. If the Burst Size is too small, SYN Rate Limits do not work.
- `syn_per_second` (Number) The Allowed SYNs per Second value (>= 1). The number of allowed SYN packets per second.
- `tcp_reset_sensitivity` (String) The DOS TCP Reset Sensitivity that defines the sensitivity level for TCP reset attacks. It can be set to 'off', 'low', 'medium', or 'high'.
- `timezone` (String) The timezone setting for the firewall, which determines the time zone used for logging and scheduling operations.
- `tls_crl_checks` (Boolean) Indicates whether the engine performs CRL checks for TLS certificates.
- `tls_cryptography_suite_set_ref` (String) This represents a TLS Cryptography Suite Set Element, which contains a set of cryptographic suites used in SSL VPN configurations.
- `tls_deny_decrypting` (Boolean) Indicates whether the engine denies decrypting TLS traffic.
- `tools_profile_ref` (String) This represents a Tools Profile. Tools Profiles add commands to the right-click menus of elements, allowing dynamic information inclusion from the element definition. Only one Tools Profile can be selected for each element, but each can include several commands. Commands are launched on the workstation running the Management Client and are operating-system-specific.
- `tracking_mode` (String) The Connection Tracking Mode. 'normal' When selected, the engine drops ICMP error messages related to connections that are not currently active in connection tracking. A valid, complete TCP handshake is required for TCP traffic. The engine checks the traffic direction and the port parameters of UDP traffic. 'strict' When selected, the engine does not permit TCP traffic to pass through before a complete, valid TCP handshake is performed. 'loose' When selected, the engine allows some connection patterns and address translation operations that are not allowed in the Normal mode. Note! You can override this general setting and configure connection tracking for TCP, UDP, and ICMP traffic in Access rules.
- `user_id_service_ref` (String) This represents a User ID Service element, which is used to manage user identification services. It includes attributes for port, IP addresses, monitored user domains, cache expiration, connect timeout, TLS profile, and TLS identity.

## Nested Attributes
- `advanced_multilink_settings` (Single Block, see [here](attr_advanced_multi_link_settings.md)) 
- `alias_value` (List of Blocks, see [here](attr_alias_value.md)) The alias values for the engine.
- `antivirus` (Single Block, see [here](attr_antivirus_settings.md)) 
- `automatic_rules_settings` (Single Block, see [here](attr_automatic_rules_settings.md)) 
- `certificate_validation_settings` (Single Block, see [here](attr_certificate_validation_settings.md)) 
- `connection_timeout` (List of Blocks, see [here](attr_idle_timeout.md)) The definitions of timeout by protocol or by TCP connection state.
- `dns_relay_interface` (List of Blocks, see [here](attr_dns_relay_interface_entry.md)) The DNS Relay Interface Entries that define the interfaces used for DNS Relay in the Firewall Cluster.
- `domain_server_address` (List of Blocks, see [here](attr_dns_element.md)) The DNS elements that define the IP addresses of the DNS server used by the Virtual IPS. You can enter a single IP address manually or define an IP address using a network element.
- `dynamic_routing` (Single Block, see [here](attr_dynamic_routing_settings.md)) 
- `eca_settings` (Single Block, see [here](attr_eca_settings.md)) 
- `file_reputation_settings` (Single Block, see [here](attr_file_reputation_settings.md)) 
- `ipv6_transition_mechanism` (Single Block, see [here](attr_nat464_settings.md)) 
- `l2fw_settings` (Single Block, see [here](attr_l2fw_settings.md)) 
- `link_usage_exception_rules` (List of Blocks, see [here](attr_link_usage_exception_rule.md)) The Link Usage Exception Rules that define specific routing rules for outbound network traffic in the Firewall Cluster.
- `local_log_storage` (Single Block, see [here](attr_local_log_storage_settings.md)) 
- `log_moderation` (List of Blocks, see [here](attr_log_moderation.md)) The Log Moderation settings that define how log entries are moderated in the Firewall Cluster.
- `nat_definition` (List of Blocks, see [here](attr_nat_definition.md)) The NAT Definitions that define the Network Address Translation rules for the Firewall Cluster.
- `nodes` (List of Blocks, see [here](attr_engine_node_wrapper.md)) The nodes that are part of this cluster. Each node represents a device that shares the workload.
- `ntp_settings` (Single Block, see [here](attr_ntp_settings.md)) 
- `physical_interfaces` (List of Blocks, see [here](attr_abstract_physical_interface_wrapper.md)) The specific physical interfaces for the engine.
- `pim_settings` (Single Block, see [here](attr_pim_settings.md)) 
- `policy_route` (List of Blocks, see [here](attr_policy_route.md)) The Policy Routes that define the routing rules for the Firewall Cluster. These routes are used to control the flow of traffic based on specific criteria.
- `saml_settings` (List of Blocks, see [here](attr_saml_settings.md)) The SAML Settings for user authentication (BBA) and Application Access Portal.
- `sandbox_settings` (Single Block, see [here](attr_sandbox_settings.md)) 
- `scan_detection` (Single Block, see [here](attr_scan_detection_settings.md)) 
- `snmp_interface` (List of Blocks, see [here](attr_snmp_interface_entry.md)) The SNMP listening interfaces for the engine.
- `ssh_host_key` (List of Blocks, see [here](ssh_host_key.md)) The SSH Host Keys that define the SSH keys used for secure communication in the Firewall Cluster.
- `ssm_advanced_setting` (List of Blocks, see [here](attr_sidewinder_proxy_advanced_settings.md)) The Sidewinder Proxy Advanced Settings that define additional configurations for the Sidewinder Proxy.
- `static_multicast_route` (List of Blocks, see [here](attr_static_multicast_route.md)) The Static Multicast Route entries that define the static multicast routing configuration.
- `tester_parameters` (Single Block, see [here](attr_tester_parameters.md)) 
- `tests` (List of Blocks, see [here](attr_test_entry_wrapper.md)) The test entries for the engine.
- `tls_client_protection` (Single Block, see [here](attr_tls_client_protection_wrapper.md)) 
- `ts_settings` (Single Block, see [here](attr_threat_seeker_settings.md)) 
- `web_authentication` (Single Block, see [here](attr_user_authentication_bba.md)) 
- `ztna_connector_settings` (Single Block, see [here](attr_ztna_connector_settings.md)) 

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
