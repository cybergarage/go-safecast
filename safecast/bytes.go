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

// FromBytes casts an interface to a byte slice type.
func FromBytes(from []byte, to any) error {
	switch to := to.(type) {
	case *string:
		*to = string(from)
	case *[]byte:
		*to = from
	default:
		return newErrorCast(from, to)
	}
	return nil
}

// ToBytes casts an interface to a byte slice type.
func ToBytes(from any, to *[]byte) error {
	switch from := from.(type) {
	case string:
		*to = []byte(from)
	case []byte:
		*to = from
	default:
		return newErrorCast(from, to)
	}
	return nil
}
