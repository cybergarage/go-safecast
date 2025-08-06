// Copyright (C) 2023 Satoshi Konno. All rights reserved.
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

// Package test provides extended test coverage specifically for from.go functions
// This file focuses on testing the generic From function and improving coverage
// of all FromXXX functions to ensure comprehensive type handling

package test

import (
	"math"
	"testing"

	"github.com/cybergarage/go-safecast/safecast"
)

// TestFrom_ComprehensiveTypeHandling provides comprehensive test coverage for From function
// This test covers all type cases in the From function switch statement to achieve maximum coverage
// Focus: from.go generic From function with complete type branch coverage.
func TestFrom_ComprehensiveTypeHandling(t *testing.T) {
	tests := []struct {
		name     string
		from     interface{}
		to       interface{}
		expected interface{}
		wantErr  bool
	}{
		// int and *int - comprehensive coverage
		{"int to int32", 42, func() *int32 { var i int32; return &i }(), int32(42), false},
		{"*int to int32", func() *int { i := 42; return &i }(), func() *int32 { var i int32; return &i }(), int32(42), false},
		{"int overflow", 3000000000, func() *int32 { var i int32; return &i }(), int32(0), true},

		// int8 and *int8 - comprehensive coverage
		{"int8 to int32", int8(100), func() *int32 { var i int32; return &i }(), int32(100), false},
		{"*int8 to int32", func() *int8 { i := int8(100); return &i }(), func() *int32 { var i int32; return &i }(), int32(100), false},
		{"int8 to uint8 negative", int8(-1), func() *uint8 { var i uint8; return &i }(), uint8(0), true},

		// int16 and *int16 - comprehensive coverage
		{"int16 to int32", int16(1000), func() *int32 { var i int32; return &i }(), int32(1000), false},
		{"*int16 to int32", func() *int16 { i := int16(1000); return &i }(), func() *int32 { var i int32; return &i }(), int32(1000), false},
		{"int16 overflow to int8", int16(300), func() *int8 { var i int8; return &i }(), int8(0), true},

		// int32 and *int32 - comprehensive coverage
		{"int32 to int64", int32(50000), func() *int64 { var i int64; return &i }(), int64(50000), false},
		{"*int32 to int64", func() *int32 { i := int32(50000); return &i }(), func() *int64 { var i int64; return &i }(), int64(50000), false},
		{"int32 overflow to int16", int32(100000), func() *int16 { var i int16; return &i }(), int16(0), true},

		// int64 and *int64 - comprehensive coverage
		{"int64 to string", int64(123456), func() *string { var s string; return &s }(), "123456", false},
		{"*int64 to string", func() *int64 { i := int64(123456); return &i }(), func() *string { var s string; return &s }(), "123456", false},
		{"int64 overflow to int32", int64(3000000000), func() *int32 { var i int32; return &i }(), int32(0), true},

		// uint and *uint - comprehensive coverage
		{"uint to uint64", uint(50000), func() *uint64 { var i uint64; return &i }(), uint64(50000), false},
		{"*uint to uint64", func() *uint { i := uint(50000); return &i }(), func() *uint64 { var i uint64; return &i }(), uint64(50000), false},
		{"uint overflow to uint8", uint(300), func() *uint8 { var i uint8; return &i }(), uint8(0), true},

		// uint8 and *uint8 - comprehensive coverage
		{"uint8 to uint32", uint8(200), func() *uint32 { var i uint32; return &i }(), uint32(200), false},
		{"*uint8 to uint32", func() *uint8 { i := uint8(200); return &i }(), func() *uint32 { var i uint32; return &i }(), uint32(200), false},
		{"uint8 overflow", uint8(255), func() *int8 { var i int8; return &i }(), int8(0), true},

		// uint16 and *uint16 - comprehensive coverage
		{"uint16 to uint32", uint16(60000), func() *uint32 { var i uint32; return &i }(), uint32(60000), false},
		{"*uint16 to uint32", func() *uint16 { i := uint16(60000); return &i }(), func() *uint32 { var i uint32; return &i }(), uint32(60000), false},
		{"uint16 overflow to int16", uint16(40000), func() *int16 { var i int16; return &i }(), int16(0), true},

		// uint32 and *uint32 - comprehensive coverage
		{"uint32 to uint64", uint32(4000000000), func() *uint64 { var i uint64; return &i }(), uint64(4000000000), false},
		{"*uint32 to uint64", func() *uint32 { i := uint32(4000000000); return &i }(), func() *uint64 { var i uint64; return &i }(), uint64(4000000000), false},
		{"uint32 overflow to int32", uint32(3000000000), func() *int32 { var i int32; return &i }(), int32(0), true},

		// uint64 and *uint64 - comprehensive coverage
		{"uint64 to string", uint64(9000000000000000000), func() *string { var s string; return &s }(), "9000000000000000000", false},
		{"*uint64 to string", func() *uint64 { i := uint64(9000000000000000000); return &i }(), func() *string { var s string; return &s }(), "9000000000000000000", false},
		{"uint64 overflow to int64", uint64(10000000000000000000), func() *int64 { var i int64; return &i }(), int64(0), true},

		// float32 and *float32 - comprehensive coverage
		{"float32 to float64", float32(3.14159), func() *float64 { var f float64; return &f }(), float64(float32(3.14159)), false},
		{"*float32 to float64", func() *float32 { f := float32(3.14159); return &f }(), func() *float64 { var f float64; return &f }(), float64(float32(3.14159)), false},
		{"float32 to int32", float32(12345.67), func() *int32 { var i int32; return &i }(), int32(12345), false},

		// float64 and *float64 - comprehensive coverage
		{"float64 to string", 2.71828, func() *string { var s string; return &s }(), "2.71828", false},
		{"*float64 to string", func() *float64 { f := 2.71828; return &f }(), func() *string { var s string; return &s }(), "2.71828", false},
		// {"float64 NaN to int32", math.NaN(), func() *int32 { var i int32; return &i }(), int32(0), false}, // Removed - platform dependent
		{"float64 +Inf to int32", math.Inf(1), func() *int32 { var i int32; return &i }(), int32(0), true},
		{"float64 -Inf to int32", math.Inf(-1), func() *int32 { var i int32; return &i }(), int32(0), true},

		// string and *string - comprehensive coverage
		{"string to int64", "987654321", func() *int64 { var i int64; return &i }(), int64(987654321), false},
		{"*string to int64", func() *string { s := "987654321"; return &s }(), func() *int64 { var i int64; return &i }(), int64(987654321), false},
		{"string to float64", "3.14159", func() *float64 { var f float64; return &f }(), 3.14159, false},
		{"string invalid to int", "invalid", func() *int { var i int; return &i }(), 0, true},
		{"string empty to int", "", func() *int { var i int; return &i }(), 0, true},

		// bool and *bool - comprehensive coverage
		{"bool true to int", true, func() *int { var i int; return &i }(), 1, false},
		{"bool false to int", false, func() *int { var i int; return &i }(), 0, false},
		{"*bool true to int", func() *bool { b := true; return &b }(), func() *int { var i int; return &i }(), 1, false},
		{"*bool false to int", func() *bool { b := false; return &b }(), func() *int { var i int; return &i }(), 0, false},
		{"bool to string", true, func() *string { var s string; return &s }(), "true", false},

		// []byte - only supports string and []byte conversion
		{"[]byte to string", []byte("hello world"), func() *string { var s string; return &s }(), "hello world", false},
		{"[]byte to int unsupported", []byte("12345"), func() *int { var i int; return &i }(), 0, true},
		{"[]byte unicode to string", []byte("こんにちは"), func() *string { var s string; return &s }(), "こんにちは", false},

		// Unsupported types - default case coverage
		{"map unsupported", map[string]int{"key": 1}, func() *int { var i int; return &i }(), 0, true},
		{"slice unsupported", []int{1, 2, 3}, func() *int { var i int; return &i }(), 0, true},
		{"struct unsupported", struct{ A int }{42}, func() *int { var i int; return &i }(), 0, true},
		{"func unsupported", func() {}, func() *int { var i int; return &i }(), 0, true},
		{"chan unsupported", make(chan int), func() *int { var i int; return &i }(), 0, true},
		{"interface{} nil", nil, func() *int { var i int; return &i }(), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := safecast.From(tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("From(%v, %T) error = %v, wantErr %v", tt.from, tt.to, err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Get the actual result by dereferencing the pointer
				var actualResult interface{}
				switch ptr := tt.to.(type) {
				case *int:
					actualResult = *ptr
				case *int8:
					actualResult = *ptr
				case *int16:
					actualResult = *ptr
				case *int32:
					actualResult = *ptr
				case *int64:
					actualResult = *ptr
				case *uint:
					actualResult = *ptr
				case *uint8:
					actualResult = *ptr
				case *uint16:
					actualResult = *ptr
				case *uint32:
					actualResult = *ptr
				case *uint64:
					actualResult = *ptr
				case *float32:
					actualResult = *ptr
				case *float64:
					actualResult = *ptr
				case *string:
					actualResult = *ptr
				case *bool:
					actualResult = *ptr
				}

				if actualResult != tt.expected {
					t.Errorf("From(%v, %T) = %v, want %v", tt.from, tt.to, actualResult, tt.expected)
				}
			}
		})
	}
}

