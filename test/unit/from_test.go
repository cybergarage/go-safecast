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

package test

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/cybergarage/go-safecast/safecast"
)

// nolint: gocyclo, maintidx
func TestFromCast(t *testing.T) {
	var vi int
	var vi8 int8
	var vi16 int16
	var vi32 int32
	var vi64 int64
	var uvi uint
	var uvi8 uint8
	var uvi16 uint16
	var uvi32 uint32
	var uvi64 uint64
	var vf32 float32
	var vf64 float64
	var vs string
	var vb bool

	t.Run("safecast.FromInt64", func(t *testing.T) {
		froms := []int64{
			math.MinInt,
			math.MaxInt,
			math.MinInt8,
			math.MaxInt8,
			math.MinInt16,
			math.MaxInt16,
			math.MinInt32,
			math.MaxInt32,
			math.MinInt,
			math.MaxInt,
			math.MinInt64,
			math.MaxInt64,
			0,
			math.MaxInt,
			0,
			math.MaxInt8,
			0,
			math.MaxInt16,
			0,
			math.MaxInt32,
			0,
			math.MaxInt,
			0,
			math.MaxInt64,
			math.MaxInt64,
			math.MaxInt64,
			1,
			0,
		}
		tos := []any{
			&vi,
			&vi,
			&vi8,
			&vi8,
			&vi16,
			&vi16,
			&vi32,
			&vi32,
			&vi64,
			&vi64,
			&vi64,
			&vi64,
			&uvi,
			&uvi,
			&uvi8,
			&uvi8,
			&uvi16,
			&uvi16,
			&uvi32,
			&uvi32,
			&uvi64,
			&uvi64,
			&uvi64,
			&uvi64,
			&vf32,
			&vf64,
			&vs,
			&vb,
			&vb,
		}
		for n, from := range froms {
			to := tos[n]
			t.Run(fmt.Sprintf("%v=>%T", from, to), func(t *testing.T) {
				if err := safecast.FromInt64(from, to); err != nil {
					t.Error(err)
					return
				}
			})
		}
	})

	t.Run("safecast.FromUint64", func(t *testing.T) {
		froms := []uint64{
			0,
			math.MaxInt,
			0,
			math.MaxInt8,
			0,
			math.MaxInt16,
			0,
			math.MaxInt32,
			0,
			math.MaxInt,
			0,
			math.MaxInt64,
			0,
			math.MaxUint,
			0,
			math.MaxUint8,
			0,
			math.MaxUint16,
			0,
			math.MaxUint32,
			0,
			math.MaxUint,
			0,
			math.MaxUint64,
			math.MaxUint64,
			math.MaxUint64,
			math.MaxUint64,
			1,
			0,
		}
		tos := []any{
			&vi,
			&vi,
			&vi8,
			&vi8,
			&vi16,
			&vi16,
			&vi32,
			&vi32,
			&vi64,
			&vi64,
			&vi64,
			&vi64,
			&uvi,
			&uvi,
			&uvi8,
			&uvi8,
			&uvi16,
			&uvi16,
			&uvi32,
			&uvi32,
			&uvi64,
			&uvi64,
			&uvi64,
			&uvi64,
			&vf32,
			&vf64,
			&vs,
			&vb,
			&vb,
		}
		for n, from := range froms {
			to := tos[n]
			t.Run(fmt.Sprintf("%v=>%T", from, to), func(t *testing.T) {
				if err := safecast.FromUint64(from, to); err != nil {
					t.Error(err)
					return
				}
			})
		}
	})

	t.Run("FromFloat64", func(t *testing.T) {
		froms := []float64{
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
			1,
		}
		tos := []any{
			&vi,
			&vi8,
			&vi16,
			&vi32,
			&vi64,
			&uvi,
			&uvi8,
			&uvi16,
			&uvi32,
			&uvi64,
			&vf32,
			&vf64,
			&vs,
		}
		for n, from := range froms {
			to := tos[n]
			t.Run(fmt.Sprintf("%v=>%T", from, to), func(t *testing.T) {
				if err := safecast.FromFloat64(from, to); err != nil {
					t.Error(err)
					return
				}
			})
		}
	})

	t.Run("FromString", func(t *testing.T) {
		froms := []string{
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
			strconv.Itoa(1),
		}
		tos := []any{
			&vi,
			&vi8,
			&vi16,
			&vi32,
			&vi64,
			&uvi,
			&uvi8,
			&uvi16,
			&uvi32,
			&uvi64,
			&vf32,
			&vf64,
			&vs,
			&vb,
		}
		for n, from := range froms {
			to := tos[n]
			t.Run(fmt.Sprintf("%v=>%T", from, to), func(t *testing.T) {
				if err := safecast.FromString(from, to); err != nil {
					t.Error(err)
					return
				}
			})
		}
	})

	t.Run("FromBool", func(t *testing.T) {
		froms := []bool{
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
			true,
		}
		tos := []any{
			&vi,
			&vi8,
			&vi16,
			&vi32,
			&vi64,
			&uvi,
			&uvi8,
			&uvi16,
			&uvi32,
			&uvi64,
			&vs,
			&vb,
		}
		for n, from := range froms {
			to := tos[n]
			t.Run(fmt.Sprintf("%v=>%T", from, to), func(t *testing.T) {
				if err := safecast.FromBool(from, to); err != nil {
					t.Error(err)
					return
				}
			})
		}
	})
}

