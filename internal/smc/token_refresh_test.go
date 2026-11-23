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
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflogtest"
)

// SMC API session tokens now expire server-side (default 30 min)
// instead of staying valid forever. These tests cover the client-side
// recovery: RefreshToken() itself, and DoRequest's automatic retry-on-401
// (refresh_token first, full re-login as fallback).

func TestRefreshToken_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || !strings.Contains(r.URL.Path, "/system/refresh_token") {
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "old-token" {
			t.Errorf("Expected current token in Authorization header, got %q", got)
		}
		w.Header().Set("Authorization", "new-token")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:        server.URL,
		APIVersion:     "v1",
		Token:          "old-token",
		UseAuthHeader:  true,
		LoginConfirmed: true,
	}

	if err := client.RefreshToken(context.Background()); err != nil {
		t.Fatalf("RefreshToken failed: %v", err)
	}
	if client.Token != "new-token" {
		t.Errorf("Expected Token to be updated to new-token, got %s", client.Token)
	}
}

func TestRefreshToken_NoActiveSession(t *testing.T) {
	client := &SmcClient{LoginConfirmed: false}
	if err := client.RefreshToken(context.Background()); err == nil {
		t.Fatal("Expected an error when there is no active session")
	}
}

func TestRefreshToken_NonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// Older SMC without the refresh_token service.
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:        server.URL,
		APIVersion:     "v1",
		Token:          "old-token",
		UseAuthHeader:  true,
		LoginConfirmed: true,
	}
	if err := client.RefreshToken(context.Background()); err == nil {
		t.Fatal("Expected an error on a non-200 refresh_token response")
	}
}

func TestRefreshToken_MissingAuthorizationHeaderInResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK) // no Authorization header set
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:        server.URL,
		APIVersion:     "v1",
		Token:          "old-token",
		UseAuthHeader:  true,
		LoginConfirmed: true,
	}
	if err := client.RefreshToken(context.Background()); err == nil {
		t.Fatal("Expected an error when the response carries no new token")
	}
}

func TestDoRequest_RecoversViaRefreshTokenOn401(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var resourceCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/system/refresh_token"):
			w.Header().Set("Authorization", "new-token")
			w.WriteHeader(http.StatusOK)
		case strings.Contains(r.URL.Path, "/login"):
			t.Fatal("Full login should not be attempted when refresh_token succeeds")
		case strings.Contains(r.URL.Path, "/elements/some_resource"):
			n := atomic.AddInt32(&resourceCalls, 1)
			if n == 1 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if got := r.Header.Get("Authorization"); got != "new-token" {
				t.Errorf("Retry should use the refreshed token, got %q", got)
			}
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	var logOutput bytes.Buffer
	ctx := tflogtest.RootLogger(context.Background(), &logOutput)
	client := &SmcClient{
		BaseUrl:        server.URL,
		APIVersion:     "v1",
		Token:          "old-token",
		UseAuthHeader:  true,
		LoginConfirmed: true,
		SmcContext:     NewLogContext(ctx, "server-1"),
	}

	resp, err := client.DoRequest(Options{
		Method:  "GET",
		URL:     fmt.Sprintf("%s/v1/elements/some_resource", server.URL),
		Context: ctx,
	})
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 after recovery, got %d", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&resourceCalls); got != 2 {
		t.Errorf("Expected 2 calls to the resource endpoint (401 then retry), got %d", got)
	}

	// This is the log assertion a robot/STF test can reproduce by grepping
	// the TF_LOG_PATH debug log file for these exact markers, to prove a
	// long-running apply recovered via a lightweight refresh and not a full
	// re-login.
	entries, err := tflogtest.MultilineJSONDecode(&logOutput)
	if err != nil {
		t.Fatalf("Failed to decode log output: %v", err)
	}
	foundRefreshLog := false
	for _, e := range entries {
		msg, _ := e["@message"].(string)
		if strings.Contains(msg, "token refreshed via refresh_token endpoint") {
			foundRefreshLog = true
		}
		if strings.Contains(msg, "full login") {
			t.Errorf("Did not expect a full-login log entry, got: %s", msg)
		}
	}
	if !foundRefreshLog {
		t.Error("Expected a log entry confirming the token was refreshed via refresh_token")
	}
}