// TestFromBool_EdgeCases provides additional coverage for FromBool function
// Ensures all boolean conversion paths are tested to improve from.go coverage.
func TestFromBool_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		input   bool
		target  interface{}
		wantErr bool
	}{
		// Basic conversions for true
		{"bool true to int", true, func() *int { var i int; return &i }(), false},
		{"bool true to int8", true, func() *int8 { var i int8; return &i }(), false},
		{"bool true to int16", true, func() *int16 { var i int16; return &i }(), false},
		{"bool true to int32", true, func() *int32 { var i int32; return &i }(), false},
		{"bool true to int64", true, func() *int64 { var i int64; return &i }(), false},
		{"bool true to uint", true, func() *uint { var i uint; return &i }(), false},
		{"bool true to uint8", true, func() *uint8 { var i uint8; return &i }(), false},
		{"bool true to uint16", true, func() *uint16 { var i uint16; return &i }(), false},
		{"bool true to uint32", true, func() *uint32 { var i uint32; return &i }(), false},
		{"bool true to uint64", true, func() *uint64 { var i uint64; return &i }(), false},
		{"bool true to string", true, func() *string { var s string; return &s }(), false},
		{"bool true to bool", true, func() *bool { var b bool; return &b }(), false},

		// Basic conversions for false
		{"bool false to int", false, func() *int { var i int; return &i }(), false},
		{"bool false to int8", false, func() *int8 { var i int8; return &i }(), false},
		{"bool false to int16", false, func() *int16 { var i int16; return &i }(), false},
		{"bool false to int32", false, func() *int32 { var i int32; return &i }(), false},
		{"bool false to int64", false, func() *int64 { var i int64; return &i }(), false},
		{"bool false to uint", false, func() *uint { var i uint; return &i }(), false},
		{"bool false to uint8", false, func() *uint8 { var i uint8; return &i }(), false},
		{"bool false to uint16", false, func() *uint16 { var i uint16; return &i }(), false},
		{"bool false to uint32", false, func() *uint32 { var i uint32; return &i }(), false},
		{"bool false to uint64", false, func() *uint64 { var i uint64; return &i }(), false},
		{"bool false to string", false, func() *string { var s string; return &s }(), false},
		{"bool false to bool", false, func() *bool { var b bool; return &b }(), false},

		// Float conversions (should error)
		{"bool to float32", true, func() *float32 { var f float32; return &f }(), true},
		{"bool to float64", false, func() *float64 { var f float64; return &f }(), true},

		// Unsupported target types
		{"bool to unsupported", true, []string{"test"}, true},
		{"bool to map", false, map[string]int{}, true},
		{"bool to struct", true, struct{ A int }{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := safecast.FromBool(tt.input, tt.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromBool(%v, %T) error = %v, wantErr %v", tt.input, tt.target, err, tt.wantErr)
			}
		})
	}
}
