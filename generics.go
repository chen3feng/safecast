package safecast

type numericType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// To converts a numeric value from the FromType to the specified ToType type safely.
// result will always be same as the usual type cast (type(value)),
// but ok is false when overflow or underflow occured.
func To[ToType numericType, FromType numericType](value FromType) (result ToType, ok bool) {
	ok = true
	switch t := any(result).(type) {
	case int:
		t, ok = toInt(value)
		result = ToType(t)
	case int8:
		t, ok = toInt8(value)
		result = ToType(t)
	case int16:
		t, ok = toInt16(value)
		result = ToType(t)
	case int32:
		t, ok = toInt32(value)
		result = ToType(t)
	case int64:
		t, ok = toInt64(value)
		result = ToType(t)
	case uint:
		t, ok = toUint(value)
		result = ToType(t)
	case uint8:
		t, ok = toUint8(value)
		result = ToType(t)
	case uint16:
		t, ok = toUint16(value)
		result = ToType(t)
	case uint32:
		t, ok = toUint32(value)
		result = ToType(t)
	case uint64:
		t, ok = toUint64(value)
		result = ToType(t)
	case float32:
		t, ok = toFloat32(value)
		result = ToType(t)
	case float64:
		t, ok = toFloat64(value)
		result = ToType(t)
	}
	return result, ok
}

func toInt[F numericType](value F) (int, bool) {
	switch f := any(value).(type) {
	case int:
		return f, true
	case int8:
		return int8ToInt(f)
	case int16:
		return int16ToInt(f)
	case int32:
		return int32ToInt(f)
	case int64:
		return int64ToInt(f)
	case uint:
		return uintToInt(f)
	case uint8:
		return uint8ToInt(f)
	case uint16:
		return uint16ToInt(f)
	case uint32:
		return uint32ToInt(f)
	case uint64:
		return uint64ToInt(f)
	case float32:
		return float32ToInt(f)
	case float64:
		return float64ToInt(f)
	}
	return int(value), false
}

func toInt8[F numericType](value F) (int8, bool) {
	switch f := any(value).(type) {
	case int:
		return intToInt8(f)
	case int8:
		return f, true
	case int16:
		return int16ToInt8(f)
	case int32:
		return int32ToInt8(f)
	case int64:
		return int64ToInt8(f)
	case uint:
		return uintToInt8(f)
	case uint8:
		return uint8ToInt8(f)
	case uint16:
		return uint16ToInt8(f)
	case uint32:
		return uint32ToInt8(f)
	case uint64:
		return uint64ToInt8(f)
	case float32:
		return float32ToInt8(f)
	case float64:
		return float64ToInt8(f)
	}
	return int8(value), false
}

func toInt16[F numericType](value F) (int16, bool) {
	switch f := any(value).(type) {
	case int:
		return intToInt16(f)
	case int8:
		return int8ToInt16(f)
	case int16:
		return f, true
	case int32:
		return int32ToInt16(f)
	case int64:
		return int64ToInt16(f)
	case uint:
		return uintToInt16(f)
	case uint8:
		return uint8ToInt16(f)
	case uint16:
		return uint16ToInt16(f)
	case uint32:
		return uint32ToInt16(f)
	case uint64:
		return uint64ToInt16(f)
	case float32:
		return float32ToInt16(f)
	case float64:
		return float64ToInt16(f)
	}
	return int16(value), false
}

func toInt32[F numericType](value F) (int32, bool) {
	switch f := any(value).(type) {
	case int:
		return intToInt32(f)
	case int8:
		return int8ToInt32(f)
	case int16:
		return int16ToInt32(f)
	case int32:
		return f, true
	case int64:
		return int64ToInt32(f)
	case uint:
		return uintToInt32(f)
	case uint8:
		return uint8ToInt32(f)
	case uint16:
		return uint16ToInt32(f)
	case uint32:
		return uint32ToInt32(f)
	case uint64:
		return uint64ToInt32(f)
	case float32:
		return float32ToInt32(f)
	case float64:
		return float64ToInt32(f)
	}
	return int32(value), false
}

func toInt64[F numericType](value F) (int64, bool) {
	switch f := any(value).(type) {
	case int:
		return intToInt64(f)
	case int8:
		return int8ToInt64(f)
	case int16:
		return int16ToInt64(f)
	case int32:
		return int32ToInt64(f)
	case int64:
		return f, true
	case uint:
		return uintToInt64(f)
	case uint8:
		return uint8ToInt64(f)
	case uint16:
		return uint16ToInt64(f)
	case uint32:
		return uint32ToInt64(f)
	case uint64:
		return uint64ToInt64(f)
	case float32:
		return float32ToInt64(f)
	case float64:
		return float64ToInt64(f)
	}
	return int64(value), false
}

