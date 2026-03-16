variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "https_service" {
  name = "HTTPS"
  type = "tcp_service"
}

data "smc_href" "AOL_service" {
  name = "AOL"
  type = "tcp_service"
}

data "smc_href" "default_policy_template" {
  type = "fw_template_policy"
  name = "Firewall Template"
}

resource "smc_fw_policy" "example" {
  name     = "tf_example_access_rules_policy_with_rule_section"
  template = data.smc_href.default_policy_template.href
  comment  = var.resource_comment
}

resource "smc_fw_ipv4_access_rule" "rule_section_allow_https" {
  from_ref         = smc_fw_policy.example.link.fw_ipv4_access_rules
  comment          = "Allow HTTPS"
  background_color = "#5EF527"
  rank             = 1.0
}

resource "smc_fw_ipv4_access_rule" "allow_https" {
  from_ref = smc_fw_policy.example.link.fw_ipv4_access_rules
  name     = "Rule 2"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.https_service.href]
  }
  action { action = ["allow"] }
  rank = 2.0
}

resource "smc_fw_ipv4_access_rule" "rule_section_discard_aol" {
  from_ref         = smc_fw_policy.example.link.fw_ipv4_access_rules
  comment          = "Discard AOL"
  background_color = "#3827F5"
  rank             = 3.0
}

resource "smc_fw_ipv4_access_rule" "discard_aol" {
  from_ref = smc_fw_policy.example.link.fw_ipv4_access_rules
  name     = "Rule 2"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.AOL_service.href]
  }
  action { action = ["discard"] }
  rank = 4.0
}
