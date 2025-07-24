package test

import (
	"math"
	"testing"

	"github.com/cybergarage/go-safecast/safecast"
)

func TestOverflowUnderflow_IntUint(t *testing.T) {
	// ToInt8
	var vi8 int8
	err := safecast.ToInt8(128, &vi8)
	if err == nil {
		t.Errorf("ToInt8: Expected overflow error, got nil")
	}
	err = safecast.ToInt8(-129, &vi8)
	if err == nil {
		t.Errorf("ToInt8: Expected underflow error, got nil")
	}

	// ToInt16
	var vi16 int16
	err = safecast.ToInt16(32768, &vi16)
	if err == nil {
		t.Errorf("ToInt16: Expected overflow error, got nil")
	}
	err = safecast.ToInt16(-32769, &vi16)
	if err == nil {
		t.Errorf("ToInt16: Expected underflow error, got nil")
	}

	// ToInt32
	var vi32 int32
	err = safecast.ToInt32(2147483648, &vi32)
	if err == nil {
		t.Errorf("ToInt32: Expected overflow error, got nil")
	}
	err = safecast.ToInt32(-2147483649, &vi32)
	if err == nil {
		t.Errorf("ToInt32: Expected underflow error, got nil")
	}

	// ToUint8
	var vu8 uint8
	err = safecast.ToUint8(-1, &vu8)
	if err == nil {
		t.Errorf("ToUint8: Expected underflow error, got nil")
	}
	err = safecast.ToUint8(256, &vu8)
	if err == nil {
		t.Errorf("ToUint8: Expected overflow error, got nil")
	}

	// ToUint16
	var vu16 uint16
	err = safecast.ToUint16(-1, &vu16)
	if err == nil {
		t.Errorf("ToUint16: Expected underflow error, got nil")
	}
	err = safecast.ToUint16(65536, &vu16)
	if err == nil {
		t.Errorf("ToUint16: Expected overflow error, got nil")
	}

	// ToUint32
	var vu32 uint32
	err = safecast.ToUint32(-1, &vu32)
	if err == nil {
		t.Errorf("ToUint32: Expected underflow error, got nil")
	}
	err = safecast.ToUint32(4294967296, &vu32)
	if err == nil {
		t.Errorf("ToUint32: Expected overflow error, got nil")
	}

	// FromInt8
	var fi8 int8
	err = safecast.FromInt(128, &fi8)
	if err == nil {
		t.Errorf("FromInt8: Expected overflow error, got nil")
	}
	err = safecast.FromInt(-129, &fi8)
	if err == nil {
		t.Errorf("FromInt8: Expected underflow error, got nil")
	}

	// FromInt16
	var fi16 int16
	err = safecast.FromInt(32768, &fi16)
	if err == nil {
		t.Errorf("FromInt16: Expected overflow error, got nil")
	}
	err = safecast.FromInt(-32769, &fi16)
	if err == nil {
		t.Errorf("FromInt16: Expected underflow error, got nil")
	}

	// FromInt32
	var fi32 int32
	err = safecast.FromInt(2147483648, &fi32)
	if err == nil {
		t.Errorf("FromInt32: Expected overflow error, got nil")
	}
	err = safecast.FromInt(-2147483649, &fi32)
	if err == nil {
		t.Errorf("FromInt32: Expected underflow error, got nil")
	}

	// FromUint8
	var fu8 uint8
	err = safecast.FromInt(256, &fu8)
	if err == nil {
		t.Errorf("FromUint8: Expected overflow error, got nil")
	}
	err = safecast.FromInt(-1, &fu8)
	if err == nil {
		t.Errorf("FromUint8: Expected underflow error, got nil")
	}

	// FromUint16
	var fu16 uint16
	err = safecast.FromInt(65536, &fu16)
	if err == nil {
		t.Errorf("FromUint16: Expected overflow error, got nil")
	}
	err = safecast.FromInt(-1, &fu16)
	if err == nil {
		t.Errorf("FromUint16: Expected underflow error, got nil")
	}

	// FromUint32
	var fu32 uint32
	err = safecast.FromInt(4294967296, &fu32)
	if err == nil {
		t.Errorf("FromUint32: Expected overflow error, got nil")
	}
	err = safecast.FromInt(-1, &fu32)
	if err == nil {
		t.Errorf("FromUint32: Expected underflow error, got nil")
	}
}

