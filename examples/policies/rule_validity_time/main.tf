variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_rule_validity_time" "tf_rule_validity_time" {
  name                   = "tf_rule_validity_time"
  rule_time_active       = "time_of_day"
  rule_time_start        = "2024-01-01 08:00"
  rule_time_end          = "2024-12-31 17:00"
  rule_time_repeat_start = "08:00"
  rule_time_repeat_end   = "17:00"
  comment                = var.resource_comment
}
