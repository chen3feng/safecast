package safecast

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Number interface {
	Integer | Float
}

func To[T Number, F Number](value F) (to T, ok bool) {
	ok = true
	switch t := any(to).(type) {
	case int:
		t, ok = toInt(value)
		to = T(t)
	case int8:
		t, ok = toInt8(value)
		to = T(t)
	case int16:
		t, ok = toInt16(value)
		to = T(t)
	case int32:
		t, ok = toInt32(value)
		to = T(t)
	case int64:
		t, ok = toInt64(value)
		to = T(t)
	case uint:
		t, ok = toUint(value)
		to = T(t)
	case uint8:
		t, ok = toUint8(value)
		to = T(t)
	case uint16:
		t, ok = toUint16(value)
		to = T(t)
	case uint32:
		t, ok = toUint32(value)
		to = T(t)
	case uint64:
		t, ok = toUint64(value)
		to = T(t)
	case float32:
		t, ok = toFloat32(value)
		to = T(t)
	case float64:
		t, ok = toFloat64(value)
		to = T(t)
	}
	return to, ok
}

func toInt[F Number](value F) (int, bool) {
	switch f := any(value).(type) {
	case int:
		return f, true
	case int8:
		return Int8ToInt(f)
	case int16:
		return Int16ToInt(f)
	case int32:
		return Int32ToInt(f)
	case int64:
		return Int64ToInt(f)
	case uint:
		return UintToInt(f)
	case uint8:
		return Uint8ToInt(f)
	case uint16:
		return Uint16ToInt(f)
	case uint32:
		return Uint32ToInt(f)
	case uint64:
		return Uint64ToInt(f)
	case float32:
		return Float32ToInt(f)
	case float64:
		return Float64ToInt(f)
	}
	return int(value), false
}

func toInt8[F Number](value F) (int8, bool) {
	switch f := any(value).(type) {
	case int:
		return IntToInt8(f)
	case int8:
		return f, true
	case int16:
		return Int16ToInt8(f)
	case int32:
		return Int32ToInt8(f)
	case int64:
		return Int64ToInt8(f)
	case uint:
		return UintToInt8(f)
	case uint8:
		return Uint8ToInt8(f)
	case uint16:
		return Uint16ToInt8(f)
	case uint32:
		return Uint32ToInt8(f)
	case uint64:
		return Uint64ToInt8(f)
	case float32:
		return Float32ToInt8(f)
	case float64:
		return Float64ToInt8(f)
	}
	return int8(value), false
}

func toInt16[F Number](value F) (int16, bool) {
	switch f := any(value).(type) {
	case int:
		return IntToInt16(f)
	case int8:
		return Int8ToInt16(f)
	case int16:
		return f, true
	case int32:
		return Int32ToInt16(f)
	case int64:
		return Int64ToInt16(f)
	case uint:
		return UintToInt16(f)
	case uint8:
		return Uint8ToInt16(f)
	case uint16:
		return Uint16ToInt16(f)
	case uint32:
		return Uint32ToInt16(f)
	case uint64:
		return Uint64ToInt16(f)
	case float32:
		return Float32ToInt16(f)
	case float64:
		return Float64ToInt16(f)
	}
	return int16(value), false
}

func toInt32[F Number](value F) (int32, bool) {
	switch f := any(value).(type) {
	case int:
		return IntToInt32(f)
	case int8:
		return Int8ToInt32(f)
	case int16:
		return Int16ToInt32(f)
	case int32:
		return f, true
	case int64:
		return Int64ToInt32(f)
	case uint:
		return UintToInt32(f)
	case uint8:
		return Uint8ToInt32(f)
	case uint16:
		return Uint16ToInt32(f)
	case uint32:
		return Uint32ToInt32(f)
	case uint64:
		return Uint64ToInt32(f)
	case float32:
		return Float32ToInt32(f)
	case float64:
		return Float64ToInt32(f)
	}
	return int32(value), false
}

func toInt64[F Number](value F) (int64, bool) {
	switch f := any(value).(type) {
	case int:
		return IntToInt64(f)
	case int8:
		return Int8ToInt64(f)
	case int16:
		return Int16ToInt64(f)
	case int32:
		return Int32ToInt64(f)
	case int64:
		return f, true
	case uint:
		return UintToInt64(f)
	case uint8:
		return Uint8ToInt64(f)
	case uint16:
		return Uint16ToInt64(f)
	case uint32:
		return Uint32ToInt64(f)
	case uint64:
		return Uint64ToInt64(f)
	case float32:
		return Float32ToInt64(f)
	case float64:
		return Float64ToInt64(f)
	}
	return int64(value), false
}

