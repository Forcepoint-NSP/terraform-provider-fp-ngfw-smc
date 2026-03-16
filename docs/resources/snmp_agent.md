---
page_title: "smc_snmp_agent"
subcategory: ""
description: |-
  This represents an SNMP Agent. It contains the configuration details for SNMP (Simple Network Management Protocol) on an engine, including SNMP version, users, monitoring settings, and trap destinations.
---

# smc_snmp_agent (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents an SNMP Agent. It contains the configuration details for SNMP (Simple Network Management Protocol) on an engine, including SNMP version, users, monitoring settings, and trap destinations.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `snmp_monitoring_contact` (String) The contact information of the person responsible for the engine. This is used for SNMP monitoring purposes.
- `snmp_monitoring_listening_port` (Number) The UDP listening port number that the SNMP agent listens to. Default is 161.
- `snmp_monitoring_user_name` (List of String) The SNMP monitoring user names. This is used for SNMP monitoring purposes.
- `snmp_trap_user_name` (String) The SNMP trap user name. This is used for SNMP trap notifications.
- `snmp_version` (String) The version of the SNMP agent.

## Nested Attributes
- `snmp_trap_destination` (List of Blocks, see [here](attr_snmp_trap_destination.md)) The SNMP trap destinations, which include the IP address and UDP port to which SNMP traps are sent.
- `snmp_user_name` (List of Blocks, see [here](attr_snmp_user.md)) The list of SNMP users.

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
