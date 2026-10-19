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
	"net/netip"
	"reflect"
	"strconv"
	"strings"

	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MergeResourceDrift overlays a freshly fetched API response (src) onto
// the current Terraform state (dest), for use by Read/drift-detection.
//
// Unlike MergeResourceModels (used after Create/Update, which only
// copies computed 'fpro' fields onto a trusted plan), this treats dest
// as the anchor: a field the user never set (null/unknown in dest) is
// left alone even if SMC reports a default for it, while a field the
// user did set is overwritten with SMC's current value so real drift
// is detected. Computed fields are always refreshed, same as today.
//
// Nested lists of identifiable objects (anything with GetSliceIds) are
// reconciled by id rather than by position: entries present in both are
// merged (state-anchored), entries only in src are additions, entries
// only in dest are removed. See reconcileSlice.
func MergeResourceDrift(ctx context.Context, src, dest any) error {
	return mergeStructs(ctx, reflect.ValueOf(src), reflect.ValueOf(dest), copyDriftField)
}

func copyDriftField(ctx context.Context, params CopyFieldParams) (bool, error) {
	fproTag := params.fieldType.Tag.Get("fpro")
	isIdField := params.fieldType.Tag.Get("tfsdk") == "id"

	if isIdField || fproTag != "" {
		// id and computed fields always come from the API, same as
		// the create/update merge path.
		return copyComputedField(ctx, params)
	}

	if params.fieldType.Name == "Lk" {
		// Lk is only ever populated as a side effect of the
		// corresponding "Link" field's fpro special handler; it has
		// no fpro tag of its own, so leave it alone here or the
		// generic scalar overlay below would overwrite it with
		// src.Lk, which is always null (json:"-").
		return true, nil
	}

	if params.fieldType.Tag.Get("tfsdk") == "from_ref" {
		// from_ref has no json tag (it's Terraform-side bookkeeping
		// for sub-resource creation, not an SMC API field), so a
		// fresh API decode always leaves it null; never overlay it
		// or it gets cleared on every refresh.
		return true, nil
	}

	switch params.srcField.Kind() {
	case reflect.Ptr:
		return copyDriftPointerField(ctx, params)
	case reflect.Slice:
		merged, err := reconcileOrReplaceSlice(ctx, params.srcField, params.destField)
		if err != nil {
			return false, err
		}
		params.destField.Set(merged)
		return true, nil
	case reflect.Struct:
		return copyDriftScalarField(params)
	default:
		return false, nil
	}
}

func copyDriftPointerField(ctx context.Context, params CopyFieldParams) (bool, error) {
	src, dest := params.srcField, params.destField

	if src.IsNil() {
		// SMC's response didn't include a value for this field this
		// time (confirmed in practice: booleans on some implicitly
		// created sub-resources, and a whole "ip" list on some
		// ip_address_list reads) - as with scalar attributes
		// (copyDriftScalarField), that's not strong enough evidence
		// of a deliberate removal to clobber a known state value with
		// it, and unlike a scalar, clearing dest here can never
		// recover on a later read once state.IsNil() is true.
		return true, nil
	}

	if src.Type().Elem().Kind() != reflect.Slice {
		// pointer to a nested object (SingleNestedBlock): if dest is
		// nil the user never configured it, and mergeStructs already
		// skips recursing in that case; otherwise let it recurse so
		// each field goes back through copyDriftField.
		return false, nil
	}

	if dest.IsNil() {
		// optional list the user never configured; don't import
		// whatever SMC currently has for it.
		return true, nil
	}

	merged, err := reconcileOrReplaceSlice(ctx, src.Elem(), dest.Elem())
	if err != nil {
		return false, err
	}
	dest.Elem().Set(merged)
	return true, nil
}

func copyDriftScalarField(params CopyFieldParams) (bool, error) {
	dest := params.destField
	if !dest.CanInterface() {
		return false, nil
	}

	destAttr, ok := dest.Interface().(attr.Value)
	if !ok {
		// not a terraform value wrapper (e.g. a plain nested struct);
		// let mergeStructs recurse into its fields.
		return false, nil
	}

	if destAttr.IsNull() || destAttr.IsUnknown() {
		// user never set this attribute; don't import SMC's default.
		return true, nil
	}

	src := params.srcField
	if !src.CanInterface() {
		return false, nil
	}

	srcAttr, ok := src.Interface().(attr.Value)
	if !ok {
		return false, nil
	}

	if srcAttr.IsNull() || srcAttr.IsUnknown() {
		// SMC's response didn't include a value for this attribute
		// this time (some fields are only conditionally rendered,
		// e.g. on implicitly-created sub-resources) - that's not
		// strong enough evidence of a deliberate removal to clobber
		// a known state value with it.
		return true, nil
	}

	if srcStr, ok := srcAttr.(types.String); ok {
		value := srcStr.ValueString()

		if isRedactedValue(value) {
			// SMC masks sensitive fields (e.g. password) in its GET
			// responses instead of echoing the real value back; that
			// placeholder is never real drift.
			return true, nil
		}

		if stripped := stripCompositeIdPrefix(params.destVal, value); stripped != value {
			// some fields come back as "{name}@#@#{value}" - the
			// name prefix is SMC's own internal composite-key
			// encoding, not part of the attribute's actual content.
			dest.Set(reflect.ValueOf(types.StringValue(stripped)))
			return true, nil
		}
	}

	if sameNormalizedIPValue(destAttr, srcAttr) {
		// same address/subnet, just formatted differently (e.g. a
		// leading zero) - not real drift.
		return true, nil
	}

	dest.Set(src)
	return true, nil
}

