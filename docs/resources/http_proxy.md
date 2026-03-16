---
page_title: "smc_http_proxy"
subcategory: ""
description: |-
  This represents an HTTP Proxy, which is used to route HTTP traffic through a proxy server. It includes attributes for IP address, port, username, and password.
---

# smc_http_proxy (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an HTTP Proxy, which is used to route HTTP traffic through a proxy server. It includes attributes for IP address, port, username, and password.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `address` (String) The primary IP address of the HTTP Proxy. Must be a valid IPv4 or IPv6 address.
- `comment` (String) An optional comment for the element. This field is not required.
- `http_proxy_password` (String) The password for the user on the HTTP Proxy. This is stored in a ciphered format.
- `http_proxy_port` (Number) The port number for the HTTP Proxy.
- `http_proxy_username` (String) The username for the HTTP Proxy. This is used for authentication.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `secondary` (List of String) A list of secondary IP addresses for the device, which can be used in policies and routing. You can add several IPv4 and IPv6 addresses (one by one).
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
