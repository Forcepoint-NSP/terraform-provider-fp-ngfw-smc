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

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

/*
   Config loading and access functions
   config file name is: tfsmc.conf.json
   config file is in local directory

   Example: enable drift detection and parallel mode, excluding one type
   from parallelization:
{
  "drift_detection": ["smc_host"],
  "parallel_mode": true,
  "parallel_mode_exclusion_list": ["smc_routing_rule", "smc_fw_cluster"]
}

   drift_detection: opt-in via TFSMC_DRIFT_DETECTION env var (precedence over config)
   parallel_mode: opt-in via TFSMC_PARALLEL_CREATE env var (precedence over config),
     or via this config field alone if the env var is unset
   parallel_mode_exclusion_list: resource types to keep serialized once
     parallel_mode is on, however it was enabled
*/

const ConfigFileName = "tfsmc.conf.json"

// Config represents the provider configuration
type Config struct {
	Hcl2                      any      `json:"hcl2"`
	DriftDetection            any      `json:"drift_detection"`
	ParallelMode              bool     `json:"parallel_mode"`
	ParallelModeExclusionList []string `json:"parallel_mode_exclusion_list"`
	Experimental              bool     `json:"experimental"`
}

var (
	globalConfig *Config
	configMutex  sync.RWMutex
	configLoaded bool = false
)

// LoadConfig loads the configuration from tfsmc.conf.json file.  If
// the file doesn't exist, it returns nil (no error) and defaults are
// used
func LoadConfig() error {
	configMutex.Lock()
	defer configMutex.Unlock()

	if configLoaded {
		return nil
	}
	return loadConfigFromPathLocked(ConfigFileName)
}

// loadConfigFromPathLocked is the internal implementation without locking
// Caller must hold configMutex. On success it marks the config as loaded
// (mirroring LoadConfig's own bookkeeping) so tests that call this directly
// with an arbitrary path don't get silently overwritten by a later
// LoadConfig() call re-reading the real ConfigFileName.
func loadConfigFromPathLocked(path string) error {
	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// No config file is not an error - use defaults
		globalConfig = nil
		configLoaded = true
		return nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse config file %s: %w", path, err)
	}

	// Validate config content
	if err := validateConfig(&cfg); err != nil {
		return fmt.Errorf("invalid config in %s: %w", path, err)
	}

	globalConfig = &cfg
	configLoaded = true
	return nil
}

// validateConfig validates the configuration content
func validateConfig(cfg *Config) error {
	if err := validateBoolOrStringArray("hcl2", cfg.Hcl2); err != nil {
		return err
	}
	return validateBoolOrStringArray("drift_detection", cfg.DriftDetection)
}

// validateBoolOrStringArray validates a config field that must be
// either a boolean or an array of resource type name strings. A nil
// value is valid (means not configured).
func validateBoolOrStringArray(fieldName string, value any) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case bool:
		return nil
	case []any:
		for i, item := range v {
			if _, ok := item.(string); !ok {
				return fmt.Errorf("%s array element at index %d must be a string, got %T", fieldName, i, item)
			}
		}
		return nil
	default:
		return fmt.Errorf("%s must be either a boolean or an array of strings, got %T", fieldName, v)
	}
}

// IsHcl2Enabled checks if HCL2 is enabled for a specific resource type
//   - hcl2 is set to true (boolean)
//   - hcl2 is an array containing the resourceType
//
// Returns false if:
//   - no config is loaded
//   - hcl2 is false
//   - hcl2 is an array that doesn't contain the resourceType
func IsHcl2Enabled(resourceType string) (bool, error) {
	if err := LoadConfig(); err != nil {
		return false, err
	}
	if globalConfig == nil {
		return false, nil
	}
	return isEnabledFor(globalConfig.Hcl2, resourceType), nil
}

// DriftDetectionEnvVar, when set to "0" or "1", overrides the
// "drift_detection" config file setting globally (for every resource
// type), without requiring a tfsmc.conf.json file.
const DriftDetectionEnvVar = "TFSMC_DRIFT_DETECTION"

// IsDriftDetectionEnabled checks if Terraform Read (drift detection) is
// enabled for a specific resource type. If the TFSMC_DRIFT_DETECTION
// environment variable is set, it takes precedence over the config
// file and applies to every resource type. Otherwise this follows the
// same true/false/array-of-resource-types convention as IsHcl2Enabled,
// and defaults to false so drift detection is opt-in per resource.
func IsDriftDetectionEnabled(resourceType string) (bool, error) {
	if v, ok := os.LookupEnv(DriftDetectionEnvVar); ok {
		switch v {
		case "1":
			return true, nil
		case "0":
			return false, nil
		default:
			return false, fmt.Errorf("%s must be \"0\" or \"1\", got %q", DriftDetectionEnvVar, v)
		}
	}

	if err := LoadConfig(); err != nil {
		return false, err
	}
	if globalConfig == nil {
		return false, nil
	}
	return isEnabledFor(globalConfig.DriftDetection, resourceType), nil
}

