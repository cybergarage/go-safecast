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
	"strings"
	"time"
)

// Compare checks if two values are equal.
func Compare(v1 any, v2 any) (int, error) {
	cmp := func(v1, v2 any) (int, error) {

		cmpInt := func(v1 *int, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 int
			if err := ToInt(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpInt8 := func(v1 *int8, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 int8
			if err := ToInt8(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpInt16 := func(v1 *int16, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 int16
			if err := ToInt16(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpInt32 := func(v1 *int32, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 int32
			if err := ToInt32(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpInt64 := func(v1 *int64, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 int64
			if err := ToInt64(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpUint := func(v1 *uint, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 uint
			if err := ToUint(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpUint8 := func(v1 *uint8, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 uint8
			if err := ToUint8(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpUint16 := func(v1 *uint16, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 uint16
			if err := ToUint16(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpUint32 := func(v1 *uint32, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 uint32
			if err := ToUint32(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpUint64 := func(v1 *uint64, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 uint64
			if err := ToUint64(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpFloat32 := func(v1 *float32, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 float32
			if err := ToFloat32(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpFloat64 := func(v1 *float64, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 float64
			if err := ToFloat64(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if cv2 < *v1 {
				return -1, nil
			}
			return 1, nil
		}

		cmpBool := func(v1 *bool, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 bool
			if err := ToBool(v2, &cv2); err != nil {
				return 0, err
			}
			if cv2 == *v1 {
				return 0, nil
			}
			if *v1 {
				return 1, nil
			}
			return -1, nil
		}

		cmpString := func(v1 *string, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 string
			if err := ToString(v2, &cv2); err != nil {
				return 0, err
			}
			return strings.Compare(*v1, cv2), nil
		}

		cmpBytes := func(v1 []byte, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 []byte
			if err := ToBytes(v2, &cv2); err != nil {
				return 0, err
			}
			return bytes.Compare(cv2, v1), nil
		}

		cmpTime := func(v1 *time.Time, v2 any) (int, error) {
			if v1 == nil {
				if v2 == nil {
					return 0, nil
				}
				return -1, nil
			}
			var cv2 time.Time
			if err := ToTime(v2, &cv2); err != nil {
				return 0, err
			}
			if v1.Equal(cv2) {
				return 0, nil
			}
			if v1.Before(cv2) {
				return -1, nil
			}
			return 1, nil
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
				return 0, nil
			}
			return 0, newCompareError(v1, v2)
		}

		return 0, newCompareError(v1, v2)
	}

	r, err := cmp(v1, v2)
	if err == nil {
		return r, err
	}

	return cmp(v2, v1)
}
