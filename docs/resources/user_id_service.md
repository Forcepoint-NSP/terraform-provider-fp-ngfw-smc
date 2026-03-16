---
page_title: "smc_user_id_service"
subcategory: "services"
description: |-
  This represents a User ID Service element, which is used to manage user identification services. It includes attributes for port, IP addresses, monitored user domains, cache expiration, connect timeout, TLS profile, and TLS identity.
---

# smc_user_id_service (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a User ID Service element, which is used to manage user identification services. It includes attributes for port, IP addresses, monitored user domains, cache expiration, connect timeout, TLS profile, and TLS identity.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `address` (String) The primary IPv4 address of the device, which is used for network communication.
- `cache_expiration` (Number) The time in seconds for the cache expiration on the engine. If not specified, it defaults to 300.
- `comment` (String) An optional comment for the element. This field is not required.
- `connect_timeout` (Number) The time in seconds for the connection from the engine to time out. If not specified, it defaults to 10.
- `ipv6_address` (String) The primary IPv6 address of the device, which is used for network communication.
- `list` (List of String) List of additional IP addresses to contact the User ID Service.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `monitored_user_domains` (List of String) The specific user domains to check. If not defined, it uses all known user domains by User ID service.
- `name` (String) Name of the object.
- `port` (Number) The port on which the User ID Service communicates with the Firewall. If you change the port from the default, you must configure the same port in the User ID Service Properties on the Windows system. You must also change the rule that allows communication between the Firewall and the User ID Service.
- `secondary` (List of String) A list of secondary IP addresses for the device, which can be used in policies and routing. You can add several IPv4 and IPv6 addresses (one by one).
- `tls_profile` (String) This represents a TLS Profile. It contains common data for establishing a TLS connection, including TLS version, cryptography suites, and trusted certificate authorities.
- `tools_profile_ref` (String) This represents a Tools Profile. Tools Profiles add commands to the right-click menus of elements, allowing dynamic information inclusion from the element definition. Only one Tools Profile can be selected for each element, but each can include several commands. Commands are launched on the workstation running the Management Client and are operating-system-specific.

## Nested Attributes
- `third_party_monitoring` (Single Block, see [here](attr_third_party_monitoring.md)) 
- `tls_identity` (Single Block, see [here](attr_tls_identity.md)) 

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
