---
page_title: "smc_dns_server"
subcategory: "network_elements"
description: |-
  This represents an External DNS (Domain Name Service) Server, which is used for DDNS updates, resolving virus signature mirrors, and URL filtering categorization services.
---

# smc_dns_server (Resource)

This represents an External DNS (Domain Name Service) Server, which is used for DDNS updates, resolving virus signature mirrors, and URL filtering categorization services.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/1.741.0/examples/network_elements/servers/external_dns_server) for a complete minimal example

This example creates a DNS server object in SMC with TTL and update interval specified.

```hcl
resource "smc_dns_server" "tf_dns_server" {
  address         = "10.10.10.1"
  name            = "tf_dns_server"
  comment         = var.resource_comment
  time_to_live    = 10
  update_interval = 5
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `address` (String) The primary IPv4 address of the server. Must be a valid IPv4 address.
- `comment` (String) An optional comment for the element. This field is not required.
- `ipv6_address` (String) The primary IPv6 address of the server. Must be a valid IPv6 address.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `secondary` (List of String) A list of secondary IP addresses for the device, which can be used in policies and routing. You can add several IPv4 and IPv6 addresses (one by one).
- `time_to_live` (Number) The Time to Live (TTL) value for DNS entries in seconds. It defines how long a DNS entry can be cached before querying the DNS server again.
- `tools_profile_ref` (String) This represents a Tools Profile. Tools Profiles add commands to the right-click menus of elements, allowing dynamic information inclusion from the element definition. Only one Tools Profile can be selected for each element, but each can include several commands. Commands are launched on the workstation running the Management Client and are operating-system-specific.
- `update_interval` (Number) The Update Interval in seconds for DNS entries. It defines how often the DNS entries can be updated to the DNS server if the link status changes constantly.

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
