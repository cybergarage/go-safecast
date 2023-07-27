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
	"time"
)

const (
	ISO8601DateLayout     = "2006-01-02"
	ISO8601TimeLayout     = "15:04:05"
	ISO860TimestampLayout = "2006-01-02T15:04:05"
)

// ToInt64 casts an interface to an int64 type.
func ToTime(from any, layout string, to *time.Time) error {
	switch from := from.(type) {
	case time.Time:
		*to = from
		return nil
	case string:
		var err error
		*to, err = time.Parse(layout, from)
		if err != nil {
			return err
		}
		return nil
	}
	return newErrorCast(from, to)
}
