---
page_title: "smc_role"
subcategory: ""
description: |-
  This represents an Administration Role. Roles are administrative privileges that allow you to limit rights of administrators.
---

# smc_role (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an Administration Role. Roles are administrative privileges that allow you to limit rights of administrators.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `admin_mgmt` (Boolean) 
- `alert_mgmt` (Boolean) 
- `approve_changes` (Boolean) 
- `archive_log_mgmt` (Boolean) 
- `backup_mgmt` (Boolean) 
- `bookmark_manage` (Boolean) 
- `comment` (String) An optional comment for the element. This field is not required.
- `delete_log_mgmt` (Boolean) 
- `element_create` (Boolean) 
- `element_delete` (Boolean) 
- `element_edit` (Boolean) 
- `element_view_content` (Boolean) 
- `export_log_mgmt` (Boolean) 
- `internal_user_mgmt` (Boolean) 
- `license_mgmt` (Boolean) 
- `log_pruning_mgmt` (Boolean) 
- `name` (String) Name of the object.
- `overview_manage` (Boolean) 
- `refresh_policy` (Boolean) 
- `send_advanced_commands` (Boolean) 
- `send_commands` (Boolean) 
- `updates_and_upgrades_mgmt` (Boolean) 
- `upload_policy` (Boolean) 
- `view_audit` (Boolean) 
- `view_edit_report` (Boolean) 
- `view_logs` (Boolean) 
- `view_system_alerts` (Boolean) 
- `vpn_mgmt` (Boolean) 


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
