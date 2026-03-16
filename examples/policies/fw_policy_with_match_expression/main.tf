variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "user_password_auth_method" {
  name = "User password"
  type = "authentication_service"
}

resource "smc_network" "network_1" {
  broadcast    = true
  comment      = var.resource_comment
  ipv4_network = "192.168.10.0/24"
  name         = "tf_network_1"
}

resource "smc_host" "host_1" {
  comment = var.resource_comment
  address = "192.168.20.1"
  name    = "tf_host_1"
}

resource "smc_match_expression" "match_expression_host1_net1" {
  name = "tf_match_expression_host1_net1"
  ref  = [smc_network.network_1.id, smc_host.host_1.id]
}

resource "smc_internal_user" "todd" {
  authentication_method = [data.smc_href.user_password_auth_method.id]
  days_left             = -1
  name                  = "tf_user"
  password              = "SuperSecretPassword"
  unique_id             = "cn=tf_user,dc=stonegate,domain=InternalDomain"
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
  name     = "allow-ssh-from-match-expression-to-any"
  sources { src = [smc_match_expression.match_expression_host1_net1.id, smc_internal_user.todd.id] }
  destinations { any = true }
  services {
    service = [data.smc_href.ssh_service.href]
  }
  action { action = ["allow"] }
}
