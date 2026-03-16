---
page_title: "external_test"
subcategory: ""
description: |-
  This represents the definition of an external test to be executed on a NGFW. It runs a custom script stored on the engine. If the script returns the code zero (0), the test is considered successful, otherwise the test is considered failed. Caution: This test allows administrators who have permissions to edit the properties of Engines to run arbitrary commands in the Engine operating system.
---

# external_test (Nested-Attribute)

This represents the definition of an external test to be executed on a NGFW. It runs a custom script stored on the engine. If the script returns the code zero (0), the test is considered successful, otherwise the test is considered failed. Caution: This test allows administrators who have permissions to edit the properties of Engines to run arbitrary commands in the Engine operating system.




## Simple Attributes
- `alert_notification` (Boolean) Indicates whether an alert notification is sent if the test fails.
- `command_line` (String) The command line to execute the external test script. The script must return an exit code of 0 (zero) if it succeeds. Any non-zero return value is considered a failure. Caution: This test allows administrators who have permissions to edit the properties of Engines to run arbitrary commands in the Engine operating system.
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `offline_state` (Boolean) Indicates whether the test is executed when the engine is offline or not.
- `online_state` (Boolean) Indicates whether the test is executed when the engine is online or not.
- `retry_count` (Number) The number of times the test will be retried if it fails. Note! We recommend always setting the retry count to more than 1 to avoid creating overly sensitive tests that burden the system unnecessarily.
- `standby_state` (Boolean) Indicates whether the test is executed when the engine is in standby state or not.
- `test_action` (String) The action to be taken if the test fails. Options include 'none', 'offline', 'forceoffline', and 'forcespeed'.
- `test_active` (Boolean) Indicates whether the test is active or not.
- `test_interval` (Number) The interval in seconds at which the test is executed. Note! Running a test too frequently can increase overhead.
- `test_timeout` (Number) The timeout in milliseconds for the external test. If the test does not return a response within this time, it is considered to have failed. To avoid overly short timeout values. We recommend a timeout of 500 - 1000 ms depending on the external test script


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
