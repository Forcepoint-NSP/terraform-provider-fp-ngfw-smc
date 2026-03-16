---
page_title: "smc_link_usage_profile"
subcategory: ""
description: |-
  This represents a Link Usage Profile, which is used to manage link usage settings in a multi-link environment. It includes configurations for link balancing, packet duplication, and forward erasure correction.
---

# smc_link_usage_profile (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Link Usage Profile, which is used to manage link usage settings in a multi-link environment. It includes configurations for link balancing, packet duplication, and forward erasure correction.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `default_link_balancing_preference` (Number) The Default SD-WAN Link balancing preference. This is a value comprised between 0 and 4. Default value is 0 and means Equal Balancing.
- `link_dont_duplicate_link_type_ref` (List of String) URI of the Link Type.
- `link_duplicate_qos_ref` (List of String) URI of the QoS Class.
- `link_fec_percentage` (Number) Forward Erasure Correction percentage indicating the ratio of correction packets to data packets.
- `link_fec_qos_ref` (List of String) URI of the QoS Class.
- `link_fec_threshold` (Number) Packet drop threshold that will start Forward Erasure Correction. This is a value comprised between 0 and 100.
- `link_fec_type_ref` (List of String) URI of the Link Type.
- `name` (String) Name of the object.

## Nested Attributes
- `link_usage_exception` (List of Blocks, see [here](attr_link_usage_entry.md)) A list of link usage exception entries that define specific link usage rules. Each entry can specify different link types and QoS classes for exceptions.

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
