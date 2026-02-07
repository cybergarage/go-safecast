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
	"testing"
	"time"

	"github.com/cybergarage/go-safecast/safecast"
)

func TestToCast(t *testing.T) {
	t.Run("ToTime", func(t *testing.T) {
		tests := []struct {
			from   string
			layout string
		}{
			{
				from:   "2022-01-01",
				layout: safecast.ISO8601DateLayout,
			},
			{
				from:   "00:00:00",
				layout: safecast.ISO8601TimeLayout,
			},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("%s (%s)", test.from, test.layout), func(t *testing.T) {
				var to time.Time
				if err := safecast.ToTime(test.from, &to, test.layout); err != nil {
					t.Error(err)
					return
				}
			})
		}
	})
}

// TestGenericTo tests the generic To() function
func TestGenericTo(t *testing.T) {
	t.Run("to int types", func(t *testing.T) {
		var result int
		err := safecast.To(42, &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		var result64 int64
		err = safecast.To(42, &result64)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("to uint types", func(t *testing.T) {
		var result uint
		err := safecast.To(42, &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("to float types", func(t *testing.T) {
		var result float64
		err := safecast.To(42.5, &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("to string", func(t *testing.T) {
		var result string
		err := safecast.To(42, &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("to bool", func(t *testing.T) {
		var result bool
		err := safecast.To(1, &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("to bytes", func(t *testing.T) {
		var result []byte
		err := safecast.To("hello", &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("unsupported type", func(t *testing.T) {
		type CustomType struct{}
		var custom CustomType
		err := safecast.To(42, &custom)
		if err == nil {
			t.Error("expected error for unsupported type")
		}
	})
}

// TestToBytes tests ToBytes function
func TestToBytes(t *testing.T) {
	t.Run("string to bytes", func(t *testing.T) {
		input := "hello world"
		var result []byte
		err := safecast.ToBytes(input, &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	})

	t.Run("string pointer to bytes", func(t *testing.T) {
		input := "test data"
		var result []byte
		err := safecast.ToBytes(&input, &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	})

	t.Run("bytes to bytes", func(t *testing.T) {
		input := []byte("test")
		var result []byte
		err := safecast.ToBytes(input, &result)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	})

	t.Run("unsupported type", func(t *testing.T) {
		input := 42
		var result []byte
		err := safecast.ToBytes(input, &result)
		if err == nil {
			t.Error("expected error for unsupported type")
		}
	})
}
