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

func ExampleFromInt() {
	var vi int
	if err := safecast.FromInt(math.MaxInt, &vi); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", vi)
	}

	var vi8 int8
	if err := safecast.FromInt(math.MaxInt, &vi8); err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", vi8)
	}

	// Output:
	// 9223372036854775807
	// cast error : overflow 9223372036854775807 (int) => *int8
}