func TestDoRequest_FallsBackToFullLoginWhenRefreshTokenUnavailable(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var loginCalls int32
	var resourceCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/system/refresh_token"):
			// Older SMC (or backport not applied): endpoint does not exist.
			w.WriteHeader(http.StatusNotFound)
		case r.Method == "POST" && strings.Contains(r.URL.Path, "/login"):
			atomic.AddInt32(&loginCalls, 1)
			w.Header().Set("Authorization", "relogin-token")
			w.WriteHeader(http.StatusOK)
		case strings.Contains(r.URL.Path, "/elements/some_resource"):
			n := atomic.AddInt32(&resourceCalls, 1)
			if n == 1 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if got := r.Header.Get("Authorization"); got != "relogin-token" {
				t.Errorf("Retry should use the token from the full re-login, got %q", got)
			}
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	var logOutput bytes.Buffer
	ctx := tflogtest.RootLogger(context.Background(), &logOutput)
	client := &SmcClient{
		BaseUrl:        server.URL,
		APIVersion:     "v1",
		APIKey:         "test-key",
		Token:          "old-token",
		UseAuthHeader:  true,
		LoginConfirmed: true,
		SmcContext:     NewLogContext(ctx, "server-1"),
	}

	resp, err := client.DoRequest(Options{
		Method:  "GET",
		URL:     fmt.Sprintf("%s/v1/elements/some_resource", server.URL),
		Context: ctx,
	})
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 after recovery, got %d", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&loginCalls); got != 1 {
		t.Errorf("Expected exactly 1 full re-login call, got %d", got)
	}

	entries, err := tflogtest.MultilineJSONDecode(&logOutput)
	if err != nil {
		t.Fatalf("Failed to decode log output: %v", err)
	}
	foundFallbackLog := false
	for _, e := range entries {
		if msg, _ := e["@message"].(string); strings.Contains(msg, "re-authenticated via full login") {
			foundFallbackLog = true
		}
	}
	if !foundFallbackLog {
		t.Error("Expected a log entry confirming the fallback to full login")
	}
}

func TestDoRequest_CookieOnlySession_SkipsRefreshTokenEndpoint(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var resourceCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/system/refresh_token"):
			t.Fatal("refresh_token should not be used for a JSESSIONID cookie-only session")
		case r.Method == "POST" && strings.Contains(r.URL.Path, "/login"):
			w.Header().Set("Set-Cookie", "JSESSIONID=relogin-session; Path=/")
			w.WriteHeader(http.StatusOK)
		case strings.Contains(r.URL.Path, "/elements/some_resource"):
			n := atomic.AddInt32(&resourceCalls, 1)
			if n == 1 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:        server.URL,
		APIVersion:     "v1",
		APIKey:         "test-key",
		Token:          "old-session",
		UseAuthHeader:  false,
		LoginConfirmed: true,
	}

	resp, err := client.DoRequest(Options{
		Method: "GET",
		URL:    fmt.Sprintf("%s/v1/elements/some_resource", server.URL),
	})
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 after recovery, got %d", resp.StatusCode)
	}
}

func TestPing_ReturnsTrueOn200(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:        server.URL,
		APIVersion:     "v1",
		Token:          "token",
		UseAuthHeader:  true,
		LoginConfirmed: true,
	}
	if !client.ping(context.Background()) {
		t.Error("Expected ping to return true on 200")
	}
}

func TestPing_ReturnsFalseWhenSessionCannotBeRecovered(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	restoreDelay := loginRetryBaseDelay
	loginRetryBaseDelay = time.Millisecond
	defer func() { loginRetryBaseDelay = restoreDelay }()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// Everything (ping, refresh_token, login) is unauthorized: the
		// session is truly dead and unrecoverable.
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:        server.URL,
		APIVersion:     "v1",
		APIKey:         "test-key",
		Token:          "dead-token",
		UseAuthHeader:  true,
		LoginConfirmed: true,
	}
	if client.ping(context.Background()) {
		t.Error("Expected ping to return false when the session cannot be recovered")
	}
}

