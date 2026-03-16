variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "default_policy_template" {
  type = "fw_template_policy"
  name = "Firewall Template"
}

data "smc_href" "interface_nic_0_net" {
  type = "interface_nic_x_net_alias"
  name = "$$ Interface ID 0.net"
}

data "smc_href" "default_nat_address_alias" {
  type = "default_nat_address_alias"
  name = "$$ Default NAT Address"
}

resource "smc_fw_policy" "example" {
  name     = "tf_nat_rule_example_policy"
  template = data.smc_href.default_policy_template.href
  comment  = var.resource_comment
}

resource "smc_fw_ipv4_nat_rule" "nat_rule1" {
  from_ref = smc_fw_policy.example.link.fw_ipv4_nat_rules
  name     = "nat_rule1"

  services {
    any = true
  }
  sources {
    src = [data.smc_href.interface_nic_0_net.href]
  }
  destinations {
    # elements = [data.smc_href.not_internal_network.href]
    any = true
  }
  options {
    dynamic_src_nat {
      automatic_proxy = true
      translation_values {
        element  = data.smc_href.default_nat_address_alias.href
        max_port = 2048
        min_port = 1024
      }
    }
  }
}
