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

func TestHasPendingVpnEdits(t *testing.T) {
	const vpnUrl = "http://localhost:8082/7.5/elements/vpn/7"
	mgr := GetVpnEditStateManager()
	if HasPendingVpnEdits() {
		t.Fatal("expected no pending VPN edits at start")
	}
	mgr.VpnEditStates.Store(vpnUrl, &VpnEditState{})
	defer mgr.VpnEditStates.Delete(vpnUrl)
	if !HasPendingVpnEdits() {
		t.Error("expected pending VPN edits once a VPN is open")
	}
}

func TestIsVpnNameLookup(t *testing.T) {
	cases := map[string]bool{
		"http://h:8082/7.5/elements/vpn/1/tunnels#Gateway Tunnel *fw2 - *":       true,
		"http://h:8082/7.5/elements/vpn/1/central_gateway_node":                  false,
		"http://h:8082/7.5/elements/single_fw/12/internal_gateway#fw1*/vpn_site": false,
		"http://h:8082/7.5/elements/vpn/1":                                       false,
	}
	for ref, want := range cases {
		if got := IsVpnNameLookup(ref); got != want {
			t.Errorf("IsVpnNameLookup(%q) = %v, want %v", ref, got, want)
		}
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

// After a full re-login the old session's
// "open for edit" lock is gone server-side, so DropEntriesForClient must
// clear only the entries bound to that client, letting OpenVpnForEdit
// re-open them under the new session instead of short-circuiting.
func TestVpnEditStateManager_DropEntriesForClient(t *testing.T) {
	mgr := GetVpnEditStateManager()
	clientA := &SmcClient{BaseUrl: "http://a"}
	clientB := &SmcClient{BaseUrl: "http://b"}
	mgr.VpnEditStates.Store("http://a/7.4/elements/vpn/1", &VpnEditState{SmcClient: clientA})
	mgr.VpnEditStates.Store("http://b/7.4/elements/vpn/2", &VpnEditState{SmcClient: clientB})
	defer mgr.DropEntriesForClient(clientB)

	mgr.DropEntriesForClient(clientA)

	if _, ok := mgr.VpnEditStates.Load("http://a/7.4/elements/vpn/1"); ok {
		t.Error("expected clientA's VPN edit-state entry to be dropped")
	}
	if _, ok := mgr.VpnEditStates.Load("http://b/7.4/elements/vpn/2"); !ok {
		t.Error("expected clientB's VPN edit-state entry to be left untouched")
	}
}
