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
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-smc/internal/apijson"
	"github.com/terraform-providers/terraform-provider-smc/internal/common"
	"github.com/terraform-providers/terraform-provider-smc/internal/customfield"
	"github.com/terraform-providers/terraform-provider-smc/internal/datastore"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

const MAX_RETRY = 4
const SLEEP_BETWEEN_RETRIES_SEC = 5 * time.Second

type ResourceBaseCustomOperations[T any] interface {
	GetCreateRequestParams(ctx context.Context, data *T) (map[string]string, error)
	BeforeDelete(ctx context.Context, data *T, id string) error
	ModelToJson(data *T) ([]byte, error)
}

type ResourceBase[T any] struct {
	ResourceType  string // name of the resource in terraform
	IsSubResource bool   // true if this resource is a sub-resource
	// without the provider prefix, eg "smc_host" -> "host"

	Config   *smc.GenericCRUDConfig
	Dispatch ResourceBaseCustomOperations[T]
}

// GetField returns the value of a struct field by name.
// Returns an error if the field does not exist or is not accessible.

// todo move me elsewhere
func GetField(mystruct any, fieldName string) (any, error) {
	// Get the reflect.Value of the struct
	val := reflect.ValueOf(mystruct)

	// Ensure the input is a struct (or a pointer to a struct)
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %s", val.Kind())
	}

	// Get the field by name
	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, fmt.Errorf("field '%s' not found", fieldName)
	}

	// Return the field's value as an any
	return field.Interface(), nil
}

// SetField sets the value of a struct field by name.
// eg
// err = SetField(&model, "ID", "new-id-value")
func SetField(obj any, fieldName string, value any) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("obj must be a non-nil pointer")
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("obj must point to a struct")
	}

	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("no such field: %s", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("cannot set field: %s", fieldName)
	}

	val := reflect.ValueOf(value)
	if field.Type() != val.Type() {
		return fmt.Errorf("provided value type (%s) doesn't match field type (%s)", val.Type(), field.Type())
	}

	field.Set(val)
	return nil
}

func getFromRefField(data any) (string, error) {
	parentUrl, err := GetField(data, "FromRef")
	if err != nil {
		return "", err
	}
	parentUrlTf := parentUrl.(types.String)
	if parentUrlTf.IsNull() || parentUrlTf.IsUnknown() {
		return "", nil
	}
	return parentUrl.(types.String).ValueString(), nil
}

func getId(data any) (string, error) {
	id, err := GetField(data, "ID")
	if err != nil {
		return "", err
	}
	if id.(types.String).IsNull() || id.(types.String).IsUnknown() {
		return "", nil
	}
	return id.(types.String).ValueString(), nil
}

func setId(data any, value string) error {
	return SetField(data, "ID", types.StringValue(value))
}

// returns additional query parameters for create requests
// overriden by specific resources as needed.
func (r *ResourceBase[T]) GetCreateRequestParams(ctx context.Context, data *T) (map[string]string, error) {
	return nil, nil
}

func (r *ResourceBase[T]) BeforeDelete(ctx context.Context, data *T, id string) error {
	return nil
}

func (r *ResourceBase[T]) ModelToJson(data *T) ([]byte, error) {
	return apijson.MarshalRoot(*data)
}

