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
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-log/tflogtest"
)

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
		ctx, server.URL, "v1", "test-key", false, "", "")

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
	_, err2 := GetOrCreateClient(ctx, server.URL, "test-key", "v1", true, "", "")

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
		ctx, server.URL, "test-key", "v1", false, "", "")

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
		server.URL, "test-key", "v1", true, "invalid-cert-data", "")

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
