variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_network" "network_1" {
  broadcast    = true
  comment      = var.resource_comment
  ipv4_network = "192.168.10.0/24"
  name         = "Network 1"
}

resource "smc_network" "network_2" {
  broadcast    = true
  comment      = var.resource_comment
  ipv4_network = "192.168.20.0/24"
  name         = "Network 2"
}

data "smc_href" "ssh_service" {
  name = "SSH"
  type = "tcp_service"
}

data "smc_href" "ftp_service" {
  name = "FTP"
  type = "tcp_service"
}

data "smc_href" "default_policy_template" {
  type = "fw_template_policy"
  name = "Firewall Template"
}

resource "smc_sub_ipv4_fw_policy" "sub_policy" {
  name    = "tf_example_subpolicy"
  comment = var.resource_comment
}

resource "smc_fw_ipv4_access_rule" "rule1" {
  from_ref = "${smc_sub_ipv4_fw_policy.sub_policy.id}/fw_ipv4_access_rule"
  name     = "allow-ftp-from-any-to-any"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.ftp_service.href]
  }
  action { action = ["allow"] }
}

resource "smc_fw_policy" "example" {
  name     = "tf_example_policy"
  template = data.smc_href.default_policy_template.href
  comment  = var.resource_comment
}

resource "smc_fw_ipv4_access_rule" "use_jump_rule" {
  from_ref = "${smc_fw_policy.example.id}/fw_ipv4_access_rule"
  name     = "jump-rule"
  sources { src = [smc_network.network_1.id] }
  destinations { dst = [smc_network.network_2.id] }
  services {
    service = [data.smc_href.ssh_service.href]
  }
  action {
    action     = ["jump"]
    sub_policy = smc_sub_ipv4_fw_policy.sub_policy.id
  }
}