// smcUpdateResource Sends an update request to the smc for the given
// resource ID with the provided data
//
// Parameters:
//   - id - url of the resource to update.
//     eg http://localhost:8082/7.4/elements/single_fw/268569103
//   - data - resource data to update (from config)
func (r *ResourceBase[T]) smcUpdateResource(
	ctx context.Context, data *T, stateData *T, id string, diag *diag.Diagnostics) (*smc.ResponseData, error) {
	// needs up-to-date etag for an actual update (maybe in the future
	// we should have a mechanism to get the etag without having to
	// reload the whole object).
	// for the "fake" create we also need the json payload to get the "key" fields.
	resolved_id, err := common.ResolveRef(ctx, r.Config, id)
	if err != nil {
		diag.AddError(
			"Error resolving resource reference",
			fmt.Sprintf("Could not resolve reference for update: %s", err.Error()),
		)
		return nil, fmt.Errorf("failed to resolve ref for update. href='%s': %w",
			id, err)
	}
	tflog.Debug(ctx, fmt.Sprintf("Resolved ref for update: %s -> %s", id, resolved_id))
	if resolved_id == "" {
		diag.AddError(
			"Error resolving resource reference",
			fmt.Sprintf("Could not resolve reference for update: %s", id),
		)
		return nil, fmt.Errorf("could not resolve ref for update: %s", id)
	}
	for range MAX_RETRY {
		readResp, err := r.Config.ReadResourceByHref(resolved_id)
		if err != nil {
			diag.AddError(
				"Error reading resource for update",
				fmt.Sprintf("Could not read current %s for update: %s", r.Config.ResourceType, err.Error()),
			)
			return nil, fmt.Errorf("failed to read current %s for update. href='%s': %w",
				r.Config.ResourceType, id, err)
		}
		if readResp == nil {
			diag.AddError(
				"Error reading resource for update",
				fmt.Sprintf("%s not found for update: %s", r.Config.ResourceType, id),
			)
			return nil, fmt.Errorf("%s not found for update: %s", r.Config.ResourceType, id)
		}

		if stateData == nil {
			stateData = new(T)
			if err := apijson.UnmarshalRoot(readResp.Body, stateData); err != nil {
				diag.AddError(
					"Error parsing resource",
					fmt.Sprintf("Could not parse: %s", err.Error()),
				)
				return nil, err
			}
		}

		// Build the outgoing payload from a deep copy of the plan, not
		// the plan itself: backfilling unset scalar attributes (e.g. a
		// comment set only via the SMC GUI) must not touch data, since
		// data is later used to set Terraform state, and Terraform
		// requires a non-Computed attribute's final state to exactly
		// match its planned value ("provider produced inconsistent
		// result after apply" otherwise).
		payloadData := new(T)
		planJson, err := apijson.MarshalRoot(*data)
		if err != nil {
			diag.AddError(
				"Error marshalling resource data",
				fmt.Sprintf("Could not marshal resource data to JSON: %s", err.Error()),
			)
			return nil, err
		}
		if err := apijson.UnmarshalRoot(planJson, payloadData); err != nil {
			diag.AddError(
				"Error parsing resource data",
				fmt.Sprintf("Could not parse resource data: %s", err.Error()),
			)
			return nil, err
		}

		// Backfill the payload copy with the freshly-read current
		// object, so any scalar attribute the user didn't configure
		// keeps its current value instead of going out as null and
		// being cleared server-side. Computed fields (at every nesting
		// level, e.g. each routing_node level's own "key") and the
		// plan's own explicit values and list/object structure are
		// unaffected, same as before.
		err = MergeResourceUpdatePayload(ctx /*src*/, stateData, payloadData /*dest*/)
		if err != nil {
			// todo diag.AddError
			return nil, fmt.Errorf("failed to merge resource models for update: %w", err)
		}

		apiData, err := r.Dispatch.ModelToJson(payloadData)
		if err != nil {
			diag.AddError(
				"Error marshalling resource data",
				fmt.Sprintf("Could not marshal resource data to JSON: %s", err.Error()),
			)
			return nil, err
		}

		resp, err := smc.UpdateResource(r.Config, apiData, resolved_id, readResp.ETag)
		if err != nil {
			if isConflict(err) {
				tflog.Warn(ctx, fmt.Sprintf("Conflict updating resource, retrying: %s", id))
				time.Sleep(SLEEP_BETWEEN_RETRIES_SEC)
				continue // retry
			}
			tflog.Error(ctx, fmt.Sprintf("Error update: %s", id))
			diag.AddError(
				"Error updating resource",
				fmt.Sprintf("Could not update %s : %s", r.ResourceType, err.Error()),
			)
			return nil, err
		} // end error check after update call
		return resp, nil
	} // for range MAX_RETRY

	// if we reach here, we exhausted retries
	diag.AddError(
		"Error updating resource",
		fmt.Sprintf("Could not update %s due to repeated conflicts. Please retry.", r.ResourceType),
	)
	return nil, fmt.Errorf("failed to update %s due to repeated conflicts",
		r.ResourceType)
}

