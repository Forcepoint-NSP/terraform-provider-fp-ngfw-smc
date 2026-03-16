# Central Gateway Node

This is an example demonstrating the usage of the Central Gateway Node configuration. This example provides a practical reference implementation to help you understand how to properly configure and deploy this component in your infrastructure.

## Types of Resources Created

- **smc_network**: Defines IPv4 networks used as local subnets in the SD-WAN topology.
- **smc_router**: Creates routers for each network, acting as gateways between the local subnet and SD-WAN fabric.
- **smc_single_fw**: Deploys single-node firewalls for each site, providing security and traffic control.
- **smc_internal_endpoint**: Configures internal endpoints for VPN connections, enabling secure IPsec tunnels.
- **smc_vpn**: Establishes the VPN configuration, specifying cryptographic profile and topology.
- **smc_gateway_node**: Represents central gateway nodes in the VPN topology, linking firewalls to the VPN and defining their role as central nodes.

## Network Concepts Used

- **SD-WAN (Software-Defined Wide Area Network)**: Simplifies WAN management and optimizes connectivity between distributed sites using software-based controls.
- **Central Gateway Node**: Acts as a hub, aggregating traffic from branches and providing secure connectivity.
- **IPsec VPN**: Encrypts traffic over public networks for secure communication.
- **Firewall**: Enforces security and inspects traffic between local networks and SD-WAN fabric.
- **Router**: Enables routing decisions and connectivity between segments.
- **Network**: Logical subnet for addressing and segmentation.

This example demonstrates how to build a secure, scalable SD-WAN environment using central gateway nodes, firewalls, routers, and VPN endpoints. Each resource illustrates a key concept in modern network architecture, focusing on security, segmentation, and connectivity.
