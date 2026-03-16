# Bgp Peering


This example demonstrates how to configure BGP (Border Gateway Protocol) peering using the Forcepoint SMC Terraform provider. It creates and configures several resources related to dynamic routing and BGP filtering.

## Purpose of Each Resource

- **IP Access List**: Controls which IP prefixes are allowed or denied in BGP updates. For example, it can block private or unauthorized subnets from being advertised or accepted.
- **AS Path Access List**: Filters BGP routes based on the AS path, allowing or denying routes that match certain AS path patterns. This helps enforce routing policies and prevent route leaks.
- **BGP Connection Profile**: Centralizes BGP session settings (timers, authentication) for reuse across multiple peerings, ensuring consistency and simplifying management.
- **BGP Peering**: Represents the actual BGP session with a peer, applying the defined filters and connection profile. It manages advanced BGP options such as route reflection, next-hop self, and soft reconfiguration.

## Network Concepts Used

- **BGP (Border Gateway Protocol)**: A dynamic routing protocol used to exchange routing information between autonomous systems (ASes) on the internet or within large enterprise networks.
- **AS Path Filtering**: The process of allowing or denying BGP routes based on the sequence of AS numbers a route advertisement has traversed. Used for policy enforcement and security.
- **IP Prefix Filtering**: The process of controlling which network prefixes are advertised or accepted in BGP updates, improving security and route control.
- **Route Reflection**: A BGP feature that allows a router to reflect routes to other BGP peers, reducing the number of peerings required in large networks.
- **BGP Session Parameters**: Settings such as hold timers, keepalive intervals, and authentication, which control the stability and security of BGP sessions.

This example can be used as a reference for implementing BGP peering with filtering and advanced options in a Forcepoint SMC environment.
