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

package apijson

import (
	"reflect"

	"github.com/tidwall/gjson"
)

type UnionVariant struct {
	TypeFilter         gjson.Type
	DiscriminatorValue interface{}
	Type               reflect.Type
}

var unionRegistry = map[reflect.Type]unionEntry{}

type unionEntry struct {
	discriminatorKey string
	variants         []UnionVariant
}

func RegisterUnion(typ reflect.Type, discriminator string, variants ...UnionVariant) {
	unionRegistry[typ] = unionEntry{
		discriminatorKey: discriminator,
		variants:         variants,
	}
}
