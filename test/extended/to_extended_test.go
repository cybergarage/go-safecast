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

// Package test provides extended test coverage specifically for to.go functions
// This file focuses on testing the generic To function and improving coverage
// of all ToXXX functions to ensure comprehensive type handling

package test

import (
	"math"
	"testing"
	"time"

	"github.com/cybergarage/go-safecast/safecast"
)

// TestTo_ComprehensiveTypeHandling provides comprehensive test coverage for To function
// This test covers all type cases in the To function switch statement to achieve maximum coverage
// Focus: to.go generic To function with complete type branch coverage
func TestTo_ComprehensiveTypeHandling(t *testing.T) {
	tests := []struct {
		name     string
		from     interface{}
		to       interface{}
		expected interface{}
		wantErr  bool
	}{
		// *int - comprehensive coverage
		{"int to *int", 42, func() *int { var i int; return &i }(), 42, false},
		{"string to *int", "123", func() *int { var i int; return &i }(), 123, false},
		{"bool true to *int", true, func() *int { var i int; return &i }(), 1, false},
		{"bool false to *int", false, func() *int { var i int; return &i }(), 0, false},
		{"float64 to *int", 3.14, func() *int { var i int; return &i }(), 3, false},
		{"invalid string to *int", "invalid", func() *int { var i int; return &i }(), 0, true},

		// *int8 - comprehensive coverage
		{"int8 to *int8", int8(100), func() *int8 { var i int8; return &i }(), int8(100), false},
		{"int to *int8", 50, func() *int8 { var i int8; return &i }(), int8(50), false},
		{"string to *int8", "75", func() *int8 { var i int8; return &i }(), int8(75), false},
		{"overflow to *int8", 300, func() *int8 { var i int8; return &i }(), int8(0), true},

		// *int16 - comprehensive coverage
		{"int16 to *int16", int16(1000), func() *int16 { var i int16; return &i }(), int16(1000), false},
		{"int to *int16", 500, func() *int16 { var i int16; return &i }(), int16(500), false},
		{"string to *int16", "750", func() *int16 { var i int16; return &i }(), int16(750), false},
		{"overflow to *int16", 40000, func() *int16 { var i int16; return &i }(), int16(0), true},

		// *int32 - comprehensive coverage
		{"int32 to *int32", int32(50000), func() *int32 { var i int32; return &i }(), int32(50000), false},
		{"int to *int32", 25000, func() *int32 { var i int32; return &i }(), int32(25000), false},
		{"string to *int32", "35000", func() *int32 { var i int32; return &i }(), int32(35000), false},
		{"float64 to *int32", 12345.67, func() *int32 { var i int32; return &i }(), int32(12345), false},

		// *int64 - comprehensive coverage
		{"int64 to *int64", int64(123456789), func() *int64 { var i int64; return &i }(), int64(123456789), false},
		{"int to *int64", 987654321, func() *int64 { var i int64; return &i }(), int64(987654321), false},
		{"string to *int64", "555666777", func() *int64 { var i int64; return &i }(), int64(555666777), false},
		{"float64 to *int64", 999.123, func() *int64 { var i int64; return &i }(), int64(999), false},

		// *uint - comprehensive coverage
		{"uint to *uint", uint(50000), func() *uint { var i uint; return &i }(), uint(50000), false},
		{"int to *uint", 25000, func() *uint { var i uint; return &i }(), uint(25000), false},
		{"string to *uint", "35000", func() *uint { var i uint; return &i }(), uint(35000), false},
		{"negative int to *uint", -100, func() *uint { var i uint; return &i }(), uint(0), true},

		// *uint8 - comprehensive coverage
		{"uint8 to *uint8", uint8(200), func() *uint8 { var i uint8; return &i }(), uint8(200), false},
		{"int to *uint8", 150, func() *uint8 { var i uint8; return &i }(), uint8(150), false},
		{"string to *uint8", "100", func() *uint8 { var i uint8; return &i }(), uint8(100), false},
		{"overflow to *uint8", 300, func() *uint8 { var i uint8; return &i }(), uint8(0), true},

		// *uint16 - comprehensive coverage
		{"uint16 to *uint16", uint16(60000), func() *uint16 { var i uint16; return &i }(), uint16(60000), false},
		{"int to *uint16", 30000, func() *uint16 { var i uint16; return &i }(), uint16(30000), false},
		{"string to *uint16", "45000", func() *uint16 { var i uint16; return &i }(), uint16(45000), false},
		{"overflow to *uint16", 70000, func() *uint16 { var i uint16; return &i }(), uint16(0), true},

		// *uint32 - comprehensive coverage
		{"uint32 to *uint32", uint32(4000000000), func() *uint32 { var i uint32; return &i }(), uint32(4000000000), false},
		{"int to *uint32", 2000000000, func() *uint32 { var i uint32; return &i }(), uint32(2000000000), false},
		{"string to *uint32", "3000000000", func() *uint32 { var i uint32; return &i }(), uint32(3000000000), false},
		{"negative int to *uint32", -1000, func() *uint32 { var i uint32; return &i }(), uint32(0), true},

		// *uint64 - comprehensive coverage
		{"uint64 to *uint64", uint64(9000000000000000000), func() *uint64 { var i uint64; return &i }(), uint64(9000000000000000000), false},
		{"int to *uint64", 5000000000, func() *uint64 { var i uint64; return &i }(), uint64(5000000000), false},
		{"string to *uint64", "7000000000000000000", func() *uint64 { var i uint64; return &i }(), uint64(7000000000000000000), false},
		{"negative int to *uint64", -5000, func() *uint64 { var i uint64; return &i }(), uint64(0), true},

		// *float32 - comprehensive coverage
		{"float32 to *float32", float32(3.14159), func() *float32 { var f float32; return &f }(), float32(3.14159), false},
		{"int to *float32", 42, func() *float32 { var f float32; return &f }(), float32(42.0), false},
		{"string to *float32", "2.71828", func() *float32 { var f float32; return &f }(), float32(2.71828), false},
		{"bool true to *float32", true, func() *float32 { var f float32; return &f }(), float32(0), true},   // bool to float32 not supported
		{"bool false to *float32", false, func() *float32 { var f float32; return &f }(), float32(0), true}, // bool to float32 not supported
		{"invalid string to *float32", "invalid", func() *float32 { var f float32; return &f }(), float32(0), true},

		// *float64 - comprehensive coverage
		{"float64 to *float64", 2.718281828, func() *float64 { var f float64; return &f }(), 2.718281828, false},
		{"int to *float64", 42, func() *float64 { var f float64; return &f }(), 42.0, false},
		{"string to *float64", "3.141592653589793", func() *float64 { var f float64; return &f }(), 3.141592653589793, false},
		{"bool true to *float64", true, func() *float64 { var f float64; return &f }(), 0.0, true},   // bool to float64 not supported
		{"bool false to *float64", false, func() *float64 { var f float64; return &f }(), 0.0, true}, // bool to float64 not supported
		{"NaN to *float64", math.NaN(), func() *float64 { var f float64; return &f }(), math.NaN(), false},
		{"Inf to *float64", math.Inf(1), func() *float64 { var f float64; return &f }(), math.Inf(1), false},
		{"invalid string to *float64", "invalid", func() *float64 { var f float64; return &f }(), 0.0, true},

		// *string - comprehensive coverage
		{"string to *string", "hello world", func() *string { var s string; return &s }(), "hello world", false},
		{"int to *string", 12345, func() *string { var s string; return &s }(), "12345", false},
		{"float64 to *string", 3.14159, func() *string { var s string; return &s }(), "3.14159", false},
		{"bool true to *string", true, func() *string { var s string; return &s }(), "true", false},
		{"bool false to *string", false, func() *string { var s string; return &s }(), "false", false},
		{"[]byte to *string", []byte("bytes to string"), func() *string { var s string; return &s }(), "bytes to string", false},
		{"time to *string", time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC), func() *string { var s string; return &s }(), "2023-01-01 12:00:00 +0000 UTC", false},

		// *bool - comprehensive coverage
		{"bool to *bool", true, func() *bool { var b bool; return &b }(), true, false},
		{"int 1 to *bool", 1, func() *bool { var b bool; return &b }(), true, false},
		{"int 0 to *bool", 0, func() *bool { var b bool; return &b }(), false, false},
		{"int 42 to *bool", 42, func() *bool { var b bool; return &b }(), true, false},
		{"string true to *bool", "true", func() *bool { var b bool; return &b }(), true, false},
		{"string false to *bool", "false", func() *bool { var b bool; return &b }(), false, false},
		{"string 1 to *bool", "1", func() *bool { var b bool; return &b }(), true, false},
		{"string 0 to *bool", "0", func() *bool { var b bool; return &b }(), false, false},
		{"float64 0.0 to *bool", 0.0, func() *bool { var b bool; return &b }(), false, true},  // float to bool not supported
		{"float64 3.14 to *bool", 3.14, func() *bool { var b bool; return &b }(), true, true}, // float to bool not supported
		{"[]byte true to *bool", []byte("true"), func() *bool { var b bool; return &b }(), true, false},
		{"[]byte false to *bool", []byte("false"), func() *bool { var b bool; return &b }(), false, false},

		// *[]byte - comprehensive coverage
		{"[]byte to *[]byte", []byte("hello"), func() *[]byte { var b []byte; return &b }(), []byte("hello"), false},
		{"string to *[]byte", "world", func() *[]byte { var b []byte; return &b }(), []byte("world"), false},
		{"int to *[]byte", 12345, func() *[]byte { var b []byte; return &b }(), []byte("12345"), true},        // int to []byte not supported
		{"float64 to *[]byte", 3.14, func() *[]byte { var b []byte; return &b }(), []byte("3.14"), true},      // float to []byte not supported
		{"bool true to *[]byte", true, func() *[]byte { var b []byte; return &b }(), []byte("true"), true},    // bool to []byte not supported
		{"bool false to *[]byte", false, func() *[]byte { var b []byte; return &b }(), []byte("false"), true}, // bool to []byte not supported

		// *time.Time - comprehensive coverage
		{"time.Time to *time.Time", time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC), func() *time.Time { var t time.Time; return &t }(), time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC), false},
		{"string RFC3339 to *time.Time", "2023-05-15T14:30:00Z", func() *time.Time { var t time.Time; return &t }(), time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC), false},
		{"int unix timestamp to *time.Time", int64(1684158600), func() *time.Time { var t time.Time; return &t }(), time.Unix(1684158600, 0).UTC(), true},                 // int to time not supported
		{"float64 unix timestamp to *time.Time", float64(1684158600.5), func() *time.Time { var t time.Time; return &t }(), time.Unix(1684158600, 500000000).UTC(), true}, // float to time not supported
		{"[]byte RFC3339 to *time.Time", []byte("2023-05-15T14:30:00Z"), func() *time.Time { var t time.Time; return &t }(), time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC), false},
		{"bool true to *time.Time", true, func() *time.Time { var t time.Time; return &t }(), time.Unix(1, 0).UTC(), true},   // bool to time not supported
		{"bool false to *time.Time", false, func() *time.Time { var t time.Time; return &t }(), time.Unix(0, 0).UTC(), true}, // bool to time not supported
		{"invalid string to *time.Time", "invalid-time", func() *time.Time { var t time.Time; return &t }(), time.Time{}, true},

		// default case - unsupported types
		{"map unsupported", map[string]int{"key": 1}, []string{"unsupported"}, nil, true},
		{"slice unsupported", []int{1, 2, 3}, []string{"unsupported"}, nil, true},
		{"struct unsupported", struct{ A int }{42}, []string{"unsupported"}, nil, true},
		{"func unsupported", func() {}, []string{"unsupported"}, nil, true},
		{"chan unsupported", make(chan int), []string{"unsupported"}, nil, true},
		{"interface{} nil", nil, []string{"unsupported"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := safecast.To(tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("To(%v, %T) error = %v, wantErr %v", tt.from, tt.to, err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.expected != nil {
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
					// Special handling for NaN comparison
					if math.IsNaN(tt.expected.(float64)) {
						if !math.IsNaN(*ptr) {
							t.Errorf("To(%v, %T) = %v, want NaN", tt.from, tt.to, *ptr)
						}
						return
					}
					// Special handling for Inf comparison
					if math.IsInf(tt.expected.(float64), 0) {
						if !math.IsInf(*ptr, 0) || math.Signbit(tt.expected.(float64)) != math.Signbit(*ptr) {
							t.Errorf("To(%v, %T) = %v, want %v", tt.from, tt.to, *ptr, tt.expected)
						}
						return
					}
					actualResult = *ptr
				case *string:
					actualResult = *ptr
				case *bool:
					actualResult = *ptr
				case *[]byte:
					actualResult = *ptr
				case *time.Time:
					actualResult = *ptr
				}

				// Special comparison for []byte
				if expectedBytes, ok := tt.expected.([]byte); ok {
					if actualBytes, ok := actualResult.([]byte); ok {
						if string(expectedBytes) != string(actualBytes) {
							t.Errorf("To(%v, %T) = %v, want %v", tt.from, tt.to, actualResult, tt.expected)
						}
						return
					}
				}

				// Special comparison for time.Time
				if expectedTime, ok := tt.expected.(time.Time); ok {
					if actualTime, ok := actualResult.(time.Time); ok {
						if !expectedTime.Equal(actualTime) {
							t.Errorf("To(%v, %T) = %v, want %v", tt.from, tt.to, actualResult, tt.expected)
						}
						return
					}
				}

				if actualResult != tt.expected {
					t.Errorf("To(%v, %T) = %v, want %v", tt.from, tt.to, actualResult, tt.expected)
				}
			}
		})
	}
}

