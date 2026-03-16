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

		err = MergeResourceModels(ctx /*src*/, stateData /*dest*/, data,
			"key", "tag", "parent_insert_point", "parent_policy")
		if err != nil {
			// todo diag.AddError
			return nil, fmt.Errorf("failed to merge resource models for update: %w", err)
		}

		apiData, err := r.Dispatch.ModelToJson(data)
		if err != nil {
			diag.AddError(
				"Error marshalling resource data",
				fmt.Sprintf("Could not marshal resource data to JSON: %s", err.Error()),
			)
			return nil, err
		}

		resp, err := smc.UpdateResource(r.Config, apiData, resolved_id, readResp.ETag)
		if err != nil {
			var statusErr *smc.StatusError
			if errors.As(err, &statusErr) && statusErr.Code == http.StatusConflict {
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

	requestResp, err := smc.CreateResource(r.Config, apiData)
	if err != nil {
		diag.AddError(
			"Error creating resource",
			fmt.Sprintf("Could not create %s: %s", r.ResourceType, err.Error()),
		)
		return nil, err
	}
	return requestResp, nil
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

	requestResp, err := smc.CreateSubResourceWithQueryParams(
		r.Config, apiData, parentUrl, queryParams)
	if err != nil {
		diag.AddError(
			"Error creating resource",
			fmt.Sprintf("Could not create %s: %s", r.ResourceType, err.Error()),
		)
		return nil, err
	}
	return requestResp, nil
}

func (r *ResourceBase[T]) updateTfState(ctx context.Context, data *T, id string,
	state *tfsdk.State, diag *diag.Diagnostics) (*smc.ResponseData, error) {

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
