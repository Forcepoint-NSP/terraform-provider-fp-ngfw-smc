# Site Mode Internal Gateway

This example demonstrates how to configure a route-based VPN using site-mode (network association) between two internal SMC-managed gateways (firewalls).

## What Does This Example Do?
- **Creates two SMC-managed firewall engines**: Sets up two internal gateways for use in the VPN.
- **Defines local and remote networks**: Assigns unique sample networks to each gateway and configures site associations.
- **Configures site-based VPN routing**: Associates networks with the corresponding gateway via site mode, showing network-to-network VPN.
- **Demonstrates route-based VPN**: Illustrates how to use site-based associations for routing over route-based VPN within SMC-managed estate.

## Typical Use Cases
- Deploying SMC-to-SMC VPN for branch or campus environments with explicit network or subnet policy associations.
- Demonstrations requiring clear separation of local and remote network sites across the VPN fabric.

**Note:** Make sure to adapt all addressing and credentials as needed. Intended for demo/learning/reference usage.