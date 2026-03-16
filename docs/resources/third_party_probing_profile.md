---
page_title: "smc_probing_profile"
subcategory: ""
description: |-
  This represents a Probing Profile used in Third Party Monitoring. It contains settings that define how a Log Server monitors third-party components.
---

# smc_probing_profile (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Probing Profile used in Third Party Monitoring. It contains settings that define how a Log Server monitors third-party components.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `cpu_usage_oid` (String) The OID for monitoring CPU usage. It is used to specify the SNMP OID that provides CPU usage information.
- `disk_space_oid` (String) The OID for monitoring disk space. It is used to specify the SNMP OID that provides disk space information.
- `interface_monitored` (Boolean) Indicates whether interfaces should be monitored.
- `mem_free_oid` (String) The OID for monitoring free memory. It is used to specify the SNMP OID that provides free memory information.
- `mem_used_oid` (String) The OID for monitoring memory usage. It is used to specify the SNMP OID that provides memory usage information.
- `name` (String) Name of the object.
- `probing_interval` (Number) The interval in seconds at which the probing is performed.
- `probing_method` (String) The probing method used for monitoring.
- `retry_count` (Number) The number of retries to perform if the probing fails.
- `snmp_community` (String) The SNMP community string used for SNMP monitoring.
- `snmp_port` (Number) The SNMP port used for monitoring.
- `snmpv3_auth_password` (String) The SNMP V3 authentication password used for securing SNMP communications.
- `snmpv3_auth_protocol` (String) The SNMP V3 authentication protocol used for securing SNMP communications.
- `snmpv3_context_name` (String) The SNMP V3 context name used for monitoring. It is used to specify the context in which the SNMP operations are performed.
- `snmpv3_priv_password` (String) The SNMP V3 privacy password used for encrypting SNMP communications.
- `snmpv3_priv_protocol` (String) The SNMP V3 privacy protocol used for encrypting SNMP communications.
- `snmpv3_sec_name` (String) The SNMP V3 security name used for authentication.
- `timeout` (Number) The timeout in milliseconds for the probing operation.


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
