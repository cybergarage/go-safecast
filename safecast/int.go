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
		if int64(math.MaxUint8) < from {
			return newErrorOverRange(from, to)
		}
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
		*to = uint8(from)
	case *uint16:
		if int64(math.MaxUint16) < from {
			return newErrorOverRange(from, to)
		}
		if from < 0 {
			return newErrorUnderRange(from, to)
		}
		*to = uint16(from)
	case *uint32:
		if int64(math.MaxUint32) < from {
			return newErrorOverRange(from, to)
		}
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
	fromBool := func(v bool) int8 {
		if v {
			return 1
		}
		return 0
	}

	fromString := func(v string) (int8, error) {
		iv, err := strconv.Atoi(v)
		if err == nil {
			return *to, ToInt8(iv, to)
		}
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return *to, ToInt8(fv, to)
		}
		return 0, newErrorCast(v, to)
	}

	var err error
	switch from := from.(type) {
	case int:
		if math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		if from < math.MinInt8 {
			return newErrorUnderRange(from, to)
		}
		*to = int8(from)
	case int8:
		*to = from
	case int16:
		if math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		if from < math.MinInt8 {
			return newErrorUnderRange(from, to)
		}
		*to = int8(from)
	case int32:
		if math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		if from < math.MinInt8 {
			return newErrorUnderRange(from, to)
		}
		*to = int8(from)
	case int64:
		if math.MaxInt8 < from {
			return newErrorOverRange(from, to)
		}
		if from < math.MinInt8 {
			return newErrorUnderRange(from, to)
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
	case float32:
		return ToInt8(int64(from), to)
	case float64:
		return ToInt8(int64(from), to)
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

// ToInt16 casts an interface to an int16 type.
func ToInt16(from any, to *int16) error {
	fromInt := func(v int) (int16, error) {
		if math.MaxInt16 < v {
			return 0, newErrorOverRange(v, to)
		}
		if v < math.MinInt16 {
			return 0, newErrorUnderRange(v, to)
		}
		return int16(v), nil
	}

	fromInt32 := func(v int32) (int16, error) {
		if math.MaxInt16 < v {
			return 0, newErrorOverRange(v, to)
		}
		if v < math.MinInt16 {
			return 0, newErrorUnderRange(v, to)
		}
		return int16(v), nil
	}

	fromInt64 := func(v int64) (int16, error) {
		if math.MaxInt16 < v {
			return 0, newErrorOverRange(v, to)
		}
		if v < math.MinInt16 {
			return 0, newErrorUnderRange(v, to)
		}
		return int16(v), nil
	}

	fromUint := func(v uint) (int16, error) {
		if math.MaxInt16 < v {
			return 0, newErrorOverRange(v, to)
		}
		return int16(v), nil
	}

	fromUint16 := func(v uint16) (int16, error) {
		if math.MaxInt16 < v {
			return 0, newErrorOverRange(v, to)
		}
		return int16(v), nil
	}

	fromUint32 := func(v uint32) (int16, error) {
		if math.MaxInt16 < v {
			return 0, newErrorOverRange(v, to)
		}
		return int16(v), nil
	}

	fromUint64 := func(v uint64) (int16, error) {
		if math.MaxInt16 < v {
			return 0, newErrorOverRange(v, to)
		}
		return int16(v), nil
	}

	fromBool := func(v bool) int16 {
		if v {
			return 1
		}
		return 0
	}

	fromString := func(v string) (int16, error) {
		iv, err := strconv.Atoi(v)
		if err == nil {
			return *to, ToInt16(iv, to)
		}
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return *to, ToInt16(fv, to)
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
		*to = int16(from)
	case *int8:
		*to = int16(*from)
	case int16:
		*to = from
	case *int16:
		*to = *from
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
		*to = int16(from)
	case *uint8:
		*to = int16(*from)
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
		return ToInt16(int64(from), to)
	case *float32:
		return ToInt16(int64(*from), to)
	case float64:
		return ToInt16(int64(from), to)
	case *float64:
		return ToInt16(int64(*from), to)
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

// ToInt32 casts an interface to an int32 type.
func ToInt32(from any, to *int32) error {
	fromInt := func(v int) (int32, error) {
		if math.MaxInt32 < v {
			return 0, newErrorOverRange(v, to)
		}
		if v < math.MinInt32 {
			return 0, newErrorUnderRange(v, to)
		}
		return int32(v), nil
	}

	fromInt64 := func(v int64) (int32, error) {
		if math.MaxInt32 < v {
			return 0, newErrorOverRange(v, to)
		}
		if v < math.MinInt32 {
			return 0, newErrorUnderRange(v, to)
		}
		return int32(v), nil
	}

	fromUint := func(v uint) (int32, error) {
		if math.MaxInt32 < v {
			return 0, newErrorOverRange(v, to)
		}
		return int32(v), nil
	}

	fromUint32 := func(v uint32) (int32, error) {
		if math.MaxInt32 < v {
			return 0, newErrorOverRange(v, to)
		}
		return int32(v), nil
	}

	fromUint64 := func(v uint64) (int32, error) {
		if math.MaxInt32 < v {
			return 0, newErrorOverRange(v, to)
		}
		return int32(v), nil
	}

	fromBool := func(v bool) int32 {
		if v {
			return 1
		}
		return 0
	}

	fromString := func(v string) (int32, error) {
		iv, err := strconv.Atoi(v)
		if err == nil {
			return *to, ToInt32(iv, to)
		}
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return *to, ToInt32(fv, to)
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
		*to = int32(from)
	case *int8:
		*to = int32(*from)
	case int16:
		*to = int32(from)
	case *int16:
		*to = int32(*from)
	case int32:
		*to = from
	case *int32:
		*to = *from
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
		*to = int32(from)
	case *uint8:
		*to = int32(*from)
	case uint16:
		*to = int32(from)
	case *uint16:
		*to = int32(*from)
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
		return ToInt32(int64(from), to)
	case *float32:
		return ToInt32(int64(*from), to)
	case float64:
		return ToInt32(int64(from), to)
	case *float64:
		return ToInt32(int64(*from), to)
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

// ToInt64 casts an interface to an int64 type.
func ToInt64(from any, to *int64) error {
	fromUint := func(v uint) (int64, error) {
		if math.MaxInt < v {
			return 0, newErrorOverRange(v, to)
		}
		return int64(v), nil
	}

	fromUint64 := func(v uint64) (int64, error) {
		if math.MaxInt < v {
			return 0, newErrorOverRange(v, to)
		}
		return int64(v), nil
	}

	fromBool := func(v bool) int64 {
		if v {
			return 1
		}
		return 0
	}

	fromString := func(v string) (int64, error) {
		iv, err := strconv.Atoi(v)
		if err == nil {
			return *to, ToInt64(iv, to)
		}
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return *to, ToInt64(fv, to)
		}
		return 0, newErrorCast(from, to)
	}

	var err error
	switch from := from.(type) {
	case int:
		*to = int64(from)
	case *int:
		*to = int64(*from)
	case int8:
		*to = int64(from)
	case *int8:
		*to = int64(*from)
	case int16:
		*to = int64(from)
	case *int16:
		*to = int64(*from)
	case int32:
		*to = int64(from)
	case *int32:
		*to = int64(*from)
	case int64:
		*to = from
	case *int64:
		*to = *from
	case uint:
		if *to, err = fromUint(from); err != nil {
			return err
		}
	case *uint:
		if *to, err = fromUint(*from); err != nil {
			return err
		}
	case uint8:
		*to = int64(from)
	case *uint8:
		*to = int64(*from)
	case uint16:
		*to = int64(from)
	case *uint16:
		*to = int64(*from)
	case uint32:
		*to = int64(from)
	case *uint32:
		*to = int64(*from)
	case uint64:
		if *to, err = fromUint64(from); err != nil {
			return err
		}
	case *uint64:
		if *to, err = fromUint64(*from); err != nil {
			return err
		}
	case float32:
		*to = int64(from)
	case *float32:
		*to = int64(*from)
	case float64:
		*to = int64(from)
	case *float64:
		*to = int64(*from)
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

// ToInt casts an interface to an int type.
func ToInt(from any, to *int) error {
	fromUint := func(v uint) (int, error) {
		if math.MaxInt < v {
			return 0, newErrorOverRange(v, to)
		}
		return int(v), nil
	}

	fromUint64 := func(v uint64) (int, error) {
		if math.MaxInt < v {
			return 0, newErrorOverRange(v, to)
		}
		return int(v), nil
	}

	fromBool := func(v bool) int {
		if v {
			return 1
		}
		return 0
	}

	fromString := func(v string) (int, error) {
		iv, err := strconv.Atoi(v)
		if err == nil {
			return iv, nil
		}
		fv, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return int(fv), nil
		}
		return 0, newErrorCast(v, to)
	}

	var err error
	switch from := from.(type) {
	case int:
		*to = from
	case *int:
		*to = *from
	case int8:
		*to = int(from)
	case *int8:
		*to = int(*from)
	case int16:
		*to = int(from)
	case *int16:
		*to = int(*from)
	case int32:
		*to = int(from)
	case *int32:
		*to = int(*from)
	case int64:
		*to = int(from)
	case *int64:
		*to = int(*from)
	case uint:
		if *to, err = fromUint(from); err != nil {
			return err
		}
	case *uint:
		if *to, err = fromUint(*from); err != nil {
			return err
		}
	case uint8:
		*to = int(from)
	case *uint8:
		*to = int(*from)
	case uint16:
		*to = int(from)
	case *uint16:
		*to = int(*from)
	case uint32:
		*to = int(from)
	case *uint32:
		*to = int(*from)
	case uint64:
		if *to, err = fromUint64(from); err != nil {
			return err
		}
	case *uint64:
		if *to, err = fromUint64(*from); err != nil {
			return err
		}
	case float32:
		*to = int(from)
	case *float32:
		*to = int(*from)
	case float64:
		*to = int(from)
	case *float64:
		*to = int(*from)
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
