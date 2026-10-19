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

package resource

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

func TestIsRetryableParallelCreateError(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"conflict", smc.NewStatusError(http.StatusConflict, "dup"), true},
		{"locked 422", smc.NewStatusError(http.StatusUnprocessableEntity, "Element is Locked by admin"), true},
		{"blocked 422", smc.NewStatusError(http.StatusUnprocessableEntity, "operation is blocked"), false},
		{"unlocked 422", smc.NewStatusError(http.StatusUnprocessableEntity, "element is unlocked"), false},
		{"validation 422", smc.NewStatusError(http.StatusUnprocessableEntity, "invalid field"), false},
		{"wrapped conflict", fmt.Errorf("x: %w", smc.NewStatusError(http.StatusConflict, "")), true},
		{"other status", smc.NewStatusError(http.StatusInternalServerError, "locked"), false},
		{"non-status error", fmt.Errorf("boom"), false},
	}
	for _, tc := range cases {
		if got := isRetryableParallelCreateError(tc.err); got != tc.want {
			t.Errorf("%s: got %v, want %v", tc.name, got, tc.want)
		}
	}
}
