# GRE Tunnel No Encryption Example

This example demonstrates how to set up a route-based VPN between SMC-managed gateways using a Generic Routing Encapsulation (GRE) tunnel without any encryption.

## What Does This Example Do?
- **Configures two SMC firewalls**: Deploys devices with relevant interfaces for GRE tunneling.
- **Allocates distinct networks**: Assigns example subnets and routing for each endpoint.
- **Establishes an unencrypted GRE tunnel**: Shows how to encapsulate network traffic directly, ideal for trusted or isolated environments where traffic encryption is not a requirement.

## Typical Use Cases
- Building overlays for internal, trusted segments or for performance testing.
- Demonstrating basic encapsulation and routing via GRE without encryption overhead.

**Security Note:** This configuration passes data unencrypted across the GRE tunnel. Only use in secured, isolated, or testing environments. Change all demo addresses and credentials for real deployment.