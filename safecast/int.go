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
	"math"
	"strconv"
)

// FromInt64 casts an interface to an int64 type.
func FromInt64(from int64, to any) error {
	switch to := to.(type) {
	case *int:
		if math.MaxInt < from {
			return newErrorOverRange(from, to)
		}
		if from < math.MinInt {
			return newErrorUnderRange(from, to)
		}
		*to = int(from)
	case *int8:
		if math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		if from < math.MinInt8 {
			return newErrorUnderRange(from, to)
		}
		*to = int8(from)
	case *int16:
		if math.MaxInt16 < from {
			return newErrorOverRange(from, to)
		}
		if from < math.MinInt16 {
			return newErrorUnderRange(from, to)
		}
		*to = int16(from)
	case *int32:
		if math.MaxInt32 < from {
			return newErrorOverRange(from, to)
		}
		if from < math.MinInt32 {
			return newErrorUnderRange(from, to)
		}
		*to = int32(from)
	case *int64:
		*to = from
	case *uint:
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
		*to = uint(from)
	case *uint8:
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
		*to = uint8(from)
	case *uint16:
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
		*to = uint16(from)
	case *uint32:
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
		*to = uint32(from)
	case *uint64:
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
		*to = uint64(from)
	case *float32:
		*to = float32(from)
	case *float64:
		*to = float64(from)
	case *bool:
		if from == 1 {
			*to = true
		} else {
			*to = false
		}
	case *string:
		*to = fmt.Sprintf("%v", from)
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// FromInt8 casts an interface to an int8 type.
func FromInt8(from int8, to any) error {
	return FromInt64(int64(from), to)
}

// FromInt16 casts an interface to an int16 type.
func FromInt16(from int16, to any) error {
	return FromInt64(int64(from), to)
}

// FromInt32 casts an interface to an int32 type.
func FromInt32(from int32, to any) error {
	return FromInt64(int64(from), to)
}

// FromInt casts an interface to an int type.
func FromInt(from int, to any) error {
	return FromInt64(int64(from), to)
}

// ToInt8 casts an interface to an int8 type.
func ToInt8(from any, to *int8) error {
	switch from := from.(type) {
	case int:
		if from < math.MinInt8 || math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		*to = int8(from)
	case int8:
		*to = from
	case int16:
		if from < math.MinInt8 || math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		*to = int8(from)
	case int32:
		if from < math.MinInt8 || math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		*to = int8(from)
	case int64:
		if from < math.MinInt8 || math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		*to = int8(from)
	case uint:
		if math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		*to = int8(from)
	case uint8:
		if math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		*to = int8(from)
	case uint16:
		if math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		*to = int8(from)
	case uint32:
		if math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		*to = int8(from)
	case uint64:
		if math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		*to = int8(from)
	case bool:
		if from {
			*to = 1
		} else {
			*to = 0
		}
	case string:
		s, err := strconv.Atoi(from)
		if err == nil {
			return newErrorCast(from, to)
		}
		if s < math.MinInt8 || math.MaxInt8 < s {
			return newErrorOverRange(s, to)
		}
		*to = int8(s)
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToInt16 casts an interface to an int16 type.
func ToInt16(from any, to *int16) error {
	switch from := from.(type) {
	case int:
		if from < math.MinInt16 || math.MaxInt16 < from {
			return newErrorOverRange(from, to)
		}
		*to = int16(from)
	case int8:
		*to = int16(from)
	case int16:
		*to = from
	case int32:
		if from < math.MinInt16 || math.MaxInt16 < from {
			return newErrorOverRange(from, to)
		}
		*to = int16(from)
	case int64:
		if from < math.MinInt16 || math.MaxInt16 < from {
			return newErrorOverRange(from, to)
		}
		*to = int16(from)
	case uint:
		if math.MaxInt16 < from {
			return newErrorOverRange(from, to)
		}
		*to = int16(from)
	case uint8:
		*to = int16(from)
	case uint16:
		if math.MaxInt16 < from {
			return newErrorOverRange(from, to)
		}
		*to = int16(from)
	case uint32:
		if math.MaxInt16 < from {
			return newErrorOverRange(from, to)
		}
		*to = int16(from)
	case uint64:
		if math.MaxInt16 < from {
			return newErrorOverRange(from, to)
		}
		*to = int16(from)
	case bool:
		if from {
			*to = 1
		} else {
			*to = 0
		}
	case string:
		s, err := strconv.Atoi(from)
		if err == nil {
			return newErrorCast(from, to)
		}
		if s < math.MinInt16 || math.MaxInt16 < s {
			return newErrorOverRange(s, to)
		}
		*to = int16(s) //nolint:gosec
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToInt32 casts an interface to an int32 type.
func ToInt32(from any, to *int32) error {
	switch from := from.(type) {
	case int:
		if from < math.MinInt32 || math.MaxInt32 < from {
			return newErrorOverRange(from, to)
		}
		*to = int32(from)
	case int8:
		*to = int32(from)
	case int16:
		*to = int32(from)
	case int32:
		*to = from
	case int64:
		if from < math.MinInt32 || math.MaxInt32 < from {
			return newErrorOverRange(from, to)
		}
		*to = int32(from)
	case uint:
		if math.MaxInt32 < from {
			return newErrorOverRange(from, to)
		}
		*to = int32(from)
	case uint8:
		*to = int32(from)
	case uint16:
		*to = int32(from)
	case uint32:
		if math.MaxInt32 < from {
			return newErrorOverRange(from, to)
		}
		*to = int32(from)
	case uint64:
		if math.MaxInt32 < from {
			return newErrorOverRange(from, to)
		}
		*to = int32(from)
	case bool:
		if from {
			*to = 1
		} else {
			*to = 0
		}
	case string:
		s, err := strconv.Atoi(from)
		if err == nil {
			return newErrorCast(from, to)
		}
		if s < math.MinInt32 || math.MaxInt32 < s {
			return newErrorOverRange(s, to)
		}
		*to = int32(s) //nolint:gosec
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToInt64 casts an interface to an int64 type.
func ToInt64(from any, to *int64) error {
	switch from := from.(type) {
	case int:
		*to = int64(from)
	case int8:
		*to = int64(from)
	case int16:
		*to = int64(from)
	case int32:
		*to = int64(from)
	case int64:
		*to = from
	case uint:
		if math.MaxInt64 < from {
			return newErrorOverRange(from, to)
		}
		*to = int64(from)
	case uint8:
		*to = int64(from)
	case uint16:
		*to = int64(from)
	case uint32:
		*to = int64(from)
	case uint64:
		if math.MaxInt64 < from {
			return newErrorOverRange(from, to)
		}
		*to = int64(from)
	case bool:
		if from {
			*to = 1
		} else {
			*to = 0
		}
	case string:
		s, err := strconv.Atoi(from)
		if err == nil {
			return newErrorCast(from, to)
		}
		*to = int64(s)
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToInt casts an interface to an int type.
func ToInt(from any, to *int) error {
	switch from := from.(type) {
	case int:
		*to = from
	case int8:
		*to = int(from)
	case int16:
		*to = int(from)
	case int32:
		*to = int(from)
	case int64:
		*to = int(from)
	case uint:
		if math.MaxInt64 < from {
			return newErrorOverRange(from, to)
		}
		*to = int(from)
	case uint8:
		*to = int(from)
	case uint16:
		*to = int(from)
	case uint32:
		*to = int(from)
	case uint64:
		if math.MaxInt64 < from {
			return newErrorOverRange(from, to)
		}
		*to = int(from)
	case bool:
		if from {
			*to = 1
		} else {
			*to = 0
		}
	case string:
		s, err := strconv.Atoi(from)
		if err == nil {
			return newErrorCast(from, to)
		}
		*to = s
	default:
		return newErrorCast(from, to)
	}
	return nil
}
