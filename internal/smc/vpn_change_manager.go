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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"regexp"
	"strings"
	"sync"
)

type VpnEditState struct {
	SmcClient *SmcClient
	// we don't actually need the state for now ("open", "saved",
	// "closed"): if an entry is present in vpnEditStates, it means the
	// vpn is being edited.
	// the vpn are saved and closed when logging out from the smcClient
}

// keeps track of all the edit states for VPNs across multiple
// smcClient sessions
type VpnEditStateManager struct {
	// key: vpn url eg  "https://host:port/7.4/elements/vpn/268438663"
	// value: VpnEditState
	VpnEditStates sync.Map
}

var vpnEditStateManagerInstance = VpnEditStateManager{}

func GetVpnEditStateManager() *VpnEditStateManager {
	return &vpnEditStateManagerInstance
}

// DropEntriesForClient removes any VPN edit-state entries bound to the given
// client. Used when a client's session is fully replaced by a re-login
// (the full-login fallback of recoverExpiredSession): the old
// session's "open for edit" lock on the SMC dies with it, so OpenVpnForEdit
// must be allowed to re-open the VPN under the new session on the next
// request instead of short-circuiting on the now-stale "already open"
// marker (which would otherwise let a retried write fail with 400/409, and
// leave SaveAndCloseVpns acting on a session that never re-opened it).
func (mgr *VpnEditStateManager) DropEntriesForClient(smcClient *SmcClient) {
	mgr.VpnEditStates.Range(func(vpnUrl, value any) bool {
		if state, ok := value.(*VpnEditState); ok && state.SmcClient == smcClient {
			mgr.VpnEditStates.Delete(vpnUrl)
		}
		return true
	})
}

func isVpnUrl(url string) string {
	re := regexp.MustCompile(`^(https?://[^/]+/[0-9\.]+/elements/vpn/\d+)/`)
	match := re.FindStringSubmatch(url)

	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

// HasPendingVpnEdits reports whether a VPN is currently open for edit, i.e.
// its changes are not committed to the SMC yet.
func HasPendingVpnEdits() bool {
	return len(GetVpnEditStateManager().trackedUrls()) > 0
}

// IsVpnNameLookup reports whether ref is a VPN sub-element reference with a
// "#name" fragment to resolve (e.g. ".../elements/vpn/1/tunnels#Gateway Tunnel*").
func IsVpnNameLookup(ref string) bool {
	return isVpnUrl(ref) != "" && strings.Contains(ref, "#")
}

func OpenVpnForEdit(smcClient *SmcClient, opts *Options) error {
	vpnUrl := isVpnUrl(opts.URL)
	if vpnUrl == "" {
		// not a vpn url
		return nil
	}
	mgr := GetVpnEditStateManager()
	_, ok := mgr.VpnEditStates.Load(vpnUrl)
	if ok {
		// already "open" for edit
		return nil
	}

	// you cannot have concurrent access for the same smcClient
	// because of session lock. So no need to lock here.
	mgr.VpnEditStates.LoadOrStore(vpnUrl, &VpnEditState{
		SmcClient: smcClient,
	})

	editOpts := Options{
		URL:     vpnUrl + "/open",
		Method:  "POST",
		Headers: opts.Headers,
		Context: opts.Context,
	}
	// call unsafeRequest directly because we already have the lock
	resp, err := smcClient.unsafeRequest(&editOpts, smcClient.currentAuth())
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("failed to open vpn for edit, status code: %d",
			resp.StatusCode)
	}
	return nil
}

func SendVpnCommand(smcClient *SmcClient, vpnUrl string, command string) error {
	cmdOpts := Options{
		URL:     vpnUrl + "/" + command,
		Method:  "POST",
		Headers: smcClient.GetAuthHeaders(),
	}
	// call unsafeRequest directly because we already have the lock
	resp, err := smcClient.unsafeRequest(&cmdOpts, smcClient.currentAuth())
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return fmt.Errorf("failed to %s vpn for edit, status code: %d",
			command,
			resp.StatusCode)
	}
	return nil

}

// trackedUrls returns the vpnUrl of every currently-tracked VPN edit state.
func (mgr *VpnEditStateManager) trackedUrls() []string {
	var urls []string
	mgr.VpnEditStates.Range(func(vpnUrl, _ any) bool {
		urls = append(urls, vpnUrl.(string))
		return true
	})
	return urls
}

// sendVpnCommand is a variable so tests can stub the SMC calls.
var sendVpnCommand = SendVpnCommand

// SaveAndCloseVpns is not reentrant with DoRequest: same per-host lock. Returns any save/close failures.
// A VPN whose save fails is not closed (closing would discard its edits). Its tracked state is
// dropped in every case, so a failure cannot block later uploads/refreshes.
func SaveAndCloseVpns(ctx context.Context) error {
	mgr := GetVpnEditStateManager()
	var errs []string
	for _, vpnUrl := range mgr.trackedUrls() {
		lockId := GetLockId(vpnUrl)
		GetLockRequests().Enter(lockId)

		// Locked, so a second caller sees the state already gone.
		value, ok := mgr.VpnEditStates.Load(vpnUrl)
		if ok {
			vpnState := value.(*VpnEditState)
			for _, op := range []string{"save", "close"} {
				err := sendVpnCommand(vpnState.SmcClient, vpnUrl, op)
				if err == nil {
					continue
				}
				tflog.Error(ctx, fmt.Sprintf("Failed to %s vpn %s: %s", op, vpnUrl, err.Error()))
				errs = append(errs, fmt.Sprintf("%s vpn %s: %s", op, vpnUrl, err.Error()))
				break
			}
			mgr.VpnEditStates.Delete(vpnUrl)
		}

		GetLockRequests().Leave(lockId)
	}
	if len(errs) > 0 {
		return fmt.Errorf("failed to save/close %d vpn(s): %s", len(errs), strings.Join(errs, "; "))
	}
	return nil
}
