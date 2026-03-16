---
page_title: "smc_routing_node"
subcategory: "routing"
description: |-
  This represents a Routing Node in the Security Management Client, which is used to configure routing for network traffic on firewalls.
---

# smc_routing_node (Sub-resource)

This represents a Routing Node in the Security Management Client, which is used to configure routing for network traffic on firewalls.

## Examples

- [getting_started/main.tf](https://github.com/Forcepoint-NSP/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/routing_node): Declares a routing node for a firewall engine.

This snippet shows use of the `smc_routing_node` resource to define the routing structure for a firewall cluster, referencing interfaces and routes.

```hcl
resource "smc_routing_node" "tf_single_fw" {
  id    = smc_single_fw.tf_single_fw1.link.routing
  level = "engine_cluster"
  routing_node {
    level = "interface"
    name  = "Interface 0"
    routing_node {
      level = "network"
      href  = smc_network.tf_sample_network_single_1.id
      routing_node {
        level = "gateway"
        href  = smc_router.tf_sample_router_1.id
        routing_node {
          level = "any"
          href  = data.smc_href.any_network.id
        }
      }
    }
  }
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `from_ref` (String) parent href of this sub-resource
- `comment` (String) An optional comment for the element. This field is not required.
- `communication_mode` (String) The type of communication mode used for this interface in case of OSPF dynamic routing configuration, such as 'not_forced', 'point_to_point', 'passive', or 'unicast'.
- `custom_configuration` (String) Custom configuration for the routing node, which can be used to specify additional settings or parameters.
- `dynamic_classid` (String) The Dynamic Type of the associated element, such as 'interface', 'modem_interface', 'adsl_interface', 'network', 'gateway', or 'netlink'.
- `dynamic_ipv6` (Boolean) Indicates whether the dynamic element is IPv6 or not. By default, it is not IPv6.
- `dynamic_nicid` (String) The Dynamic Interface ID of the associated Physical Interface, required for dynamic interface levels.
- `exclude_from_ip_counting` (Boolean) Indicates whether the Interface Routing element should be excluded from IP counting, which means it will not be considered in IP address usage calculations.
- `href` (String) This is a base class for objects that can be stored in the database. It includes properties such as key, comment, and links.
- `invalid` (Boolean) Indicates whether the Routing element is valid or not.
- `ip` (String) The IP address of the Routing/Antispoofing, required in case of secondary IP Addresses.
- `level` (String) The level of the Routing/Antispoofing, such as 'engine_cluster', 'interface', 'network', 'gateway', or 'any'.
- `mroute_only` (Boolean) Indicates whether the PIM Any routing level is marked as MRoute only, which means it is used for multicast routing only.
- `name` (String) Name of the object.
- `nic_id` (String) The Interface ID of the associated Physical Interface, required in case of Interface level.
- `probe_ecmp` (Number) The ECMP (Equal-Cost Multi-Path) value for the routing probe, which is used to determine the number of equal-cost paths for load balancing.
- `probe_interval` (Number) The interval in seconds between routing probe tests when the probe test type is 'ping'. This value determines how often the probe will check the reachability of the next hop.
- `probe_ipaddress` (String) The IP address used for the routing probe when the probe test type is 'ping'. This address is used to check the reachability of the next hop.
- `probe_metric` (Number) The metric value for the routing probe, which is used to determine the cost of the route. A lower metric indicates a preferred route.
- `probe_retry_count` (Number) The number of retries for the routing probe when the probe test type is 'ping'. This value determines how many times the probe will attempt to reach the next hop before considering it unreachable.
- `probe_test` (String) The type of probe test used for the routing node, such as 'ping', 'traceroute', or 'not_set'.
- `related_element_type` (String) The type of the related element.
- `to_delete` (Boolean) Indicates whether the Routing element should be deleted or not.

## Nested Attributes
- `routing_node` (List of Blocks, see [here](routing_node.md)) A list of Routing child nodes, which represent individual routing configurations for interfaces.

## Readonly Attributes
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `read_only` (Boolean) Indicates whether the view node is read-only or not.
- `system` (Boolean) Indicates whether the view node is a system node or not.
