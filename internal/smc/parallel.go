// Copyright 2026 Forcepoint LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package smc

import (
	"log"
	"sync"

	"github.com/terraform-providers/terraform-provider-smc/internal/config"
)

// Parallel mode is opt-in via the TFSMC_PARALLEL_CREATE env var, or via the
// "parallel_mode" boolean in tfsmc.conf.json (see internal/config).
// Concurrency is controlled by terraform apply --parallelism (default recommended: 20,
// matching SMC's DB connection limit). Config file (tfsmc.conf.json) specifies
// resource types to exclude from parallelization via the
// parallel_mode_exclusion_list field.

// alwaysSerializedResourceTypes cannot be parallelized due to concurrent update risks:
// - *_rule types: SMC holds exclusive per-policy write locks, causing 422 errors
// - Firewall types (single_fw, fw_cluster): complex cross-resource dependencies
// - VPN types: complex bring-up dependencies
var alwaysSerializedResourceTypes = map[string]bool{
	// Rank/position-ordered rules: take exclusive write locks
	"fw_ipv4_access_rule":           true,
	"fw_ipv4_nat_rule":              true,
	"fw_ipv6_access_rule":           true,
	"fw_ipv6_nat_rule":              true,
	"ips_ipv6_access_rule":          true,
	"ips_ethernet_rule":             true,
	"l2_interface_ipv4_access_rule": true,
	"l2_interface_ipv6_access_rule": true,
	"l2_interface_ethernet_rule":    true,
	"layer2_ipv4_access_rule":       true,
	"layer2_ipv6_access_rule":       true,
	"layer2_ethernet_rule":          true,
	"inspection_exception_rule":     true,
	"inspection_global_rule":        true,
	"alert_chain_rule":              true,
	"file_filtering_rule":           true,
	"limit_rule":                    true,
	"route_map_rule":                true,
	"rule_validity_time":            true,

	// Firewall and VPN resources: complex dependencies
	"single_fw":          true,
	"fw_cluster":         true,
	"vpn":                true,
	"rbvpn_tunnel":       true,
	"rbvpn_tunnel_group": true,
	"gateway_node":       true,
	"internal_endpoint":  true,
	"vpn_site":           true,
}

// terraformTypePrefix mirrors common.PROVIDER_TYPE_NAME + "_", the
// convention already used for drift_detection: tfsmc.conf.json blacklists
// list fully-qualified Terraform type names (e.g. "smc_fw_cluster"), not the
// bare API resource type (e.g. "fw_cluster") passed around internally.
// Duplicated as a literal instead of importing internal/common, which would
// create an import cycle (internal/common -> internal/smc -> internal/config).
const terraformTypePrefix = "smc_"

var parallelConfigErrOnce sync.Once

// IsResourceTypeParallel checks if a resource type can be parallelized.
// Requires: parallel mode enabled globally (env var or config, see
// internal/config.IsParallelModeEnabled) AND type not hard-guarded or
// excluded via tfsmc.conf.json's parallel_mode_exclusion_list.
func IsResourceTypeParallel(resourceType string) bool {
	if alwaysSerializedResourceTypes[resourceType] {
		return false
	}
	// Check if enabled and not excluded via tfsmc.conf.json, keyed by the
	// fully-qualified Terraform type name.
	ok, err := config.IsParallelModeEnabled(terraformTypePrefix + resourceType)
	if err != nil {
		parallelConfigErrOnce.Do(func() {
			log.Printf("[WARN] parallel mode disabled: %v", err)
		})
		return false
	}
	return ok
}
