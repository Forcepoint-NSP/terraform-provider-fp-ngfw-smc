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

func isVpnUrl(url string) string {
	re := regexp.MustCompile(`^(https?://[^/]+/[0-9\.]+/elements/vpn/\d+)/`)
	match := re.FindStringSubmatch(url)

	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
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
	resp, err := smcClient.unsafeRequest(&editOpts)
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
	resp, err := smcClient.unsafeRequest(&cmdOpts)
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

func SaveAndCloseVpns(ctx context.Context) {
	mgr := GetVpnEditStateManager()
	mgr.VpnEditStates.Range(func(vpnUrl, value any) bool {
		for _, op := range []string{"save", "close"} {
			vpnState := value.(*VpnEditState)
			err := SendVpnCommand(vpnState.SmcClient, vpnUrl.(string), op)
			if err != nil {
				tflog.Error(ctx, fmt.Sprintf("Failed to save vpn %s: %s",
					vpnUrl.(string), err.Error()))
			}
		}
		return true
	})
}
