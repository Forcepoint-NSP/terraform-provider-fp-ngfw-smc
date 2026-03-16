---
page_title: "smc_geolocation"
subcategory: ""
description: |-
  This represents the definition of a geolocation. A Geolocation object keeps a list of Network Elements belonging to the same geolocation.
---

# smc_geolocation (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents the definition of a geolocation. A Geolocation object keeps a list of Network Elements belonging to the same geolocation.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `city` (String) City where the geolocation is situated.
- `comment` (String) An optional comment for the element. This field is not required.
- `country_code` (String) Country code where the geolocation is situated. Default value is 'O1' which corresponds to 'Other'.
- `element` (List of String) URI of the Storable element.
- `latitude` (Number) Latitude of the geolocation.
- `logo_ref` (String) This is a Logo File. It represents a logo file used in the system, typically for branding or identification purposes.
- `longitude` (Number) Longitude of the geolocation.
- `name` (String) Name of the object.
- `postal_code` (String) Postal code for the geolocation.
- `region` (String) Region where the geolocation is situated.
- `street_address1` (String) Street address (part 1).
- `street_address2` (String) Street address (part 2).


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
