# IP-IP Mode External Gateway

This example demonstrates how to configure a route-based VPN between a SMC-managed internal gateway and an external gateway (representing a third-party device or remote site) using IP-IP encapsulation.

## What Does This Example Do?
- **Provisions an internal SMC-managed firewall**: Sets up configuration and tunnel interfaces for a managed device.
- **Defines an external gateway**: Represents a remote VPN endpoint, such as a third-party firewall, cloud device, or partner.
- **Configures tunnel interfaces and routes**: Prepares sample networks and routers on both sides.
- **Establishes route-based VPN**: Demonstrates IP-IP encapsulated VPN connectivity between SMC and a non-SMC/third-party gateway.

## Typical Use Cases
- Hybrid deployments connecting SMC-managed gateways to non-SMC devices (such as public cloud firewalls or business partner gateways).
- Scenarios requiring explicit route control and lightweight encapsulation with a third-party peer.

**Note:** Always customize addressing and authentication settings for your environment. This example is intended for reference and demonstration usage.