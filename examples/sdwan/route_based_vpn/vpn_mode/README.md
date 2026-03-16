# VPN Mode Example

This example demonstrates how to provision a basic route-based VPN using the 'vpn_mode' in the Forcepoint SMC Terraform provider.

## What Does This Example Do?
- **Deploys two SMC-managed firewalls**: Fully configures devices and relevant interfaces for VPN.
- **Defines VPN tunnel groups**: Creates the route-based VPN group that provides structure for individual tunnels.
- **Creates sample networks and routers**: Assigns test networks to each firewall.
- **Demonstrates a route-based approach**: Unlike policy-based VPN, explicitly sets up tunnel interfaces for traffic steering and dynamic routing across the VPN link.

## Typical Use Cases
- Deploy a simple route-based VPN between SMC-managed sites or datacenters.
- Build a proof-of-concept or test changes in SMC VPN provisioning.

Route-based VPNs offer more flexible network design than policy-based (static) VPNs, making this example useful for learning or as a basis for customized deployments.

**Note:** Adapt the networks, addresses, and secret credentials for your own environment before actual use.