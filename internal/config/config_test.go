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
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig_NoFile(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Try to load non-existent file
	err := loadConfigFromPathLocked("nonexistent.json")
	if err != nil {
		t.Errorf("loadConfigFromPathLocked should not error when file doesn't exist, got: %v", err)
	}

	if GetConfig() != nil {
		t.Error("Config should be nil when file doesn't exist")
	}
}

func TestLoadConfig_ValidBooleanTrue(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": true}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config
	err := loadConfigFromPathLocked(configPath)
	if err != nil {
		t.Fatalf("loadConfigFromPathLocked failed: %v", err)
	}

	cfg := GetConfig()
	if cfg == nil {
		t.Fatal("Config should not be nil")
	}

	if hcl2, ok := cfg.Hcl2.(bool); !ok || !hcl2 {
		t.Errorf("Expected hcl2 to be true, got: %v", cfg.Hcl2)
	}
}

func TestLoadConfig_ValidBooleanFalse(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": false}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config
	err := loadConfigFromPathLocked(configPath)
	if err != nil {
		t.Fatalf("loadConfigFromPathLocked failed: %v", err)
	}

	cfg := GetConfig()
	if cfg == nil {
		t.Fatal("Config should not be nil")
	}

	if hcl2, ok := cfg.Hcl2.(bool); !ok || hcl2 {
		t.Errorf("Expected hcl2 to be false, got: %v", cfg.Hcl2)
	}
}

func TestLoadConfig_ValidArray(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": ["smc_single_fw", "smc_routing_node"]}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config
	err := loadConfigFromPathLocked(configPath)
	if err != nil {
		t.Fatalf("loadConfigFromPathLocked failed: %v", err)
	}

	cfg := GetConfig()
	if cfg == nil {
		t.Fatal("Config should not be nil")
	}

	hcl2Array, ok := cfg.Hcl2.([]interface{})
	if !ok {
		t.Fatalf("Expected hcl2 to be array, got: %T", cfg.Hcl2)
	}

	if len(hcl2Array) != 2 {
		t.Errorf("Expected array length 2, got: %d", len(hcl2Array))
	}
}

func TestLoadConfig_InvalidJSON(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file with invalid JSON
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": invalid}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config should fail
	err := loadConfigFromPathLocked(configPath)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestLoadConfig_EmptyFile(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary empty config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config
	err := loadConfigFromPathLocked(configPath)
	if err != nil {
		t.Fatalf("loadConfigFromPathLocked failed: %v", err)
	}

	cfg := GetConfig()
	if cfg == nil {
		t.Fatal("Config should not be nil")
	}

	if cfg.Hcl2 != nil {
		t.Errorf("Expected hcl2 to be nil, got: %v", cfg.Hcl2)
	}
}

func TestIsHcl2Enabled_NoConfig(t *testing.T) {
	// Reset state before test
	ResetConfig()

	enabled, err := IsHcl2Enabled("smc_single_fw")
	if err != nil {
		t.Fatalf("IsHcl2Enabled failed: %v", err)
	}
	if enabled {
		t.Error("Expected IsHcl2Enabled to return false when no config is loaded")
	}
}

func TestIsHcl2Enabled_BooleanTrue(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": true}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	if err := loadConfigFromPathLocked(configPath); err != nil {
		t.Fatalf("loadConfigFromPathLocked failed: %v", err)
	}

	// All resource types should be enabled
	testCases := []string{
		"smc_single_fw",
		"smc_routing_node",
		"any_resource_type",
	}

	for _, resourceType := range testCases {
		enabled, err := IsHcl2Enabled(resourceType)
		if err != nil {
			t.Fatalf("IsHcl2Enabled failed: %v", err)
		}
		if !enabled {
			t.Errorf("Expected IsHcl2Enabled(%q) to return true", resourceType)
		}
	}
}

func TestIsHcl2Enabled_BooleanFalse(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": false}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	if err := loadConfigFromPathLocked(configPath); err != nil {
		t.Fatalf("loadConfigFromPathLocked failed: %v", err)
	}

	// All resource types should be disabled
	testCases := []string{
		"smc_single_fw",
		"smc_routing_node",
		"any_resource_type",
	}

	for _, resourceType := range testCases {
		enabled, err := IsHcl2Enabled(resourceType)
		if err != nil {
			t.Fatalf("IsHcl2Enabled failed: %v", err)
		}
		if enabled {
			t.Errorf("Expected IsHcl2Enabled(%q) to return false", resourceType)
		}
	}
}

func TestIsHcl2Enabled_ArrayIncluded(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": ["smc_single_fw", "smc_routing_node"]}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	if err := loadConfigFromPathLocked(configPath); err != nil {
		t.Fatalf("loadConfigFromPathLocked failed: %v", err)
	}

	// These should be enabled
	enabledTypes := []string{"smc_single_fw", "smc_routing_node"}
	for _, resourceType := range enabledTypes {
		enabled, err := IsHcl2Enabled(resourceType)
		if err != nil {
			t.Fatalf("IsHcl2Enabled failed: %v", err)
		}
		if !enabled {
			t.Errorf("Expected IsHcl2Enabled(%q) to return true", resourceType)
		}
	}

	// These should be disabled
	disabledTypes := []string{"other_resource", "smc_cluster"}
	for _, resourceType := range disabledTypes {
		enabled, err := IsHcl2Enabled(resourceType)
		if err != nil {
			t.Fatalf("IsHcl2Enabled failed: %v", err)
		}
		if enabled {
			t.Errorf("Expected IsHcl2Enabled(%q) to return false", resourceType)
		}
	}
}

