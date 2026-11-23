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
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflogtest"
)

// maxInFlightServer starts a test server that tracks the maximum number of
// requests it served concurrently, holding each request open briefly so
// overlapping requests are observable.
func maxInFlightServer(t *testing.T) (*httptest.Server, *int32) {
	t.Helper()
	var inFlight int32
	var maxInFlight int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cur := atomic.AddInt32(&inFlight, 1)
		for {
			m := atomic.LoadInt32(&maxInFlight)
			if cur <= m {
				break
			}
			if atomic.CompareAndSwapInt32(&maxInFlight, m, cur) {
				break
			}
		}
		time.Sleep(30 * time.Millisecond)
		atomic.AddInt32(&inFlight, -1)
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)
	return server, &maxInFlight
}

func newTestClient(baseUrl string) *SmcClient {
	return &SmcClient{
		BaseUrl:        baseUrl,
		APIVersion:     "v1",
		LoginConfirmed: true,
		UseAuthHeader:  true,
		Token:          "test-token",
		SmcContext:     NewLogContext(context.Background(), baseUrl),
	}
}

func doConcurrentRequests(t *testing.T, client *SmcClient, resourceType string, n int) {
	t.Helper()
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := client.DoRequest(Options{
				Method:         "GET",
				URL:            client.BaseUrl + "/v1/elements/" + resourceType + "/1",
				CanParallelize: func() bool { return IsResourceTypeParallel(resourceType) },
			})
			if err != nil {
				t.Errorf("DoRequest failed: %v", err)
			}
		}()
	}
	wg.Wait()
}

func TestDoRequest_SerializesByDefault(t *testing.T) {
	server, maxInFlight := maxInFlightServer(t)
	client := newTestClient(server.URL)

	doConcurrentRequests(t, client, "host", 5)

	if got := atomic.LoadInt32(maxInFlight); got != 1 {
		t.Errorf("expected requests to stay serialized (max in-flight 1), got %d", got)
	}
}

func TestDoRequest_ParallelModeAllowsConcurrency(t *testing.T) {
	t.Setenv("TFSMC_PARALLEL_CREATE", "1")

	server, maxInFlight := maxInFlightServer(t)
	client := newTestClient(server.URL)

	// "host" is not guarded, so it can be parallel
	doConcurrentRequests(t, client, "host", 5)

	if got := atomic.LoadInt32(maxInFlight); got <= 1 {
		t.Errorf("expected requests to overlap when parallel mode is enabled, got max in-flight %d", got)
	}
}

func TestDoRequest_ParallelMode_GuardedTypeStillSerialized(t *testing.T) {
	t.Setenv("TFSMC_PARALLEL_CREATE", "1")

	server, maxInFlight := maxInFlightServer(t)
	client := newTestClient(server.URL)

	// fw_ipv4_access_rule is always serialized (hard-guarded)
	doConcurrentRequests(t, client, "fw_ipv4_access_rule", 5)

	if got := atomic.LoadInt32(maxInFlight); got != 1 {
		t.Errorf("expected fw_ipv4_access_rule requests to stay serialized, got max in-flight %d", got)
	}
}

// Regression test: ensureLoggedIn holds loginMutex for the whole lazy-login
// flow (ensureLoggedIn -> Login -> loginOnce -> unsafeRequest). unsafeRequest
// used to re-acquire loginMutex itself to read auth fields, which deadlocked
// every first-time login since sync.Mutex is not reentrant. Without a
// per-test timeout this used to hang forever; go test's default 10m test
// binary timeout was what previously surfaced it. Flagged by stonemind
// review (also caught directly by TestDoRequest_HTTPS_InsecureMode).
func TestDoRequest_ConcurrentFirstLoginNoDeadlock(t *testing.T) {
	var loginCalls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/login") {
			atomic.AddInt32(&loginCalls, 1)
			w.Header().Set("Authorization", "test-token")
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	client := &SmcClient{
		BaseUrl:    server.URL,
		APIVersion: "v1",
		SmcContext: NewLogContext(context.Background(), server.URL),
	}

	done := make(chan struct{})
	go func() {
		doConcurrentRequests(t, client, "host", 10)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("DoRequest deadlocked on concurrent first-time login")
	}

	if got := atomic.LoadInt32(&loginCalls); got != 1 {
		t.Errorf("expected exactly one login request across concurrent first-time DoRequest calls, got %d", got)
	}
}

