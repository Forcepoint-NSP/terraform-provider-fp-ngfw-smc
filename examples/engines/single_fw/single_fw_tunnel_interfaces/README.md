# Single Firewall Tunnel Interfaces Example

This example demonstrates how to configure and attach tunnel interfaces to a single firewall (SMC Engine) node using Terraform. It is intended as a clear implementation reference for provisioning both IPv4 and IPv6 tunnel interfaces, suitable for VPN, site-to-site, dynamic routing, or advanced segmentation scenarios in Forcepoint SMC-managed infrastructures.

## What This Example Does

- **Creates a single firewall engine** named `tf_single_fw_tunnel_interfaces`.
- **Associates the firewall node with an existing log server** (discovered via a wildcard).
- **Provisions interfaces:**
  - One physical interface with a primary management IPv4 address.
  - Two distinct tunnel interfaces:
    - **Tunnel Interface 6483**
      - IPv6 address and /128 network.
      - Example for advanced dynamic routing or IPv6 tunnel endpoints.
    - **Tunnel Interface 1000**
      - IPv4 address and /32 network.
      - Example for virtual or point-to-point link simulation, or as a VPN tunnel endpoint.
- Both tunnel interfaces are mapped to the single firewall node and provide the basis for secure, logical link creation over non-physical media within your SMC topology.

## How Tunnel Interfaces Are Used

Tunnel interfaces enable your firewall to support logical/virtualized network paths separate from physical connections. Common uses include:
- Inter-site VPN/IPSec tunnels
- GRE/IP-in-IP or dynamic routing overlays
- Segmentation for cloud or hybrid environments
- Scalable site-to-site topologies where overlay interfaces are required

