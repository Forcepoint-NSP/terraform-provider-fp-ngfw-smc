---
page_title: "smc_ntp"
subcategory: "network_elements"
description: |-
  This represents an NTP server, which is used to synchronize the time across network devices. It includes attributes for host name, authentication key type, and optional authentication key ID and key.
---

# smc_ntp (Resource)

This represents an NTP server, which is used to synchronize the time across network devices. It includes attributes for host name, authentication key type, and optional authentication key ID and key.

## Examples

No usage found in the current example set.

A stub configuration:

```hcl
# resource "smc_ntp" "example" {
#   # Example attributes here
# }
```

> This resource type does not appear in any official example yet.


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `address` (String) The primary IPv4 address of the device, which is used for network communication.
- `comment` (String) An optional comment for the element. This field is not required.
- `ipv6_address` (String) The primary IPv6 address of the device, which is used for network communication.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `ntp_auth_key` (String) The authentication key used by the NTP server. This is optional and can be null.
- `ntp_auth_key_id` (Number) The ID of the authentication key used by the NTP server. This is optional and can be null.
- `ntp_auth_key_type` (String) The type of authentication key used by the NTP server.
- `ntp_host_name` (String) The host name of the NTP server.
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