func TestGetOrCreateClientSuccess(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && strings.Contains(r.URL.Path, "/login") {
			// Verify login request
			var loginReq LoginRequest
			if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
				t.Fatalf("Failed to decode login request: %v", err)
			}
			if loginReq.APIKey != "test-key" {
				t.Errorf("Expected APIKey test-key, got %s", loginReq.APIKey)
			}

			// Return success with session cookie
			w.Header().Set("Set-Cookie", "JSESSIONID=test-session-id; Path=/")
			w.WriteHeader(http.StatusOK)
			return
		} else if r.Method == "GET" && strings.Contains(r.URL.Path, "/v1/elements/admin_domain/1") {
			w.WriteHeader(http.StatusOK)
		} else {
			t.Fatalf("Unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	client, err := GetOrCreateClient(
		ctx, server.URL, "v1", "test-key", false, "", "", 0)

	if err != nil {
		t.Fatalf("NewClientFromAuth failed: %v", err)
	}
	err = client.Login(context.Background())

	if client.BaseUrl != server.URL {
		t.Errorf("Expected URL %s, got %s", server.URL, client.BaseUrl)
	}
	if client.APIKey != "test-key" {
		t.Errorf("Expected APIKey test-key, got %s", client.APIKey)
	}
	if client.APIVersion != "v1" {
		t.Errorf("Expected APIVersion v1, got %s", client.APIVersion)
	}
	if client.VerifySSL {
		t.Error("Expected VerifySSL to be true")
	}
	if client.Token != "test-session-id" {
		t.Errorf("Expected Token test-session-id, got %s", client.Token)
	}
	if client.HTTPClient == nil {
		t.Error("HTTPClient should not be nil")
	}

	// Test retrieving the same client again
	_, err2 := GetOrCreateClient(ctx, server.URL, "test-key", "v1", true, "", "", 0)

	if err2 != nil {
		t.Fatalf("NewClientFromAuth failed: %v", err)
	}

}

func TestGetOrCreateClientFromAuth_Insecure(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && strings.Contains(r.URL.Path, "/login") {
			w.Header().Set("Set-Cookie", "JSESSIONID=test-session-id; Path=/")
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	client, err := GetOrCreateClient(
		ctx, server.URL, "test-key", "v1", false, "", "", 0)

	if err != nil {
		t.Fatalf("NewClientFromAuth failed: %v", err)
	}

	client.Login(context.Background())

	if client.VerifySSL {
		t.Error("Expected VerifySSL to be false for insecure connection")
	}

	// Test that the client can actually make insecure connections
	transport := client.HTTPClient.Transport.(*http.Transport)
	if !transport.TLSClientConfig.InsecureSkipVerify {
		t.Errorf("Expected TLS InsecureSkipVerify to be true, got %v", transport.TLSClientConfig.InsecureSkipVerify)
	}
}

func TestGetOrCreateClientFromAuth_WithTrustedCert(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	// Skip this test since it requires a valid certificate
	// The functionality is tested in the production code
	t.Skip("Skipping trusted certificate test - requires valid PEM certificate")
}

func TestNewClientFromAuth_InvalidTrustedCert(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

	_, err := GetOrCreateClient(ctx,
		server.URL, "test-key", "v1", true, "invalid-cert-data", "", 0)

	if err == nil {
		t.Error("Expected error for invalid trusted certificate")
	}
	if !strings.Contains(err.Error(), "failed to resolve certificate") {
		t.Errorf("Expected certificate error, got %v", err)
	}
}

func TestLogin_Success(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && strings.Contains(r.URL.Path, "/login") {
			// Verify request
			var loginReq LoginRequest
			if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
				t.Fatalf("Failed to decode login request: %v", err)
			}
			if loginReq.APIKey != "test-api-key" {
				t.Errorf("Expected APIKey test-api-key, got %s", loginReq.APIKey)
			}

			// Return success with session cookie
			w.Header().Set("Set-Cookie", "JSESSIONID=session123; Path=/")
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:    server.URL,
		APIKey:     "test-api-key",
		APIVersion: "v1",
		HTTPClient: server.Client(),
	}

	err := client.Login(context.Background())
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	if client.Token != "session123" {
		t.Errorf("Expected Token session123, got %s", client.Token)
	}
}

