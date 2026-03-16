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

resource "smc_fw_template_policy" "example" {
  name    = "tf_example_template_policy"
  comment = var.resource_comment
}

### IPv4 Access Rules
resource "smc_fw_ipv4_access_rule" "rule1" {
  from_ref = "${smc_fw_template_policy.example.id}/fw_ipv4_access_rule"
  name     = "allow-ssh-from-any-to-any"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.ssh_service.href]
  }
  action { action = ["allow"] }
  rank = 1.0
}

resource "smc_fw_ipv4_access_rule" "ipv4_automatic_insert_point" {
  from_ref = "${smc_fw_template_policy.example.id}/fw_ipv4_access_rule"
  type     = "automatic"
  name     = "IPv4 Access Rule Automatic Insert Point"
  rank     = 2.0
}

resource "smc_fw_ipv4_access_rule" "ipv4_insert_point" {
  from_ref = "${smc_fw_template_policy.example.id}/fw_ipv4_access_rule"
  type     = "normal"
  name     = "IPv4 Access Rule Insert Point"
  rank     = 3.0
}

#### IPv6 Access
resource "smc_fw_ipv6_access_rule" "ipv6_automatic_insert_point" {
  from_ref = "${smc_fw_template_policy.example.id}/fw_ipv6_access_rule"
  type     = "automatic"
  name     = "IPv6 Access Rule Automatic Insert Point"
  rank     = 1.0
}

resource "smc_fw_ipv6_access_rule" "ipv6_insert_point" {
  from_ref = "${smc_fw_template_policy.example.id}/fw_ipv6_access_rule"
  type     = "normal"
  name     = "IPv6 Access Rule Insert Point"
  rank     = 2.0
}

#### IPv4 NAT
resource "smc_fw_ipv4_access_rule" "ipv4_insert_point_nat" {
  from_ref = "${smc_fw_template_policy.example.id}/fw_ipv4_nat_rule"
  type     = "normal"
  name     = "IPv4 Access Rule Insert Point - NAT"
  rank     = 1.0
}

#### IPv6 NAT
resource "smc_fw_ipv6_access_rule" "ipv6_insert_point_nat" {
  from_ref = "${smc_fw_template_policy.example.id}/fw_ipv6_nat_rule"
  type     = "normal"
  name     = "IPv6 Access Rule Insert Point - NAT"
  rank     = 1.0
}