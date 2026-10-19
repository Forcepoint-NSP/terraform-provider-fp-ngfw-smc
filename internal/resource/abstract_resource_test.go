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
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terraform-providers/terraform-provider-smc/internal/customfield"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

type HostResourceModel struct {
	ID types.String `tfsdk:"id"`
}

type linkedResourceModel struct {
	ID types.String                  `tfsdk:"id"`
	Lk customfield.Map[types.String] `tfsdk:"link" json:"-"`
}

// Regression test for a scheme-switch bug (fb1769e): the "id" and "link"
// schema attributes carry a UseStateForUnknown plan modifier, so a stale
// http:// href left in state after an SMC http->https switch gets carried
// forward as the planned value; Update then normalizes it for real and
// produces "Provider produced inconsistent result after apply". Read must
// normalize both, regardless of whether drift detection is enabled, since
// normalizeStateHrefs is called unconditionally in abstract_resource.go's
// Read() before the drift_detection opt-in check.
func TestNormalizeStateHrefs_RewritesSchemeUnconditionally(t *testing.T) {
	ctx := context.Background()
	client := &smc.SmcClient{BaseUrl: "https://smc:8443", APIVersion: "7.5"}

	links, diags := customfield.NewMap[types.String](ctx, map[string]types.String{
		"add_route": types.StringValue("http://smc:8082/7.4/elements/single_fw/4820/add_route"),
	})
	if diags.HasError() {
		t.Fatalf("failed to build test link map: %v", diags)
	}

	state := &linkedResourceModel{
		ID: types.StringValue("http://smc:8082/7.4/elements/single_fw/4820"),
		Lk: links,
	}

	changed, err := normalizeStateHrefs(ctx, client, state)
	if err != nil {
		t.Fatalf("normalizeStateHrefs returned an error: %v", err)
	}
	if !changed {
		t.Fatal("expected normalizeStateHrefs to report a change")
	}

	if got, want := state.ID.ValueString(), "https://smc:8443/7.5/elements/single_fw/4820"; got != want {
		t.Errorf("id: got %q, want %q", got, want)
	}

	values, diags := state.Lk.Value(ctx)
	if diags.HasError() {
		t.Fatalf("failed to read link map: %v", diags)
	}
	if got, want := values["add_route"].ValueString(), "https://smc:8443/7.5/elements/single_fw/4820/add_route"; got != want {
		t.Errorf("link[add_route]: got %q, want %q", got, want)
	}
}

type nestedLinkedModel struct {
	ID   types.String                  `tfsdk:"id"`
	Lk   customfield.Map[types.String] `tfsdk:"link" json:"-"`
	Name types.String                  `tfsdk:"name"`
}

func (n nestedLinkedModel) GetSliceIds(ctx context.Context) []string {
	return []string{n.Name.ValueString()}
}

type parentWithNestedLinksModel struct {
	ID         types.String                  `tfsdk:"id"`
	Lk         customfield.Map[types.String] `tfsdk:"link" json:"-"`
	Interfaces *[]nestedLinkedModel          `tfsdk:"interfaces"`
}

// Regression test: nested blocks (e.g. physical_interface under
// smc_single_fw) have their own "id"/"link" attributes, populated the
// same way as the top-level resource. normalizeStateHrefs must reach
// them too, not just the top-level struct.
func TestNormalizeStateHrefs_RewritesNestedBlocks(t *testing.T) {
	ctx := context.Background()
	client := &smc.SmcClient{BaseUrl: "https://smc:8443", APIVersion: "7.5"}

	nestedLinks, diags := customfield.NewMap[types.String](ctx, map[string]types.String{
		"self": types.StringValue("http://smc:8082/7.4/elements/single_fw/4820/interfaces/0"),
	})
	if diags.HasError() {
		t.Fatalf("failed to build nested test link map: %v", diags)
	}

	interfaces := []nestedLinkedModel{
		{
			ID:   types.StringValue("http://smc:8082/7.4/elements/single_fw/4820/interfaces/0"),
			Lk:   nestedLinks,
			Name: types.StringValue("eth0"),
		},
	}

	state := &parentWithNestedLinksModel{
		ID:         types.StringValue("http://smc:8082/7.4/elements/single_fw/4820"),
		Interfaces: &interfaces,
	}

	changed, err := normalizeStateHrefs(ctx, client, state)
	if err != nil {
		t.Fatalf("normalizeStateHrefs returned an error: %v", err)
	}
	if !changed {
		t.Fatal("expected normalizeStateHrefs to report a change")
	}

	if got, want := (*state.Interfaces)[0].ID.ValueString(), "https://smc:8443/7.5/elements/single_fw/4820/interfaces/0"; got != want {
		t.Errorf("interfaces[0].id: got %q, want %q", got, want)
	}

	values, diags := (*state.Interfaces)[0].Lk.Value(ctx)
	if diags.HasError() {
		t.Fatalf("failed to read nested link map: %v", diags)
	}
	if got, want := values["self"].ValueString(), "https://smc:8443/7.5/elements/single_fw/4820/interfaces/0"; got != want {
		t.Errorf("interfaces[0].link[self]: got %q, want %q", got, want)
	}
}

