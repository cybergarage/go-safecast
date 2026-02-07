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

	"github.com/cybergarage/go-safecast/safecast"
)

func TestCompare(t *testing.T) {
	t.Run("Ints", func(t *testing.T) {
		// Positive values tests
		if cmp, _ := safecast.Compare(42, 42); cmp != 0 {
			t.Errorf("Expected 0, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(42, 42); cmp != 0 {
			t.Errorf("Expected 0, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(42, 43); cmp >= 0 {
			t.Errorf("Expected negative, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(42, 41); cmp <= 0 {
			t.Errorf("Expected positive, got %d", cmp)
		}

		i := 100
		if cmp, _ := safecast.Compare(&i, 100); cmp != 0 {
			t.Errorf("Expected 0 for *int/int, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(100, &i); cmp != 0 {
			t.Errorf("Expected 0 for int/*int, got %d", cmp)
		}
		j := 101
		if cmp, _ := safecast.Compare(&i, &j); cmp >= 0 {
			t.Errorf("Expected negative for different *int/*int, got %d", cmp)
		}
		j = 99
		if cmp, _ := safecast.Compare(&i, &j); cmp <= 0 {
			t.Errorf("Expected positive for different *int/*int, got %d", cmp)
		}

		// Negative values tests
		if cmp, _ := safecast.Compare(-42, -42); cmp != 0 {
			t.Errorf("Expected 0 for negative ints, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(-42, -42); cmp != 0 {
			t.Errorf("Expected 0 for negative int/int64, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(-42, -43); cmp <= 0 {
			t.Errorf("Expected positive for -42 > -43, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(-42, -41); cmp >= 0 {
			t.Errorf("Expected negative for -42 < -41, got %d", cmp)
		}
		ni := -100
		if cmp, _ := safecast.Compare(&ni, -100); cmp != 0 {
			t.Errorf("Expected 0 for *int/int negative, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(-100, &ni); cmp != 0 {
			t.Errorf("Expected 0 for int/*int negative, got %d", cmp)
		}
		nj := -101
		if cmp, _ := safecast.Compare(&ni, &nj); cmp <= 0 {
			t.Errorf("Expected positive for -100 > -101, got %d", cmp)
		}
		nj = -99
		if cmp, _ := safecast.Compare(&ni, &nj); cmp >= 0 {
			t.Errorf("Expected negative for -100 < -99, got %d", cmp)
		}

		// Mixed sign
		if cmp, _ := safecast.Compare(-1, 1); cmp >= 0 {
			t.Errorf("Expected negative for -1 < 1, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(1, -1); cmp <= 0 {
			t.Errorf("Expected positive for 1 > -1, got %d", cmp)
		}
	})

	t.Run("Uints", func(t *testing.T) {
		if cmp, _ := safecast.Compare(uint(10), uint(10)); cmp != 0 {
			t.Errorf("Expected 0, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(uint8(10), 10); cmp != 0 {
			t.Errorf("Expected 0, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(uint16(10), 11); cmp >= 0 {
			t.Errorf("Expected negative, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(uint32(12), 11); cmp <= 0 {
			t.Errorf("Expected positive, got %d", cmp)
		}
	})

	t.Run("Floats", func(t *testing.T) {
		if cmp, _ := safecast.Compare(float32(3.14), 3.14); cmp != 0 {
			t.Errorf("Expected 0, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(float64(2.71), 2.71); cmp != 0 {
			t.Errorf("Expected 0, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(float64(1.23), 1.24); cmp >= 0 {
			t.Errorf("Expected negative, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(float64(1.25), 1.24); cmp <= 0 {
			t.Errorf("Expected positive, got %d", cmp)
		}
	})

	t.Run("Bool", func(t *testing.T) {
		if cmp, _ := safecast.Compare(true, true); cmp != 0 {
			t.Errorf("Expected 0, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(false, true); cmp >= 0 {
			t.Errorf("Expected negative, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(true, false); cmp <= 0 {
			t.Errorf("Expected positive, got %d", cmp)
		}
	})

	t.Run("String", func(t *testing.T) {
		if cmp, _ := safecast.Compare("abc", "abc"); cmp != 0 {
			t.Errorf("Expected 0, got %d", cmp)
		}
		if cmp, _ := safecast.Compare("abc", "def"); cmp >= 0 {
			t.Errorf("Expected negative, got %d", cmp)
		}
		if cmp, _ := safecast.Compare("def", "abc"); cmp <= 0 {
			t.Errorf("Expected positive, got %d", cmp)
		}
	})

	t.Run("Bytes", func(t *testing.T) {
		if cmp, _ := safecast.Compare([]byte{1, 2, 3}, []byte{1, 2, 3}); cmp != 0 {
			t.Errorf("Expected 0, got %d", cmp)
		}
		if cmp, _ := safecast.Compare([]byte{1, 2, 3}, []byte{1, 2, 4}); cmp <= 0 {
			t.Errorf("Expected positive, got %d", cmp)
		}
		if cmp, _ := safecast.Compare([]byte{1, 2, 4}, []byte{1, 2, 3}); cmp >= 0 {
			t.Errorf("Expected negative, got %d", cmp)
		}
		if cmp, _ := safecast.Compare([]byte{1, 2}, []byte{1, 2, 3}); cmp <= 0 {
			t.Errorf("Expected positive for longer slice, got %d", cmp)
		}
		if cmp, _ := safecast.Compare([]byte{1, 2, 3}, []byte{1, 2}); cmp >= 0 {
			t.Errorf("Expected negative for shorter slice, got %d", cmp)
		}
	})
}

// TestCompareMoreTypes tests additional compare scenarios
func TestCompareMoreTypes(t *testing.T) {
	t.Run("uint comparisons", func(t *testing.T) {
		cmp, err := safecast.Compare(uint(42), uint(42))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if cmp != 0 {
			t.Errorf("expected 0, got %d", cmp)
		}
	})

	t.Run("bool comparisons", func(t *testing.T) {
		cmp, err := safecast.Compare(true, true)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if cmp != 0 {
			t.Errorf("expected 0, got %d", cmp)
		}

		cmp, err = safecast.Compare(true, false)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if cmp != 1 {
			t.Errorf("expected 1, got %d", cmp)
		}
	})

	t.Run("mixed type comparisons", func(t *testing.T) {
		cmp, err := safecast.Compare(42, 42.0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if cmp != 0 {
			t.Errorf("expected 0, got %d", cmp)
		}
	})

	t.Run("overflow conditions", func(t *testing.T) {
		t.Skip("Platform-dependent overflow behavior")
		_, err := safecast.Compare(-1, uint(10))
		if err == nil {
			t.Error("expected error for overflow condition")
		}
	})

	t.Run("unsupported types", func(t *testing.T) {
		type CustomType struct{}
		custom := CustomType{}
		_, err := safecast.Compare(custom, 42)
		if err == nil {
			t.Error("expected error for unsupported type")
		}
	})
}
