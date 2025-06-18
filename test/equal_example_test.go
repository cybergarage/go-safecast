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

func ExampleEqual_int() {
	a := 42
	b := 42
	eq := safecast.Equal(a, b)
	fmt.Println(eq)
	// Output: true
}

func ExampleEqual_string() {
	a := "hello"
	b := "world"
	eq := safecast.Equal(a, b)
	fmt.Println(eq)
	// Output: false
}

func ExampleEqual_slice() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	eq := safecast.Equal(a, b)
	fmt.Println(eq)
	// Output: true
}

func ExampleEqual_map() {
	a := map[string]int{"foo": 1, "bar": 2}
	b := map[string]int{"foo": 1, "bar": 2}
	eq := safecast.Equal(a, b)
	fmt.Println(eq)
	// Output: true
}

func ExampleEqual_time() {
	a := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	b := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	eq := safecast.Equal(a, b)
	fmt.Println(eq)
	// Output: true
}

func ExampleEqual_struct() {
	type Point struct {
		X int
		Y int
	}
	a := Point{X: 1, Y: 2}
	b := Point{X: 1, Y: 2}
	eq := safecast.Equal(a, b)
	fmt.Println(eq)
	// Output: true
}