func TestLogout_SkipsAuthRetryOn401(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var loginCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/login"):
			atomic.AddInt32(&loginCalls, 1)
			w.WriteHeader(http.StatusOK)
		case strings.Contains(r.URL.Path, "/logout"):
			// The session is already dead server-side; logging back in
			// just to log out again would be pointless.
			w.WriteHeader(http.StatusUnauthorized)
		default:
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:        server.URL,
		APIVersion:     "v1",
		Token:          "old-token",
		UseAuthHeader:  true,
		LoginConfirmed: true,
	}

	if err := client.Logout(context.Background()); err != nil {
		t.Fatalf("Logout should not return an error even on 401: %v", err)
	}
	if got := atomic.LoadInt32(&loginCalls); got != 0 {
		t.Errorf("Expected no re-login attempt on logout 401, got %d calls", got)
	}
}

// Proactive refresh (TokenRefreshInterval) keeps a long-running
// apply ahead of the SMC's configured SMC_API_TOKEN_VALIDITY instead of
// relying solely on the reactive 401 recovery above.

func TestDoRequest_ProactivelyRefreshesTokenAfterInterval(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var refreshCalls, resourceCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/system/refresh_token"):
			atomic.AddInt32(&refreshCalls, 1)
			w.Header().Set("Authorization", "new-token")
			w.WriteHeader(http.StatusOK)
		case strings.Contains(r.URL.Path, "/elements/some_resource"):
			n := atomic.AddInt32(&resourceCalls, 1)
			if n == 1 {
				if got := r.Header.Get("Authorization"); got != "new-token" {
					t.Errorf("Expected the request to use the proactively refreshed token, got %q", got)
				}
			}
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	var logOutput bytes.Buffer
	ctx := tflogtest.RootLogger(context.Background(), &logOutput)
	client := &SmcClient{
		BaseUrl:              server.URL,
		APIVersion:           "v1",
		Token:                "old-token",
		UseAuthHeader:        true,
		LoginConfirmed:       true,
		TokenRefreshInterval: time.Millisecond,
		TokenIssuedAt:        time.Now().Add(-time.Hour), // long past the interval
		SmcContext:           NewLogContext(ctx, "server-1"),
	}

	resp, err := client.DoRequest(Options{
		Method:  "GET",
		URL:     fmt.Sprintf("%s/v1/elements/some_resource", server.URL),
		Context: ctx,
	})
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&refreshCalls); got != 1 {
		t.Errorf("Expected exactly 1 proactive refresh_token call, got %d", got)
	}
	if got := atomic.LoadInt32(&resourceCalls); got != 1 {
		t.Errorf("Expected exactly 1 resource call (no 401 should ever occur), got %d", got)
	}

	entries, err := tflogtest.MultilineJSONDecode(&logOutput)
	if err != nil {
		t.Fatalf("Failed to decode log output: %v", err)
	}
	foundProactiveLog := false
	for _, e := range entries {
		msg, _ := e["@message"].(string)
		if strings.Contains(msg, "proactively refreshed via refresh_token endpoint") {
			foundProactiveLog = true
		}
		if strings.Contains(msg, "Received HTTP 401") {
			t.Errorf("Did not expect any 401 to occur with proactive refresh, got log: %s", msg)
		}
	}
	if !foundProactiveLog {
		t.Error("Expected a log entry confirming the token was proactively refreshed")
	}
}

func TestDoRequest_DoesNotProactivelyRefreshWhenDisabled(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var refreshCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/system/refresh_token") {
			atomic.AddInt32(&refreshCalls, 1)
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:        server.URL,
		APIVersion:     "v1",
		Token:          "old-token",
		UseAuthHeader:  true,
		LoginConfirmed: true,
		// TokenRefreshInterval left at zero value: proactive refresh disabled.
		TokenIssuedAt: time.Now().Add(-time.Hour),
	}

	_, err := client.DoRequest(Options{
		Method: "GET",
		URL:    fmt.Sprintf("%s/v1/elements/some_resource", server.URL),
	})
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}
	if got := atomic.LoadInt32(&refreshCalls); got != 0 {
		t.Errorf("Expected no proactive refresh when TokenRefreshInterval is disabled, got %d calls", got)
	}
}

