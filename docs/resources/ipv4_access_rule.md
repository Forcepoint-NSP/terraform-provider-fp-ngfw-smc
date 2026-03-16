---
page_title: "smc_fw_ipv4_access_rule"
subcategory: "policy"
description: |-
  This represents an IPv4 Access Rule. It defines how one type of IPv4 connection is handled by providing matching criteria based on the source, destination, and protocol information.
---

# smc_fw_ipv4_access_rule (Sub-resource)

This represents an IPv4 Access Rule. It defines how one type of IPv4 connection is handled by providing matching criteria based on the source, destination, and protocol information.


- see [here](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/policies/fw_ipv4_access_rule)
  for a working example of ipv4 access rule creation

Example:

```hcl
resource "smc_fw_policy" "example" {
  name     = "tf_example_access_rules_policy2"
  template = data.smc_href.default_policy_template.id
  comment  = var.resource_comment
}

resource "smc_fw_ipv4_access_rule" "allow_https" {
  from_ref = smc_fw_policy.example.link.fw_ipv4_access_rules
  name     = "allow-https-from-any-to-any"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.https_service.id]
  }
  action { action = ["allow"] }
}
```

the `from_ref` parameter must either point to:
- the link of the policy representing the list of access rules (eg
`smc_fw_policy.example.link.fw_ipv4_access_rules`)

- the add_after (or add_before) link of the previous rule (eg
  `smc_fw_ipv4_access_rule.allow_https.link.add_after`)


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `background_color` (String) The background color for the comment (rule section) and insert point rules. It is represented by its hexadecimal representation ('#RRGGBB').
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `rank` (Number) The rank of the Rule within the Policy. If not specified, the Rule will be the first one in the Policy.
- `rule_validity_times` (List of String) URI of the Rule Validity Time.
- `type` (String) The type of insert point for the rule. It can be 'normal' for a standard insert point or 'automatic' for an automatic rule insert point.

## Nested Attributes
- `action` (Single Block, see [here](attr_firewall_rule_action.md)) 
- `authentication_options` (Single Block, see [here](attr_user_authentication_match.md)) 
- `destinations` (Single Block, see [here](attr_destination_matching_part.md)) 
- `match_vpn_options` (Single Block, see [here](attr_match_vpn_option.md)) 
- `options` (Single Block, see [here](attr_firewall_rule_options.md)) 
- `services` (Single Block, see [here](attr_service_match_part.md)) 
- `sources` (Single Block, see [here](attr_source_match_part.md)) 

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
