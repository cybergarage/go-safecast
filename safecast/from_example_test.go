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
	"strconv"
)

func ExampleFromInt() {
	var v int
	if err := FromInt(math.MaxInt, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	var uv int8
	if err := FromInt(math.MaxInt, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	// Output:
	// 9223372036854775807
	// cast error : out of range 9223372036854775807 > *int8
}

func ExampleFromInt8() {
	var v int
	if err := FromInt8(math.MaxInt8, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	var uv uint8
	if err := FromInt8(math.MinInt8, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	// Output:
	// 127
	// cast error : out of range -128 < *uint8
}

func ExampleFromInt16() {
	var v int
	if err := FromInt16(math.MaxInt16, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	var uv uint16
	if err := FromInt16(math.MinInt16, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	// Output:
	// 32767
	// cast error : out of range -32768 < *uint16
}

func ExampleFromInt32() {
	var v int
	if err := FromInt32(math.MaxInt32, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	var uv uint32
	if err := FromInt32(math.MinInt32, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	// Output:
	// 2147483647
	// cast error : out of range -2147483648 < *uint32
}

func ExampleFromInt64() {
	var v int
	if err := FromInt64(math.MaxInt64, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	var uv uint64
	if err := FromInt64(math.MinInt64, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	// Output:
	// 9223372036854775807
	// cast error : out of range -9223372036854775808 < *uint64
}

func ExampleFromUint() {
	var uv uint
	if err := FromUint(math.MaxUint, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	var v int
	if err := FromUint(math.MaxUint, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	// Output:
	// 18446744073709551615
	// cast error : out of range 18446744073709551615 > *int
}

func ExampleFromUint8() {
	var uv uint
	if err := FromUint8(math.MaxUint8, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	var v int8
	if err := FromUint8(math.MaxUint8, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	// Output:
	// 255
	// cast error : out of range 255 > *int8
}

func ExampleFromUint16() {
	var uv uint
	if err := FromUint16(math.MaxUint16, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	var v int16
	if err := FromUint16(math.MaxUint16, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	// Output:
	// 65535
	// cast error : out of range 65535 > *int16
}

func ExampleFromUint32() {
	var uv uint
	if err := FromUint32(math.MaxUint32, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	var v int32
	if err := FromUint32(math.MaxUint32, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	// Output:
	// 4294967295
	// cast error : out of range 4294967295 > *int32
}

func ExampleFromUint64() {
	var uv uint
	if err := FromUint64(math.MaxUint64, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	var v int64
	if err := FromUint64(math.MaxUint64, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	// Output:
	// 18446744073709551615
	// cast error : out of range 18446744073709551615 > *int64
}

func ExampleFromFloat64() {
	var uv uint
	if err := FromFloat64(1.0, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	var v int64
	if err := FromFloat64(-1.0, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	// Output:
	// 1
	// -1
}

func ExampleFromFloat32() {
	var uv uint
	if err := FromFloat32(1, &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	var v int64
	if err := FromFloat32(-1, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	// Output:
	// 1
	// -1
}

func ExampleFromString() {
	var uv uint
	if err := FromString(strconv.Itoa(math.MaxInt64), &uv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", uv)
	}

	var v int8
	if err := FromString(strconv.Itoa(math.MaxInt64), &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	// Output:
	// 9223372036854775807
	// cast error : strconv.ParseInt: parsing "9223372036854775807": value out of range
}

func ExampleFromBool() {
	var v bool
	if err := FromBool(true, &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	var iv int
	if err := FromBool(true, &iv); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", iv)
	}

	// Output:
	// true
	// 1
}

func ExampleFromBytes() {
	var v string
	if err := FromBytes([]byte("abc"), &v); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", v)
	}

	// Output:
	// abc
}