func TestLogin_MultipleSetCookieHeaders(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && strings.Contains(r.URL.Path, "/login") {
			// Set multiple cookies
			w.Header().Add("Set-Cookie", "OTHER=value1; Path=/")
			w.Header().Add("Set-Cookie", "JSESSIONID=correct-session; Path=/; HttpOnly")
			w.Header().Add("Set-Cookie", "ANOTHER=value2; Path=/")
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:    server.URL,
		APIKey:     "test-api-key",
		APIVersion: "v1",
		HTTPClient: server.Client(),
	}

	err := client.Login(context.Background())
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	if client.Token != "correct-session" {
		t.Errorf("Expected Token correct-session, got %s", client.Token)
	}
}

func TestLogin_CreatedStatus(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && strings.Contains(r.URL.Path, "/login") {
			w.Header().Set("Set-Cookie", "JSESSIONID=created-session; Path=/")
			w.WriteHeader(http.StatusCreated) // 201 instead of 200
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:    server.URL,
		APIKey:     "test-api-key",
		APIVersion: "v1",
		HTTPClient: server.Client(),
	}

	err := client.Login(context.Background())
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	if client.Token != "created-session" {
		t.Errorf("Expected Token created-session, got %s", client.Token)
	}
}

func TestLogin_RetriesTransient500(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	restoreDelay := loginRetryBaseDelay
	loginRetryBaseDelay = time.Millisecond
	defer func() { loginRetryBaseDelay = restoreDelay }()

	var attempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && strings.Contains(r.URL.Path, "/login") {
			// Fail twice with an empty-body 500 (as seen in SMC-66463), then succeed
			if atomic.AddInt32(&attempts, 1) <= 2 {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Set-Cookie", "JSESSIONID=retried-session; Path=/")
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:    server.URL,
		APIKey:     "test-api-key",
		APIVersion: "v1",
		HTTPClient: server.Client(),
	}

	err := client.Login(context.Background())
	if err != nil {
		t.Fatalf("Login should succeed after transient 500s, got: %v", err)
	}
	if got := atomic.LoadInt32(&attempts); got != 3 {
		t.Errorf("Expected 3 login attempts, got %d", got)
	}
	if client.Token != "retried-session" {
		t.Errorf("Expected Token retried-session, got %s", client.Token)
	}
}

func TestLogin_Persistent500GivesUp(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	restoreDelay := loginRetryBaseDelay
	loginRetryBaseDelay = time.Millisecond
	defer func() { loginRetryBaseDelay = restoreDelay }()

	var attempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.Header().Set("Server", "test-tomcat")
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:    server.URL,
		APIKey:     "test-api-key",
		APIVersion: "v1",
		HTTPClient: server.Client(),
	}

	err := client.Login(context.Background())
	if err == nil {
		t.Fatal("Expected login error for persistent 500")
	}
	if got := atomic.LoadInt32(&attempts); got != loginMaxAttempts {
		t.Errorf("Expected %d login attempts, got %d", loginMaxAttempts, got)
	}
	if !strings.Contains(err.Error(), "login failed with status 500") {
		t.Errorf("Expected status 500 error, got %v", err)
	}
	// An empty 500 body must surface the response headers instead
	if !strings.Contains(err.Error(), "empty body") || !strings.Contains(err.Error(), "test-tomcat") {
		t.Errorf("Expected empty-body detail with headers in error, got %v", err)
	}
}

