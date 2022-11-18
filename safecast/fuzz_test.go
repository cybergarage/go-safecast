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

package safecast

import (
	"math"
	"testing"
)

func FuzzFromInt(f *testing.F) {
	f.Add(int(0))
	f.Add(int(math.MinInt))
	f.Add(int(math.MaxInt))
	f.Fuzz(func(t *testing.T, from int) {
		var to int
		if err := FromInt(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromInt8(f *testing.F) {
	f.Add(int8(0))
	f.Add(int8(math.MinInt8))
	f.Add(int8(math.MaxInt8))
	f.Fuzz(func(t *testing.T, from int8) {
		var to int8
		if err := FromInt8(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromInt16(f *testing.F) {
	f.Add(int16(0))
	f.Add(int16(math.MinInt16))
	f.Add(int16(math.MaxInt16))
	f.Fuzz(func(t *testing.T, from int16) {
		var to int16
		if err := FromInt16(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromInt32(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(math.MinInt32))
	f.Add(int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, from int32) {
		var to int32
		if err := FromInt32(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromInt64(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, from int64) {
		var to int64
		if err := FromInt64(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}
