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
	"testing"
	"time"

	"github.com/cybergarage/go-safecast/safecast"
)

func TestFromGeneric(t *testing.T) {
	// Test From with various types
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
		wantErr  bool
	}{
		{"int to int64", 42, int64(42), false},
		{"float64 to int64", 42.0, int64(42), false},
		{"string to int64", "42", int64(42), false},
		{"bool to int64", true, int64(1), false},
		{"bool false to int64", false, int64(0), false},
		{"invalid string", "invalid", int64(0), true},
		{"float with decimal", 42.7, int64(42), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result int64
			err := safecast.From(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("From() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("From() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToGeneric(t *testing.T) {
	// Test To with various types
	tests := []struct {
		name     string
		input    interface{}
		wantErr  bool
		testFunc func() error
	}{
		{"int to int64", 42, false, func() error {
			var result int64
			return safecast.To(42, &result)
		}},
		{"float64 to int64", 42.0, false, func() error {
			var result int64
			return safecast.To(42.0, &result)
		}},
		{"string to int64", "42", false, func() error {
			var result int64
			return safecast.To("42", &result)
		}},
		{"bool to int64", true, false, func() error {
			var result int64
			return safecast.To(true, &result)
		}},
		{"int to string", 42, false, func() error {
			var result string
			return safecast.To(42, &result)
		}},
		{"float to string", 3.14, false, func() error {
			var result string
			return safecast.To(3.14, &result)
		}},
		{"bool to string", true, false, func() error {
			var result string
			return safecast.To(true, &result)
		}},
		{"string to time", "2022-01-01T00:00:00Z", false, func() error {
			var result time.Time
			return safecast.To("2022-01-01T00:00:00Z", &result)
		}},
		{"invalid string to time", "invalid", true, func() error {
			var result time.Time
			return safecast.To("invalid", &result)
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.testFunc()
			if (err != nil) != tt.wantErr {
				t.Errorf("To() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFromErrors(t *testing.T) {
	// Test error cases for better coverage
	var result int64

	// Test with nil destination
	err := safecast.From(42, nil)
	if err == nil {
		t.Error("Expected error for nil destination")
	}

	// Test with non-pointer destination
	err = safecast.From(42, result)
	if err == nil {
		t.Error("Expected error for non-pointer destination")
	}

	// Test with unsupported type conversion
	var strResult string
	err = safecast.From([]int{1, 2, 3}, &strResult)
	if err == nil {
		t.Error("Expected error for unsupported type conversion")
	}
}

func TestToErrors(t *testing.T) {
	// Test error cases for better coverage

	// Test with unsupported target type
	err := safecast.To(42, 123) // non-pointer
	if err == nil {
		t.Error("Expected error for non-pointer target")
	}

	// Test with unsupported source type
	var result int64
	err = safecast.To([]int{1, 2, 3}, &result)
	if err == nil {
		t.Error("Expected error for unsupported source type")
	}
}

func TestFromGeneric_ExtensiveTypes(t *testing.T) {
	tests := []struct {
		name     string
		from     interface{}
		to       interface{}
		wantErr  bool
		validate func(interface{}) bool
	}{
		// Test all integer pointer conversions to improve From() coverage
		{"*int to int8", func() interface{} { i := 42; return &i }(), new(int8), false, func(v interface{}) bool { return *(v.(*int8)) == 42 }},
		{"*int8 to int", func() interface{} { i := int8(42); return &i }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 42 }},
		{"*int16 to int", func() interface{} { i := int16(1000); return &i }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 1000 }},
		{"*int32 to int", func() interface{} { i := int32(100000); return &i }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 100000 }},
		{"*int64 to int", func() interface{} { i := int64(1000000); return &i }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 1000000 }},

		// Test all unsigned integer pointer conversions
		{"*uint to int", func() interface{} { i := uint(42); return &i }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 42 }},
		{"*uint8 to int", func() interface{} { i := uint8(200); return &i }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 200 }},
		{"*uint16 to int", func() interface{} { i := uint16(50000); return &i }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 50000 }},
		{"*uint32 to int", func() interface{} { i := uint32(1000000); return &i }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 1000000 }},
		{"*uint64 to int", func() interface{} { i := uint64(1000000); return &i }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 1000000 }},

		// Test float pointer conversions
		{"*float32 to int", func() interface{} { f := float32(42.0); return &f }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 42 }},
		{"*float64 to int", func() interface{} { f := 42.0; return &f }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 42 }},

		// Test string pointer conversions
		{"*string to int", func() interface{} { s := "42"; return &s }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 42 }},

		// Test bool pointer conversions
		{"*bool true to int", func() interface{} { b := true; return &b }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 1 }},
		{"*bool false to int", func() interface{} { b := false; return &b }(), new(int), false, func(v interface{}) bool { return *(v.(*int)) == 0 }},

		// Test []byte conversions
		{"[]byte to string", []byte("hello world"), new(string), false, func(v interface{}) bool { return *(v.(*string)) == "hello world" }},
		{"[]byte to []byte", []byte("test"), new([]byte), false, func(v interface{}) bool {
			result := *(v.(*[]byte))
			return string(result) == "test"
		}},

		// Error cases for unsupported types
		{"unsupported map type", map[string]int{"key": 1}, new(int), true, nil},
		{"unsupported slice type", []int{1, 2, 3}, new(int), true, nil},
		{"unsupported struct type", struct{ A int }{42}, new(int), true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := safecast.From(tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("From(%v, %T) error = %v, wantErr %v", tt.from, tt.to, err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.validate != nil && !tt.validate(tt.to) {
				t.Errorf("From(%v, %T) validation failed", tt.from, tt.to)
			}
		})
	}
}

