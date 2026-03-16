# Routing Node Example

This example demonstrates how to configure explicit, hierarchical routing on a Forcepoint SMC firewall engine using Terraform. It provides a reference for creating networks, routers (gateways), and assigning intricate routing logic to interfaces using the SMC provider's `smc_routing_node` resource.

## What This Example Does

- **Creates two test networks**:
  - `192.168.100.0/24` (`tf_sample_network`)
  - `192.168.101.0/24` (`tf_sample_network2`)
- **Defines two routers (gateways)** at `.254` of each network.
- **Provisions a single firewall (engine) node** with a management interface.
- **Assigns a routing hierarchy**:
  - Starts at engine cluster level.
  - Drills down to a specific interface.
  - Specifies routing via the defined network.
  - Sets a gateway router for that network.
  - Completes with an 'any network' destination scenario.

