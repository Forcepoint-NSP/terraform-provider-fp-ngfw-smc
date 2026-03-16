---
page_title: "smc_known_host"
subcategory: ""
description: |-
  This represents a Known Host object used with the Sidewinder SSH Proxy. A Known Host is identified by an IP address, port number and an SSH Key. Only certain SSH Key types are supported.
---

# smc_known_host (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Known Host object used with the Sidewinder SSH Proxy. A Known Host is identified by an IP address, port number and an SSH Key. Only certain SSH Key types are supported.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `ipaddress` (String) Single valid IPv4 address. Either an IPv4 or IPv6 address is mandatory.
- `ipv6_address` (String) Single valid IPv6 address. Either an IPv4 or IPv6 address is mandatory.
- `known_host_public_key` (String) The SSH key value in OpenSSH public key format.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `port` (Number) Port number for the SSH connection.
- `sshkey_type` (String) Type of the SSH key used for the connection.


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