// isConflict reports whether err wraps a smc.StatusError for HTTP 409
// Conflict.
func isConflict(err error) bool {
	var statusErr *smc.StatusError
	return errors.As(err, &statusErr) && statusErr.Code == http.StatusConflict
}

// isRetryableParallelCreateError reports whether a failed create should be
// retried in parallel mode: 409 (transient conflict) or a 422 whose message
// says the parent is locked ("locked by..."). Other 422s are validation
// failures and must surface immediately.
func isRetryableParallelCreateError(err error) bool {
	var statusErr *smc.StatusError
	if !errors.As(err, &statusErr) {
		return false
	}
	return statusErr.Code == http.StatusConflict ||
		(statusErr.Code == http.StatusUnprocessableEntity &&
			strings.Contains(strings.ToLower(statusErr.Message), "locked by"))
}

// smcCreateResource sends a create request for the given
// sub-resource with the provided data
//
// Parameters:
//   - data - resource data to update (from config)
func (r *ResourceBase[T]) smcCreateResource(
	ctx context.Context, data *T, diag *diag.Diagnostics) (*smc.ResponseData, error) {

	apiData, err := r.Dispatch.ModelToJson(data)
	if err != nil {
		diag.AddError(
			"Error marshalling resource data",
			fmt.Sprintf("Could not marshal resource data to JSON: %s", err.Error()),
		)
		return nil, err
	}

	var lastErr error
	for range MAX_RETRY {
		requestResp, err := smc.CreateResource(r.Config, apiData)
		if err == nil {
			return requestResp, nil
		}
		lastErr = err
		// Only retry in parallel mode; serially, 409 is a legitimate user error
		// (duplicate element name) that should fail immediately, not stall for ~20s.
		// A genuine duplicate name is indistinguishable here and is retried too.
		if isRetryableParallelCreateError(err) && smc.IsResourceTypeParallel(r.ResourceType) {
			tflog.Warn(ctx, fmt.Sprintf("Conflict creating resource, retrying: %s", r.ResourceType))
			time.Sleep(SLEEP_BETWEEN_RETRIES_SEC)
			continue // retry
		}
		diag.AddError(
			"Error creating resource",
			fmt.Sprintf("Could not create %s: %s", r.ResourceType, err.Error()),
		)
		return nil, err
	} // for range MAX_RETRY

	// if we reach here, we exhausted retries on conflict (parallel mode only)
	diag.AddError(
		"Error creating resource",
		fmt.Sprintf("Could not create %s after %d retries due to conflicts.", r.ResourceType, MAX_RETRY),
	)
	return nil, fmt.Errorf("failed to create %s after %d conflict retries: %w", r.ResourceType, MAX_RETRY, lastErr)
}

