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

// To casts an interface to an interface type.
func To(from any, to any) error {
	switch to := to.(type) {
	case *int:
		return ToInt(from, to)
	case *int8:
		return ToInt8(from, to)
	case *int16:
		return ToInt16(from, to)
	case *int32:
		return ToInt32(from, to)
	case *int64:
		return ToInt64(from, to)
	case *uint:
		return ToUint(from, to)
	case *uint8:
		return ToUint8(from, to)
	case *uint16:
		return ToUint16(from, to)
	case *uint32:
		return ToUint32(from, to)
	case *uint64:
		return ToUint64(from, to)
	case *float32:
		return ToFloat32(from, to)
	case *float64:
		return ToFloat64(from, to)
	case *string:
		return ToString(from, to)
	case *bool:
		return ToBool(from, to)
	case *[]byte:
		return ToBytes(from, to)
	case *time.Time:
		return ToTime(from, to)
	default:
		return newErrorCast(from, to)
	}
}
