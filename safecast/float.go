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

// FromFloat64 casts an interface to an float64 type.
func FromFloat64(from float64, to any) error {
	switch to := to.(type) {
	case *int:
		if float64(math.MaxInt) < from {
			return newErrorOverRange(from, to)
		}
		if from < float64(math.MinInt) {
			return newErrorUnderRange(from, to)
		}
		*to = int(from)
	case *int8:
		if float64(math.MaxInt8) < from {
			return newErrorOverRange(from, to)
		}
		if from < float64(math.MinInt8) {
			return newErrorUnderRange(from, to)
		}
		*to = int8(from)
	case *int16:
		if float64(math.MaxInt16) < from {
			return newErrorOverRange(from, to)
		}
		if from < float64(math.MinInt16) {
			return newErrorUnderRange(from, to)
		}
		*to = int16(from)
	case *int32:
		if float64(math.MaxInt32) < from {
			return newErrorOverRange(from, to)
		}
		if from < float64(math.MinInt32) {
			return newErrorUnderRange(from, to)
		}
		*to = int32(from)
	case *int64:
		if float64(math.MaxInt64) < from {
			return newErrorOverRange(from, to)
		}
		if from < float64(math.MinInt64) {
			return newErrorUnderRange(from, to)
		}
		*to = int64(from)
	case *uint:
		if float64(math.MaxUint) < from {
			return newErrorOverRange(from, to)
		}
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
		*to = uint(from)
	case *uint8:
		if float64(math.MaxUint8) < from {
			return newErrorOverRange(from, to)
		}
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
		*to = uint8(from)
	case *uint16:
		if float64(math.MaxUint16) < from {
			return newErrorOverRange(from, to)
		}
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
		*to = uint16(from)
	case *uint32:
		if float64(math.MaxUint32) < from {
			return newErrorOverRange(from, to)
		}
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
		*to = uint32(from)
	case *uint64:
		if float64(math.MaxUint64) < from {
			return newErrorOverRange(from, to)
		}
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
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
