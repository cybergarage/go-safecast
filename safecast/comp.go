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

// Compare checks if two values are equal.
func Compare(v1 any, v2 any) int {
	cmp := func(v1, v2 any) int {

		cmpInt := func(v1 *int, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 int
			if err := ToInt(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpInt8 := func(v1 *int8, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 int8
			if err := ToInt8(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpInt16 := func(v1 *int16, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 int16
			if err := ToInt16(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return cv2 == *v1
		}

		cmpInt32 := func(v1 *int32, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 int32
			if err := ToInt32(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpInt64 := func(v1 *int64, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 int64
			if err := ToInt64(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpUint := func(v1 *uint, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 uint
			if err := ToUint(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpUint8 := func(v1 *uint8, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 uint8
			if err := ToUint8(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpUint16 := func(v1 *uint16, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 uint16
			if err := ToUint16(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpUint32 := func(v1 *uint32, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 uint32
			if err := ToUint32(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpUint64 := func(v1 *uint64, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 uint64
			if err := ToUint64(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpFloat32 := func(v1 *float32, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 float32
			if err := ToFloat32(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpFloat64 := func(v1 *float64, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 float64
			if err := ToFloat64(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpBool := func(v1 *bool, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 bool
			if err := ToBool(v2, &cv2); err != nil {
				return -1
			}
			if cv2 == *v1 {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpString := func(v1 *string, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 string
			if err := ToString(v2, &cv2); err != nil {
				return -1
			}
			if strings.EqualFold(cv2, *v1) {
				return 0
			}
			if cv2 < *v1 {
				return -1
			}
			return 1
		}

		cmpBytes := func(v1 []byte, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 []byte
			if err := ToBytes(v2, &cv2); err != nil {
				return -1
			}
			if bytes.Equal(cv2, v1) {
				return 0
			}
			if bytes.Compare(cv2, v1) < 0 {
				return -1
			}
			return 1
		}

		cmpTime := func(v1 *time.Time, v2 any) int {
			if v1 == nil {
				if v2 == nil {
					return 0
				}
				return -1
			}
			var cv2 time.Time
			if err := ToTime(v2, &cv2); err != nil {
				return -1
			}
			if v1.Equal(cv2) {
				return 0
			}
			if v1.Before(cv2) {
				return -1
			}
			return 1
		}

		switch v1 := v1.(type) {
		case int:
			return cmpInt(&v1, v2)
		case *int:
			return cmpInt(v1, v2)
		case int8:
			return cmpInt8(&v1, v2)
		case *int8:
			return cmpInt8(v1, v2)
		case int16:
			return cmpInt16(&v1, v2)
		case *int16:
			return cmpInt16(v1, v2)
		case int32:
			return cmpInt32(&v1, v2)
		case *int32:
			return cmpInt32(v1, v2)
		case int64:
			return cmpInt64(&v1, v2)
		case *int64:
			return cmpInt64(v1, v2)
		case uint:
			return cmpUint(&v1, v2)
		case *uint:
			return cmpUint(v1, v2)
		case uint8:
			return cmpUint8(&v1, v2)
		case *uint8:
			return cmpUint8(v1, v2)
		case uint16:
			return cmpUint16(&v1, v2)
		case *uint16:
			return cmpUint16(v1, v2)
		case uint32:
			return cmpUint32(&v1, v2)
		case *uint32:
			return cmpUint32(v1, v2)
		case uint64:
			return cmpUint64(&v1, v2)
		case *uint64:
			return cmpUint64(v1, v2)
		case float32:
			return cmpFloat32(&v1, v2)
		case *float32:
			return cmpFloat32(v1, v2)
		case float64:
			return cmpFloat64(&v1, v2)
		case *float64:
			return cmpFloat64(v1, v2)
		case bool:
			return cmpBool(&v1, v2)
		case *bool:
			return cmpBool(v1, v2)
		case string:
			return cmpString(&v1, v2)
		case *string:
			return cmpString(v1, v2)
		case []byte:
			return cmpBytes(v1, v2)
		case time.Time:
			return cmpTime(&v1, v2)
		case *time.Time:
			return cmpTime(v1, v2)
		case nil:
			if v2 == nil {
				return 0
			}
			return -1
		}

		return reflect.DeepEqual(v1, v2)
	}

	return cmp(v2, v1)
}
