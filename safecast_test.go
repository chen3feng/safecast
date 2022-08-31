package safecast_test

import (
	_ "fmt"
	"math"
	"testing"

	"github.com/chen3feng/safecast"
)

func expectFalse(t *testing.T, value bool) {
	if value {
		t.Fail()
	}
}

func expectTrue(t *testing.T, value bool) {
	if !value {
		t.Fail()
	}
}

func TestTo_int8_to_uint8(t *testing.T) {
	var i int8 = -1
	_, ok := safecast.To[uint8](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint8](i)
	expectTrue(t, ok)
}

func TestTo_int8_to_uint16(t *testing.T) {
	var i int8 = -1
	_, ok := safecast.To[uint16](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint16](i)
	expectTrue(t, ok)
}

func TestTo_int8_to_uint32(t *testing.T) {
	var i int8 = -1
	_, ok := safecast.To[uint32](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint32](i)
	expectTrue(t, ok)
}

func TestTo_int8_to_uint64(t *testing.T) {
	var i int8 = -1
	_, ok := safecast.To[uint64](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_int8_to_uint(t *testing.T) {
	var i int8 = -1
	_, ok := safecast.To[uint](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_uint8(t *testing.T) {
	var i int16 = -1
	_, ok := safecast.To[uint8](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint8](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_uint16(t *testing.T) {
	var i int16 = -1
	_, ok := safecast.To[uint16](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint16](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_uint32(t *testing.T) {
	var i int16 = -1
	_, ok := safecast.To[uint32](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint32](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_uint64(t *testing.T) {
	var i int16 = -1
	_, ok := safecast.To[uint64](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_uint(t *testing.T) {
	var i int16 = -1
	_, ok := safecast.To[uint](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_uint8(t *testing.T) {
	var i int32 = -1
	_, ok := safecast.To[uint8](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint8](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_uint16(t *testing.T) {
	var i int32 = -1
	_, ok := safecast.To[uint16](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint16](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_uint32(t *testing.T) {
	var i int32 = -1
	_, ok := safecast.To[uint32](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint32](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_uint64(t *testing.T) {
	var i int32 = -1
	_, ok := safecast.To[uint64](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_uint(t *testing.T) {
	var i int32 = -1
	_, ok := safecast.To[uint](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_uint8(t *testing.T) {
	var i int64 = -1
	_, ok := safecast.To[uint8](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint8](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_uint16(t *testing.T) {
	var i int64 = -1
	_, ok := safecast.To[uint16](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint16](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_uint32(t *testing.T) {
	var i int64 = -1
	_, ok := safecast.To[uint32](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint32](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_uint64(t *testing.T) {
	var i int64 = -1
	_, ok := safecast.To[uint64](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_uint(t *testing.T) {
	var i int64 = -1
	_, ok := safecast.To[uint](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_int_to_uint8(t *testing.T) {
	var i int = -1
	_, ok := safecast.To[uint8](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint8](i)
	expectTrue(t, ok)
}

func TestTo_int_to_uint16(t *testing.T) {
	var i int = -1
	_, ok := safecast.To[uint16](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint16](i)
	expectTrue(t, ok)
}

func TestTo_int_to_uint32(t *testing.T) {
	var i int = -1
	_, ok := safecast.To[uint32](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint32](i)
	expectTrue(t, ok)
}

func TestTo_int_to_uint64(t *testing.T) {
	var i int = -1
	_, ok := safecast.To[uint64](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_int_to_uint(t *testing.T) {
	var i int = -1
	_, ok := safecast.To[uint](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_int8_to_int8(t *testing.T) {
	var i int8 = -1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int8](i)
	expectTrue(t, ok)
}

func TestTo_int8_to_int16(t *testing.T) {
	var i int8 = -1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int16](i)
	expectTrue(t, ok)
}

func TestTo_int8_to_int32(t *testing.T) {
	var i int8 = -1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int32](i)
	expectTrue(t, ok)
}

func TestTo_int8_to_int64(t *testing.T) {
	var i int8 = -1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int64](i)
	expectTrue(t, ok)
}

func TestTo_int8_to_int(t *testing.T) {
	var i int8 = -1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_int8(t *testing.T) {
	var i int16 = -1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int8](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_int16(t *testing.T) {
	var i int16 = -1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int16](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_int32(t *testing.T) {
	var i int16 = -1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int32](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_int64(t *testing.T) {
	var i int16 = -1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int64](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_int(t *testing.T) {
	var i int16 = -1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_int8(t *testing.T) {
	var i int32 = -1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int8](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_int16(t *testing.T) {
	var i int32 = -1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int16](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_int32(t *testing.T) {
	var i int32 = -1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int32](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_int64(t *testing.T) {
	var i int32 = -1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int64](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_int(t *testing.T) {
	var i int32 = -1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_int8(t *testing.T) {
	var i int64 = -1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int8](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_int16(t *testing.T) {
	var i int64 = -1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int16](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_int32(t *testing.T) {
	var i int64 = -1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int32](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_int64(t *testing.T) {
	var i int64 = -1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int64](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_int(t *testing.T) {
	var i int64 = -1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int](i)
	expectTrue(t, ok)
}

func TestTo_int_to_int8(t *testing.T) {
	var i int = -1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int8](i)
	expectTrue(t, ok)
}

func TestTo_int_to_int16(t *testing.T) {
	var i int = -1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int16](i)
	expectTrue(t, ok)
}

func TestTo_int_to_int32(t *testing.T) {
	var i int = -1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int32](i)
	expectTrue(t, ok)
}

func TestTo_int_to_int64(t *testing.T) {
	var i int = -1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int64](i)
	expectTrue(t, ok)
}

func TestTo_int_to_int(t *testing.T) {
	var i int = -1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[int](i)
	expectTrue(t, ok)
}

func TestTo_uint8_to_int8(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)

	i = math.MaxInt8 + 1
	_, ok = safecast.To[int8](i)
	expectFalse(t, ok)
}

func TestTo_uint8_to_int16(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)
}

func TestTo_uint8_to_int32(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)
}

func TestTo_uint8_to_int64(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)
}

func TestTo_uint8_to_int(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
}

func TestTo_uint16_to_int8(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)

	i = math.MaxInt8 + 1
	_, ok = safecast.To[int8](i)
	expectFalse(t, ok)
}

func TestTo_uint16_to_int16(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)

	i = math.MaxInt16 + 1
	_, ok = safecast.To[int16](i)
	expectFalse(t, ok)
}

func TestTo_uint16_to_int32(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)
}

func TestTo_uint16_to_int64(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)
}

func TestTo_uint16_to_int(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
}

func TestTo_uint32_to_int8(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)

	i = math.MaxInt8 + 1
	_, ok = safecast.To[int8](i)
	expectFalse(t, ok)
}

func TestTo_uint32_to_int16(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)

	i = math.MaxInt16 + 1
	_, ok = safecast.To[int16](i)
	expectFalse(t, ok)
}

func TestTo_uint32_to_int32(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)

	i = math.MaxInt32 + 1
	_, ok = safecast.To[int32](i)
	expectFalse(t, ok)
}

func TestTo_uint32_to_int64(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)
}

func TestTo_uint32_to_int(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
}

func TestTo_uint64_to_int8(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)

	i = math.MaxInt8 + 1
	_, ok = safecast.To[int8](i)
	expectFalse(t, ok)
}

func TestTo_uint64_to_int16(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)

	i = math.MaxInt16 + 1
	_, ok = safecast.To[int16](i)
	expectFalse(t, ok)
}

func TestTo_uint64_to_int32(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)

	i = math.MaxInt32 + 1
	_, ok = safecast.To[int32](i)
	expectFalse(t, ok)
}

func TestTo_uint64_to_int64(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)

	i = math.MaxInt64 + 1
	_, ok = safecast.To[int64](i)
	expectFalse(t, ok)
}

func TestTo_uint64_to_int(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_int8(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_int16(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_int32(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_int64(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_int(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
}

func TestTo_uint8_to_uint8(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[uint8](i)
	expectTrue(t, ok)
}

func TestTo_uint8_to_uint16(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[uint16](i)
	expectTrue(t, ok)
}

func TestTo_uint8_to_uint32(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[uint32](i)
	expectTrue(t, ok)
}

func TestTo_uint8_to_uint64(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_uint8_to_uint(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_uint16_to_uint8(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[uint8](i)
	expectTrue(t, ok)

	i = math.MaxUint8 + 1
	_, ok = safecast.To[uint8](i)
	expectFalse(t, ok)
}

func TestTo_uint16_to_uint16(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[uint16](i)
	expectTrue(t, ok)
}

func TestTo_uint16_to_uint32(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[uint32](i)
	expectTrue(t, ok)
}

func TestTo_uint16_to_uint64(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_uint16_to_uint(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_uint32_to_uint8(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[uint8](i)
	expectTrue(t, ok)

	i = math.MaxUint8 + 1
	_, ok = safecast.To[uint8](i)
	expectFalse(t, ok)
}

func TestTo_uint32_to_uint16(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[uint16](i)
	expectTrue(t, ok)

	i = math.MaxUint16 + 1
	_, ok = safecast.To[uint16](i)
	expectFalse(t, ok)
}

func TestTo_uint32_to_uint32(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[uint32](i)
	expectTrue(t, ok)
}

func TestTo_uint32_to_uint64(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_uint32_to_uint(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_uint64_to_uint8(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[uint8](i)
	expectTrue(t, ok)

	i = math.MaxUint8 + 1
	_, ok = safecast.To[uint8](i)
	expectFalse(t, ok)
}

func TestTo_uint64_to_uint16(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[uint16](i)
	expectTrue(t, ok)

	i = math.MaxUint16 + 1
	_, ok = safecast.To[uint16](i)
	expectFalse(t, ok)
}

func TestTo_uint64_to_uint32(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[uint32](i)
	expectTrue(t, ok)

	i = math.MaxUint32 + 1
	_, ok = safecast.To[uint32](i)
	expectFalse(t, ok)
}

func TestTo_uint64_to_uint64(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_uint64_to_uint(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_uint8(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[uint8](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_uint16(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[uint16](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_uint32(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[uint32](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_uint64(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_uint(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_int8_to_float32(t *testing.T) {
	var i int8 = 1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_uint8_to_float32(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_int8_to_float64(t *testing.T) {
	var i int8 = 1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
}

func TestTo_uint8_to_float64(t *testing.T) {
	var i uint8 = 1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_float32(t *testing.T) {
	var i int16 = 1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_uint16_to_float32(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_int16_to_float64(t *testing.T) {
	var i int16 = 1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
}

func TestTo_uint16_to_float64(t *testing.T) {
	var i uint16 = 1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_float32(t *testing.T) {
	var i int32 = 1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_uint32_to_float32(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_int32_to_float64(t *testing.T) {
	var i int32 = 1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
}

func TestTo_uint32_to_float64(t *testing.T) {
	var i uint32 = 1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_float32(t *testing.T) {
	var i int64 = 1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_uint64_to_float32(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_int64_to_float64(t *testing.T) {
	var i int64 = 1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
}

func TestTo_uint64_to_float64(t *testing.T) {
	var i uint64 = 1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
}

func TestTo_int_to_float32(t *testing.T) {
	var i int = 1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_float32(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_int_to_float64(t *testing.T) {
	var i int = 1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
}

func TestTo_uint_to_float64(t *testing.T) {
	var i uint = 1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
}

func TestTo_float32_to_int8(t *testing.T) {
	var i float32 = 1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)
	i = math.MaxInt8
	i *= 2
	_, ok = safecast.To[int8](i)
	expectFalse(t, ok)
}

func TestTo_float32_to_int16(t *testing.T) {
	var i float32 = 1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)
	i = math.MaxInt16
	i *= 2
	_, ok = safecast.To[int16](i)
	expectFalse(t, ok)
}

func TestTo_float32_to_int32(t *testing.T) {
	var i float32 = 1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)
	i = math.MaxInt32
	i *= 2
	_, ok = safecast.To[int32](i)
	expectFalse(t, ok)
}

func TestTo_float32_to_int64(t *testing.T) {
	var i float32 = 1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)
	i = math.MaxInt64
	i *= 2
	_, ok = safecast.To[int64](i)
	expectFalse(t, ok)
}

func TestTo_float32_to_int(t *testing.T) {
	var i float32 = 1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
	i = math.MaxInt
	i *= 2
	_, ok = safecast.To[int](i)
	expectFalse(t, ok)
}

func TestTo_float64_to_int8(t *testing.T) {
	var i float64 = 1
	_, ok := safecast.To[int8](i)
	expectTrue(t, ok)
	i = math.MaxInt8
	i *= 2
	_, ok = safecast.To[int8](i)
	expectFalse(t, ok)
}

func TestTo_float64_to_int16(t *testing.T) {
	var i float64 = 1
	_, ok := safecast.To[int16](i)
	expectTrue(t, ok)
	i = math.MaxInt16
	i *= 2
	_, ok = safecast.To[int16](i)
	expectFalse(t, ok)
}

func TestTo_float64_to_int32(t *testing.T) {
	var i float64 = 1
	_, ok := safecast.To[int32](i)
	expectTrue(t, ok)
	i = math.MaxInt32
	i *= 2
	_, ok = safecast.To[int32](i)
	expectFalse(t, ok)
}

func TestTo_float64_to_int64(t *testing.T) {
	var i float64 = 1
	_, ok := safecast.To[int64](i)
	expectTrue(t, ok)
	i = math.MaxInt64
	i *= 2
	_, ok = safecast.To[int64](i)
	expectFalse(t, ok)
}

func TestTo_float64_to_int(t *testing.T) {
	var i float64 = 1
	_, ok := safecast.To[int](i)
	expectTrue(t, ok)
	i = math.MaxInt
	i *= 2
	_, ok = safecast.To[int](i)
	expectFalse(t, ok)
}

func TestTo_float32_to_uint8(t *testing.T) {
	var i float32 = -1
	_, ok := safecast.To[uint8](i)
	expectFalse(t, ok)
	i = math.MaxUint8
	i *= 2
	_, ok = safecast.To[uint8](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint8](i)
	expectTrue(t, ok)
}

func TestTo_float32_to_uint16(t *testing.T) {
	var i float32 = -1
	_, ok := safecast.To[uint16](i)
	expectFalse(t, ok)
	i = math.MaxUint16
	i *= 2
	_, ok = safecast.To[uint16](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint16](i)
	expectTrue(t, ok)
}

func TestTo_float32_to_uint32(t *testing.T) {
	var i float32 = -1
	_, ok := safecast.To[uint32](i)
	expectFalse(t, ok)
	i = math.MaxUint32
	i *= 2
	_, ok = safecast.To[uint32](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint32](i)
	expectTrue(t, ok)
}

func TestTo_float32_to_uint64(t *testing.T) {
	var i float32 = -1
	_, ok := safecast.To[uint64](i)
	expectFalse(t, ok)
	i = math.MaxUint64
	i *= 2
	_, ok = safecast.To[uint64](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_float32_to_uint(t *testing.T) {
	var i float32 = -1
	_, ok := safecast.To[uint](i)
	expectFalse(t, ok)
	i = math.MaxUint
	i *= 2
	_, ok = safecast.To[uint](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_float64_to_uint8(t *testing.T) {
	var i float64 = -1
	_, ok := safecast.To[uint8](i)
	expectFalse(t, ok)
	i = math.MaxUint8
	i *= 2
	_, ok = safecast.To[uint8](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint8](i)
	expectTrue(t, ok)
}

func TestTo_float64_to_uint16(t *testing.T) {
	var i float64 = -1
	_, ok := safecast.To[uint16](i)
	expectFalse(t, ok)
	i = math.MaxUint16
	i *= 2
	_, ok = safecast.To[uint16](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint16](i)
	expectTrue(t, ok)
}

func TestTo_float64_to_uint32(t *testing.T) {
	var i float64 = -1
	_, ok := safecast.To[uint32](i)
	expectFalse(t, ok)
	i = math.MaxUint32
	i *= 2
	_, ok = safecast.To[uint32](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint32](i)
	expectTrue(t, ok)
}

func TestTo_float64_to_uint64(t *testing.T) {
	var i float64 = -1
	_, ok := safecast.To[uint64](i)
	expectFalse(t, ok)
	i = math.MaxUint64
	i *= 2
	_, ok = safecast.To[uint64](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint64](i)
	expectTrue(t, ok)
}

func TestTo_float64_to_uint(t *testing.T) {
	var i float64 = -1
	_, ok := safecast.To[uint](i)
	expectFalse(t, ok)
	i = math.MaxUint
	i *= 2
	_, ok = safecast.To[uint](i)
	expectFalse(t, ok)
	i = 1
	_, ok = safecast.To[uint](i)
	expectTrue(t, ok)
}

func TestTo_float32_to_float32(t *testing.T) {
	var i float32 = -1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_float32_to_float64(t *testing.T) {
	var i float32 = -1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[float64](i)
	expectTrue(t, ok)
}

func TestTo_float64_to_float32(t *testing.T) {
	var i float64 = -1
	_, ok := safecast.To[float32](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[float32](i)
	expectTrue(t, ok)
}

func TestTo_float64_to_float64(t *testing.T) {
	var i float64 = -1
	_, ok := safecast.To[float64](i)
	expectTrue(t, ok)
	i = 1
	_, ok = safecast.To[float64](i)
	expectTrue(t, ok)
}
