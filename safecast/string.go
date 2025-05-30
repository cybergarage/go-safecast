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

// FromString casts an interface to a string type.
func FromString(from string, to any) error {
	switch to := to.(type) {
	case *int:
		v, err := strconv.ParseInt(from, 10, 64)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = int(v)
	case *int8:
		v, err := strconv.ParseInt(from, 10, 8)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = int8(v)
	case *int16:
		v, err := strconv.ParseInt(from, 10, 16)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = int16(v)
	case *int32:
		v, err := strconv.ParseInt(from, 10, 32)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = int32(v)
	case *int64:
		v, err := strconv.ParseInt(from, 10, 64)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = v
	case *uint:
		v, err := strconv.ParseUint(from, 10, 64)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = uint(v)
	case *uint8:
		v, err := strconv.ParseUint(from, 10, 8)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = uint8(v)
	case *uint16:
		v, err := strconv.ParseUint(from, 10, 16)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = uint16(v)
	case *uint32:
		v, err := strconv.ParseUint(from, 10, 32)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = uint32(v)
	case *uint64:
		v, err := strconv.ParseUint(from, 10, 64)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = v
	case *float32:
		v, err := strconv.ParseFloat(from, 32)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = float32(v)
	case *float64:
		v, err := strconv.ParseFloat(from, 64)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = v
	case *string:
		*to = from
	case *[]byte:
		*to = []byte(from)
	case *bool:
		v, err := strconv.ParseBool(from)
		if err != nil {
			return newErrorWithError(err)
		}
		*to = v
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToString casts an interface to a string type.
func ToString(from any, to *string) error {
	switch from := from.(type) {
	case int:
		*to = strconv.Itoa(from)
	case int8:
		*to = strconv.Itoa(int(from))
	case int16:
		*to = strconv.Itoa(int(from))
	case int32:
		*to = strconv.Itoa(int(from))
	case int64:
		*to = strconv.FormatInt(from, 10)
	case uint:
		*to = strconv.FormatUint(uint64(from), 10)
	case uint8:
		*to = strconv.Itoa(int(from))
	case uint16:
		*to = strconv.Itoa(int(from))
	case uint32:
		*to = strconv.Itoa(int(from))
	case uint64:
		*to = strconv.FormatUint(from, 10)
	case float32:
		*to = strconv.FormatFloat(float64(from), 'f', -1, 32)
	case float64:
		*to = strconv.FormatFloat(from, 'f', -1, 64)
	case bool:
		*to = strconv.FormatBool(from)
	case string:
		*to = from
	case *string:
		*to = *from
	case []byte:
		*to = string(from)
	case fmt.Stringer:
		*to = from.String()
	case nil:
		*to = ""
	default:
		*to = fmt.Sprintf("%v", from)
	}
	return nil
}
