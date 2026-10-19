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
	"os"
	"path/filepath"
	"testing"

	"github.com/terraform-providers/terraform-provider-smc/internal/config"
)

func withParallelEnv(t *testing.T, env map[string]string) {
	t.Helper()
	for k, v := range env {
		t.Setenv(k, v)
	}
	config.ResetConfig()
	t.Cleanup(func() {
		config.ResetConfig()
	})
}

func TestIsResourceTypeParallel_DisabledByDefault(t *testing.T) {
	withParallelEnv(t, map[string]string{})

	for _, rt := range []string{"host", "network", "fw_ipv4_access_rule", "single_fw"} {
		if IsResourceTypeParallel(rt) {
			t.Errorf("expected %q to stay serialized when TFSMC_PARALLEL_CREATE is unset", rt)
		}
	}
}

func TestIsResourceTypeParallel_EnabledWithEnvVar(t *testing.T) {
	withParallelEnv(t, map[string]string{config.ParallelModeEnvVar: "1"})

	// Non-guarded types should be parallel
	if !IsResourceTypeParallel("host") || !IsResourceTypeParallel("network") {
		t.Error("expected non-guarded types to be parallel when enabled")
	}
}

func TestIsResourceTypeParallel_HardGuardedAlwaysSerial(t *testing.T) {
	withParallelEnv(t, map[string]string{config.ParallelModeEnvVar: "1"})

	for rt := range alwaysSerializedResourceTypes {
		if IsResourceTypeParallel(rt) {
			t.Errorf("expected %q to stay serialized (hard-guarded)", rt)
		}
	}
}

func TestIsResourceTypeParallel_TruthyEnvValues(t *testing.T) {
	for _, v := range []string{"1", "true", "TRUE", "yes", "on"} {
		withParallelEnv(t, map[string]string{config.ParallelModeEnvVar: v})
		if !IsResourceTypeParallel("host") {
			t.Errorf("expected %q to enable parallel mode", v)
		}
	}
	for _, v := range []string{"0", "false", "", "nope"} {
		withParallelEnv(t, map[string]string{config.ParallelModeEnvVar: v})
		if IsResourceTypeParallel("host") {
			t.Errorf("expected %q to disable parallel mode", v)
		}
	}
}

func TestIsResourceTypeParallel_fw_clusterGuarded(t *testing.T) {
	withParallelEnv(t, map[string]string{config.ParallelModeEnvVar: "1"})

	if IsResourceTypeParallel("fw_cluster") {
		t.Error("expected fw_cluster to be hard-guarded (never parallel)")
	}
	if IsResourceTypeParallel("single_fw") {
		t.Error("expected single_fw to be hard-guarded (never parallel)")
	}
}

// Regression test: parallel_mode_exclusion_list in tfsmc.conf.json uses
// fully-qualified Terraform type names (e.g. "smc_routing_rule"), matching
// the guide's examples and drift_detection's established convention, while
// IsResourceTypeParallel is called with the bare API resource type (e.g.
// "routing_rule"). IsResourceTypeParallel used to pass the bare type
// straight through to config.IsParallelModeEnabled, so a documented
// exclusion like "smc_routing_rule" never matched anything and had no
// effect. Flagged by stonemind review.
func TestIsResourceTypeParallel_ConfigBlacklistUsesTerraformTypeName(t *testing.T) {
	withParallelEnv(t, map[string]string{config.ParallelModeEnvVar: "1"})

	tmpDir := t.TempDir()
	t.Chdir(tmpDir)
	content := `{"parallel_mode_exclusion_list": ["smc_routing_rule"]}`
	if err := os.WriteFile(filepath.Join(tmpDir, config.ConfigFileName), []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test config file: %v", err)
	}

	if IsResourceTypeParallel("routing_rule") {
		t.Error("expected \"routing_rule\" to be excluded via the \"smc_routing_rule\" blacklist entry")
	}
	if !IsResourceTypeParallel("host") {
		t.Error("expected \"host\" to remain parallel (not excluded)")
	}
}
