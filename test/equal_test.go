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
	})

	t.Run("Floats", func(t *testing.T) {
		if !safecast.Equal(float32(3.14), 3.14) {
			t.Error("Equal failed for float32/float64")
		}
		if !safecast.Equal(float64(2.71), 2.71) {
			t.Error("Equal failed for float64/float64")
		}
		if safecast.Equal(float64(1.23), 1.24) {
			t.Error("Equal failed for different float64")
		}
	})

	t.Run("Strings", func(t *testing.T) {
		if !safecast.Equal("hello", "hello") {
			t.Error("Equal failed for string/string")
		}
		if safecast.Equal("hello", "world") {
			t.Error("Equal failed for different strings")
		}
	})

	t.Run("Booleans", func(t *testing.T) {
		if !safecast.Equal(true, true) {
			t.Error("Equal failed for bool/bool")
		}
		if safecast.Equal(true, false) {
			t.Error("Equal failed for true/false")
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
	})
}
