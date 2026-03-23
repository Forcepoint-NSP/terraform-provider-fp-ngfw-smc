## 1.741.0 (2026-03-23)

Targets SMC API **7.4 / 7.4.1**. First production-ready release covering:

**Engines**
- `smc_single_fw` – single firewall with physical, VLAN, tunnel, layer-2, dynamic, modem, wireless, and VPN-broker interfaces; contact address; custom domain
- `smc_fw_cluster` – firewall cluster with CVI/NDI, layer-2, tunnel, and VPN-broker interfaces
- `smc_routing_node` / `smc_routing_node` with netlinks and tunnels
- `smc_location`

**Dynamic Routing**
- BGP: autonomous system, BGP profile, connection profile, peering, external BGP peers, AS-path / community / extended-community access lists
- OSPFv2: domain settings, areas, interface settings, key chain, profile
- PIM: profiles, IPv4 interface settings, IGMP querier settings
- Shared: IP / IPv6 access lists, IP / IPv6 prefix lists, route maps

**Network Elements**
- Host, network, address range, router, group, alias, domain name, expression, IP list, interface zone, netlink
- Servers: Active Directory, DHCP, Elasticsearch cluster, external DNS, ICAP, LDAP, NTP, proxy, RADIUS, SMTP, TACACS+

**Policies**
- `smc_fw_policy`, `smc_fw_template_policy`, `smc_sub_ipv4_fw_policy`
- IPv4 access rules (rank, sections), IPv4 NAT rules
- Policy options: logging, VPN, match expressions

**SD-WAN / VPN**
- External gateway, internal endpoint
- Policy-based VPN: central gateway, satellite gateway
- Route-based VPN: GRE (no-encryption, transport, tunnel mode), IP-IP, GENEVE, site mode, VPN mode

**Services**
- TCP, UDP, ICMP, IP-proto, SUN-RPC service objects and service groups

**Data Sources & Actions**
- `smc_href` / `smc_sub_href` – element lookup by type and name pattern
- `smc_command` – generic engine actions (initial contact, bind license, policy upload/refresh, backup, reboot, go offline/online/standby)
- `smc_initial_contact` – node bootstrapping (SSH, root password, timezone, keyboard layout)

---

## 0.0.3 (2026-03-18)
"development version"
- Terraform Provider 0.0.3 – version bump, self-contained examples, and version update script
- Fix copyright headers in internal packages

## 0.0.2

"development version"
fix new provider name reference in docs

## 0.0.1

"development version"
- Block syntax will evolve without backward compatibility in the future
- Internal ID/name/block will change (some will disappear, others may be added) without backward compatibility in the future
- State generated with this version will have to be deleted in the future to switch to a new version
