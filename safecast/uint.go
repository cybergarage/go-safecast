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

import "math"

func FromUint64(from uint64, to any) error {
	switch to := to.(type) {
	case *int:
		if math.MaxInt < from {
			return newErrorOverflow(from, to)
		}
		*to = int(from)
	case *int8:
		if math.MaxInt8 < from {
			return newErrorOverflow(from, to)
		}
		*to = int8(from)
	case *int16:
		if math.MaxInt16 < from {
			return newErrorOverflow(from, to)
		}
		*to = int16(from)
	case *int32:
		if math.MaxInt32 < from {
			return newErrorOverflow(from, to)
		}
		*to = int32(from)
	case *int64:
		if math.MaxInt64 < from {
			return newErrorOverflow(from, to)
		}
		*to = int64(from)
	case *uint:
		*to = uint(from)
	case *uint8:
		if math.MaxUint8 < from {
			return newErrorOverflow(from, to)
		}
		*to = uint8(from)
	case *uint16:
		if math.MaxUint16 < from {
			return newErrorOverflow(from, to)
		}
		*to = uint16(from)
	case *uint32:
		if math.MaxUint32 < from {
			return newErrorOverflow(from, to)
		}
		*to = uint32(from)
	case *uint64:
		*to = from
	case *float32:
		*to = float32(from)
	case *float64:
		*to = float64(from)
	default:
		return newErrorCast(from, to)
	}
	return nil
}

func FromUint8(from uint8, to any) error {
	return FromUint64(uint64(from), to)
}

func FromUint16(from uint16, to any) error {
	return FromUint64(uint64(from), to)
}

func FromUint32(from uint32, to any) error {
	return FromUint64(uint64(from), to)
}

func FromUint(from uint, to any) error {
	return FromUint64(uint64(from), to)
}
