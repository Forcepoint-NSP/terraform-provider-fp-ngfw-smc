# BGP Extended Community Access List Example

This example demonstrates how to create a BGP extended community
access list using the `smc_extended_community_access_list` resource
from the Forcepoint NGFW SMC Terraform provider.

## What is an Extended Community Access List?

BGP (Border Gateway Protocol) extended communities are a feature that
allows additional route tagging in BGP, enabling more flexible and
granular control over route selection and propagation. 

Extended community lists can match, filter, and set complex community
attributes on BGP routes, often used in service provider and advanced
enterprise networks for route policy control.

## What This Example Does
- **Creates an extended community access list** named `tf_extended_community_access_list`.
- **Type:** `expanded`, supporting regex and rich matching for communities.
- **Entry:**
  - `action`: `permit` – allows matching routes.
  - `community`: `.*1545` – regex to match any BGP extended community ending with `1545`.
  - `rank`: `0.0` – priority of the entry.
- **Comment:** Description is set via a variable (defaults to "Created by Terraform").

## Typical Use Case

Extended community access lists like this are used to:
- Filter or match BGP routes based on extended community values.
- Enforce routing policies in route-maps or policy statements.
- Tag or restrict routes during redistribution or propagation between BGP peers.

For instance, you may permit specific routes with certain community tags to be accepted from, or advertised to, a peer.

For more detail on supported parameters, see the [provider documentation](https://registry.terraform.io/providers/forcepoint/fp-ngfw-smc/latest/docs/resources/extended_community_access_list).
