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

// Code not auto-generated
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-smc/internal/common"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

// engineNodeStatusResponse mirrors the relevant fields of the SMC
// Engine_Node_Status response returned by GET .../firewall_node/{key}/status.
type engineNodeStatusResponse struct {
	Name                string `json:"name"`
	EngineNodeStatus    string `json:"engine_node_status"`
	ConfigurationStatus string `json:"configuration_status"`
	ConnectivityStatus  string `json:"connectivity_status"`
}

// extractParentElementURL derives the cluster element URL from any sub-element URL.
//
// Example:
//
//	in:  http://smc1:18082/7.6/elements/single_fw/4858/internal_gateway/4859/generate_certificate
//	out: http://smc1:18082/7.6/elements/single_fw/4858
func extractParentElementURL(subURL string) (string, error) {
	u, err := url.Parse(subURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL %q: %w", subURL, err)
	}
	parts := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")
	for i, p := range parts {
		if p == "elements" && i+3 <= len(parts) {
			// parts[i+1] = element type, parts[i+2] = element key
			u.Path = "/" + strings.Join(parts[:i+3], "/")
			return u.String(), nil
		}
	}
	return "", fmt.Errorf("could not find /elements/{type}/{key} in URL %q", subURL)
}

// resolveNodeStatusHref discovers the status URL of the first node of a
// cluster element using ResolveRef2. The "*_node" glob matches any node type
// link (firewall_node, fw_cluster_node, master_node, etc.) so this works for
// single_fw, fw_cluster, and master_engine alike.
func resolveNodeStatusHref(ctx context.Context, client *smc.SmcClient, parentURL string) (string, error) {
	reader := &smc.GenericCRUDConfig{Client: client}
	return common.ResolveRef2(ctx, reader, parentURL,
		"*_node/result[0].href/link[?rel=='status'] | [0].href")
}

// waitForEngineStatus polls statusHref until statusGetter(response) == expectedStatus
// or the timeout elapses. progressFn is called with the current status before
// each sleep so callers can emit progress events.
func waitForEngineStatus(
	ctx context.Context,
	client *smc.SmcClient,
	statusHref string,
	statusGetter func(engineNodeStatusResponse) string,
	expectedStatus string,
	timeout time.Duration,
	pollInterval time.Duration,
	progressFn func(status engineNodeStatusResponse),
) error {
	timer := time.NewTimer(timeout)
	ticker := time.NewTicker(pollInterval)
	defer timer.Stop()
	defer ticker.Stop()

	var lastStatus engineNodeStatusResponse
	var lastErr error

	for {
		s, err := fetchEngineNodeStatus(ctx, client, statusHref)
		if err != nil {
			// Transient fetch errors (timeout, non-200, connection blip) shouldn't
			// abort the wait — keep polling and only surface this on timeout.
			lastErr = err
		} else {
			lastErr = nil
			lastStatus = s
			if statusGetter(s) == expectedStatus {
				return nil
			}
			progressFn(s)
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("context cancelled while waiting for engine status %q: %w", expectedStatus, ctx.Err())
		case <-timer.C:
			if lastErr != nil {
				return fmt.Errorf("engine did not reach %q within %v; last status fetch failed: %w",
					expectedStatus, timeout, lastErr)
			}
			return fmt.Errorf("engine did not reach %q within %v; last connectivity: %s (engine: %s, configuration: %s)",
				expectedStatus, timeout, lastStatus.ConnectivityStatus, lastStatus.EngineNodeStatus, lastStatus.ConfigurationStatus)
		case <-ticker.C:
		}
	}
}

func fetchEngineNodeStatus(ctx context.Context, client *smc.SmcClient, href string) (engineNodeStatusResponse, error) {
	respData, err := client.DoRequest(smc.Options{Method: "GET", URL: href, Context: ctx})
	if err != nil {
		return engineNodeStatusResponse{}, fmt.Errorf("GET %s failed: %w", href, err)
	}
	if respData.StatusCode != 200 {
		return engineNodeStatusResponse{}, fmt.Errorf("GET %s returned status %d: %s",
			href, respData.StatusCode, string(respData.Body))
	}
	var s engineNodeStatusResponse
	if err := json.Unmarshal(respData.Body, &s); err != nil {
		return engineNodeStatusResponse{}, fmt.Errorf("failed to parse engine status: %w", err)
	}
	return s, nil
}
