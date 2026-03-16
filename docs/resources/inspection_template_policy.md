---
page_title: "smc_inspection_template_policy"
subcategory: "policy"
description: |-
  This represents an Inspection Template Policy, which is used to define a set of rules and insert points that can be inherited by other policies or template policies.
---

# smc_inspection_template_policy (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an Inspection Template Policy, which is used to define a set of rules and insert points that can be inherited by other policies or template policies.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `file_filtering_policy` (String) This represents a File Filtering Policy, which is used to define the action and inspection rules for File Filtering Engines.
- `inspection_policy` (String) This represents an Inspection Template Policy, which is used to define a set of rules and insert points that can be inherited by other policies or template policies.
- `name` (String) Name of the object.
- `template` (String) This represents a normal policy that can be applied to a Firewall, IPS, or Layer 2 device. It includes references to template policies for access control and inspection, as well as file filtering policies.


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
