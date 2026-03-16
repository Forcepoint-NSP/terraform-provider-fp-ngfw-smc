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

// Package smc provides the core functionality for the SMC Terraform provider.
package smc

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	SMC_DEFAULT_REQUEST_TIMEOUT_SEC = 300
)

const PingEntryPoint = "elements/admin_domain/1"

// SmcClient represents an SMC API client with authentication capabilities.
// URL will be used as a unique identifier for the provider (SMC Server ID)
type SmcClient struct {
	BaseUrl       string // Base URL of the SMC API. expecting: "http(s)://<host>[:port]"
	APIKey        string
	VerifySSL     bool   // Whether to verify SSL certificates
	TrustedCert   string // PEM encoded trusted certificate
	Token         string // Authentication token (JSESSIONID or Authorization header)
	APIVersion    string // API version to use. eg "7.4"
	Domain        string
	UseAuthHeader bool // Track whether to use Authorization header
	// or Cookie. This is set during login.
	LoginConfirmed bool        // Track if login was successful.
	SmcContext     *LogContext // Context for structured logging
	OriginalCtx    *context.Context

	BeforeRequestHooks []func(client *SmcClient, opts *Options) error // Hooks to modify request options before sending

	// todo currently not used: a client is created per request
	HTTPClient *http.Client
}

// Options define options for the generic request
type Options struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []byte
	Timeout time.Duration
	Context context.Context
}

// ResponseData contains only the body, headers, and status code
// for easier consumption in callers
type ResponseData struct {
	Body       []byte
	Headers    http.Header
	StatusCode int
	ETag       string
	Location   string
}

// LoginRequest represents the payload for SMC API login
type LoginRequest struct {
	APIKey string `json:"authenticationkey"`
	Domain string `json:"domain,omitempty"`
}

// ValidateBaseUrl checks if the provided URL is valid for SMC API usage
// valid url like: http(s)://<host>[:port]
func ValidateBaseUrl(urlStr string) (bool, error) {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return false, err
	}

	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return false, fmt.Errorf("invalid URL scheme: %s", parsedUrl.Scheme)
	}

	if parsedUrl.Path != "" && parsedUrl.Path != "/" {
		return false, fmt.Errorf("base URL should not contain a path: %s", parsedUrl.Path)
	}
	isHttps := parsedUrl.Scheme == "https"
	return isHttps, nil
}

// createTLSConfig creates a TLS configuration with trusted certificate
func createTLSConfig(trustedCert string, insecure bool) (*tls.Config, error) {
	tlsConfig := &tls.Config{}

	// If insecure is explicitly set to true, skip certificate verification
	if insecure {
		tlsConfig.InsecureSkipVerify = true
		return tlsConfig, nil
	}

	// Create certificate pool and add the trusted certificate content
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM([]byte(trustedCert)) {
		return nil, fmt.Errorf("failed to append trusted certificate")
	}

	tlsConfig.RootCAs = caCertPool
	tlsConfig.InsecureSkipVerify = false

	return tlsConfig, nil
}

// unsafeRequest performs the actual HTTP request without
// locking. Should not be called directly, use DoRequest instead.
func (c *SmcClient) unsafeRequest(opts *Options) (*ResponseData, error) {
	ctx := opts.Context
	if ctx == nil {
		ctx = context.Background()
	}

	var body io.Reader
	if opts.Body != nil {
		body = bytes.NewReader(opts.Body)
	}

	// replace the scheme, host and port with those from the
	// client (reverse proxy)
	newUrl, err := ReplaceBaseInURL(opts.URL, c.BaseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to replace base URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, opts.Method, newUrl, body)
	if err != nil {
		return nil, err
	}

	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	if (opts.Body != nil ||
		strings.EqualFold(opts.Method, "put") ||
		strings.EqualFold(opts.Method, "post")) && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.UseAuthHeader {
		req.Header.Set("Authorization", c.Token)
	} else {
		req.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%s", c.Token))
	}

	// Create HTTP client with optional TLS configuration
	client := &http.Client{}
	if opts.Timeout > 0 {
		client.Timeout = opts.Timeout
	} else {
		client.Timeout = SMC_DEFAULT_REQUEST_TIMEOUT_SEC * time.Second
	}

	if req.URL.Scheme == "https" {
		tlsConfig, err := createTLSConfig(c.TrustedCert, !c.VerifySSL)
		if err != nil {
			return nil, fmt.Errorf("failed to create TLS config: %w", err)
		}

		// todo reuse client.Transport if already set?
		transport := &http.Transport{
			TLSClientConfig: tlsConfig,
		}
		client.Transport = transport
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	etag := resp.Header.Get("ETag")
	if etag != "" {
		etag = strings.ReplaceAll(etag, "\"", "")
	} else {
		etag = ""
	}

	location := resp.Header.Get("Location")

	return &ResponseData{
		Body:       respBody,
		Headers:    resp.Header,
		StatusCode: resp.StatusCode,
		ETag:       etag,
		Location:   location,
	}, nil
}

// ping checks connectivity to the SMC API by sending a GET request to
// the login endpoint and returns true if successful.
func (c *SmcClient) ping(ctx context.Context) bool {
	url := fmt.Sprintf("%s/%s/%s", c.BaseUrl, c.APIVersion, PingEntryPoint)
	opts := Options{
		Method:  "GET",
		URL:     url,
		Headers: map[string]string{"Content-Type": "application/json"},
		Context: ctx,
	}
	_, err := c.DoRequest(opts)
	if err != nil {
		return false
	}
	return true
}

// GetJSONHeaders returns standard headers for JSON API requests with
// authentication
func (c *SmcClient) GetJSONHeaders() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if c.UseAuthHeader {
		headers["Authorization"] = c.Token
	} else {
		headers["Cookie"] = fmt.Sprintf("JSESSIONID=%s", c.Token)
	}

	return headers
}

