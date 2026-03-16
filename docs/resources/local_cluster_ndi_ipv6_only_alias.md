---
page_title: "smc_local_cluster_ndi_for_ipv6_only_alias"
subcategory: "network_elements"
description: |-
  This represents the System alias for '$$ Local Cluster(NDI IPv6 addresses only)', which is used to substitute all NDI IPv6 addresses of the local cluster object where the policy is uploaded. If the firewall is a single node, all defined addresses are used.
---

# smc_local_cluster_ndi_for_ipv6_only_alias (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents the System alias for '$$ Local Cluster(NDI IPv6 addresses only)', which is used to substitute all NDI IPv6 addresses of the local cluster object where the policy is uploaded. If the firewall is a single node, all defined addresses are used.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `default_alias_value` (List of String) URI of the default value Storable.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.

## Nested Attributes
- `alias_value` (List of Blocks, see [here](attr_alias_value.md)) A list of alias values that define specific configurations for the Alias. Each alias value can specify different configurations for the Alias in policies.

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
