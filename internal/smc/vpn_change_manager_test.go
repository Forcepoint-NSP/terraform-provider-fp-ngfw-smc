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
	"context"
	"fmt"
	"testing"
)

func TestIsVpnUrl(t *testing.T) {
	url := "http://localhost:8082/7.4/elements/vpn/268438672/gateway_tree_nodes/central"
	prefix := isVpnUrl(url)
	if prefix != "http://localhost:8082/7.4/elements/vpn/268438672" {
		t.Errorf("expected vpn url prefix, got: %s", prefix)
	}

	prefix = isVpnUrl("http://localhost:8082/7.4/elements/single_fw/268571118")
	if prefix != "" {
		t.Errorf("expected empty string for non-vpn url, got: %s", prefix)
	}
}

func TestSaveAndCloseVpnsSkipsCloseAfterSaveFailure(t *testing.T) {
	const vpnUrl = "http://localhost:8082/7.4/elements/vpn/1"
	mgr := GetVpnEditStateManager()
	mgr.VpnEditStates.Store(vpnUrl, &VpnEditState{})

	var ops []string
	orig := sendVpnCommand
	defer func() { sendVpnCommand = orig }()
	sendVpnCommand = func(_ *SmcClient, _ string, command string) error {
		ops = append(ops, command)
		if command == "save" {
			return fmt.Errorf("save failed")
		}
		return nil
	}

	err := SaveAndCloseVpns(context.Background())

	if err == nil {
		t.Fatal("expected an error when save fails")
	}
	if len(ops) != 1 || ops[0] != "save" {
		t.Errorf("expected only save to be sent, got: %v", ops)
	}
	if _, ok := mgr.VpnEditStates.Load(vpnUrl); ok {
		t.Error("tracked state must be dropped after a failure")
	}
}

func TestSaveAndCloseVpnsSavesThenCloses(t *testing.T) {
	const vpnUrl = "http://localhost:8082/7.4/elements/vpn/2"
	mgr := GetVpnEditStateManager()
	mgr.VpnEditStates.Store(vpnUrl, &VpnEditState{})

	var ops []string
	orig := sendVpnCommand
	defer func() { sendVpnCommand = orig }()
	sendVpnCommand = func(_ *SmcClient, _ string, command string) error {
		ops = append(ops, command)
		return nil
	}

	err := SaveAndCloseVpns(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ops) != 2 || ops[0] != "save" || ops[1] != "close" {
		t.Errorf("expected save then close, got: %v", ops)
	}
	if _, ok := mgr.VpnEditStates.Load(vpnUrl); ok {
		t.Error("tracked state must be dropped after success")
	}
}
