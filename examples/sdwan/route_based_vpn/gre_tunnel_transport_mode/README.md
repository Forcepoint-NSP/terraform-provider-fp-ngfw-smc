# GRE Tunnel Transport Mode Example

This example demonstrates the use of GRE tunnel interfaces in transport mode for setting up a route-based VPN between SMC-managed gateways.

## What Does This Example Do?
- **Creates two SMC-managed firewalls:** Establishes tunnel interfaces for each device.
- **Configures GRE in transport mode:** GRE transport mode encapsulates only the payload, providing lower overhead than tunnel mode.
- **Assigns networks and routes:** Sample routing demonstrates connectivity between local and remote subnets over the GRE link.
- **Shows route-based VPN controls:** Useful for overlay topologies or scenarios where custom tunnel optimization is important.

## Typical Use Cases
- Interconnecting datacenters or branches using transport-mode tunnels to reduce protocol overhead.
- Demonstrating the flexibility of SMC in terms of route-based VPN tunnel types.

**Note:** Adjust topology, addresses, and security settings per your environment. Example is for reference/learning.