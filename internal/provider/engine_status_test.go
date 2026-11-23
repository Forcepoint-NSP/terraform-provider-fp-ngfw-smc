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

package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflogtest"
	"github.com/stretchr/testify/assert"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

func TestExtractParentElementURL(t *testing.T) {
	cases := []struct {
		in   string
		want string
		err  bool
	}{
		{
			in:   "http://smc1:18082/7.6/elements/single_fw/4858/internal_gateway/4859/generate_certificate",
			want: "http://smc1:18082/7.6/elements/single_fw/4858",
		},
		{
			in:   "http://smc1:18082/7.6/elements/fw_cluster/123/internal_gateway/456/gateway_certificate_request/789/certificate_import",
			want: "http://smc1:18082/7.6/elements/fw_cluster/123",
		},
		{
			in:  "http://smc1:18082/7.6/no-elements-path/single_fw/4858",
			err: true,
		},
	}
	for _, c := range cases {
		got, err := extractParentElementURL(c.in)
		if c.err {
			assert.Error(t, err, "input: %s", c.in)
		} else {
			assert.NoError(t, err, "input: %s", c.in)
			assert.Equal(t, c.want, got)
		}
	}
}

// newTestClientWithStatusServer starts an httptest server that handles SMC login
// and serves engine node status JSON from statusFn on GET /status.
func newTestClientWithStatusServer(t *testing.T, statusFn func() engineNodeStatusResponse) (*smc.SmcClient, func()) {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && strings.Contains(r.URL.Path, "login") {
			w.Header().Set("Set-Cookie", "JSESSIONID=test-session; Path=/")
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method == "GET" && strings.Contains(r.URL.Path, "status") {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(statusFn())
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	smc.GetClientManager().CleanupAll(context.Background())
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client, err := smc.GetOrCreateClient(ctx, server.URL, "7.6", "test-key", false, "", "", 0)
	if err != nil {
		t.Fatalf("GetOrCreateClient: %v", err)
	}
	return client, server.Close
}

func TestWaitForEngineStatus_AlreadyConnected(t *testing.T) {
	client, cleanup := newTestClientWithStatusServer(t, func() engineNodeStatusResponse {
		return engineNodeStatusResponse{Name: "node1", ConnectivityStatus: "OK", EngineNodeStatus: "Online"}
	})
	defer cleanup()

	statusHref := client.BaseUrl + "/7.6/elements/single_fw/1/firewall_node/1/status"
	err := waitForEngineStatus(
		context.Background(), client, statusHref,
		func(s engineNodeStatusResponse) string { return s.ConnectivityStatus },
		"OK",
		2*time.Second, 10*time.Millisecond,
		func(engineNodeStatusResponse) {},
	)
	assert.NoError(t, err)
}

func TestWaitForEngineStatus_PollsUntilConnected(t *testing.T) {
	var callCount atomic.Int32
	client, cleanup := newTestClientWithStatusServer(t, func() engineNodeStatusResponse {
		n := callCount.Add(1)
		if n < 3 {
			return engineNodeStatusResponse{Name: "node1", ConnectivityStatus: "NO_CONNECTIONS", EngineNodeStatus: "Not Monitored"}
		}
		return engineNodeStatusResponse{Name: "node1", ConnectivityStatus: "OK", EngineNodeStatus: "No Policy Installed"}
	})
	defer cleanup()

	statusHref := client.BaseUrl + "/7.6/elements/single_fw/1/firewall_node/1/status"
	err := waitForEngineStatus(
		context.Background(), client, statusHref,
		func(s engineNodeStatusResponse) string { return s.ConnectivityStatus },
		"OK",
		5*time.Second, 10*time.Millisecond,
		func(engineNodeStatusResponse) {},
	)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, callCount.Load(), int32(3))
}

// newTestClientWithFlakyStatusServer is like newTestClientWithStatusServer but
// statusFn also controls the HTTP status code, letting tests simulate transient
// status-fetch failures (non-200 responses).
func newTestClientWithFlakyStatusServer(t *testing.T, statusFn func() (int, engineNodeStatusResponse)) (*smc.SmcClient, func()) {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && strings.Contains(r.URL.Path, "login") {
			w.Header().Set("Set-Cookie", "JSESSIONID=test-session; Path=/")
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method == "GET" && strings.Contains(r.URL.Path, "status") {
			code, s := statusFn()
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			_ = json.NewEncoder(w).Encode(s)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	smc.GetClientManager().CleanupAll(context.Background())
	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)
	client, err := smc.GetOrCreateClient(ctx, server.URL, "7.6", "test-key", false, "", "", 0)
	if err != nil {
		t.Fatalf("GetOrCreateClient: %v", err)
	}
	return client, server.Close
}

func TestWaitForEngineStatus_SurvivesTransientFetchError(t *testing.T) {
	var callCount atomic.Int32
	client, cleanup := newTestClientWithFlakyStatusServer(t, func() (int, engineNodeStatusResponse) {
		n := callCount.Add(1)
		if n < 3 {
			return http.StatusInternalServerError, engineNodeStatusResponse{}
		}
		return http.StatusOK, engineNodeStatusResponse{Name: "node1", ConnectivityStatus: "OK", EngineNodeStatus: "No Policy Installed"}
	})
	defer cleanup()

	statusHref := client.BaseUrl + "/7.6/elements/single_fw/1/firewall_node/1/status"
	err := waitForEngineStatus(
		context.Background(), client, statusHref,
		func(s engineNodeStatusResponse) string { return s.ConnectivityStatus },
		"OK",
		5*time.Second, 10*time.Millisecond,
		func(engineNodeStatusResponse) {},
	)
	assert.NoError(t, err, "a transient fetch error should not abort the wait")
	assert.GreaterOrEqual(t, callCount.Load(), int32(3))
}

func TestWaitForEngineStatus_TimeoutAfterOnlyErrors(t *testing.T) {
	client, cleanup := newTestClientWithFlakyStatusServer(t, func() (int, engineNodeStatusResponse) {
		return http.StatusInternalServerError, engineNodeStatusResponse{}
	})
	defer cleanup()

	statusHref := client.BaseUrl + "/7.6/elements/single_fw/1/firewall_node/1/status"
	err := waitForEngineStatus(
		context.Background(), client, statusHref,
		func(s engineNodeStatusResponse) string { return s.ConnectivityStatus },
		"OK",
		100*time.Millisecond, 20*time.Millisecond,
		func(engineNodeStatusResponse) {},
	)
	assert.ErrorContains(t, err, "did not reach")
	assert.ErrorContains(t, err, "status fetch failed")
}

func TestWaitForEngineStatus_Timeout(t *testing.T) {
	client, cleanup := newTestClientWithStatusServer(t, func() engineNodeStatusResponse {
		return engineNodeStatusResponse{Name: "node1", ConnectivityStatus: "NO_CONNECTIONS", EngineNodeStatus: "Not Monitored"}
	})
	defer cleanup()

	statusHref := client.BaseUrl + "/7.6/elements/single_fw/1/firewall_node/1/status"
	err := waitForEngineStatus(
		context.Background(), client, statusHref,
		func(s engineNodeStatusResponse) string { return s.ConnectivityStatus },
		"OK",
		100*time.Millisecond, 20*time.Millisecond,
		func(engineNodeStatusResponse) {},
	)
	assert.ErrorContains(t, err, "did not reach")
}
