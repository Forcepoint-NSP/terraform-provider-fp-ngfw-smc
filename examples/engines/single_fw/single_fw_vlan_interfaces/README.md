# Single Firewall with VLAN Interfaces Example

This example demonstrates how to deploy a single firewall engine (`single_fw`) in Forcepoint SMC with multiple VLAN sub-interfaces configured on a physical port. It provides a reference implementation for setting up VLAN trunking and multi-segmented firewall connectivity using the SMC Terraform provider.

## What this example does

- **Deploys a single firewall engine** with:
  - A standard management interface on a dedicated physical port for provisioning, monitoring, and management access (interface 0).
  - An additional physical port (`interface_id = 1`) configured to act as a VLAN trunk.

- **Defines VLAN sub-interfaces**:
  - Two VLAN sub-interfaces (`1.2000` and `1.100`), each mapped to the parent physical port.
  - Each VLAN interface can carry its own network segment, distinct addressing scheme, and is attached to the firewall node:
    - `VLAN 1.2000`: Uses IPv6 addressing (`fc00:2000:2000:2000::2/64`)
    - `VLAN 1.100`: Uses IPv4 addressing (`100.100.100.100/24`)

- **Illustrates segmentation and tagging:**
  - By configuring multiple `vlan_interfaces` blocks within a single `physical_interface`, this example shows how to provision one port as a trunk for several logically separated networks (common in data centers and complex topologies).

## Usage Notes

- Adjust VLAN IDs, addresses, or network CIDRs to fit your environment, physical cabling, and tagging practices.
- Each VLAN interface is represented as a sub-interface (`1.<vlan-id>`) under its parent physical port number.
- This pattern can be extended to add more VLAN sub-interfaces or combine with additional physical ports as needed.

## How to use this example

1. **Initialize Terraform:**
   ```sh
   terraform init
   ```
2. **Plan and review changes:**
   ```sh
   terraform plan
   ```
3. **Apply to provision firewall and VLAN interfaces:**
   ```sh
   terraform apply
   ```

