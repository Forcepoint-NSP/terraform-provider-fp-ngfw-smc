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
