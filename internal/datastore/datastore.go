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
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// this structure gathers the pending delete information so that
// deletions can be retried later. In particular, the BaseUrl is
// needed to re-instantiate the SMC client in order to perform the
// delete (multi-smc case)
type PendingDelete struct {
	BaseUrl      string `json:"base_url"`
	ResourceType string `json:"resource_type"`
	Href         string `json:"href"`
}

type StateFile struct {
	PendingDeletes []PendingDelete `json:"pending_deletes"`
}

var (
	// Mutex for synchronizing access to the state file across goroutines
	stateFileMutex = &sync.Mutex{}
)

const stateFileName = "tfsmc.state.json"
const lockTimeout = 30 * time.Second

// PendingDeleteOperation defines the type of operation to perform
type PendingDeleteOperation string

const (
	AddPendingDelete    PendingDeleteOperation = "add"
	UpdatePendingDelete PendingDeleteOperation = "update"
	RemovePendingDelete PendingDeleteOperation = "remove"
)

// readStateFile acquires a lock and reads the current state from the state file.
// Returns the state and releases the lock via the returned unlock function.
func readStateFile(ctx context.Context) (*StateFile, func(), error) {
	// Acquire exclusive lock. Using goroutine and channel to implement timeout
	lockAcquired := make(chan bool, 1)
	go func() {
		stateFileMutex.Lock()
		lockAcquired <- true
	}()

	// Wait for lock with timeout
	select {
	case <-lockAcquired:
		// Lock acquired successfully
	case <-time.After(lockTimeout):
		return nil, nil, fmt.Errorf("failed to acquire lock on %s within %v", stateFileName, lockTimeout)
	case <-ctx.Done():
		return nil, nil, fmt.Errorf("context cancelled while acquiring lock on %s", stateFileName)
	}

	unlock := func() {
		stateFileMutex.Unlock()
	}
	// return empty list if file does not exist
	if _, err := os.Stat(stateFileName); errors.Is(err, os.ErrNotExist) {
		return &StateFile{PendingDeletes: []PendingDelete{}}, unlock, nil
	}

	// Try to read existing state file
	var state StateFile
	if data, err := os.ReadFile(stateFileName); err == nil {
		if err := json.Unmarshal(data, &state); err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Failed to parse existing %s: %s", stateFileName, err.Error()))
			// Continue with empty state
			state = StateFile{PendingDeletes: []PendingDelete{}}
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		// File exists but couldn't be read
		stateFileMutex.Unlock()
		return nil, nil, fmt.Errorf("failed to read %s: %w", stateFileName, err)
	}

	return &state, unlock, nil
}

func GetPendingDeletes(ctx context.Context) ([]PendingDelete, error) {
	state, unlock, err := readStateFile(ctx)
	if err != nil {
		return nil, err
	}
	defer unlock()

	return state.PendingDeletes, nil
}

// ManagePendingDelete manages pending delete operations (add, update, remove).
// For add: creates a new entry with the provided details
// For update: updates an existing entry using href as the key
// For remove: deletes an entry using href as the key
func ManagePendingDelete(ctx context.Context, operation PendingDeleteOperation, baseUrl, resourceType, href string) error {
	state, unlock, err := readStateFile(ctx)
	if err != nil {
		return err
	}
	defer unlock()

	// Find index of entry with matching href
	hrefIndex := -1
	for i, pendingDelete := range state.PendingDeletes {
		if pendingDelete.Href == href {
			hrefIndex = i
			break
		}
	}

	// Perform the requested operation
	switch operation {
	case AddPendingDelete:
		// Add the new pending delete entry (must not exist)
		if hrefIndex != -1 {
			tflog.Debug(ctx, fmt.Sprintf("Pending delete for %s already exists, skipping add", href))
			return nil
		}
		state.PendingDeletes = append(state.PendingDeletes, PendingDelete{
			BaseUrl:      baseUrl,
			ResourceType: resourceType,
			Href:         href,
		})
		tflog.Info(ctx, fmt.Sprintf("Added pending delete for %s", href))

	case UpdatePendingDelete:
		if hrefIndex == -1 {
			return fmt.Errorf("pending delete entry not found for href: %s", href)
		}
		// Update the existing entry
		state.PendingDeletes[hrefIndex] = PendingDelete{
			BaseUrl:      baseUrl,
			ResourceType: resourceType,
			Href:         href,
		}
		tflog.Info(ctx, fmt.Sprintf("Updated pending delete for %s", href))

	case RemovePendingDelete:
		if hrefIndex == -1 {
			return fmt.Errorf("pending delete entry not found for href: %s", href)
		}
		// Remove the entry by slicing
		state.PendingDeletes = append(state.PendingDeletes[:hrefIndex], state.PendingDeletes[hrefIndex+1:]...)
		tflog.Info(ctx, fmt.Sprintf("Removed pending delete for %s", href))

	default:
		return fmt.Errorf("unknown operation: %s", operation)
	}

	// Marshal to JSON with indentation
	jsonData, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal state to JSON: %w", err)
	}

	// Write back to file
	if err := os.WriteFile(stateFileName, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", stateFileName, err)
	}

	return nil
}

// SavePendingDelete adds a new pending delete entry (convenience function).
// Deprecated: Use ManagePendingDelete with AddPendingDelete operation instead.
func SavePendingDelete(ctx context.Context, baseUrl, resourceType, href string) error {
	return ManagePendingDelete(ctx, AddPendingDelete, baseUrl, resourceType, href)
}
