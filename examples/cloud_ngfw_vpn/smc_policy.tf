# Shared north-south firewall policy for both engines, plus the protected
# LAN network elements the rules and VPN sites reference.

resource "smc_network" "lan" {
  for_each     = local.sites
  name         = "net-${each.value.name}-lan"
  ipv4_network = each.value.lan_cidr
  comment      = "${each.value.name} protected subnet"
}

resource "smc_fw_policy" "north_south" {
  name     = "cloud-ngfw-north-south"
  template = data.smc_href.default_policy_template.id
  comment  = var.resource_comment
}

# Rule order is controlled by `rank` (lower = evaluated earlier). The
# site-to-site rule must rank above the internet rule: inter-site traffic
# also matches the internet rule's dst=any, so if that were evaluated first
# the LAN-to-LAN flow would leave in the clear and never enter the VPN.

# Rank 1 - LAN-to-LAN between the two sites, carried over the IPsec VPN.
resource "smc_fw_ipv4_access_rule" "site_to_site" {
  from_ref = smc_fw_policy.north_south.link.fw_ipv4_access_rules
  name     = "site-to-site VPN"
  sources {
    src = [for k in keys(local.sites) : smc_network.lan[k].id]
  }
  destinations {
    dst = [for k in keys(local.sites) : smc_network.lan[k].id]
  }
  services { any = true }
  action {
    action = ["allow", "apply_vpn"]
    vpn    = smc_vpn.site_to_site.id
  }
  rank = 1
}

# Rank 2 - internet egress for both LAN subnets.
resource "smc_fw_ipv4_access_rule" "allow_internet" {
  from_ref = smc_fw_policy.north_south.link.fw_ipv4_access_rules
  name     = "allow internet traffic"
  sources {
    src = [for k in keys(local.sites) : smc_network.lan[k].id]
  }
  destinations { any = true }
  services { any = true }
  action { action = ["allow"] }
  rank = 2
}

# Rank 3 - SSH from anywhere.
resource "smc_fw_ipv4_access_rule" "allow_ssh" {
  from_ref = smc_fw_policy.north_south.link.fw_ipv4_access_rules
  name     = "allow ssh"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.service_ssh.id]
  }
  action { action = ["allow"] }
  rank = 3
}

# Rank 100 - explicit default-deny tail.
resource "smc_fw_ipv4_access_rule" "discard_all" {
  from_ref = smc_fw_policy.north_south.link.fw_ipv4_access_rules
  name     = "discard all"
  sources { any = true }
  destinations { any = true }
  services { any = true }
  action { action = ["discard"] }
  rank = 100
}