func TestDoRequest_DoesNotProactivelyRefreshBeforeIntervalElapsed(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var refreshCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/system/refresh_token") {
			atomic.AddInt32(&refreshCalls, 1)
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:              server.URL,
		APIVersion:           "v1",
		Token:                "old-token",
		UseAuthHeader:        true,
		LoginConfirmed:       true,
		TokenRefreshInterval: time.Hour,
		TokenIssuedAt:        time.Now(), // just issued, well within the interval
	}

	_, err := client.DoRequest(Options{
		Method: "GET",
		URL:    fmt.Sprintf("%s/v1/elements/some_resource", server.URL),
	})
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}
	if got := atomic.LoadInt32(&refreshCalls); got != 0 {
		t.Errorf("Expected no proactive refresh before the interval elapses, got %d calls", got)
	}
}

func TestDoRequest_ProactiveRefreshFailureIsSwallowed(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var resourceCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/system/refresh_token"):
			// Older SMC without the refresh_token service.
			w.WriteHeader(http.StatusNotFound)
		case strings.Contains(r.URL.Path, "/elements/some_resource"):
			atomic.AddInt32(&resourceCalls, 1)
			if got := r.Header.Get("Authorization"); got != "old-token" {
				t.Errorf("Expected the request to still use the old token when refresh failed, got %q", got)
			}
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	var logOutput bytes.Buffer
	ctx := tflogtest.RootLogger(context.Background(), &logOutput)
	client := &SmcClient{
		BaseUrl:              server.URL,
		APIVersion:           "v1",
		Token:                "old-token",
		UseAuthHeader:        true,
		LoginConfirmed:       true,
		TokenRefreshInterval: time.Millisecond,
		TokenIssuedAt:        time.Now().Add(-time.Hour),
		SmcContext:           NewLogContext(ctx, "server-1"),
	}

	resp, err := client.DoRequest(Options{
		Method:  "GET",
		URL:     fmt.Sprintf("%s/v1/elements/some_resource", server.URL),
		Context: ctx,
	})
	if err != nil {
		t.Fatalf("DoRequest should still succeed using the old token when the proactive refresh fails: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&resourceCalls); got != 1 {
		t.Errorf("Expected the resource call to proceed despite the failed proactive refresh, got %d calls", got)
	}
}

func TestDoRequest_ProactiveRefreshFailureBacksOffForAFullInterval(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var refreshCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/system/refresh_token"):
			atomic.AddInt32(&refreshCalls, 1)
			w.WriteHeader(http.StatusNotFound)
		case strings.Contains(r.URL.Path, "/elements/some_resource"):
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:              server.URL,
		APIVersion:           "v1",
		Token:                "old-token",
		UseAuthHeader:        true,
		LoginConfirmed:       true,
		TokenRefreshInterval: time.Hour,
		TokenIssuedAt:        time.Now().Add(-2 * time.Hour),
		SmcContext:           NewLogContext(context.Background(), "server-1"),
	}

	for i := 0; i < 3; i++ {
		_, err := client.DoRequest(Options{
			Method: "GET",
			URL:    fmt.Sprintf("%s/v1/elements/some_resource", server.URL),
		})
		if err != nil {
			t.Fatalf("request %d failed: %v", i, err)
		}
	}

	if got := atomic.LoadInt32(&refreshCalls); got != 1 {
		t.Errorf("Expected a single refresh attempt per interval after a failure, got %d", got)
	}
}

func TestLogout_SkipsProactiveRefresh(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var refreshCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/system/refresh_token"):
			atomic.AddInt32(&refreshCalls, 1)
			w.WriteHeader(http.StatusOK)
		case strings.Contains(r.URL.Path, "/logout"):
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:              server.URL,
		APIVersion:           "v1",
		Token:                "old-token",
		UseAuthHeader:        true,
		LoginConfirmed:       true,
		TokenRefreshInterval: time.Millisecond,
		TokenIssuedAt:        time.Now().Add(-time.Hour), // long past the interval
	}

	if err := client.Logout(context.Background()); err != nil {
		t.Fatalf("Logout failed: %v", err)
	}
	if got := atomic.LoadInt32(&refreshCalls); got != 0 {
		t.Errorf("Expected no proactive refresh before logging out, got %d calls", got)
	}
}

