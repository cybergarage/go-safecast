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
	"time"

	"github.com/cybergarage/go-safecast/safecast"
)

func ExampleCompare_int() {
	cmp, _ := safecast.Compare(42, 42)
	fmt.Println(cmp)
	cmp, _ = safecast.Compare(42, 43)
	fmt.Println(cmp)
	cmp, _ = safecast.Compare(43, 42)
	fmt.Println(cmp)
	// Output:
	// 0
	// -1
	// 1
}

func ExampleCompare_string() {
	cmp, _ := safecast.Compare("abc", "abc")
	fmt.Println(cmp)
	cmp, _ = safecast.Compare("abc", "def")
	fmt.Println(cmp)
	cmp, _ = safecast.Compare("def", "abc")
	fmt.Println(cmp)
	// Output:
	// 0
	// -1
	// 1
}

func ExampleCompare_float() {
	cmp, _ := safecast.Compare(3.14, 3.14)
	fmt.Println(cmp)
	cmp, _ = safecast.Compare(3.14, 2.71)
	fmt.Println(cmp)
	cmp, _ = safecast.Compare(2.71, 3.14)
	fmt.Println(cmp)
	// Output:
	// 0
	// 1
	// -1
}

func ExampleCompare_time() {
	t1 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	cmp, _ := safecast.Compare(t1, t2)
	fmt.Println(cmp)
	cmp, _ = safecast.Compare(t1, t3)
	fmt.Println(cmp)
	cmp, _ = safecast.Compare(t3, t1)
	fmt.Println(cmp)
	// Output:
	// 0
	// -1
	// 1
}
