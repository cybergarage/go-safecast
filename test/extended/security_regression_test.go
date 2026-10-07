package test

import (
	"errors"
	"math"
	"strconv"
	"testing"

	"github.com/cybergarage/go-safecast/safecast"
)

func TestUnsignedNarrowingBounds(t *testing.T) {
	checkUnsignedBounds(t, safecast.ToUint8, int16(0), int16(255), int16(256), int16(257), int16(-1))
	checkUnsignedBounds(t, safecast.ToUint8, int32(0), int32(255), int32(256), int32(257), int32(-1))
	checkUnsignedBounds(t, safecast.ToUint16, int32(0), int32(65535), int32(65536), int32(65537), int32(-1))
}

func checkUnsignedBounds[S ~int16 | ~int32, D ~uint8 | ~uint16](t *testing.T, convert func(any, *D) error, values ...S) {
	t.Helper()
	for i, value := range values {
		for _, input := range []any{value, &value} {
			for _, generic := range []bool{false, true} {
				var out D
				var err error
				if generic {
					err = safecast.To(input, &out)
				} else {
					err = convert(input, &out)
				}
				if i < 2 {
					if err != nil || out != D(value) {
						t.Fatalf("valid %T(%v) -> %T: out=%v err=%v", input, input, out, out, err)
					}
				} else if !errors.Is(err, safecast.ErrCast) {
					t.Fatalf("invalid %T(%v) -> %T: expected ErrCast, got out=%v err=%v", input, input, out, out, err)
				}
			}
		}
	}
}

func TestFromFloatExclusiveBounds(t *testing.T) {
	checkFloatBounds[int64](t, 0x1p63, -0x1p63)
	checkFloatBounds[uint64](t, 0x1p64, 0)
	checkFloatBounds[int](t, math.Ldexp(1, strconv.IntSize-1), -math.Ldexp(1, strconv.IntSize-1))
	checkFloatBounds[uint](t, math.Ldexp(1, strconv.IntSize), 0)
}

func checkFloatBounds[D ~int | ~int64 | ~uint | ~uint64](t *testing.T, upper, lower float64) {
	t.Helper()
	for _, single := range []bool{false, true} {
		below := math.Floor(math.Nextafter(upper, 0))
		above := math.Nextafter(upper, math.Inf(1))
		if single {
			below = float64(math.Nextafter32(float32(upper), 0))
			above = float64(math.Nextafter32(float32(upper), float32(math.Inf(1))))
		}
		for i, value := range []float64{lower, 0, 42.75, below, upper, above} {
			for _, route := range []string{"direct", "generic", "pointer"} {
				out := D(7)
				var err error
				switch {
				case single:
					input := float32(value)
					switch route {
					case "direct":
						err = safecast.FromFloat32(input, &out)
					case "generic":
						err = safecast.From(input, &out)
					default:
						err = safecast.From(&input, &out)
					}
				case route == "direct":
					err = safecast.FromFloat64(value, &out)
				case route == "generic":
					err = safecast.From(value, &out)
				default:
					err = safecast.From(&value, &out)
				}
				if i < 4 {
					if err != nil || out != D(value) {
						t.Fatalf("valid float(%v) -> %T (%s, single=%v): out=%v err=%v", value, out, route, single, out, err)
					}
				} else if !errors.Is(err, safecast.ErrCast) || out != D(7) {
					t.Fatalf("overflow float(%v) -> %T (%s, single=%v): out=%v err=%v", value, out, route, single, out, err)
				}
			}
		}
	}
}

func TestCompareFallbackOrdering(t *testing.T) {
	for _, pair := range [][2]any{
		{int8(100), int16(256)},
		{int16(100), int32(65536)},
		{int32(100), int64(1 << 32)},
		{uint8(100), uint16(256)},
		{int8(-100), int16(-256)},
	} {
		want := -1
		if pair[0] == int8(-100) {
			want = 1
		}
		for i := range 2 {
			got, err := safecast.Compare(pair[i], pair[1-i])
			if err != nil || got != want {
				t.Fatalf("Compare(%v, %v) = %v, %v; want %v", pair[i], pair[1-i], got, err, want)
			}
			want = -want
		}
	}
	if got, err := safecast.Compare(int8(100), int16(100)); err != nil || got != 0 {
		t.Fatalf("equal mixed-width values: got %v, %v", got, err)
	}
	if _, err := safecast.Compare(struct{}{}, 1); err == nil {
		t.Fatal("unsupported values must retain a comparison error")
	}
}
