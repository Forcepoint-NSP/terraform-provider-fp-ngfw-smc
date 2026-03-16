# Geneve Mode Example

This example demonstrates how to set up a route-based VPN between SMC-managed devices using a Geneve tunnel as the transport mechanism.

## What Does This Example Do?
- **Creates two internal SMC-managed firewalls**: Deploys and configures two devices for VPN testing.
- **Assigns networks**: Allocates local and remote sample subnets to each gateway.
- **Configures Geneve tunnels**: Utilizes the Geneve protocol to encapsulate traffic over the VPN tunnel interface, demonstrating flexibility for overlay transport.
- **Demonstrates route-based VPN with overlays**: Ideal for cloud, datacenter, or SD-WAN overlays that need dynamic, non-IPsec encapsulation.

## Typical Use Cases
- Securely connecting network segments with flexible tunnel overlays.
- Building POCs for new tunnel types in service-provider or hybrid environments.

**Note:** Geneve support and compatibility should be validated in your environment. Update all addresses and credentials per your use case.