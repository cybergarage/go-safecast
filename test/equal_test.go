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

package safecasttest

import (
	"testing"
	"time"

	"github.com/cybergarage/go-safecast/safecast"
)

func TestEqual(t *testing.T) {
	t.Run("Ints", func(t *testing.T) {
		if !safecast.Equal(42, 42) {
			t.Error("Equal failed for int/int")
		}
		if !safecast.Equal(42, int64(42)) {
			t.Error("Equal failed for int/int64")
		}
		if safecast.Equal(42, 43) {
			t.Error("Equal failed for different ints")
		}

		// Pointer comparisons

		i := 100
		if !safecast.Equal(&i, 100) {
			t.Error("Equal failed for *int/int")
		}
		if !safecast.Equal(100, &i) {
			t.Error("Equal failed for int/*int")
		}
		j := 101
		if safecast.Equal(&i, &j) {
			t.Error("Equal failed for different *int/*int")
		}
	})

	t.Run("Uints", func(t *testing.T) {
		if !safecast.Equal(uint(10), uint(10)) {
			t.Error("Equal failed for uint/uint")
		}
		if !safecast.Equal(uint8(10), 10) {
			t.Error("Equal failed for uint8/int")
		}
		if safecast.Equal(uint16(10), 11) {
			t.Error("Equal failed for different uint16/int")
		}

		// Pointer comparisons

		u := uint32(20)
		if !safecast.Equal(&u, uint32(20)) {
			t.Error("Equal failed for *uint32/uint32")
		}
		if safecast.Equal(&u, uint32(21)) {
			t.Error("Equal failed for *uint32/different uint32")
		}
		u2 := uint64(30)
		if safecast.Equal(&u, &u2) {
			t.Error("Equal failed for different *uint32/*uint64")
		}
	})

	t.Run("Floats", func(t *testing.T) {
		if !safecast.Equal(float32(3.14), float32(3.14)) {
			t.Error("Equal failed for float32/float64")
		}
		if !safecast.Equal(float64(2.71), float64(2.71)) {
			t.Error("Equal failed for float64/float64")
		}
		if safecast.Equal(float64(1.23), float64(1.24)) {
			t.Error("Equal failed for different float64")
		}

		// Pointer comparisons

		f := 1.23
		if !safecast.Equal(&f, 1.23) {
			t.Error("Equal failed for *float64/float64")
		}
		if safecast.Equal(&f, 2.34) {
			t.Error("Equal failed for *float64/different float64")
		}
	})

	t.Run("Strings", func(t *testing.T) {
		if !safecast.Equal("hello", "hello") {
			t.Error("Equal failed for string/string")
		}
		if safecast.Equal("hello", "world") {
			t.Error("Equal failed for different strings")
		}

		// Pointer comparisons

		s := "GoLang"
		if safecast.Equal(&s, "golang") {
			t.Error("Equal failed for *string/string (case-insensitive)")
		}
		if safecast.Equal("golang", &s) {
			t.Error("Equal failed for string/*string (case-insensitive)")
		}
		s2 := "Python"
		if safecast.Equal(&s, &s2) {
			t.Error("Equal failed for different *string/*string")
		}
	})

	t.Run("Booleans", func(t *testing.T) {
		if !safecast.Equal(true, true) {
			t.Error("Equal failed for bool/bool")
		}
		if safecast.Equal(true, false) {
			t.Error("Equal failed for true/false")
		}

		// Pointer comparisons

		b := true
		if !safecast.Equal(&b, true) {
			t.Error("Equal failed for *bool/bool")
		}
		if safecast.Equal(&b, false) {
			t.Error("Equal failed for *bool/false")
		}
	})

	t.Run("Bytes", func(t *testing.T) {
		if !safecast.Equal([]byte{1, 2, 3}, []byte{1, 2, 3}) {
			t.Error("Equal failed for byte slice/byte slice")
		}
		if safecast.Equal([]byte{1, 2, 3}, []byte{1, 2, 4}) {
			t.Error("Equal failed for different byte slices")
		}
		if safecast.Equal([]byte{1, 2}, []byte{1, 2, 3}) {
			t.Error("Equal failed for different length byte slices")
		}
	})

	t.Run("Nil", func(t *testing.T) {
		if !safecast.Equal(nil, nil) {
			t.Error("Equal failed for nil/nil")
		}
		if safecast.Equal(nil, 0) {
			t.Error("Equal failed for nil/0")
		}
	})

	t.Run("Time", func(t *testing.T) {
		t1 := time.Now()
		t2 := t1.Add(0)
		if !safecast.Equal(t1, t2) {
			t.Error("Equal failed for time/time")
		}
		t3 := t1.Add(1 * time.Second)
		if safecast.Equal(t1, t3) {
			t.Error("Equal failed for different times")
		}

		now := time.Now()
		if !safecast.Equal(&now, now) {
			t.Error("Equal failed for *time.Time/time.Time")
		}
		later := now.Add(time.Minute)
		if safecast.Equal(&now, later) {
			t.Error("Equal failed for *time.Time/different time.Time")
		}
	})

	t.Run("Nil Pointers", func(t *testing.T) {
		if !safecast.Equal(nil, nil) {
			t.Error("Equal failed for nil/nil")
		}

		var pi *int = nil
		if !safecast.Equal(pi, nil) {
			t.Error("Equal failed for *int(nil)/nil")
		}
		var ps *string = nil
		if !safecast.Equal(nil, ps) {
			t.Error("Equal failed for nil/*string(nil)")
		}
	})
}
