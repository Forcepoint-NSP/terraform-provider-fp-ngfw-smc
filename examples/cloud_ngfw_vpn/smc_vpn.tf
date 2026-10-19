# Policy-based (IPsec) site-to-site VPN between the two engines.
# Central gateway = site A, satellite gateway = site B. Each side's LAN is
# published as its VPN site, and IPsec is enabled on each WAN endpoint.

resource "smc_vpn" "site_to_site" {
  name                     = "cloud-site-to-site"
  mobile_vpn_topology_mode = "none"
  nat                      = false
  vpn_profile              = data.smc_href.suite_b_gcm_128.href
  comment                  = var.resource_comment
}

# --- contact addresses on the WAN interfaces ----------------------------
# Advertise each engine's mocked public IP so the peer resolves the tunnel
# endpoint to the "public" address rather than the private WAN IP.

data "smc_sub_href" "wan_intf" {
  for_each = local.sites
  from_ref = smc_single_fw.ngfw[each.key].link.physical_interface
  match    = format(local.contact_addr_match, "Interface 1", local.wan_ip[each.key])
}

resource "smc_contact_address" "wan" {
  for_each = local.sites
  id       = data.smc_sub_href.wan_intf[each.key].id
  contact_addresses {
    address      = each.value.public_ip
    dynamic      = false
    location_ref = data.smc_href.default_location.id
  }
}

# --- internal gateways + IPsec endpoints --------------------------------

data "smc_sub_href" "internal_gateway" {
  for_each = local.sites
  from_ref = smc_single_fw.ngfw[each.key].link.internal_gateway
  match    = "${each.value.name}*"
}

# Enable IPsec on each engine's WAN endpoint (named after its WAN IP).
resource "smc_internal_endpoint" "wan" {
  for_each = local.sites
  id = format(
    "%s#%s",
    smc_single_fw.ngfw[each.key].link.internal_gateway,
    "*/internal_endpoint/${local.wan_ip[each.key]}",
  )
  enabled   = true
  ipsec_vpn = true
}

# --- gateway nodes ------------------------------------------------------

resource "smc_gateway_node" "central" {
  from_ref   = smc_vpn.site_to_site.link.central_gateway_node
  name       = "${local.sites["a"].name} central"
  node_usage = "central"
  gateway    = data.smc_sub_href.internal_gateway["a"].id
}

resource "smc_gateway_node" "satellite" {
  from_ref   = smc_vpn.site_to_site.link.satellite_gateway_node
  name       = "${local.sites["b"].name} satellite"
  node_usage = "satellite"
  gateway    = data.smc_sub_href.internal_gateway["b"].id
}

# --- protected networks published on each gateway (vpn_site) -------------

resource "smc_vpn_site" "site" {
  for_each     = local.sites
  from_ref     = format("%s#%s", smc_single_fw.ngfw[each.key].link.internal_gateway, "*/vpn_site")
  name         = "${each.value.name} site"
  site_element = [smc_network.lan[each.key].id]
}
