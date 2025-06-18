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
	"fmt"
	"math"
	"time"
)

func ExampleToInt8() {
	var from int8
	var to int8

	from = 1
	if err := ToInt8(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt8
	if err := ToInt8(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt8
	if err := ToInt8(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToInt8("123", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	var from64 int
	from64 = math.MaxInt
	if err := ToInt8(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from64 = math.MinInt
	if err := ToInt8(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 127
	// -128
	// 123
	// cast error : out of range 9223372036854775807 > *int8
	// cast error : out of range -9223372036854775808 < *int8
}

func ExampleToInt16() {
	var from int16
	var to int16

	from = 1
	if err := ToInt16(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt16
	if err := ToInt16(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt16
	if err := ToInt16(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToInt16("123", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	var from64 int
	from64 = math.MaxInt
	if err := ToInt16(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from64 = math.MinInt
	if err := ToInt16(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 32767
	// -32768
	// 123
	// cast error : out of range 9223372036854775807 > *int16
	// cast error : out of range -9223372036854775808 < *int16
}

func ExampleToInt32() {
	var from int32
	var to int32

	from = 1
	if err := ToInt32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt32
	if err := ToInt32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt32
	if err := ToInt32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToInt32("123", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	var from64 int
	from64 = math.MaxInt
	if err := ToInt32(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from64 = math.MinInt
	if err := ToInt32(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 2147483647
	// -2147483648
	// 123
	// cast error : out of range 9223372036854775807 > *int32
	// cast error : out of range -9223372036854775808 < *int32
}

func ExampleToInt64() {
	var from int64
	var to int64

	from = 1
	if err := ToInt64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt64
	if err := ToInt64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt64
	if err := ToInt64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt
	if err := ToInt64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt
	if err := ToInt64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToInt64("123", &to); err != nil {
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
	// 123
}

func ExampleToInt() {
	var from int
	var to int

	from = 1
	if err := ToInt(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt
	if err := ToInt(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt
	if err := ToInt(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxInt64
	if err := ToInt(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MinInt64
	if err := ToInt(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToInt("123", &to); err != nil {
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
	// 123
}

func ExampleToUint8() {
	var from uint8
	var to uint8

	from = 1
	if err := ToUint8(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxUint8
	if err := ToUint8(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = 0
	if err := ToUint8(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToUint8("123", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from64 := uint(math.MaxUint)
	if err := ToUint8(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 255
	// 0
	// 123
	// cast error : out of range 18446744073709551615 > *uint8
}

func ExampleToUint16() {
	var from uint16
	var to uint16

	from = 1
	if err := ToUint16(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxUint16
	if err := ToUint16(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = 0
	if err := ToUint16(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToUint16("123", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from64 := uint(math.MaxUint)
	if err := ToUint16(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 65535
	// 0
	// 123
	// cast error : out of range 18446744073709551615 > *uint16
}

func ExampleToUint32() {
	var from uint32
	var to uint32

	from = 1
	if err := ToUint32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxUint32
	if err := ToUint32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = 0
	if err := ToUint32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToUint32("123", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from64 := uint(math.MaxUint)
	if err := ToUint32(from64, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 4294967295
	// 0
	// 123
	// cast error : out of range 18446744073709551615 > *uint32
}

func ExampleToUint64() {
	var from uint64
	var to uint64

	from = 1
	if err := ToUint64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxUint64
	if err := ToUint64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = 0
	if err := ToUint64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToUint64("123", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 18446744073709551615
	// 0
	// 123
}

func ExampleToUint() {
	var from uint
	var to uint

	from = 1
	if err := ToUint(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxUint
	if err := ToUint(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = 0
	if err := ToUint(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToUint("123", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 18446744073709551615
	// 0
	// 123
}

func ExampleToBool() {
	var to bool

	if err := ToBool(1, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToBool("TRUE", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// true
	// true
}

func ExampleToFloat64() {
	var from float64
	var to float64

	from = 1.0
	if err := ToFloat64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxFloat64
	if err := ToFloat64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.SmallestNonzeroFloat64
	if err := ToFloat64(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToFloat64("123.0", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 1.7976931348623157e+308
	// 5e-324
	// 123
}

func ExampleToFloat32() {
	var from float32
	var to float32

	from = 1.0
	if err := ToFloat32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.MaxFloat32
	if err := ToFloat32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	from = math.SmallestNonzeroFloat32
	if err := ToFloat32(from, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToFloat32("123.0", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 1
	// 3.4028235e+38
	// 1e-45
	// 123
}

func ExampleToString() {
	var to string

	if err := ToString(123, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToString(123.0, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	if err := ToString(true, &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 123
	// 123
	// true
}

func ExampleToTime() {
	var to time.Time

	if err := ToTime("2022-01-01T00:00:00", &to, ISO860TimestampLayout); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(to)
	}

	// Output:
	// 2022-01-01 00:00:00 +0000 UTC
}

func ExampleToBytes() {
	var to []byte

	if err := ToBytes("abc", &to); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(to))
	}

	// Output:
	// abc
}
