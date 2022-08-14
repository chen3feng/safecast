package safecast

import "math"

// intToInt8 converts the int value to int8 safely.
func intToInt8(value int) (result int8, ok bool) {
	if int(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// intToInt16 converts the int value to int16 safely.
func intToInt16(value int) (result int16, ok bool) {
	if int(int16(value)) != value {
		return int16(value), false
	}
	return int16(value), true
}

// intToInt32 converts the int value to int32 safely.
func intToInt32(value int) (result int32, ok bool) {
	if math.MaxInt == math.MaxInt64 && int(int32(value)) != value {
		return int32(value), false
	}
	return int32(value), true
}

// intToInt64 converts the int value to int64 safely.
func intToInt64(value int) (result int64, ok bool) {
	return int64(value), true
}

// intToUint converts the int value to uint safely.
func intToUint(value int) (result uint, ok bool) {
	if value < 0 {
		return uint(value), false
	}
	return uint(value), true
}

// intToUint8 converts the int value to uint8 safely.
func intToUint8(value int) (result uint8, ok bool) {
	if value < 0 {
		return uint8(value), false
	}
	if value > math.MaxUint8 {
		return uint8(value), false
	}
	return uint8(value), true
}

// intToUint16 converts the int value to uint16 safely.
func intToUint16(value int) (result uint16, ok bool) {
	if value < 0 {
		return uint16(value), false
	}
	if value > math.MaxUint16 {
		return uint16(value), false
	}
	return uint16(value), true
}

// intToUint32 converts the int value to uint32 safely.
func intToUint32(value int) (result uint32, ok bool) {
	if value < 0 {
		return uint32(value), false
	}
	if value > math.MaxUint32 {
		return uint32(value), false
	}
	return uint32(value), true
}

// intToUint64 converts the int value to uint64 safely.
func intToUint64(value int) (result uint64, ok bool) {
	if value < 0 {
		return uint64(value), false
	}
	return uint64(value), true
}

// int8ToInt converts the int8 value to int safely.
func int8ToInt(value int8) (result int, ok bool) {
	return int(value), true
}

// int8ToInt16 converts the int8 value to int16 safely.
func int8ToInt16(value int8) (result int16, ok bool) {
	return int16(value), true
}

// int8ToInt32 converts the int8 value to int32 safely.
func int8ToInt32(value int8) (result int32, ok bool) {
	return int32(value), true
}

// int8ToInt64 converts the int8 value to int64 safely.
func int8ToInt64(value int8) (result int64, ok bool) {
	return int64(value), true
}

// int8ToUint converts the int8 value to uint safely.
func int8ToUint(value int8) (result uint, ok bool) {
	if value < 0 {
		return uint(value), false
	}
	return uint(value), true
}

// int8ToUint8 converts the int8 value to uint8 safely.
func int8ToUint8(value int8) (result uint8, ok bool) {
	if value < 0 {
		return uint8(value), false
	}
	return uint8(value), true
}

// int8ToUint16 converts the int8 value to uint16 safely.
func int8ToUint16(value int8) (result uint16, ok bool) {
	if value < 0 {
		return uint16(value), false
	}
	return uint16(value), true
}

// int8ToUint32 converts the int8 value to uint32 safely.
func int8ToUint32(value int8) (result uint32, ok bool) {
	if value < 0 {
		return uint32(value), false
	}
	return uint32(value), true
}

// int8ToUint64 converts the int8 value to uint64 safely.
func int8ToUint64(value int8) (result uint64, ok bool) {
	if value < 0 {
		return uint64(value), false
	}
	return uint64(value), true
}

// int16ToInt converts the int16 value to int safely.
func int16ToInt(value int16) (result int, ok bool) {
	return int(value), true
}

// int16ToInt8 converts the int16 value to int8 safely.
func int16ToInt8(value int16) (result int8, ok bool) {
	if int16(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// int16ToInt32 converts the int16 value to int32 safely.
func int16ToInt32(value int16) (result int32, ok bool) {
	return int32(value), true
}

// int16ToInt64 converts the int16 value to int64 safely.
func int16ToInt64(value int16) (result int64, ok bool) {
	return int64(value), true
}

// int16ToUint converts the int16 value to uint safely.
func int16ToUint(value int16) (result uint, ok bool) {
	if value < 0 {
		return uint(value), false
	}
	return uint(value), true
}

// int16ToUint8 converts the int16 value to uint8 safely.
func int16ToUint8(value int16) (result uint8, ok bool) {
	if value < 0 {
		return uint8(value), false
	}
	if int16(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// int16ToUint16 converts the int16 value to uint16 safely.
func int16ToUint16(value int16) (result uint16, ok bool) {
	if value < 0 {
		return uint16(value), false
	}
	return uint16(value), true
}

// int16ToUint32 converts the int16 value to uint32 safely.
func int16ToUint32(value int16) (result uint32, ok bool) {
	if value < 0 {
		return uint32(value), false
	}
	return uint32(value), true
}

// int16ToUint64 converts the int16 value to uint64 safely.
func int16ToUint64(value int16) (result uint64, ok bool) {
	if value < 0 {
		return uint64(value), false
	}
	return uint64(value), true
}

// int32ToInt converts the int32 value to int safely.
func int32ToInt(value int32) (result int, ok bool) {
	return int(value), true
}

// int32ToInt8 converts the int32 value to int8 safely.
func int32ToInt8(value int32) (result int8, ok bool) {
	if int32(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// int32ToInt16 converts the int32 value to int16 safely.
func int32ToInt16(value int32) (result int16, ok bool) {
	if int32(int16(value)) != value {
		return int16(value), false
	}
	return int16(value), true
}

// int32ToInt64 converts the int32 value to int64 safely.
func int32ToInt64(value int32) (result int64, ok bool) {
	return int64(value), true
}

// int32ToUint converts the int32 value to uint safely.
func int32ToUint(value int32) (result uint, ok bool) {
	if value < 0 {
		return uint(value), false
	}
	return uint(value), true
}

// int32ToUint8 converts the int32 value to uint8 safely.
func int32ToUint8(value int32) (result uint8, ok bool) {
	if value < 0 {
		return uint8(value), false
	}
	if int32(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// int32ToUint16 converts the int32 value to uint16 safely.
func int32ToUint16(value int32) (result uint16, ok bool) {
	if value < 0 {
		return uint16(value), false
	}
	if int32(uint16(value)) != value {
		return uint16(value), false
	}
	return uint16(value), true
}

// int32ToUint32 converts the int32 value to uint32 safely.
func int32ToUint32(value int32) (result uint32, ok bool) {
	if value < 0 {
		return uint32(value), false
	}
	return uint32(value), true
}

// int32ToUint64 converts the int32 value to uint64 safely.
func int32ToUint64(value int32) (result uint64, ok bool) {
	if value < 0 {
		return uint64(value), false
	}
	return uint64(value), true
}

// int64ToInt converts the int64 value to int safely.
func int64ToInt(value int64) (result int, ok bool) {
	if math.MaxInt == math.MaxInt32 && int64(int(value)) != value {
		return int(value), false
	}
	return int(value), true
}

// int64ToInt8 converts the int64 value to int8 safely.
func int64ToInt8(value int64) (result int8, ok bool) {
	if int64(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// int64ToInt16 converts the int64 value to int16 safely.
func int64ToInt16(value int64) (result int16, ok bool) {
	if int64(int16(value)) != value {
		return int16(value), false
	}
	return int16(value), true
}

// int64ToInt32 converts the int64 value to int32 safely.
func int64ToInt32(value int64) (result int32, ok bool) {
	if int64(int32(value)) != value {
		return int32(value), false
	}
	return int32(value), true
}

// int64ToUint converts the int64 value to uint safely.
func int64ToUint(value int64) (result uint, ok bool) {
	if value < 0 {
		return uint(value), false
	}
	if math.MaxInt == math.MaxInt32 && value > math.MaxUint32 {
		return uint(value), false
	}
	return uint(value), true
}

// int64ToUint8 converts the int64 value to uint8 safely.
func int64ToUint8(value int64) (result uint8, ok bool) {
	if value < 0 {
		return uint8(value), false
	}
	if int64(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// int64ToUint16 converts the int64 value to uint16 safely.
func int64ToUint16(value int64) (result uint16, ok bool) {
	if value < 0 {
		return uint16(value), false
	}
	if int64(uint16(value)) != value {
		return uint16(value), false
	}
	return uint16(value), true
}

// int64ToUint32 converts the int64 value to uint32 safely.
func int64ToUint32(value int64) (result uint32, ok bool) {
	if value < 0 {
		return uint32(value), false
	}
	if int64(uint32(value)) != value {
		return uint32(value), false
	}
	return uint32(value), true
}

// int64ToUint64 converts the int64 value to uint64 safely.
func int64ToUint64(value int64) (result uint64, ok bool) {
	if value < 0 {
		return uint64(value), false
	}
	return uint64(value), true
}

// uintToInt converts the uint value to int safely.
func uintToInt(value uint) (result int, ok bool) {
	if value > math.MaxInt {
		return int(value), false
	}
	return int(value), true
}

// uintToInt8 converts the uint value to int8 safely.
func uintToInt8(value uint) (result int8, ok bool) {
	if value > math.MaxInt {
		return int8(value), false
	}
	return int8(value), true
}

// uintToInt16 converts the uint value to int16 safely.
func uintToInt16(value uint) (result int16, ok bool) {
	if value > math.MaxInt {
		return int16(value), false
	}
	return int16(value), true
}

// uintToInt32 converts the uint value to int32 safely.
func uintToInt32(value uint) (result int32, ok bool) {
	if value > math.MaxInt {
		return int32(value), false
	}
	return int32(value), true
}

// uintToInt64 converts the uint value to int64 safely.
func uintToInt64(value uint) (result int64, ok bool) {
	return int64(value), true
}

// uintToUint8 converts the uint value to uint8 safely.
func uintToUint8(value uint) (result uint8, ok bool) {
	if uint(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// uintToUint16 converts the uint value to uint16 safely.
func uintToUint16(value uint) (result uint16, ok bool) {
	if uint(uint16(value)) != value {
		return uint16(value), false
	}
	return uint16(value), true
}

// uintToUint32 converts the uint value to uint32 safely.
func uintToUint32(value uint) (result uint32, ok bool) {
	if math.MaxInt == math.MaxInt64 && uint(uint32(value)) != value {
		return uint32(value), false
	}
	return uint32(value), true
}

// uintToUint64 converts the uint value to uint64 safely.
func uintToUint64(value uint) (result uint64, ok bool) {
	return uint64(value), true
}

// uint8ToInt converts the uint8 value to int safely.
func uint8ToInt(value uint8) (result int, ok bool) {
	return int(value), true
}

// uint8ToInt8 converts the uint8 value to int8 safely.
func uint8ToInt8(value uint8) (result int8, ok bool) {
	if value > math.MaxInt8 {
		return int8(value), false
	}
	return int8(value), true
}

// uint8ToInt16 converts the uint8 value to int16 safely.
func uint8ToInt16(value uint8) (result int16, ok bool) {
	return int16(value), true
}

// uint8ToInt32 converts the uint8 value to int32 safely.
func uint8ToInt32(value uint8) (result int32, ok bool) {
	return int32(value), true
}

// uint8ToInt64 converts the uint8 value to int64 safely.
func uint8ToInt64(value uint8) (result int64, ok bool) {
	return int64(value), true
}

// uint8ToUint converts the uint8 value to uint safely.
func uint8ToUint(value uint8) (result uint, ok bool) {
	return uint(value), true
}

// uint8ToUint16 converts the uint8 value to uint16 safely.
func uint8ToUint16(value uint8) (result uint16, ok bool) {
	return uint16(value), true
}

// uint8ToUint32 converts the uint8 value to uint32 safely.
func uint8ToUint32(value uint8) (result uint32, ok bool) {
	return uint32(value), true
}

// uint8ToUint64 converts the uint8 value to uint64 safely.
func uint8ToUint64(value uint8) (result uint64, ok bool) {
	return uint64(value), true
}

// uint16ToInt converts the uint16 value to int safely.
func uint16ToInt(value uint16) (result int, ok bool) {
	return int(value), true
}

// uint16ToInt8 converts the uint16 value to int8 safely.
func uint16ToInt8(value uint16) (result int8, ok bool) {
	if uint16(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// uint16ToInt16 converts the uint16 value to int16 safely.
func uint16ToInt16(value uint16) (result int16, ok bool) {
	if value > math.MaxInt16 {
		return int16(value), false
	}
	return int16(value), true
}

// uint16ToInt32 converts the uint16 value to int32 safely.
func uint16ToInt32(value uint16) (result int32, ok bool) {
	return int32(value), true
}

// uint16ToInt64 converts the uint16 value to int64 safely.
func uint16ToInt64(value uint16) (result int64, ok bool) {
	return int64(value), true
}

// uint16ToUint converts the uint16 value to uint safely.
func uint16ToUint(value uint16) (result uint, ok bool) {
	return uint(value), true
}

// uint16ToUint8 converts the uint16 value to uint8 safely.
func uint16ToUint8(value uint16) (result uint8, ok bool) {
	if uint16(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// uint16ToUint32 converts the uint16 value to uint32 safely.
func uint16ToUint32(value uint16) (result uint32, ok bool) {
	return uint32(value), true
}

// uint16ToUint64 converts the uint16 value to uint64 safely.
func uint16ToUint64(value uint16) (result uint64, ok bool) {
	return uint64(value), true
}

// uint32ToInt converts the uint32 value to int safely.
func uint32ToInt(value uint32) (result int, ok bool) {
	if value > math.MaxInt32 {
		return int(value), false
	}
	return int(value), true
}

// uint32ToInt8 converts the uint32 value to int8 safely.
func uint32ToInt8(value uint32) (result int8, ok bool) {
	if uint32(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// uint32ToInt16 converts the uint32 value to int16 safely.
func uint32ToInt16(value uint32) (result int16, ok bool) {
	if uint32(int16(value)) != value {
		return int16(value), false
	}
	return int16(value), true
}

// uint32ToInt32 converts the uint32 value to int32 safely.
func uint32ToInt32(value uint32) (result int32, ok bool) {
	if value > math.MaxInt32 {
		return int32(value), false
	}
	return int32(value), true
}

// uint32ToInt64 converts the uint32 value to int64 safely.
func uint32ToInt64(value uint32) (result int64, ok bool) {
	return int64(value), true
}

// uint32ToUint converts the uint32 value to uint safely.
func uint32ToUint(value uint32) (result uint, ok bool) {
	return uint(value), true
}

// uint32ToUint8 converts the uint32 value to uint8 safely.
func uint32ToUint8(value uint32) (result uint8, ok bool) {
	if uint32(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// uint32ToUint16 converts the uint32 value to uint16 safely.
func uint32ToUint16(value uint32) (result uint16, ok bool) {
	if uint32(uint16(value)) != value {
		return uint16(value), false
	}
	return uint16(value), true
}

// uint32ToUint64 converts the uint32 value to uint64 safely.
func uint32ToUint64(value uint32) (result uint64, ok bool) {
	return uint64(value), true
}

// uint64ToInt converts the uint64 value to int safely.
func uint64ToInt(value uint64) (result int, ok bool) {
	if value > math.MaxInt64 {
		return int(value), false
	}
	return int(value), true
}

// uint64ToInt8 converts the uint64 value to int8 safely.
func uint64ToInt8(value uint64) (result int8, ok bool) {
	if uint64(int8(value)) != value {
		return int8(value), false
	}
	return int8(value), true
}

// uint64ToInt16 converts the uint64 value to int16 safely.
func uint64ToInt16(value uint64) (result int16, ok bool) {
	if uint64(int16(value)) != value {
		return int16(value), false
	}
	return int16(value), true
}

// uint64ToInt32 converts the uint64 value to int32 safely.
func uint64ToInt32(value uint64) (result int32, ok bool) {
	if uint64(int32(value)) != value {
		return int32(value), false
	}
	return int32(value), true
}

// uint64ToInt64 converts the uint64 value to int64 safely.
func uint64ToInt64(value uint64) (result int64, ok bool) {
	if value > math.MaxInt64 {
		return int64(value), false
	}
	return int64(value), true
}

// uint64ToUint converts the uint64 value to uint safely.
func uint64ToUint(value uint64) (result uint, ok bool) {
	if math.MaxInt == math.MaxInt32 && uint64(uint(value)) != value {
		return uint(value), false
	}
	return uint(value), true
}

// uint64ToUint8 converts the uint64 value to uint8 safely.
func uint64ToUint8(value uint64) (result uint8, ok bool) {
	if uint64(uint8(value)) != value {
		return uint8(value), false
	}
	return uint8(value), true
}

// uint64ToUint16 converts the uint64 value to uint16 safely.
func uint64ToUint16(value uint64) (result uint16, ok bool) {
	if uint64(uint16(value)) != value {
		return uint16(value), false
	}
	return uint16(value), true
}

// uint64ToUint32 converts the uint64 value to uint32 safely.
func uint64ToUint32(value uint64) (result uint32, ok bool) {
	if uint64(uint32(value)) != value {
		return uint32(value), false
	}
	return uint32(value), true
}

// float32ToInt converts the float32 value to int safely.
func float32ToInt(value float32) (result int, ok bool) {
	if value < float32(math.MinInt) || value > float32(math.MaxInt) {
		return int(value), false
	}
	return int(value), true
}

func intToFloat32(value int) (float32, bool) {
	return float32(value), true
}
// float32ToInt8 converts the float32 value to int8 safely.
func float32ToInt8(value float32) (result int8, ok bool) {
	if value < float32(math.MinInt8) || value > float32(math.MaxInt8) {
		return int8(value), false
	}
	return int8(value), true
}

func int8ToFloat32(value int8) (float32, bool) {
	return float32(value), true
}
// float32ToInt16 converts the float32 value to int16 safely.
func float32ToInt16(value float32) (result int16, ok bool) {
	if value < float32(math.MinInt16) || value > float32(math.MaxInt16) {
		return int16(value), false
	}
	return int16(value), true
}

func int16ToFloat32(value int16) (float32, bool) {
	return float32(value), true
}
// float32ToInt32 converts the float32 value to int32 safely.
func float32ToInt32(value float32) (result int32, ok bool) {
	if value < float32(math.MinInt32) || value > float32(math.MaxInt32) {
		return int32(value), false
	}
	return int32(value), true
}

func int32ToFloat32(value int32) (float32, bool) {
	return float32(value), true
}
// float32ToInt64 converts the float32 value to int64 safely.
func float32ToInt64(value float32) (result int64, ok bool) {
	if value < float32(math.MinInt64) || value > float32(math.MaxInt64) {
		return int64(value), false
	}
	return int64(value), true
}

func int64ToFloat32(value int64) (float32, bool) {
	return float32(value), true
}
// float32ToUint converts the float32 value to uint safely.
func float32ToUint(value float32) (result uint, ok bool) {
	if value < 0 || value > float32(math.MaxUint) {
		return uint(value), false
	}
	return uint(value), true
}

func uintToFloat32(value uint) (float32, bool) {
	return float32(value), true
}
// float32ToUint8 converts the float32 value to uint8 safely.
func float32ToUint8(value float32) (result uint8, ok bool) {
	if value < 0 || value > float32(math.MaxUint8) {
		return uint8(value), false
	}
	return uint8(value), true
}

func uint8ToFloat32(value uint8) (float32, bool) {
	return float32(value), true
}
// float32ToUint16 converts the float32 value to uint16 safely.
func float32ToUint16(value float32) (result uint16, ok bool) {
	if value < 0 || value > float32(math.MaxUint16) {
		return uint16(value), false
	}
	return uint16(value), true
}

func uint16ToFloat32(value uint16) (float32, bool) {
	return float32(value), true
}
// float32ToUint32 converts the float32 value to uint32 safely.
func float32ToUint32(value float32) (result uint32, ok bool) {
	if value < 0 || value > float32(math.MaxUint32) {
		return uint32(value), false
	}
	return uint32(value), true
}

func uint32ToFloat32(value uint32) (float32, bool) {
	return float32(value), true
}
// float32ToUint64 converts the float32 value to uint64 safely.
func float32ToUint64(value float32) (result uint64, ok bool) {
	if value < 0 || value > float32(math.MaxUint64) {
		return uint64(value), false
	}
	return uint64(value), true
}

func uint64ToFloat32(value uint64) (float32, bool) {
	return float32(value), true
}
// float64ToInt converts the float64 value to int safely.
func float64ToInt(value float64) (result int, ok bool) {
	if value < float64(math.MinInt) || value > float64(math.MaxInt) {
		return int(value), false
	}
	return int(value), true
}

func intToFloat64(value int) (float64, bool) {
	return float64(value), true
}
// float64ToInt8 converts the float64 value to int8 safely.
func float64ToInt8(value float64) (result int8, ok bool) {
	if value < float64(math.MinInt8) || value > float64(math.MaxInt8) {
		return int8(value), false
	}
	return int8(value), true
}

func int8ToFloat64(value int8) (float64, bool) {
	return float64(value), true
}
// float64ToInt16 converts the float64 value to int16 safely.
func float64ToInt16(value float64) (result int16, ok bool) {
	if value < float64(math.MinInt16) || value > float64(math.MaxInt16) {
		return int16(value), false
	}
	return int16(value), true
}

func int16ToFloat64(value int16) (float64, bool) {
	return float64(value), true
}
// float64ToInt32 converts the float64 value to int32 safely.
func float64ToInt32(value float64) (result int32, ok bool) {
	if value < float64(math.MinInt32) || value > float64(math.MaxInt32) {
		return int32(value), false
	}
	return int32(value), true
}

func int32ToFloat64(value int32) (float64, bool) {
	return float64(value), true
}
// float64ToInt64 converts the float64 value to int64 safely.
func float64ToInt64(value float64) (result int64, ok bool) {
	if value < float64(math.MinInt64) || value > float64(math.MaxInt64) {
		return int64(value), false
	}
	return int64(value), true
}

func int64ToFloat64(value int64) (float64, bool) {
	return float64(value), true
}
// float64ToUint converts the float64 value to uint safely.
func float64ToUint(value float64) (result uint, ok bool) {
	if value < 0 || value > float64(math.MaxUint) {
		return uint(value), false
	}
	return uint(value), true
}

func uintToFloat64(value uint) (float64, bool) {
	return float64(value), true
}
// float64ToUint8 converts the float64 value to uint8 safely.
func float64ToUint8(value float64) (result uint8, ok bool) {
	if value < 0 || value > float64(math.MaxUint8) {
		return uint8(value), false
	}
	return uint8(value), true
}

func uint8ToFloat64(value uint8) (float64, bool) {
	return float64(value), true
}
// float64ToUint16 converts the float64 value to uint16 safely.
func float64ToUint16(value float64) (result uint16, ok bool) {
	if value < 0 || value > float64(math.MaxUint16) {
		return uint16(value), false
	}
	return uint16(value), true
}

func uint16ToFloat64(value uint16) (float64, bool) {
	return float64(value), true
}
// float64ToUint32 converts the float64 value to uint32 safely.
func float64ToUint32(value float64) (result uint32, ok bool) {
	if value < 0 || value > float64(math.MaxUint32) {
		return uint32(value), false
	}
	return uint32(value), true
}

func uint32ToFloat64(value uint32) (float64, bool) {
	return float64(value), true
}
// float64ToUint64 converts the float64 value to uint64 safely.
func float64ToUint64(value float64) (result uint64, ok bool) {
	if value < 0 || value > float64(math.MaxUint64) {
		return uint64(value), false
	}
	return uint64(value), true
}

func uint64ToFloat64(value uint64) (float64, bool) {
	return float64(value), true
}

    func float32ToFloat64(value float32) (float64, bool) {
        return float64(value), true
    }

    func float64ToFloat32(value float64) (float32, bool) {
        if value > math.MaxFloat32 || value < -math.MaxFloat32 {
            return float32(value), false
        }
        return float32(value), true
    }
    
