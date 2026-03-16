---
page_title: "smc_server_pool"
subcategory: ""
description: |-
  This represents a Server Pool, which is a group of servers used for inbound traffic management. It includes attributes for DNS name, monitoring frequency, monitoring mode, and server allocation.
---

# smc_server_pool (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Server Pool, which is a group of servers used for inbound traffic management. It includes attributes for DNS name, monitoring frequency, monitoring mode, and server allocation.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `domain_name` (String) The DNS name of the Server Pool, which is used to route traffic to the pool.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `monitoring_frequency` (Number) The frequency in seconds at which the Server Pool monitors the availability of its servers.
- `monitoring_host` (String) The host name of the web server used for monitoring. This is used in HTTP mode to check server availability.
- `monitoring_mode` (String) The monitoring mode used to check the availability of servers in the Server Pool. 'ping' uses ICMP echo requests, 'agent' uses Monitoring Agents, 'tcp' checks TCP services, and 'http' checks HTTP services.
- `monitoring_port` (Number) The port number used for monitoring the servers in the Server Pool.
- `monitoring_request` (String) The request string sent to the server for monitoring. This is used in HTTP mode to check server availability.
- `monitoring_response` (String) The expected response string from the server for monitoring. This is used in HTTP mode to verify server availability.
- `monitoring_url` (String) The URL path to the web page used for monitoring the server. This is used in HTTP mode to check server availability.
- `name` (String) Name of the object.
- `server_allocation` (String) The server allocation method used to distribute traffic among servers in the Server Pool. 'order' allocates by priority, 'network' by source network, 'host' by source host, 'connection' per connection, and 'notdefined' behaves like 'network'.

## Nested Attributes
- `ip_netlink_weight` (List of Blocks, see [here](attr_ip_net_link_weight.md)) List of IP NetLink Weights used to influence traffic distribution among servers in the Server Pool.
- `members_list` (List of Blocks, see [here](attr_server_pool_member.md)) List of members in the Server Pool. Each member represents a server that handles incoming traffic.

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