func TestIsHcl2Enabled_ArrayEmpty(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": []}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	if err := loadConfigFromPathLocked(configPath); err != nil {
		t.Fatalf("loadConfigFromPathLocked failed: %v", err)
	}

	// No resource types should be enabled
	enabled, err := IsHcl2Enabled("smc_single_fw")
	if err != nil {
		t.Fatalf("IsHcl2Enabled failed: %v", err)
	}
	if enabled {
		t.Error("Expected IsHcl2Enabled to return false for empty array")
	}
}

func TestIsHcl2Enabled_InvalidType(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file with invalid type (number)
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": 123}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Should fail validation now
	err := loadConfigFromPathLocked(configPath)
	if err == nil {
		t.Error("Expected validation error for invalid type, got nil")
	}

	if err != nil && !contains(err.Error(), "must be either a boolean or an array of strings") {
		t.Errorf("Expected validation error about type, got: %v", err)
	}
}

func TestIsHcl2Enabled_ArrayWithNonStringValues(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file with mixed types in array
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": ["smc_single_fw", 123, true, "smc_routing_node"]}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Should fail validation now
	err := loadConfigFromPathLocked(configPath)
	if err == nil {
		t.Error("Expected validation error for non-string elements in array, got nil")
	}

	if err != nil && !contains(err.Error(), "must be a string") {
		t.Errorf("Expected validation error about non-string element, got: %v", err)
	}
}

func TestResetConfig(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create and load a config
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": true}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	if err := loadConfigFromPathLocked(configPath); err != nil {
		t.Fatalf("loadConfigFromPathLocked failed: %v", err)
	}

	if GetConfig() == nil {
		t.Fatal("Config should not be nil after loading")
	}

	// Reset config
	ResetConfig()

	if GetConfig() != nil {
		t.Error("Config should be nil after reset")
	}
}

func TestLoadConfig_ReadPermissionError(t *testing.T) {
	// Skip on Windows as it handles permissions differently
	if os.Getenv("GOOS") == "windows" {
		t.Skip("Skipping permission test on Windows")
	}

	// Reset state before test
	ResetConfig()

	// Create temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": true}`

	if err := os.WriteFile(configPath, []byte(content), 0000); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Try to load config - should fail due to permissions
	err := loadConfigFromPathLocked(configPath)
	if err == nil {
		t.Error("Expected error when reading file with no permissions")
	}
}

func TestLoadConfig_ValidationInvalidType(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file with invalid type (number)
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": 123}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config should fail validation
	err := loadConfigFromPathLocked(configPath)
	if err == nil {
		t.Error("Expected validation error for invalid type, got nil")
	}

	if err != nil && !contains(err.Error(), "must be either a boolean or an array of strings") {
		t.Errorf("Expected validation error about type, got: %v", err)
	}
}

func TestLoadConfig_ValidationInvalidTypeString(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file with string type (should be bool or array)
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": "string_value"}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config should fail validation
	err := loadConfigFromPathLocked(configPath)
	if err == nil {
		t.Error("Expected validation error for string type, got nil")
	}

	if err != nil && !contains(err.Error(), "must be either a boolean or an array of strings") {
		t.Errorf("Expected validation error about type, got: %v", err)
	}
}

func TestLoadConfig_ValidationArrayWithInvalidElement(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file with invalid element in array
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": ["valid_string", 123, "another_string"]}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config should fail validation
	err := loadConfigFromPathLocked(configPath)
	if err == nil {
		t.Error("Expected validation error for non-string element in array, got nil")
	}

	if err != nil && !contains(err.Error(), "must be a string") {
		t.Errorf("Expected validation error about non-string element, got: %v", err)
	}
}

func TestLoadConfig_ValidationArrayWithBoolElement(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file with boolean element in array
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": ["smc_single_fw", true]}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config should fail validation
	err := loadConfigFromPathLocked(configPath)
	if err == nil {
		t.Error("Expected validation error for boolean element in array, got nil")
	}

	if err != nil && !contains(err.Error(), "must be a string") {
		t.Errorf("Expected validation error about non-string element, got: %v", err)
	}
}

func TestLoadConfig_ValidationArrayWithObjectElement(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file with object element in array
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": ["smc_single_fw", {"key": "value"}]}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config should fail validation
	err := loadConfigFromPathLocked(configPath)
	if err == nil {
		t.Error("Expected validation error for object element in array, got nil")
	}

	if err != nil && !contains(err.Error(), "must be a string") {
		t.Errorf("Expected validation error about non-string element, got: %v", err)
	}
}

func TestLoadConfig_ValidationEmptyArrayValid(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file with empty array (should be valid)
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{"hcl2": []}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config should succeed
	err := loadConfigFromPathLocked(configPath)
	if err != nil {
		t.Errorf("Empty array should be valid, got error: %v", err)
	}

	cfg := GetConfig()
	if cfg == nil {
		t.Fatal("Config should not be nil")
	}
}

func TestLoadConfig_ValidationNilHcl2Valid(t *testing.T) {
	// Reset state before test
	ResetConfig()

	// Create temporary config file with no hcl2 field (should be valid)
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.json")
	content := `{}`

	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}

	// Load config should succeed
	err := loadConfigFromPathLocked(configPath)
	if err != nil {
		t.Errorf("Empty config should be valid, got error: %v", err)
	}

	cfg := GetConfig()
	if cfg == nil {
		t.Fatal("Config should not be nil")
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