// isRedactedValue reports whether s looks like SMC's masked-value
// placeholder for a sensitive field (observed as a string of asterisk
// characters, e.g. "*****", in place of a real password).
func isRedactedValue(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c != '*' {
			return false
		}
	}
	return true
}

// stripCompositeIdPrefix removes a "{name}@#@#" prefix from s, where
// {name} matches the resource's own Name field, if present. Some SMC
// fields (e.g. Internal_User.unique_id) come back from the API with
// this composite encoding even though the attribute is configured and
// stored as the plain value. Returns s unchanged if the prefix isn't
// present or doesn't match the resource's Name.
func stripCompositeIdPrefix(destVal reflect.Value, s string) string {
	const separator = "@#@#"

	prefix, rest, found := strings.Cut(s, separator)
	if !found {
		return s
	}

	nameField := destVal.FieldByName("Name")
	if !nameField.IsValid() || !nameField.CanInterface() {
		return s
	}
	nameAttr, ok := nameField.Interface().(types.String)
	if !ok || nameAttr.IsNull() || nameAttr.ValueString() != prefix {
		return s
	}

	return rest
}

// reconcileOrReplaceSlice decides, per field, whether a slice holds
// plain terraform values (no per-element identity, e.g. a list of
// strings) or identifiable nested objects, and dispatches accordingly.
func reconcileOrReplaceSlice(ctx context.Context, src, dest reflect.Value) (reflect.Value, error) {
	elemType := src.Type().Elem()

	if _, ok := reflect.New(elemType).Elem().Interface().(attr.Value); ok {
		if dest.IsNil() {
			// user never configured this list; don't import SMC's value.
			return dest, nil
		}
		if sameElementSet(dest, src) {
			// same members, just returned in a different order (SMC
			// does not guarantee element order for these lists) - not
			// real drift, keep dest's existing order.
			return dest, nil
		}
		return src, nil
	}

	return reconcileSlice(ctx, src, dest)
}

