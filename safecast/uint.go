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

// FromUint64 casts an interface to an uint64 type.
func FromUint64(from uint64, to any) error {
	switch to := to.(type) {
	case *int:
		if math.MaxInt < from {
			return newErrorOverRange(from, to)
		}
		*to = int(from)
	case *int8:
		if math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		*to = int8(from)
	case *int16:
		if math.MaxInt16 < from {
			return newErrorOverRange(from, to)
		}
		*to = int16(from)
	case *int32:
		if math.MaxInt32 < from {
			return newErrorOverRange(from, to)
		}
		*to = int32(from)
	case *int64:
		if math.MaxInt64 < from {
			return newErrorOverRange(from, to)
		}
		*to = int64(from)
	case *uint:
		*to = uint(from)
	case *uint8:
		if math.MaxUint8 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint8(from)
	case *uint16:
		if math.MaxUint16 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint16(from)
	case *uint32:
		if math.MaxUint32 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint32(from)
	case *uint64:
		*to = from
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

// FromUint8 casts an interface to an uint8 type.
func FromUint8(from uint8, to any) error {
	return FromUint64(uint64(from), to)
}

// FromUint16 casts an interface to an uint16 type.
func FromUint16(from uint16, to any) error {
	return FromUint64(uint64(from), to)
}

// FromUint32 casts an interface to an uint32 type.
func FromUint32(from uint32, to any) error {
	return FromUint64(uint64(from), to)
}

// FromUint casts an interface to an uint type.
func FromUint(from uint, to any) error {
	return FromUint64(uint64(from), to)
}

// ToUint8 casts an interface to an uint8 type.
func ToUint8(from any, to *uint8) error {
	switch from := from.(type) {
	case int:
		if from < 0 || math.MaxUint8 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint8(from)
	case int8:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint8(from)
	case int16:
		if from < 0 || math.MaxUint8 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint8(from)
	case int32:
		if from < 0 || math.MaxUint8 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint8(from)
	case int64:
		if from < 0 || math.MaxUint8 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint8(from)
	case uint:
		if math.MaxUint8 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint8(from)
	case uint8:
		*to = from
	case uint16:
		if math.MaxUint8 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint8(from)
	case uint32:
		if math.MaxUint8 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint8(from)
	case uint64:
		if math.MaxUint8 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint8(from)
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
		if s < 0 || math.MaxUint8 < s {
			return newErrorOverRange(s, to)
		}
		*to = uint8(s)
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToUint16 casts an interface to an uint16 type.
func ToUint16(from any, to *uint16) error {
	switch from := from.(type) {
	case int:
		if from < 0 || math.MaxUint16 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint16(from)
	case int8:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint16(from)
	case int16:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint16(from)
	case int32:
		if from < 0 || math.MaxUint16 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint16(from)
	case int64:
		if from < 0 || math.MaxUint16 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint16(from)
	case uint:
		if math.MaxUint16 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint16(from)
	case uint8:
		*to = uint16(from)
	case uint16:
		*to = from
	case uint32:
		if math.MaxUint16 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint16(from)
	case uint64:
		if math.MaxUint16 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint16(from)
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
		if s < 0 || math.MaxUint16 < s {
			return newErrorOverRange(s, to)
		}
		*to = uint16(s)
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToUint32 casts an interface to an uint32 type.
func ToUint32(from any, to *uint32) error {
	switch from := from.(type) {
	case int:
		if from < 0 || math.MaxUint32 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint32(from)
	case int8:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint32(from)
	case int16:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint32(from)
	case int32:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint32(from)
	case int64:
		if from < 0 || math.MaxUint32 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint32(from)
	case uint:
		if math.MaxUint32 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint32(from)
	case uint8:
		*to = uint32(from)
	case uint16:
		*to = uint32(from)
	case uint32:
		*to = from
	case uint64:
		if math.MaxUint32 < from {
			return newErrorOverRange(from, to)
		}
		*to = uint32(from)
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
		if s < 0 || math.MaxUint32 < s {
			return newErrorOverRange(s, to)
		}
		*to = uint32(s)
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToUint64 casts an interface to an uint64 type.
func ToUint64(from any, to *uint64) error {
	switch from := from.(type) {
	case int:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint64(from)
	case int8:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint64(from)
	case int16:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint64(from)
	case int32:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint64(from)
	case int64:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint64(from)
	case uint:
		*to = uint64(from)
	case uint8:
		*to = uint64(from)
	case uint16:
		*to = uint64(from)
	case uint32:
		*to = uint64(from)
	case uint64:
		*to = from
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
		if s < 0 {
			return newErrorOverRange(s, to)
		}
		*to = uint64(s)
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToUint casts an interface to an uint type.
func ToUint(from any, to *uint) error {
	switch from := from.(type) {
	case int:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint(from)
	case int8:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint(from)
	case int16:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint(from)
	case int32:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint(from)
	case int64:
		if from < 0 {
			return newErrorOverRange(from, to)
		}
		*to = uint(from)
	case uint:
		if math.MaxUint < from {
			return newErrorOverRange(from, to)
		}
		*to = from
	case uint8:
		*to = uint(from)
	case uint16:
		*to = uint(from)
	case uint32:
		*to = uint(from)
	case uint64:
		if math.MaxUint < from {
			return newErrorOverRange(from, to)
		}
		*to = uint(from)
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
		if s < 0 {
			return newErrorOverRange(s, to)
		}
		*to = uint(s)
	default:
		return newErrorCast(from, to)
	}
	return nil
}