// smcCreateResource sends a create request for the given
// sub-resource with the provided data
func (r *ResourceBase[T]) smcCreateSubResource(
	ctx context.Context, data *T, diag *diag.Diagnostics) (*smc.ResponseData, error) {

	parentUrl, err := getFromRefField(data)
	if err != nil {
		diag.AddError(
			"Missing from_ref for sub-resource creation",
			"The 'from_ref' attribute must be set to the parent resource's href when creating a sub-resource.",
		)
		return nil, err
	}

	parentUrl, err = common.ResolveRef(ctx, r.Config, parentUrl)
	if err != nil {
		diag.AddError(
			"Error resolving from_ref for sub-resource creation",
			fmt.Sprintf("Could not resolve from_ref '%s': %s", parentUrl, err.Error()),
		)
		return nil, err
	}

	apiData, err := r.Dispatch.ModelToJson(data)
	if err != nil {
		diag.AddError(
			"Error marshalling resource data",
			fmt.Sprintf("Could not marshal resource data to JSON: %s", err.Error()),
		)
		return nil, err
	}

	// get additional query params for this resource if any
	queryParams, err := r.Dispatch.GetCreateRequestParams(ctx, data)
	if err != nil {
		diag.AddError(
			"Error creating resource",
			fmt.Sprintf("Internal error: could not get request params for %s: %s",
				r.ResourceType, err.Error()),
		)
		return nil, err
	}

	var lastErr error
	for range MAX_RETRY {
		requestResp, err := smc.CreateSubResourceWithQueryParams(
			r.Config, apiData, parentUrl, queryParams)
		if err == nil {
			return requestResp, nil
		}
		lastErr = err
		// Only retry in parallel mode; serially, 409 is a legitimate user error.
		// 422 is retried only when the parent is locked ("locked by...").
		if isRetryableParallelCreateError(err) && smc.IsResourceTypeParallel(r.ResourceType) {
			tflog.Warn(ctx, fmt.Sprintf("Conflict creating sub-resource, retrying: %s", r.ResourceType))
			time.Sleep(SLEEP_BETWEEN_RETRIES_SEC)
			continue // retry
		}
		diag.AddError(
			"Error creating resource",
			fmt.Sprintf("Could not create %s: %s", r.ResourceType, err.Error()),
		)
		return nil, err
	} // for range MAX_RETRY

	// if we reach here, we exhausted retries on conflict (parallel mode only)
	diag.AddError(
		"Error creating resource",
		fmt.Sprintf("Could not create %s after %d retries due to conflicts.", r.ResourceType, MAX_RETRY),
	)
	return nil, fmt.Errorf("failed to create %s after %d conflict retries: %w", r.ResourceType, MAX_RETRY, lastErr)
}

func (r *ResourceBase[T]) updateTfState(ctx context.Context, data *T, id string,
	state *tfsdk.State, diag *diag.Diagnostics) (*smc.ResponseData, error) {

	// Normalize the href against the current base URL and API version so the
	// id persisted to state stays valid after an SMC upgrade or http->https
	// switch (SMC-66510, SMC-64735). SMC element ids are stable across
	// versions, so only the scheme/host/port and version prefix change.
	normalizedId, err := r.Config.Client.NormalizeHref(id)
	if err != nil {
		diag.AddError(
			"Error normalizing resource href",
			fmt.Sprintf("Could not normalize href '%s': %s", id, err.Error()),
		)
		return nil, err
	}
	id = normalizedId

	// Read the created resource to get full details
	readResp, err := r.Config.ReadResourceByHref(id)
	if err != nil {
		diag.AddError(
			"Error reading created resource",
			fmt.Sprintf("Could not read created %s: %s", r.ResourceType, err.Error()),
		)
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("Reading resource after create: %v", string(readResp.Body[:])))

	// Unmarshal the response into our model
	var smcData T
	if err := apijson.UnmarshalRoot(readResp.Body, &smcData); err != nil {
		diag.AddError(
			"Error parsing created resource",
			fmt.Sprintf("Could not parse created %s: %s", r.ResourceType, err.Error()),
		)
		return nil, err
	}

	MergeResourceModels(ctx /*src*/, &smcData /*dest*/, data)
	setId(data, id)

	diag.Append(state.Set(ctx, data)...)
	return readResp, nil
}

