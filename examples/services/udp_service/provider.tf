terraform {
  required_providers {
    smc = {
      source  = "forcepoint-nsp/fp-ngfw-smc"
      version = "1.752.0"
    }
  }
}

variable "api_key" {
  type    = string
  default = null
}

variable "url" {
  type    = string
  default = "http://localhost:8082"
}

variable "api_version" {
  type    = string
  default = null
}

variable "trusted_cert" {
  type    = string
  default = null
}

variable "verify_ssl" {
  type    = bool
  default = false
}

# Overrides the default interval (15 minutes) after which the SMC API
# session token is proactively refreshed. Left null for every topology
# that doesn't need it.
variable "token_refresh_interval_seconds" {
  type    = number
  default = null
}

provider "smc" {
  url                            = var.url
  api_key                        = var.api_key
  api_version                    = var.api_version
  trusted_cert                   = var.trusted_cert
  verify_ssl                     = var.verify_ssl
  token_refresh_interval_seconds = var.token_refresh_interval_seconds
}
