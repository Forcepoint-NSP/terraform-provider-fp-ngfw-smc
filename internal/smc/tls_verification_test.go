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

// Tests for the provider's TLS verification path (SMC-64735): the SmcClient
// must be able to establish and *verify* an SSL connection to an HTTPS endpoint
// using a supplied trusted_cert (verify_ssl=true), enforce that verification (a
// wrong cert is rejected), allow skipping it (verify_ssl=false), and still
// reach the endpoint when a stored href carries a stale scheme/host/version
// that NormalizeHref must rewrite. The switch_api_scheme robot suite exercises
// the same behavior end-to-end against a real SMC.

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// unrelatedCertPEM generates a standalone self-signed certificate, unrelated to
// any test server, to act as a "wrong" trusted_cert.
func unrelatedCertPEM(t *testing.T) string {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "unrelated"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}))
}

// certPEM returns the PEM-encoded leaf certificate presented by an httptest
// TLS server (acts as the "trusted_cert" the operator would pass to the
// provider).
func certPEM(t *testing.T, ts *httptest.Server) string {
	t.Helper()
	cert := ts.Certificate()
	block := &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}
	return string(pem.EncodeToMemory(block))
}

// newTLSSMC spins an in-process HTTPS server that answers a GET on any
// /<ver>/elements/... path with a small JSON body, standing in for the SMC API
// behind TLS.
func newTLSSMC(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("ETag", "\"test-etag\"")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"name":"tlstest","link":[{"rel":"self","href":"` + r.URL.String() + `"}]}`))
	}))
}

func newTLSTestClient(baseURL, trustedCert string, verifySSL bool) *SmcClient {
	c, err := NewClient(context.Background(), baseURL, "7.6", "test-key", verifySSL, trustedCert, "")
	if err != nil {
		// NewClient only errors on an invalid/unresolvable cert; callers that
		// expect that should construct the client directly.
		return nil
	}
	// Skip the login round-trip; we only want to exercise the TLS path here.
	c.LoginConfirmed = true
	c.UseAuthHeader = true
	c.Token = "test-token"
	return c
}

// 1. Happy path: verify_ssl=true with the server's cert as inline PEM trusted_cert.
func TestTLSTrustedCertVerifies(t *testing.T) {
	ts := newTLSSMC(t)
	defer ts.Close()

	c := newTLSTestClient(ts.URL, certPEM(t, ts), true)
	cfg := &GenericCRUDConfig{ResourceType: "host", Client: c}

	resp, err := cfg.ReadResourceByHref(ts.URL + "/7.6/elements/host/1")
	if err != nil {
		t.Fatalf("expected verified TLS read to succeed, got error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	t.Logf("OK: provider verified the SSL connection using the supplied trusted_cert")
}

// 2. Enforcement: verify_ssl=true with a *wrong* cert must be rejected.
func TestTLSWrongCertRejected(t *testing.T) {
	ts := newTLSSMC(t)
	defer ts.Close()

	// A syntactically valid but unrelated self-signed cert.
	c := newTLSTestClient(ts.URL, unrelatedCertPEM(t), true)
	cfg := &GenericCRUDConfig{ResourceType: "host", Client: c}

	_, err := cfg.ReadResourceByHref(ts.URL + "/7.6/elements/host/1")
	if err == nil {
		t.Fatalf("expected TLS verification to FAIL with a wrong trusted_cert, but it succeeded")
	}
	t.Logf("OK: verification enforced, wrong cert rejected: %v", err)
}

// 3. verify_ssl=false skips verification (the demo-style insecure path).
func TestTLSInsecureSkipsVerification(t *testing.T) {
	ts := newTLSSMC(t)
	defer ts.Close()

	c := newTLSTestClient(ts.URL, "", false)
	cfg := &GenericCRUDConfig{ResourceType: "host", Client: c}

	if _, err := cfg.ReadResourceByHref(ts.URL + "/7.6/elements/host/1"); err != nil {
		t.Fatalf("expected insecure read to succeed, got: %v", err)
	}
	t.Logf("OK: verify_ssl=false connects without a trusted_cert")
}

// 4. The actual SMC-64735 scenario: a stored href with the OLD http scheme
// (and stale host/port) is normalized to the current https base + verified.
func TestTLSStaleHTTPHrefNormalizedToHTTPS(t *testing.T) {
	ts := newTLSSMC(t)
	defer ts.Close()

	c := newTLSTestClient(ts.URL, certPEM(t, ts), true)
	cfg := &GenericCRUDConfig{ResourceType: "host", Client: c}

	// Simulate an id stored in tfstate before the http->https switch:
	// wrong scheme, wrong host/port, even a stale version.
	staleHref := "http://old-host:8082/7.4/elements/host/1"

	resp, err := cfg.ReadResourceByHref(staleHref)
	if err != nil {
		t.Fatalf("expected stale http href to be normalized to verified https and succeed, got: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if !strings.Contains(string(resp.Body), "/7.6/elements/host/1") {
		t.Fatalf("expected request to be rewritten to current base+version, body=%s", string(resp.Body))
	}
	t.Logf("OK: stale http href normalized to verified https + current version")
}
