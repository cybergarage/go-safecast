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
	"time"
)

// Equal checks if two values are equal.
// It returns true if both values are equal, otherwise false.
func Equal(v1 any, v2 any) bool {
	eq := func(v1, v2 any) bool {
		switch v1 := v1.(type) {
		case int:
			var cv2 int
			if err := ToInt(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case int8:
			var cv2 int8
			if err := ToInt8(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case int16:
			var cv2 int16
			if err := ToInt16(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case int32:
			var cv2 int32
			if err := ToInt32(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case int64:
			var cv2 int64
			if err := ToInt64(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case uint:
			var cv2 uint
			if err := ToUint(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case uint8:
			var cv2 uint8
			if err := ToUint8(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case uint16:
			var cv2 uint16
			if err := ToUint16(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case uint32:
			var cv2 uint32
			if err := ToUint32(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case uint64:
			var cv2 uint64
			if err := ToUint64(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case float32:
			var cv2 float32
			if err := ToFloat32(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case float64:
			var cv2 float64
			if err := ToFloat64(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case bool:
			var cv2 bool
			if err := ToBool(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case string:
			var cv2 string
			if err := ToString(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		case []byte:
			var cv2 []byte
			if err := ToBytes(v2, &cv2); err != nil {
				return false
			}
			if len(cv2) != len(v1) {
				return false
			}
			for i := range v1 {
				if cv2[i] != v1[i] {
					return false
				}
			}
		case time.Time:
			var cv2 time.Time
			if err := ToTime(v2, &cv2); err != nil {
				return false
			}
			return v1.Equal(cv2)
		case nil:
			if v2 == nil {
				return true
			}
			return false
		}

		return reflect.DeepEqual(v1, v2)
	}

	if eq(v1, v2) {
		return true
	}

	return eq(v2, v1)
}