func TestFromErrorCast(t *testing.T) {
	var vi int
	var vi8 int8
	var vi16 int16
	var vi32 int32
	var vi64 int64
	var uvi uint
	var uvi8 uint8
	var uvi16 uint16
	var uvi32 uint32
	var uvi64 uint64
	var vf32 float32
	var vf64 float64

	t.Run("safecast.FromInt64", func(t *testing.T) {
		froms := []int64{
			math.MaxInt8 + 1,
			math.MinInt8 - 1,
			math.MaxInt16 + 1,
			math.MinInt16 - 1,
			math.MaxInt32 + 1,
			math.MinInt32 - 1,
			-1,
			-1,
			-1,
			-1,
			-1,
		}
		tos := []any{
			&vi8,
			&vi8,
			&vi16,
			&vi16,
			&vi32,
			&vi32,
			&uvi,
			&uvi8,
			&uvi16,
			&uvi32,
			&uvi64,
		}
		for n, from := range froms {
			to := tos[n]
			t.Run(fmt.Sprintf("%v=>%T", from, to), func(t *testing.T) {
				if err := safecast.FromInt64(from, to); err == nil {
					t.Errorf("%v=>%T", from, to)
					return
				}
			})
		}
	})

	t.Run("safecast.FromUint64", func(t *testing.T) {
		froms := []uint64{
			math.MaxInt64 + 1,
			math.MaxInt8 + 1,
			math.MaxInt16 + 1,
			math.MaxInt32 + 1,
			math.MaxInt64 + 1,
			math.MaxInt8 + 1,
			math.MaxInt16 + 1,
			math.MaxInt32 + 1,
			math.MaxUint8 + 1,
			math.MaxUint16 + 1,
			math.MaxUint32 + 1,
		}
		tos := []any{
			&vi,
			&vi8,
			&vi16,
			&vi32,
			&vi64,
			&vi8,
			&vi16,
			&vi32,
			&uvi8,
			&uvi16,
			&uvi32,
		}
		for n, from := range froms {
			to := tos[n]
			t.Run(fmt.Sprintf("%v=>%T", from, to), func(t *testing.T) {
				if err := safecast.FromUint64(from, to); err == nil {
					t.Errorf("%v=>%T", from, to)
					return
				}
			})
		}
	})

	t.Run("FromString", func(t *testing.T) {
		froms := []string{
			"abc",
			"abc",
			"abc",
			"abc",
			"abc",
			"abc",
			"abc",
			"abc",
			"abc",
			"abc",
			"abc",
			"abc",
		}
		tos := []any{
			&vi,
			&vi8,
			&vi16,
			&vi32,
			&vi64,
			&uvi,
			&uvi8,
			&uvi16,
			&uvi32,
			&uvi64,
			&vf32,
			&vf64,
		}
		for n, from := range froms {
			to := tos[n]
			t.Run(fmt.Sprintf("%v=>%T", from, to), func(t *testing.T) {
				if err := safecast.FromString(from, to); err == nil {
					t.Errorf("%v=>%T", from, to)
					return
				}
			})
		}
	})
}

// TestGenericFrom tests the generic From() function
func TestGenericFrom(t *testing.T) {
	t.Run("int types", func(t *testing.T) {
		var result int64
		err := safecast.From(int(42), &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		var result2 int
		err = safecast.From(int64(42), &result2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("uint types", func(t *testing.T) {
		var result int
		err := safecast.From(uint(42), &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("float types", func(t *testing.T) {
		var result float64
		err := safecast.From(float32(42.5), &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("string types", func(t *testing.T) {
		var result int
		err := safecast.From("42", &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("bool types", func(t *testing.T) {
		var result int
		err := safecast.From(true, &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("bytes types", func(t *testing.T) {
		var result string
		err := safecast.From([]byte("hello"), &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("unsupported type", func(t *testing.T) {
		type CustomType struct{}
		custom := CustomType{}
		var result int
		err := safecast.From(custom, &result)
		if err == nil {
			t.Error("expected error for unsupported type")
		}
	})
}

// TestFromBytes tests FromBytes function
func TestFromBytes(t *testing.T) {
	t.Run("bytes to string", func(t *testing.T) {
		input := []byte("hello world")
		var result string
		err := safecast.FromBytes(input, &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		if result != "hello world" {
			t.Errorf("expected 'hello world', got '%s'", result)
		}
	})

	t.Run("bytes to bytes", func(t *testing.T) {
		input := []byte("test data")
		var result []byte
		err := safecast.FromBytes(input, &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	})

	t.Run("unsupported type", func(t *testing.T) {
		input := []byte("test")
		var result int
		err := safecast.FromBytes(input, &result)
		if err == nil {
			t.Error("expected error for unsupported type")
		}
	})
}
