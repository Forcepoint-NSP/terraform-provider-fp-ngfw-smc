---
page_title: "ntp_settings"
subcategory: "engines"
description: |-
  This represents the definition of NTP settings, which includes enabling NTP, defining NTP servers, and setting a preferred NTP server.
---

# ntp_settings (Nested-Attribute)

This represents the definition of NTP settings, which includes enabling NTP, defining NTP servers, and setting a preferred NTP server.




## Simple Attributes
- `ntp_enable` (Boolean) Indicates if NTP is enabled.
- `ntp_preferred_server_ref` (String) This represents an NTP server, which is used to synchronize the time across network devices. It includes attributes for host name, authentication key type, and optional authentication key ID and key.
- `ntp_server_ref` (List of String) URI of the NTP Server.


