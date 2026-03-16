---
page_title: "log_moderation"
subcategory: ""
description: |-
  This represents the Log Moderation settings for an engine or interface, which allows for compression of log entries to save resources.
---

# log_moderation (Nested-Attribute)

This represents the Log Moderation settings for an engine or interface, which allows for compression of log entries to save resources.




## Simple Attributes
- `burst` (Number) The maximum number of matching log entries in a single burst. The default value for Antispoofing entries is 1000. By default, Discard log entries are not compressed.
- `log_event` (String) The type of log event for moderation.
- `rate` (Number) The maximum number of log entries per second.


