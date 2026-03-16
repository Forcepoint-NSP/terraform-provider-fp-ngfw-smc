---
page_title: "smc_alert_chain_rule"
subcategory: "policy"
description: |-
  This represents a Alert Chain Rule for Alert Chain Policy. It contains information about the alert channel, destination, delay, and admin users.
---

# smc_alert_chain_rule (Sub-resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a Alert Chain Rule for Alert Chain Policy. It contains information about the alert channel, destination, delay, and admin users.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `admin_name` (List of String) URI of the administrator in case of user notification channel.
- `alert_channel` (String) The alert channel.
- `amount` (Number) The maximum number of notifications to be sent before activating moderation. This is used to control the frequency of notifications.
- `background_color` (String) The background color for the comment (rule section) and insert point rules. It is represented by its hexadecimal representation ('#RRGGBB').
- `comment` (String) An optional comment for the element. This field is not required.
- `delay` (Number) The delay before the next notification in minutes. This is applicable for channels that support delays.
- `destination` (String) The destination address for the alert channel. This is required for channels like SMTP, SMS, and Custom Script.
- `is_disabled` (Boolean) Indicates whether the rule is disabled.
- `name` (String) Name of the object.
- `notify_first_block` (Number) Indicates whether to notify the first blocked notification upon moderation activation. This is used to alert the user about the first blocked notification.
- `period` (Number) The period during which notifications are tracked before activating moderation. This is used to determine the time frame for monitoring notifications.
- `rank` (Number) The rank of the Rule within the Policy. If not specified, the Rule will be the first one in the Policy.
- `type` (String) The type of insert point for the rule. It can be 'normal' for a standard insert point or 'automatic' for an automatic rule insert point.


## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `parent_insert_point` (String) This represents a rule in a policy, including its parent policy, parent insert point, rank, tag, background color, and insert point type.
- `parent_policy` (String) This represents a policy that can be applied to various elements in the system, such as network elements, inspection rules, etc.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `tag` (String) The unique identifier of the rule in this policy. It contains a static part that does not change when rules are added, removed, or moved, and a changing part that indicates the version of the rule.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
