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
	"bytes"
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-log/tflogtest"
)

// Structured fields must be emitted by the subsystem logger. They are only
// picked up when set with tflog.SubsystemSetField; the root-level
// tflog.SetField silently drops them (SMC-66463 debugging was blind to
// status_code/body/headers because of this).
func TestLogContext_FieldsAreEmitted(t *testing.T) {
	var output bytes.Buffer
	ctx := tflogtest.RootLogger(context.Background(), &output)

	logCtx := NewLogContext(ctx, "server-1")
	logCtx.Error("Login failed with non-success status", map[string]any{
		"status_code": 500,
		"body":        "some body",
	})

	entries, err := tflogtest.MultilineJSONDecode(&output)
	if err != nil {
		t.Fatalf("Failed to decode log output: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("Expected 1 log entry, got %d", len(entries))
	}
	entry := entries[0]
	if entry["status_code"] != float64(500) {
		t.Errorf("Expected status_code field 500, got %v", entry["status_code"])
	}
	if entry["body"] != "some body" {
		t.Errorf("Expected body field, got %v", entry["body"])
	}
	if entry["mgt_server_id"] != "server-1" {
		t.Errorf("Expected mgt_server_id field from NewLogContext, got %v", entry["mgt_server_id"])
	}
}
