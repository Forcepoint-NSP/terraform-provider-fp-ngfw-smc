---
page_title: "smc_rbvpn_tunnel_side"
subcategory: "vpn"
description: |-
  This defines one part of a Route-Based Tunnel, which can be a gateway, a gateway with its EndPoint, or just an external IP Address. None of the elements are mandatory, as the configuration of the Route-Based Tunnel determines which elements are required.
---

# smc_rbvpn_tunnel_side (Resource)


⚠️ __Experimental feature, use with caution__. This feature is not yet fully supported and may change without deprecation in future releases.
This defines one part of a Route-Based Tunnel, which can be a gateway, a gateway with its EndPoint, or just an external IP Address. None of the elements are mandatory, as the configuration of the Route-Based Tunnel determines which elements are required.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `endpoint_ref` (String) This represents an abstract endpoint element in the VPN Gateway, which can be extended to represent specific types of endpoints.
- `gateway_ref` (String) This is the base class for all storable elements.
- `ip_address` (String) The external IP address used in the Route-Based VPN Tunnel Side.
- `tunnel_interface_ref` (String) This represents a Tunnel Interface, which is used for defining endpoints for tunnels in the Route-Based VPN. It allows traffic to be routed into the tunnel based on Firewall Access rules.


