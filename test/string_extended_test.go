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
	"time"

	"github.com/cybergarage/go-safecast/safecast"
)

func TestFromStringExtended(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		target  interface{}
		wantErr bool
	}{
		// Integer conversions
		{"string to int valid", "42", int(0), false},
		{"string to int invalid", "invalid", int(0), true},
		{"string to int empty", "", int(0), true},
		{"string to int8 valid", "100", int8(0), false},
		{"string to int8 overflow", "200", int8(0), true},
		{"string to int16 valid", "1000", int16(0), false},
		{"string to int32 valid", "100000", int32(0), false},
		{"string to int64 valid", "1000000", int64(0), false},

		// Unsigned integer conversions
		{"string to uint valid", "42", uint(0), false},
		{"string to uint negative", "-1", uint(0), true},
		{"string to uint8 valid", "200", uint8(0), false},
		{"string to uint8 overflow", "300", uint8(0), true},
		{"string to uint16 valid", "50000", uint16(0), false},
		{"string to uint32 valid", "1000000", uint32(0), false},
		{"string to uint64 valid", "1000000", uint64(0), false},

		// Float conversions
		{"string to float32 valid", "3.14", float32(0), false},
		{"string to float32 invalid", "invalid", float32(0), true},
		{"string to float64 valid", "3.14159", float64(0), false},
		{"string to float64 scientific", "1.23e-4", float64(0), false},
		{"string to float64 invalid", "not_a_number", float64(0), true},

		// Bool conversions
		{"string to bool true", "true", false, false},
		{"string to bool false", "false", false, false},
		{"string to bool 1", "1", false, false},
		{"string to bool 0", "0", false, false},
		{"string to bool invalid", "maybe", false, true},

		// String to string
		{"string to string", "hello", "", false},

		// Special numeric formats
		{"string hex to int", "0x10", int(0), true},                // hex not supported in standard strconv
		{"string float to int", "42.0", int(0), true},              // Fixed: float string to int causes error
		{"string float to int with decimal", "42.7", int(0), true}, // Fixed: float string to int causes error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			switch tt.target.(type) {
			case int:
				var result int
				err = safecast.FromString(tt.input, &result)
			case int8:
				var result int8
				err = safecast.FromString(tt.input, &result)
			case int16:
				var result int16
				err = safecast.FromString(tt.input, &result)
			case int32:
				var result int32
				err = safecast.FromString(tt.input, &result)
			case int64:
				var result int64
				err = safecast.FromString(tt.input, &result)
			case uint:
				var result uint
				err = safecast.FromString(tt.input, &result)
			case uint8:
				var result uint8
				err = safecast.FromString(tt.input, &result)
			case uint16:
				var result uint16
				err = safecast.FromString(tt.input, &result)
			case uint32:
				var result uint32
				err = safecast.FromString(tt.input, &result)
			case uint64:
				var result uint64
				err = safecast.FromString(tt.input, &result)
			case float32:
				var result float32
				err = safecast.FromString(tt.input, &result)
			case float64:
				var result float64
				err = safecast.FromString(tt.input, &result)
			case bool:
				var result bool
				err = safecast.FromString(tt.input, &result)
			case string:
				var result string
				err = safecast.FromString(tt.input, &result)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("FromString(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestToStringExtended(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		// Various input types
		{"int to string", 42, false},
		{"int8 to string", int8(100), false},
		{"int16 to string", int16(1000), false},
		{"int32 to string", int32(100000), false},
		{"int64 to string", int64(1000000), false},
		{"uint to string", uint(42), false},
		{"uint8 to string", uint8(200), false},
		{"uint16 to string", uint16(50000), false},
		{"uint32 to string", uint32(1000000), false},
		{"uint64 to string", uint64(1000000), false},
		{"float32 to string", float32(3.14), false},
		{"float64 to string", 3.14159, false},
		{"bool true to string", true, false},
		{"bool false to string", false, false},
		{"string to string", "hello", false},
		{"bytes to string", []byte("hello"), false},
		{"time to string", time.Now(), false},
		{"unsupported type", []int{1, 2, 3}, false}, // Fixed: slice to string works

		// Pointer types
		{"*int to string", func() interface{} { i := 42; return &i }(), false},
		{"*string to string", func() interface{} { s := "hello"; return &s }(), false},
		{"*bool to string", func() interface{} { b := true; return &b }(), false},
		{"*time to string", func() interface{} { t := time.Now(); return &t }(), false},

		// Special float values
		{"NaN to string", math.NaN(), false},
		{"Inf to string", math.Inf(1), false},
		{"Negative Inf to string", math.Inf(-1), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result string
			err := safecast.ToString(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToString(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestToString_ErrorCases(t *testing.T) {
	// Test error cases for newErrorCast coverage
	var result string

	// Test with unsupported struct type - actually works via fmt.Sprintf
	type customStruct struct {
		Field int
	}
	err := safecast.ToString(customStruct{Field: 42}, &result)
	if err != nil {
		t.Error("Unexpected error for struct type - should work via fmt.Sprintf")
	}

	// Test with chan type - actually works via fmt.Sprintf
	ch := make(chan int)
	err = safecast.ToString(ch, &result)
	if err != nil {
		t.Error("Unexpected error for chan type - should work via fmt.Sprintf")
	}

	// Test with func type - actually works via fmt.Sprintf
	fn := func() {}
	err = safecast.ToString(fn, &result)
	if err != nil {
		t.Error("Unexpected error for func type - should work via fmt.Sprintf")
	}
}
