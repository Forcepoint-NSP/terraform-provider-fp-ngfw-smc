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
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func DoRequest(opts Options) (*ResponseData, error) {
	c := SmcClient{
		BaseUrl:        opts.URL,
		LoginConfirmed: true,
	}
	return c.DoRequest(opts)
}

func TestDoRequest_GET(t *testing.T) {
	expectedBody := `{"message": "hello"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("ETag", `"test-etag"`)
		w.Header().Set("Location", "/test/location")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, expectedBody)
	}))
	defer server.Close()

	opts := Options{
		Method: "GET",
		URL:    server.URL,
		Headers: map[string]string{
			"Accept": "application/json",
		},
	}

	resp, err := DoRequest(opts)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	if string(resp.Body) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, string(resp.Body))
	}

	if resp.ETag != "test-etag" {
		t.Errorf("Expected ETag test-etag, got %s", resp.ETag)
	}

	if resp.Location != "/test/location" {
		t.Errorf("Expected Location /test/location, got %s", resp.Location)
	}

	contentType := resp.Headers.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}
}

func TestDoRequest_POST_WithBody(t *testing.T) {
	expectedRequestBody := `{"name": "test"}`
	var receivedBody string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		body := make([]byte, r.ContentLength)
		_, _ = r.Body.Read(body)
		receivedBody = string(body)

		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, `{"id": 123}`)
	}))
	defer server.Close()

	opts := Options{
		Method: "POST",
		URL:    server.URL,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: []byte(expectedRequestBody),
	}

	resp, err := DoRequest(opts)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

	if receivedBody != expectedRequestBody {
		t.Errorf("Expected request body %s, got %s", expectedRequestBody, receivedBody)
	}

	if string(resp.Body) != `{"id": 123}` {
		t.Errorf("Expected response body {\"id\": 123}, got %s", string(resp.Body))
	}
}

func TestDoRequest_PUT(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	opts := Options{
		Method: "PUT",
		URL:    server.URL,
		Body:   []byte(`{"update": "data"}`),
	}

	resp, err := DoRequest(opts)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestDoRequest_DELETE(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	opts := Options{
		Method: "DELETE",
		URL:    server.URL,
	}

	resp, err := DoRequest(opts)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status 204, got %d", resp.StatusCode)
	}
}

func TestDoRequest_WithHeaders(t *testing.T) {
	expectedHeaders := map[string]string{
		"Authorization": "Bearer token123",
		"User-Agent":    "test-client/1.0",
		"X-Custom":      "custom-value",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, expectedValue := range expectedHeaders {
			actualValue := r.Header.Get(key)
			if actualValue != expectedValue {
				t.Errorf("Expected header %s: %s, got %s", key, expectedValue, actualValue)
			}
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	opts := Options{
		Method:  "GET",
		URL:     server.URL,
		Headers: expectedHeaders,
	}

	_, err := DoRequest(opts)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}
}

func TestDoRequest_WithTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(200 * time.Millisecond) // Delay longer than timeout
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	opts := Options{
		Method:  "GET",
		URL:     server.URL,
		Timeout: 100 * time.Millisecond,
	}

	_, err := DoRequest(opts)
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
}

func TestDoRequest_WithContext(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	opts := Options{
		Method:  "GET",
		URL:     server.URL,
		Context: ctx,
	}

	_, err := DoRequest(opts)
	if err == nil {
		t.Error("Expected context timeout error, got nil")
	}
}

func TestDoRequest_NoBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	opts := Options{
		Method: "GET",
		URL:    server.URL,
		Body:   nil, // Explicitly nil
	}

	resp, err := DoRequest(opts)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestDoRequest_ETagProcessing(t *testing.T) {
	tests := []struct {
		name         string
		serverETag   string
		expectedETag string
	}{
		{
			name:         "ETag with quotes",
			serverETag:   `"quoted-etag"`,
			expectedETag: "quoted-etag",
		},
		{
			name:         "ETag without quotes",
			serverETag:   "unquoted-etag",
			expectedETag: "unquoted-etag",
		},
		{
			name:         "Empty ETag",
			serverETag:   "",
			expectedETag: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				if tt.serverETag != "" {
					w.Header().Set("ETag", tt.serverETag)
				}
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			opts := Options{
				Method: "GET",
				URL:    server.URL,
			}

			resp, err := DoRequest(opts)
			if err != nil {
				t.Fatalf("DoRequest failed: %v", err)
			}

			if resp.ETag != tt.expectedETag {
				t.Errorf("Expected ETag %s, got %s", tt.expectedETag, resp.ETag)
			}
		})
	}
}

func TestDoRequest_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprint(w, "Internal Server Error")
	}))
	defer server.Close()

	opts := Options{
		Method: "GET",
		URL:    server.URL,
	}

	resp, err := DoRequest(opts)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status 500, got %d", resp.StatusCode)
	}

	if string(resp.Body) != "Internal Server Error" {
		t.Errorf("Expected error message, got %s", string(resp.Body))
	}
}

func TestDoRequest_InvalidURL(t *testing.T) {
	opts := Options{
		Method: "GET",
		URL:    "://invalid-url",
	}

	_, err := DoRequest(opts)
	if err == nil {
		t.Error("Expected error for invalid URL, got nil")
	}
}

func TestDoRequest_DefaultContext(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	opts := Options{
		Method:  "GET",
		URL:     server.URL,
		Context: nil, // Test with nil context
	}

	resp, err := DoRequest(opts)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestResponseData_Fields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("ETag", `"test-etag"`)
		w.Header().Set("Location", "/test/location")
		w.Header().Set("Custom-Header", "custom-value")
		w.WriteHeader(http.StatusAccepted)
		_, _ = fmt.Fprint(w, "response body")
	}))
	defer server.Close()

	opts := Options{
		Method: "POST",
		URL:    server.URL,
	}

	resp, err := DoRequest(opts)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}

	// Test all ResponseData fields
	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected StatusCode 202, got %d", resp.StatusCode)
	}

	if string(resp.Body) != "response body" {
		t.Errorf("Expected Body 'response body', got %s", string(resp.Body))
	}

	if resp.ETag != "test-etag" {
		t.Errorf("Expected ETag 'test-etag', got %s", resp.ETag)
	}
}

// Helper function to generate a test certificate
func generateTestCert() (string, *tls.Certificate, error) {
	// Generate private key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", nil, err
	}

	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "localhost",
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.IPv4(127, 0, 0, 1)},
		DNSNames:    []string{"localhost"},
	}

	// Create certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return "", nil, err
	}

	// Create PEM encoded certificate
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	keyDER, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return "", nil, err
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyDER})

	// Create tls.Certificate
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return "", nil, err
	}

	return string(certPEM), &cert, nil
}

// TestDoRequest_HTTPS_WithTrustedCert tests HTTPS requests with trusted certificates
func TestDoRequest_HTTPS_WithTrustedCert(t *testing.T) {
	certPEM, cert, err := generateTestCert()
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	// Create HTTPS test server with custom certificate
	server := httptest.NewUnstartedServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected GET request, got %s", r.Method)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"status": "success"}`)
		}))

	// Configure server with custom certificate
	server.TLS = &tls.Config{
		Certificates: []tls.Certificate{*cert},
	}
	server.StartTLS()
	defer server.Close()

	opts := Options{
		Method: "GET",
		URL:    server.URL,
		Headers: map[string]string{
			"Accept": "application/json",
		},
	}

	c := SmcClient{
		BaseUrl:        opts.URL,
		TrustedCert:    certPEM,
		LoginConfirmed: true,
	}

	resp, err := c.DoRequest(opts)
	if err != nil {
		t.Fatalf("DoRequest with trusted cert failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	expectedBody := `{"status": "success"}`
	if string(resp.Body) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, string(resp.Body))
	}
}

