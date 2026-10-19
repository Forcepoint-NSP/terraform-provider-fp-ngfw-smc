# WAN routing for each engine: reach the remote site's public endpoint and
# its LAN through the local WAN gateway. Only the WAN interface carries
# declared routes; the connected mgmt and LAN subnets are handled by SMC.

# Connected WAN network (routing anchor) and its upstream gateway router.
resource "smc_network" "wan" {
  for_each     = local.sites
  name         = "net-${each.value.name}-wan"
  ipv4_network = each.value.wan_cidr
  comment      = "${each.value.name} WAN subnet"
}

resource "smc_router" "wan_gw" {
  for_each = local.sites
  name     = "router-${each.value.name}-wan-gw"
  address  = local.wan_gw_ip[each.key]
  comment  = "${each.value.name} WAN gateway"
}

# The peer's public endpoint as a host element (SMC rejects a /32 network).
resource "smc_host" "public" {
  for_each = local.sites
  name     = "host-${each.value.name}-public"
  address  = each.value.public_ip
  comment  = "${each.value.name} mocked public WAN address (VPN peer endpoint)"
}

# Interface 1 (WAN) routes toward the remote site: its public VPN endpoint
# and its protected LAN, both via the local WAN gateway. These static routes
# must be in the engine config before the policy is installed, so the policy
# upload (smc_policy_upload.tf) depends on this resource.
resource "smc_routing_node" "wan" {
  for_each = local.sites
  id       = smc_single_fw.ngfw[each.key].link.routing
  level    = "engine_cluster"

  routing_node {
    level = "interface"
    name  = "Interface 1"
    routing_node {
      level = "network"
      href  = smc_network.wan[each.key].id
      routing_node {
        level = "gateway"
        href  = smc_router.wan_gw[each.key].id
        routing_node {
          level = "any"
          href  = smc_host.public[each.value.peer].id
        }
        routing_node {
          level = "any"
          href  = smc_network.lan[each.value.peer].id
        }
      }
    }
  }
}
