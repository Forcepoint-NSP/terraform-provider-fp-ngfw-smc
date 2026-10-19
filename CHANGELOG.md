## 1.751.0 (2026-10-19)

**New features**
- New `smc_engine_upload_policy` and `smc_engine_refresh_policy` actions, replacing the generic `smc_command` approach: they wait for the SMC operation to complete and display a progress bar
- New `smc_generate_vpn_certificate` action to generate VPN certificates
- Drift detection: : optional and experimental
- Optional parallel SMC request mode; on parallel creates, element lock errors ("locked by") are now recognised

**Bug fixes**
- Login: transient failures (HTTP 5xx, connection refused or reset, timeouts) are now retried with backoff, and login errors carry more diagnostic detail
- Login: improved handling of login failures returning HTTP 401
- Fixed a provider crash on destroy; resource operations are now guarded against an unconfigured provider
- URLs containing whitespace are now escaped consistently
- Engines no longer end up with invalid routes after an update (timeout issue); routing diagnostics and retry validation are improved
- VPN topology changes are now committed before a policy upload or refresh, and a failed VPN save no longer blocks policy upload
- `smc_internal_endpoint` is no longer rejected by SMC 7.5+ because of `ssl_vpn_portal`
- Fixed unknown computed values on `smc_expression` `sub_expression`

**Documentation**
- Policy upload/refresh actions: alternative slots and their types, the default colour progress bar and the `preserve_connection` option are now documented correctly
- New section on the `before_create` proxy pattern for `smc_initial_contact`

## 1.750.0 (2026-06-23)

Targets SMC API **7.5 / 7.5.0** (base path `/7.4/` → `/7.5/`). Main changes vs 7.4.1:

**SSL VPN Portal → Application Access Portal** _(API only — no example yet)_
- SSL VPN Portal feature renamed and reworked: all `SSL_VPN_*` schemas removed, replaced by `Application_Access_*` (portal, portal policy, rules, service profile, web service, SSO domain, logo/icon/background assets)
- Internal gateway: `ssl_vpn_portal_setting` → `application_access_portal_setting`

**ZTNA Connector → Replica Internal Gateways** _(API only — no example yet)_
- `ZTNA_Connector_Settings` removed; `ztna_connector_settings` dropped from all engines
- New `replica_internal_gateway` / `replica_internal_endpoint` elements; `replica_internal_gateways` added on single/cluster/cloud/virtual firewalls, master engine, VSS container

**Interfaces**
- New `APN_Interface`; `Modem_Interface` gains `apn_interfaces` + `pin_code`, drops `modem_auth_method` — example: `engines/single_fw/single_fw_modem_interfaces`
- Legacy `apn` / `phone_number` / `pin_code` removed from node interfaces
- `Physical_Interface` gains `native_vlan_id` _(API only — no example yet)_

**VPN / IPsec** _(API only — no example yet)_
- VPN profile: TFC traffic-flow confidentiality (`tfc_padding`, `tfc_disable_dscp_copy`)
- Gateway settings: dead-peer-detection tuning (`dpd_interval`, `dpd_timeout`)

**Policies & logging** _(API only — no example yet)_
- Firewall rule action: authentication & confirmation tracking + caching
- Log tasks (archive/delete/export): `use_elasticsearch`

---

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
