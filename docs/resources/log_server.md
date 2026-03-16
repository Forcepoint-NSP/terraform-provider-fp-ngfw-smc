---
page_title: "smc_log_server"
subcategory: "network_elements"
description: |-
  This represents a Log Server, which is a component of the Management Center responsible for storing and managing log (and alert) data, and analyzing and correlating events detected by multiple NGFW Engines.
---

# smc_log_server (Resource)

This represents a Log Server, which is a component of the Management Center responsible for storing and managing log (and alert) data, and analyzing and correlating events detected by multiple NGFW Engines.

## Examples

Defines an SMC Log Server resource for receiving logs from engines.

```hcl
resource "smc_log_server" "new_log_server" {
  address = "127.0.0.10"
  channel_port = 3020
  log_disk_space_handling_mode = "overwrite_oldest"
  name = "Log Server - New"
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `address` (String) The primary IPv4 address of the device, which is used for network communication.
- `channel_port` (Number) The TCP port number for the Log Server. We recommend using default port 3020 if possible. To use a non-standard port, manually add Access rules to allow communications using the new port from the NGFW Engines to the Log Server.
- `comment` (String) An optional comment for the element. This field is not required.
- `elasticsearch_indexing_active` (Boolean) Indicates whether Elasticsearch indexing is active on this server. Note: indexing also needs to be active on both the server and the Elasticsearch cluster itself.
- `inactive` (Boolean) Indicates whether the Log Server is inactive. If true, it is excluded from Log Browsing, Reporting and Statistics.
- `ipv6_address` (String) The primary IPv6 address of the device, which is used for network communication.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `log_disk_space_handling_mode` (String) The mode chosen to handle extra logs when disk runs out of space.
- `name` (String) Name of the object.
- `secondary` (List of String) A list of secondary IP addresses for the device, which can be used in policies and routing. You can add several IPv4 and IPv6 addresses (one by one).
- `tools_profile_ref` (String) This represents a Tools Profile. Tools Profiles add commands to the right-click menus of elements, allowing dynamic information inclusion from the element definition. Only one Tools Profile can be selected for each element, but each can include several commands. Commands are launched on the workstation running the Management Client and are operating-system-specific.
- `uiid` (String) The unique installation ID (UIID) of the Log Server, used to identify the server in the system.

## Nested Attributes
- `backup_log_server` (List of Blocks, see [here](attr_secondary_log_server_container_for_storage.md)) The backup Log Server for this Log Server. You can specify several backup Log Servers, and the same Log Server can simultaneously be the main Log Server for some components and a backup Log Server for others.
- `elasticsearch_authentication_settings` (Single Block, see [here](attr_elasticsearch_authentication_settings.md)) 
- `external_pki_certificate_settings` (Single Block, see [here](attr_certificate_settings.md)) 
- `forwarding_tls_settings` (Single Block, see [here](attr_tls_settings.md)) 
- `netflow_collector` (List of Blocks, see [here](attr_netflow_collector.md)) The Netflow Collector associated with the Log Server. Log Servers can be configured to forward log data to external hosts, and you can define which type of log data you want to forward and in which format.
- `third_party_monitoring` (Single Block, see [here](attr_third_party_monitoring.md)) 

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
