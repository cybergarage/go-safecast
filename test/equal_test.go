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

	"github.com/cybergarage/go-safecast/safecast"
)

func TestEqual_Ints(t *testing.T) {
	if !safecast.Equal(42, 42) {
		t.Error("Equal failed for int/int")
	}
	if !safecast.Equal(42, int64(42)) {
		t.Error("Equal failed for int/int64")
	}
	if safecast.Equal(42, 43) {
		t.Error("Equal failed for different ints")
	}
}

func TestEqual_Uints(t *testing.T) {
	if !safecast.Equal(uint(10), uint(10)) {
		t.Error("Equal failed for uint/uint")
	}
	if !safecast.Equal(uint8(10), 10) {
		t.Error("Equal failed for uint8/int")
	}
	if safecast.Equal(uint16(10), 11) {
		t.Error("Equal failed for different uint16/int")
	}
}

func TestEqual_Floats(t *testing.T) {
	if !safecast.Equal(float32(3.14), 3.14) {
		t.Error("Equal failed for float32/float64")
	}
	if !safecast.Equal(float64(2.71), 2.71) {
		t.Error("Equal failed for float64/float64")
	}
	if safecast.Equal(float64(1.23), 1.24) {
		t.Error("Equal failed for different float64")
	}
}

func TestEqual_Bool(t *testing.T) {
	if !safecast.Equal(true, true) {
		t.Error("Equal failed for true/true")
	}
	if safecast.Equal(true, false) {
		t.Error("Equal failed for true/false")
	}
}

func TestEqual_String(t *testing.T) {
	if !safecast.Equal("abc", "abc") {
		t.Error("Equal failed for string/string")
	}
	if safecast.Equal("abc", "def") {
		t.Error("Equal failed for different strings")
	}
}

func TestEqual_Bytes(t *testing.T) {
	if !safecast.Equal([]byte{1, 2, 3}, []byte{1, 2, 3}) {
		t.Error("Equal failed for equal byte slices")
	}
	if safecast.Equal([]byte{1, 2, 3}, []byte{1, 2, 4}) {
		t.Error("Equal failed for different byte slices")
	}
	if safecast.Equal([]byte{1, 2}, []byte{1, 2, 3}) {
		t.Error("Equal failed for different length byte slices")
	}
}

func TestEqual_Nil(t *testing.T) {
	if !safecast.Equal(nil, nil) {
		t.Error("Equal failed for nil/nil")
	}
	if safecast.Equal(nil, 0) {
		t.Error("Equal failed for nil/0")
	}
}

func TestEqual_DeepEqualFallback(t *testing.T) {
	type foo struct{ A int }
	a := foo{A: 1}
	b := foo{A: 1}
	c := foo{A: 2}
	if !safecast.Equal(a, b) {
		t.Error("Equal failed for struct DeepEqual")
	}
	if safecast.Equal(a, c) {
		t.Error("Equal failed for different struct DeepEqual")
	}
}
