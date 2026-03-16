---
page_title: "smc_limit_rule"
subcategory: "policy"
description: |-
  This represents a QoS Limit Rule. It can be used to set limits on network traffic based on various parameters.
---

# smc_limit_rule (Sub-resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This represents a QoS Limit Rule. It can be used to set limits on network traffic based on various parameters.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `background_color` (String) The background color for the comment (rule section) and insert point rules. It is represented by its hexadecimal representation ('#RRGGBB').
- `comment` (String) An optional comment for the element. This field is not required.
- `dscp_match_mask` (Number) Mask for matching the six-bit DSCP field in incoming packets. This is used to filter packets based on their DSCP value.
- `dscp_match_value` (Number) Match value for the six-bit DSCP field. This is used to match packets based on their DSCP value.
- `dscp_set_mask` (Number) Mask for setting the six-bit DSCP field in outgoing packets. This is used to control which bits of the DSCP value are modified.
- `dscp_set_value` (Number) Value to set for outgoing packets for the six-bit DSCP field. This is used to modify the DSCP value of packets being sent out.
- `guarantee` (Number) The bandwidth guarantee for this QoS rule in kbit/s. Negative values indicate percents of the interface capacity.
- `latency` (Number) The optional latency for configuring active queue management scheduling algorithm. This value helps in managing the delay in packet processing.
- `limit` (Number) The bandwidth limit for this QoS rule in kbit/s. Negative values indicate percents of the interface capacity.
- `name` (String) Name of the object.
- `priority` (Number) The traffic priority for this QoS rule. This value determines the priority of the traffic handled by this rule.
- `qos_class_ref` (String) This represents a QoS Class, which is an element that links a rule in a QoS Policy to one or more Firewall Actions. The traffic allowed in the access rule is assigned the QoS Class defined for the rule, and the QoS class is used as the matching criteria for applying QoS Policy rules.
- `rank` (Number) The rank of the Rule within the Policy. If not specified, the Rule will be the first one in the Policy.
- `type` (String) The type of insert point for the rule. It can be 'normal' for a standard insert point or 'automatic' for an automatic rule insert point.
- `weight` (Number) The optional weight in percents used for assigning to priority queues when the class guarantee is reached. This helps in managing traffic prioritization effectively.


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
