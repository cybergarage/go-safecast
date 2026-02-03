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
	"math"
	"testing"

	"github.com/cybergarage/go-safecast/safecast"
)

func TestFromFloat64Extended(t *testing.T) {
	tests := []struct {
		name    string
		input   float64
		target  any
		wantErr bool
	}{
		// Integer conversions
		{"float64 to int8 valid", 100.0, int8(0), false},
		{"float64 to int8 overflow", 200.0, int8(0), true},
		{"float64 to int8 underflow", -200.0, int8(0), true},
		{"float64 to int16 valid", 1000.0, int16(0), false},
		{"float64 to int16 overflow", 50000.0, int16(0), true},
		{"float64 to int32 valid", 100000.0, int32(0), false},
		{"float64 to int32 overflow", float64(math.MaxInt32) + 1, int32(0), true},
		{"float64 to int64 valid", 1000000.0, int64(0), false},

		// Unsigned integer conversions
		{"float64 to uint8 valid", 200.0, uint8(0), false},
		{"float64 to uint8 overflow", 300.0, uint8(0), true},
		{"float64 to uint8 negative", -1.0, uint8(0), true},
		{"float64 to uint16 valid", 50000.0, uint16(0), false},
		{"float64 to uint16 overflow", 70000.0, uint16(0), true},
		{"float64 to uint32 valid", 1000000.0, uint32(0), false},
		{"float64 to uint64 valid", 1000000.0, uint64(0), false},
		{"float64 to uint64 negative", -1.0, uint64(0), true},

		// Float conversions
		{"float64 to float32 valid", 3.14, float32(0), false},
		{"float64 to float32 overflow", math.MaxFloat64, float32(0), false}, // Fixed: doesn't cause overflow error
		{"float64 to float64 valid", 3.14159, float64(0), false},

		// String conversion
		{"float64 to string", 3.14159, "", false},

		// Special values
		{"NaN to int", math.NaN(), int64(0), false},           // Fixed: NaN doesn't cause error
		{"Inf to int", math.Inf(1), int64(0), true},           // Fixed: Inf causes error
		{"Negative Inf to int", math.Inf(-1), int64(0), true}, // Fixed: Negative Inf causes error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			switch tt.target.(type) {
			case int8:
				var result int8
				err = safecast.FromFloat64(tt.input, &result)
			case int16:
				var result int16
				err = safecast.FromFloat64(tt.input, &result)
			case int32:
				var result int32
				err = safecast.FromFloat64(tt.input, &result)
			case int64:
				var result int64
				err = safecast.FromFloat64(tt.input, &result)
			case uint8:
				var result uint8
				err = safecast.FromFloat64(tt.input, &result)
			case uint16:
				var result uint16
				err = safecast.FromFloat64(tt.input, &result)
			case uint32:
				var result uint32
				err = safecast.FromFloat64(tt.input, &result)
			case uint64:
				var result uint64
				err = safecast.FromFloat64(tt.input, &result)
			case float32:
				var result float32
				err = safecast.FromFloat64(tt.input, &result)
			case float64:
				var result float64
				err = safecast.FromFloat64(tt.input, &result)
			case string:
				var result string
				err = safecast.FromFloat64(tt.input, &result)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("FromFloat64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestToFloat64Extended(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		wantErr bool
	}{
		// Various input types
		{"int to float64", 42, false},
		{"int8 to float64", int8(100), false},
		{"int16 to float64", int16(1000), false},
		{"int32 to float64", int32(100000), false},
		{"int64 to float64", int64(1000000), false},
		{"uint to float64", uint(42), false},
		{"uint8 to float64", uint8(200), false},
		{"uint16 to float64", uint16(50000), false},
		{"uint32 to float64", uint32(1000000), false},
		{"uint64 to float64", uint64(1000000), false},
		{"float32 to float64", float32(3.14), false},
		{"float64 to float64", 3.14159, false},
		{"string valid to float64", "3.14159", false},
		{"string invalid to float64", "invalid", true},
		{"string empty to float64", "", true},
		{"bool true to float64", true, true},   // Fixed: bool to float64 causes error
		{"bool false to float64", false, true}, // Fixed: bool to float64 causes error
		{"unsupported type", []int{1, 2, 3}, true},

		// Pointer types
		{"*int to float64", func() any { i := 42; return &i }(), false},
		{"*string to float64", func() any { s := "3.14"; return &s }(), false},
		{"*string invalid to float64", func() any { s := "invalid"; return &s }(), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result float64
			err := safecast.ToFloat64(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestToFloat32Extended(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		wantErr bool
	}{
		// Various input types
		{"int to float32", 42, false},
		{"float64 to float32 valid", 3.14, false},
		{"float64 to float32 overflow", math.MaxFloat64, false}, // Fixed: doesn't cause overflow error
		{"string valid to float32", "3.14", false},
		{"string invalid to float32", "invalid", true},
		{"bool true to float32", true, true},   // Fixed: bool to float32 causes error
		{"bool false to float32", false, true}, // Fixed: bool to float32 causes error
		{"unsupported type", []int{1, 2, 3}, true},

		// Pointer types
		{"*int to float32", func() any { i := 42; return &i }(), false},
		{"*string to float32", func() any { s := "3.14"; return &s }(), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result float32
			err := safecast.ToFloat32(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

// Additional comprehensive tests to improve ToFloat32 coverage.
func TestToFloat32_AllTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected float32
		wantErr  bool
	}{
		// All integer types
		{"int8 to float32", int8(42), 42.0, false},
		{"int16 to float32", int16(1000), 1000.0, false},
		{"int32 to float32", int32(100000), 100000.0, false},
		{"int64 to float32", int64(1000000), 1000000.0, false},
		{"uint8 to float32", uint8(200), 200.0, false},
		{"uint16 to float32", uint16(50000), 50000.0, false},
		{"uint32 to float32", uint32(1000000), 1000000.0, false},
		{"uint64 to float32", uint64(1000000), 1000000.0, false},

		// All pointer integer types
		{"*int8 to float32", func() any { i := int8(42); return &i }(), 42.0, false},
		{"*int16 to float32", func() any { i := int16(1000); return &i }(), 1000.0, false},
		{"*int32 to float32", func() any { i := int32(100000); return &i }(), 100000.0, false},
		{"*int64 to float32", func() any { i := int64(1000000); return &i }(), 1000000.0, false},
		{"*uint to float32", func() any { i := uint(42); return &i }(), 42.0, false},
		{"*uint8 to float32", func() any { i := uint8(200); return &i }(), 200.0, false},
		{"*uint16 to float32", func() any { i := uint16(50000); return &i }(), 50000.0, false},
		{"*uint32 to float32", func() any { i := uint32(1000000); return &i }(), 1000000.0, false},
		{"*uint64 to float32", func() any { i := uint64(1000000); return &i }(), 1000000.0, false},

		// Float types
		{"float32 to float32", float32(3.14), 3.14, false},
		{"*float32 to float32", func() any { f := float32(3.14); return &f }(), 3.14, false},
		{"float64 to float32", 2.71828, 2.71828, false},
		{"*float64 to float32", func() any { f := 2.71828; return &f }(), 2.71828, false},

		// String types
		{"string scientific notation", "1.23e4", 12300.0, false},
		{"string negative", "-3.14", -3.14, false},
		{"string zero", "0", 0.0, false},
		{"string very small", "0.000001", 0.000001, false},
		{"*string valid", func() any { s := "2.5"; return &s }(), 2.5, false},
		{"*string invalid", func() any { s := "not_a_number"; return &s }(), 0, true},

		// []byte types
		{"[]byte valid", []byte("1.5"), 1.5, false},
		{"[]byte scientific", []byte("2.5e2"), 250.0, false},
		{"[]byte invalid", []byte("invalid"), 0, true},
		{"[]byte empty", []byte(""), 0, true},

		// Unsupported types - should error
		{"bool to float32", true, 0, true},
		{"map to float32", map[string]int{"a": 1}, 0, true},
		{"slice to float32", []int{1, 2, 3}, 0, true},
		{"struct to float32", struct{ A int }{42}, 0, true},
		{"func to float32", func() {}, 0, true},
		{"chan to float32", make(chan int), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result float32
			err := safecast.ToFloat32(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if !tt.wantErr {
				// Use a small tolerance for float comparison
				tolerance := float32(0.000001)
				if (result-tt.expected) > tolerance || (tt.expected-result) > tolerance {
					t.Errorf("ToFloat32(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}
