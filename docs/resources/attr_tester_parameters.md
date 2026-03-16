---
page_title: "tester_parameters"
subcategory: ""
description: |-
  This represents the Tester Global Settings defined for an engine, including alert intervals, boot delays, and recovery options.
---

# tester_parameters (Nested-Attribute)

This represents the Tester Global Settings defined for an engine, including alert intervals, boot delays, and recovery options.




## Simple Attributes
- `alert_interval` (Number) The interval in seconds before sending a new alert when the same test keeps failing repeatedly. The default value is 3600 seconds (60 minutes).
- `auto_recovery` (Boolean) Indicates whether the engine should automatically go back online after a reboot if all offline tests report success.
- `boot_delay` (Number) The boot delay in seconds that the engine waits before resuming tests after listed events. The default value is 30 seconds (maximum: 1800 seconds).
- `boot_recovery` (Boolean) Indicates whether the engine should automatically go back online after a reboot if all offline tests report success.
- `restart_delay` (Number) The restart delay in seconds that the engine waits before resuming tests after listed events. The default value is 5 seconds (maximum: 1800 seconds).
- `status_delay` (Number) The status delay in seconds that the engine waits before resuming tests after listed events. The default value is 5 seconds (maximum: 1800 seconds).