// normalizeStateHrefs rewrites the scheme/host/port/version of every
// "id" attribute and every href in every "link" map (the Lk field
// populated by CopyLinkField), at every nesting level (e.g. each
// physical_interface block has its own), against the client's current
// BaseUrl/APIVersion.
//
// Only needed on the drift-detection-enabled Read() path, after
// MergeResourceDrift: copyDriftField in model_drift.go deliberately
// never refreshes "id"/"link" from the freshly-read apiData (to avoid
// clobbering state with an SMC response that can omit them), so a
// stale scheme/host/port baked into an href from before an SMC upgrade
// or http->https switch survives that merge untouched. Since drift
// detection also gates whether "id"/"link" get a UseStateForUnknown
// plan modifier at all (schema_objects.j2), leaving them stale here
// would carry that stale value into the next plan, which then
// mismatches what Update's own (unconditional) NormalizeHref actually
// produces - "Provider produced inconsistent result after apply". See
// NormalizeHref, SMC-64735, SMC-66510. Pure string rewrite, no network
// call.
//
// Reuses mergeStructs' traversal (the same one copyComputedField relies
// on for create/update) with src and dest both pointing at state, so
// nested objects/lists are reached the same way create/update already
// does, without hand-rolling a second walker.
//
// Returns whether anything was actually rewritten, so the caller can
// skip persisting state back when there's nothing to do.
func normalizeStateHrefs(ctx context.Context, client *smc.SmcClient, state any) (bool, error) {
	changed := false

	copyFunc := func(ctx context.Context, params CopyFieldParams) (bool, error) {
		if params.fieldType.Name == "Lk" {
			ok, err := normalizeLinkMapField(ctx, client, params.destField)
			if err != nil {
				return false, err
			}
			changed = changed || ok
			return true, nil
		}

		if params.fieldType.Tag.Get("tfsdk") == "id" {
			ok, err := normalizeIdField(client, params.destField)
			if err != nil {
				return false, err
			}
			changed = changed || ok
			return true, nil
		}

		if params.fieldType.Tag.Get("fpro") != "" || params.fieldType.Tag.Get("tfsdk") == "from_ref" {
			// no hrefs to normalize on any other fpro/from_ref field,
			// and recursing would tunnel into framework-internal
			// types (e.g. the raw "Link" NestedObjectList) that
			// aren't meant to be walked outside CopyLinkField.
			return true, nil
		}

		// not handled here: let mergeStructs' default traversal
		// recurse into nested objects/lists so their own "id"/"link"
		// attributes are reached too.
		return false, nil
	}

	v := reflect.ValueOf(state)
	if err := mergeStructs(ctx, v, v, copyFunc); err != nil {
		return false, err
	}
	return changed, nil
}

// normalizeIdField normalizes a single "id" types.String field in
// place. Null/unknown/empty values are left alone.
func normalizeIdField(client *smc.SmcClient, field reflect.Value) (bool, error) {
	if !field.IsValid() || !field.CanSet() {
		return false, nil
	}
	id, ok := field.Interface().(types.String)
	if !ok || id.IsNull() || id.IsUnknown() || id.ValueString() == "" {
		return false, nil
	}

	normalized, err := client.NormalizeHref(id.ValueString())
	if err != nil {
		return false, err
	}
	if normalized == id.ValueString() {
		return false, nil
	}
	field.Set(reflect.ValueOf(types.StringValue(normalized)))
	return true, nil
}

// normalizeLinkMapField normalizes every href in a single "link" map
// (the Lk field) in place. Missing/null/unknown maps are left alone.
func normalizeLinkMapField(ctx context.Context, client *smc.SmcClient, field reflect.Value) (bool, error) {
	if !field.IsValid() || !field.CanSet() {
		return false, nil
	}

	lk, ok := field.Interface().(customfield.Map[types.String])
	if !ok || lk.IsNull() || lk.IsUnknown() {
		return false, nil
	}

	values, diags := lk.Value(ctx)
	if diags.HasError() {
		return false, fmt.Errorf("failed to read link map: %v", diags)
	}

	changed := false
	normalized := make(map[string]types.String, len(values))
	for rel, href := range values {
		if href.IsNull() || href.IsUnknown() {
			normalized[rel] = href
			continue
		}
		normalizedHref, err := client.NormalizeHref(href.ValueString())
		if err != nil {
			return false, err
		}
		if normalizedHref != href.ValueString() {
			changed = true
		}
		normalized[rel] = types.StringValue(normalizedHref)
	}
	if !changed {
		return false, nil
	}

	newMap, diags := customfield.NewMap[types.String](ctx, normalized)
	if diags.HasError() {
		return false, fmt.Errorf("failed to rebuild link map: %v", diags)
	}
	field.Set(reflect.ValueOf(newMap))
	return true, nil
}

