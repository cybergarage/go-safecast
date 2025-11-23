// Copyright (C) 2022 The go-safecast Authors All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package safecast

import (
	"reflect"
)

// Equal checks if two values are equal.
// It returns true if both values are equal, otherwise false.
func Equal(v1 any, v2 any) bool {
	if reflect.DeepEqual(v1, v2) {
		return true
	}

	equalSlices := func(v1, v2 []any) bool {
		if len(v1) != len(v2) {
			return false
		}
		for i := range v1 {
			if !Equal(v1[i], v2[i]) {
				return false
			}
		}
		return true
	}

	equalMaps := func(v1Map, v2Map map[any]any) bool {
		if len(v1Map) != len(v2Map) {
			return false
		}
		for key, v1Val := range v1Map {
			v2Val, ok := v2Map[key]
			if !ok {
				return false
			}
			if !Equal(v1Val, v2Val) {
				return false
			}
		}
		return true
	}

	convertMapStringToAny := func(m map[string]any) map[any]any {
		result := make(map[any]any, len(m))
		for k, v := range m {
			result[k] = v
		}
		return result
	}

	switch v1 := v1.(type) {
	case []any:
		if v2Slice, ok := v2.([]any); ok {
			return equalSlices(v1, v2Slice)
		}
		return false
	case map[string]any:
		switch v2 := v2.(type) {
		case map[string]any:
			return equalMaps(convertMapStringToAny(v1), convertMapStringToAny(v2))
		case map[any]any:
			return equalMaps(convertMapStringToAny(v1), v2)
		}
		return false
	case map[any]any:
		switch v2 := v2.(type) {
		case map[any]any:
			return equalMaps(v1, v2)
		case map[string]any:
			return equalMaps(v1, convertMapStringToAny(v2))
		}
		return false
	}

	cmp, err := Compare(v1, v2)
	if err == nil {
		switch cmp {
		case 0:
			return true
		}
	}

	return false
}
