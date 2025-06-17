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

func TestCompare(t *testing.T) {
	t.Run("Ints", func(t *testing.T) {
		if cmp, _ := safecast.Compare(42, 42); cmp != 0 {
			t.Errorf("Expected 0, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(42, int64(42)); cmp != 0 {
			t.Errorf("Expected 0, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(42, 43); cmp >= 0 {
			t.Errorf("Expected negative, got %d", cmp)
		}

		i := 100
		if cmp, _ := safecast.Compare(&i, 100); cmp != 0 {
			t.Errorf("Expected 0 for *int/int, got %d", cmp)
		}
		if cmp, _ := safecast.Compare(100, &i); cmp != 0 {
			t.Errorf("Expected 0 for int/*int, got %d", cmp)
		}
		j := 101
		if cmp, _ := safecast.Compare(&i, &j); cmp >= 0 {
			t.Errorf("Expected negative for different *int/*int, got %d", cmp)
		}
	})
}
