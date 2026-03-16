---
page_title: "smc_gateway_node"
subcategory: "vpn"
description: |-
  This represents a gateway node, which is used to manage the VPN topology and its nodes, including their usage and relationships.
---

# smc_gateway_node (Sub-resource)

This represents a gateway node, which is used to manage the VPN topology and its nodes, including their usage and relationships.

## Examples

- [external_gateway example](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.2/examples/sdwan/external_gateway/main.tf): Adds gateway nodes (central) in a policy-based VPN SD-WAN scenario.

This example illustrates the configuration of `smc_gateway_node`
resources for central gateways in a VPN configuration. Each gateway
node is linked to a VPN resource and references an internal gateway
for SD-WAN.

```hcl
resource "smc_gateway_node" "tf_sample_central_gateway" {
  from_ref   = smc_vpn.tf_sample_vpn.link.central_gateway_node
  name       = "tf_sample_central_gateway"
  node_usage = "central"
  gateway    = data.smc_sub_href.internal_gateway["fw1"].id
}

resource "smc_gateway_node" "tf_sample_satellite_gateway2" {
  from_ref   = smc_vpn.tf_sample_vpn.link.satellite_gateway_node
  name       = "tf_sample_satellite_external_gw"
  node_usage = "satellite"
  gateway    = smc_external_gateway.external_gateway.id
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `child_node` (List of String) URI of the child Gateway Node.
- `comment` (String) An optional comment for the element. This field is not required.
- `gateway` (String) This is the base class for all storable elements.
- `name` (String) Name of the object.
- `node_usage` (String) The usage type of the gateway node, indicating its role in the VPN topology, such as 'central', 'satellite', or 'mobile'.
- `parent_node` (String) This represents a gateway node, which is used to manage the VPN topology and its nodes, including their usage and relationships.
- `vpn_key` (Number) The unique identifier for the VPN node, used to reference this node in the VPN topology.


## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
