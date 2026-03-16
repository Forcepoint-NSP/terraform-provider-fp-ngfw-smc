
# data "smc_host" "dns-google" {
#   name = "DNS Google"
# }

# data "smc_host" "all_router_pims" {
#     name = "ALL-PIM-ROUTERS"
# }

data "smc_href" "host_routers" {
  name = "*Router*"
  type = "host"
}

output "host_routers_len" {
  value = length(data.smc_href.host_routers)
}
# output "all_router_pims_len" {
#   value = length(data.smc_host.all_router_pims)
# }

# output "dns-google" {
#   value = data.smc_host.dns-google.address
# }
