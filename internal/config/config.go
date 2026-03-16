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
	"sync"
)

/*
   Config loading and access functions
   config file name is: tfsmc.conf.json
   config file is in local directory

   first example config file content:
{
  "hcl2": ["smc_single_fw", "smc_routing_node"]
}

   second example config file content:
{
  "hcl2": true
}


*/

const ConfigFileName = "tfsmc.conf.json"

// Config represents the provider configuration
type Config struct {
	Hcl2         any  `json:"hcl2"`
	Experimental bool `json:"experimental"`
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
	err := loadConfigFromPathLocked(ConfigFileName)
	if err != nil {
		return err
	}
	configLoaded = true
	return nil
}

// loadConfigFromPathLocked is the internal implementation without locking
// Caller must hold configMutex
func loadConfigFromPathLocked(path string) error {
	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// No config file is not an error - use defaults
		globalConfig = nil
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
	return nil
}

// validateConfig validates the configuration content
func validateConfig(cfg *Config) error {
	// Validate experimental field
	// No additional validation needed - bool type is validated during JSON unmarshaling

	// Validate hcl2 field
	// If hcl2 is nil, it's valid (means not configured)
	if cfg.Hcl2 == nil {
		return nil
	}

	switch v := cfg.Hcl2.(type) {
	case bool:
		// Boolean is valid
		return nil
	case []any:
		// Array must contain only strings
		for i, item := range v {
			if _, ok := item.(string); !ok {
				return fmt.Errorf("hcl2 array element at index %d must be a string, got %T", i, item)
			}
		}
		return nil
	default:
		return fmt.Errorf("hcl2 must be either a boolean or an array of strings, got %T", v)
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

	err := LoadConfig()
	if err != nil {
		return false, err
	}

	if globalConfig == nil {
		return false, nil
	}

	switch v := globalConfig.Hcl2.(type) {
	case bool:
		return v, nil
	case []interface{}:
		for _, item := range v {
			if str, ok := item.(string); ok && str == resourceType {
				return true, nil
			}
		}
		return false, nil
	default:
		return false, nil
	}
}

// GetConfig returns the current global configuration (for testing)
func GetConfig() *Config {
	configMutex.Lock()
	defer configMutex.Unlock()
	return globalConfig
}

// ResetConfig resets the global configuration (for testing)
func ResetConfig() {
	globalConfig = nil
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
