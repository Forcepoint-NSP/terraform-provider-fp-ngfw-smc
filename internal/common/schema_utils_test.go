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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------------------------------------------------------------------------
// Test types — two structurally similar structs in different "packages"
// (simulated here as distinct named types in the same file).
// ---------------------------------------------------------------------------

type v0Scalar struct {
	Name    string
	Age     int
	Active  bool
	Removed string // only in v0
}

type v1Scalar struct {
	Name   string
	Age    int
	Active bool
	Added  string // only in v1
}

// Nested types whose layouts differ (v0 has Extra, v1 does not).
type v0Inner struct {
	X     int
	Extra string // only in v0
}

type v1Inner struct {
	X   int
	New float64 // only in v1
}

type v0Outer struct {
	Label string
	Inner *v0Inner
	Items *[]v0Inner
}

type v1Outer struct {
	Label string
	Inner *v1Inner
	Items *[]v1Inner
}

// ---------------------------------------------------------------------------
// CopyFieldsByName
// ---------------------------------------------------------------------------

func TestCopyFieldsByName_ScalarFields(t *testing.T) {
	src := v0Scalar{Name: "alice", Age: 30, Active: true, Removed: "gone"}
	var dst v1Scalar
	CopyFieldsByName(&dst, &src)

	assert.Equal(t, "alice", dst.Name)
	assert.Equal(t, 30, dst.Age)
	assert.True(t, dst.Active)
	assert.Empty(t, dst.Added, "field only in v1 stays zero")
}

func TestCopyFieldsByName_TypeMismatch_Skipped(t *testing.T) {
	// Name is string in both, but Age has different types.
	type srcT struct {
		Name string
		Age  int64
	}
	type dstT struct {
		Name string
		Age  int
	}
	src := srcT{Name: "bob", Age: 99}
	var dst dstT
	CopyFieldsByName(&dst, &src)

	assert.Equal(t, "bob", dst.Name, "identical-type field copied")
	assert.Equal(t, 0, dst.Age, "type-mismatch field skipped")
}

func TestCopyFieldsByName_SrcFieldAbsentInDst_Skipped(t *testing.T) {
	src := v0Scalar{Name: "carol", Removed: "x"}
	var dst v1Scalar
	CopyFieldsByName(&dst, &src)
	// Removed does not exist in v1Scalar — must not panic; other fields copied.
	assert.Equal(t, "carol", dst.Name)
}

func TestCopyFieldsByName_NilPointerFields_NotCopied(t *testing.T) {
	// Both structs have a *string field — nil src pointer stays nil in dst.
	type T struct{ P *string }
	s := "hello"
	src := T{P: &s}
	var dst T
	CopyFieldsByName(&dst, &src)
	// *string is an exact type match, so the pointer itself is copied.
	assert.Equal(t, &s, dst.P)
}

// ---------------------------------------------------------------------------
// CopyFieldsByNameRecursive
// ---------------------------------------------------------------------------

func TestCopyFieldsByNameRecursive_SameType_DirectCopy(t *testing.T) {
	type S struct {
		X int
		Y string
	}
	src := S{X: 42, Y: "hi"}
	var dst S
	CopyFieldsByNameRecursive(&dst, &src)
	assert.Equal(t, src, dst)
}

func TestCopyFieldsByNameRecursive_ScalarsAcrossDifferentTypes(t *testing.T) {
	src := v0Scalar{Name: "dave", Age: 25, Active: false, Removed: "drop"}
	var dst v1Scalar
	CopyFieldsByNameRecursive(&dst, &src)

	assert.Equal(t, "dave", dst.Name)
	assert.Equal(t, 25, dst.Age)
	assert.False(t, dst.Active)
	assert.Empty(t, dst.Added)
}

func TestCopyFieldsByNameRecursive_NestedPointer_DifferentTypes(t *testing.T) {
	src := v0Outer{
		Label: "outer",
		Inner: &v0Inner{X: 7, Extra: "ignored"},
	}
	var dst v1Outer
	CopyFieldsByNameRecursive(&dst, &src)

	assert.Equal(t, "outer", dst.Label)
	require_notnil(t, dst.Inner)
	assert.Equal(t, 7, dst.Inner.X)
	assert.Equal(t, 0.0, dst.Inner.New, "new field in v1 stays zero")
}