func TestToInt8_AllBranches(t *testing.T) {
	var v int8
	// int
	if err := safecast.ToInt8(42, &v); err != nil || v != 42 {
		t.Errorf("int: got %v, err %v", v, err)
	}
	// *int
	ival := 43
	if err := safecast.ToInt8(&ival, &v); err != nil || v != 43 {
		t.Errorf("*int: got %v, err %v", v, err)
	}
	// int8
	if err := safecast.ToInt8(int8(44), &v); err != nil || v != 44 {
		t.Errorf("int8: got %v, err %v", v, err)
	}
	// *int8
	ival8 := int8(45)
	if err := safecast.ToInt8(&ival8, &v); err != nil || v != 45 {
		t.Errorf("*int8: got %v, err %v", v, err)
	}
	// int16
	if err := safecast.ToInt8(int16(46), &v); err != nil || v != 46 {
		t.Errorf("int16: got %v, err %v", v, err)
	}
	// *int16
	ival16 := int16(47)
	if err := safecast.ToInt8(&ival16, &v); err != nil || v != 47 {
		t.Errorf("*int16: got %v, err %v", v, err)
	}
	// int32
	if err := safecast.ToInt8(int32(48), &v); err != nil || v != 48 {
		t.Errorf("int32: got %v, err %v", v, err)
	}
	// *int32
	ival32 := int32(49)
	if err := safecast.ToInt8(&ival32, &v); err != nil || v != 49 {
		t.Errorf("*int32: got %v, err %v", v, err)
	}
	// int64
	if err := safecast.ToInt8(int64(50), &v); err != nil || v != 50 {
		t.Errorf("int64: got %v, err %v", v, err)
	}
	// *int64
	ival64 := int64(51)
	if err := safecast.ToInt8(&ival64, &v); err != nil || v != 51 {
		t.Errorf("*int64: got %v, err %v", v, err)
	}
	// uint
	if err := safecast.ToInt8(uint(52), &v); err != nil || v != 52 {
		t.Errorf("uint: got %v, err %v", v, err)
	}
	// *uint
	uval := uint(53)
	if err := safecast.ToInt8(&uval, &v); err != nil || v != 53 {
		t.Errorf("*uint: got %v, err %v", v, err)
	}
	// uint8
	if err := safecast.ToInt8(uint8(54), &v); err != nil || v != 54 {
		t.Errorf("uint8: got %v, err %v", v, err)
	}
	// *uint8
	uval8 := uint8(55)
	if err := safecast.ToInt8(&uval8, &v); err != nil || v != 55 {
		t.Errorf("*uint8: got %v, err %v", v, err)
	}
	// uint16
	if err := safecast.ToInt8(uint16(56), &v); err != nil || v != 56 {
		t.Errorf("uint16: got %v, err %v", v, err)
	}
	// *uint16
	uval16 := uint16(57)
	if err := safecast.ToInt8(&uval16, &v); err != nil || v != 57 {
		t.Errorf("*uint16: got %v, err %v", v, err)
	}
	// uint32
	if err := safecast.ToInt8(uint32(58), &v); err != nil || v != 58 {
		t.Errorf("uint32: got %v, err %v", v, err)
	}
	// *uint32
	uval32 := uint32(59)
	if err := safecast.ToInt8(&uval32, &v); err != nil || v != 59 {
		t.Errorf("*uint32: got %v, err %v", v, err)
	}
	// uint64
	if err := safecast.ToInt8(uint64(60), &v); err != nil || v != 60 {
		t.Errorf("uint64: got %v, err %v", v, err)
	}
	// *uint64
	uval64 := uint64(61)
	if err := safecast.ToInt8(&uval64, &v); err != nil || v != 61 {
		t.Errorf("*uint64: got %v, err %v", v, err)
	}
	// float32
	if err := safecast.ToInt8(float32(62.0), &v); err != nil || v != 62 {
		t.Errorf("float32: got %v, err %v", v, err)
	}
	// *float32
	fval32 := float32(63.0)
	if err := safecast.ToInt8(&fval32, &v); err != nil || v != 63 {
		t.Errorf("*float32: got %v, err %v", v, err)
	}
	// float64
	if err := safecast.ToInt8(float64(64.0), &v); err != nil || v != 64 {
		t.Errorf("float64: got %v, err %v", v, err)
	}
	// *float64
	fval64 := float64(65.0)
	if err := safecast.ToInt8(&fval64, &v); err != nil || v != 65 {
		t.Errorf("*float64: got %v, err %v", v, err)
	}
	// bool
	if err := safecast.ToInt8(true, &v); err != nil || v != 1 {
		t.Errorf("bool true: got %v, err %v", v, err)
	}
	if err := safecast.ToInt8(false, &v); err != nil || v != 0 {
		t.Errorf("bool false: got %v, err %v", v, err)
	}
	// *bool
	bval := true
	if err := safecast.ToInt8(&bval, &v); err != nil || v != 1 {
		t.Errorf("*bool true: got %v, err %v", v, err)
	}
	bval = false
	if err := safecast.ToInt8(&bval, &v); err != nil || v != 0 {
		t.Errorf("*bool false: got %v, err %v", v, err)
	}
	// string
	if err := safecast.ToInt8("66", &v); err != nil || v != 66 {
		t.Errorf("string: got %v, err %v", v, err)
	}
	// *string
	sval := "67"
	if err := safecast.ToInt8(&sval, &v); err != nil || v != 67 {
		t.Errorf("*string: got %v, err %v", v, err)
	}
	// []byte
	if err := safecast.ToInt8([]byte("68"), &v); err != nil || v != 68 {
		t.Errorf("[]byte: got %v, err %v", v, err)
	}
	// unsupported type
	if err := safecast.ToInt8([]int{1, 2, 3}, &v); err == nil {
		t.Errorf("unsupported type: expected error, got nil")
	}
}

