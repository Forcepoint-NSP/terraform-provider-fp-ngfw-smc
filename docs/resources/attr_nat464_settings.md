---
page_title: "ipv6_transition_mechanism"
subcategory: ""
description: |-
  This represents the IPv6 Transition Mechanism Settings for the Engine Cluster, including CLAT, NAT64, and SIIT EAM modes.
---

# ipv6_transition_mechanism (Nested-Attribute)

This represents the IPv6 Transition Mechanism Settings for the Engine Cluster, including CLAT, NAT64, and SIIT EAM modes.




## Simple Attributes
- `clat_local_ipv6_prefix_ip` (String) The local IPv6 prefix address for 464XLAT CLAT mode, which is not required.
- `clat_local_ipv6_prefix_ref` (String) This represents a Network, which is a Network Element that represents a (sub)network of computers. It includes attributes for IPv4 and IPv6 networks, broadcast address, and validation patterns.
- `clat_remote_ipv6_prefix_ip` (String) The remote IPv6 prefix address for 464XLAT CLAT mode, which is not required.
- `clat_remote_ipv6_prefix_ref` (String) This represents a Network, which is a Network Element that represents a (sub)network of computers. It includes attributes for IPv4 and IPv6 networks, broadcast address, and validation patterns.
- `mode` (String) The mode for the IPv6 Transition Mechanism, which can be CLAT, NAT64, or SIIT EAM.
- `nat64_ipv6_prefix_ip` (String) The IPv6 prefix address for NAT64 mode, which is not required.
- `nat64_ipv6_prefix_ref` (String) This represents a Network, which is a Network Element that represents a (sub)network of computers. It includes attributes for IPv4 and IPv6 networks, broadcast address, and validation patterns.
- `nat64_local_pools_ref` (List of String) URI of the Network representing the local IPv4 pools in NAT64 mode.
- `siit_eam_ipv6_pool_ip` (String) The default IPv6 pool address for SIIT EAM mode, which is not required.
- `siit_eam_ipv6_pool_ref` (String) This represents a Network, which is a Network Element that represents a (sub)network of computers. It includes attributes for IPv4 and IPv6 networks, broadcast address, and validation patterns.

## Nested Attributes
- `nat64_mapping` (List of Blocks, see [here](attr_nat64_mapping.md)) The static mappings for NAT64 mode, which is not required.
- `siit_eam_mapping` (List of Blocks, see [here](attr_siit_mapping.md)) The mappings for SIIT EAM mode, which is not required.

