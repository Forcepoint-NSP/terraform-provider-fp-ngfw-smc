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
	"testing"
)

func TestReplaceInURL(t *testing.T) {
	tests := []struct {
		name      string
		rawURL    string
		scheme    string
		host      string
		port      string
		want      string
		wantError bool
	}{
		{
			name:      "change scheme, host, and port",
			rawURL:    "http://example.com/path?a=1#frag",
			scheme:    "https",
			host:      "api.example.org",
			port:      "8443",
			want:      "https://api.example.org:8443/path?a=1#frag",
			wantError: false,
		},
		{
			name:      "keep host, change scheme and port",
			rawURL:    "https://user:pass@host:8080/p",
			scheme:    "http",
			host:      "",
			port:      "9090",
			want:      "http://user:pass@host:9090/p",
			wantError: false,
		},
		{
			name:      "IPv6 literal with scheme and port change",
			rawURL:    "http://[2001:db8::1]/x",
			scheme:    "https",
			host:      "[2001:db8::1]",
			port:      "443",
			want:      "https://[2001:db8::1]:443/x",
			wantError: false,
		},
		{
			name:      "replace IPv6 host only",
			rawURL:    "http://[2001:db8::1]:8080/x",
			scheme:    "",
			host:      "2001:db8::2",
			port:      "",
			want:      "http://[2001:db8::2]:8080/x",
			wantError: false,
		},
		{
			name:      "change nothing",
			rawURL:    "https://example.com",
			scheme:    "",
			host:      "",
			port:      "",
			want:      "https://example.com",
			wantError: false,
		},
		{
			name:      "same host no change",
			rawURL:    "http://example.com",
			scheme:    "",
			host:      "example.com",
			port:      "",
			want:      "http://example.com",
			wantError: false,
		},
		{
			name:      "invalid URL",
			rawURL:    "ht!tp://invalid",
			scheme:    "https",
			host:      "example.com",
			port:      "",
			want:      "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReplaceInURL(tt.rawURL, tt.scheme, tt.host, tt.port)
			if (err != nil) != tt.wantError {
				t.Errorf("ReplaceInURL() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("ReplaceInURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceVersionInURL(t *testing.T) {
	tests := []struct {
		name      string
		rawURL    string
		version   string
		want      string
		wantError bool
	}{
		{
			name:    "replace version in element href",
			rawURL:  "https://smc.example.com:8082/7.4/elements/single_fw/1409",
			version: "7.6",
			want:    "https://smc.example.com:8082/7.6/elements/single_fw/1409",
		},
		{
			name:    "replace version in nested sub-resource href",
			rawURL:  "http://localhost:8082/7.4/elements/single_fw/268571118/internal_gateway/268439099",
			version: "7.6",
			want:    "http://localhost:8082/7.6/elements/single_fw/268571118/internal_gateway/268439099",
		},
		{
			name:    "version only path",
			rawURL:  "https://smc.example.com:8082/7.4/login",
			version: "7.6",
			want:    "https://smc.example.com:8082/7.6/login",
		},
		{
			name:    "preserve query and fragment",
			rawURL:  "https://smc.example.com:8082/7.4/elements/host?filter=foo#frag",
			version: "7.6",
			want:    "https://smc.example.com:8082/7.6/elements/host?filter=foo#frag",
		},
		{
			name:    "minor.major with two-digit minor",
			rawURL:  "https://smc.example.com:8082/6.11/elements/host/1",
			version: "7.6",
			want:    "https://smc.example.com:8082/7.6/elements/host/1",
		},
		{
			name:    "empty version leaves URL unchanged",
			rawURL:  "https://smc.example.com:8082/7.4/elements/host/1",
			version: "",
			want:    "https://smc.example.com:8082/7.4/elements/host/1",
		},
		{
			name:    "non-version first segment left untouched",
			rawURL:  "https://smc.example.com:8082/api/elements/host/1",
			version: "7.6",
			want:    "https://smc.example.com:8082/api/elements/host/1",
		},
		{
			name:    "no path left untouched",
			rawURL:  "https://smc.example.com:8082",
			version: "7.6",
			want:    "https://smc.example.com:8082",
		},
		{
			name:    "IPv6 host",
			rawURL:  "https://[2001:db8::1]:8082/7.4/elements/host/1",
			version: "7.6",
			want:    "https://[2001:db8::1]:8082/7.6/elements/host/1",
		},
		{
			name:      "invalid URL",
			rawURL:    "ht!tp://invalid/7.4/x",
			version:   "7.6",
			want:      "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReplaceVersionInURL(tt.rawURL, tt.version)
			if (err != nil) != tt.wantError {
				t.Errorf("ReplaceVersionInURL() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("ReplaceVersionInURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeHref(t *testing.T) {
	tests := []struct {
		name       string
		baseURL    string
		apiVersion string
		href       string
		want       string
	}{
		{
			name:       "version upgrade only (SMC-66510)",
			baseURL:    "https://smc.example.com:8082",
			apiVersion: "7.6",
			href:       "https://smc.example.com:8082/7.4/elements/single_fw/1409",
			want:       "https://smc.example.com:8082/7.6/elements/single_fw/1409",
		},
		{
			name:       "http to https switch (SMC-64735)",
			baseURL:    "https://smc.example.com:8082",
			apiVersion: "7.4",
			href:       "http://smc.example.com:8082/7.4/elements/host/1",
			want:       "https://smc.example.com:8082/7.4/elements/host/1",
		},
		{
			name:       "scheme, host, port and version all change",
			baseURL:    "https://proxy.internal:443",
			apiVersion: "7.6",
			href:       "http://smc.example.com:8082/7.4/elements/host/1",
			want:       "https://proxy.internal:443/7.6/elements/host/1",
		},
		{
			name:       "already normalized is a no-op",
			baseURL:    "https://smc.example.com:8082",
			apiVersion: "7.6",
			href:       "https://smc.example.com:8082/7.6/elements/host/1",
			want:       "https://smc.example.com:8082/7.6/elements/host/1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SmcClient{BaseUrl: tt.baseURL, APIVersion: tt.apiVersion}
			got, err := c.NormalizeHref(tt.href)
			if err != nil {
				t.Fatalf("NormalizeHref() unexpected error = %v", err)
			}
			if got != tt.want {
				t.Errorf("NormalizeHref() got = %v, want %v", got, tt.want)
			}
		})
	}
}
