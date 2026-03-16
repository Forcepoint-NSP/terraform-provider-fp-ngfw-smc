# FW Cluster with CVI + NDI Interface Example

This example shows how to configure a firewall cluster (two nodes)
with both Cluster Virtual Interface (CVI) and Node Dedicated
Interfaces (NDI) using Terraform and the Forcepoint SMC provider.

- **Creates a balancing-mode firewall cluster with 2 nodes**
- **Configures one physical interface (Interface 0):**
  - CVI: Shared cluster address
  - NDI: Unique address per node
  - Attached to the "Internal" zone
  - Management and heartbeat enabled on node interfaces

