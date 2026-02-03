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
	"bytes"
	"testing"

	"github.com/cybergarage/go-safecast/safecast"
)

func TestToBool_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected bool
		wantErr  bool
	}{
		// Integer types - only values > 0 return true
		{"int 0 to bool", 0, false, false},
		{"int 1 to bool", 1, true, false},
		{"int 42 to bool", 42, true, false},
		{"int -1 to bool", -1, false, false},
		{"int8 0 to bool", int8(0), false, false},
		{"int8 1 to bool", int8(1), true, false},
		{"int16 0 to bool", int16(0), false, false},
		{"int16 1 to bool", int16(1), true, false},
		{"int32 0 to bool", int32(0), false, false},
		{"int32 1 to bool", int32(1), true, false},
		{"int64 0 to bool", int64(0), false, false},
		{"int64 1 to bool", int64(1), true, false},

		// Unsigned integer types
		{"uint 0 to bool", uint(0), false, false},
		{"uint 1 to bool", uint(1), true, false},
		{"uint8 0 to bool", uint8(0), false, false},
		{"uint8 1 to bool", uint8(1), true, false},
		{"uint16 0 to bool", uint16(0), false, false},
		{"uint16 1 to bool", uint16(1), true, false},
		{"uint32 0 to bool", uint32(0), false, false},
		{"uint32 1 to bool", uint32(1), true, false},
		{"uint64 0 to bool", uint64(0), false, false},
		{"uint64 1 to bool", uint64(1), true, false},

		// Float types - these should error (not supported by ToBool)
		{"float32 0.0 to bool", float32(0.0), false, true},
		{"float32 1.0 to bool", float32(1.0), false, true},
		{"float32 3.14 to bool", float32(3.14), false, true},
		{"float64 0.0 to bool", 0.0, false, true},
		{"float64 1.0 to bool", 1.0, false, true},
		{"float64 3.14 to bool", 3.14, false, true},

		// String types
		{"string 'true' to bool", "true", true, false},
		{"string 'false' to bool", "false", false, false},
		{"string '1' to bool", "1", true, false},
		{"string '0' to bool", "0", false, false},
		{"string 'yes' to bool", "yes", false, true}, // should error
		{"string 'no' to bool", "no", false, true},   // should error
		{"string empty to bool", "", false, true},    // should error

		// Bool type
		{"bool true to bool", true, true, false},
		{"bool false to bool", false, false, false},

		// Byte slice type
		{"[]byte 'true' to bool", []byte("true"), true, false},
		{"[]byte 'false' to bool", []byte("false"), false, false},
		{"[]byte '1' to bool", []byte("1"), true, false},
		{"[]byte '0' to bool", []byte("0"), false, false},

		// Pointer types
		{"*int 0 to bool", func() any { i := 0; return &i }(), false, false},
		{"*int 1 to bool", func() any { i := 1; return &i }(), true, false},
		{"*bool true to bool", func() any { b := true; return &b }(), true, false},
		{"*bool false to bool", func() any { b := false; return &b }(), false, false},
		{"*string 'true' to bool", func() any { s := "true"; return &s }(), true, false},

		// Error cases
		{"unsupported type", []int{1, 2, 3}, false, true},
		{"map type", map[string]int{"a": 1}, false, true},
		{"struct type", struct{ A int }{42}, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result bool
			err := safecast.ToBool(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBool(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ToBool(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToBytes_Extended(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected []byte
		wantErr  bool
	}{
		// String types - supported
		{"string to bytes", "hello", []byte("hello"), false},
		{"empty string to bytes", "", []byte(""), false},
		{"unicode string to bytes", "こんにちは", []byte("こんにちは"), false},

		// *string types - supported
		{"*string to bytes", func() any { s := "hello"; return &s }(), []byte("hello"), false},
		{"*empty string to bytes", func() any { s := ""; return &s }(), []byte(""), false},

		// []byte types - supported
		{"bytes to bytes", []byte("hello"), []byte("hello"), false},
		{"empty bytes to bytes", []byte(""), []byte(""), false},
		{"binary bytes to bytes", []byte{1, 2, 3, 255}, []byte{1, 2, 3, 255}, false},

		// Unsupported types - these should error
		{"int to bytes", 42, nil, true},
		{"int8 to bytes", int8(100), nil, true},
		{"int16 to bytes", int16(1000), nil, true},
		{"int32 to bytes", int32(100000), nil, true},
		{"int64 to bytes", int64(1000000), nil, true},
		{"uint to bytes", uint(42), nil, true},
		{"uint8 to bytes", uint8(200), nil, true},
		{"uint16 to bytes", uint16(50000), nil, true},
		{"uint32 to bytes", uint32(1000000), nil, true},
		{"uint64 to bytes", uint64(1000000), nil, true},
		{"float32 to bytes", float32(3.14), nil, true},
		{"float64 to bytes", 3.14159, nil, true},
		{"bool true to bytes", true, nil, true},
		{"bool false to bytes", false, nil, true},
		{"*int to bytes", func() any { i := 42; return &i }(), nil, true},
		{"*bool to bytes", func() any { b := true; return &b }(), nil, true},
		{"unsupported type", []int{1, 2, 3}, nil, true},
		{"map type", map[string]int{"a": 1}, nil, true},
		{"struct type", struct{ A int }{42}, nil, true},
		{"function type", func() {}, nil, true},
		{"channel type", make(chan int), nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result []byte
			err := safecast.ToBytes(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBytes(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if !bytes.Equal(result, tt.expected) {
					t.Errorf("ToBytes(%v) = %v, want %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestFromBytes_Extended(t *testing.T) {
	// FromBytes function supports string and []byte conversions
	t.Run("string_conversion", func(t *testing.T) {
		var result string
		err := safecast.FromBytes([]byte("hello"), &result)
		if err != nil {
			t.Errorf("FromBytes should convert to string, but got error: %v", err)
		}
		if result != "hello" {
			t.Errorf("FromBytes result = %s, want hello", result)
		}
	})

	t.Run("bytes_conversion", func(t *testing.T) {
		var result []byte
		err := safecast.FromBytes([]byte("hello"), &result)
		if err != nil {
			t.Errorf("FromBytes should convert to []byte, but got error: %v", err)
		}
		if string(result) != "hello" {
			t.Errorf("FromBytes result = %s, want hello", string(result))
		}
	})

	t.Run("int_conversion", func(t *testing.T) {
		var result int
		err := safecast.FromBytes([]byte("42"), &result)
		if err == nil {
			t.Errorf("FromBytes should error when converting to int, but got nil error")
		}
	})
}
