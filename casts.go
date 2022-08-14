package safecast

import "math"

// IntToInt8 converts the int value to int8 safely.
func IntToInt8(value int) (result int8, ok bool) {
	if int(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// IntToInt16 converts the int value to int16 safely.
func IntToInt16(value int) (result int16, ok bool) {
	if int(int16(value)) != value {
		return int16(value), false
	}
	return int16(value), true
}

// IntToInt32 converts the int value to int32 safely.
func IntToInt32(value int) (result int32, ok bool) {
	if math.MaxInt == math.MaxInt64 && int(int32(value)) != value {
		return int32(value), false
	}
	return int32(value), true
}

// IntToInt64 converts the int value to int64 safely.
func IntToInt64(value int) (result int64, ok bool) {
	return int64(value), true
}

// IntToUint converts the int value to uint safely.
func IntToUint(value int) (result uint, ok bool) {
	if value < 0 {
		return uint(value), false
	}
	return uint(value), true
}

// IntToUint8 converts the int value to uint8 safely.
func IntToUint8(value int) (result uint8, ok bool) {
	if value < 0 {
		return uint8(value), false
	}
	if value > math.MaxUint8 {
		return uint8(value), false
	}
	return uint8(value), true
}

// IntToUint16 converts the int value to uint16 safely.
func IntToUint16(value int) (result uint16, ok bool) {
	if value < 0 {
		return uint16(value), false
	}
	if value > math.MaxUint16 {
		return uint16(value), false
	}
	return uint16(value), true
}

// IntToUint32 converts the int value to uint32 safely.
func IntToUint32(value int) (result uint32, ok bool) {
	if value < 0 {
		return uint32(value), false
	}
	if value > math.MaxUint32 {
		return uint32(value), false
	}
	return uint32(value), true
}

// IntToUint64 converts the int value to uint64 safely.
func IntToUint64(value int) (result uint64, ok bool) {
	if value < 0 {
		return uint64(value), false
	}
	return uint64(value), true
}

// Int8ToInt converts the int8 value to int safely.
func Int8ToInt(value int8) (result int, ok bool) {
	return int(value), true
}

// Int8ToInt16 converts the int8 value to int16 safely.
func Int8ToInt16(value int8) (result int16, ok bool) {
	return int16(value), true
}

// Int8ToInt32 converts the int8 value to int32 safely.
func Int8ToInt32(value int8) (result int32, ok bool) {
	return int32(value), true
}

// Int8ToInt64 converts the int8 value to int64 safely.
func Int8ToInt64(value int8) (result int64, ok bool) {
	return int64(value), true
}

// Int8ToUint converts the int8 value to uint safely.
func Int8ToUint(value int8) (result uint, ok bool) {
	if value < 0 {
		return uint(value), false
	}
	return uint(value), true
}

// Int8ToUint8 converts the int8 value to uint8 safely.
func Int8ToUint8(value int8) (result uint8, ok bool) {
	if value < 0 {
		return uint8(value), false
	}
	return uint8(value), true
}

// Int8ToUint16 converts the int8 value to uint16 safely.
func Int8ToUint16(value int8) (result uint16, ok bool) {
	if value < 0 {
		return uint16(value), false
	}
	return uint16(value), true
}

// Int8ToUint32 converts the int8 value to uint32 safely.
func Int8ToUint32(value int8) (result uint32, ok bool) {
	if value < 0 {
		return uint32(value), false
	}
	return uint32(value), true
}

// Int8ToUint64 converts the int8 value to uint64 safely.
func Int8ToUint64(value int8) (result uint64, ok bool) {
	if value < 0 {
		return uint64(value), false
	}
	return uint64(value), true
}

// Int16ToInt converts the int16 value to int safely.
func Int16ToInt(value int16) (result int, ok bool) {
	return int(value), true
}

// Int16ToInt8 converts the int16 value to int8 safely.
func Int16ToInt8(value int16) (result int8, ok bool) {
	if int16(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// Int16ToInt32 converts the int16 value to int32 safely.
func Int16ToInt32(value int16) (result int32, ok bool) {
	return int32(value), true
}

// Int16ToInt64 converts the int16 value to int64 safely.
func Int16ToInt64(value int16) (result int64, ok bool) {
	return int64(value), true
}

// Int16ToUint converts the int16 value to uint safely.
func Int16ToUint(value int16) (result uint, ok bool) {
	if value < 0 {
		return uint(value), false
	}
	return uint(value), true
}

// Int16ToUint8 converts the int16 value to uint8 safely.
func Int16ToUint8(value int16) (result uint8, ok bool) {
	if value < 0 {
		return uint8(value), false
	}
	if int16(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// Int16ToUint16 converts the int16 value to uint16 safely.
func Int16ToUint16(value int16) (result uint16, ok bool) {
	if value < 0 {
		return uint16(value), false
	}
	return uint16(value), true
}

// Int16ToUint32 converts the int16 value to uint32 safely.
func Int16ToUint32(value int16) (result uint32, ok bool) {
	if value < 0 {
		return uint32(value), false
	}
	return uint32(value), true
}

// Int16ToUint64 converts the int16 value to uint64 safely.
func Int16ToUint64(value int16) (result uint64, ok bool) {
	if value < 0 {
		return uint64(value), false
	}
	return uint64(value), true
}

// Int32ToInt converts the int32 value to int safely.
func Int32ToInt(value int32) (result int, ok bool) {
	return int(value), true
}

// Int32ToInt8 converts the int32 value to int8 safely.
func Int32ToInt8(value int32) (result int8, ok bool) {
	if int32(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// Int32ToInt16 converts the int32 value to int16 safely.
func Int32ToInt16(value int32) (result int16, ok bool) {
	if int32(int16(value)) != value {
		return int16(value), false
	}
	return int16(value), true
}

// Int32ToInt64 converts the int32 value to int64 safely.
func Int32ToInt64(value int32) (result int64, ok bool) {
	return int64(value), true
}

// Int32ToUint converts the int32 value to uint safely.
func Int32ToUint(value int32) (result uint, ok bool) {
	if value < 0 {
		return uint(value), false
	}
	return uint(value), true
}

// Int32ToUint8 converts the int32 value to uint8 safely.
func Int32ToUint8(value int32) (result uint8, ok bool) {
	if value < 0 {
		return uint8(value), false
	}
	if int32(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// Int32ToUint16 converts the int32 value to uint16 safely.
func Int32ToUint16(value int32) (result uint16, ok bool) {
	if value < 0 {
		return uint16(value), false
	}
	if int32(uint16(value)) != value {
		return uint16(value), false
	}
	return uint16(value), true
}

// Int32ToUint32 converts the int32 value to uint32 safely.
func Int32ToUint32(value int32) (result uint32, ok bool) {
	if value < 0 {
		return uint32(value), false
	}
	return uint32(value), true
}

// Int32ToUint64 converts the int32 value to uint64 safely.
func Int32ToUint64(value int32) (result uint64, ok bool) {
	if value < 0 {
		return uint64(value), false
	}
	return uint64(value), true
}

// Int64ToInt converts the int64 value to int safely.
func Int64ToInt(value int64) (result int, ok bool) {
	if math.MaxInt == math.MaxInt32 && int64(int(value)) != value {
		return int(value), false
	}
	return int(value), true
}

// Int64ToInt8 converts the int64 value to int8 safely.
func Int64ToInt8(value int64) (result int8, ok bool) {
	if int64(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// Int64ToInt16 converts the int64 value to int16 safely.
func Int64ToInt16(value int64) (result int16, ok bool) {
	if int64(int16(value)) != value {
		return int16(value), false
	}
	return int16(value), true
}

// Int64ToInt32 converts the int64 value to int32 safely.
func Int64ToInt32(value int64) (result int32, ok bool) {
	if int64(int32(value)) != value {
		return int32(value), false
	}
	return int32(value), true
}

// Int64ToUint converts the int64 value to uint safely.
func Int64ToUint(value int64) (result uint, ok bool) {
	if value < 0 {
		return uint(value), false
	}
	if math.MaxInt == math.MaxInt32 && value > math.MaxUint32 {
		return uint(value), false
	}
	return uint(value), true
}

// Int64ToUint8 converts the int64 value to uint8 safely.
func Int64ToUint8(value int64) (result uint8, ok bool) {
	if value < 0 {
		return uint8(value), false
	}
	if int64(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// Int64ToUint16 converts the int64 value to uint16 safely.
func Int64ToUint16(value int64) (result uint16, ok bool) {
	if value < 0 {
		return uint16(value), false
	}
	if int64(uint16(value)) != value {
		return uint16(value), false
	}
	return uint16(value), true
}

// Int64ToUint32 converts the int64 value to uint32 safely.
func Int64ToUint32(value int64) (result uint32, ok bool) {
	if value < 0 {
		return uint32(value), false
	}
	if int64(uint32(value)) != value {
		return uint32(value), false
	}
	return uint32(value), true
}

// Int64ToUint64 converts the int64 value to uint64 safely.
func Int64ToUint64(value int64) (result uint64, ok bool) {
	if value < 0 {
		return uint64(value), false
	}
	return uint64(value), true
}

// UintToInt converts the uint value to int safely.
func UintToInt(value uint) (result int, ok bool) {
	if value > math.MaxInt {
		return int(value), false
	}
	return int(value), true
}

// UintToInt8 converts the uint value to int8 safely.
func UintToInt8(value uint) (result int8, ok bool) {
	if value > math.MaxInt {
		return int8(value), false
	}
	return int8(value), true
}

// UintToInt16 converts the uint value to int16 safely.
func UintToInt16(value uint) (result int16, ok bool) {
	if value > math.MaxInt {
		return int16(value), false
	}
	return int16(value), true
}

// UintToInt32 converts the uint value to int32 safely.
func UintToInt32(value uint) (result int32, ok bool) {
	if value > math.MaxInt {
		return int32(value), false
	}
	return int32(value), true
}

// UintToInt64 converts the uint value to int64 safely.
func UintToInt64(value uint) (result int64, ok bool) {
	return int64(value), true
}

// UintToUint8 converts the uint value to uint8 safely.
func UintToUint8(value uint) (result uint8, ok bool) {
	if uint(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// UintToUint16 converts the uint value to uint16 safely.
func UintToUint16(value uint) (result uint16, ok bool) {
	if uint(uint16(value)) != value {
		return uint16(value), false
	}
	return uint16(value), true
}

// UintToUint32 converts the uint value to uint32 safely.
func UintToUint32(value uint) (result uint32, ok bool) {
	if math.MaxInt == math.MaxInt64 && uint(uint32(value)) != value {
		return uint32(value), false
	}
	return uint32(value), true
}

// UintToUint64 converts the uint value to uint64 safely.
func UintToUint64(value uint) (result uint64, ok bool) {
	return uint64(value), true
}

// Uint8ToInt converts the uint8 value to int safely.
func Uint8ToInt(value uint8) (result int, ok bool) {
	return int(value), true
}

// Uint8ToInt8 converts the uint8 value to int8 safely.
func Uint8ToInt8(value uint8) (result int8, ok bool) {
	if value > math.MaxInt8 {
		return int8(value), false
	}
	return int8(value), true
}

// Uint8ToInt16 converts the uint8 value to int16 safely.
func Uint8ToInt16(value uint8) (result int16, ok bool) {
	return int16(value), true
}

// Uint8ToInt32 converts the uint8 value to int32 safely.
func Uint8ToInt32(value uint8) (result int32, ok bool) {
	return int32(value), true
}

// Uint8ToInt64 converts the uint8 value to int64 safely.
func Uint8ToInt64(value uint8) (result int64, ok bool) {
	return int64(value), true
}

// Uint8ToUint converts the uint8 value to uint safely.
func Uint8ToUint(value uint8) (result uint, ok bool) {
	return uint(value), true
}

// Uint8ToUint16 converts the uint8 value to uint16 safely.
func Uint8ToUint16(value uint8) (result uint16, ok bool) {
	return uint16(value), true
}

// Uint8ToUint32 converts the uint8 value to uint32 safely.
func Uint8ToUint32(value uint8) (result uint32, ok bool) {
	return uint32(value), true
}

// Uint8ToUint64 converts the uint8 value to uint64 safely.
func Uint8ToUint64(value uint8) (result uint64, ok bool) {
	return uint64(value), true
}

// Uint16ToInt converts the uint16 value to int safely.
func Uint16ToInt(value uint16) (result int, ok bool) {
	return int(value), true
}

// Uint16ToInt8 converts the uint16 value to int8 safely.
func Uint16ToInt8(value uint16) (result int8, ok bool) {
	if uint16(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// Uint16ToInt16 converts the uint16 value to int16 safely.
func Uint16ToInt16(value uint16) (result int16, ok bool) {
	if value > math.MaxInt16 {
		return int16(value), false
	}
	return int16(value), true
}

// Uint16ToInt32 converts the uint16 value to int32 safely.
func Uint16ToInt32(value uint16) (result int32, ok bool) {
	return int32(value), true
}

// Uint16ToInt64 converts the uint16 value to int64 safely.
func Uint16ToInt64(value uint16) (result int64, ok bool) {
	return int64(value), true
}

// Uint16ToUint converts the uint16 value to uint safely.
func Uint16ToUint(value uint16) (result uint, ok bool) {
	return uint(value), true
}

// Uint16ToUint8 converts the uint16 value to uint8 safely.
func Uint16ToUint8(value uint16) (result uint8, ok bool) {
	if uint16(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// Uint16ToUint32 converts the uint16 value to uint32 safely.
func Uint16ToUint32(value uint16) (result uint32, ok bool) {
	return uint32(value), true
}

// Uint16ToUint64 converts the uint16 value to uint64 safely.
func Uint16ToUint64(value uint16) (result uint64, ok bool) {
	return uint64(value), true
}

// Uint32ToInt converts the uint32 value to int safely.
func Uint32ToInt(value uint32) (result int, ok bool) {
	if value > math.MaxInt32 {
		return int(value), false
	}
	return int(value), true
}

// Uint32ToInt8 converts the uint32 value to int8 safely.
func Uint32ToInt8(value uint32) (result int8, ok bool) {
	if uint32(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// Uint32ToInt16 converts the uint32 value to int16 safely.
func Uint32ToInt16(value uint32) (result int16, ok bool) {
	if uint32(int16(value)) != value {
		return int16(value), false
	}
	return int16(value), true
}

// Uint32ToInt32 converts the uint32 value to int32 safely.
func Uint32ToInt32(value uint32) (result int32, ok bool) {
	if value > math.MaxInt32 {
		return int32(value), false
	}
	return int32(value), true
}

// Uint32ToInt64 converts the uint32 value to int64 safely.
func Uint32ToInt64(value uint32) (result int64, ok bool) {
	return int64(value), true
}

// Uint32ToUint converts the uint32 value to uint safely.
func Uint32ToUint(value uint32) (result uint, ok bool) {
	return uint(value), true
}

// Uint32ToUint8 converts the uint32 value to uint8 safely.
func Uint32ToUint8(value uint32) (result uint8, ok bool) {
	if uint32(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// Uint32ToUint16 converts the uint32 value to uint16 safely.
func Uint32ToUint16(value uint32) (result uint16, ok bool) {
	if uint32(uint16(value)) != value {
		return uint16(value), false
	}
	return uint16(value), true
}

// Uint32ToUint64 converts the uint32 value to uint64 safely.
func Uint32ToUint64(value uint32) (result uint64, ok bool) {
	return uint64(value), true
}

// Uint64ToInt converts the uint64 value to int safely.
func Uint64ToInt(value uint64) (result int, ok bool) {
	if value > math.MaxInt64 {
		return int(value), false
	}
	return int(value), true
}

// Uint64ToInt8 converts the uint64 value to int8 safely.
func Uint64ToInt8(value uint64) (result int8, ok bool) {
	if uint64(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// Uint64ToInt16 converts the uint64 value to int16 safely.
func Uint64ToInt16(value uint64) (result int16, ok bool) {
	if uint64(int16(value)) != value {
		return int16(value), false
	}
	return int16(value), true
}

// Uint64ToInt32 converts the uint64 value to int32 safely.
func Uint64ToInt32(value uint64) (result int32, ok bool) {
	if uint64(int32(value)) != value {
		return int32(value), false
	}
	return int32(value), true
}

// Uint64ToInt64 converts the uint64 value to int64 safely.
func Uint64ToInt64(value uint64) (result int64, ok bool) {
	if value > math.MaxInt64 {
		return int64(value), false
	}
	return int64(value), true
}

// Uint64ToUint converts the uint64 value to uint safely.
func Uint64ToUint(value uint64) (result uint, ok bool) {
	if math.MaxInt == math.MaxInt32 && uint64(uint(value)) != value {
		return uint(value), false
	}
	return uint(value), true
}

// Uint64ToUint8 converts the uint64 value to uint8 safely.
func Uint64ToUint8(value uint64) (result uint8, ok bool) {
	if uint64(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// Uint64ToUint16 converts the uint64 value to uint16 safely.
func Uint64ToUint16(value uint64) (result uint16, ok bool) {
	if uint64(uint16(value)) != value {
		return uint16(value), false
	}
	return uint16(value), true
}

// Uint64ToUint32 converts the uint64 value to uint32 safely.
func Uint64ToUint32(value uint64) (result uint32, ok bool) {
	if uint64(uint32(value)) != value {
		return uint32(value), false
	}
	return uint32(value), true
}

// Float32ToInt converts the float32 value to int safely.
func Float32ToInt(value float32) (result int, ok bool) {
	if value < float32(math.MinInt) || value > float32(math.MaxInt) {
		return int(value), false
	}
	return int(value), true
}

func IntToFloat32(value int) (float32, bool) {
	return float32(value), true
}
// Float32ToInt8 converts the float32 value to int8 safely.
func Float32ToInt8(value float32) (result int8, ok bool) {
	if value < float32(math.MinInt8) || value > float32(math.MaxInt8) {
		return int8(value), false
	}
	return int8(value), true
}

func Int8ToFloat32(value int8) (float32, bool) {
	return float32(value), true
}
// Float32ToInt16 converts the float32 value to int16 safely.
func Float32ToInt16(value float32) (result int16, ok bool) {
	if value < float32(math.MinInt16) || value > float32(math.MaxInt16) {
		return int16(value), false
	}
	return int16(value), true
}

func Int16ToFloat32(value int16) (float32, bool) {
	return float32(value), true
}
// Float32ToInt32 converts the float32 value to int32 safely.
func Float32ToInt32(value float32) (result int32, ok bool) {
	if value < float32(math.MinInt32) || value > float32(math.MaxInt32) {
		return int32(value), false
	}
	return int32(value), true
}

func Int32ToFloat32(value int32) (float32, bool) {
	return float32(value), true
}
// Float32ToInt64 converts the float32 value to int64 safely.
func Float32ToInt64(value float32) (result int64, ok bool) {
	if value < float32(math.MinInt64) || value > float32(math.MaxInt64) {
		return int64(value), false
	}
	return int64(value), true
}

func Int64ToFloat32(value int64) (float32, bool) {
	return float32(value), true
}
// Float32ToUint converts the float32 value to uint safely.
func Float32ToUint(value float32) (result uint, ok bool) {
	if value < 0 || value > float32(math.MaxUint) {
		return uint(value), false
	}
	return uint(value), true
}

func UintToFloat32(value uint) (float32, bool) {
	return float32(value), true
}
// Float32ToUint8 converts the float32 value to uint8 safely.
func Float32ToUint8(value float32) (result uint8, ok bool) {
	if value < 0 || value > float32(math.MaxUint8) {
		return uint8(value), false
	}
	return uint8(value), true
}

func Uint8ToFloat32(value uint8) (float32, bool) {
	return float32(value), true
}
// Float32ToUint16 converts the float32 value to uint16 safely.
func Float32ToUint16(value float32) (result uint16, ok bool) {
	if value < 0 || value > float32(math.MaxUint16) {
		return uint16(value), false
	}
	return uint16(value), true
}

func Uint16ToFloat32(value uint16) (float32, bool) {
	return float32(value), true
}
// Float32ToUint32 converts the float32 value to uint32 safely.
func Float32ToUint32(value float32) (result uint32, ok bool) {
	if value < 0 || value > float32(math.MaxUint32) {
		return uint32(value), false
	}
	return uint32(value), true
}

func Uint32ToFloat32(value uint32) (float32, bool) {
	return float32(value), true
}
// Float32ToUint64 converts the float32 value to uint64 safely.
func Float32ToUint64(value float32) (result uint64, ok bool) {
	if value < 0 || value > float32(math.MaxUint64) {
		return uint64(value), false
	}
	return uint64(value), true
}

func Uint64ToFloat32(value uint64) (float32, bool) {
	return float32(value), true
}
// Float64ToInt converts the float64 value to int safely.
func Float64ToInt(value float64) (result int, ok bool) {
	if value < float64(math.MinInt) || value > float64(math.MaxInt) {
		return int(value), false
	}
	return int(value), true
}

func IntToFloat64(value int) (float64, bool) {
	return float64(value), true
}
// Float64ToInt8 converts the float64 value to int8 safely.
func Float64ToInt8(value float64) (result int8, ok bool) {
	if value < float64(math.MinInt8) || value > float64(math.MaxInt8) {
		return int8(value), false
	}
	return int8(value), true
}

func Int8ToFloat64(value int8) (float64, bool) {
	return float64(value), true
}
// Float64ToInt16 converts the float64 value to int16 safely.
func Float64ToInt16(value float64) (result int16, ok bool) {
	if value < float64(math.MinInt16) || value > float64(math.MaxInt16) {
		return int16(value), false
	}
	return int16(value), true
}

func Int16ToFloat64(value int16) (float64, bool) {
	return float64(value), true
}
// Float64ToInt32 converts the float64 value to int32 safely.
func Float64ToInt32(value float64) (result int32, ok bool) {
	if value < float64(math.MinInt32) || value > float64(math.MaxInt32) {
		return int32(value), false
	}
	return int32(value), true
}

func Int32ToFloat64(value int32) (float64, bool) {
	return float64(value), true
}
// Float64ToInt64 converts the float64 value to int64 safely.
func Float64ToInt64(value float64) (result int64, ok bool) {
	if value < float64(math.MinInt64) || value > float64(math.MaxInt64) {
		return int64(value), false
	}
	return int64(value), true
}

func Int64ToFloat64(value int64) (float64, bool) {
	return float64(value), true
}
// Float64ToUint converts the float64 value to uint safely.
func Float64ToUint(value float64) (result uint, ok bool) {
	if value < 0 || value > float64(math.MaxUint) {
		return uint(value), false
	}
	return uint(value), true
}

func UintToFloat64(value uint) (float64, bool) {
	return float64(value), true
}
// Float64ToUint8 converts the float64 value to uint8 safely.
func Float64ToUint8(value float64) (result uint8, ok bool) {
	if value < 0 || value > float64(math.MaxUint8) {
		return uint8(value), false
	}
	return uint8(value), true
}

func Uint8ToFloat64(value uint8) (float64, bool) {
	return float64(value), true
}
// Float64ToUint16 converts the float64 value to uint16 safely.
func Float64ToUint16(value float64) (result uint16, ok bool) {
	if value < 0 || value > float64(math.MaxUint16) {
		return uint16(value), false
	}
	return uint16(value), true
}

func Uint16ToFloat64(value uint16) (float64, bool) {
	return float64(value), true
}
// Float64ToUint32 converts the float64 value to uint32 safely.
func Float64ToUint32(value float64) (result uint32, ok bool) {
	if value < 0 || value > float64(math.MaxUint32) {
		return uint32(value), false
	}
	return uint32(value), true
}

func Uint32ToFloat64(value uint32) (float64, bool) {
	return float64(value), true
}
// Float64ToUint64 converts the float64 value to uint64 safely.
func Float64ToUint64(value float64) (result uint64, ok bool) {
	if value < 0 || value > float64(math.MaxUint64) {
		return uint64(value), false
	}
	return uint64(value), true
}

func Uint64ToFloat64(value uint64) (float64, bool) {
	return float64(value), true
}

    func Float32ToFloat64(value float32) (float64, bool) {
        return float64(value), true
    }

    func Float64ToFloat32(value float64) (float32, bool) {
        if value > math.MaxFloat32 || value < -math.MaxFloat32 {
            return float32(value), false
        }
        return float32(value), true
    }
    
