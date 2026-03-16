---
page_title: "antivirus"
subcategory: ""
description: |-
  This represents the settings for the antivirus feature on a firewall element, including antivirus enablement, log levels, update frequencies, and proxy settings.
---

# antivirus (Nested-Attribute)

This represents the settings for the antivirus feature on a firewall element, including antivirus enablement, log levels, update frequencies, and proxy settings.




## Simple Attributes
- `antivirus_enabled` (Boolean) Indicates whether the antivirus feature is enabled on the firewall element.
- `antivirus_http_proxy_enabled` (Boolean) Indicates whether the HTTP Proxy is enabled for antivirus updates. By default, it is not enabled.
- `antivirus_proxy_password` (String) The HTTP Proxy password used for authentication when accessing the antivirus update server.
- `antivirus_proxy_port` (Number) The HTTP Proxy port used for antivirus updates.
- `antivirus_proxy_user` (String) The HTTP Proxy username used for authentication when accessing the antivirus update server.
- `antivirus_update` (String) The frequency at which the antivirus updates are performed.
- `antivirus_update_day` (String) The day of the week when antivirus updates are performed, represented as a two-letter abbreviation.
- `antivirus_update_time` (Number) The time of the day the Antivirus updates are performed, given as a long value in a 24-hour format.
- `virus_alert` (String) This represents an abstract Alert, which is used to display messages when certain conditions are met.
- `virus_log_level` (String) The log level for antivirus events, which determines how antivirus-related logs are handled.
- `virus_mirror` (String) The URL of the antivirus signature mirrors used for updating malware signatures.


