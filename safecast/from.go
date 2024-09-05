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

// From casts an interface to an interface type.
func From(from any, to any) error {
	switch from := from.(type) {
	case int:
		return FromInt(from, to)
	case int8:
		return FromInt8(from, to)
	case int16:
		return FromInt16(from, to)
	case int32:
		return FromInt32(from, to)
	case int64:
		return FromInt64(from, to)
	case uint:
		return FromUint(from, to)
	case uint8:
		return FromUint8(from, to)
	case uint16:
		return FromUint16(from, to)
	case uint32:
		return FromUint32(from, to)
	case uint64:
		return FromUint64(from, to)
	case float32:
		return FromFloat32(from, to)
	case float64:
		return FromFloat64(from, to)
	case string:
		return FromString(from, to)
	case bool:
		return FromBool(from, to)
	case []byte:
		return FromBytes(from, to)
	default:
		return newErrorCast(from, to)
	}
}
