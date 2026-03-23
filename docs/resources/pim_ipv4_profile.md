---
page_title: "smc_pim_ipv4_profile"
subcategory: "routing"
description: |-
  This represents the PIM IPv4 Profile for Dynamic Routing Firewall functionality. It is used to configure PIM (Protocol Independent Multicast) settings in the firewall's dynamic routing capabilities.
---

# smc_pim_ipv4_profile (Resource)

This represents the PIM IPv4 Profile for Dynamic Routing Firewall functionality. It is used to configure PIM (Protocol Independent Multicast) settings in the firewall's dynamic routing capabilities.

## Examples

- [PIM IPv4 Profile Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/1.741.0/examples/engines/dynamic_routing/PIM/pim_profiles/main.tf)

Defines a PIM (Protocol Independent Multicast) IPv4 profile for use in multicast configurations.

```hcl
resource "smc_pim_ipv4_profile" "pim_profile" {
  hello_interval = 45
  joined_prune   = 80
  name           = "tf_pim_profile"
  comment        = var.resource_comment
  pim_multicast_group_entry {
    mapping              = "10.0.1.100"
    mode                 = "pim_sm"
    multicast_ip_network = "224.0.1.0/24"
  }
  pim_multicast_group_entry {
    mapping              = "source1.example.com"
    mode                 = "pim_ssm"
    multicast_ip_network = "232.1.1.0/24"
  }
  pim_multicast_group_entry {
    mapping              = ""
    mode                 = "pim_sm"
    multicast_ip_network = "239.255.0.0/16"
  }
  pim_multicast_group_entry {
    mapping              = "10.0.2.50"
    mode                 = "pim_ssm"
    multicast_ip_network = "232.10.10.0/24"
  }
  smart_multicast_antispoofing = true
  # spt_switch_interval          = 90
  # spt_switch_threshold         = 10
  # spt_switch_threshold_unit    = "packets"
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `hello_interval` (Number) The Hello Interval in seconds. This is the interval at which PIM Hello messages are sent to maintain neighbor relationships.
- `joined_prune` (Number) The Joined-Prune interval in seconds. This value determines how long a PIM router will wait before sending a prune message after a join.
- `name` (String) Name of the object.
- `smart_multicast_antispoofing` (Boolean) Indicates whether Smart Multicast Antispoofing is enabled. When enabled, it helps prevent multicast spoofing attacks.

## Nested Attributes
- `pim_multicast_group_entry` (List of Blocks, see [here](attr_pim_ipv4_multicast_group_entry.md)) The PIM Multicast Group entries. These entries define the multicast groups that the PIM Profile will manage.

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
