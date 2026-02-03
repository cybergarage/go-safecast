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
	"time"

	"github.com/cybergarage/go-safecast/safecast"
)

func TestToTime_Extended(t *testing.T) {
	// Test with various supported time formats
	tests := []struct {
		name    string
		input   any
		wantErr bool
	}{
		// time.Time and *time.Time - supported
		{"time.Time to time", time.Now(), false},
		{"*time.Time to time", func() any { now := time.Now(); return &now }(), false},

		// String formats that should work with supported layouts
		{"RFC3339 format", "2022-01-01T15:04:05Z", false},
		{"RFC3339Nano format", "2022-01-01T15:04:05.123456789Z", false},
		{"RFC1123 format", "Sat, 01 Jan 2022 15:04:05 GMT", false},
		{"RFC1123Z format", "Sat, 01 Jan 2022 15:04:05 +0000", false},

		// String pointers
		{"*string RFC3339", func() any { s := "2022-01-01T15:04:05Z"; return &s }(), false},

		// []byte format
		{"[]byte RFC3339", []byte("2022-01-01T15:04:05Z"), false},

		// Unsupported string formats - these should error
		{"Date only string", "2022-01-01", true},
		{"Time only string", "15:04:05", true},
		{"Unix timestamp string", "1640995200", true},
		{"MM/DD/YYYY", "01/01/2022", true},
		{"DD/MM/YYYY", "01/01/2022", true},
		{"YYYY/MM/DD", "2022/01/01", true},
		{"YYYYMMDD", "20220101", true},
		{"HHMMSS", "000000", true},
		{"Kitchen time format", "3:04PM", true},
		{"RFC822 format", "02 Jan 06 15:04 MST", true},
		{"RFC850 format", "Monday, 02-Jan-06 15:04:05 MST", true},
		{"ANSIC format", "Mon Jan _2 15:04:05 2006", true},
		{"Ruby date format", "Mon Jan 02 15:04:05 -0700 2006", true},

		// Unsupported types - these should error
		{"int Unix timestamp", 1640995200, true},
		{"int8 Unix timestamp", int8(127), true},
		{"int16 Unix timestamp", int16(32767), true},
		{"int32 Unix timestamp", int32(1640995200), true},
		{"int64 Unix timestamp", int64(1640995200), true},
		{"uint Unix timestamp", uint(1640995200), true},
		{"uint8 Unix timestamp", uint8(255), true},
		{"uint16 Unix timestamp", uint16(65535), true},
		{"uint32 Unix timestamp", uint32(1640995200), true},
		{"uint64 Unix timestamp", uint64(1640995200), true},
		{"float32 Unix timestamp", float32(1.6409952e+09), true},
		{"float64 Unix timestamp", float64(1.640995200123456e+09), true},
		{"*int to time", func() any { i := 1640995200; return &i }(), true},
		{"bool true to time", true, true},
		{"bool false to time", false, true},
		{"zero Unix timestamp", 0, true},
		{"negative Unix timestamp", -1, true},
		{"very large Unix timestamp", int64(253402300799), true},

		// Invalid string formats
		{"invalid string", "not a time", true},
		{"empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result time.Time
			err := safecast.ToTime(tt.input, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToTime(%v) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestToTime_ErrorCases(t *testing.T) {
	tests := []struct {
		name  string
		input any
	}{
		{"unsupported type", []int{1, 2, 3}},
		{"map type", map[string]int{"a": 1}},
		{"struct type", struct{ A int }{42}},
		{"function type", func() {}},
		{"channel type", make(chan int)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result time.Time
			err := safecast.ToTime(tt.input, &result)
			if err == nil {
				t.Errorf("ToTime(%v) should return error, but got nil", tt.input)
			}
		})
	}
}