// GetAuthHeaders returns headers for authenticated requests (no Content-Type)
func (c *SmcClient) GetAuthHeaders() map[string]string {
	if c.UseAuthHeader {
		return map[string]string{
			"Authorization": c.Token,
		}
	}
	return map[string]string{
		"Cookie": fmt.Sprintf("JSESSIONID=%s", c.Token),
	}
}

// GetJSONHeadersWithEtag returns JSON headers with If-Match header
// for conditional updates
func (c *SmcClient) GetJSONHeadersWithEtag(etag string) map[string]string {
	headers := c.GetJSONHeaders()
	headers["If-Match"] = etag
	return headers
}

// GetAuthHeadersWithEtag returns auth headers with If-Match header
// for conditional operations
func (c *SmcClient) GetAuthHeadersWithEtag(etag string) map[string]string {
	headers := c.GetAuthHeaders()
	headers["If-Match"] = etag
	return headers
}

// GetLogContext returns the client's log context
func (c *SmcClient) GetLogContext() *LogContext {
	return c.SmcContext
}

// UpdateLogContext updates the client's log context with a new MgtServerId
func (c *SmcClient) UpdateLogContext(ctx context.Context, mgtServerId string) {
	c.SmcContext = NewLogContext(ctx, mgtServerId)
}

// ResolveCertificate resolves the certificate from either a file path or PEM-encoded content.
// If the input looks like a file path (exists as a file), it reads the file.
// Otherwise, it assumes the input is PEM-encoded certificate content.
func ResolveCertificate(certInput string) (string, error) {
	// First, try to check if it's a valid file path
	if _, err := os.Stat(certInput); err == nil {
		// File exists, read it
		certContent, err := os.ReadFile(certInput)
		if err != nil {
			return "", fmt.Errorf("failed to read certificate file: %w", err)
		}
		return string(certContent), nil
	}

	// If file doesn't exist, assume it's PEM content directly
	// Check if it looks like PEM content (starts with -----BEGIN)
	trimmedInput := strings.TrimSpace(certInput)
	if strings.HasPrefix(trimmedInput, "-----BEGIN") {
		return trimmedInput, nil
	}

	// If it's neither a file nor PEM content, return an error
	return "", fmt.Errorf("certificate input is neither a valid file path nor PEM-encoded content")
}

// NewClient creates a new SmcClient
// use GetOrCreateClientFromAuth if you need to reuse sessions
func NewClient(ctx context.Context,
	url string, apiVersion string, apikey string, verifySSL bool,
	trustedCert string, domain string) (*SmcClient, error) {

	// Create the initial log context.
	logCtx := NewLogContext(ctx, url)

	logCtx.Info(fmt.Sprintf("New SMC client session on %s", url), map[string]any{
		"url":         url,
		"api_version": apiVersion,
	})

	tlsConfig := &tls.Config{InsecureSkipVerify: !verifySSL}

	// todo the client/transport configured here is not used
	if trustedCert != "" {
		// Resolve the certificate - either from file path or direct content
		certContent, err := ResolveCertificate(trustedCert)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve certificate: %w", err)
		}

		// Load custom trusted cert
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM([]byte(certContent)) {
			return nil, fmt.Errorf("failed to append trusted certificate")
		}
		tlsConfig.RootCAs = certPool
	}
	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	httpClient := &http.Client{Transport: tr}

	c := &SmcClient{
		BaseUrl:     url,
		APIKey:      apikey,
		VerifySSL:   verifySSL,
		HTTPClient:  httpClient, // todo not used
		APIVersion:  apiVersion,
		TrustedCert: trustedCert,
		Domain:      domain,
		SmcContext:  logCtx,
		OriginalCtx: &ctx,
	}
	return c, nil
}

