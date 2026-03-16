# IP-IP Mode Internal Gateway

This example demonstrates how to configure a route-based VPN between two SMC-managed internal gateways (firewalls) using IP-IP mode (a simple tunneling protocol) as the transport for the VPN tunnel interface.

## What Does This Example Do?
- **Provisions two internal firewalls**: Sets up and configures two single firewall engines managed by SMC.
- **Assigns local and remote networks**: Creates sample networks and routing for demonstration.
- **Configures tunnel interfaces**: Each firewall is given a tunnel interface for IP-IP encapsulation.
- **Establishes a route-based VPN**: Demonstrates how to use IP-IP mode to connect these two gateways securely using dynamic route-based VPN, rather than policy-based VPN.
- **Reference topology for secure SD-WAN**: Useful for scenarios where you want exact route control and simple encapsulation for VPN links between managed endpoints.

## Typical Use Cases
- Internal (on-premises or hybrid) datacenter or branch-to-branch VPNs using SMC-managed devices.
- Lab, proof-of-concept, or learning scenarios requiring route-based VPN with IP-IP encapsulation between SMC gateways.

**Note:** Adjust network addressing and authentication settings as needed. Example values are for demonstration purposes only.
