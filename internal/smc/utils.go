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
	"fmt"
	"net"
	"net/url"
	"strings"
)

func ReplaceBaseInURL(origUrl, baseUrl string) (string, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}
	return ReplaceInURL(origUrl, u.Scheme, u.Hostname(), u.Port())
}

// ReplaceBaseInURL takes an input URL string and replaces its scheme, host, and port.
// Pass empty values for any field you don't want to change (e.g., scheme="", host="", port="").
// Returns the updated URL string.
func ReplaceInURL(rawURL, scheme, host, port string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	// Scheme
	if scheme != "" {
		u.Scheme = scheme
	}

	// Host and port
	// u.Host may include userinfo@host:port or just host:port (userinfo is actually in u.User)
	hostname, currentPort := splitHostPortSafe(u.Host)
	if host == "" && hostname == "" && u.Host != "" {
		// Handle cases where u.Host is set but hostname extraction failed (rare)
		hostname = u.Host
	}
	if host != "" {
		hostname = host
	}
	// Decide which port to use
	usePort := currentPort
	if port != "" {
		usePort = port
	}

	// Rebuild host: ensure IPv6 literals are bracketed
	if hostname != "" {
		// If host is an IPv6 literal without brackets, add them
		if ip := net.ParseIP(hostname); ip != nil && strings.Contains(hostname, ":") && !strings.HasPrefix(hostname, "[") {
			hostname = "[" + hostname + "]"
		}
		if usePort != "" {
			u.Host = net.JoinHostPort(stripBrackets(hostname), usePort)
		} else {
			u.Host = hostname
		}
	}

	return u.String(), nil
}

// splitHostPortSafe splits a host[:port] allowing for IPv6 literals with brackets.
// If no port is present, returns (host, "").
// If input is empty, returns ("", "").
func splitHostPortSafe(hostport string) (string, string) {
	if hostport == "" {
		return "", ""
	}
	// If it already has brackets, use stdlib
	if strings.HasPrefix(hostport, "[") {
		h, p, err := net.SplitHostPort(hostport)
		if err == nil {
			return stripBrackets(h), p
		}
		// Might be bracketed host without port: [::1]
		return stripBrackets(hostport), ""
	}
	// Try regular split
	h, p, err := net.SplitHostPort(hostport)
	if err == nil {
		return h, p
	}
	// Fallback: no port present
	return hostport, ""
}

// stripBrackets removes enclosing [] from IPv6 literal if present.
func stripBrackets(h string) string {
	if strings.HasPrefix(h, "[") && strings.HasSuffix(h, "]") {
		return h[1 : len(h)-1]
	}
	return h
}