// Login Login authenticates with the SMC API using the provided API key
func (c *SmcClient) Login(ctx context.Context) error {
	if c.LoginConfirmed {
		c.SmcContext.Debug("Already logged to SMC API", map[string]any{
			"url": c.BaseUrl,
		})
		return nil
	}

	c.SmcContext.Info("Attempting to login to SMC API", map[string]any{
		"url": c.BaseUrl,
	})

	loginRequest := LoginRequest{
		APIKey: c.APIKey,
		Domain: c.Domain,
	}

	reqBody, err := json.Marshal(loginRequest)
	if err != nil {
		c.SmcContext.Error("Failed to marshal login request", map[string]any{
			"error": err.Error(),
		})
		return err
	}

	url := fmt.Sprintf("%s/%s/login", c.BaseUrl, c.APIVersion)
	opts := Options{
		Method:  "POST",
		URL:     url,
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    reqBody,
		Context: ctx,
	}
	resp, err := c.unsafeRequest(&opts)
	if err != nil {
		c.SmcContext.Error("Login request failed", map[string]any{
			"error": err.Error(),
		})
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		c.SmcContext.Error("Login failed with non-success status", map[string]any{
			"status_code": resp.StatusCode,
			"body":        string(resp.Body),
		})
		return fmt.Errorf("login failed with status %d: %s", resp.StatusCode, string(resp.Body))
	}

	// Try to get token from Authorization header first (for HTTPS)
	if authHeader := resp.Headers.Get("Authorization"); authHeader != "" {
		c.Token = authHeader
		c.UseAuthHeader = true
		c.SmcContext.Info("Successfully logged in using Authorization header")
		return nil
	}

	// Fallback: Extract JSESSIONID from Set-Cookie header (for HTTP)
	cookies := resp.Headers["Set-Cookie"]
	for _, cookieStr := range cookies {
		if strings.Contains(cookieStr, "JSESSIONID=") {
			parts := strings.Split(cookieStr, ";")
			for _, part := range parts {
				if strings.HasPrefix(part, "JSESSIONID=") {
					c.Token = strings.TrimPrefix(part, "JSESSIONID=")
					c.UseAuthHeader = false
					c.SmcContext.Info("Successfully logged in using JSESSIONID cookie")
					break
				}
			}
		}
	}
	c.LoginConfirmed = true
	return nil
}

// Logout closes the SMC API session and cleans up resources
func (c *SmcClient) Logout(ctx context.Context) error {
	if !c.LoginConfirmed {
		c.SmcContext.Debug("No active login session to logout", map[string]any{
			"url": c.BaseUrl,
		})
		return nil
	}

	c.SmcContext.Info("Logging out from SMC API", map[string]any{
		"url": c.BaseUrl,
	})

	url := fmt.Sprintf("%s/%s/logout", c.BaseUrl, c.APIVersion)
	opts := Options{
		Method:  "PUT",
		URL:     url,
		Headers: c.GetAuthHeaders(),
		Context: ctx,
	}

	resp, err := c.DoRequest(opts)
	if err != nil {
		c.SmcContext.Warn("Logout request failed", map[string]any{
			"error": err.Error(),
		})
		// Don't return error - we still want to mark as logged out
	} else if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		c.SmcContext.Warn("Logout failed with non-success status", map[string]any{
			"status_code": resp.StatusCode,
			"body":        string(resp.Body),
		})
		// Don't return error - we still want to mark as logged out
	} else {
		c.SmcContext.Info("Successfully logged out from SMC API", map[string]any{
			"url": c.BaseUrl,
		})
	}

	// Clear the authentication state
	c.Token = ""
	c.LoginConfirmed = false

	return nil
}

// DoRequest performs a generic HTTP request and returns ResponseData
// or error. In case of HTTP errors (4xx, 5xx), no error is returned,
// the caller must check the StatusCode in ResponseData.
func (c *SmcClient) DoRequest(opts Options) (*ResponseData, error) {
	// Lock by base URL to prevent concurrent requests
	// Using different lock by base URL will allow
	// processing if several SMC providers are set.
	lockId := GetLockId(opts.URL)
	GetLockRequests().Enter(lockId)
	defer GetLockRequests().Leave(lockId)

	if !c.LoginConfirmed {
		err := c.Login(opts.Context)
		if err != nil {
			return nil, fmt.Errorf("failed to login before request: %w", err)
		}
	}

	for _, hook := range c.BeforeRequestHooks {
		err := hook(c, &opts)
		if err != nil {
			return nil, err
		}
	}

	return c.unsafeRequest(&opts)
}
