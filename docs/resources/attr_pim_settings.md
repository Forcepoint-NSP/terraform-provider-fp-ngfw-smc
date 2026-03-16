---
page_title: "pim_settings"
subcategory: ""
description: |-
  This represents the definition of PIM Multicast Routing Settings for a Firewall, including PIM Profile, MRoute Preference, RP Candidate Interface, BSR Candidate Interface, and RP Candidates.
---

# pim_settings (Nested-Attribute)

This represents the definition of PIM Multicast Routing Settings for a Firewall, including PIM Profile, MRoute Preference, RP Candidate Interface, BSR Candidate Interface, and RP Candidates.




## Simple Attributes
- `bsr_priority` (Number) The BSR Priority for multicast routing settings, with a default value of 64.
- `mroute_preference` (String) The MRoute preference for multicast routing, which can be 'best_match_preferred' or 'mroute_preferred'.
- `pim_profile_ref` (String) This represents the PIM IPv4 Profile for Dynamic Routing Firewall functionality. It is used to configure PIM (Protocol Independent Multicast) settings in the firewall's dynamic routing capabilities.
- `rp_priority` (Number) The RP Priority for multicast routing settings, with a default value of 64.

## Nested Attributes
- `bsr_candidate_interface` (Single Block, see [here](attr_bsr_candidate_interface_entry.md)) 
- `rp_candidate_interface` (Single Block, see [here](attr_rp_candidate_interface_entry.md)) 
- `sm_rp_candidate_entry` (List of Blocks, see [here](attr_pim_sm_bootstrap_settings_rp_candidate_entry.md)) The PIM-SM RP Candidate entries for multicast groups.

