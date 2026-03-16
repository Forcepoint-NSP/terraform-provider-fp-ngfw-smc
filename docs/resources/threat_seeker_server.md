---
page_title: "smc_threatseeker_server"
subcategory: "engines"
description: |-
  This represents a Threat Seeker Server, which is a System object holding configuration needed to use Threat Seeker URL Categories. It includes attributes for the complete URL, license key, expiration date, and CA certificate.
---

# smc_threatseeker_server (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Threat Seeker Server, which is a System object holding configuration needed to use Threat Seeker URL Categories. It includes attributes for the complete URL, license key, expiration date, and CA certificate.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `dds_api_key` (String) The API key for accessing the ThreatSeeker DDS service.
- `dds_url` (String) The URL to the ThreatSeeker DDS service.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `ts_certificate` (String) The CA certificate for the Threat Seeker Server, which is used to establish a secure connection.
- `tskey_exp` (String) The expiration date of the Threat Seeker License Key
- `url` (String) The complete URL to the Threat Seeker Server.


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
