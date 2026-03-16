---
page_title: "smc_bgp_connection_profile"
subcategory: "routing"
description: |-
  This represents the BGP Connection Profile for Dynamic Routing Firewall functionality. It contains settings for BGP connections, such as session timers and password.
---

# smc_bgp_connection_profile (Resource)

This represents the BGP Connection Profile for Dynamic Routing Firewall functionality. It contains settings for BGP connections, such as session timers and password.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/dynamic_routing/BGP/bgp_peering) for an example

This example creates a BGP connection profile with session and authentication settings.

```hcl
resource "smc_bgp_connection_profile" "bgp_connection_profile" {
  connect            = 150
  md5_password       = "TheSuperPassword00!!"
  name               = "tf_bgp_connection_profile"
  session_hold_timer = 190
  session_keep_alive = 70
  comment            = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `connect` (Number) The connect retry timer, in seconds. It is not required and defaults to 60 seconds.
- `md5_password` (String) The MD5 password for the BGP connection. It is not required but can be used for authentication.
- `name` (String) Name of the object.
- `session_hold_timer` (Number) The session hold timer, in seconds. It is not required and defaults to 180 seconds.
- `session_keep_alive` (Number) The session keep alive timer, in seconds. It is not required and defaults to 60 seconds.


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