func TestNormalizeStateHrefs_NoopWhenAlreadyNormalized(t *testing.T) {
	ctx := context.Background()
	client := &smc.SmcClient{BaseUrl: "https://smc:8443", APIVersion: "7.5"}

	links, diags := customfield.NewMap[types.String](ctx, map[string]types.String{
		"add_route": types.StringValue("https://smc:8443/7.5/elements/single_fw/4820/add_route"),
	})
	if diags.HasError() {
		t.Fatalf("failed to build test link map: %v", diags)
	}

	state := &linkedResourceModel{
		ID: types.StringValue("https://smc:8443/7.5/elements/single_fw/4820"),
		Lk: links,
	}

	changed, err := normalizeStateHrefs(ctx, client, state)
	if err != nil {
		t.Fatalf("normalizeStateHrefs returned an error: %v", err)
	}
	if changed {
		t.Fatal("expected normalizeStateHrefs to report no change when state is already normalized")
	}
}

func TestGetId(t *testing.T) {
	type SampleResourceModel struct {
		ID types.String `tfsdk:"id"`
	}

	sample := SampleResourceModel{
		ID: types.StringValue("sample-id"),
	}

	idValue, err := GetField(sample, "ID")
	if err != nil {
		t.Fatalf("GetField returned an error: %v", err)
	}
	id, ok := idValue.(types.String)
	if !ok {
		t.Fatalf("Expected types.String, got %T", idValue)
	}
	if id.ValueString() != "sample-id" {
		t.Fatalf("Expected id 'sample-id', got '%s'", id.ValueString())
	}
}

func TestSetId(t *testing.T) {
	const newID = "new-id"
	const testID = "test-id"

	// Test successful ID setting
	model := &HostResourceModel{
		ID: types.StringValue("old-id"),
	}

	err := setId(model, newID)
	if err != nil {
		t.Fatalf("setId returned error: %v", err)
	}

	if model.ID.ValueString() != newID {
		t.Errorf("expected ID to be '%s', got '%s'", newID, model.ID.ValueString())
	}

	// Test error on non-pointer input
	err = setId(HostResourceModel{}, testID)
	if err == nil {
		t.Error("expected error for non-pointer input")
	}

	// Test error on nil pointer
	var nilModel *HostResourceModel
	err = setId(nilModel, testID)
	if err == nil {
		t.Error("expected error for nil pointer")
	}

	// Test error on struct without ID field
	type NoIdModel struct {
		Name types.String `tfsdk:"name"`
	}
	noIdModel := &NoIdModel{Name: types.StringValue("test")}
	err = setId(noIdModel, testID)
	if err == nil {
		t.Error("expected error for struct without ID field")
	}

	// Test error on wrong field type
	type WrongTypeModel struct {
		ID string `tfsdk:"id"`
	}
	wrongTypeModel := &WrongTypeModel{ID: "old-id"}
	err = setId(wrongTypeModel, newID)
	if err == nil {
		t.Error("expected error for wrong field type")
	}
}

func TestIsConflict(t *testing.T) {
	conflictErr := fmt.Errorf("create host failed: %w", smc.NewStatusError(http.StatusConflict, "conflict"))
	if !isConflict(conflictErr) {
		t.Error("expected a wrapped 409 StatusError to be detected as a conflict")
	}

	notFoundErr := fmt.Errorf("read host failed: %w", smc.NewStatusError(http.StatusNotFound, "not found"))
	if isConflict(notFoundErr) {
		t.Error("expected a 404 StatusError not to be detected as a conflict")
	}

	if isConflict(errors.New("some other error")) {
		t.Error("expected a non-StatusError not to be detected as a conflict")
	}

	if isConflict(nil) {
		t.Error("expected a nil error not to be detected as a conflict")
	}
}
