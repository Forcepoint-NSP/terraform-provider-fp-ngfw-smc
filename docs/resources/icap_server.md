---
page_title: "smc_icap"
subcategory: "network_elements"
description: |-
  This represents an ICAP server, which is a Network Element that represents an ICAP instance of server. It includes attributes for port, path, secure ICAP, TLS profile, and X-Headers.
---

# smc_icap (Resource)

This represents an ICAP server, which is a Network Element that represents an ICAP instance of server. It includes attributes for port, path, secure ICAP, TLS profile, and X-Headers.

## Examples

No usage found in the current example set.

A stub configuration:

```hcl
# resource "smc_icap" "example" {
#   # Example attributes here
# }
```

> This resource type does not appear in the repository examples at this time.


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `address` (String) The primary IPv4 address of the device, which is used for network communication.
- `comment` (String) An optional comment for the element. This field is not required.
- `icap_include_xhdrs` (Boolean) Flag indicating if X-Headers should be included in the ICAP requests. Defaults to false.
- `icap_path` (String) The path for the ICAP server.
- `icap_port` (Number) The port number for the ICAP server.
- `icap_secure` (Boolean) Flag indicating if secure ICAP is enabled. Defaults to false.
- `ipv6_address` (String) The primary IPv6 address of the device, which is used for network communication.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `secondary` (List of String) A list of secondary IP addresses for the device, which can be used in policies and routing. You can add several IPv4 and IPv6 addresses (one by one).
- `tls_profile_ref` (String) This represents a TLS Profile. It contains common data for establishing a TLS connection, including TLS version, cryptography suites, and trusted certificate authorities.
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
