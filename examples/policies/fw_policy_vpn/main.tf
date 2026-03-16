variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "ssh_service" {
  name = "SSH"
  type = "tcp_service"
}

data "smc_href" "default_policy_template" {
  type = "fw_template_policy"
  name = "Firewall Template"
}

resource "smc_fw_policy" "example" {
  name     = "tf_example_policy"
  template = data.smc_href.default_policy_template.href
  comment  = var.resource_comment
}

resource "smc_vpn_profile" "tf_keepalive_vpn_profile" {
  name       = "tf_keepalive_vpn_profile"
  keep_alive = true
}

# Create VPN with a specific profile to keep alive tunnels
resource "smc_vpn" "policy_vpn" {
  name        = "tf_vpn"
  nat         = false
  vpn_profile = smc_vpn_profile.tf_keepalive_vpn_profile.id
}

resource "smc_fw_ipv4_access_rule" "vpn_rule" {
  from_ref = "${smc_fw_policy.example.id}/fw_ipv4_access_rule"
  name     = "vpn-rule"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.ssh_service.href]
  }
  action {
    action = ["allow", "apply_vpn"]
    vpn    = smc_vpn.policy_vpn.id
  }
}

resource "smc_fw_ipv4_access_rule" "source_vpn_rule" {
  from_ref = "${smc_fw_policy.example.id}/fw_ipv4_access_rule"
  name     = "source-vpn-rule"
  sources { any = true }
  destinations { any = true }
  services { any = true }
  action {
    action = ["allow"]
  }
  match_vpn_options {
    match_type = "normal"
    match_vpns = [smc_vpn.policy_vpn.id]
  }
}
