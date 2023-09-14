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

// SupportedTimeLayouts is a list of supported time layouts.
var SupportedTimeLayouts = []string{
	ISO860TimestampLayout,
	time.DateTime,
}

// ToTime casts an interface to a time.Time.
func ToTime(from any, to *time.Time, layouts ...string) error {
	switch from := from.(type) {
	case time.Time:
		*to = from
		return nil
	case string:
		var err error
		if len(layouts) == 0 {
			layouts = SupportedTimeLayouts
		}
		for _, layout := range layouts {
			*to, err = time.Parse(layout, from)
			if err == nil {
				return nil
			}
		}
		return err
	}
	return newErrorCast(from, to)
}
