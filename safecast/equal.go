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
	"bytes"
	"reflect"
	"strings"
	"time"
)

// Equal checks if two values are equal.
// It returns true if both values are equal, otherwise false.
func Equal(v1 any, v2 any) bool {
	eq := func(v1, v2 any) bool {

		eqInt := func(v1 int, v2 any) bool {
			var cv2 int
			if err := ToInt(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqInt8 := func(v1 int8, v2 any) bool {
			var cv2 int8
			if err := ToInt8(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqInt16 := func(v1 int16, v2 any) bool {
			var cv2 int16
			if err := ToInt16(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqInt32 := func(v1 int32, v2 any) bool {
			var cv2 int32
			if err := ToInt32(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqInt64 := func(v1 int64, v2 any) bool {
			var cv2 int64
			if err := ToInt64(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqUint := func(v1 uint, v2 any) bool {
			var cv2 uint
			if err := ToUint(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}
		eqUint8 := func(v1 uint8, v2 any) bool {
			var cv2 uint8
			if err := ToUint8(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqUint16 := func(v1 uint16, v2 any) bool {
			var cv2 uint16
			if err := ToUint16(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqUint32 := func(v1 uint32, v2 any) bool {
			var cv2 uint32
			if err := ToUint32(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqUint64 := func(v1 uint64, v2 any) bool {
			var cv2 uint64
			if err := ToUint64(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqFloat32 := func(v1 float32, v2 any) bool {
			var cv2 float32
			if err := ToFloat32(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqFloat64 := func(v1 float64, v2 any) bool {
			var cv2 float64
			if err := ToFloat64(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqBool := func(v1 bool, v2 any) bool {
			var cv2 bool
			if err := ToBool(v2, &cv2); err != nil {
				return false
			}
			return cv2 == v1
		}

		eqString := func(v1 string, v2 any) bool {
			var cv2 string
			if err := ToString(v2, &cv2); err != nil {
				return false
			}
			return strings.EqualFold(cv2, v1)
		}

		eqBytes := func(v1 []byte, v2 any) bool {
			var cv2 []byte
			if err := ToBytes(v2, &cv2); err != nil {
				return false
			}
			return bytes.Equal(cv2, v1)
		}

		eqTime := func(v1 time.Time, v2 any) bool {
			var cv2 time.Time
			if err := ToTime(v2, &cv2); err != nil {
				return false
			}
			return v1.Equal(cv2)
		}

		switch v1 := v1.(type) {
		case int:
			return eqInt(v1, v2)
		case *int:
			if v1 == nil {
				return v2 == nil
			}
			return eqInt(*v1, v2)
		case int8:
			return eqInt8(v1, v2)
		case *int8:
			if v1 == nil {
				return v2 == nil
			}
			return eqInt8(*v1, v2)
		case int16:
			return eqInt16(v1, v2)
		case *int16:
			if v1 == nil {
				return v2 == nil
			}
			return eqInt16(*v1, v2)
		case int32:
			return eqInt32(v1, v2)
		case *int32:
			if v1 == nil {
				return v2 == nil
			}
			return eqInt32(*v1, v2)
		case int64:
			return eqInt64(v1, v2)
		case *int64:
			if v1 == nil {
				return v2 == nil
			}
			return eqInt64(*v1, v2)
		case uint:
			return eqUint(v1, v2)
		case *uint:
			if v1 == nil {
				return v2 == nil
			}
			return eqUint(*v1, v2)
		case uint8:
			return eqUint8(v1, v2)
		case *uint8:
			if v1 == nil {
				return v2 == nil
			}
			return eqUint8(*v1, v2)
		case uint16:
			return eqUint16(v1, v2)
		case *uint16:
			if v1 == nil {
				return v2 == nil
			}
			return eqUint16(*v1, v2)
		case uint32:
			return eqUint32(v1, v2)
		case *uint32:
			if v1 == nil {
				return v2 == nil
			}
			return eqUint32(*v1, v2)
		case uint64:
			return eqUint64(v1, v2)
		case *uint64:
			if v1 == nil {
				return v2 == nil
			}
			return eqUint64(*v1, v2)
		case float32:
			return eqFloat32(v1, v2)
		case *float32:
			if v1 == nil {
				return v2 == nil
			}
			return eqFloat32(*v1, v2)
		case float64:
			return eqFloat64(v1, v2)
		case *float64:
			if v1 == nil {
				return v2 == nil
			}
			return eqFloat64(*v1, v2)
		case bool:
			return eqBool(v1, v2)
		case *bool:
			if v1 == nil {
				return v2 == nil
			}
			return eqBool(*v1, v2)
		case string:
			return eqString(v1, v2)
		case *string:
			if v1 == nil {
				return v2 == nil
			}
			return eqString(*v1, v2)
		case []byte:
			if v1 == nil {
				return v2 == nil
			}
			return eqBytes(v1, v2)
		case time.Time:
			return eqTime(v1, v2)
		case *time.Time:
			if v1 == nil {
				return v2 == nil
			}
			return eqTime(*v1, v2)
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
