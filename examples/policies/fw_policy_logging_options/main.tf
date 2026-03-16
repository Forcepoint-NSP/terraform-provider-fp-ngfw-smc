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

data "smc_href" "ids_alert" {
  name = "Default alert"
  type = "ids_alert"
}

resource "smc_fw_policy" "example" {
  name     = "tf_example_policy"
  template = data.smc_href.default_policy_template.href
  comment  = var.resource_comment
}

resource "smc_fw_ipv4_access_rule" "rule1" {
  from_ref = "${smc_fw_policy.example.id}/fw_ipv4_access_rule"
  name     = "allow-ssh-from-any-to-any-with-logging"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.ssh_service.href]
  }
  action { action = ["allow"] }
  options {
    log_accounting_info_mode = false
    log_alert                = data.smc_href.ids_alert.href
    log_closing_mode         = true
    log_compression          = "off"
    log_level                = "alert"
    log_payload_excerpt      = false
    log_payload_record       = false
    log_severity             = 7
  }
}

resource "smc_fw_ipv4_access_rule" "rule2" {
  from_ref = "${smc_fw_policy.example.id}/fw_ipv4_access_rule"
  name     = "allow-ssh-from-any-to-any-with-action-options-logging"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.ssh_service.href]
  }
  action {
    action                                 = ["allow"]
    decrypting                             = true
    deep_inspection                        = true
    file_filtering                         = true
    network_application_latency_monitoring = "probing"
    scan_detection                         = "undefined"
    snort                                  = true
  }
}