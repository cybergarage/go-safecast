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

package safecast_test

import (
	"fmt"
	"math"

	"github.com/cybergarage/go-safecast/safecast"
)

func ExampleToInt8() {
	var from int8
	var to int8

	from = 1
	if err := safecast.ToInt8(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt8
	if err := safecast.ToInt8(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt8
	if err := safecast.ToInt8(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	var from64 int
	from64 = math.MaxInt
	if err := safecast.ToInt8(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from64 = math.MinInt
	if err := safecast.ToInt8(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 127
	// -128
	// cast error : out of range 9223372036854775807 > *int8
	// cast error : out of range -9223372036854775808 > *int8
}

func ExampleToInt16() {
	var from int16
	var to int16

	from = 1
	if err := safecast.ToInt16(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt16
	if err := safecast.ToInt16(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt16
	if err := safecast.ToInt16(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	var from64 int
	from64 = math.MaxInt
	if err := safecast.ToInt16(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from64 = math.MinInt
	if err := safecast.ToInt16(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 32767
	// -32768
	// cast error : out of range 9223372036854775807 > *int16
	// cast error : out of range -9223372036854775808 > *int16
}

func ExampleToInt32() {
	var from int32
	var to int32

	from = 1
	if err := safecast.ToInt32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt32
	if err := safecast.ToInt32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt32
	if err := safecast.ToInt32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	var from64 int
	from64 = math.MaxInt
	if err := safecast.ToInt32(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from64 = math.MinInt
	if err := safecast.ToInt32(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 2147483647
	// -2147483648
	// cast error : out of range 9223372036854775807 > *int32
	// cast error : out of range -9223372036854775808 > *int32
}

func ExampleToInt64() {
	var from int64
	var to int64

	from = 1
	if err := safecast.ToInt64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt64
	if err := safecast.ToInt64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt64
	if err := safecast.ToInt64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt
	if err := safecast.ToInt64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt
	if err := safecast.ToInt64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 9223372036854775807
	// -9223372036854775808
	// 9223372036854775807
	// -9223372036854775808
}
