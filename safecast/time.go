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
	DateTime              = "2006-01-02 15:04:05"
	DateTimeZ             = "2006-01-02 15:04:05 MST"
	DateTimeNZ            = "2006-01-02 15:04:05 -0700"
)

// SupportedTimeLayouts is a list of supported time layouts.
var SupportedTimeLayouts = []string{
	DateTime,
	DateTimeZ,
	DateTimeNZ,
	ISO860TimestampLayout,
	time.RFC3339,
	time.RFC3339Nano,
	time.RFC1123,
	time.RFC1123Z,
}

// ToTime casts an interface to a time.Time.
func ToTime(from any, to *time.Time, layouts ...string) error {
	parseTimeString := func(s string, to *time.Time) error {
		var err error
		if len(layouts) == 0 {
			layouts = SupportedTimeLayouts
		}
		for _, layout := range layouts {
			*to, err = time.Parse(layout, s)
			if err == nil {
				return nil
			}
		}
		return err
	}
	switch from := from.(type) {
	case time.Time:
		*to = from
		return nil
	case *time.Time:
		*to = *from
		return nil
	case string:
		return parseTimeString(from, to)
	case *string:
		return parseTimeString(*from, to)
	case []byte:
		return parseTimeString(string(from), to)
	}
	return newErrorCast(from, to)
}
