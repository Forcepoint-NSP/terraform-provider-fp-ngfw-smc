---
page_title: "local_log_storage"
subcategory: ""
description: |-
  This represents the Local Log Storage settings for the FWIPSCluster, which is not applicable to virtual engines.
---

# local_log_storage (Nested-Attribute)

This represents the Local Log Storage settings for the FWIPSCluster, which is not applicable to virtual engines.




## Simple Attributes
- `lls_guaranteed_free_percent` (Number) Defines a minimum amount of spool space that must be left available for other uses in percentage. By default, this is set to -1, which means that the value is not defined.
- `lls_guaranteed_free_size_in_mb` (Number) Defines a minimum amount of spool space that must be left available for other uses in MegaBytes. By default, this is set to -1, which means that the value is not defined.
- `lls_max_time` (Number) Defines the maximum amount of hours before the stored logs are deleted. By default, this is set to -1, which means that the logs will not be deleted automatically.
- `local_log_storage_activated` (Boolean) Activate the Local Log Storage feature. At least one of the Guaranteed free disk partition values must be set up.


