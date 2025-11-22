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
	fromInt := func(v int) (uint8, error) {
		if math.MaxUint8 < v {
			return 0, newErrorOverRange(v, to)
		}
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint8(v), nil
	}

	fromInt8 := func(v int8) (uint8, error) {
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint8(v), nil
	}

	fromInt16 := func(v int16) (uint8, error) {
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint8(v), nil
	}

	fromInt32 := func(v int32) (uint8, error) {
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint8(v), nil
	}

	fromInt64 := func(v int64) (uint8, error) {
		if math.MaxUint8 < v {
			return 0, newErrorOverRange(v, to)
		}
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint8(v), nil
	}

	fromUint := func(v uint) (uint8, error) {
		if math.MaxUint8 < v {
			return 0, newErrorOverRange(v, to)
		}
		return uint8(v), nil
	}

	fromUint16 := func(v uint16) (uint8, error) {
		if math.MaxUint8 < v {
			return 0, newErrorOverRange(v, to)
		}
		return uint8(v), nil
	}

	fromUint32 := func(v uint32) (uint8, error) {
		if math.MaxUint8 < v {
			return 0, newErrorOverRange(v, to)
		}
		return uint8(v), nil
	}

	fromUint64 := func(v uint64) (uint8, error) {
		if math.MaxUint8 < v {
			return 0, newErrorOverRange(v, to)
		}
		return uint8(v), nil
	}
	fromBool := func(v bool) uint8 {
		if v {
			return 1
		}
		return 0
	}

	fromString := func(v string) (uint8, error) {
		iv, err := strconv.Atoi(v)
		if err == nil {
			return *to, ToUint8(iv, to)
		}
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return *to, ToUint8(fv, to)
		}
		return 0, newErrorCast(from, to)
	}

	var err error
	switch from := from.(type) {
	case int:
		if *to, err = fromInt(from); err != nil {
			return err
		}
	case *int:
		if *to, err = fromInt(*from); err != nil {
			return err
		}
	case int8:
		if *to, err = fromInt8(from); err != nil {
			return err
		}
	case *int8:
		if *to, err = fromInt8(*from); err != nil {
			return err
		}
	case int16:
		if *to, err = fromInt16(from); err != nil {
			return err
		}
	case *int16:
		if *to, err = fromInt16(*from); err != nil {
			return err
		}
	case int32:
		if *to, err = fromInt32(from); err != nil {
			return err
		}
	case *int32:
		if *to, err = fromInt32(*from); err != nil {
			return err
		}
	case int64:
		if *to, err = fromInt64(from); err != nil {
			return err
		}
	case *int64:
		if *to, err = fromInt64(*from); err != nil {
			return err
		}
	case uint:
		if *to, err = fromUint(from); err != nil {
			return err
		}
	case *uint:
		if *to, err = fromUint(*from); err != nil {
			return err
		}
	case uint8:
		*to = from
	case *uint8:
		*to = *from
	case uint16:
		if *to, err = fromUint16(from); err != nil {
			return err
		}
	case *uint16:
		if *to, err = fromUint16(*from); err != nil {
			return err
		}
	case uint32:
		if *to, err = fromUint32(from); err != nil {
			return err
		}
	case *uint32:
		if *to, err = fromUint32(*from); err != nil {
			return err
		}
	case uint64:
		if *to, err = fromUint64(from); err != nil {
			return err
		}
	case *uint64:
		if *to, err = fromUint64(*from); err != nil {
			return err
		}
	case float32:
		return ToUint8(int64(from), to)
	case *float32:
		return ToUint8(int64(*from), to)
	case float64:
		return ToUint8(int64(from), to)
	case *float64:
		return ToUint8(int64(*from), to)
	case bool:
		*to = fromBool(from)
	case *bool:
		*to = fromBool(*from)
	case string:
		if *to, err = fromString(from); err != nil {
			return err
		}
	case *string:
		if *to, err = fromString(*from); err != nil {
			return err
		}
	case []byte:
		if *to, err = fromString(string(from)); err != nil {
			return err
		}
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToUint16 casts an interface to an uint16 type.
func ToUint16(from any, to *uint16) error {
	fromInt := func(v int) (uint16, error) {
		if math.MaxUint16 < v {
			return 0, newErrorOverRange(v, to)
		}
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint16(v), nil
	}

	fromInt8 := func(v int8) (uint16, error) {
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint16(v), nil
	}

	fromInt16 := func(v int16) (uint16, error) {
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint16(v), nil
	}

	fromInt32 := func(v int32) (uint16, error) {
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint16(v), nil
	}

	fromInt64 := func(v int64) (uint16, error) {
		if math.MaxUint16 < v {
			return 0, newErrorOverRange(v, to)
		}
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint16(v), nil
	}

	fromUint := func(v uint) (uint16, error) {
		if math.MaxUint16 < v {
			return 0, newErrorOverRange(v, to)
		}
		return uint16(v), nil
	}

	fromUint32 := func(v uint32) (uint16, error) {
		if math.MaxUint16 < v {
			return 0, newErrorOverRange(v, to)
		}
		return uint16(v), nil
	}

	fromUint64 := func(v uint64) (uint16, error) {
		if math.MaxUint16 < v {
			return 0, newErrorOverRange(v, to)
		}
		return uint16(v), nil
	}

	fromBool := func(v bool) uint16 {
		if v {
			return 1
		}
		return 0
	}

	fromString := func(v string) (uint16, error) {
		iv, err := strconv.Atoi(v)
		if err == nil {
			return *to, ToUint16(iv, to)
		}
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return *to, ToUint16(fv, to)
		}
		return 0, newErrorCast(from, to)
	}

	var err error
	switch from := from.(type) {
	case int:
		if *to, err = fromInt(from); err != nil {
			return err
		}
	case *int:
		if *to, err = fromInt(*from); err != nil {
			return err
		}
	case int8:
		if *to, err = fromInt8(from); err != nil {
			return err
		}
	case *int8:
		if *to, err = fromInt8(*from); err != nil {
			return err
		}
	case int16:
		if *to, err = fromInt16(from); err != nil {
			return err
		}
	case *int16:
		if *to, err = fromInt16(*from); err != nil {
			return err
		}
	case int32:
		if *to, err = fromInt32(from); err != nil {
			return err
		}
	case *int32:
		if *to, err = fromInt32(*from); err != nil {
			return err
		}
	case int64:
		if *to, err = fromInt64(from); err != nil {
			return err
		}
	case *int64:
		if *to, err = fromInt64(*from); err != nil {
			return err
		}
	case uint:
		if *to, err = fromUint(from); err != nil {
			return err
		}
	case *uint:
		if *to, err = fromUint(*from); err != nil {
			return err
		}
	case uint8:
		*to = uint16(from)
	case *uint8:
		*to = uint16(*from)
	case uint16:
		*to = from
	case *uint16:
		*to = *from
	case uint32:
		if *to, err = fromUint32(from); err != nil {
			return err
		}
	case *uint32:
		if *to, err = fromUint32(*from); err != nil {
			return err
		}
	case uint64:
		if *to, err = fromUint64(from); err != nil {
			return err
		}
	case *uint64:
		if *to, err = fromUint64(*from); err != nil {
			return err
		}
	case float32:
		return ToUint16(int64(from), to)
	case *float32:
		return ToUint16(int64(*from), to)
	case float64:
		return ToUint16(int64(from), to)
	case *float64:
		return ToUint16(int64(*from), to)
	case bool:
		*to = fromBool(from)
	case *bool:
		*to = fromBool(*from)
	case string:
		if *to, err = fromString(from); err != nil {
			return err
		}
	case *string:
		if *to, err = fromString(*from); err != nil {
			return err
		}
	case []byte:
		if *to, err = fromString(string(from)); err != nil {
			return err
		}
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToUint32 casts an interface to an uint32 type.
func ToUint32(from any, to *uint32) error {
	fromInt := func(v int) (uint32, error) {
		if math.MaxUint32 < v {
			return 0, newErrorOverRange(v, to)
		}
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint32(v), nil
	}

	fromInt8 := func(v int8) (uint32, error) {
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint32(v), nil
	}

	fromInt16 := func(v int16) (uint32, error) {
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint32(v), nil
	}

	fromInt32 := func(v int32) (uint32, error) {
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint32(v), nil
	}

	fromInt64 := func(v int64) (uint32, error) {
		if math.MaxUint32 < v {
			return 0, newErrorOverRange(v, to)
		}
		if v < 0 {
			return 0, newErrorUnderRange(v, to)
		}
		return uint32(v), nil
	}

	fromUint := func(v uint) (uint32, error) {
		if math.MaxUint32 < v {
			return 0, newErrorOverRange(v, to)
		}
		return uint32(v), nil
	}

	fromUint64 := func(v uint64) (uint32, error) {
		if math.MaxUint32 < v {
			return 0, newErrorOverRange(v, to)
		}
		return uint32(v), nil
	}

	fromBool := func(v bool) uint32 {
		if v {
			return 1
		}
		return 0
	}

	fromString := func(v string) (uint32, error) {
		iv, err := strconv.Atoi(v)
		if err == nil {
			return *to, ToUint32(iv, to)
		}
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return *to, ToUint32(fv, to)
		}
		return 0, newErrorCast(from, to)
	}

	var err error
	switch from := from.(type) {
	case int:
		if *to, err = fromInt(from); err != nil {
			return err
		}
	case *int:
		if *to, err = fromInt(*from); err != nil {
			return err
		}
	case int8:
		if *to, err = fromInt8(from); err != nil {
			return err
		}
	case *int8:
		if *to, err = fromInt8(*from); err != nil {
			return err
		}
	case int16:
		if *to, err = fromInt16(from); err != nil {
			return err
		}
	case *int16:
		if *to, err = fromInt16(*from); err != nil {
			return err
		}
	case int32:
		if *to, err = fromInt32(from); err != nil {
			return err
		}
	case *int32:
		if *to, err = fromInt32(*from); err != nil {
			return err
		}
	case int64:
		if *to, err = fromInt64(from); err != nil {
			return err
		}
	case *int64:
		if *to, err = fromInt64(*from); err != nil {
			return err
		}
	case uint:
		if *to, err = fromUint(from); err != nil {
			return err
		}
	case *uint:
		if *to, err = fromUint(*from); err != nil {
			return err
		}
	case uint8:
		*to = uint32(from)
	case *uint8:
		*to = uint32(*from)
	case uint16:
		*to = uint32(from)
	case *uint16:
		*to = uint32(*from)
	case uint32:
		*to = from
	case *uint32:
		*to = *from
	case uint64:
		if *to, err = fromUint64(uint64(from)); err != nil {
			return err
		}
	case *uint64:
		if *to, err = fromUint64(*from); err != nil {
			return err
		}
	case float32:
		return ToUint32(int64(from), to)
	case *float32:
		return ToUint32(int64(*from), to)
	case float64:
		return ToUint32(int64(from), to)
	case *float64:
		return ToUint32(int64(*from), to)
	case bool:
		*to = fromBool(from)
	case *bool:
		*to = fromBool(*from)
	case string:
		if *to, err = fromString(from); err != nil {
			return err
		}
	case *string:
		if *to, err = fromString(*from); err != nil {
			return err
		}
	case []byte:
		if *to, err = fromString(string(from)); err != nil {
			return err
		}
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToUint64 casts an interface to an uint64 type.
func ToUint64(from any, to *uint64) error {
	fromInt := func(v int) (uint64, error) {
		if v < 0 {
			return 0, newErrorUnderRange(from, to)
		}
		return uint64(v), nil
	}

	fromInt8 := func(v int8) (uint64, error) {
		if v < 0 {
			return 0, newErrorUnderRange(from, to)
		}
		return uint64(v), nil
	}

	fromInt16 := func(v int16) (uint64, error) {
		if v < 0 {
			return 0, newErrorUnderRange(from, to)
		}
		return uint64(v), nil
	}

	fromInt32 := func(v int32) (uint64, error) {
		if v < 0 {
			return 0, newErrorUnderRange(from, to)
		}
		return uint64(v), nil
	}

	fromInt64 := func(v int64) (uint64, error) {
		if v < 0 {
			return 0, newErrorUnderRange(from, to)
		}
		return uint64(v), nil
	}

	fromBool := func(v bool) uint64 {
		if v {
			return 1
		}
		return 0
	}

	fromString := func(v string) (uint64, error) {
		uv, err := strconv.ParseUint(v, 10, 64)
		if err == nil {
			return uv, nil
		}
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return *to, ToUint64(fv, to)
		}
		return 0, newErrorCast(from, to)
	}

	var err error
	switch from := from.(type) {
	case int:
		if *to, err = fromInt(from); err != nil {
			return err
		}
	case *int:
		if *to, err = fromInt(*from); err != nil {
			return err
		}
	case int8:
		if *to, err = fromInt8(from); err != nil {
			return err
		}
	case *int8:
		if *to, err = fromInt8(*from); err != nil {
			return err
		}
	case int16:
		if *to, err = fromInt16(from); err != nil {
			return err
		}
	case *int16:
		if *to, err = fromInt16(*from); err != nil {
			return err
		}
	case int32:
		if *to, err = fromInt32(from); err != nil {
			return err
		}
	case *int32:
		if *to, err = fromInt32(*from); err != nil {
			return err
		}
	case int64:
		if *to, err = fromInt64(from); err != nil {
			return err
		}
	case *int64:
		if *to, err = fromInt64(*from); err != nil {
			return err
		}
	case uint:
		*to = uint64(from)
	case *uint:
		*to = uint64(*from)
	case uint8:
		*to = uint64(from)
	case *uint8:
		*to = uint64(*from)
	case uint16:
		*to = uint64(from)
	case *uint16:
		*to = uint64(*from)
	case uint32:
		*to = uint64(from)
	case *uint32:
		*to = uint64(*from)
	case uint64:
		*to = from
	case *uint64:
		*to = *from
	case float32:
		return ToUint64(int64(from), to)
	case *float32:
		return ToUint64(int64(*from), to)
	case float64:
		return ToUint64(int64(from), to)
	case *float64:
		return ToUint64(int64(*from), to)
	case bool:
		*to = fromBool(from)
	case *bool:
		*to = fromBool(*from)
	case string:
		if *to, err = fromString(from); err != nil {
			return err
		}
	case *string:
		if *to, err = fromString(*from); err != nil {
			return err
		}
	case []byte:
		if *to, err = fromString(string(from)); err != nil {
			return err
		}
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToUint casts an interface to an uint type.
func ToUint(from any, to *uint) error {
	fromInt := func(v int) (uint, error) {
		if v < 0 {
			return 0, newErrorUnderRange(from, to)
		}
		return uint(v), nil
	}

	fromInt8 := func(v int8) (uint, error) {
		if v < 0 {
			return 0, newErrorUnderRange(from, to)
		}
		return uint(v), nil
	}

	fromInt16 := func(v int16) (uint, error) {
		if v < 0 {
			return 0, newErrorUnderRange(from, to)
		}
		return uint(v), nil
	}

	fromInt32 := func(v int32) (uint, error) {
		if v < 0 {
			return 0, newErrorUnderRange(from, to)
		}
		return uint(v), nil
	}

	fromInt64 := func(v int64) (uint, error) {
		if v < 0 {
			return 0, newErrorUnderRange(from, to)
		}
		return uint(v), nil
	}

	fromBool := func(v bool) uint {
		if v {
			return 1
		}
		return 0
	}

	fromString := func(v string) (uint, error) {
		uv, err := strconv.ParseUint(v, 10, 64)
		if err == nil {
			return uint(uv), nil
		}
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return *to, ToUint(fv, to)
		}
		return 0, newErrorCast(v, to)
	}

	var err error
	switch from := from.(type) {
	case int:
		if *to, err = fromInt(from); err != nil {
			return err
		}
	case *int:
		if *to, err = fromInt(*from); err != nil {
			return err
		}
	case int8:
		if *to, err = fromInt8(from); err != nil {
			return err
		}
	case *int8:
		if *to, err = fromInt8(*from); err != nil {
			return err
		}
	case int16:
		if *to, err = fromInt16(from); err != nil {
			return err
		}
	case *int16:
		if *to, err = fromInt16(*from); err != nil {
			return err
		}
	case int32:
		if *to, err = fromInt32(from); err != nil {
			return err
		}
	case *int32:
		if *to, err = fromInt32(*from); err != nil {
			return err
		}
	case int64:
		if *to, err = fromInt64(from); err != nil {
			return err
		}
	case *int64:
		if *to, err = fromInt64(*from); err != nil {
			return err
		}
		*to = uint(*from)
	case uint:
		*to = from
	case *uint:
		*to = *from
	case uint8:
		*to = uint(from)
	case *uint8:
		*to = uint(*from)
	case uint16:
		*to = uint(from)
	case *uint16:
		*to = uint(*from)
	case uint32:
		*to = uint(from)
	case *uint32:
		*to = uint(*from)
	case uint64:
		*to = uint(from)
	case *uint64:
		*to = uint(*from)
	case float32:
		return ToUint(int64(from), to)
	case *float32:
		return ToUint(int64(*from), to)
	case float64:
		return ToUint(int64(from), to)
	case *float64:
		return ToUint(int64(*from), to)
	case bool:
		*to = fromBool(from)
	case *bool:
		*to = fromBool(*from)
	case string:
		if *to, err = fromString(from); err != nil {
			return err
		}
	case *string:
		if *to, err = fromString(*from); err != nil {
			return err
		}
	case []byte:
		if *to, err = fromString(string(from)); err != nil {
			return err
		}
	default:
		return newErrorCast(from, to)
	}
	return nil
}
