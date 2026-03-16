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

package datastore

import (
	"context"
	"encoding/json"
	"os"
	"testing"
)

func TestManagePendingDelete_AddOperation(t *testing.T) {
	// Setup: use temp directory
	tmpDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(oldWd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	ctx := context.Background()
	baseUrl := "http://localhost:8082/7.4"
	resourceType := "host"
	href := "http://localhost:8082/7.4/hosts/123456"

	// Act
	err = ManagePendingDelete(ctx, AddPendingDelete, baseUrl, resourceType, href)

	// Assert
	if err != nil {
		t.Errorf("ManagePendingDelete failed: %v", err)
	}

	// Verify file was created
	data, err := os.ReadFile("tfsmc.state.json")
	if err != nil {
		t.Fatalf("State file not created: %v", err)
	}

	var state StateFile
	if err := json.Unmarshal(data, &state); err != nil {
		t.Fatalf("Failed to unmarshal state file: %v", err)
	}

	if len(state.PendingDeletes) != 1 {
		t.Errorf("Expected 1 pending delete, got %d", len(state.PendingDeletes))
	}

	if state.PendingDeletes[0].Href != href {
		t.Errorf("Expected Href %q, got %q", href, state.PendingDeletes[0].Href)
	}
}

func TestManagePendingDelete_UpdateOperation(t *testing.T) {
	// Setup: use temp directory
	tmpDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(oldWd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	ctx := context.Background()
	href := "http://localhost:8082/7.4/hosts/123456"

	// Create initial state with an entry
	initialState := StateFile{
		PendingDeletes: []PendingDelete{
			{
				BaseUrl:      "http://localhost:8082/7.4",
				ResourceType: "host",
				Href:         href,
			},
		},
	}
	data, _ := json.MarshalIndent(initialState, "", "  ")
	os.WriteFile("tfsmc.state.json", data, 0644)

	// Act: update the entry with new baseUrl
	newBaseUrl := "http://localhost:8082/7.5"
	err = ManagePendingDelete(ctx, UpdatePendingDelete, newBaseUrl, "host", href)

	// Assert
	if err != nil {
		t.Errorf("ManagePendingDelete failed: %v", err)
	}

	// Verify update
	data, err = os.ReadFile("tfsmc.state.json")
	if err != nil {
		t.Fatalf("Failed to read state file: %v", err)
	}

	var state StateFile
	if err := json.Unmarshal(data, &state); err != nil {
		t.Fatalf("Failed to unmarshal state file: %v", err)
	}

	if len(state.PendingDeletes) != 1 {
		t.Errorf("Expected 1 pending delete, got %d", len(state.PendingDeletes))
	}

	if state.PendingDeletes[0].BaseUrl != newBaseUrl {
		t.Errorf("Expected BaseUrl %q, got %q", newBaseUrl, state.PendingDeletes[0].BaseUrl)
	}
}

func TestManagePendingDelete_UpdateOperationNotFound(t *testing.T) {
	// Setup: use temp directory
	tmpDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(oldWd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	ctx := context.Background()

	// Act: try to update non-existent entry
	err = ManagePendingDelete(ctx, UpdatePendingDelete, "http://localhost:8082/7.4", "host", "http://localhost:8082/7.4/hosts/notfound")

	// Assert: should fail
	if err == nil {
		t.Error("Expected error for non-existent entry, but got nil")
	}
	if err != nil && err.Error() != "pending delete entry not found for href: http://localhost:8082/7.4/hosts/notfound" {
		t.Errorf("Expected 'not found' error, got: %v", err)
	}
}

func TestManagePendingDelete_RemoveOperation(t *testing.T) {
	// Setup: use temp directory
	tmpDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(oldWd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	ctx := context.Background()
	href1 := "http://localhost:8082/7.4/hosts/123456"
	href2 := "http://localhost:8082/7.4/hosts/654321"

	// Create initial state with two entries
	initialState := StateFile{
		PendingDeletes: []PendingDelete{
			{
				BaseUrl:      "http://localhost:8082/7.4",
				ResourceType: "host",
				Href:         href1,
			},
			{
				BaseUrl:      "http://localhost:8082/7.4",
				ResourceType: "host",
				Href:         href2,
			},
		},
	}
	data, _ := json.MarshalIndent(initialState, "", "  ")
	os.WriteFile("tfsmc.state.json", data, 0644)

	// Act: remove first entry
	err = ManagePendingDelete(ctx, RemovePendingDelete, "", "", href1)

	// Assert
	if err != nil {
		t.Errorf("ManagePendingDelete failed: %v", err)
	}

	// Verify removal
	data, err = os.ReadFile("tfsmc.state.json")
	if err != nil {
		t.Fatalf("Failed to read state file: %v", err)
	}

	var state StateFile
	if err := json.Unmarshal(data, &state); err != nil {
		t.Fatalf("Failed to unmarshal state file: %v", err)
	}

	if len(state.PendingDeletes) != 1 {
		t.Errorf("Expected 1 pending delete after removal, got %d", len(state.PendingDeletes))
	}

	if state.PendingDeletes[0].Href != href2 {
		t.Errorf("Expected remaining entry to be %q, got %q", href2, state.PendingDeletes[0].Href)
	}
}

func TestManagePendingDelete_RemoveOperationNotFound(t *testing.T) {
	// Setup: use temp directory
	tmpDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(oldWd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	ctx := context.Background()

	// Act: try to remove non-existent entry
	err = ManagePendingDelete(ctx, RemovePendingDelete, "", "", "http://localhost:8082/7.4/hosts/notfound")

	// Assert: should fail
	if err == nil {
		t.Error("Expected error for non-existent entry, but got nil")
	}
}

func TestSavePendingDelete_CreatesNewStateFile(t *testing.T) {
	// Setup: use temp directory
	tmpDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(oldWd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	ctx := context.Background()
	baseUrl := "http://localhost:8082/7.4"
	resourceType := "host"
	href := "http://localhost:8082/7.4/hosts/123456"

	// Act
	err = SavePendingDelete(ctx, baseUrl, resourceType, href)

	// Assert
	if err != nil {
		t.Errorf("SavePendingDelete failed: %v", err)
	}

	// Verify file was created
	data, err := os.ReadFile("tfsmc.state.json")
	if err != nil {
		t.Fatalf("State file not created: %v", err)
	}

	var state StateFile
	if err := json.Unmarshal(data, &state); err != nil {
		t.Fatalf("Failed to unmarshal state file: %v", err)
	}

	if len(state.PendingDeletes) != 1 {
		t.Errorf("Expected 1 pending delete, got %d", len(state.PendingDeletes))
	}

	if state.PendingDeletes[0].BaseUrl != baseUrl {
		t.Errorf("Expected BaseUrl %q, got %q", baseUrl, state.PendingDeletes[0].BaseUrl)
	}
	if state.PendingDeletes[0].ResourceType != resourceType {
		t.Errorf("Expected ResourceType %q, got %q", resourceType, state.PendingDeletes[0].ResourceType)
	}
	if state.PendingDeletes[0].Href != href {
		t.Errorf("Expected Href %q, got %q", href, state.PendingDeletes[0].Href)
	}
}

func TestSavePendingDelete_AppendsToExistingStateFile(t *testing.T) {
	// Setup: use temp directory
	tmpDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(oldWd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	ctx := context.Background()

	// Create initial state file
	initialState := `{
  "pending_deletes": [
    {
      "base_url": "http://localhost:8082/7.4",
      "resource_type": "firewall",
      "href": "http://localhost:8082/7.4/elements/single_fw/111111"
    }
  ]
}`
	if err := os.WriteFile("tfsmc.state.json", []byte(initialState), 0644); err != nil {
		t.Fatalf("Failed to create initial state file: %v", err)
	}

	// Act: add another pending delete
	baseUrl := "http://localhost:8082/7.4"
	resourceType := "host"
	href := "http://localhost:8082/7.4/hosts/123456"
	err = SavePendingDelete(ctx, baseUrl, resourceType, href)

	// Assert
	if err != nil {
		t.Errorf("SavePendingDelete failed: %v", err)
	}

	// Verify content
	data, err := os.ReadFile("tfsmc.state.json")
	if err != nil {
		t.Fatalf("Failed to read state file: %v", err)
	}

	var state StateFile
	if err := json.Unmarshal(data, &state); err != nil {
		t.Fatalf("Failed to unmarshal state file: %v", err)
	}

	if len(state.PendingDeletes) != 2 {
		t.Errorf("Expected 2 pending deletes, got %d", len(state.PendingDeletes))
	}

	// Verify second entry
	if state.PendingDeletes[1].BaseUrl != baseUrl {
		t.Errorf("Expected BaseUrl %q, got %q", baseUrl, state.PendingDeletes[1].BaseUrl)
	}
	if state.PendingDeletes[1].ResourceType != resourceType {
		t.Errorf("Expected ResourceType %q, got %q", resourceType, state.PendingDeletes[1].ResourceType)
	}
	if state.PendingDeletes[1].Href != href {
		t.Errorf("Expected Href %q, got %q", href, state.PendingDeletes[1].Href)
	}
}

func TestSavePendingDelete_HandlesInvalidJsonGracefully(t *testing.T) {
	// Setup: use temp directory
	tmpDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(oldWd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	ctx := context.Background()

	// Create state file with invalid JSON
	if err := os.WriteFile("tfsmc.state.json", []byte("{invalid json"), 0644); err != nil {
		t.Fatalf("Failed to create state file: %v", err)
	}

	// Act: should not fail, should recover from invalid JSON
	baseUrl := "http://localhost:8082/7.4"
	resourceType := "host"
	href := "http://localhost:8082/7.4/hosts/123456"
	err = SavePendingDelete(ctx, baseUrl, resourceType, href)

	// Assert
	if err != nil {
		t.Errorf("SavePendingDelete failed: %v", err)
	}

	// Verify file was updated with new entry
	data, err := os.ReadFile("tfsmc.state.json")
	if err != nil {
		t.Fatalf("Failed to read state file: %v", err)
	}

	var state StateFile
	if err := json.Unmarshal(data, &state); err != nil {
		t.Fatalf("Failed to unmarshal state file: %v", err)
	}

	if len(state.PendingDeletes) != 1 {
		t.Errorf("Expected 1 pending delete, got %d", len(state.PendingDeletes))
	}
}

func TestSavePendingDelete_FilePermissions(t *testing.T) {
	// Setup: use temp directory
	tmpDir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(oldWd)

	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Act
	ctx := context.Background()
	baseUrl := "http://localhost:8082/7.4"
	resourceType := "host"
	href := "http://localhost:8082/7.4/hosts/123456"
	err = SavePendingDelete(ctx, baseUrl, resourceType, href)

	if err != nil {
		t.Errorf("SavePendingDelete failed: %v", err)
	}

	// Verify file permissions (should be readable)
	fileInfo, err := os.Stat("tfsmc.state.json")
	if err != nil {
		t.Fatalf("Failed to stat state file: %v", err)
	}

	// Check that file is readable
	if fileInfo.Mode()&0400 == 0 {
		t.Error("State file is not readable by owner")
	}
	// Check that file is writable
	if fileInfo.Mode()&0200 == 0 {
		t.Error("State file is not writable by owner")
	}
}