func TestToInt_AllBranches(t *testing.T) {
	var v int
	// int
	if err := safecast.ToInt(42, &v); err != nil || v != 42 {
		t.Errorf("int: got %v, err %v", v, err)
	}
	// *int
	ival := 43
	if err := safecast.ToInt(&ival, &v); err != nil || v != 43 {
		t.Errorf("*int: got %v, err %v", v, err)
	}
	// int8
	if err := safecast.ToInt(int8(44), &v); err != nil || v != 44 {
		t.Errorf("int8: got %v, err %v", v, err)
	}
	// *int8
	ival8 := int8(45)
	if err := safecast.ToInt(&ival8, &v); err != nil || v != 45 {
		t.Errorf("*int8: got %v, err %v", v, err)
	}
	// int16
	if err := safecast.ToInt(int16(46), &v); err != nil || v != 46 {
		t.Errorf("int16: got %v, err %v", v, err)
	}
	// *int16
	ival16 := int16(47)
	if err := safecast.ToInt(&ival16, &v); err != nil || v != 47 {
		t.Errorf("*int16: got %v, err %v", v, err)
	}
	// int32
	if err := safecast.ToInt(int32(48), &v); err != nil || v != 48 {
		t.Errorf("int32: got %v, err %v", v, err)
	}
	// *int32
	ival32 := int32(49)
	if err := safecast.ToInt(&ival32, &v); err != nil || v != 49 {
		t.Errorf("*int32: got %v, err %v", v, err)
	}
	// int64
	if err := safecast.ToInt(int64(50), &v); err != nil || v != 50 {
		t.Errorf("int64: got %v, err %v", v, err)
	}
	// *int64
	ival64 := int64(51)
	if err := safecast.ToInt(&ival64, &v); err != nil || v != 51 {
		t.Errorf("*int64: got %v, err %v", v, err)
	}
	// uint
	if err := safecast.ToInt(uint(52), &v); err != nil || v != 52 {
		t.Errorf("uint: got %v, err %v", v, err)
	}
	// *uint
	uval := uint(53)
	if err := safecast.ToInt(&uval, &v); err != nil || v != 53 {
		t.Errorf("*uint: got %v, err %v", v, err)
	}
	// uint8
	if err := safecast.ToInt(uint8(54), &v); err != nil || v != 54 {
		t.Errorf("uint8: got %v, err %v", v, err)
	}
	// *uint8
	uval8 := uint8(55)
	if err := safecast.ToInt(&uval8, &v); err != nil || v != 55 {
		t.Errorf("*uint8: got %v, err %v", v, err)
	}
	// uint16
	if err := safecast.ToInt(uint16(56), &v); err != nil || v != 56 {
		t.Errorf("uint16: got %v, err %v", v, err)
	}
	// *uint16
	uval16 := uint16(57)
	if err := safecast.ToInt(&uval16, &v); err != nil || v != 57 {
		t.Errorf("*uint16: got %v, err %v", v, err)
	}
	// uint32
	if err := safecast.ToInt(uint32(58), &v); err != nil || v != 58 {
		t.Errorf("uint32: got %v, err %v", v, err)
	}
	// *uint32
	uval32 := uint32(59)
	if err := safecast.ToInt(&uval32, &v); err != nil || v != 59 {
		t.Errorf("*uint32: got %v, err %v", v, err)
	}
	// uint64
	if err := safecast.ToInt(uint64(60), &v); err != nil || v != 60 {
		t.Errorf("uint64: got %v, err %v", v, err)
	}
	// *uint64
	uval64 := uint64(61)
	if err := safecast.ToInt(&uval64, &v); err != nil || v != 61 {
		t.Errorf("*uint64: got %v, err %v", v, err)
	}
	// float32
	if err := safecast.ToInt(float32(62.0), &v); err != nil || v != 62 {
		t.Errorf("float32: got %v, err %v", v, err)
	}
	// *float32
	fval32 := float32(63.0)
	if err := safecast.ToInt(&fval32, &v); err != nil || v != 63 {
		t.Errorf("*float32: got %v, err %v", v, err)
	}
	// float64
	if err := safecast.ToInt(float64(64.0), &v); err != nil || v != 64 {
		t.Errorf("float64: got %v, err %v", v, err)
	}
	// *float64
	fval64 := float64(65.0)
	if err := safecast.ToInt(&fval64, &v); err != nil || v != 65 {
		t.Errorf("*float64: got %v, err %v", v, err)
	}
	// bool
	if err := safecast.ToInt(true, &v); err != nil || v != 1 {
		t.Errorf("bool true: got %v, err %v", v, err)
	}
	if err := safecast.ToInt(false, &v); err != nil || v != 0 {
		t.Errorf("bool false: got %v, err %v", v, err)
	}
	// *bool
	bval := true
	if err := safecast.ToInt(&bval, &v); err != nil || v != 1 {
		t.Errorf("*bool true: got %v, err %v", v, err)
	}
	bval = false
	if err := safecast.ToInt(&bval, &v); err != nil || v != 0 {
		t.Errorf("*bool false: got %v, err %v", v, err)
	}
	// string
	if err := safecast.ToInt("66", &v); err != nil || v != 66 {
		t.Errorf("string: got %v, err %v", v, err)
	}
	// *string
	sval := "67"
	if err := safecast.ToInt(&sval, &v); err != nil || v != 67 {
		t.Errorf("*string: got %v, err %v", v, err)
	}
	// []byte
	if err := safecast.ToInt([]byte("68"), &v); err != nil || v != 68 {
		t.Errorf("[]byte: got %v, err %v", v, err)
	}
	// unsupported type
	if err := safecast.ToInt([]int{1, 2, 3}, &v); err == nil {
		t.Errorf("unsupported type: expected error, got nil")
	}
}

