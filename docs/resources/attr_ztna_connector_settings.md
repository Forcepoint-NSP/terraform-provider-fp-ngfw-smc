---
page_title: "ztna_connector_settings"
subcategory: ""
description: |-
  This represents the settings for the ZTNA connector feature on a Firewall element, including the installer key, data center, and auto-update option.
---

# ztna_connector_settings (Nested-Attribute)

This represents the settings for the ZTNA connector feature on a Firewall element, including the installer key, data center, and auto-update option.




## Simple Attributes
- `auto_update` (Boolean) If true, the ZTNA connector version will be checked upon the next policy push, and the latest ZTNA connector image will be downloaded if needed.
- `bgkey` (String) The installer key for the ZTNA connector, which can be found in the admin portal under protect/ZTNA/Installer Key.
- `datacenter` (String) The name of the data center accessed by this ZTNA connector, which can be found in the admin portal under protect/ZTNA/DATA CENTERS.


