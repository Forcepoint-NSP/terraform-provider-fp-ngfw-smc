---
page_title: "smc_radius_server"
subcategory: ""
description: |-
  This represents a RADIUS Server, which is an external authentication server that supports the RADIUS protocol. It can be used as an authentication method for administrators and includes attributes for default authentication port.
---

# smc_radius_server (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a RADIUS Server, which is an external authentication server that supports the RADIUS protocol. It can be used as an authentication method for administrators and includes attributes for default authentication port.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `address` (String) The primary IPv4 address of the server. Must be a valid IPv4 address.
- `comment` (String) An optional comment for the element. This field is not required.
- `ipv6_address` (String) The primary IPv6 address of the server. Must be a valid IPv6 address.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `port` (Number) The port number for the authentication server. If the server communicates on a port other than the default port. The predefined Firewall Template allows the engines to connect to the default port. If you change to a custom port, you must add a new IPv4 Access Rule to allow the traffic.
- `provided_method` (List of String) URI of the Authentication Method.
- `retries` (Number) The number of retries for the authentication server connection. It defines how many times Firewalls try to connect to the RADIUS or TACACS+ authentication server if the connection fails.
- `secondary` (List of String) A list of secondary IP addresses for the device, which can be used in policies and routing. You can add several IPv4 and IPv6 addresses (one by one).
- `shared_secret` (String) The shared secret for the authentication server. It is used for RADIUS clients on the Active Directory server.
- `timeout` (Number) The timeout value in seconds for the authentication server. It defines how long Firewalls wait for the RADIUS or TACACS+ authentication server to reply.
- `tools_profile_ref` (String) This represents a Tools Profile. Tools Profiles add commands to the right-click menus of elements, allowing dynamic information inclusion from the element definition. Only one Tools Profile can be selected for each element, but each can include several commands. Commands are launched on the workstation running the Management Client and are operating-system-specific.

## Nested Attributes
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
