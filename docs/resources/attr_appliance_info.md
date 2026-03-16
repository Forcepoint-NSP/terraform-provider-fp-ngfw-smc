---
page_title: "appliance_info"
subcategory: "engines"
description: |-
  This represents the Appliance data for an engine Node, including product name, hardware version, initial contact time, first upload time, proof of serial, software version, software features, cloud identifiers, and remaining days before initial license expiration.
---

# appliance_info (Nested-Attribute)

This represents the Appliance data for an engine Node, including product name, hardware version, initial contact time, first upload time, proof of serial, software version, software features, cloud identifiers, and remaining days before initial license expiration.




## Simple Attributes
- `cloud_id` (String) The unique identifier for the appliance in the cloud service, such as an AWS instance ID.
- `cloud_type` (String) The type of cloud service associated with the appliance, such as AWS or Azure.
- `first_upload_time` (Number) The first upload time of the appliance in milliseconds since epoch.
- `hardware_version` (String) The hardware version of the appliance.
- `initial_contact_time` (Number) The initial contact time of the appliance in milliseconds since epoch.
- `initial_license_remaining_days` (Number) The number of remaining days before the initial license expires.
- `product_name` (String) The name of the product associated with the appliance.
- `proof_of_serial` (String) The proof of serial number for the appliance, used for validation purposes.
- `software_features` (String) The features available in the software running on the appliance.
- `software_version` (String) The version of the software running on the appliance.