func (r *ResourceBase[T]) smcDeleteResource(
	ctx context.Context, href string, diag *diag.Diagnostics) {

	// used by specific resources to perform actions, eg to close
	// and save vpn
	r.Dispatch.BeforeDelete(ctx, nil, href)

	for range MAX_RETRY {
		// Read the resource to get the ETag for deletion.

		// todo: optimize pass the last known etag and check for 304
		// Not Modified
		readResp, err := r.Config.ReadResourceByHref(href)
		if err != nil {
			var statusErr *smc.StatusError
			if errors.As(err, &statusErr) && statusErr.Code == http.StatusNotFound {
				// got a 404 Not Found - resource already deleted. Ignore error
				tflog.Warn(ctx, fmt.Sprintf("Resource '%s' not found before deletion, assuming already deleted", href))
				return
			}

			diag.AddError(
				"Error reading resource before deletion",
				fmt.Sprintf("Could not read resource: %s", err.Error()),
			)
			return
		}

		err = smc.DeleteResourceByHref(r.Config, href, (*readResp).ETag)
		if err == nil {
			tflog.Debug(ctx, fmt.Sprintf("Deleted resource '%s' successfully", href))
			return
		}

		var statusErr *smc.StatusError
		if errors.As(err, &statusErr) {
			switch statusErr.Code {
			case http.StatusNotFound:
				// got a 404 Not Found - resource already
				// deleted. Ignore error
				return
			case http.StatusConflict:
				// got a 409 Conflict - reread the etag and retry
				tflog.Warn(ctx,
					fmt.Sprintf("Conflict delete resource, retrying: %s", href))

				diag.AddWarning(
					"Conflict detected while deleting resource",
					fmt.Sprintf("Could not delete resource: %s", err.Error()),
				)

				time.Sleep(SLEEP_BETWEEN_RETRIES_SEC)
				continue // retry

			case http.StatusBadRequest:
				// got a 400 Bad Request - likely due to dependencies
				if strings.Contains(statusErr.Message, "Impossible to delete the element") {
					diag.AddWarning("Resource has dependencies",
						fmt.Sprintf("Resource '%s' has dependencies and cannot be deleted now: %s",
							href, statusErr.Message))

					tflog.Warn(ctx,
						fmt.Sprintf("Resource '%s' has dependencies, cannot delete now", href))

					// note silently ignored if already pending delete
					err := datastore.ManagePendingDelete(ctx,
						datastore.AddPendingDelete,
						r.Config.Client.BaseUrl,
						r.Config.ResourceType, href)
					if err != nil {
						diag.AddError(
							"Error scheduling resource for pending deletion",
							fmt.Sprintf("Could not schedule resource '%s' for pending deletion: %s",
								href, err.Error()),
						)
					}
					return
				}
				// fallthrough to default handler
			default:
				// other status codes handled below
			}
		}

		diag.AddError(
			"Error deleting resource",
			fmt.Sprintf("Could not delete resource: %s", err.Error()),
		)
		return
	} // for range MAX_RETRY

	// if we reach here, we exhausted retries
	diag.AddError(
		"Error deleting resource",
		fmt.Sprintf("Could not delete resource '%s' after %d retries", href, MAX_RETRY),
	)
}
