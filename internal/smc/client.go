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
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	SMC_DEFAULT_REQUEST_TIMEOUT_SEC = 300
)

const PingEntryPoint = "elements/admin_domain/1"

// SmcClient represents an SMC API client with authentication capabilities.
// URL will be used as a unique identifier for the provider (SMC Server ID)
type SmcClient struct {
	BaseUrl     string // Base URL of the SMC API. expecting: "http(s)://<host>[:port]"
	APIKey      string
	VerifySSL   bool   // Whether to verify SSL certificates
	TrustedCert string // PEM encoded trusted certificate
	Token       string // Authentication token (JSESSIONID or Authorization header value)
	// SessionCookie holds the JSESSIONID returned by the SMC login response.
	// When the SMC option "Use SSL for session ID" is disabled, the server
	// tracks sessions via JSESSIONID even on HTTPS, so this cookie must be
	// sent alongside the Authorization header on every request.
	SessionCookie string
	APIVersion    string // API version to use. eg "7.4"
	Domain        string
	UseAuthHeader bool // Track whether to use Authorization header
	// or Cookie. This is set during login.
	LoginConfirmed bool        // Track if login was successful.
	SmcContext     *LogContext // Context for structured logging
	OriginalCtx    *context.Context

	// TokenRefreshInterval, when > 0, makes DoRequest proactively refresh the
	// SMC API session token once this much time has elapsed
	// since it was last issued, instead of waiting for the SMC to reject a
	// request with 401. Set it comfortably below the server's configured
	// SMC_API_TOKEN_VALIDITY so a long-running apply keeps refreshing ahead
	// of expiry and never actually hits a 401 in normal use.
	TokenRefreshInterval time.Duration
	// TokenIssuedAt is when Token was last obtained via Login or RefreshToken.
	TokenIssuedAt time.Time
	// lastRefreshAttempt is when a proactive refresh last failed; it makes the
	// next attempt wait a full TokenRefreshInterval.
	lastRefreshAttempt time.Time

	BeforeRequestHooks []func(client *SmcClient, opts *Options) error // Hooks to modify request options before sending

	// loginMutex protects concurrent access to login-related fields (Token,
	// SessionCookie, UseAuthHeader, LoginConfirmed, TokenIssuedAt). Required
	// because the client is a per-URL singleton and login is lazy (triggered
	// on first request). Without this, concurrent parallel requests can race
	// to login or refresh and create multiple sessions, invalidating tokens
	// (401 errors).
	loginMutex sync.Mutex

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
	// CanParallelize reports whether this request may bypass the global
	// per-host lock (see parallel.go). Callers that know about SMC
	// resource types provide this (e.g. GenericCRUDConfig.canParallelize);
	// nil or false means the request stays serialized.
	CanParallelize func() bool
	// SkipAuthRetry disables the automatic session-recovery retry on a 401
	// response and the proactive token refresh (see DoRequest). Used by
	// Logout, where refreshing or re-authenticating just to immediately log
	// out again would be pointless, and by ping, which must return quickly
	// and let the caller (GetClient) decide how to recover.
	SkipAuthRetry bool
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

// NormalizeHref rewrites an SMC href so it targets the currently configured
// connection: it replaces the scheme/host/port with the client BaseUrl and the
// API version path segment with the client APIVersion.
//
// SMC element ids are stable across versions, so an href stored in Terraform
// state (e.g. from a previous SMC version, or before an http->https switch)
// remains valid after being normalized. See SMC-66510 and SMC-64735.
func (c *SmcClient) NormalizeHref(href string) (string, error) {
	normalized, err := ReplaceBaseInURL(href, c.BaseUrl)
	if err != nil {
		return "", err
	}
	return ReplaceVersionInURL(normalized, c.APIVersion)
}

// authSnapshot is a point-in-time copy of the auth fields needed to send a
// request, taken under loginMutex so it can be read outside the lock.
type authSnapshot struct {
	useAuthHeader bool
	token         string
	sessionCookie string
}

// currentAuth reads the client's auth fields under loginMutex. Must not be
// called by anything already holding loginMutex (e.g. the login/refresh
// flows) since sync.Mutex is not reentrant; those callers already have
// exclusive access and should read the fields directly instead.
func (c *SmcClient) currentAuth() authSnapshot {
	c.loginMutex.Lock()
	defer c.loginMutex.Unlock()
	return authSnapshot{
		useAuthHeader: c.UseAuthHeader,
		token:         c.Token,
		sessionCookie: c.SessionCookie,
	}
}

// unsafeRequest performs the actual HTTP request without
// locking. Should not be called directly, use DoRequest instead.
func (c *SmcClient) unsafeRequest(opts *Options, auth authSnapshot) (*ResponseData, error) {
	ctx := opts.Context
	if ctx == nil {
		ctx = context.Background()
	}

	var body io.Reader
	if opts.Body != nil {
		body = bytes.NewReader(opts.Body)
	}

	// replace the scheme, host, port and API version with those from
	// the client, so a stale href stored in state still targets the
	// currently configured SMC (reverse proxy, http->https, version upgrade)
	newUrl, err := c.NormalizeHref(opts.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to normalize URL: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, opts.Method, newUrl, body)
	if err != nil {
		return nil, err
	}

	for k, v := range opts.Headers {
		// Skip auth headers — they are always set below from c.Token/c.UseAuthHeader.
		// opts.Headers may contain stale auth values captured before login (e.g. an
		// empty JSESSIONID cookie from GetJSONHeaders() on an unauthenticated client).
		// Letting them through causes the SMC to reject the request with 401 even
		// when a valid Authorization header is also present.
		if strings.EqualFold(k, "Authorization") || strings.EqualFold(k, "Cookie") {
			continue
		}
		req.Header.Set(k, v)
	}

	if (opts.Body != nil ||
		strings.EqualFold(opts.Method, "put") ||
		strings.EqualFold(opts.Method, "post")) && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	useAuthHeader := auth.useAuthHeader
	token := auth.token
	sessionCookie := auth.sessionCookie

	if useAuthHeader {
		req.Header.Set("Authorization", token)
		if sessionCookie != "" {
			req.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%s", sessionCookie))
		}
		c.SmcContext.Debug(fmt.Sprintf(
			"Sending %s %s with Authorization header: token=%q jsessionid=%q",
			opts.Method, newUrl, token[:min(len(token), 60)], sessionCookie,
		))
	} else {
		req.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%s", token))
		c.SmcContext.Debug(fmt.Sprintf(
			"Sending %s %s with JSESSIONID cookie: token=%q",
			opts.Method, newUrl, token[:min(len(token), 60)],
		))
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

	c.SmcContext.Trace(fmt.Sprintf("→ %s %s", opts.Method, newUrl))
	resp, err := client.Do(req)
	if err != nil {
		c.SmcContext.Trace(fmt.Sprintf("← %s %s error: %s", opts.Method, newUrl, err.Error()))
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.SmcContext.Trace(fmt.Sprintf(
		"← %s %s %d headers=%v body=%s",
		opts.Method, newUrl, resp.StatusCode, resp.Header, string(respBody),
	))

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
// the login endpoint and returns true if the session is still usable.
//
// It sets SkipAuthRetry: ping runs while GetClient holds the ClientManager's
// RWMutex (client_management.go), and the 401 recovery/proactive-refresh
// machinery can take up to ~31s (login retries) -- doing that here would
// block every other reader/writer of the ClientManager for that whole time.
// A definitive 401 is reported as "not usable" so the caller replaces the
// client and lets the normal lazy Login() on the next real request
// re-authenticate, outside this lock. Any other non-200 status (5xx, etc.)
// is inconclusive about session validity, so it does not tear down a
// possibly-still-good session over a transient server-side hiccup on the
// ping endpoint itself.
func (c *SmcClient) ping(ctx context.Context) bool {
	url := fmt.Sprintf("%s/%s/%s", c.BaseUrl, c.APIVersion, PingEntryPoint)
	opts := Options{
		Method:        "GET",
		URL:           url,
		Headers:       map[string]string{"Content-Type": "application/json"},
		Context:       ctx,
		SkipAuthRetry: true,
	}
	resp, err := c.DoRequest(opts)
	if err != nil {
		return false
	}
	return resp.StatusCode != http.StatusUnauthorized
}

// GetJSONHeaders returns standard headers for JSON API requests with
// authentication
func (c *SmcClient) GetJSONHeaders() map[string]string {
	c.loginMutex.Lock()
	useAuthHeader := c.UseAuthHeader
	token := c.Token
	c.loginMutex.Unlock()

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if useAuthHeader {
		headers["Authorization"] = token
	} else {
		headers["Cookie"] = fmt.Sprintf("JSESSIONID=%s", token)
	}

	return headers
}

// GetAuthHeaders returns headers for authenticated requests (no Content-Type)
func (c *SmcClient) GetAuthHeaders() map[string]string {
	c.loginMutex.Lock()
	useAuthHeader := c.UseAuthHeader
	token := c.Token
	c.loginMutex.Unlock()

	if useAuthHeader {
		return map[string]string{
			"Authorization": token,
		}
	}
	return map[string]string{
		"Cookie": fmt.Sprintf("JSESSIONID=%s", token),
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
	trustedCert string, domain string, tokenRefreshInterval time.Duration) (*SmcClient, error) {

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
		BaseUrl:              url,
		APIKey:               apikey,
		VerifySSL:            verifySSL,
		HTTPClient:           httpClient, // todo not used
		APIVersion:           apiVersion,
		TrustedCert:          trustedCert,
		Domain:               domain,
		SmcContext:           logCtx,
		OriginalCtx:          &ctx,
		TokenRefreshInterval: tokenRefreshInterval,
	}
	return c, nil
}

const loginMaxAttempts = 5

// loginRetryBaseDelay is the initial backoff between login attempts,
// doubled after each retry. A var so tests can shorten it.
var loginRetryBaseDelay = 1 * time.Second

// isTransientNetworkError reports whether a transport-level error is worth
// retrying. Certificate verification and DNS not-found failures are permanent;
// connection refused/reset, EOF, proxy and timeout errors are transient
// (the SMC or an intermediate may just not be ready yet).
func isTransientNetworkError(err error) bool {
	var urlErr *url.Error
	if !errors.As(err, &urlErr) {
		return false
	}
	var certErr *tls.CertificateVerificationError
	if errors.As(err, &certErr) {
		return false
	}
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
		return false
	}
	return true
}

// Login authenticates with the SMC API using the provided API key.
// Transient failures (HTTP 5xx, connection refused/reset, EOF, timeouts)
// are retried with exponential backoff: the SMC login endpoint is known to
// fail sporadically while the server is busy or still warming up (SMC-66463).
func (c *SmcClient) Login(ctx context.Context) error {
	if c.LoginConfirmed {
		c.SmcContext.Debug("Already logged to SMC API", map[string]any{
			"url": c.BaseUrl,
		})
		return nil
	}
	if ctx == nil {
		ctx = context.Background()
	}

	delay := loginRetryBaseDelay
	for attempt := 1; ; attempt++ {
		retryable, err := c.loginOnce(ctx)
		if err == nil {
			return nil
		}
		if !retryable || attempt >= loginMaxAttempts {
			return err
		}
		c.SmcContext.Warn("Transient login failure, retrying", map[string]any{
			"url":     c.BaseUrl,
			"attempt": attempt,
			"of":      loginMaxAttempts,
			"wait":    delay.String(),
			"error":   err.Error(),
		})
		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return err
		}
		delay *= 2
	}
}

// loginOnce performs a single login attempt. It returns the error of the
// attempt and whether that error is transient (worth retrying).
func (c *SmcClient) loginOnce(ctx context.Context) (retryable bool, err error) {
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
		return false, err
	}

	url := fmt.Sprintf("%s/%s/login", c.BaseUrl, c.APIVersion)
	opts := Options{
		Method:  "POST",
		URL:     url,
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    reqBody,
		Context: ctx,
	}
	// loginOnce always runs either under loginMutex (via ensureLoggedIn or
	// recoverExpiredSession) or single-threaded in tests calling Login
	// directly, so a direct field read is safe here; calling currentAuth()
	// would deadlock in the former case since sync.Mutex is not reentrant.
	auth := authSnapshot{
		useAuthHeader: c.UseAuthHeader,
		token:         c.Token,
		sessionCookie: c.SessionCookie,
	}
	resp, err := c.unsafeRequest(&opts, auth)
	if err != nil {
		c.SmcContext.Error("Login request failed", map[string]any{
			"error": err.Error(),
		})
		return isTransientNetworkError(err), err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		c.SmcContext.Error("Login failed with non-success status", map[string]any{
			"status_code": resp.StatusCode,
			"url":         url,
			"body":        string(resp.Body),
			"headers":     resp.Headers,
		})
		detail := string(resp.Body)
		if detail == "" {
			// An empty body means the SMC authentication logic likely never
			// answered; the headers help identify what did (Tomcat, proxy, ...).
			detail = fmt.Sprintf("(empty body; headers: %v)", resp.Headers)
		}
		// 5xx responses are transient (server busy / initializing); 4xx
		// (bad API key, forbidden) will not improve with a retry.
		return resp.StatusCode >= http.StatusInternalServerError,
			fmt.Errorf("login failed with status %d at %s: %s", resp.StatusCode, url, detail)
	}

	// Always extract JSESSIONID from Set-Cookie (needed alongside Authorization on HTTPS).
	for _, cookieStr := range resp.Headers["Set-Cookie"] {
		if !strings.Contains(cookieStr, "JSESSIONID=") {
			continue
		}
		for _, part := range strings.Split(cookieStr, ";") {
			part = strings.TrimSpace(part)
			if strings.HasPrefix(part, "JSESSIONID=") {
				c.SessionCookie = strings.TrimPrefix(part, "JSESSIONID=")
				break
			}
		}
	}

	// Try to get token from Authorization header first (for HTTPS)
	if authHeader := resp.Headers.Get("Authorization"); authHeader != "" {
		c.Token = authHeader
		c.UseAuthHeader = true
		c.LoginConfirmed = true
		c.TokenIssuedAt = time.Now()
		c.SmcContext.Debug(fmt.Sprintf(
			"Login OK (Authorization header): token=%q jsessionid=%q",
			authHeader[:min(len(authHeader), 60)],
			c.SessionCookie,
		))
		return false, nil
	}

	// Fallback: use JSESSIONID cookie only (for HTTP)
	if c.SessionCookie != "" {
		c.Token = c.SessionCookie
		c.UseAuthHeader = false
		c.LoginConfirmed = true
		c.SmcContext.Info(fmt.Sprintf("Login OK (JSESSIONID cookie): jsessionid=%q", c.SessionCookie))
	}
	return false, nil
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
		// Re-authenticating just to immediately log out again would be
		// pointless (and would leave a fresh, now-orphaned session open).
		SkipAuthRetry: true,
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

	// Clear the authentication state under loginMutex to prevent race with
	// concurrent requests reading Token/LoginConfirmed in GetJSONHeaders/GetAuthHeaders
	c.loginMutex.Lock()
	c.Token = ""
	c.LoginConfirmed = false
	c.loginMutex.Unlock()

	return nil
}

// ensureLoggedIn guards login with a mutex to prevent concurrent login races.
// Required because the client is a per-URL singleton and login is lazy.
// Multiple concurrent requests (especially in parallel mode) must not attempt
// to login simultaneously, as that creates multiple sessions which invalidate
// each other's tokens.
func (c *SmcClient) ensureLoggedIn(ctx context.Context) error {
	c.loginMutex.Lock()
	defer c.loginMutex.Unlock()
	if c.LoginConfirmed {
		return nil
	}
	return c.Login(ctx)
}

// DoRequest performs a generic HTTP request and returns ResponseData
// or error. In case of HTTP errors (4xx, 5xx), no error is returned,
// the caller must check the StatusCode in ResponseData.
//
// If the SMC rejects the request with 401 (the session token expired or
// was killed server-side), the session is transparently
// recovered and the request is retried once, unless opts.SkipAuthRetry
// is set.
func (c *SmcClient) DoRequest(opts Options) (*ResponseData, error) {
	// Requests the caller marked parallelizable (see parallel.go, disabled
	// by default) skip the per-host lock below; concurrency for them is
	// bounded by terraform apply --parallelism, not by this client.
	// Everything else keeps today's behavior: a full per-host lock that
	// serializes every request end-to-end.
	canParallelize := opts.CanParallelize != nil && opts.CanParallelize()
	if !canParallelize {
		// Lock by base URL to prevent concurrent requests
		// Using different lock by base URL will allow
		// processing if several SMC providers are set.
		lockId := GetLockId(opts.URL)
		GetLockRequests().Enter(lockId)
		defer GetLockRequests().Leave(lockId)
	}

	// Ensure login with mutex protection to prevent concurrent login races
	// that would create multiple sessions and invalidate tokens (401 errors).
	if err := c.ensureLoggedIn(opts.Context); err != nil {
		return nil, fmt.Errorf("failed to login before request: %w", err)
	}

	if !opts.SkipAuthRetry {
		c.maybeProactivelyRefreshToken(opts.Context)
	}

	for _, hook := range c.BeforeRequestHooks {
		err := hook(c, &opts)
		if err != nil {
			return nil, err
		}
	}

	auth := c.currentAuth()
	resp, err := c.unsafeRequest(&opts, auth)
	if err != nil || resp.StatusCode != http.StatusUnauthorized || opts.SkipAuthRetry {
		return resp, err
	}

	ctx := opts.Context
	if ctx == nil {
		ctx = context.Background()
	}
	c.SmcContext.Debug("Received HTTP 401, session token likely expired; attempting recovery", map[string]any{
		"url": opts.URL,
	})
	if recoverErr := c.recoverExpiredSession(ctx, auth.token); recoverErr != nil {
		c.SmcContext.Error("Failed to recover SMC API session after token expiry", map[string]any{
			"url":   opts.URL,
			"error": recoverErr.Error(),
		})
		return resp, nil
	}

	// Re-run hooks before the retried request too: a hook may need to
	// re-establish per-session state (e.g. OpenVpnForEdit) under the new
	// token/session obtained by recoverExpiredSession.
	for _, hook := range c.BeforeRequestHooks {
		if err := hook(c, &opts); err != nil {
			return nil, err
		}
	}

	return c.unsafeRequest(&opts, c.currentAuth())
}

// recoverExpiredSession attempts to restore a usable SMC API session after
// a request was rejected with 401 (session tokens now expire
// server-side, default after 30 min). It first tries the lightweight
// refresh_token service, which mints a new token bound to the same session
// without a full re-authentication. If that is unavailable (older SMC
// without the endpoint) or fails for any other reason, it falls back to a
// full Login.
//
// Holds loginMutex for its whole duration, like ensureLoggedIn: in parallel
// mode, multiple concurrent requests can hit a 401 for the same expired
// token around the same time. failedToken is the token the failing request
// was sent with (from the auth snapshot DoRequest took before sending it);
// if c.Token no longer matches it once the lock is acquired, a concurrent
// request already recovered the session, and refreshing or re-logging in
// again here would invalidate that other goroutine's still-fresh token out
// from under its own in-flight retry.
func (c *SmcClient) recoverExpiredSession(ctx context.Context, failedToken string) error {
	c.loginMutex.Lock()
	defer c.loginMutex.Unlock()

	if c.Token != failedToken {
		c.SmcContext.Debug("Session already recovered by a concurrent request, skipping", map[string]any{
			"url": c.BaseUrl,
		})
		return nil
	}

	if c.UseAuthHeader {
		refreshErr := c.refreshTokenLocked(ctx)
		if refreshErr == nil {
			c.SmcContext.Info("SMC API session token refreshed via refresh_token endpoint", map[string]any{
				"url": c.BaseUrl,
			})
			return nil
		}
		c.SmcContext.Warn("SMC API token refresh failed, falling back to full re-login", map[string]any{
			"url":   c.BaseUrl,
			"error": refreshErr.Error(),
		})
	}

	c.LoginConfirmed = false
	if err := c.Login(ctx); err != nil {
		return fmt.Errorf("failed to re-authenticate after session expiry: %w", err)
	}
	// A full re-login replaces the session entirely: any VPN this client had
	// open for edit under the old (now-dead) session no longer is, on the
	// server. Drop those markers so the hooks DoRequest re-runs after this
	// returns will genuinely re-open the VPN under the new session instead
	// of short-circuiting on a stale "already open" entry.
	GetVpnEditStateManager().DropEntriesForClient(c)
	c.SmcContext.Info("SMC API session re-authenticated via full login after token expiry", map[string]any{
		"url": c.BaseUrl,
	})
	return nil
}

// maybeProactivelyRefreshToken refreshes the SMC API session token before it
// expires, if the caller configured TokenRefreshInterval. This
// keeps a long-running apply (many resources over the SMC's configured
// SMC_API_TOKEN_VALIDITY) refreshing ahead of expiry so it never actually
// hits a 401 in normal use, rather than relying solely on the reactive
// recovery in DoRequest. Best-effort: a failure here is logged and
// swallowed, since the reactive 401 path still covers it on the next
// request. Holds loginMutex for the same reason as recoverExpiredSession.
func (c *SmcClient) maybeProactivelyRefreshToken(ctx context.Context) {
	c.loginMutex.Lock()
	defer c.loginMutex.Unlock()

	if !c.UseAuthHeader || c.TokenRefreshInterval <= 0 || !c.LoginConfirmed {
		return
	}
	lastActivity := c.TokenIssuedAt
	if c.lastRefreshAttempt.After(lastActivity) {
		lastActivity = c.lastRefreshAttempt
	}
	if time.Since(lastActivity) < c.TokenRefreshInterval {
		return
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if err := c.refreshTokenLocked(ctx); err != nil {
		c.lastRefreshAttempt = time.Now()
		c.SmcContext.Warn("Proactive SMC API token refresh failed, will rely on reactive recovery on the next 401", map[string]any{
			"url":   c.BaseUrl,
			"error": err.Error(),
		})
		return
	}
	c.SmcContext.Info("SMC API session token proactively refreshed via refresh_token endpoint", map[string]any{
		"url": c.BaseUrl,
	})
}

// RefreshToken requests a new SMC API session token bound to the current
// session and replaces the stored one. The SMC invalidates the
// previous token, so every subsequent request transparently uses the new
// one. This targets the POST /<version>/system/refresh_token service,
// available on SMC >= 7.6.0 (and its 7.5.1/7.4.2/7.3.5 backports). It only
// applies to the Authorization-header auth mode: JSESSIONID-cookie-only
// sessions (HTTP, "Use SSL for session ID" disabled) fall back to a full
// re-login in recoverExpiredSession.
func (c *SmcClient) RefreshToken(ctx context.Context) error {
	c.loginMutex.Lock()
	defer c.loginMutex.Unlock()
	return c.refreshTokenLocked(ctx)
}

// refreshTokenLocked is RefreshToken's implementation. Callers must already
// hold loginMutex (recoverExpiredSession, maybeProactivelyRefreshToken); like
// loginOnce, it reads/writes the client's auth fields directly rather than
// through currentAuth()/RefreshToken(), since sync.Mutex is not reentrant.
func (c *SmcClient) refreshTokenLocked(ctx context.Context) error {
	if !c.LoginConfirmed {
		return fmt.Errorf("no active SMC API session to refresh the token for")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	url := fmt.Sprintf("%s/%s/system/refresh_token", c.BaseUrl, c.APIVersion)
	opts := Options{
		Method:  "POST",
		URL:     url,
		Context: ctx,
	}
	auth := authSnapshot{
		useAuthHeader: c.UseAuthHeader,
		token:         c.Token,
		sessionCookie: c.SessionCookie,
	}
	resp, err := c.unsafeRequest(&opts, auth)
	if err != nil {
		return fmt.Errorf("token refresh request failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("token refresh failed with status %d at %s: %s", resp.StatusCode, url, string(resp.Body))
	}
	newToken := resp.Headers.Get("Authorization")
	if newToken == "" {
		return fmt.Errorf("token refresh response did not contain a new Authorization token")
	}
	c.Token = newToken
	c.TokenIssuedAt = time.Now()
	return nil
}