func TestFromInt64_UnderflowBranches(t *testing.T) {
	var u uint
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64

	// Underflow for uint
	if err := safecast.FromInt64(-1, &u); err == nil {
		t.Errorf("FromInt64(-1, &uint): Expected underflow error, got nil")
	}
	if err := safecast.FromInt64(-1, &u8); err == nil {
		t.Errorf("FromInt64(-1, &uint8): Expected underflow error, got nil")
	}
	if err := safecast.FromInt64(-1, &u16); err == nil {
		t.Errorf("FromInt64(-1, &uint16): Expected underflow error, got nil")
	}
	if err := safecast.FromInt64(-1, &u32); err == nil {
		t.Errorf("FromInt64(-1, &uint32): Expected underflow error, got nil")
	}
	if err := safecast.FromInt64(-1, &u64); err == nil {
		t.Errorf("FromInt64(-1, &uint64): Expected underflow error, got nil")
	}
}

func TestFromUint64_OverBranches(t *testing.T) {
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	// Overflow for int
	if err := safecast.FromUint64(uint64(math.MaxInt)+1, &i); err == nil {
		t.Errorf("FromUint64(MaxInt+1, &int): Expected overflow error, got nil")
	}
	if err := safecast.FromUint64(uint64(math.MaxInt8)+1, &i8); err == nil {
		t.Errorf("FromUint64(MaxInt8+1, &int8): Expected overflow error, got nil")
	}
	if err := safecast.FromUint64(uint64(math.MaxInt16)+1, &i16); err == nil {
		t.Errorf("FromUint64(MaxInt16+1, &int16): Expected overflow error, got nil")
	}
	if err := safecast.FromUint64(uint64(math.MaxInt32)+1, &i32); err == nil {
		t.Errorf("FromUint64(MaxInt32+1, &int32): Expected overflow error, got nil")
	}
	if err := safecast.FromUint64(uint64(math.MaxInt64)+1, &i64); err == nil {
		t.Errorf("FromUint64(MaxInt64+1, &int64): Expected overflow error, got nil")
	}
}
