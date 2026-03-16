variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_proxy_server" "generic_proxy" {
  add_x_forwarded_for = true
  address             = "10.1.1.1"
  comment             = var.resource_comment
  balancing_mode      = "ha"
  fp_proxy_key_id     = 0
  http_proxy          = "generic"
  inspected_service {
    name         = "FTP"
    port         = 21
    service_type = "FTP"
  }
  inspected_service {
    name         = "HTTP"
    port         = 8080
    service_type = "HTTP"
  }
  inspected_service {
    name         = "HTTPS"
    port         = 8080
    service_type = "HTTPS"
  }
  inspected_service {
    name         = "SMTP"
    port         = 25
    service_type = "SMTP"
  }
  inspected_service {
    name         = "Default"
    port         = 8080
    service_type = "Default"
  }
  ip_address        = ["10.1.1.11"]
  name              = "tf_generic_proxy_server"
  trust_host_header = false
}
resource "smc_proxy_server" "redirect_proxy" {
  add_x_forwarded_for = false
  address             = "10.1.1.2"
  balancing_mode      = "srcdst"
  fp_proxy_key_id     = 0
  http_proxy          = "redirect"
  inspected_service {
    name         = "FTP"
    comment      = var.resource_comment
    port         = 21
    service_type = "FTP"
  }
  inspected_service {
    name         = "HTTP"
    port         = 8080
    service_type = "HTTP"
  }
  inspected_service {
    name         = "HTTPS"
    port         = 8080
    service_type = "HTTPS"
  }
  inspected_service {
    name         = "SMTP"
    port         = 25
    service_type = "SMTP"
  }
  inspected_service {
    name         = "Default"
    port         = 8080
    service_type = "Default"
  }
  name              = "tf_proxy_redirect"
  trust_host_header = false
}

resource "smc_proxy_server" "fp_web_cloud_proxy" {
  add_x_forwarded_for = false
  address             = "10.1.1.3"
  balancing_mode      = "src"
  fp_proxy_key        = "SuperPassword00!!"
  fp_proxy_key_hashed = "$6$iIFMVfOLuNtwhtfz$up4y6ZmWBxkmeuc8lnmxn.hYjM7xTiD5qeBmZkWre4yzo6iZGe67E42y0/xZ74HBJiOSnNauYTKqrvpFVU8vr."
  fp_proxy_key_id     = 1
  fp_proxy_user_id    = "12345"
  http_proxy          = "forcepoint_ap-web_cloud"
  inspected_service {
    name         = "Default"
    comment      = var.resource_comment
    port         = 8080
    service_type = "Default"
  }
  name              = "tf_proxy_fp_web_cloud"
  trust_host_header = false
}