func TestToGeneric_ExtensiveTypes(t *testing.T) {
	tests := []struct {
		name     string
		from     interface{}
		to       interface{}
		wantErr  bool
		validate func(interface{}) bool
	}{
		// Test comprehensive To conversions for each integer type
		{"multiple types to int8", "100", new(int8), false, func(v interface{}) bool { return *(v.(*int8)) == 100 }},
		{"multiple types to int16", "30000", new(int16), false, func(v interface{}) bool { return *(v.(*int16)) == 30000 }},
		{"multiple types to int32", "2000000", new(int32), false, func(v interface{}) bool { return *(v.(*int32)) == 2000000 }},
		{"multiple types to int64", "9000000000", new(int64), false, func(v interface{}) bool { return *(v.(*int64)) == 9000000000 }},

		{"bool to int8", true, new(int8), false, func(v interface{}) bool { return *(v.(*int8)) == 1 }},
		{"float to int8", 50.0, new(int8), false, func(v interface{}) bool { return *(v.(*int8)) == 50 }},

		// Test comprehensive uint conversions
		{"string to uint8", "200", new(uint8), false, func(v interface{}) bool { return *(v.(*uint8)) == 200 }},
		{"string to uint16", "60000", new(uint16), false, func(v interface{}) bool { return *(v.(*uint16)) == 60000 }},
		{"string to uint32", "4000000000", new(uint32), false, func(v interface{}) bool { return *(v.(*uint32)) == 4000000000 }},
		{"string to uint64", "9000000000000000000", new(uint64), false, func(v interface{}) bool { return *(v.(*uint64)) == 9000000000000000000 }},

		{"bool to uint", true, new(uint), false, func(v interface{}) bool { return *(v.(*uint)) == 1 }},
		{"float to uint", 100.0, new(uint), false, func(v interface{}) bool { return *(v.(*uint)) == 100 }},

		// Test comprehensive float conversions
		{"int to float32", 42, new(float32), false, func(v interface{}) bool { return *(v.(*float32)) == 42.0 }},
		{"string to float32", "3.14159", new(float32), false, func(v interface{}) bool {
			result := *(v.(*float32))
			return result > 3.14 && result < 3.15
		}},

		{"int to float64", 123, new(float64), false, func(v interface{}) bool { return *(v.(*float64)) == 123.0 }},

		// Test comprehensive string conversions
		{"int8 to string", int8(42), new(string), false, func(v interface{}) bool { return *(v.(*string)) == "42" }},
		{"int16 to string", int16(1000), new(string), false, func(v interface{}) bool { return *(v.(*string)) == "1000" }},
		{"int32 to string", int32(100000), new(string), false, func(v interface{}) bool { return *(v.(*string)) == "100000" }},
		{"int64 to string", int64(1000000), new(string), false, func(v interface{}) bool { return *(v.(*string)) == "1000000" }},
		{"uint8 to string", uint8(200), new(string), false, func(v interface{}) bool { return *(v.(*string)) == "200" }},
		{"uint16 to string", uint16(50000), new(string), false, func(v interface{}) bool { return *(v.(*string)) == "50000" }},
		{"uint32 to string", uint32(1000000), new(string), false, func(v interface{}) bool { return *(v.(*string)) == "1000000" }},
		{"uint64 to string", uint64(1000000), new(string), false, func(v interface{}) bool { return *(v.(*string)) == "1000000" }},
		{"float32 to string", float32(3.14), new(string), false, func(v interface{}) bool { return *(v.(*string)) == "3.14" }},
		{"float64 to string", 2.71828, new(string), false, func(v interface{}) bool { return *(v.(*string)) == "2.71828" }},

		// Test comprehensive bool conversions
		{"int negative to bool", -5, new(bool), false, func(v interface{}) bool { return *(v.(*bool)) == false }},
		{"int positive to bool", 100, new(bool), false, func(v interface{}) bool { return *(v.(*bool)) == true }},
		{"uint zero to bool", uint(0), new(bool), false, func(v interface{}) bool { return *(v.(*bool)) == false }},
		{"uint positive to bool", uint(50), new(bool), false, func(v interface{}) bool { return *(v.(*bool)) == true }},
		{"string false to bool", "false", new(bool), false, func(v interface{}) bool { return *(v.(*bool)) == false }},
		{"string 0 to bool", "0", new(bool), false, func(v interface{}) bool { return *(v.(*bool)) == false }},
		{"string 1 to bool", "1", new(bool), false, func(v interface{}) bool { return *(v.(*bool)) == true }},

		// Test []byte conversions
		{"string to []byte", "hello", new([]byte), false, func(v interface{}) bool {
			result := *(v.(*[]byte))
			return string(result) == "hello"
		}},

		// Test time conversions
		{"RFC3339 to time", "2022-12-25T10:30:00Z", new(time.Time), false, func(v interface{}) bool {
			result := *(v.(*time.Time))
			return result.Year() == 2022 && result.Month() == 12 && result.Day() == 25
		}},
		{"RFC1123 to time", "Sun, 25 Dec 2022 10:30:00 GMT", new(time.Time), false, func(v interface{}) bool {
			result := *(v.(*time.Time))
			return result.Year() == 2022
		}},

		// Error cases - unsupported target types
		{"int to interface{}", 42, new(interface{}), true, nil},
		{"string to chan", "test", new(chan int), true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := safecast.To(tt.from, tt.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("To(%v, %T) error = %v, wantErr %v", tt.from, tt.to, err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.validate != nil && !tt.validate(tt.to) {
				t.Errorf("To(%v, %T) validation failed", tt.from, tt.to)
			}
		})
	}
}
