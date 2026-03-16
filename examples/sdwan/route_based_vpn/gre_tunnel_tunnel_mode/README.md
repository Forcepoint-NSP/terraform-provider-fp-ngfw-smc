# GRE Tunnel Tunnel Mode Example

This example demonstrates the configuration of a route-based VPN between SMC-managed gateways using GRE tunnel mode as the transport mechanism.

## What Does This Example Do?
- **Provisions two SMC-managed firewall engines:** Sets up both endpoints of the VPN.
- **Creates tunnel interfaces:** Assigns GRE tunnel interfaces to each gateway.
- **Allocates networks:** Sample subnets and routing illustrate tunneling between sites.
- **Uses GRE in tunnel mode:** Encapsulates full IP packets, allowing easy extension of routing domains across the VPN link.

## Typical Use Cases
- Bridging networks where flexible overlay tunnels are required.
- Demonstrating alternative tunnel modes available within SMC-based route VPNs.

**Note:** Tunnel mode supports multiprotocol and is ideal for extending network reachability. Update IP addressing and credentials for your needs.