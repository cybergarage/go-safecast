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
