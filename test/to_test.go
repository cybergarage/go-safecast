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
	"testing"
	"time"

	"github.com/cybergarage/go-safecast/safecast"
)

func TestToCast(t *testing.T) {
	t.Run("ToTime", func(t *testing.T) {
		tests := []struct {
			from   string
			layout string
		}{
			{
				from:   "2022-01-01",
				layout: safecast.ISO8601DateLayout,
			},
			{
				from:   "00:00:00",
				layout: safecast.ISO8601TimeLayout,
			},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("%s (%s)", test.from, test.layout), func(t *testing.T) {
				var to time.Time
				if err := safecast.ToTime(test.from, test.layout, &to); err != nil {
					t.Error(err)
					return
				}
			})
		}
	})
}