func toUint[F numericType](value F) (uint, bool) {
	switch f := any(value).(type) {
	case int:
		return intToUint(f)
	case int8:
		return int8ToUint(f)
	case int16:
		return int16ToUint(f)
	case int32:
		return int32ToUint(f)
	case int64:
		return int64ToUint(f)
	case uint:
		return f, true
	case uint8:
		return uint8ToUint(f)
	case uint16:
		return uint16ToUint(f)
	case uint32:
		return uint32ToUint(f)
	case uint64:
		return uint64ToUint(f)
	case float32:
		return float32ToUint(f)
	case float64:
		return float64ToUint(f)
	}
	return uint(value), false
}

func toUint8[F numericType](value F) (uint8, bool) {
	switch f := any(value).(type) {
	case int:
		return intToUint8(f)
	case int8:
		return int8ToUint8(f)
	case int16:
		return int16ToUint8(f)
	case int32:
		return int32ToUint8(f)
	case int64:
		return int64ToUint8(f)
	case uint:
		return uintToUint8(f)
	case uint8:
		return f, true
	case uint16:
		return uint16ToUint8(f)
	case uint32:
		return uint32ToUint8(f)
	case uint64:
		return uint64ToUint8(f)
	case float32:
		return float32ToUint8(f)
	case float64:
		return float64ToUint8(f)
	}
	return uint8(value), false
}

func toUint16[F numericType](value F) (uint16, bool) {
	switch f := any(value).(type) {
	case int:
		return intToUint16(f)
	case int8:
		return int8ToUint16(f)
	case int16:
		return int16ToUint16(f)
	case int32:
		return int32ToUint16(f)
	case int64:
		return int64ToUint16(f)
	case uint:
		return uintToUint16(f)
	case uint8:
		return uint8ToUint16(f)
	case uint16:
		return f, true
	case uint32:
		return uint32ToUint16(f)
	case uint64:
		return uint64ToUint16(f)
	case float32:
		return float32ToUint16(f)
	case float64:
		return float64ToUint16(f)
	}
	return uint16(value), false
}

func toUint32[F numericType](value F) (uint32, bool) {
	switch f := any(value).(type) {
	case int:
		return intToUint32(f)
	case int8:
		return int8ToUint32(f)
	case int16:
		return int16ToUint32(f)
	case int32:
		return int32ToUint32(f)
	case int64:
		return int64ToUint32(f)
	case uint:
		return uintToUint32(f)
	case uint8:
		return uint8ToUint32(f)
	case uint16:
		return uint16ToUint32(f)
	case uint32:
		return f, true
	case uint64:
		return uint64ToUint32(f)
	case float32:
		return float32ToUint32(f)
	case float64:
		return float64ToUint32(f)
	}
	return uint32(value), false
}

func toUint64[F numericType](value F) (uint64, bool) {
	switch f := any(value).(type) {
	case int:
		return intToUint64(f)
	case int8:
		return int8ToUint64(f)
	case int16:
		return int16ToUint64(f)
	case int32:
		return int32ToUint64(f)
	case int64:
		return int64ToUint64(f)
	case uint:
		return uintToUint64(f)
	case uint8:
		return uint8ToUint64(f)
	case uint16:
		return uint16ToUint64(f)
	case uint32:
		return uint32ToUint64(f)
	case uint64:
		return f, true
	case float32:
		return float32ToUint64(f)
	case float64:
		return float64ToUint64(f)
	}
	return uint64(value), false
}

func toFloat32[F numericType](value F) (float32, bool) {
	switch f := any(value).(type) {
	case int:
		return intToFloat32(f)
	case int8:
		return int8ToFloat32(f)
	case int16:
		return int16ToFloat32(f)
	case int32:
		return int32ToFloat32(f)
	case int64:
		return int64ToFloat32(f)
	case uint:
		return uintToFloat32(f)
	case uint8:
		return uint8ToFloat32(f)
	case uint16:
		return uint16ToFloat32(f)
	case uint32:
		return uint32ToFloat32(f)
	case uint64:
		return uint64ToFloat32(f)
	case float32:
		return f, true
	case float64:
		return float64ToFloat32(f)
	}
	return float32(value), false
}

func toFloat64[F numericType](value F) (float64, bool) {
	switch f := any(value).(type) {
	case int:
		return intToFloat64(f)
	case int8:
		return int8ToFloat64(f)
	case int16:
		return int16ToFloat64(f)
	case int32:
		return int32ToFloat64(f)
	case int64:
		return int64ToFloat64(f)
	case uint:
		return uintToFloat64(f)
	case uint8:
		return uint8ToFloat64(f)
	case uint16:
		return uint16ToFloat64(f)
	case uint32:
		return uint32ToFloat64(f)
	case uint64:
		return uint64ToFloat64(f)
	case float32:
		return float32ToFloat64(f)
	case float64:
		return f, true
	}
	return float64(value), false
}
