---
page_title: "smc_broker_config_file"
subcategory: "vpn"
description: |-
  This is a configuration file for the VPN Broker Domain. It contains settings and configurations specific to the broker domain, used for managing VPN connections and policies.
---

# smc_broker_config_file (Resource)

This is a configuration file for the VPN Broker Domain. It contains settings and configurations specific to the broker domain, used for managing VPN connections and policies.

## Examples

- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/fw_cluster/fw_cluster_with_vpn_broker_interface) for an example

This example defines a VPN Broker configuration file for use in VPN Broker domains.

```hcl
resource "smc_broker_config_file" "vpn_broker_config" {
  name           = "tf_vpn_broker_config"
  binary_content = "aGVsbG8K"
  comment        = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `binary_content` (String) The binary content of the file, encoded in base64 format. This is the actual binary data.
- `comment` (String) An optional comment for the element. This field is not required.
- `comptype` (String) The compression type of the binary file, such as 'gzip', 'zip', etc.
- `max_version` (String) The maximum supported engine version for this binary.
- `min_version` (String) The minimum supported engine version for this binary.
- `name` (String) Name of the object.
- `type` (String) The type of the binary file, such as 'Firmware', 'Configuration', etc.
- `upload_always` (Boolean) Indicates whether the binary should always be uploaded, regardless of version compatibility.


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
