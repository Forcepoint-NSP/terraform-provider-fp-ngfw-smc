variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ospfv2_key_chain" "ospfv2_key_chain" {
  name    = "tf_ospfv2_key_chain"
  comment = var.resource_comment
  ospfv2_key_chain_entry {
    algorithm = "md5"
    comment   = "md5"
    key_id    = 1
    send_key  = true
    key       = "123"
  }
  ospfv2_key_chain_entry {
    algorithm = "hmac-sha-512"
    comment   = ""
    key_id    = 2
    send_key  = false
    key       = "564"
  }
}
