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

resource "smc_fw_ipv4_access_rule" "rule1" {
  from_ref = "${smc_fw_policy.example.id}/fw_ipv4_access_rule"
  name     = "allow-ssh-from-any-to-any"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.ssh_service.href]
  }
  action { action = ["allow"] }
}
