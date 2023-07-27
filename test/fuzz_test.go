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

	"github.com/cybergarage/go-safecast/safecast"
)

func FuzzFromInt(f *testing.F) {
	f.Add(int(0))
	f.Add(int(math.MinInt))
	f.Add(int(math.MaxInt))
	f.Fuzz(func(t *testing.T, from int) {
		var to int
		if err := safecast.FromInt(from, &to); err != nil {
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
		if err := safecast.FromInt8(from, &to); err != nil {
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
		if err := safecast.FromInt16(from, &to); err != nil {
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
		if err := safecast.FromInt32(from, &to); err != nil {
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
		if err := safecast.FromInt64(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromUint(f *testing.F) {
	f.Add(uint(0))
	f.Add(uint(math.MaxUint))
	f.Fuzz(func(t *testing.T, from uint) {
		var to uint
		if err := safecast.FromUint(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromUint8(f *testing.F) {
	f.Add(uint8(0))
	f.Add(uint8(math.MaxUint8))
	f.Fuzz(func(t *testing.T, from uint8) {
		var to uint8
		if err := safecast.FromUint8(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromUint16(f *testing.F) {
	f.Add(uint16(0))
	f.Add(uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, from uint16) {
		var to uint16
		if err := safecast.FromUint16(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromUint32(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, from uint32) {
		var to uint32
		if err := safecast.FromUint32(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromUint64(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, from uint64) {
		var to uint64
		if err := safecast.FromUint64(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromFloat64(f *testing.F) {
	f.Add(float64(-math.MaxFloat64))
	f.Add(float64(0))
	f.Add(float64(math.MaxFloat64))
	f.Fuzz(func(t *testing.T, from float64) {
		var to float64
		if err := safecast.FromFloat64(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromFloat32(f *testing.F) {
	f.Add(float32(-math.MaxFloat32))
	f.Add(float32(0))
	f.Add(float32(math.MaxFloat32))
	f.Fuzz(func(t *testing.T, from float32) {
		var to float32
		if err := safecast.FromFloat32(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromString(f *testing.F) {
	f.Add("a")
	f.Add("ab")
	f.Add("abc")
	f.Fuzz(func(t *testing.T, from string) {
		var to string
		if err := safecast.FromString(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}

func FuzzFromBool(f *testing.F) {
	f.Add(true)
	f.Add(false)
	f.Fuzz(func(t *testing.T, from bool) {
		var to bool
		if err := safecast.FromBool(from, &to); err != nil {
			t.Error(err)
			return
		}
		if to != from {
			t.Errorf("%v != %v", to, from)
		}
	})
}