func toUint[F Number](value F) (uint, bool) {
	switch f := any(value).(type) {
	case int:
		return IntToUint(f)
	case int8:
		return Int8ToUint(f)
	case int16:
		return Int16ToUint(f)
	case int32:
		return Int32ToUint(f)
	case int64:
		return Int64ToUint(f)
	case uint:
		return f, true
	case uint8:
		return Uint8ToUint(f)
	case uint16:
		return Uint16ToUint(f)
	case uint32:
		return Uint32ToUint(f)
	case uint64:
		return Uint64ToUint(f)
	case float32:
		return Float32ToUint(f)
	case float64:
		return Float64ToUint(f)
	}
	return uint(value), false
}

func toUint8[F Number](value F) (uint8, bool) {
	switch f := any(value).(type) {
	case int:
		return IntToUint8(f)
	case int8:
		return Int8ToUint8(f)
	case int16:
		return Int16ToUint8(f)
	case int32:
		return Int32ToUint8(f)
	case int64:
		return Int64ToUint8(f)
	case uint:
		return UintToUint8(f)
	case uint8:
		return f, true
	case uint16:
		return Uint16ToUint8(f)
	case uint32:
		return Uint32ToUint8(f)
	case uint64:
		return Uint64ToUint8(f)
	case float32:
		return Float32ToUint8(f)
	case float64:
		return Float64ToUint8(f)
	}
	return uint8(value), false
}

func toUint16[F Number](value F) (uint16, bool) {
	switch f := any(value).(type) {
	case int:
		return IntToUint16(f)
	case int8:
		return Int8ToUint16(f)
	case int16:
		return Int16ToUint16(f)
	case int32:
		return Int32ToUint16(f)
	case int64:
		return Int64ToUint16(f)
	case uint:
		return UintToUint16(f)
	case uint8:
		return Uint8ToUint16(f)
	case uint16:
		return f, true
	case uint32:
		return Uint32ToUint16(f)
	case uint64:
		return Uint64ToUint16(f)
	case float32:
		return Float32ToUint16(f)
	case float64:
		return Float64ToUint16(f)
	}
	return uint16(value), false
}

func toUint32[F Number](value F) (uint32, bool) {
	switch f := any(value).(type) {
	case int:
		return IntToUint32(f)
	case int8:
		return Int8ToUint32(f)
	case int16:
		return Int16ToUint32(f)
	case int32:
		return Int32ToUint32(f)
	case int64:
		return Int64ToUint32(f)
	case uint:
		return UintToUint32(f)
	case uint8:
		return Uint8ToUint32(f)
	case uint16:
		return Uint16ToUint32(f)
	case uint32:
		return f, true
	case uint64:
		return Uint64ToUint32(f)
	case float32:
		return Float32ToUint32(f)
	case float64:
		return Float64ToUint32(f)
	}
	return uint32(value), false
}

func toUint64[F Number](value F) (uint64, bool) {
	switch f := any(value).(type) {
	case int:
		return IntToUint64(f)
	case int8:
		return Int8ToUint64(f)
	case int16:
		return Int16ToUint64(f)
	case int32:
		return Int32ToUint64(f)
	case int64:
		return Int64ToUint64(f)
	case uint:
		return UintToUint64(f)
	case uint8:
		return Uint8ToUint64(f)
	case uint16:
		return Uint16ToUint64(f)
	case uint32:
		return Uint32ToUint64(f)
	case uint64:
		return f, true
	case float32:
		return Float32ToUint64(f)
	case float64:
		return Float64ToUint64(f)
	}
	return uint64(value), false
}

func toFloat32[F Number](value F) (float32, bool) {
	switch f := any(value).(type) {
	case int:
		return IntToFloat32(f)
	case int8:
		return Int8ToFloat32(f)
	case int16:
		return Int16ToFloat32(f)
	case int32:
		return Int32ToFloat32(f)
	case int64:
		return Int64ToFloat32(f)
	case uint:
		return UintToFloat32(f)
	case uint8:
		return Uint8ToFloat32(f)
	case uint16:
		return Uint16ToFloat32(f)
	case uint32:
		return Uint32ToFloat32(f)
	case uint64:
		return Uint64ToFloat32(f)
	case float32:
		return f, true
	case float64:
		return Float64ToFloat32(f)
	}
	return float32(value), false
}

func toFloat64[F Number](value F) (float64, bool) {
	switch f := any(value).(type) {
	case int:
		return IntToFloat64(f)
	case int8:
		return Int8ToFloat64(f)
	case int16:
		return Int16ToFloat64(f)
	case int32:
		return Int32ToFloat64(f)
	case int64:
		return Int64ToFloat64(f)
	case uint:
		return UintToFloat64(f)
	case uint8:
		return Uint8ToFloat64(f)
	case uint16:
		return Uint16ToFloat64(f)
	case uint32:
		return Uint32ToFloat64(f)
	case uint64:
		return Uint64ToFloat64(f)
	case float32:
		return Float32ToFloat64(f)
	case float64:
		return f, true
	}
	return float64(value), false
}
