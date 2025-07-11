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
	"errors"
	"fmt"
)

// ErrCast is returned when a value cannot be cast to the desired type.
var ErrCast = errors.New("cast error")

// ErrNil is returned when a nil value is passed to a casting function that requires a non-nil value.
var ErrNil = errors.New("nil")

const (
	errorCastType   = "%w : %T (%v) => %T"
	errorOverRange  = "%w : out of range %v > %T"
	errorUnderRange = "%w : out of range %v < %T"
	errorSimple     = "%w : %s"
	errorCompare    = "%w : %T (%v) != %T (%v)"
)

func newErrorCast(fromItem any, toItem any) error {
	return fmt.Errorf(errorCastType, ErrCast, fromItem, fromItem, toItem)
}

func newErrorOverRange(fromItem any, toItem any) error {
	return fmt.Errorf(errorOverRange, ErrCast, fromItem, toItem)
}

func newErrorUnderRange(fromItem any, toItem any) error {
	return fmt.Errorf(errorUnderRange, ErrCast, fromItem, toItem)
}

func newErrorWithError(err error) error {
	return fmt.Errorf(errorSimple, ErrCast, err.Error())
}

func newCompareError(item any, otherItem any) error {
	return fmt.Errorf(errorCompare, ErrCast, item, item, otherItem, otherItem)
}