// In parallel mode, several concurrent requests
// can share the same expired token and all observe a 401 before any of them
// finishes recovering it. recoverExpiredSession must not blindly refresh or
// re-login again once it acquires loginMutex if another goroutine already
// fixed the session in the meantime -- doing so would invalidate that other
// goroutine's brand-new token out from under its own in-flight retry.
func TestRecoverExpiredSession_SkipsWhenAlreadyRecoveredByConcurrentRequest(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var refreshCalls, loginCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/system/refresh_token"):
			atomic.AddInt32(&refreshCalls, 1)
			w.WriteHeader(http.StatusOK)
		case strings.Contains(r.URL.Path, "/login"):
			atomic.AddInt32(&loginCalls, 1)
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:    server.URL,
		APIVersion: "v1",
		// A concurrent request already refreshed the token by the time this
		// one acquires loginMutex.
		Token:          "new-token",
		UseAuthHeader:  true,
		LoginConfirmed: true,
	}

	// This request's own snapshot saw "old-token" before its 401.
	if err := client.recoverExpiredSession(context.Background(), "old-token"); err != nil {
		t.Fatalf("recoverExpiredSession failed: %v", err)
	}
	if got := atomic.LoadInt32(&refreshCalls); got != 0 {
		t.Errorf("Expected no refresh_token call when the session was already recovered, got %d", got)
	}
	if got := atomic.LoadInt32(&loginCalls); got != 0 {
		t.Errorf("Expected no re-login call when the session was already recovered, got %d", got)
	}
	if client.Token != "new-token" {
		t.Errorf("Expected the concurrently-recovered token to be left untouched, got %s", client.Token)
	}
}

// A full re-login replaces the session
// entirely, so a VPN this client already had open for edit under the old
// session is not open under the new one. Re-running BeforeRequestHooks on
// the retry (see TestDoRequest_RecoversViaRefreshTokenOn401) is not enough
// by itself -- OpenVpnForEdit short-circuits on a cached "already open"
// marker. recoverExpiredSession must drop that marker on the full-login
// fallback so the re-run hook actually re-opens the VPN before the retried
// write, instead of letting it fail 400/409 against a session that never
// opened it.
func TestDoRequest_FullReloginReopensVpnAfterExpiry(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	var vpnOpenCalls, resourceCalls, loginCalls int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/system/refresh_token"):
			// Endpoint unavailable: force the full re-login fallback.
			w.WriteHeader(http.StatusNotFound)
		case r.Method == "POST" && strings.HasSuffix(r.URL.Path, "/open"):
			atomic.AddInt32(&vpnOpenCalls, 1)
			w.WriteHeader(http.StatusOK)
		case r.Method == "POST" && strings.Contains(r.URL.Path, "/login"):
			atomic.AddInt32(&loginCalls, 1)
			w.Header().Set("Authorization", "relogin-token")
			w.WriteHeader(http.StatusOK)
		case strings.Contains(r.URL.Path, "/elements/vpn/1/gateway_tree_nodes"):
			n := atomic.AddInt32(&resourceCalls, 1)
			if n == 1 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if got := r.Header.Get("Authorization"); got != "relogin-token" {
				t.Errorf("Retry should use the token from the full re-login, got %q", got)
			}
			w.WriteHeader(http.StatusOK)
		default:
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:            server.URL,
		APIVersion:         "7.4",
		APIKey:             "test-key",
		Token:              "old-token",
		UseAuthHeader:      true,
		LoginConfirmed:     true,
		BeforeRequestHooks: []func(*SmcClient, *Options) error{OpenVpnForEdit},
	}

	vpnUrl := fmt.Sprintf("%s/7.4/elements/vpn/1", server.URL)
	// The VPN is already open for edit under this client's *old* session
	// (e.g. an earlier request in the same apply opened it).
	GetVpnEditStateManager().VpnEditStates.Store(vpnUrl, &VpnEditState{SmcClient: client})
	t.Cleanup(func() { GetVpnEditStateManager().VpnEditStates.Delete(vpnUrl) })

	resp, err := client.DoRequest(Options{
		Method: "PUT",
		URL:    vpnUrl + "/gateway_tree_nodes/central",
	})
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 after recovery, got %d", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&loginCalls); got != 1 {
		t.Errorf("Expected exactly 1 full re-login call, got %d", got)
	}
	if got := atomic.LoadInt32(&vpnOpenCalls); got != 1 {
		t.Errorf("Expected the VPN to be re-opened once under the new session, got %d calls", got)
	}
}
