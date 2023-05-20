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

// FromBool casts an interface to a bool type.
func FromBool(from bool, to any) error {
	toint := func(v bool) int {
		if v {
			return 1
		}
		return 0
	}

	switch to := to.(type) {
	case *int:
		*to = toint(from)
	case *int8:
		*to = int8(toint(from))
	case *int16:
		*to = int16(toint(from))
	case *int32:
		*to = int32(toint(from))
	case *int64:
		*to = int64(toint(from))
	case *uint:
		*to = uint(toint(from))
	case *uint8:
		*to = uint8(toint(from))
	case *uint16:
		*to = uint16(toint(from))
	case *uint32:
		*to = uint32(toint(from))
	case *uint64:
		*to = uint64(toint(from))
	case *bool:
		*to = from
	case *string:
		*to = fmt.Sprintf("%v", from)
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToBool casts an interface to a bool type.
func ToBool(from any, to *bool) error {
	tobool := func(v any) bool {
		switch v := v.(type) {
		case int:
			if 0 < v {
				return true
			}
		case uint:
			if 0 < v {
				return true
			}
		}
		return false
	}

	switch from := from.(type) {
	case int:
		*to = tobool(from)
	case int8:
		*to = tobool(int(from))
	case int16:
		*to = tobool(int(from))
	case int32:
		*to = tobool(int(from))
	case int64:
		*to = tobool(int(from))
	case uint:
		*to = tobool(from)
	case uint8:
		*to = tobool(uint(from))
	case uint16:
		*to = tobool(uint(from))
	case uint32:
		*to = tobool(uint(from))
	case uint64:
		*to = tobool(uint(from))
	case bool:
		*to = from
	case string:
		b, err := strconv.ParseBool(from)
		if err != nil {
			return newErrorCast(from, to)
		}
		*to = b
	default:
		return newErrorCast(from, to)
	}
	return nil
}
