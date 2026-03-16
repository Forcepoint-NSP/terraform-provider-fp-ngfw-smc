# FW Cluster with Layer 2 Interface Example

This example configures a firewall cluster (2 nodes) with both standard (L3) and Layer 2 interfaces using Terraform and the Forcepoint SMC provider.

- **Defines two logical interfaces:** one for inline IPS, one for capture
- **Creates a cluster with two nodes in balancing mode**
- **Configures:**
  - Standard interface (with CVI/NDI): management & data
  - Inline IPS (layer 2) interface: bypass failure mode, Q-in-Q, VLAN inspection
  - Capture interface: for passive monitoring, Q-in-Q and VLAN inspection