// TestTo_EdgeCases provides additional coverage for edge cases in To function
// Ensures all error paths and boundary conditions are tested
func TestTo_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		from    interface{}
		to      interface{}
		wantErr bool
	}{
		// Edge cases for numeric overflows
		{"large int64 to *int8", int64(1000), func() *int8 { var i int8; return &i }(), true},
		{"large uint64 to *uint8", uint64(1000), func() *uint8 { var i uint8; return &i }(), true},
		{"negative int to *uint", -1, func() *uint { var i uint; return &i }(), true},
		{"float overflow to *float32", math.MaxFloat64, func() *float32 { var f float32; return &f }(), false}, // doesn't overflow, just inf

		// Edge cases for time parsing
		{"empty string to *time.Time", "", func() *time.Time { var t time.Time; return &t }(), true},
		{"malformed time string to *time.Time", "not-a-time", func() *time.Time { var t time.Time; return &t }(), true},

		// Edge cases for bool parsing
		{"empty string to *bool", "", func() *bool { var b bool; return &b }(), true}, // empty string to bool fails
		{"invalid string to *bool", "maybe", func() *bool { var b bool; return &b }(), true},

		// Edge cases for numeric string parsing
		{"empty string to *int", "", func() *int { var i int; return &i }(), true},
		{"non-numeric string to *float64", "not-a-number", func() *float64 { var f float64; return &f }(), true},

		// Unsupported target types with various sources
		{"int to unsupported", 42, struct{}{}, true},
		{"string to unsupported", "test", make(chan int), true},
		{"bool to unsupported", true, map[string]int{}, true},
		{"float to unsupported", 3.14, []int{}, true},
		{"[]byte to unsupported", []byte("test"), func() {}, true},
		{"time to unsupported", time.Now(), complex64(1 + 2i), true},

		// nil and zero values
		{"nil to *int", nil, func() *int { var i int; return &i }(), true},
		{"nil to *string", nil, func() *string { var s string; return &s }(), false}, // nil to string succeeds with empty string
		{"nil to *bool", nil, func() *bool { var b bool; return &b }(), true},
		{"nil to *time.Time", nil, func() *time.Time { var t time.Time; return &t }(), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := safecast.To(tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("To(%v, %T) error = %v, wantErr %v", tt.from, tt.to, err, tt.wantErr)
			}
		})
	}
}
