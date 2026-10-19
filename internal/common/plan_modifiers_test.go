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

package common

import "testing"

func TestStringPlanModifiers(t *testing.T) {
	if got := StringPlanModifiers(false); got != nil {
		t.Errorf("expected nil when disabled, got %v", got)
	}
	if got := StringPlanModifiers(true); len(got) != 1 {
		t.Errorf("expected one modifier when enabled, got %v", got)
	}
}

func TestMapPlanModifiers(t *testing.T) {
	if got := MapPlanModifiers(false); got != nil {
		t.Errorf("expected nil when disabled, got %v", got)
	}
	if got := MapPlanModifiers(true); len(got) != 1 {
		t.Errorf("expected one modifier when enabled, got %v", got)
	}
}