// sameElementSet reports whether two slices of plain terraform values
// hold the same elements, ignoring order (but not count - a repeated
// value must appear the same number of times in both).
func sameElementSet(destSlice, srcSlice reflect.Value) bool {
	if destSlice.Len() != srcSlice.Len() {
		return false
	}

	consumed := make([]bool, srcSlice.Len())
	for i := 0; i < destSlice.Len(); i++ {
		destElem, ok := destSlice.Index(i).Interface().(attr.Value)
		if !ok {
			return false
		}

		matched := false
		for j := 0; j < srcSlice.Len(); j++ {
			if consumed[j] {
				continue
			}
			srcElem, ok := srcSlice.Index(j).Interface().(attr.Value)
			if !ok {
				return false
			}
			if destElem.Equal(srcElem) {
				consumed[j] = true
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

// reconcileSlice reconciles list membership by id (GetSliceIds), then
// state-anchor-merges matched pairs and fully imports src-only
// (added) entries. dest-only (removed) entries are dropped.
//
// The generated ListNestedBlock diffs config against (refreshed) state
// by position, and Read cannot see config (only prior state) - so the
// only way to avoid every already-matched entry showing as a false
// "modified" diff is to keep matched entries in dest's existing
// relative order (normally identical to config's order when nothing
// has drifted) and append genuinely new entries after them. Sorting
// into src's/API's order instead would shift every entry after an
// insertion and make them all look changed.
//
// A removal from the middle of the list will still show cascading
// "modified" diffs on later entries - that is an inherent limitation
// of position-based ListNestedBlock diffing, not something reordering
// here can fix.
//
// If any element (src or dest) has no usable id, reconciliation is not
// safe (a real entry could look like a simultaneous add+remove), so the
// whole list is left untouched for this refresh rather than guessed at.
func reconcileSlice(ctx context.Context, src, dest reflect.Value) (reflect.Value, error) {
	srcIds := make([][]string, src.Len())
	for j := 0; j < src.Len(); j++ {
		ids, err := GetElementIds(ctx, src.Index(j))
		if err != nil {
			tflog.Warn(ctx, "reconcileSlice: API element has no reliable identity, leaving list untouched", map[string]any{
				"type": elemType(src).Name(),
			})
			return dest, nil
		}
		srcIds[j] = ids
	}

	destIds := make([][]string, dest.Len())
	for i := 0; i < dest.Len(); i++ {
		ids, err := GetElementIds(ctx, dest.Index(i))
		if err != nil {
			tflog.Warn(ctx, "reconcileSlice: state element has no reliable identity, leaving list untouched", map[string]any{
				"type": elemType(dest).Name(),
			})
			return dest, nil
		}
		destIds[i] = ids
	}

	idToSrc := make(map[string]int, src.Len())
	for j, ids := range srcIds {
		for _, id := range ids {
			if _, dup := idToSrc[id]; dup {
				tflog.Warn(ctx, "reconcileSlice: duplicate id among API elements", map[string]any{"id": id})
				continue
			}
			idToSrc[id] = j
		}
	}

	srcForDest := make(map[int]int, dest.Len())
	matchedSrc := make(map[int]bool, src.Len())
	for i, ids := range destIds {
		for _, id := range ids {
			j, ok := idToSrc[id]
			if !ok || matchedSrc[j] {
				continue
			}
			srcForDest[i] = j
			matchedSrc[j] = true
			break
		}
	}

	result := reflect.MakeSlice(src.Type(), 0, dest.Len()+src.Len())
	for i := 0; i < dest.Len(); i++ {
		j, ok := srcForDest[i]
		if !ok {
			// removed: API no longer reports this entry
			continue
		}
		merged := reflect.New(src.Type().Elem()).Elem()
		merged.Set(dest.Index(i))
		if err := mergeStructs(ctx, src.Index(j), merged, copyDriftField); err != nil {
			return reflect.Value{}, err
		}
		result = reflect.Append(result, merged)
	}

	for j := 0; j < src.Len(); j++ {
		if matchedSrc[j] {
			continue
		}
		added := reflect.New(src.Type().Elem()).Elem()
		added.Set(src.Index(j))
		if err := mergeStructs(ctx, src.Index(j), added, copyComputedField); err != nil {
			return reflect.Value{}, err
		}
		result = reflect.Append(result, added)
	}

	return result, nil
}

func elemType(slice reflect.Value) reflect.Type {
	return slice.Type().Elem()
}

// sameNormalizedIPValue reports whether two string attribute values
// represent the same IP address or CIDR prefix, even if formatted
// differently (e.g. SMC returning "192.168.100.00/24" for what
// Terraform has as "192.168.100.0/24"). Values that aren't both
// parseable the same way (e.g. hostnames, plain strings) are not
// considered equal here - the caller falls back to a plain overlay.
func sameNormalizedIPValue(destAttr, srcAttr attr.Value) bool {
	destStr, ok := destAttr.(types.String)
	if !ok {
		return false
	}
	srcStr, ok := srcAttr.(types.String)
	if !ok {
		return false
	}
	a, b := destStr.ValueString(), srcStr.ValueString()

	if pa, err := netip.ParsePrefix(a); err == nil {
		if pb, err := netip.ParsePrefix(b); err == nil {
			return pa.String() == pb.String()
		}
	}
	if aa, err := netip.ParseAddr(a); err == nil {
		if ab, err := netip.ParseAddr(b); err == nil {
			return aa.String() == ab.String()
		}
	}

	// net/netip deliberately rejects a leading zero in an IPv4 octet
	// (RFC 6943 hardening against octal ambiguity), but SMC has been
	// observed to return exactly that (e.g. "192.168.100.00/24"), so
	// fall back to a lenient dotted-decimal comparison for that case.
	na, ok := normalizeLenientIPv4(a)
	if !ok {
		return false
	}
	nb, ok := normalizeLenientIPv4(b)
	return ok && na == nb
}

// normalizeLenientIPv4 normalizes a dotted-decimal IPv4 address,
// optionally with a "/n" prefix length suffix, tolerating a leading
// zero within an octet. Returns ("", false) if s isn't a plausible
// IPv4 address/prefix at all.
func normalizeLenientIPv4(s string) (string, bool) {
	addr, mask, hasMask := strings.Cut(s, "/")
	octets := strings.Split(addr, ".")
	if len(octets) != 4 {
		return "", false
	}
	for i, o := range octets {
		if o == "" || len(o) > 3 {
			return "", false
		}
		n, err := strconv.Atoi(o)
		if err != nil || n < 0 || n > 255 {
			return "", false
		}
		octets[i] = strconv.Itoa(n)
	}

	normalized := strings.Join(octets, ".")
	if hasMask {
		normalized += "/" + mask
	}
	return normalized, true
}