// TestDoRequest_HTTPS_InsecureMode tests HTTPS requests with insecure mode
func TestDoRequest_HTTPS_InsecureMode(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"insecure": "test"}`)
	}))
	defer server.Close()

	opts := Options{
		Method: "GET",
		URL:    server.URL,
		Headers: map[string]string{
			"Accept": "application/json",
		},
	}

	c := SmcClient{
		BaseUrl:     opts.URL,
		VerifySSL:   false,
		TrustedCert: "",
	}

	resp, err := c.DoRequest(opts)
	if err != nil {
		t.Fatalf("DoRequest with insecure mode failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	expectedBody := `{"insecure": "test"}`
	if string(resp.Body) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, string(resp.Body))
	}
}

// TestDoRequest_HTTPS_InvalidCertificate tests HTTPS requests with invalid certificates
func TestDoRequest_HTTPS_InvalidCertificate(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Test with invalid certificate
	opts := Options{
		Method: "GET",
		URL:    server.URL,
	}

	c := SmcClient{
		BaseUrl:     opts.URL,
		TrustedCert: "invalid-certificate-data",
		VerifySSL:   true,
	}

	_, err := c.DoRequest(opts)
	if err == nil {
		t.Error("Expected error for invalid certificate, got nil")
	}

	if !strings.Contains(err.Error(), "failed to append trusted certificate") {
		t.Errorf("Expected certificate error, got %v", err)
	}
}

// TestCreateTLSConfig tests the TLS configuration creation
func TestCreateTLSConfig(t *testing.T) {
	// Test with valid certificate
	certPEM, _, err := generateTestCert()
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	tlsConfig, err := createTLSConfig(certPEM, false)
	if err != nil {
		t.Fatalf("createTLSConfig failed: %v", err)
	}

	if tlsConfig.InsecureSkipVerify {
		t.Error("Expected InsecureSkipVerify to be false")
	}

	if tlsConfig.RootCAs == nil {
		t.Error("Expected RootCAs to be set")
	}
}

// TestCreateTLSConfig_InsecureMode tests TLS config with insecure mode
func TestCreateTLSConfig_InsecureMode(t *testing.T) {
	insecure := true
	tlsConfig, err := createTLSConfig("dummy-cert", insecure)
	if err != nil {
		t.Fatalf("createTLSConfig with insecure mode failed: %v", err)
	}

	if !tlsConfig.InsecureSkipVerify {
		t.Error("Expected InsecureSkipVerify to be true")
	}
}

// TestCreateTLSConfig_InvalidCertificate tests TLS config with invalid certificate
func TestCreateTLSConfig_InvalidCertificate(t *testing.T) {
	_, err := createTLSConfig("invalid-certificate", false)
	if err == nil {
		t.Error("Expected error for invalid certificate")
	}

	expectedError := "failed to append trusted certificate"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Expected error containing %s, got %v", expectedError, err)
	}
}
