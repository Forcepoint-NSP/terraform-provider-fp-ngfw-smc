# External Gateway

This example demonstrates how to configure an External Gateway within the Forcepoint SMC Terraform provider. 

## What Does This Example Do?

- **Creates two internal firewalls** (`fw1` and `fw2`) with different interfaces and networks, simulating two SMC-managed gateways.
- **Defines an internal VPN topology** using the `smc_vpn` resource, with fw1 as the central gateway and fw2 as an additional internal gateway.
- **Provisions a third-party (external) VPN gateway** using the `smc_external_gateway` resource, representing a remote VPN peer not managed by Forcepoint SMC (for example, a business partner or cloud service).
- **Configures VPN endpoints and sites** for both internal and external gateways, specifying which networks should be reachable over VPN.
- **Integrates both SMC-managed and external networks** into a single secure SD-WAN topology—demonstrating secure connectivity and segmentation between your managed environment and an external network.

This example is useful if you need to interconnect your SMC-managed gateways with an external peer, such as a cloud or partner system, using VPN. It illustrates the configuration of all necessary objects for full end-to-end site-to-site VPN.

## Typical Use Cases
- Connecting a remote office, IaaS cloud environment, or business partner network via a VPN to your SMC-managed SD-WAN estate.
- Demonstrating site-to-site VPN configuration, including both internal and externally-managed (third-party) gateways.
- Securely segmenting traffic between managed and unmanaged environments.