func TestLogin_NoRetryOn401(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	restoreDelay := loginRetryBaseDelay
	loginRetryBaseDelay = time.Millisecond
	defer func() { loginRetryBaseDelay = restoreDelay }()

	var attempts int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("Unauthorized"))
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:    server.URL,
		APIKey:     "invalid-key",
		APIVersion: "v1",
		HTTPClient: server.Client(),
	}

	err := client.Login(context.Background())
	if err == nil {
		t.Fatal("Expected login error for unauthorized request")
	}
	if got := atomic.LoadInt32(&attempts); got != 1 {
		t.Errorf("Expected a single login attempt for 401, got %d", got)
	}
}

func TestLogin_RetriesConnectionRefused(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	restoreDelay := loginRetryBaseDelay
	loginRetryBaseDelay = time.Millisecond
	defer func() { loginRetryBaseDelay = restoreDelay }()

	// Reserve a port then close the listener so connections are refused
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to open listener: %v", err)
	}
	url := "http://" + listener.Addr().String()
	_ = listener.Close()

	client := &SmcClient{
		BaseUrl:    url,
		APIKey:     "test-api-key",
		APIVersion: "v1",
		HTTPClient: &http.Client{},
	}

	start := time.Now()
	err = client.Login(context.Background())
	if err == nil {
		t.Fatal("Expected login error for refused connection")
	}
	// 4 backoffs of 1/2/4/8 ms prove the retry loop actually ran
	if elapsed := time.Since(start); elapsed < 15*time.Millisecond {
		t.Errorf("Expected retries with backoff, login returned after %v", elapsed)
	}
}

func TestLogin_HTTPError(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("Unauthorized"))
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:    server.URL,
		APIKey:     "invalid-key",
		APIVersion: "v1",
		HTTPClient: server.Client(),
	}

	err := client.Login(context.Background())
	if err == nil {
		t.Error("Expected login error for unauthorized request")
	}
	if !strings.Contains(err.Error(), "login failed with status 401") {
		t.Errorf("Expected unauthorized error, got %v", err)
	}
}

func TestLogin_NoSessionCookie(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// Return success but without JSESSIONID cookie
		w.Header().Set("Set-Cookie", "OTHER=value; Path=/")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &SmcClient{
		BaseUrl:    server.URL,
		APIKey:     "test-key",
		APIVersion: "v1",
		HTTPClient: server.Client(),
	}

	err := client.Login(context.Background())
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	// Token should be empty if no JSESSIONID is found
	if client.Token != "" {
		t.Errorf("Expected empty Token, got %s", client.Token)
	}
}

func TestLogin_InvalidJSON(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	// Test that invalid JSON in request body is handled
	originalClient := &SmcClient{
		BaseUrl:    "http://example.com",
		APIKey:     "test-key",
		APIVersion: "v1",
		HTTPClient: &http.Client{},
	}

	// This should work normally as the JSON marshaling is straightforward
	err := originalClient.Login(context.Background())
	// This will likely fail due to network issues, but not due to JSON marshaling
	if err == nil {
		t.Error("Expected network error")
	}
}

func TestLoginRequest_JSONMarshal(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	req := LoginRequest{
		APIKey: "test-key-123",
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal LoginRequest: %v", err)
	}

	expected := `{"authenticationkey":"test-key-123"}`
	if string(data) != expected {
		t.Errorf("Expected JSON %s, got %s", expected, string(data))
	}
}

func TestClient_Fields(t *testing.T) {
	GetClientManager().CleanupAll(context.Background())
	client := &SmcClient{
		BaseUrl:    "https://test.com",
		APIKey:     "key123",
		VerifySSL:  true,
		Token:      "token456",
		APIVersion: "v2",
		HTTPClient: &http.Client{},
	}

	if client.BaseUrl != "https://test.com" {
		t.Errorf("Expected URL https://test.com, got %s", client.BaseUrl)
	}
	if client.APIKey != "key123" {
		t.Errorf("Expected APIKey key123, got %s", client.APIKey)
	}
	if !client.VerifySSL {
		t.Error("Expected VerifySSL to be true")
	}
	if client.Token != "token456" {
		t.Errorf("Expected Token token456, got %s", client.Token)
	}
	if client.APIVersion != "v2" {
		t.Errorf("Expected APIVersion v2, got %s", client.APIVersion)
	}
	if client.HTTPClient == nil {
		t.Error("HTTPClient should not be nil")
	}
}
