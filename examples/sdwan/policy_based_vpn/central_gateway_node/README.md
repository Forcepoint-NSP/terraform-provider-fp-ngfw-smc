# Central Gateway Node

This is an example demonstrating the usage of the Central Gateway Node configuration. This example provides a practical reference implementation to help you understand how to properly configure and deploy this component in your infrastructure.

## Types of Resources Created

- **smc_network**: Defines IPv4 networks used as local subnets in the SD-WAN topology.
- **smc_router**: Creates routers for each network, acting as gateways between the local subnet and SD-WAN fabric.
- **smc_single_fw**: Deploys single-node firewalls for each site, providing security and traffic control.
- **smc_internal_endpoint**: Configures internal endpoints for VPN connections, enabling secure IPsec tunnels.
- **smc_vpn**: Establishes the VPN configuration, specifying cryptographic profile and topology.
- **smc_gateway_node**: Represents central gateway nodes in the VPN topology, linking firewalls to the VPN and defining their role as central nodes.
- **smc_gateway_tunnel**: Adopts the logical tunnel the SMC generates between two gateways and sets its pre-shared key.

## Pre-shared keys

The SMC creates the tunnels between gateways by itself, so `smc_gateway_tunnel` does not create one: it looks the tunnel up by name under `smc_vpn.<name>.link.gateway_tunnel` and updates it. The generated name is `Gateway Tunnel <gw A> - Primary-<gw B> - Primary`, and the order of A and B changes between runs. Match on one gateway with a wildcard (`Gateway Tunnel *tf_single_fw2 - *`) rather than a full name. In a star topology this hits exactly one tunnel per satellite, so a `for_each` over the satellites works the same way.

- The VPN profile used by the VPN must have the capability `pre_shared_key_for_ike = true`. The built-in `Suite-B-GCM-128` used here does.
- A lookup that matches nothing fails with `jmespath search failed: SyntaxError`, which is misleading: check the name pattern.
- The key comes from the sensitive variable `preshared_key`; the default is an example value. Set your own with `TF_VAR_preshared_key`, at most 142 characters.
- The engine reads a key starting with `0x` as hexadecimal and one starting with `0s` as Base64, so a text key must not start with either. A random key may start with one: with a `random_password` resource from the `hashicorp/random` provider, prefix its result with a fixed letter (`"K${random_password.psk.result}"`, `length = 141`, `special = false`). Otherwise the policy upload fails on the engine with `Hexadecimal encoded pre-shared key contains invalid characters`, and no tunnel of the VPN comes up.
- A tunnel between two gateways the SMC manages already has a key the SMC generated, which never starts with `0x` or `0s`. Set one only when the key must be known outside the SMC.
- The key is stored in the Terraform state. Keep the state local or in a protected backend.

## Network Concepts Used

- **SD-WAN (Software-Defined Wide Area Network)**: Simplifies WAN management and optimizes connectivity between distributed sites using software-based controls.
- **Central Gateway Node**: Acts as a hub, aggregating traffic from branches and providing secure connectivity.
- **IPsec VPN**: Encrypts traffic over public networks for secure communication.
- **Firewall**: Enforces security and inspects traffic between local networks and SD-WAN fabric.
- **Router**: Enables routing decisions and connectivity between segments.
- **Network**: Logical subnet for addressing and segmentation.

This example demonstrates how to build a secure, scalable SD-WAN environment using central gateway nodes, firewalls, routers, and VPN endpoints. Each resource illustrates a key concept in modern network architecture, focusing on security, segmentation, and connectivity.
