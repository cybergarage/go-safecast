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

package test

import (
	"testing"

	"github.com/cybergarage/go-safecast/safecast"
)

func TestToInt8_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int8
		wantErr  bool
	}{
		{"int8 to int8", int8(42), 42, false},
		{"int16 to int8", int16(42), 42, false},
		{"int16 overflow", int16(300), 0, true},
		{"int16 underflow", int16(-300), 0, true},
		{"string to int8", "42", 42, false},
		{"bool true to int8", true, 1, false},
		{"bool false to int8", false, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result int8
			err := safecast.ToInt8(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToInt8(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToUint8_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected uint8
		wantErr  bool
	}{
		{"uint8 to uint8", uint8(42), 42, false},
		{"int8 to uint8", int8(42), 42, false},
		{"negative int to uint8", int(-1), 0, true},
		{"int16 to uint8", int16(42), 42, false}, // No overflow check for int16
		{"string to uint8", "42", 42, false},
		{"bool true to uint8", true, 1, false},
		{"bool false to uint8", false, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result uint8
			err := safecast.ToUint8(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToUint8(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToInt16_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int16
		wantErr  bool
	}{
		{"int8 to int16", int8(42), 42, false},
		{"int16 to int16", int16(1000), 1000, false},
		{"int32 to int16", int32(1000), 1000, false},
		{"string to int16", "1000", 1000, false},
		{"int32 overflow", int32(40000), 0, true},
		{"int32 underflow", int32(-40000), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result int16
			err := safecast.ToInt16(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToInt16(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToUint16_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected uint16
		wantErr  bool
	}{
		{"uint8 to uint16", uint8(42), 42, false},
		{"uint16 to uint16", uint16(1000), 1000, false},
		{"int16 to uint16", int16(1000), 1000, false},
		{"string to uint16", "1000", 1000, false},
		{"negative int to uint16", int(-1), 0, true},
		{"int32 to uint16", int32(1000), 1000, false}, // No overflow check for int32
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result uint16
			err := safecast.ToUint16(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToUint16(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToInt32_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int32
		wantErr  bool
	}{
		{"int8 to int32", int8(42), 42, false},
		{"int16 to int32", int16(1000), 1000, false},
		{"int32 to int32", int32(100000), 100000, false},
		{"string to int32", "100000", 100000, false},
		{"int64 overflow", int64(3000000000), 0, true},
		{"int64 underflow", int64(-3000000000), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result int32
			err := safecast.ToInt32(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToInt32(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToUint32_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected uint32
		wantErr  bool
	}{
		{"uint8 to uint32", uint8(42), 42, false},
		{"uint16 to uint32", uint16(1000), 1000, false},
		{"uint32 to uint32", uint32(100000), 100000, false},
		{"int32 to uint32", int32(100000), 100000, false},
		{"string to uint32", "100000", 100000, false},
		{"negative int to uint32", int(-1), 0, true},
		{"uint64 overflow", uint64(5000000000), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result uint32
			err := safecast.ToUint32(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint32(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToUint32(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToInt64_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int64
		wantErr  bool
	}{
		{"int8 to int64", int8(42), 42, false},
		{"int16 to int64", int16(1000), 1000, false},
		{"int32 to int64", int32(100000), 100000, false},
		{"int64 to int64", int64(1000000), 1000000, false},
		{"uint32 to int64", uint32(100000), 100000, false},
		{"string to int64", "1000000", 1000000, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result int64
			err := safecast.ToInt64(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToInt64(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToUint64_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected uint64
		wantErr  bool
	}{
		{"uint8 to uint64", uint8(42), 42, false},
		{"uint16 to uint64", uint16(1000), 1000, false},
		{"uint32 to uint64", uint32(100000), 100000, false},
		{"uint64 to uint64", uint64(1000000), 1000000, false},
		{"int64 to uint64", int64(1000000), 1000000, false},
		{"string to uint64", "1000000", 1000000, false},
		{"negative int to uint64", int(-1), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result uint64
			err := safecast.ToUint64(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint64(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToUint64(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToInt_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int
		wantErr  bool
	}{
		{"int8 to int", int8(42), 42, false},
		{"int16 to int", int16(1000), 1000, false},
		{"int32 to int", int32(100000), 100000, false},
		{"int to int", int(42), 42, false},
		{"uint8 to int", uint8(42), 42, false},
		{"string to int", "42", 42, false},
		{"bool true to int", true, 1, false},
		{"bool false to int", false, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result int
			err := safecast.ToInt(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToInt(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToUint_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected uint
		wantErr  bool
	}{
		{"uint8 to uint", uint8(42), 42, false},
		{"uint16 to uint", uint16(1000), 1000, false},
		{"uint32 to uint", uint32(100000), 100000, false},
		{"uint to uint", uint(42), 42, false},
		{"int to uint", int(42), 42, false},
		{"string to uint", "42", 42, false},
		{"bool true to uint", true, 1, false},
		{"bool false to uint", false, 0, false},
		{"negative int to uint", int(-1), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result uint
			err := safecast.ToUint(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToUint(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// Additional comprehensive tests to improve low-coverage functions
func TestToInt8_ComprehensiveTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int8
		wantErr  bool
	}{
		// All supported input types for ToInt8
		{"int to int8", 42, 42, false},
		{"int8 to int8", int8(42), 42, false},
		{"int16 to int8", int16(100), 100, false},
		{"int32 to int8", int32(50), 50, false},
		{"int64 to int8", int64(75), 75, false},
		{"uint to int8", uint(30), 30, false},
		{"uint8 to int8", uint8(80), 80, false},
		{"uint16 to int8", uint16(90), 90, false},
		{"uint32 to int8", uint32(60), 60, false},
		{"uint64 to int8", uint64(70), 70, false},
		{"float32 to int8", float32(25.0), 25, false},
		{"float64 to int8", 35.0, 35, false},
		{"string to int8", "42", 42, false},
		{"bool true to int8", true, 1, false},
		{"bool false to int8", false, 0, false},
		{"[]byte to int8", []byte("50"), 50, false},

		// Pointer types
		{"*int to int8", func() interface{} { i := 42; return &i }(), 42, false},
		{"*string to int8", func() interface{} { s := "25"; return &s }(), 25, false},
		{"*bool to int8", func() interface{} { b := true; return &b }(), 1, false},

		// Overflow cases
		{"int overflow", 200, 0, true},
		{"int underflow", -200, 0, true},
		{"uint overflow", uint(300), 0, true},
		{"float overflow", 200.0, 0, true},
		{"float underflow", -200.0, 0, true},

		// Invalid string
		{"invalid string", "abc", 0, true},
		{"empty string", "", 0, true},

		// Unsupported types
		{"map type", map[string]int{"a": 1}, 0, true},
		{"slice type", []int{1, 2, 3}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result int8
			err := safecast.ToInt8(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToInt8(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToUint8_ComprehensiveTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected uint8
		wantErr  bool
	}{
		// All supported input types for ToUint8
		{"int to uint8", 200, 200, false},
		{"int8 to uint8", int8(100), 100, false},
		{"int16 to uint8", int16(150), 150, false},
		{"int32 to uint8", int32(200), 200, false},
		{"int64 to uint8", int64(180), 180, false},
		{"uint to uint8", uint(220), 220, false},
		{"uint8 to uint8", uint8(255), 255, false},
		{"uint16 to uint8", uint16(100), 100, false},
		{"uint32 to uint8", uint32(50), 50, false},
		{"uint64 to uint8", uint64(200), 200, false},
		{"float32 to uint8", float32(100.0), 100, false},
		{"float64 to uint8", 150.0, 150, false},
		{"string to uint8", "200", 200, false},
		{"bool true to uint8", true, 1, false},
		{"bool false to uint8", false, 0, false},
		{"[]byte to uint8", []byte("100"), 100, false},

		// Pointer types
		{"*int to uint8", func() interface{} { i := 200; return &i }(), 200, false},
		{"*string to uint8", func() interface{} { s := "100"; return &s }(), 100, false},
		{"*bool to uint8", func() interface{} { b := true; return &b }(), 1, false},

		// Overflow/underflow cases
		{"int overflow", 300, 0, true},
		{"int negative", -1, 0, true},
		{"uint overflow", uint(300), 0, true},
		{"float overflow", 300.0, 0, true},
		{"float negative", -1.0, 0, true},

		// Invalid cases
		{"invalid string", "xyz", 0, true},
		{"negative string", "-1", 0, true},
		{"unsupported type", []int{1, 2}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result uint8
			err := safecast.ToUint8(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint8(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToUint8(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToInt16_ComprehensiveTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int16
		wantErr  bool
	}{
		// All basic types
		{"int to int16", 30000, 30000, false},
		{"int8 to int16", int8(100), 100, false},
		{"int16 to int16", int16(25000), 25000, false},
		{"int32 to int16", int32(20000), 20000, false},
		{"int64 to int16", int64(15000), 15000, false},
		{"uint to int16", uint(10000), 10000, false},
		{"uint8 to int16", uint8(255), 255, false},
		{"uint16 to int16", uint16(30000), 30000, false},
		{"float32 to int16", float32(5000.0), 5000, false},
		{"float64 to int16", 8000.0, 8000, false},
		{"string to int16", "12000", 12000, false},
		{"bool true to int16", true, 1, false},
		{"bool false to int16", false, 0, false},

		// Overflow cases
		{"int overflow", 40000, 0, true},
		{"int underflow", -40000, 0, true},
		{"uint overflow", uint(70000), 0, true},

		// Invalid cases
		{"invalid string", "invalid", 0, true},
		{"unsupported type", []string{"a"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result int16
			err := safecast.ToInt16(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToInt16(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToUint16_ComprehensiveTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected uint16
		wantErr  bool
	}{
		// All basic types
		{"int to uint16", 60000, 60000, false},
		{"int8 to uint16", int8(100), 100, false},
		{"int16 to uint16", int16(30000), 30000, false},
		{"int32 to uint16", int32(50000), 50000, false},
		{"int64 to uint16", int64(40000), 40000, false},
		{"uint to uint16", uint(35000), 35000, false},
		{"uint8 to uint16", uint8(255), 255, false},
		{"uint16 to uint16", uint16(65535), 65535, false},
		{"float32 to uint16", float32(20000.0), 20000, false},
		{"float64 to uint16", 25000.0, 25000, false},
		{"string to uint16", "50000", 50000, false},
		{"bool true to uint16", true, 1, false},
		{"bool false to uint16", false, 0, false},

		// Overflow/underflow cases
		{"int overflow", 70000, 0, true},
		{"int negative", -1, 0, true},
		{"uint overflow", uint(70000), 0, true},

		// Invalid cases
		{"invalid string", "invalid", 0, true},
		{"negative string", "-1", 0, true},
		{"unsupported type", []string{"b"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result uint16
			err := safecast.ToUint16(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUint16(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToUint16(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
