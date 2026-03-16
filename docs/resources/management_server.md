---
page_title: "smc_mgt_server"
subcategory: "admin"
description: |-
  This represents a Management Server, which is the central component that stores all configurations and monitors the state of all NGFW Engines and other components in the Stonesoft Management Center. It provides access for Management Clients to change configurations or command the engines.
---

# smc_mgt_server (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Management Server, which is the central component that stores all configurations and monitors the state of all NGFW Engines and other components in the Stonesoft Management Center. It provides access for Management Clients to change configurations or command the engines.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `address` (String) The primary IPv4 address of the device, which is used for network communication.
- `alert_server_ref` (String) This represents a Log Server, which is a component of the Management Center responsible for storing and managing log (and alert) data, and analyzing and correlating events detected by multiple NGFW Engines.
- `announcement_enabled` (Boolean) Indicates whether announcements are enabled for the Management Server. If enabled, announcements are shown to all users.
- `announcement_message` (String) The announcement message for the Management Server, limited to 160 characters. You can add formatting using standard HTML tags.
- `comment` (String) An optional comment for the element. This field is not required.
- `db_replication` (Boolean) Indicates whether database replication is enabled for the Management Server. This value is deprecated when replication is based on PostgreSQL native replication.
- `elasticsearch_indexing_active` (Boolean) Indicates whether Elasticsearch indexing is active on this Management Server. Note that indexing also needs to be active on both the server and the Elasticsearch cluster itself.
- `ipv6_address` (String) The primary IPv6 address of the device, which is used for network communication.
- `license_update_enabled` (Boolean) Indicates whether the Management Server automatically generates and installs new licenses for system components when upgrading to a major new release.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `log_disk_space_handling_mode` (String) The mode chosen to handle extra logs when the disk runs out of space. Possible values are 'stop_receiving', 'delete_old_logs', and 'overwrite_old_logs'.
- `name` (String) Name of the object.
- `radius_method` (String) The RADIUS authentication method used by the Management Server for authenticating administrators. Possible values are 'pap', 'chap', 'mschap', 'mschap2', and 'eap-md5'. Default is 'eap-md5'.
- `script_path` (String) The root path for custom alert scripts used by the Management Server.
- `secondary` (List of String) A list of secondary IP addresses for the device, which can be used in policies and routing. You can add several IPv4 and IPv6 addresses (one by one).
- `sender_address` (String) The email address of the sender for notifications sent by the Management Server.
- `sender_name` (String) The name of the sender for notifications sent by the Management Server.
- `smtp_server_ref` (String) This represents a Simple Mail Transfer Protocol (SMTP) server, which is used to process notifications by e-mails. It includes attributes for port, default sender email, and default sender name.
- `snmp_gateways` (String) The SNMP gateway for the Management Server, which defines the SNMP settings for network management.
- `tacacs_method` (String) The TACACS authentication method used by the Management Server for authenticating administrators. Possible values are 'ascii', 'pap', 'chap', and 'mschap'. Default is 'mschap'.
- `tls_credentials` (String) This represents a TLS Server Credentials element, which is used to store the private key and certificate of an internal server. The certificate and the associated private key must be compatible with OpenSSL and be in PEM format. It is used for TLS inspection, securing Web Access Servers, and authenticating Authentication Servers.
- `tls_profile` (String) This represents a TLS Profile. It contains common data for establishing a TLS connection, including TLS version, cryptography suites, and trusted certificate authorities.
- `tools_profile_ref` (String) This represents a Tools Profile. Tools Profiles add commands to the right-click menus of elements, allowing dynamic information inclusion from the element definition. Only one Tools Profile can be selected for each element, but each can include several commands. Commands are launched on the workstation running the Management Client and are operating-system-specific.
- `uiid` (String) The unique identifier (UIID) for the Management Server installation, which is used to identify the server instance.
- `updates_check_enabled` (Boolean) Indicates whether the Management Server checks for updates automatically.
- `updates_proxy_address` (String) The address of the proxy server used by the Management Server for HTTPS connections to the Internet servers.
- `updates_proxy_authentication_enabled` (Boolean) Indicates whether the Management Server requires authentication for the proxy server used for HTTPS connections to the Internet servers.
- `updates_proxy_enabled` (Boolean) Indicates whether the Management Server uses a proxy server for HTTPS connections to the Internet servers.
- `updates_proxy_password` (String) The password for authentication to the proxy server used by the Management Server for HTTPS connections to the Internet servers.
- `updates_proxy_port` (Number) The port number of the proxy server used by the Management Server for HTTPS connections to the Internet servers.
- `updates_proxy_username` (String) The username for authentication to the proxy server used by the Management Server for HTTPS connections to the Internet servers.

## Nested Attributes
- `elasticsearch_authentication_settings` (Single Block, see [here](attr_elasticsearch_authentication_settings.md)) 
- `external_pki_certificate_settings` (Single Block, see [here](attr_certificate_settings.md)) 
- `forwarding_tls_settings` (Single Block, see [here](attr_tls_settings.md)) 
- `mgt_integration_container` (List of Blocks, see [here](attr_management_integration_container.md)) The Management Integration Containers for the Management Server, which define the settings for different management integrations.
- `netflow_collector` (List of Blocks, see [here](attr_netflow_collector.md)) The Netflow Collector to which the Management Server forwards audit data. You can define which type of audit data you want to forward and in which format, and use Filters to specify in detail which log data is forwarded.
- `sms_http_channel` (List of Blocks, see [here](attr_sms_http_channel.md)) The HTTP channel for the Management Server, which defines the settings for HTTP-based communication.
- `sms_script_channel` (List of Blocks, see [here](attr_sms_script_channel.md)) The Script channel for the Management Server, which defines the settings for script-based communication.
- `sms_smtp_channel` (List of Blocks, see [here](attr_sms_smtp_channel.md)) The SMTP channel for the Management Server, which defines the settings for SMTP-based communication.
- `third_party_monitoring` (Single Block, see [here](attr_third_party_monitoring.md)) 
- `web_app` (List of Blocks, see [here](attr_web_application_parameters.md)) The Web Application parameters for the Management Server, which define the settings for different web applications.

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