// ParallelModeEnvVar enables parallel request mode globally without requiring
// tfsmc.conf.json. Set to "1" to enable or "0" to disable; takes precedence
// over the "parallel_mode" config file field.
const ParallelModeEnvVar = "TFSMC_PARALLEL_CREATE"

// parseParallelEnv interprets a TFSMC_PARALLEL_CREATE value. Unrecognized
// values are an error, like TFSMC_DRIFT_DETECTION, so a typo never silently
// disables parallel mode.
func parseParallelEnv(v string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "1", "true", "yes", "on":
		return true, nil
	case "0", "false", "no", "off", "":
		return false, nil
	default:
		return false, fmt.Errorf("%s must be one of 1/0/true/false/yes/no/on/off, got %q", ParallelModeEnvVar, v)
	}
}

// ValidateParallelModeEnv returns an error if TFSMC_PARALLEL_CREATE is set to
// an unrecognized value. Called at provider configure time so a typo fails
// the run instead of silently disabling parallel mode.
func ValidateParallelModeEnv() error {
	v, ok := os.LookupEnv(ParallelModeEnvVar)
	if !ok {
		return nil
	}
	_, err := parseParallelEnv(v)
	return err
}

// IsParallelModeEnabled checks if parallel mode is active for a resource
// type. Parallel mode can be switched on globally either via the
// TFSMC_PARALLEL_CREATE env var (which takes precedence) or via the
// "parallel_mode" boolean in tfsmc.conf.json, so a config file alone is
// enough to enable it without setting any environment variable. Once
// active, resource types listed in "parallel_mode_exclusion_list" stay
// serialized; this allows users to add problematic types to
// tfsmc.conf.json if concurrent updates cause conflicts.
func IsParallelModeEnabled(resourceType string) (bool, error) {
	enabled, err := isParallelModeEnabledGlobally()
	if err != nil {
		return false, err
	}
	if !enabled {
		return false, nil
	}

	if err := LoadConfig(); err != nil {
		return false, err
	}
	if globalConfig == nil {
		return true, nil // No exclusions configured, allow parallel
	}
	if isInExclusionList(globalConfig.ParallelModeExclusionList, resourceType) {
		return false, nil // Type is excluded
	}
	return true, nil
}

// isParallelModeEnabledGlobally reports whether parallel mode is switched on
// at all, before per-type exclusions are considered. The env var takes
// precedence when set; otherwise the "parallel_mode" config file boolean
// applies, so tfsmc.conf.json alone can enable the feature.
func isParallelModeEnabledGlobally() (bool, error) {
	if v, ok := os.LookupEnv(ParallelModeEnvVar); ok {
		return parseParallelEnv(v)
	}

	if err := LoadConfig(); err != nil {
		return false, err
	}
	if globalConfig == nil {
		return false, nil
	}
	return globalConfig.ParallelMode, nil
}

// isInExclusionList checks if resourceType is in a parallel_mode_exclusion_list.
func isInExclusionList(list []string, resourceType string) bool {
	for _, excluded := range list {
		if excluded == resourceType {
			return true
		}
	}
	return false
}

// isEnabledFor implements the true/false/array-of-resource-types
// convention shared by IsHcl2Enabled and IsDriftDetectionEnabled.
func isEnabledFor(value any, resourceType string) bool {
	switch v := value.(type) {
	case bool:
		return v
	case []interface{}:
		for _, item := range v {
			if str, ok := item.(string); ok && str == resourceType {
				return true
			}
		}
		return false
	default:
		return false
	}
}

// GetConfig returns the current global configuration (for testing)
func GetConfig() *Config {
	configMutex.Lock()
	defer configMutex.Unlock()
	return globalConfig
}

// ResetConfig resets the global configuration (for testing). It also clears
// the configLoaded flag so a subsequent LoadConfig() call re-reads the file
// instead of returning the cached (possibly stale) state.
func ResetConfig() {
	configMutex.Lock()
	defer configMutex.Unlock()
	globalConfig = nil
	configLoaded = false
}

// IsExperimentalEnabled checks if experimental features are enabled
func IsExperimentalEnabled() bool {
	err := LoadConfig()
	if err != nil {
		return false
	}

	if globalConfig == nil {
		return false
	}

	return globalConfig.Experimental
}
