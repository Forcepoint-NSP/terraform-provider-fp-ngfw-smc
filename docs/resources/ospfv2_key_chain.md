---
page_title: "smc_ospfv2_key_chain"
subcategory: "routing"
description: |-
  This represents the OSPFv2 Key Chain element used as Message Digest authentication method for OSPFv2 Interface Settings for Dynamic Routing Firewall functionality.
---

# smc_ospfv2_key_chain (Resource)

This represents the OSPFv2 Key Chain element used as Message Digest authentication method for OSPFv2 Interface Settings for Dynamic Routing Firewall functionality.

## Examples

- [OSPFv2 Key Chain Example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/engines/dynamic_routing/OSPFv2/ospfv2_key_chain/main.tf)

Defines a chain of authentication keys for OSPFv2 interface security.

```hcl
resource "smc_ospfv2_key_chain" "ospfv2_key_chain" {
  name    = "tf_ospfv2_key_chain"
  comment = var.resource_comment
  ospfv2_key_chain_entry {
    algorithm = "md5"
    comment   = "md5"
    key_id    = 1
    send_key  = true
    key       = "123"
  }
  ospfv2_key_chain_entry {
    algorithm = "hmac-sha-512"
    comment   = ""
    key_id    = 2
    send_key  = false
    key       = "564"
  }
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.

## Nested Attributes
- `ospfv2_key_chain_entry` (List of Blocks, see [here](attr_ospf_key_chain_entry.md)) The OSPFv2 Key Chain Entries, which define the keys used for authentication in the OSPFv2 Key Chain.

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
