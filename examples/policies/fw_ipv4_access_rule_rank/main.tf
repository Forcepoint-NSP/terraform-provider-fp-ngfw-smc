variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "ssh_service" {
  name = "SSH"
  type = "tcp_service"
}

data "smc_href" "https_service" {
  name = "HTTPS"
  type = "tcp_service"
}

data "smc_href" "ftp_service" {
  name = "FTP"
  type = "tcp_service"
}

data "smc_href" "ping_service" {
  name = "Ping"
  type = "icmp_service_group"
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
  name     = "tf_example_access_rules_policy_with_rank"
  template = data.smc_href.default_policy_template.href
  comment  = var.resource_comment
}

resource "smc_fw_ipv4_access_rule" "allow_ssh" {
  from_ref = smc_fw_policy.example.link.fw_ipv4_access_rules
  name     = "Rule 3"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.ssh_service.href]
  }
  action { action = ["allow"] }
  rank = 3.0
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

resource "smc_fw_ipv4_access_rule" "allow_ftp" {
  from_ref = smc_fw_policy.example.link.fw_ipv4_access_rules
  name     = "Rule 5"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.ftp_service.href]
  }
  action { action = ["discard"] }
  rank = 5.0
}

resource "smc_fw_ipv4_access_rule" "allow_icmp" {
  from_ref = smc_fw_policy.example.link.fw_ipv4_access_rules
  name     = "Rule 4"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.ping_service.href]
  }
  action { action = ["allow"] }
  rank = 4.0
}
resource "smc_fw_ipv4_access_rule" "allow_aol" {
  from_ref = smc_fw_policy.example.link.fw_ipv4_access_rules
  name     = "Rule 1"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.AOL_service.href]
  }
  action { action = ["allow"] }
  rank = 1.0
}