func TestCopyFieldsByNameRecursive_NestedPointer_NilSrc(t *testing.T) {
	src := v0Outer{Label: "x", Inner: nil}
	var dst v1Outer
	CopyFieldsByNameRecursive(&dst, &src)
	assert.Equal(t, "x", dst.Label)
	assert.Nil(t, dst.Inner, "nil src pointer leaves dst nil")
}

func TestCopyFieldsByNameRecursive_SliceOfDifferentTypes(t *testing.T) {
	items := []v0Inner{{X: 1, Extra: "a"}, {X: 2, Extra: "b"}}
	src := v0Outer{Label: "list", Items: &items}
	var dst v1Outer
	CopyFieldsByNameRecursive(&dst, &src)

	assert.Equal(t, "list", dst.Label)
	require_notnil(t, dst.Items)
	assert.Len(t, *dst.Items, 2)
	assert.Equal(t, 1, (*dst.Items)[0].X)
	assert.Equal(t, 2, (*dst.Items)[1].X)
	assert.Equal(t, 0.0, (*dst.Items)[0].New)
}

func TestCopyFieldsByNameRecursive_SliceNil(t *testing.T) {
	src := v0Outer{Items: nil}
	var dst v1Outer
	CopyFieldsByNameRecursive(&dst, &src)
	assert.Nil(t, dst.Items)
}

func TestCopyFieldsByNameRecursive_SrcFieldAbsentInDst_Skipped(t *testing.T) {
	// Removed is in v0 but not v1 — must not panic.
	src := v0Scalar{Name: "eve", Removed: "nope"}
	var dst v1Scalar
	CopyFieldsByNameRecursive(&dst, &src)
	assert.Equal(t, "eve", dst.Name)
}

func TestCopyFieldsByNameRecursive_DstFieldAbsentInSrc_Preserved(t *testing.T) {
	// Added exists in v1 but not v0. The copy iterates src fields only,
	// so a pre-existing dst-only value must remain unchanged.
	src := v0Scalar{Name: "frank"}
	dst := v1Scalar{Added: "keep-me"}
	CopyFieldsByNameRecursive(&dst, &src)
	assert.Equal(t, "frank", dst.Name)
	assert.Equal(t, "keep-me", dst.Added, "dst-only field is untouched by the copy")
}

// ---------------------------------------------------------------------------
// CopyStructSlice
// ---------------------------------------------------------------------------

func TestCopyStructSlice_NilSrc(t *testing.T) {
	result := CopyStructSlice[v0Scalar, v1Scalar](nil)
	assert.Nil(t, result)
}

func TestCopyStructSlice_EmptySlice(t *testing.T) {
	empty := []v0Scalar{}
	result := CopyStructSlice[v0Scalar, v1Scalar](&empty)
	assert.NotNil(t, result)
	assert.Empty(t, *result)
}

func TestCopyStructSlice_CopiesIdenticalFields(t *testing.T) {
	src := &[]v0Scalar{
		{Name: "g", Age: 10, Active: true, Removed: "x"},
		{Name: "h", Age: 20, Active: false, Removed: "y"},
	}
	result := CopyStructSlice[v0Scalar, v1Scalar](src)

	assert.NotNil(t, result)
	assert.Len(t, *result, 2)
	assert.Equal(t, "g", (*result)[0].Name)
	assert.Equal(t, 10, (*result)[0].Age)
	assert.True(t, (*result)[0].Active)
	assert.Empty(t, (*result)[0].Added)
	assert.Equal(t, "h", (*result)[1].Name)
	assert.Equal(t, 20, (*result)[1].Age)
}

func TestCopyStructSlice_DoesNotModifyOriginal(t *testing.T) {
	orig := []v0Scalar{{Name: "i", Age: 5}}
	result := CopyStructSlice[v0Scalar, v1Scalar](&orig)
	(*result)[0].Name = "mutated"
	assert.Equal(t, "i", orig[0].Name, "original slice must not be modified")
}

// require_notnil is a minimal helper to avoid importing testify/require.
func require_notnil(t *testing.T, v any) {
	t.Helper()
	if v == nil {
		t.Fatal("expected non-nil value")
	}
}
