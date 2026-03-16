---
page_title: "smc_snapshot"
subcategory: "policy"
description: |-
  This represents a Policy Snapshot, which is a record of policy configuration that shows the configuration in the form it was installed or refreshed. It includes rules, elements, properties, upload time, and the administrator who performed the upload. It helps in tracking configuration changes.
---

# smc_snapshot (Sub-resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Policy Snapshot, which is a record of policy configuration that shows the configuration in the form it was installed or refreshed. It includes rules, elements, properties, upload time, and the administrator who performed the upload. It helps in tracking configuration changes.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.


## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `alternative_slot` (Number) The alternative slot number if the policy was uploaded as an alternative policy, indicating which slot it occupies.
- `config_id` (Number) The configuration ID associated with this policy snapshot, used for dynamic uploads.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `package_id` (Number) The ID of the update package associated with this policy snapshot, indicating which package was activated when the policy was uploaded.
- `policy_name` (String) The name of the policy associated with this snapshot, if the policy does not exist.
- `policy_ref` (String) This represents a policy that can be applied to various elements in the system, such as network elements, inspection rules, etc.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `target` (String) This is the base class for all storable elements.
- `target_name` (String) The name of the target (Cluster/...) associated with this snapshot, if the target does not exist.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
- `upload_time` (Number) The date and time when the policy was uploaded, represented as a Unix timestamp in milliseconds.
- `uploaded_rule_tags` (String) A comma-separated list of rule tags that were uploaded with this policy snapshot, used for tracking rule counters.
- `uploader` (String) The name of the person who started the upload, such as an administrator.
