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
	"fmt"
	"strconv"
)

// FromFloat64 casts an interface to an float64 type.
func FromFloat64(from float64, to any) error {
	switch to := to.(type) {
	case *int:
		*to = int(from)
	case *int8:
		*to = int8(from)
	case *int16:
		*to = int16(from)
	case *int32:
		*to = int32(from)
	case *int64:
		*to = int64(from)
	case *uint:
		*to = uint(from)
	case *uint8:
		*to = uint8(from)
	case *uint16:
		*to = uint16(from)
	case *uint32:
		*to = uint32(from)
	case *uint64:
		*to = uint64(from)
	case *float32:
		*to = float32(from)
	case *float64:
		*to = from
	case *string:
		*to = fmt.Sprintf("%v", from)
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// FromFloat32 casts an interface to an float32 type.
func FromFloat32(from float32, to any) error {
	return FromFloat64(float64(from), to)
}

// ToFloat64 casts an interface to an float64 type.
func ToFloat64(from any, to *float64) error {
	switch from := from.(type) {
	case int:
		*to = float64(from)
	case int8:
		*to = float64(from)
	case int16:
		*to = float64(from)
	case int32:
		*to = float64(from)
	case int64:
		*to = float64(from)
	case uint:
		*to = float64(from)
	case uint8:
		*to = float64(from)
	case uint16:
		*to = float64(from)
	case uint32:
		*to = float64(from)
	case uint64:
		*to = float64(from)
	case float32:
		*to = float64(from)
	case float64:
		*to = from
	case string:
		f, err := strconv.ParseFloat(from, 64)
		if err != nil {
			return newErrorCast(from, to)
		}
		*to = f
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToFloat32 casts an interface to an float64 type.
func ToFloat32(from any, to *float32) error {
	switch from := from.(type) {
	case int:
		*to = float32(from)
	case int8:
		*to = float32(from)
	case int16:
		*to = float32(from)
	case int32:
		*to = float32(from)
	case int64:
		*to = float32(from)
	case uint:
		*to = float32(from)
	case uint8:
		*to = float32(from)
	case uint16:
		*to = float32(from)
	case uint32:
		*to = float32(from)
	case uint64:
		*to = float32(from)
	case float32:
		*to = from
	case float64:
		*to = float32(from)
	case string:
		f, err := strconv.ParseFloat(from, 32)
		if err != nil {
			return newErrorCast(from, to)
		}
		*to = float32(f)
	default:
		return newErrorCast(from, to)
	}
	return nil
}
