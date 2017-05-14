package call

import (
	"time"
)

func lutMonadicVoidIdx0(fn interface{}) Caller {
	if to, ok := fn.(func(int)); ok {
		return Caller(monadicVoidIntCallFn(to))
	}
	if to, ok := fn.(func(int64)); ok {
		return Caller(monadicVoidInt64CallFn(to))
	}
	if to, ok := fn.(func(int32)); ok {
		return Caller(monadicVoidInt32CallFn(to))
	}
	if to, ok := fn.(func(time.Duration)); ok {
		return Caller(monadicVoidDurationCallFn(to))
	}
	return nil
}

func lutDyadicVoid0x0(fn interface{}) Caller {
	if to, ok := fn.(func(int, int)); ok {
		return Caller(dyadicVoidIntIntCallFn(to))
	}
	if to, ok := fn.(func(int64, int)); ok {
		return Caller(dyadicVoidInt64IntCallFn(to))
	}
	if to, ok := fn.(func(int32, int)); ok {
		return Caller(dyadicVoidInt32IntCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, int)); ok {
		return Caller(dyadicVoidDurationIntCallFn(to))
	}
	if to, ok := fn.(func(int, int64)); ok {
		return Caller(dyadicVoidIntInt64CallFn(to))
	}
	if to, ok := fn.(func(int64, int64)); ok {
		return Caller(dyadicVoidInt64Int64CallFn(to))
	}
	if to, ok := fn.(func(int32, int64)); ok {
		return Caller(dyadicVoidInt32Int64CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, int64)); ok {
		return Caller(dyadicVoidDurationInt64CallFn(to))
	}
	if to, ok := fn.(func(int, int32)); ok {
		return Caller(dyadicVoidIntInt32CallFn(to))
	}
	if to, ok := fn.(func(int64, int32)); ok {
		return Caller(dyadicVoidInt64Int32CallFn(to))
	}
	if to, ok := fn.(func(int32, int32)); ok {
		return Caller(dyadicVoidInt32Int32CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, int32)); ok {
		return Caller(dyadicVoidDurationInt32CallFn(to))
	}
	if to, ok := fn.(func(int, time.Duration)); ok {
		return Caller(dyadicVoidIntDurationCallFn(to))
	}
	if to, ok := fn.(func(int64, time.Duration)); ok {
		return Caller(dyadicVoidInt64DurationCallFn(to))
	}
	if to, ok := fn.(func(int32, time.Duration)); ok {
		return Caller(dyadicVoidInt32DurationCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, time.Duration)); ok {
		return Caller(dyadicVoidDurationDurationCallFn(to))
	}
	return nil
}

func lutDyadicVoid0x1(fn interface{}) Caller {
	if to, ok := fn.(func(int, uint)); ok {
		return Caller(dyadicVoidIntUintCallFn(to))
	}
	if to, ok := fn.(func(int64, uint)); ok {
		return Caller(dyadicVoidInt64UintCallFn(to))
	}
	if to, ok := fn.(func(int32, uint)); ok {
		return Caller(dyadicVoidInt32UintCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, uint)); ok {
		return Caller(dyadicVoidDurationUintCallFn(to))
	}
	if to, ok := fn.(func(int, uint64)); ok {
		return Caller(dyadicVoidIntUint64CallFn(to))
	}
	if to, ok := fn.(func(int64, uint64)); ok {
		return Caller(dyadicVoidInt64Uint64CallFn(to))
	}
	if to, ok := fn.(func(int32, uint64)); ok {
		return Caller(dyadicVoidInt32Uint64CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, uint64)); ok {
		return Caller(dyadicVoidDurationUint64CallFn(to))
	}
	if to, ok := fn.(func(int, uint32)); ok {
		return Caller(dyadicVoidIntUint32CallFn(to))
	}
	if to, ok := fn.(func(int64, uint32)); ok {
		return Caller(dyadicVoidInt64Uint32CallFn(to))
	}
	if to, ok := fn.(func(int32, uint32)); ok {
		return Caller(dyadicVoidInt32Uint32CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, uint32)); ok {
		return Caller(dyadicVoidDurationUint32CallFn(to))
	}
	if to, ok := fn.(func(int, uint8)); ok {
		return Caller(dyadicVoidIntUint8CallFn(to))
	}
	if to, ok := fn.(func(int64, uint8)); ok {
		return Caller(dyadicVoidInt64Uint8CallFn(to))
	}
	if to, ok := fn.(func(int32, uint8)); ok {
		return Caller(dyadicVoidInt32Uint8CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, uint8)); ok {
		return Caller(dyadicVoidDurationUint8CallFn(to))
	}
	return nil
}

func lutDyadicVoid0x2(fn interface{}) Caller {
	if to, ok := fn.(func(int, float64)); ok {
		return Caller(dyadicVoidIntFloat64CallFn(to))
	}
	if to, ok := fn.(func(int64, float64)); ok {
		return Caller(dyadicVoidInt64Float64CallFn(to))
	}
	if to, ok := fn.(func(int32, float64)); ok {
		return Caller(dyadicVoidInt32Float64CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, float64)); ok {
		return Caller(dyadicVoidDurationFloat64CallFn(to))
	}
	if to, ok := fn.(func(int, float32)); ok {
		return Caller(dyadicVoidIntFloat32CallFn(to))
	}
	if to, ok := fn.(func(int64, float32)); ok {
		return Caller(dyadicVoidInt64Float32CallFn(to))
	}
	if to, ok := fn.(func(int32, float32)); ok {
		return Caller(dyadicVoidInt32Float32CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, float32)); ok {
		return Caller(dyadicVoidDurationFloat32CallFn(to))
	}
	if to, ok := fn.(func(int, complex128)); ok {
		return Caller(dyadicVoidIntComplex128CallFn(to))
	}
	if to, ok := fn.(func(int64, complex128)); ok {
		return Caller(dyadicVoidInt64Complex128CallFn(to))
	}
	if to, ok := fn.(func(int32, complex128)); ok {
		return Caller(dyadicVoidInt32Complex128CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, complex128)); ok {
		return Caller(dyadicVoidDurationComplex128CallFn(to))
	}
	return nil
}

func lutDyadicVoid0x3(fn interface{}) Caller {
	if to, ok := fn.(func(int, string)); ok {
		return Caller(dyadicVoidIntStringCallFn(to))
	}
	if to, ok := fn.(func(int64, string)); ok {
		return Caller(dyadicVoidInt64StringCallFn(to))
	}
	if to, ok := fn.(func(int32, string)); ok {
		return Caller(dyadicVoidInt32StringCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, string)); ok {
		return Caller(dyadicVoidDurationStringCallFn(to))
	}
	if to, ok := fn.(func(int, []byte)); ok {
		return Caller(dyadicVoidIntBytesCallFn(to))
	}
	if to, ok := fn.(func(int64, []byte)); ok {
		return Caller(dyadicVoidInt64BytesCallFn(to))
	}
	if to, ok := fn.(func(int32, []byte)); ok {
		return Caller(dyadicVoidInt32BytesCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, []byte)); ok {
		return Caller(dyadicVoidDurationBytesCallFn(to))
	}
	if to, ok := fn.(func(int, bool)); ok {
		return Caller(dyadicVoidIntBoolCallFn(to))
	}
	if to, ok := fn.(func(int64, bool)); ok {
		return Caller(dyadicVoidInt64BoolCallFn(to))
	}
	if to, ok := fn.(func(int32, bool)); ok {
		return Caller(dyadicVoidInt32BoolCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, bool)); ok {
		return Caller(dyadicVoidDurationBoolCallFn(to))
	}
	if to, ok := fn.(func(int, time.Time)); ok {
		return Caller(dyadicVoidIntTimeCallFn(to))
	}
	if to, ok := fn.(func(int64, time.Time)); ok {
		return Caller(dyadicVoidInt64TimeCallFn(to))
	}
	if to, ok := fn.(func(int32, time.Time)); ok {
		return Caller(dyadicVoidInt32TimeCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, time.Time)); ok {
		return Caller(dyadicVoidDurationTimeCallFn(to))
	}
	return nil
}

func lutMonadicVoidIdx1(fn interface{}) Caller {
	if to, ok := fn.(func(uint)); ok {
		return Caller(monadicVoidUintCallFn(to))
	}
	if to, ok := fn.(func(uint64)); ok {
		return Caller(monadicVoidUint64CallFn(to))
	}
	if to, ok := fn.(func(uint32)); ok {
		return Caller(monadicVoidUint32CallFn(to))
	}
	if to, ok := fn.(func(uint8)); ok {
		return Caller(monadicVoidUint8CallFn(to))
	}
	return nil
}

func lutDyadicVoid1x0(fn interface{}) Caller {
	if to, ok := fn.(func(uint, int)); ok {
		return Caller(dyadicVoidUintIntCallFn(to))
	}
	if to, ok := fn.(func(uint64, int)); ok {
		return Caller(dyadicVoidUint64IntCallFn(to))
	}
	if to, ok := fn.(func(uint32, int)); ok {
		return Caller(dyadicVoidUint32IntCallFn(to))
	}
	if to, ok := fn.(func(uint8, int)); ok {
		return Caller(dyadicVoidUint8IntCallFn(to))
	}
	if to, ok := fn.(func(uint, int64)); ok {
		return Caller(dyadicVoidUintInt64CallFn(to))
	}
	if to, ok := fn.(func(uint64, int64)); ok {
		return Caller(dyadicVoidUint64Int64CallFn(to))
	}
	if to, ok := fn.(func(uint32, int64)); ok {
		return Caller(dyadicVoidUint32Int64CallFn(to))
	}
	if to, ok := fn.(func(uint8, int64)); ok {
		return Caller(dyadicVoidUint8Int64CallFn(to))
	}
	if to, ok := fn.(func(uint, int32)); ok {
		return Caller(dyadicVoidUintInt32CallFn(to))
	}
	if to, ok := fn.(func(uint64, int32)); ok {
		return Caller(dyadicVoidUint64Int32CallFn(to))
	}
	if to, ok := fn.(func(uint32, int32)); ok {
		return Caller(dyadicVoidUint32Int32CallFn(to))
	}
	if to, ok := fn.(func(uint8, int32)); ok {
		return Caller(dyadicVoidUint8Int32CallFn(to))
	}
	if to, ok := fn.(func(uint, time.Duration)); ok {
		return Caller(dyadicVoidUintDurationCallFn(to))
	}
	if to, ok := fn.(func(uint64, time.Duration)); ok {
		return Caller(dyadicVoidUint64DurationCallFn(to))
	}
	if to, ok := fn.(func(uint32, time.Duration)); ok {
		return Caller(dyadicVoidUint32DurationCallFn(to))
	}
	if to, ok := fn.(func(uint8, time.Duration)); ok {
		return Caller(dyadicVoidUint8DurationCallFn(to))
	}
	return nil
}

func lutDyadicVoid1x1(fn interface{}) Caller {
	if to, ok := fn.(func(uint, uint)); ok {
		return Caller(dyadicVoidUintUintCallFn(to))
	}
	if to, ok := fn.(func(uint64, uint)); ok {
		return Caller(dyadicVoidUint64UintCallFn(to))
	}
	if to, ok := fn.(func(uint32, uint)); ok {
		return Caller(dyadicVoidUint32UintCallFn(to))
	}
	if to, ok := fn.(func(uint8, uint)); ok {
		return Caller(dyadicVoidUint8UintCallFn(to))
	}
	if to, ok := fn.(func(uint, uint64)); ok {
		return Caller(dyadicVoidUintUint64CallFn(to))
	}
	if to, ok := fn.(func(uint64, uint64)); ok {
		return Caller(dyadicVoidUint64Uint64CallFn(to))
	}
	if to, ok := fn.(func(uint32, uint64)); ok {
		return Caller(dyadicVoidUint32Uint64CallFn(to))
	}
	if to, ok := fn.(func(uint8, uint64)); ok {
		return Caller(dyadicVoidUint8Uint64CallFn(to))
	}
	if to, ok := fn.(func(uint, uint32)); ok {
		return Caller(dyadicVoidUintUint32CallFn(to))
	}
	if to, ok := fn.(func(uint64, uint32)); ok {
		return Caller(dyadicVoidUint64Uint32CallFn(to))
	}
	if to, ok := fn.(func(uint32, uint32)); ok {
		return Caller(dyadicVoidUint32Uint32CallFn(to))
	}
	if to, ok := fn.(func(uint8, uint32)); ok {
		return Caller(dyadicVoidUint8Uint32CallFn(to))
	}
	if to, ok := fn.(func(uint, uint8)); ok {
		return Caller(dyadicVoidUintUint8CallFn(to))
	}
	if to, ok := fn.(func(uint64, uint8)); ok {
		return Caller(dyadicVoidUint64Uint8CallFn(to))
	}
	if to, ok := fn.(func(uint32, uint8)); ok {
		return Caller(dyadicVoidUint32Uint8CallFn(to))
	}
	if to, ok := fn.(func(uint8, uint8)); ok {
		return Caller(dyadicVoidUint8Uint8CallFn(to))
	}
	return nil
}

func lutDyadicVoid1x2(fn interface{}) Caller {
	if to, ok := fn.(func(uint, float64)); ok {
		return Caller(dyadicVoidUintFloat64CallFn(to))
	}
	if to, ok := fn.(func(uint64, float64)); ok {
		return Caller(dyadicVoidUint64Float64CallFn(to))
	}
	if to, ok := fn.(func(uint32, float64)); ok {
		return Caller(dyadicVoidUint32Float64CallFn(to))
	}
	if to, ok := fn.(func(uint8, float64)); ok {
		return Caller(dyadicVoidUint8Float64CallFn(to))
	}
	if to, ok := fn.(func(uint, float32)); ok {
		return Caller(dyadicVoidUintFloat32CallFn(to))
	}
	if to, ok := fn.(func(uint64, float32)); ok {
		return Caller(dyadicVoidUint64Float32CallFn(to))
	}
	if to, ok := fn.(func(uint32, float32)); ok {
		return Caller(dyadicVoidUint32Float32CallFn(to))
	}
	if to, ok := fn.(func(uint8, float32)); ok {
		return Caller(dyadicVoidUint8Float32CallFn(to))
	}
	if to, ok := fn.(func(uint, complex128)); ok {
		return Caller(dyadicVoidUintComplex128CallFn(to))
	}
	if to, ok := fn.(func(uint64, complex128)); ok {
		return Caller(dyadicVoidUint64Complex128CallFn(to))
	}
	if to, ok := fn.(func(uint32, complex128)); ok {
		return Caller(dyadicVoidUint32Complex128CallFn(to))
	}
	if to, ok := fn.(func(uint8, complex128)); ok {
		return Caller(dyadicVoidUint8Complex128CallFn(to))
	}
	return nil
}

func lutDyadicVoid1x3(fn interface{}) Caller {
	if to, ok := fn.(func(uint, string)); ok {
		return Caller(dyadicVoidUintStringCallFn(to))
	}
	if to, ok := fn.(func(uint64, string)); ok {
		return Caller(dyadicVoidUint64StringCallFn(to))
	}
	if to, ok := fn.(func(uint32, string)); ok {
		return Caller(dyadicVoidUint32StringCallFn(to))
	}
	if to, ok := fn.(func(uint8, string)); ok {
		return Caller(dyadicVoidUint8StringCallFn(to))
	}
	if to, ok := fn.(func(uint, []byte)); ok {
		return Caller(dyadicVoidUintBytesCallFn(to))
	}
	if to, ok := fn.(func(uint64, []byte)); ok {
		return Caller(dyadicVoidUint64BytesCallFn(to))
	}
	if to, ok := fn.(func(uint32, []byte)); ok {
		return Caller(dyadicVoidUint32BytesCallFn(to))
	}
	if to, ok := fn.(func(uint8, []byte)); ok {
		return Caller(dyadicVoidUint8BytesCallFn(to))
	}
	if to, ok := fn.(func(uint, bool)); ok {
		return Caller(dyadicVoidUintBoolCallFn(to))
	}
	if to, ok := fn.(func(uint64, bool)); ok {
		return Caller(dyadicVoidUint64BoolCallFn(to))
	}
	if to, ok := fn.(func(uint32, bool)); ok {
		return Caller(dyadicVoidUint32BoolCallFn(to))
	}
	if to, ok := fn.(func(uint8, bool)); ok {
		return Caller(dyadicVoidUint8BoolCallFn(to))
	}
	if to, ok := fn.(func(uint, time.Time)); ok {
		return Caller(dyadicVoidUintTimeCallFn(to))
	}
	if to, ok := fn.(func(uint64, time.Time)); ok {
		return Caller(dyadicVoidUint64TimeCallFn(to))
	}
	if to, ok := fn.(func(uint32, time.Time)); ok {
		return Caller(dyadicVoidUint32TimeCallFn(to))
	}
	if to, ok := fn.(func(uint8, time.Time)); ok {
		return Caller(dyadicVoidUint8TimeCallFn(to))
	}
	return nil
}

func lutMonadicVoidIdx2(fn interface{}) Caller {
	if to, ok := fn.(func(float64)); ok {
		return Caller(monadicVoidFloat64CallFn(to))
	}
	if to, ok := fn.(func(float32)); ok {
		return Caller(monadicVoidFloat32CallFn(to))
	}
	if to, ok := fn.(func(complex128)); ok {
		return Caller(monadicVoidComplex128CallFn(to))
	}
	return nil
}

func lutDyadicVoid2x0(fn interface{}) Caller {
	if to, ok := fn.(func(float64, int)); ok {
		return Caller(dyadicVoidFloat64IntCallFn(to))
	}
	if to, ok := fn.(func(float32, int)); ok {
		return Caller(dyadicVoidFloat32IntCallFn(to))
	}
	if to, ok := fn.(func(complex128, int)); ok {
		return Caller(dyadicVoidComplex128IntCallFn(to))
	}
	if to, ok := fn.(func(float64, int64)); ok {
		return Caller(dyadicVoidFloat64Int64CallFn(to))
	}
	if to, ok := fn.(func(float32, int64)); ok {
		return Caller(dyadicVoidFloat32Int64CallFn(to))
	}
	if to, ok := fn.(func(complex128, int64)); ok {
		return Caller(dyadicVoidComplex128Int64CallFn(to))
	}
	if to, ok := fn.(func(float64, int32)); ok {
		return Caller(dyadicVoidFloat64Int32CallFn(to))
	}
	if to, ok := fn.(func(float32, int32)); ok {
		return Caller(dyadicVoidFloat32Int32CallFn(to))
	}
	if to, ok := fn.(func(complex128, int32)); ok {
		return Caller(dyadicVoidComplex128Int32CallFn(to))
	}
	if to, ok := fn.(func(float64, time.Duration)); ok {
		return Caller(dyadicVoidFloat64DurationCallFn(to))
	}
	if to, ok := fn.(func(float32, time.Duration)); ok {
		return Caller(dyadicVoidFloat32DurationCallFn(to))
	}
	if to, ok := fn.(func(complex128, time.Duration)); ok {
		return Caller(dyadicVoidComplex128DurationCallFn(to))
	}
	return nil
}

func lutDyadicVoid2x1(fn interface{}) Caller {
	if to, ok := fn.(func(float64, uint)); ok {
		return Caller(dyadicVoidFloat64UintCallFn(to))
	}
	if to, ok := fn.(func(float32, uint)); ok {
		return Caller(dyadicVoidFloat32UintCallFn(to))
	}
	if to, ok := fn.(func(complex128, uint)); ok {
		return Caller(dyadicVoidComplex128UintCallFn(to))
	}
	if to, ok := fn.(func(float64, uint64)); ok {
		return Caller(dyadicVoidFloat64Uint64CallFn(to))
	}
	if to, ok := fn.(func(float32, uint64)); ok {
		return Caller(dyadicVoidFloat32Uint64CallFn(to))
	}
	if to, ok := fn.(func(complex128, uint64)); ok {
		return Caller(dyadicVoidComplex128Uint64CallFn(to))
	}
	if to, ok := fn.(func(float64, uint32)); ok {
		return Caller(dyadicVoidFloat64Uint32CallFn(to))
	}
	if to, ok := fn.(func(float32, uint32)); ok {
		return Caller(dyadicVoidFloat32Uint32CallFn(to))
	}
	if to, ok := fn.(func(complex128, uint32)); ok {
		return Caller(dyadicVoidComplex128Uint32CallFn(to))
	}
	if to, ok := fn.(func(float64, uint8)); ok {
		return Caller(dyadicVoidFloat64Uint8CallFn(to))
	}
	if to, ok := fn.(func(float32, uint8)); ok {
		return Caller(dyadicVoidFloat32Uint8CallFn(to))
	}
	if to, ok := fn.(func(complex128, uint8)); ok {
		return Caller(dyadicVoidComplex128Uint8CallFn(to))
	}
	return nil
}

func lutDyadicVoid2x2(fn interface{}) Caller {
	if to, ok := fn.(func(float64, float64)); ok {
		return Caller(dyadicVoidFloat64Float64CallFn(to))
	}
	if to, ok := fn.(func(float32, float64)); ok {
		return Caller(dyadicVoidFloat32Float64CallFn(to))
	}
	if to, ok := fn.(func(complex128, float64)); ok {
		return Caller(dyadicVoidComplex128Float64CallFn(to))
	}
	if to, ok := fn.(func(float64, float32)); ok {
		return Caller(dyadicVoidFloat64Float32CallFn(to))
	}
	if to, ok := fn.(func(float32, float32)); ok {
		return Caller(dyadicVoidFloat32Float32CallFn(to))
	}
	if to, ok := fn.(func(complex128, float32)); ok {
		return Caller(dyadicVoidComplex128Float32CallFn(to))
	}
	if to, ok := fn.(func(float64, complex128)); ok {
		return Caller(dyadicVoidFloat64Complex128CallFn(to))
	}
	if to, ok := fn.(func(float32, complex128)); ok {
		return Caller(dyadicVoidFloat32Complex128CallFn(to))
	}
	if to, ok := fn.(func(complex128, complex128)); ok {
		return Caller(dyadicVoidComplex128Complex128CallFn(to))
	}
	return nil
}

func lutDyadicVoid2x3(fn interface{}) Caller {
	if to, ok := fn.(func(float64, string)); ok {
		return Caller(dyadicVoidFloat64StringCallFn(to))
	}
	if to, ok := fn.(func(float32, string)); ok {
		return Caller(dyadicVoidFloat32StringCallFn(to))
	}
	if to, ok := fn.(func(complex128, string)); ok {
		return Caller(dyadicVoidComplex128StringCallFn(to))
	}
	if to, ok := fn.(func(float64, []byte)); ok {
		return Caller(dyadicVoidFloat64BytesCallFn(to))
	}
	if to, ok := fn.(func(float32, []byte)); ok {
		return Caller(dyadicVoidFloat32BytesCallFn(to))
	}
	if to, ok := fn.(func(complex128, []byte)); ok {
		return Caller(dyadicVoidComplex128BytesCallFn(to))
	}
	if to, ok := fn.(func(float64, bool)); ok {
		return Caller(dyadicVoidFloat64BoolCallFn(to))
	}
	if to, ok := fn.(func(float32, bool)); ok {
		return Caller(dyadicVoidFloat32BoolCallFn(to))
	}
	if to, ok := fn.(func(complex128, bool)); ok {
		return Caller(dyadicVoidComplex128BoolCallFn(to))
	}
	if to, ok := fn.(func(float64, time.Time)); ok {
		return Caller(dyadicVoidFloat64TimeCallFn(to))
	}
	if to, ok := fn.(func(float32, time.Time)); ok {
		return Caller(dyadicVoidFloat32TimeCallFn(to))
	}
	if to, ok := fn.(func(complex128, time.Time)); ok {
		return Caller(dyadicVoidComplex128TimeCallFn(to))
	}
	return nil
}

func lutMonadicVoidIdx3(fn interface{}) Caller {
	if to, ok := fn.(func(string)); ok {
		return Caller(monadicVoidStringCallFn(to))
	}
	if to, ok := fn.(func([]byte)); ok {
		return Caller(monadicVoidBytesCallFn(to))
	}
	if to, ok := fn.(func(bool)); ok {
		return Caller(monadicVoidBoolCallFn(to))
	}
	if to, ok := fn.(func(time.Time)); ok {
		return Caller(monadicVoidTimeCallFn(to))
	}
	return nil
}

func lutDyadicVoid3x0(fn interface{}) Caller {
	if to, ok := fn.(func(string, int)); ok {
		return Caller(dyadicVoidStringIntCallFn(to))
	}
	if to, ok := fn.(func([]byte, int)); ok {
		return Caller(dyadicVoidBytesIntCallFn(to))
	}
	if to, ok := fn.(func(bool, int)); ok {
		return Caller(dyadicVoidBoolIntCallFn(to))
	}
	if to, ok := fn.(func(time.Time, int)); ok {
		return Caller(dyadicVoidTimeIntCallFn(to))
	}
	if to, ok := fn.(func(string, int64)); ok {
		return Caller(dyadicVoidStringInt64CallFn(to))
	}
	if to, ok := fn.(func([]byte, int64)); ok {
		return Caller(dyadicVoidBytesInt64CallFn(to))
	}
	if to, ok := fn.(func(bool, int64)); ok {
		return Caller(dyadicVoidBoolInt64CallFn(to))
	}
	if to, ok := fn.(func(time.Time, int64)); ok {
		return Caller(dyadicVoidTimeInt64CallFn(to))
	}
	if to, ok := fn.(func(string, int32)); ok {
		return Caller(dyadicVoidStringInt32CallFn(to))
	}
	if to, ok := fn.(func([]byte, int32)); ok {
		return Caller(dyadicVoidBytesInt32CallFn(to))
	}
	if to, ok := fn.(func(bool, int32)); ok {
		return Caller(dyadicVoidBoolInt32CallFn(to))
	}
	if to, ok := fn.(func(time.Time, int32)); ok {
		return Caller(dyadicVoidTimeInt32CallFn(to))
	}
	if to, ok := fn.(func(string, time.Duration)); ok {
		return Caller(dyadicVoidStringDurationCallFn(to))
	}
	if to, ok := fn.(func([]byte, time.Duration)); ok {
		return Caller(dyadicVoidBytesDurationCallFn(to))
	}
	if to, ok := fn.(func(bool, time.Duration)); ok {
		return Caller(dyadicVoidBoolDurationCallFn(to))
	}
	if to, ok := fn.(func(time.Time, time.Duration)); ok {
		return Caller(dyadicVoidTimeDurationCallFn(to))
	}
	return nil
}

func lutDyadicVoid3x1(fn interface{}) Caller {
	if to, ok := fn.(func(string, uint)); ok {
		return Caller(dyadicVoidStringUintCallFn(to))
	}
	if to, ok := fn.(func([]byte, uint)); ok {
		return Caller(dyadicVoidBytesUintCallFn(to))
	}
	if to, ok := fn.(func(bool, uint)); ok {
		return Caller(dyadicVoidBoolUintCallFn(to))
	}
	if to, ok := fn.(func(time.Time, uint)); ok {
		return Caller(dyadicVoidTimeUintCallFn(to))
	}
	if to, ok := fn.(func(string, uint64)); ok {
		return Caller(dyadicVoidStringUint64CallFn(to))
	}
	if to, ok := fn.(func([]byte, uint64)); ok {
		return Caller(dyadicVoidBytesUint64CallFn(to))
	}
	if to, ok := fn.(func(bool, uint64)); ok {
		return Caller(dyadicVoidBoolUint64CallFn(to))
	}
	if to, ok := fn.(func(time.Time, uint64)); ok {
		return Caller(dyadicVoidTimeUint64CallFn(to))
	}
	if to, ok := fn.(func(string, uint32)); ok {
		return Caller(dyadicVoidStringUint32CallFn(to))
	}
	if to, ok := fn.(func([]byte, uint32)); ok {
		return Caller(dyadicVoidBytesUint32CallFn(to))
	}
	if to, ok := fn.(func(bool, uint32)); ok {
		return Caller(dyadicVoidBoolUint32CallFn(to))
	}
	if to, ok := fn.(func(time.Time, uint32)); ok {
		return Caller(dyadicVoidTimeUint32CallFn(to))
	}
	if to, ok := fn.(func(string, uint8)); ok {
		return Caller(dyadicVoidStringUint8CallFn(to))
	}
	if to, ok := fn.(func([]byte, uint8)); ok {
		return Caller(dyadicVoidBytesUint8CallFn(to))
	}
	if to, ok := fn.(func(bool, uint8)); ok {
		return Caller(dyadicVoidBoolUint8CallFn(to))
	}
	if to, ok := fn.(func(time.Time, uint8)); ok {
		return Caller(dyadicVoidTimeUint8CallFn(to))
	}
	return nil
}

func lutDyadicVoid3x2(fn interface{}) Caller {
	if to, ok := fn.(func(string, float64)); ok {
		return Caller(dyadicVoidStringFloat64CallFn(to))
	}
	if to, ok := fn.(func([]byte, float64)); ok {
		return Caller(dyadicVoidBytesFloat64CallFn(to))
	}
	if to, ok := fn.(func(bool, float64)); ok {
		return Caller(dyadicVoidBoolFloat64CallFn(to))
	}
	if to, ok := fn.(func(time.Time, float64)); ok {
		return Caller(dyadicVoidTimeFloat64CallFn(to))
	}
	if to, ok := fn.(func(string, float32)); ok {
		return Caller(dyadicVoidStringFloat32CallFn(to))
	}
	if to, ok := fn.(func([]byte, float32)); ok {
		return Caller(dyadicVoidBytesFloat32CallFn(to))
	}
	if to, ok := fn.(func(bool, float32)); ok {
		return Caller(dyadicVoidBoolFloat32CallFn(to))
	}
	if to, ok := fn.(func(time.Time, float32)); ok {
		return Caller(dyadicVoidTimeFloat32CallFn(to))
	}
	if to, ok := fn.(func(string, complex128)); ok {
		return Caller(dyadicVoidStringComplex128CallFn(to))
	}
	if to, ok := fn.(func([]byte, complex128)); ok {
		return Caller(dyadicVoidBytesComplex128CallFn(to))
	}
	if to, ok := fn.(func(bool, complex128)); ok {
		return Caller(dyadicVoidBoolComplex128CallFn(to))
	}
	if to, ok := fn.(func(time.Time, complex128)); ok {
		return Caller(dyadicVoidTimeComplex128CallFn(to))
	}
	return nil
}

func lutDyadicVoid3x3(fn interface{}) Caller {
	if to, ok := fn.(func(string, string)); ok {
		return Caller(dyadicVoidStringStringCallFn(to))
	}
	if to, ok := fn.(func([]byte, string)); ok {
		return Caller(dyadicVoidBytesStringCallFn(to))
	}
	if to, ok := fn.(func(bool, string)); ok {
		return Caller(dyadicVoidBoolStringCallFn(to))
	}
	if to, ok := fn.(func(time.Time, string)); ok {
		return Caller(dyadicVoidTimeStringCallFn(to))
	}
	if to, ok := fn.(func(string, []byte)); ok {
		return Caller(dyadicVoidStringBytesCallFn(to))
	}
	if to, ok := fn.(func([]byte, []byte)); ok {
		return Caller(dyadicVoidBytesBytesCallFn(to))
	}
	if to, ok := fn.(func(bool, []byte)); ok {
		return Caller(dyadicVoidBoolBytesCallFn(to))
	}
	if to, ok := fn.(func(time.Time, []byte)); ok {
		return Caller(dyadicVoidTimeBytesCallFn(to))
	}
	if to, ok := fn.(func(string, bool)); ok {
		return Caller(dyadicVoidStringBoolCallFn(to))
	}
	if to, ok := fn.(func([]byte, bool)); ok {
		return Caller(dyadicVoidBytesBoolCallFn(to))
	}
	if to, ok := fn.(func(bool, bool)); ok {
		return Caller(dyadicVoidBoolBoolCallFn(to))
	}
	if to, ok := fn.(func(time.Time, bool)); ok {
		return Caller(dyadicVoidTimeBoolCallFn(to))
	}
	if to, ok := fn.(func(string, time.Time)); ok {
		return Caller(dyadicVoidStringTimeCallFn(to))
	}
	if to, ok := fn.(func([]byte, time.Time)); ok {
		return Caller(dyadicVoidBytesTimeCallFn(to))
	}
	if to, ok := fn.(func(bool, time.Time)); ok {
		return Caller(dyadicVoidBoolTimeCallFn(to))
	}
	if to, ok := fn.(func(time.Time, time.Time)); ok {
		return Caller(dyadicVoidTimeTimeCallFn(to))
	}
	return nil
}

func lutMonadicErrorIdx0(fn interface{}) Caller {
	if to, ok := fn.(func(int) error); ok {
		return Caller(monadicErrorIntCallFn(to))
	}
	if to, ok := fn.(func(int64) error); ok {
		return Caller(monadicErrorInt64CallFn(to))
	}
	if to, ok := fn.(func(int32) error); ok {
		return Caller(monadicErrorInt32CallFn(to))
	}
	if to, ok := fn.(func(time.Duration) error); ok {
		return Caller(monadicErrorDurationCallFn(to))
	}
	return nil
}

func lutDyadicError0x0(fn interface{}) Caller {
	if to, ok := fn.(func(int, int) error); ok {
		return Caller(dyadicErrorIntIntCallFn(to))
	}
	if to, ok := fn.(func(int64, int) error); ok {
		return Caller(dyadicErrorInt64IntCallFn(to))
	}
	if to, ok := fn.(func(int32, int) error); ok {
		return Caller(dyadicErrorInt32IntCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, int) error); ok {
		return Caller(dyadicErrorDurationIntCallFn(to))
	}
	if to, ok := fn.(func(int, int64) error); ok {
		return Caller(dyadicErrorIntInt64CallFn(to))
	}
	if to, ok := fn.(func(int64, int64) error); ok {
		return Caller(dyadicErrorInt64Int64CallFn(to))
	}
	if to, ok := fn.(func(int32, int64) error); ok {
		return Caller(dyadicErrorInt32Int64CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, int64) error); ok {
		return Caller(dyadicErrorDurationInt64CallFn(to))
	}
	if to, ok := fn.(func(int, int32) error); ok {
		return Caller(dyadicErrorIntInt32CallFn(to))
	}
	if to, ok := fn.(func(int64, int32) error); ok {
		return Caller(dyadicErrorInt64Int32CallFn(to))
	}
	if to, ok := fn.(func(int32, int32) error); ok {
		return Caller(dyadicErrorInt32Int32CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, int32) error); ok {
		return Caller(dyadicErrorDurationInt32CallFn(to))
	}
	if to, ok := fn.(func(int, time.Duration) error); ok {
		return Caller(dyadicErrorIntDurationCallFn(to))
	}
	if to, ok := fn.(func(int64, time.Duration) error); ok {
		return Caller(dyadicErrorInt64DurationCallFn(to))
	}
	if to, ok := fn.(func(int32, time.Duration) error); ok {
		return Caller(dyadicErrorInt32DurationCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, time.Duration) error); ok {
		return Caller(dyadicErrorDurationDurationCallFn(to))
	}
	return nil
}

func lutDyadicError0x1(fn interface{}) Caller {
	if to, ok := fn.(func(int, uint) error); ok {
		return Caller(dyadicErrorIntUintCallFn(to))
	}
	if to, ok := fn.(func(int64, uint) error); ok {
		return Caller(dyadicErrorInt64UintCallFn(to))
	}
	if to, ok := fn.(func(int32, uint) error); ok {
		return Caller(dyadicErrorInt32UintCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, uint) error); ok {
		return Caller(dyadicErrorDurationUintCallFn(to))
	}
	if to, ok := fn.(func(int, uint64) error); ok {
		return Caller(dyadicErrorIntUint64CallFn(to))
	}
	if to, ok := fn.(func(int64, uint64) error); ok {
		return Caller(dyadicErrorInt64Uint64CallFn(to))
	}
	if to, ok := fn.(func(int32, uint64) error); ok {
		return Caller(dyadicErrorInt32Uint64CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, uint64) error); ok {
		return Caller(dyadicErrorDurationUint64CallFn(to))
	}
	if to, ok := fn.(func(int, uint32) error); ok {
		return Caller(dyadicErrorIntUint32CallFn(to))
	}
	if to, ok := fn.(func(int64, uint32) error); ok {
		return Caller(dyadicErrorInt64Uint32CallFn(to))
	}
	if to, ok := fn.(func(int32, uint32) error); ok {
		return Caller(dyadicErrorInt32Uint32CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, uint32) error); ok {
		return Caller(dyadicErrorDurationUint32CallFn(to))
	}
	if to, ok := fn.(func(int, uint8) error); ok {
		return Caller(dyadicErrorIntUint8CallFn(to))
	}
	if to, ok := fn.(func(int64, uint8) error); ok {
		return Caller(dyadicErrorInt64Uint8CallFn(to))
	}
	if to, ok := fn.(func(int32, uint8) error); ok {
		return Caller(dyadicErrorInt32Uint8CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, uint8) error); ok {
		return Caller(dyadicErrorDurationUint8CallFn(to))
	}
	return nil
}

func lutDyadicError0x2(fn interface{}) Caller {
	if to, ok := fn.(func(int, float64) error); ok {
		return Caller(dyadicErrorIntFloat64CallFn(to))
	}
	if to, ok := fn.(func(int64, float64) error); ok {
		return Caller(dyadicErrorInt64Float64CallFn(to))
	}
	if to, ok := fn.(func(int32, float64) error); ok {
		return Caller(dyadicErrorInt32Float64CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, float64) error); ok {
		return Caller(dyadicErrorDurationFloat64CallFn(to))
	}
	if to, ok := fn.(func(int, float32) error); ok {
		return Caller(dyadicErrorIntFloat32CallFn(to))
	}
	if to, ok := fn.(func(int64, float32) error); ok {
		return Caller(dyadicErrorInt64Float32CallFn(to))
	}
	if to, ok := fn.(func(int32, float32) error); ok {
		return Caller(dyadicErrorInt32Float32CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, float32) error); ok {
		return Caller(dyadicErrorDurationFloat32CallFn(to))
	}
	if to, ok := fn.(func(int, complex128) error); ok {
		return Caller(dyadicErrorIntComplex128CallFn(to))
	}
	if to, ok := fn.(func(int64, complex128) error); ok {
		return Caller(dyadicErrorInt64Complex128CallFn(to))
	}
	if to, ok := fn.(func(int32, complex128) error); ok {
		return Caller(dyadicErrorInt32Complex128CallFn(to))
	}
	if to, ok := fn.(func(time.Duration, complex128) error); ok {
		return Caller(dyadicErrorDurationComplex128CallFn(to))
	}
	return nil
}

func lutDyadicError0x3(fn interface{}) Caller {
	if to, ok := fn.(func(int, string) error); ok {
		return Caller(dyadicErrorIntStringCallFn(to))
	}
	if to, ok := fn.(func(int64, string) error); ok {
		return Caller(dyadicErrorInt64StringCallFn(to))
	}
	if to, ok := fn.(func(int32, string) error); ok {
		return Caller(dyadicErrorInt32StringCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, string) error); ok {
		return Caller(dyadicErrorDurationStringCallFn(to))
	}
	if to, ok := fn.(func(int, []byte) error); ok {
		return Caller(dyadicErrorIntBytesCallFn(to))
	}
	if to, ok := fn.(func(int64, []byte) error); ok {
		return Caller(dyadicErrorInt64BytesCallFn(to))
	}
	if to, ok := fn.(func(int32, []byte) error); ok {
		return Caller(dyadicErrorInt32BytesCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, []byte) error); ok {
		return Caller(dyadicErrorDurationBytesCallFn(to))
	}
	if to, ok := fn.(func(int, bool) error); ok {
		return Caller(dyadicErrorIntBoolCallFn(to))
	}
	if to, ok := fn.(func(int64, bool) error); ok {
		return Caller(dyadicErrorInt64BoolCallFn(to))
	}
	if to, ok := fn.(func(int32, bool) error); ok {
		return Caller(dyadicErrorInt32BoolCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, bool) error); ok {
		return Caller(dyadicErrorDurationBoolCallFn(to))
	}
	if to, ok := fn.(func(int, time.Time) error); ok {
		return Caller(dyadicErrorIntTimeCallFn(to))
	}
	if to, ok := fn.(func(int64, time.Time) error); ok {
		return Caller(dyadicErrorInt64TimeCallFn(to))
	}
	if to, ok := fn.(func(int32, time.Time) error); ok {
		return Caller(dyadicErrorInt32TimeCallFn(to))
	}
	if to, ok := fn.(func(time.Duration, time.Time) error); ok {
		return Caller(dyadicErrorDurationTimeCallFn(to))
	}
	return nil
}

func lutMonadicErrorIdx1(fn interface{}) Caller {
	if to, ok := fn.(func(uint) error); ok {
		return Caller(monadicErrorUintCallFn(to))
	}
	if to, ok := fn.(func(uint64) error); ok {
		return Caller(monadicErrorUint64CallFn(to))
	}
	if to, ok := fn.(func(uint32) error); ok {
		return Caller(monadicErrorUint32CallFn(to))
	}
	if to, ok := fn.(func(uint8) error); ok {
		return Caller(monadicErrorUint8CallFn(to))
	}
	return nil
}

func lutDyadicError1x0(fn interface{}) Caller {
	if to, ok := fn.(func(uint, int) error); ok {
		return Caller(dyadicErrorUintIntCallFn(to))
	}
	if to, ok := fn.(func(uint64, int) error); ok {
		return Caller(dyadicErrorUint64IntCallFn(to))
	}
	if to, ok := fn.(func(uint32, int) error); ok {
		return Caller(dyadicErrorUint32IntCallFn(to))
	}
	if to, ok := fn.(func(uint8, int) error); ok {
		return Caller(dyadicErrorUint8IntCallFn(to))
	}
	if to, ok := fn.(func(uint, int64) error); ok {
		return Caller(dyadicErrorUintInt64CallFn(to))
	}
	if to, ok := fn.(func(uint64, int64) error); ok {
		return Caller(dyadicErrorUint64Int64CallFn(to))
	}
	if to, ok := fn.(func(uint32, int64) error); ok {
		return Caller(dyadicErrorUint32Int64CallFn(to))
	}
	if to, ok := fn.(func(uint8, int64) error); ok {
		return Caller(dyadicErrorUint8Int64CallFn(to))
	}
	if to, ok := fn.(func(uint, int32) error); ok {
		return Caller(dyadicErrorUintInt32CallFn(to))
	}
	if to, ok := fn.(func(uint64, int32) error); ok {
		return Caller(dyadicErrorUint64Int32CallFn(to))
	}
	if to, ok := fn.(func(uint32, int32) error); ok {
		return Caller(dyadicErrorUint32Int32CallFn(to))
	}
	if to, ok := fn.(func(uint8, int32) error); ok {
		return Caller(dyadicErrorUint8Int32CallFn(to))
	}
	if to, ok := fn.(func(uint, time.Duration) error); ok {
		return Caller(dyadicErrorUintDurationCallFn(to))
	}
	if to, ok := fn.(func(uint64, time.Duration) error); ok {
		return Caller(dyadicErrorUint64DurationCallFn(to))
	}
	if to, ok := fn.(func(uint32, time.Duration) error); ok {
		return Caller(dyadicErrorUint32DurationCallFn(to))
	}
	if to, ok := fn.(func(uint8, time.Duration) error); ok {
		return Caller(dyadicErrorUint8DurationCallFn(to))
	}
	return nil
}

func lutDyadicError1x1(fn interface{}) Caller {
	if to, ok := fn.(func(uint, uint) error); ok {
		return Caller(dyadicErrorUintUintCallFn(to))
	}
	if to, ok := fn.(func(uint64, uint) error); ok {
		return Caller(dyadicErrorUint64UintCallFn(to))
	}
	if to, ok := fn.(func(uint32, uint) error); ok {
		return Caller(dyadicErrorUint32UintCallFn(to))
	}
	if to, ok := fn.(func(uint8, uint) error); ok {
		return Caller(dyadicErrorUint8UintCallFn(to))
	}
	if to, ok := fn.(func(uint, uint64) error); ok {
		return Caller(dyadicErrorUintUint64CallFn(to))
	}
	if to, ok := fn.(func(uint64, uint64) error); ok {
		return Caller(dyadicErrorUint64Uint64CallFn(to))
	}
	if to, ok := fn.(func(uint32, uint64) error); ok {
		return Caller(dyadicErrorUint32Uint64CallFn(to))
	}
	if to, ok := fn.(func(uint8, uint64) error); ok {
		return Caller(dyadicErrorUint8Uint64CallFn(to))
	}
	if to, ok := fn.(func(uint, uint32) error); ok {
		return Caller(dyadicErrorUintUint32CallFn(to))
	}
	if to, ok := fn.(func(uint64, uint32) error); ok {
		return Caller(dyadicErrorUint64Uint32CallFn(to))
	}
	if to, ok := fn.(func(uint32, uint32) error); ok {
		return Caller(dyadicErrorUint32Uint32CallFn(to))
	}
	if to, ok := fn.(func(uint8, uint32) error); ok {
		return Caller(dyadicErrorUint8Uint32CallFn(to))
	}
	if to, ok := fn.(func(uint, uint8) error); ok {
		return Caller(dyadicErrorUintUint8CallFn(to))
	}
	if to, ok := fn.(func(uint64, uint8) error); ok {
		return Caller(dyadicErrorUint64Uint8CallFn(to))
	}
	if to, ok := fn.(func(uint32, uint8) error); ok {
		return Caller(dyadicErrorUint32Uint8CallFn(to))
	}
	if to, ok := fn.(func(uint8, uint8) error); ok {
		return Caller(dyadicErrorUint8Uint8CallFn(to))
	}
	return nil
}

func lutDyadicError1x2(fn interface{}) Caller {
	if to, ok := fn.(func(uint, float64) error); ok {
		return Caller(dyadicErrorUintFloat64CallFn(to))
	}
	if to, ok := fn.(func(uint64, float64) error); ok {
		return Caller(dyadicErrorUint64Float64CallFn(to))
	}
	if to, ok := fn.(func(uint32, float64) error); ok {
		return Caller(dyadicErrorUint32Float64CallFn(to))
	}
	if to, ok := fn.(func(uint8, float64) error); ok {
		return Caller(dyadicErrorUint8Float64CallFn(to))
	}
	if to, ok := fn.(func(uint, float32) error); ok {
		return Caller(dyadicErrorUintFloat32CallFn(to))
	}
	if to, ok := fn.(func(uint64, float32) error); ok {
		return Caller(dyadicErrorUint64Float32CallFn(to))
	}
	if to, ok := fn.(func(uint32, float32) error); ok {
		return Caller(dyadicErrorUint32Float32CallFn(to))
	}
	if to, ok := fn.(func(uint8, float32) error); ok {
		return Caller(dyadicErrorUint8Float32CallFn(to))
	}
	if to, ok := fn.(func(uint, complex128) error); ok {
		return Caller(dyadicErrorUintComplex128CallFn(to))
	}
	if to, ok := fn.(func(uint64, complex128) error); ok {
		return Caller(dyadicErrorUint64Complex128CallFn(to))
	}
	if to, ok := fn.(func(uint32, complex128) error); ok {
		return Caller(dyadicErrorUint32Complex128CallFn(to))
	}
	if to, ok := fn.(func(uint8, complex128) error); ok {
		return Caller(dyadicErrorUint8Complex128CallFn(to))
	}
	return nil
}

func lutDyadicError1x3(fn interface{}) Caller {
	if to, ok := fn.(func(uint, string) error); ok {
		return Caller(dyadicErrorUintStringCallFn(to))
	}
	if to, ok := fn.(func(uint64, string) error); ok {
		return Caller(dyadicErrorUint64StringCallFn(to))
	}
	if to, ok := fn.(func(uint32, string) error); ok {
		return Caller(dyadicErrorUint32StringCallFn(to))
	}
	if to, ok := fn.(func(uint8, string) error); ok {
		return Caller(dyadicErrorUint8StringCallFn(to))
	}
	if to, ok := fn.(func(uint, []byte) error); ok {
		return Caller(dyadicErrorUintBytesCallFn(to))
	}
	if to, ok := fn.(func(uint64, []byte) error); ok {
		return Caller(dyadicErrorUint64BytesCallFn(to))
	}
	if to, ok := fn.(func(uint32, []byte) error); ok {
		return Caller(dyadicErrorUint32BytesCallFn(to))
	}
	if to, ok := fn.(func(uint8, []byte) error); ok {
		return Caller(dyadicErrorUint8BytesCallFn(to))
	}
	if to, ok := fn.(func(uint, bool) error); ok {
		return Caller(dyadicErrorUintBoolCallFn(to))
	}
	if to, ok := fn.(func(uint64, bool) error); ok {
		return Caller(dyadicErrorUint64BoolCallFn(to))
	}
	if to, ok := fn.(func(uint32, bool) error); ok {
		return Caller(dyadicErrorUint32BoolCallFn(to))
	}
	if to, ok := fn.(func(uint8, bool) error); ok {
		return Caller(dyadicErrorUint8BoolCallFn(to))
	}
	if to, ok := fn.(func(uint, time.Time) error); ok {
		return Caller(dyadicErrorUintTimeCallFn(to))
	}
	if to, ok := fn.(func(uint64, time.Time) error); ok {
		return Caller(dyadicErrorUint64TimeCallFn(to))
	}
	if to, ok := fn.(func(uint32, time.Time) error); ok {
		return Caller(dyadicErrorUint32TimeCallFn(to))
	}
	if to, ok := fn.(func(uint8, time.Time) error); ok {
		return Caller(dyadicErrorUint8TimeCallFn(to))
	}
	return nil
}

func lutMonadicErrorIdx2(fn interface{}) Caller {
	if to, ok := fn.(func(float64) error); ok {
		return Caller(monadicErrorFloat64CallFn(to))
	}
	if to, ok := fn.(func(float32) error); ok {
		return Caller(monadicErrorFloat32CallFn(to))
	}
	if to, ok := fn.(func(complex128) error); ok {
		return Caller(monadicErrorComplex128CallFn(to))
	}
	return nil
}

func lutDyadicError2x0(fn interface{}) Caller {
	if to, ok := fn.(func(float64, int) error); ok {
		return Caller(dyadicErrorFloat64IntCallFn(to))
	}
	if to, ok := fn.(func(float32, int) error); ok {
		return Caller(dyadicErrorFloat32IntCallFn(to))
	}
	if to, ok := fn.(func(complex128, int) error); ok {
		return Caller(dyadicErrorComplex128IntCallFn(to))
	}
	if to, ok := fn.(func(float64, int64) error); ok {
		return Caller(dyadicErrorFloat64Int64CallFn(to))
	}
	if to, ok := fn.(func(float32, int64) error); ok {
		return Caller(dyadicErrorFloat32Int64CallFn(to))
	}
	if to, ok := fn.(func(complex128, int64) error); ok {
		return Caller(dyadicErrorComplex128Int64CallFn(to))
	}
	if to, ok := fn.(func(float64, int32) error); ok {
		return Caller(dyadicErrorFloat64Int32CallFn(to))
	}
	if to, ok := fn.(func(float32, int32) error); ok {
		return Caller(dyadicErrorFloat32Int32CallFn(to))
	}
	if to, ok := fn.(func(complex128, int32) error); ok {
		return Caller(dyadicErrorComplex128Int32CallFn(to))
	}
	if to, ok := fn.(func(float64, time.Duration) error); ok {
		return Caller(dyadicErrorFloat64DurationCallFn(to))
	}
	if to, ok := fn.(func(float32, time.Duration) error); ok {
		return Caller(dyadicErrorFloat32DurationCallFn(to))
	}
	if to, ok := fn.(func(complex128, time.Duration) error); ok {
		return Caller(dyadicErrorComplex128DurationCallFn(to))
	}
	return nil
}

func lutDyadicError2x1(fn interface{}) Caller {
	if to, ok := fn.(func(float64, uint) error); ok {
		return Caller(dyadicErrorFloat64UintCallFn(to))
	}
	if to, ok := fn.(func(float32, uint) error); ok {
		return Caller(dyadicErrorFloat32UintCallFn(to))
	}
	if to, ok := fn.(func(complex128, uint) error); ok {
		return Caller(dyadicErrorComplex128UintCallFn(to))
	}
	if to, ok := fn.(func(float64, uint64) error); ok {
		return Caller(dyadicErrorFloat64Uint64CallFn(to))
	}
	if to, ok := fn.(func(float32, uint64) error); ok {
		return Caller(dyadicErrorFloat32Uint64CallFn(to))
	}
	if to, ok := fn.(func(complex128, uint64) error); ok {
		return Caller(dyadicErrorComplex128Uint64CallFn(to))
	}
	if to, ok := fn.(func(float64, uint32) error); ok {
		return Caller(dyadicErrorFloat64Uint32CallFn(to))
	}
	if to, ok := fn.(func(float32, uint32) error); ok {
		return Caller(dyadicErrorFloat32Uint32CallFn(to))
	}
	if to, ok := fn.(func(complex128, uint32) error); ok {
		return Caller(dyadicErrorComplex128Uint32CallFn(to))
	}
	if to, ok := fn.(func(float64, uint8) error); ok {
		return Caller(dyadicErrorFloat64Uint8CallFn(to))
	}
	if to, ok := fn.(func(float32, uint8) error); ok {
		return Caller(dyadicErrorFloat32Uint8CallFn(to))
	}
	if to, ok := fn.(func(complex128, uint8) error); ok {
		return Caller(dyadicErrorComplex128Uint8CallFn(to))
	}
	return nil
}

func lutDyadicError2x2(fn interface{}) Caller {
	if to, ok := fn.(func(float64, float64) error); ok {
		return Caller(dyadicErrorFloat64Float64CallFn(to))
	}
	if to, ok := fn.(func(float32, float64) error); ok {
		return Caller(dyadicErrorFloat32Float64CallFn(to))
	}
	if to, ok := fn.(func(complex128, float64) error); ok {
		return Caller(dyadicErrorComplex128Float64CallFn(to))
	}
	if to, ok := fn.(func(float64, float32) error); ok {
		return Caller(dyadicErrorFloat64Float32CallFn(to))
	}
	if to, ok := fn.(func(float32, float32) error); ok {
		return Caller(dyadicErrorFloat32Float32CallFn(to))
	}
	if to, ok := fn.(func(complex128, float32) error); ok {
		return Caller(dyadicErrorComplex128Float32CallFn(to))
	}
	if to, ok := fn.(func(float64, complex128) error); ok {
		return Caller(dyadicErrorFloat64Complex128CallFn(to))
	}
	if to, ok := fn.(func(float32, complex128) error); ok {
		return Caller(dyadicErrorFloat32Complex128CallFn(to))
	}
	if to, ok := fn.(func(complex128, complex128) error); ok {
		return Caller(dyadicErrorComplex128Complex128CallFn(to))
	}
	return nil
}

func lutDyadicError2x3(fn interface{}) Caller {
	if to, ok := fn.(func(float64, string) error); ok {
		return Caller(dyadicErrorFloat64StringCallFn(to))
	}
	if to, ok := fn.(func(float32, string) error); ok {
		return Caller(dyadicErrorFloat32StringCallFn(to))
	}
	if to, ok := fn.(func(complex128, string) error); ok {
		return Caller(dyadicErrorComplex128StringCallFn(to))
	}
	if to, ok := fn.(func(float64, []byte) error); ok {
		return Caller(dyadicErrorFloat64BytesCallFn(to))
	}
	if to, ok := fn.(func(float32, []byte) error); ok {
		return Caller(dyadicErrorFloat32BytesCallFn(to))
	}
	if to, ok := fn.(func(complex128, []byte) error); ok {
		return Caller(dyadicErrorComplex128BytesCallFn(to))
	}
	if to, ok := fn.(func(float64, bool) error); ok {
		return Caller(dyadicErrorFloat64BoolCallFn(to))
	}
	if to, ok := fn.(func(float32, bool) error); ok {
		return Caller(dyadicErrorFloat32BoolCallFn(to))
	}
	if to, ok := fn.(func(complex128, bool) error); ok {
		return Caller(dyadicErrorComplex128BoolCallFn(to))
	}
	if to, ok := fn.(func(float64, time.Time) error); ok {
		return Caller(dyadicErrorFloat64TimeCallFn(to))
	}
	if to, ok := fn.(func(float32, time.Time) error); ok {
		return Caller(dyadicErrorFloat32TimeCallFn(to))
	}
	if to, ok := fn.(func(complex128, time.Time) error); ok {
		return Caller(dyadicErrorComplex128TimeCallFn(to))
	}
	return nil
}

func lutMonadicErrorIdx3(fn interface{}) Caller {
	if to, ok := fn.(func(string) error); ok {
		return Caller(monadicErrorStringCallFn(to))
	}
	if to, ok := fn.(func([]byte) error); ok {
		return Caller(monadicErrorBytesCallFn(to))
	}
	if to, ok := fn.(func(bool) error); ok {
		return Caller(monadicErrorBoolCallFn(to))
	}
	if to, ok := fn.(func(time.Time) error); ok {
		return Caller(monadicErrorTimeCallFn(to))
	}
	return nil
}

func lutDyadicError3x0(fn interface{}) Caller {
	if to, ok := fn.(func(string, int) error); ok {
		return Caller(dyadicErrorStringIntCallFn(to))
	}
	if to, ok := fn.(func([]byte, int) error); ok {
		return Caller(dyadicErrorBytesIntCallFn(to))
	}
	if to, ok := fn.(func(bool, int) error); ok {
		return Caller(dyadicErrorBoolIntCallFn(to))
	}
	if to, ok := fn.(func(time.Time, int) error); ok {
		return Caller(dyadicErrorTimeIntCallFn(to))
	}
	if to, ok := fn.(func(string, int64) error); ok {
		return Caller(dyadicErrorStringInt64CallFn(to))
	}
	if to, ok := fn.(func([]byte, int64) error); ok {
		return Caller(dyadicErrorBytesInt64CallFn(to))
	}
	if to, ok := fn.(func(bool, int64) error); ok {
		return Caller(dyadicErrorBoolInt64CallFn(to))
	}
	if to, ok := fn.(func(time.Time, int64) error); ok {
		return Caller(dyadicErrorTimeInt64CallFn(to))
	}
	if to, ok := fn.(func(string, int32) error); ok {
		return Caller(dyadicErrorStringInt32CallFn(to))
	}
	if to, ok := fn.(func([]byte, int32) error); ok {
		return Caller(dyadicErrorBytesInt32CallFn(to))
	}
	if to, ok := fn.(func(bool, int32) error); ok {
		return Caller(dyadicErrorBoolInt32CallFn(to))
	}
	if to, ok := fn.(func(time.Time, int32) error); ok {
		return Caller(dyadicErrorTimeInt32CallFn(to))
	}
	if to, ok := fn.(func(string, time.Duration) error); ok {
		return Caller(dyadicErrorStringDurationCallFn(to))
	}
	if to, ok := fn.(func([]byte, time.Duration) error); ok {
		return Caller(dyadicErrorBytesDurationCallFn(to))
	}
	if to, ok := fn.(func(bool, time.Duration) error); ok {
		return Caller(dyadicErrorBoolDurationCallFn(to))
	}
	if to, ok := fn.(func(time.Time, time.Duration) error); ok {
		return Caller(dyadicErrorTimeDurationCallFn(to))
	}
	return nil
}

func lutDyadicError3x1(fn interface{}) Caller {
	if to, ok := fn.(func(string, uint) error); ok {
		return Caller(dyadicErrorStringUintCallFn(to))
	}
	if to, ok := fn.(func([]byte, uint) error); ok {
		return Caller(dyadicErrorBytesUintCallFn(to))
	}
	if to, ok := fn.(func(bool, uint) error); ok {
		return Caller(dyadicErrorBoolUintCallFn(to))
	}
	if to, ok := fn.(func(time.Time, uint) error); ok {
		return Caller(dyadicErrorTimeUintCallFn(to))
	}
	if to, ok := fn.(func(string, uint64) error); ok {
		return Caller(dyadicErrorStringUint64CallFn(to))
	}
	if to, ok := fn.(func([]byte, uint64) error); ok {
		return Caller(dyadicErrorBytesUint64CallFn(to))
	}
	if to, ok := fn.(func(bool, uint64) error); ok {
		return Caller(dyadicErrorBoolUint64CallFn(to))
	}
	if to, ok := fn.(func(time.Time, uint64) error); ok {
		return Caller(dyadicErrorTimeUint64CallFn(to))
	}
	if to, ok := fn.(func(string, uint32) error); ok {
		return Caller(dyadicErrorStringUint32CallFn(to))
	}
	if to, ok := fn.(func([]byte, uint32) error); ok {
		return Caller(dyadicErrorBytesUint32CallFn(to))
	}
	if to, ok := fn.(func(bool, uint32) error); ok {
		return Caller(dyadicErrorBoolUint32CallFn(to))
	}
	if to, ok := fn.(func(time.Time, uint32) error); ok {
		return Caller(dyadicErrorTimeUint32CallFn(to))
	}
	if to, ok := fn.(func(string, uint8) error); ok {
		return Caller(dyadicErrorStringUint8CallFn(to))
	}
	if to, ok := fn.(func([]byte, uint8) error); ok {
		return Caller(dyadicErrorBytesUint8CallFn(to))
	}
	if to, ok := fn.(func(bool, uint8) error); ok {
		return Caller(dyadicErrorBoolUint8CallFn(to))
	}
	if to, ok := fn.(func(time.Time, uint8) error); ok {
		return Caller(dyadicErrorTimeUint8CallFn(to))
	}
	return nil
}

func lutDyadicError3x2(fn interface{}) Caller {
	if to, ok := fn.(func(string, float64) error); ok {
		return Caller(dyadicErrorStringFloat64CallFn(to))
	}
	if to, ok := fn.(func([]byte, float64) error); ok {
		return Caller(dyadicErrorBytesFloat64CallFn(to))
	}
	if to, ok := fn.(func(bool, float64) error); ok {
		return Caller(dyadicErrorBoolFloat64CallFn(to))
	}
	if to, ok := fn.(func(time.Time, float64) error); ok {
		return Caller(dyadicErrorTimeFloat64CallFn(to))
	}
	if to, ok := fn.(func(string, float32) error); ok {
		return Caller(dyadicErrorStringFloat32CallFn(to))
	}
	if to, ok := fn.(func([]byte, float32) error); ok {
		return Caller(dyadicErrorBytesFloat32CallFn(to))
	}
	if to, ok := fn.(func(bool, float32) error); ok {
		return Caller(dyadicErrorBoolFloat32CallFn(to))
	}
	if to, ok := fn.(func(time.Time, float32) error); ok {
		return Caller(dyadicErrorTimeFloat32CallFn(to))
	}
	if to, ok := fn.(func(string, complex128) error); ok {
		return Caller(dyadicErrorStringComplex128CallFn(to))
	}
	if to, ok := fn.(func([]byte, complex128) error); ok {
		return Caller(dyadicErrorBytesComplex128CallFn(to))
	}
	if to, ok := fn.(func(bool, complex128) error); ok {
		return Caller(dyadicErrorBoolComplex128CallFn(to))
	}
	if to, ok := fn.(func(time.Time, complex128) error); ok {
		return Caller(dyadicErrorTimeComplex128CallFn(to))
	}
	return nil
}

func lutDyadicError3x3(fn interface{}) Caller {
	if to, ok := fn.(func(string, string) error); ok {
		return Caller(dyadicErrorStringStringCallFn(to))
	}
	if to, ok := fn.(func([]byte, string) error); ok {
		return Caller(dyadicErrorBytesStringCallFn(to))
	}
	if to, ok := fn.(func(bool, string) error); ok {
		return Caller(dyadicErrorBoolStringCallFn(to))
	}
	if to, ok := fn.(func(time.Time, string) error); ok {
		return Caller(dyadicErrorTimeStringCallFn(to))
	}
	if to, ok := fn.(func(string, []byte) error); ok {
		return Caller(dyadicErrorStringBytesCallFn(to))
	}
	if to, ok := fn.(func([]byte, []byte) error); ok {
		return Caller(dyadicErrorBytesBytesCallFn(to))
	}
	if to, ok := fn.(func(bool, []byte) error); ok {
		return Caller(dyadicErrorBoolBytesCallFn(to))
	}
	if to, ok := fn.(func(time.Time, []byte) error); ok {
		return Caller(dyadicErrorTimeBytesCallFn(to))
	}
	if to, ok := fn.(func(string, bool) error); ok {
		return Caller(dyadicErrorStringBoolCallFn(to))
	}
	if to, ok := fn.(func([]byte, bool) error); ok {
		return Caller(dyadicErrorBytesBoolCallFn(to))
	}
	if to, ok := fn.(func(bool, bool) error); ok {
		return Caller(dyadicErrorBoolBoolCallFn(to))
	}
	if to, ok := fn.(func(time.Time, bool) error); ok {
		return Caller(dyadicErrorTimeBoolCallFn(to))
	}
	if to, ok := fn.(func(string, time.Time) error); ok {
		return Caller(dyadicErrorStringTimeCallFn(to))
	}
	if to, ok := fn.(func([]byte, time.Time) error); ok {
		return Caller(dyadicErrorBytesTimeCallFn(to))
	}
	if to, ok := fn.(func(bool, time.Time) error); ok {
		return Caller(dyadicErrorBoolTimeCallFn(to))
	}
	if to, ok := fn.(func(time.Time, time.Time) error); ok {
		return Caller(dyadicErrorTimeTimeCallFn(to))
	}
	return nil
}

func lutMonadicVoidPtrIdx0(fn interface{}) Caller {
	if to, ok := fn.(func(*int)); ok {
		return Caller(monadicVoidPtrIntCallFn(to))
	}
	if to, ok := fn.(func(*int64)); ok {
		return Caller(monadicVoidPtrInt64CallFn(to))
	}
	if to, ok := fn.(func(*int32)); ok {
		return Caller(monadicVoidPtrInt32CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration)); ok {
		return Caller(monadicVoidPtrDurationCallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr0x0(fn interface{}) Caller {
	if to, ok := fn.(func(*int, *int)); ok {
		return Caller(dyadicVoidPtrIntIntCallFn(to))
	}
	if to, ok := fn.(func(*int64, *int)); ok {
		return Caller(dyadicVoidPtrInt64IntCallFn(to))
	}
	if to, ok := fn.(func(*int32, *int)); ok {
		return Caller(dyadicVoidPtrInt32IntCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *int)); ok {
		return Caller(dyadicVoidPtrDurationIntCallFn(to))
	}
	if to, ok := fn.(func(*int, *int64)); ok {
		return Caller(dyadicVoidPtrIntInt64CallFn(to))
	}
	if to, ok := fn.(func(*int64, *int64)); ok {
		return Caller(dyadicVoidPtrInt64Int64CallFn(to))
	}
	if to, ok := fn.(func(*int32, *int64)); ok {
		return Caller(dyadicVoidPtrInt32Int64CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *int64)); ok {
		return Caller(dyadicVoidPtrDurationInt64CallFn(to))
	}
	if to, ok := fn.(func(*int, *int32)); ok {
		return Caller(dyadicVoidPtrIntInt32CallFn(to))
	}
	if to, ok := fn.(func(*int64, *int32)); ok {
		return Caller(dyadicVoidPtrInt64Int32CallFn(to))
	}
	if to, ok := fn.(func(*int32, *int32)); ok {
		return Caller(dyadicVoidPtrInt32Int32CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *int32)); ok {
		return Caller(dyadicVoidPtrDurationInt32CallFn(to))
	}
	if to, ok := fn.(func(*int, *time.Duration)); ok {
		return Caller(dyadicVoidPtrIntDurationCallFn(to))
	}
	if to, ok := fn.(func(*int64, *time.Duration)); ok {
		return Caller(dyadicVoidPtrInt64DurationCallFn(to))
	}
	if to, ok := fn.(func(*int32, *time.Duration)); ok {
		return Caller(dyadicVoidPtrInt32DurationCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *time.Duration)); ok {
		return Caller(dyadicVoidPtrDurationDurationCallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr0x1(fn interface{}) Caller {
	if to, ok := fn.(func(*int, *uint)); ok {
		return Caller(dyadicVoidPtrIntUintCallFn(to))
	}
	if to, ok := fn.(func(*int64, *uint)); ok {
		return Caller(dyadicVoidPtrInt64UintCallFn(to))
	}
	if to, ok := fn.(func(*int32, *uint)); ok {
		return Caller(dyadicVoidPtrInt32UintCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *uint)); ok {
		return Caller(dyadicVoidPtrDurationUintCallFn(to))
	}
	if to, ok := fn.(func(*int, *uint64)); ok {
		return Caller(dyadicVoidPtrIntUint64CallFn(to))
	}
	if to, ok := fn.(func(*int64, *uint64)); ok {
		return Caller(dyadicVoidPtrInt64Uint64CallFn(to))
	}
	if to, ok := fn.(func(*int32, *uint64)); ok {
		return Caller(dyadicVoidPtrInt32Uint64CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *uint64)); ok {
		return Caller(dyadicVoidPtrDurationUint64CallFn(to))
	}
	if to, ok := fn.(func(*int, *uint32)); ok {
		return Caller(dyadicVoidPtrIntUint32CallFn(to))
	}
	if to, ok := fn.(func(*int64, *uint32)); ok {
		return Caller(dyadicVoidPtrInt64Uint32CallFn(to))
	}
	if to, ok := fn.(func(*int32, *uint32)); ok {
		return Caller(dyadicVoidPtrInt32Uint32CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *uint32)); ok {
		return Caller(dyadicVoidPtrDurationUint32CallFn(to))
	}
	if to, ok := fn.(func(*int, *uint8)); ok {
		return Caller(dyadicVoidPtrIntUint8CallFn(to))
	}
	if to, ok := fn.(func(*int64, *uint8)); ok {
		return Caller(dyadicVoidPtrInt64Uint8CallFn(to))
	}
	if to, ok := fn.(func(*int32, *uint8)); ok {
		return Caller(dyadicVoidPtrInt32Uint8CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *uint8)); ok {
		return Caller(dyadicVoidPtrDurationUint8CallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr0x2(fn interface{}) Caller {
	if to, ok := fn.(func(*int, *float64)); ok {
		return Caller(dyadicVoidPtrIntFloat64CallFn(to))
	}
	if to, ok := fn.(func(*int64, *float64)); ok {
		return Caller(dyadicVoidPtrInt64Float64CallFn(to))
	}
	if to, ok := fn.(func(*int32, *float64)); ok {
		return Caller(dyadicVoidPtrInt32Float64CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *float64)); ok {
		return Caller(dyadicVoidPtrDurationFloat64CallFn(to))
	}
	if to, ok := fn.(func(*int, *float32)); ok {
		return Caller(dyadicVoidPtrIntFloat32CallFn(to))
	}
	if to, ok := fn.(func(*int64, *float32)); ok {
		return Caller(dyadicVoidPtrInt64Float32CallFn(to))
	}
	if to, ok := fn.(func(*int32, *float32)); ok {
		return Caller(dyadicVoidPtrInt32Float32CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *float32)); ok {
		return Caller(dyadicVoidPtrDurationFloat32CallFn(to))
	}
	if to, ok := fn.(func(*int, *complex128)); ok {
		return Caller(dyadicVoidPtrIntComplex128CallFn(to))
	}
	if to, ok := fn.(func(*int64, *complex128)); ok {
		return Caller(dyadicVoidPtrInt64Complex128CallFn(to))
	}
	if to, ok := fn.(func(*int32, *complex128)); ok {
		return Caller(dyadicVoidPtrInt32Complex128CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *complex128)); ok {
		return Caller(dyadicVoidPtrDurationComplex128CallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr0x3(fn interface{}) Caller {
	if to, ok := fn.(func(*int, *string)); ok {
		return Caller(dyadicVoidPtrIntStringCallFn(to))
	}
	if to, ok := fn.(func(*int64, *string)); ok {
		return Caller(dyadicVoidPtrInt64StringCallFn(to))
	}
	if to, ok := fn.(func(*int32, *string)); ok {
		return Caller(dyadicVoidPtrInt32StringCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *string)); ok {
		return Caller(dyadicVoidPtrDurationStringCallFn(to))
	}
	if to, ok := fn.(func(*int, *[]byte)); ok {
		return Caller(dyadicVoidPtrIntBytesCallFn(to))
	}
	if to, ok := fn.(func(*int64, *[]byte)); ok {
		return Caller(dyadicVoidPtrInt64BytesCallFn(to))
	}
	if to, ok := fn.(func(*int32, *[]byte)); ok {
		return Caller(dyadicVoidPtrInt32BytesCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *[]byte)); ok {
		return Caller(dyadicVoidPtrDurationBytesCallFn(to))
	}
	if to, ok := fn.(func(*int, *bool)); ok {
		return Caller(dyadicVoidPtrIntBoolCallFn(to))
	}
	if to, ok := fn.(func(*int64, *bool)); ok {
		return Caller(dyadicVoidPtrInt64BoolCallFn(to))
	}
	if to, ok := fn.(func(*int32, *bool)); ok {
		return Caller(dyadicVoidPtrInt32BoolCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *bool)); ok {
		return Caller(dyadicVoidPtrDurationBoolCallFn(to))
	}
	if to, ok := fn.(func(*int, *time.Time)); ok {
		return Caller(dyadicVoidPtrIntTimeCallFn(to))
	}
	if to, ok := fn.(func(*int64, *time.Time)); ok {
		return Caller(dyadicVoidPtrInt64TimeCallFn(to))
	}
	if to, ok := fn.(func(*int32, *time.Time)); ok {
		return Caller(dyadicVoidPtrInt32TimeCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *time.Time)); ok {
		return Caller(dyadicVoidPtrDurationTimeCallFn(to))
	}
	return nil
}

func lutMonadicVoidPtrIdx1(fn interface{}) Caller {
	if to, ok := fn.(func(*uint)); ok {
		return Caller(monadicVoidPtrUintCallFn(to))
	}
	if to, ok := fn.(func(*uint64)); ok {
		return Caller(monadicVoidPtrUint64CallFn(to))
	}
	if to, ok := fn.(func(*uint32)); ok {
		return Caller(monadicVoidPtrUint32CallFn(to))
	}
	if to, ok := fn.(func(*uint8)); ok {
		return Caller(monadicVoidPtrUint8CallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr1x0(fn interface{}) Caller {
	if to, ok := fn.(func(*uint, *int)); ok {
		return Caller(dyadicVoidPtrUintIntCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *int)); ok {
		return Caller(dyadicVoidPtrUint64IntCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *int)); ok {
		return Caller(dyadicVoidPtrUint32IntCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *int)); ok {
		return Caller(dyadicVoidPtrUint8IntCallFn(to))
	}
	if to, ok := fn.(func(*uint, *int64)); ok {
		return Caller(dyadicVoidPtrUintInt64CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *int64)); ok {
		return Caller(dyadicVoidPtrUint64Int64CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *int64)); ok {
		return Caller(dyadicVoidPtrUint32Int64CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *int64)); ok {
		return Caller(dyadicVoidPtrUint8Int64CallFn(to))
	}
	if to, ok := fn.(func(*uint, *int32)); ok {
		return Caller(dyadicVoidPtrUintInt32CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *int32)); ok {
		return Caller(dyadicVoidPtrUint64Int32CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *int32)); ok {
		return Caller(dyadicVoidPtrUint32Int32CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *int32)); ok {
		return Caller(dyadicVoidPtrUint8Int32CallFn(to))
	}
	if to, ok := fn.(func(*uint, *time.Duration)); ok {
		return Caller(dyadicVoidPtrUintDurationCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *time.Duration)); ok {
		return Caller(dyadicVoidPtrUint64DurationCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *time.Duration)); ok {
		return Caller(dyadicVoidPtrUint32DurationCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *time.Duration)); ok {
		return Caller(dyadicVoidPtrUint8DurationCallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr1x1(fn interface{}) Caller {
	if to, ok := fn.(func(*uint, *uint)); ok {
		return Caller(dyadicVoidPtrUintUintCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *uint)); ok {
		return Caller(dyadicVoidPtrUint64UintCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *uint)); ok {
		return Caller(dyadicVoidPtrUint32UintCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *uint)); ok {
		return Caller(dyadicVoidPtrUint8UintCallFn(to))
	}
	if to, ok := fn.(func(*uint, *uint64)); ok {
		return Caller(dyadicVoidPtrUintUint64CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *uint64)); ok {
		return Caller(dyadicVoidPtrUint64Uint64CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *uint64)); ok {
		return Caller(dyadicVoidPtrUint32Uint64CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *uint64)); ok {
		return Caller(dyadicVoidPtrUint8Uint64CallFn(to))
	}
	if to, ok := fn.(func(*uint, *uint32)); ok {
		return Caller(dyadicVoidPtrUintUint32CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *uint32)); ok {
		return Caller(dyadicVoidPtrUint64Uint32CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *uint32)); ok {
		return Caller(dyadicVoidPtrUint32Uint32CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *uint32)); ok {
		return Caller(dyadicVoidPtrUint8Uint32CallFn(to))
	}
	if to, ok := fn.(func(*uint, *uint8)); ok {
		return Caller(dyadicVoidPtrUintUint8CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *uint8)); ok {
		return Caller(dyadicVoidPtrUint64Uint8CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *uint8)); ok {
		return Caller(dyadicVoidPtrUint32Uint8CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *uint8)); ok {
		return Caller(dyadicVoidPtrUint8Uint8CallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr1x2(fn interface{}) Caller {
	if to, ok := fn.(func(*uint, *float64)); ok {
		return Caller(dyadicVoidPtrUintFloat64CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *float64)); ok {
		return Caller(dyadicVoidPtrUint64Float64CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *float64)); ok {
		return Caller(dyadicVoidPtrUint32Float64CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *float64)); ok {
		return Caller(dyadicVoidPtrUint8Float64CallFn(to))
	}
	if to, ok := fn.(func(*uint, *float32)); ok {
		return Caller(dyadicVoidPtrUintFloat32CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *float32)); ok {
		return Caller(dyadicVoidPtrUint64Float32CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *float32)); ok {
		return Caller(dyadicVoidPtrUint32Float32CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *float32)); ok {
		return Caller(dyadicVoidPtrUint8Float32CallFn(to))
	}
	if to, ok := fn.(func(*uint, *complex128)); ok {
		return Caller(dyadicVoidPtrUintComplex128CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *complex128)); ok {
		return Caller(dyadicVoidPtrUint64Complex128CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *complex128)); ok {
		return Caller(dyadicVoidPtrUint32Complex128CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *complex128)); ok {
		return Caller(dyadicVoidPtrUint8Complex128CallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr1x3(fn interface{}) Caller {
	if to, ok := fn.(func(*uint, *string)); ok {
		return Caller(dyadicVoidPtrUintStringCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *string)); ok {
		return Caller(dyadicVoidPtrUint64StringCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *string)); ok {
		return Caller(dyadicVoidPtrUint32StringCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *string)); ok {
		return Caller(dyadicVoidPtrUint8StringCallFn(to))
	}
	if to, ok := fn.(func(*uint, *[]byte)); ok {
		return Caller(dyadicVoidPtrUintBytesCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *[]byte)); ok {
		return Caller(dyadicVoidPtrUint64BytesCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *[]byte)); ok {
		return Caller(dyadicVoidPtrUint32BytesCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *[]byte)); ok {
		return Caller(dyadicVoidPtrUint8BytesCallFn(to))
	}
	if to, ok := fn.(func(*uint, *bool)); ok {
		return Caller(dyadicVoidPtrUintBoolCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *bool)); ok {
		return Caller(dyadicVoidPtrUint64BoolCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *bool)); ok {
		return Caller(dyadicVoidPtrUint32BoolCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *bool)); ok {
		return Caller(dyadicVoidPtrUint8BoolCallFn(to))
	}
	if to, ok := fn.(func(*uint, *time.Time)); ok {
		return Caller(dyadicVoidPtrUintTimeCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *time.Time)); ok {
		return Caller(dyadicVoidPtrUint64TimeCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *time.Time)); ok {
		return Caller(dyadicVoidPtrUint32TimeCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *time.Time)); ok {
		return Caller(dyadicVoidPtrUint8TimeCallFn(to))
	}
	return nil
}

func lutMonadicVoidPtrIdx2(fn interface{}) Caller {
	if to, ok := fn.(func(*float64)); ok {
		return Caller(monadicVoidPtrFloat64CallFn(to))
	}
	if to, ok := fn.(func(*float32)); ok {
		return Caller(monadicVoidPtrFloat32CallFn(to))
	}
	if to, ok := fn.(func(*complex128)); ok {
		return Caller(monadicVoidPtrComplex128CallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr2x0(fn interface{}) Caller {
	if to, ok := fn.(func(*float64, *int)); ok {
		return Caller(dyadicVoidPtrFloat64IntCallFn(to))
	}
	if to, ok := fn.(func(*float32, *int)); ok {
		return Caller(dyadicVoidPtrFloat32IntCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *int)); ok {
		return Caller(dyadicVoidPtrComplex128IntCallFn(to))
	}
	if to, ok := fn.(func(*float64, *int64)); ok {
		return Caller(dyadicVoidPtrFloat64Int64CallFn(to))
	}
	if to, ok := fn.(func(*float32, *int64)); ok {
		return Caller(dyadicVoidPtrFloat32Int64CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *int64)); ok {
		return Caller(dyadicVoidPtrComplex128Int64CallFn(to))
	}
	if to, ok := fn.(func(*float64, *int32)); ok {
		return Caller(dyadicVoidPtrFloat64Int32CallFn(to))
	}
	if to, ok := fn.(func(*float32, *int32)); ok {
		return Caller(dyadicVoidPtrFloat32Int32CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *int32)); ok {
		return Caller(dyadicVoidPtrComplex128Int32CallFn(to))
	}
	if to, ok := fn.(func(*float64, *time.Duration)); ok {
		return Caller(dyadicVoidPtrFloat64DurationCallFn(to))
	}
	if to, ok := fn.(func(*float32, *time.Duration)); ok {
		return Caller(dyadicVoidPtrFloat32DurationCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *time.Duration)); ok {
		return Caller(dyadicVoidPtrComplex128DurationCallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr2x1(fn interface{}) Caller {
	if to, ok := fn.(func(*float64, *uint)); ok {
		return Caller(dyadicVoidPtrFloat64UintCallFn(to))
	}
	if to, ok := fn.(func(*float32, *uint)); ok {
		return Caller(dyadicVoidPtrFloat32UintCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *uint)); ok {
		return Caller(dyadicVoidPtrComplex128UintCallFn(to))
	}
	if to, ok := fn.(func(*float64, *uint64)); ok {
		return Caller(dyadicVoidPtrFloat64Uint64CallFn(to))
	}
	if to, ok := fn.(func(*float32, *uint64)); ok {
		return Caller(dyadicVoidPtrFloat32Uint64CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *uint64)); ok {
		return Caller(dyadicVoidPtrComplex128Uint64CallFn(to))
	}
	if to, ok := fn.(func(*float64, *uint32)); ok {
		return Caller(dyadicVoidPtrFloat64Uint32CallFn(to))
	}
	if to, ok := fn.(func(*float32, *uint32)); ok {
		return Caller(dyadicVoidPtrFloat32Uint32CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *uint32)); ok {
		return Caller(dyadicVoidPtrComplex128Uint32CallFn(to))
	}
	if to, ok := fn.(func(*float64, *uint8)); ok {
		return Caller(dyadicVoidPtrFloat64Uint8CallFn(to))
	}
	if to, ok := fn.(func(*float32, *uint8)); ok {
		return Caller(dyadicVoidPtrFloat32Uint8CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *uint8)); ok {
		return Caller(dyadicVoidPtrComplex128Uint8CallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr2x2(fn interface{}) Caller {
	if to, ok := fn.(func(*float64, *float64)); ok {
		return Caller(dyadicVoidPtrFloat64Float64CallFn(to))
	}
	if to, ok := fn.(func(*float32, *float64)); ok {
		return Caller(dyadicVoidPtrFloat32Float64CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *float64)); ok {
		return Caller(dyadicVoidPtrComplex128Float64CallFn(to))
	}
	if to, ok := fn.(func(*float64, *float32)); ok {
		return Caller(dyadicVoidPtrFloat64Float32CallFn(to))
	}
	if to, ok := fn.(func(*float32, *float32)); ok {
		return Caller(dyadicVoidPtrFloat32Float32CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *float32)); ok {
		return Caller(dyadicVoidPtrComplex128Float32CallFn(to))
	}
	if to, ok := fn.(func(*float64, *complex128)); ok {
		return Caller(dyadicVoidPtrFloat64Complex128CallFn(to))
	}
	if to, ok := fn.(func(*float32, *complex128)); ok {
		return Caller(dyadicVoidPtrFloat32Complex128CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *complex128)); ok {
		return Caller(dyadicVoidPtrComplex128Complex128CallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr2x3(fn interface{}) Caller {
	if to, ok := fn.(func(*float64, *string)); ok {
		return Caller(dyadicVoidPtrFloat64StringCallFn(to))
	}
	if to, ok := fn.(func(*float32, *string)); ok {
		return Caller(dyadicVoidPtrFloat32StringCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *string)); ok {
		return Caller(dyadicVoidPtrComplex128StringCallFn(to))
	}
	if to, ok := fn.(func(*float64, *[]byte)); ok {
		return Caller(dyadicVoidPtrFloat64BytesCallFn(to))
	}
	if to, ok := fn.(func(*float32, *[]byte)); ok {
		return Caller(dyadicVoidPtrFloat32BytesCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *[]byte)); ok {
		return Caller(dyadicVoidPtrComplex128BytesCallFn(to))
	}
	if to, ok := fn.(func(*float64, *bool)); ok {
		return Caller(dyadicVoidPtrFloat64BoolCallFn(to))
	}
	if to, ok := fn.(func(*float32, *bool)); ok {
		return Caller(dyadicVoidPtrFloat32BoolCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *bool)); ok {
		return Caller(dyadicVoidPtrComplex128BoolCallFn(to))
	}
	if to, ok := fn.(func(*float64, *time.Time)); ok {
		return Caller(dyadicVoidPtrFloat64TimeCallFn(to))
	}
	if to, ok := fn.(func(*float32, *time.Time)); ok {
		return Caller(dyadicVoidPtrFloat32TimeCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *time.Time)); ok {
		return Caller(dyadicVoidPtrComplex128TimeCallFn(to))
	}
	return nil
}

func lutMonadicVoidPtrIdx3(fn interface{}) Caller {
	if to, ok := fn.(func(*string)); ok {
		return Caller(monadicVoidPtrStringCallFn(to))
	}
	if to, ok := fn.(func(*[]byte)); ok {
		return Caller(monadicVoidPtrBytesCallFn(to))
	}
	if to, ok := fn.(func(*bool)); ok {
		return Caller(monadicVoidPtrBoolCallFn(to))
	}
	if to, ok := fn.(func(*time.Time)); ok {
		return Caller(monadicVoidPtrTimeCallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr3x0(fn interface{}) Caller {
	if to, ok := fn.(func(*string, *int)); ok {
		return Caller(dyadicVoidPtrStringIntCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *int)); ok {
		return Caller(dyadicVoidPtrBytesIntCallFn(to))
	}
	if to, ok := fn.(func(*bool, *int)); ok {
		return Caller(dyadicVoidPtrBoolIntCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *int)); ok {
		return Caller(dyadicVoidPtrTimeIntCallFn(to))
	}
	if to, ok := fn.(func(*string, *int64)); ok {
		return Caller(dyadicVoidPtrStringInt64CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *int64)); ok {
		return Caller(dyadicVoidPtrBytesInt64CallFn(to))
	}
	if to, ok := fn.(func(*bool, *int64)); ok {
		return Caller(dyadicVoidPtrBoolInt64CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *int64)); ok {
		return Caller(dyadicVoidPtrTimeInt64CallFn(to))
	}
	if to, ok := fn.(func(*string, *int32)); ok {
		return Caller(dyadicVoidPtrStringInt32CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *int32)); ok {
		return Caller(dyadicVoidPtrBytesInt32CallFn(to))
	}
	if to, ok := fn.(func(*bool, *int32)); ok {
		return Caller(dyadicVoidPtrBoolInt32CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *int32)); ok {
		return Caller(dyadicVoidPtrTimeInt32CallFn(to))
	}
	if to, ok := fn.(func(*string, *time.Duration)); ok {
		return Caller(dyadicVoidPtrStringDurationCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *time.Duration)); ok {
		return Caller(dyadicVoidPtrBytesDurationCallFn(to))
	}
	if to, ok := fn.(func(*bool, *time.Duration)); ok {
		return Caller(dyadicVoidPtrBoolDurationCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *time.Duration)); ok {
		return Caller(dyadicVoidPtrTimeDurationCallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr3x1(fn interface{}) Caller {
	if to, ok := fn.(func(*string, *uint)); ok {
		return Caller(dyadicVoidPtrStringUintCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *uint)); ok {
		return Caller(dyadicVoidPtrBytesUintCallFn(to))
	}
	if to, ok := fn.(func(*bool, *uint)); ok {
		return Caller(dyadicVoidPtrBoolUintCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *uint)); ok {
		return Caller(dyadicVoidPtrTimeUintCallFn(to))
	}
	if to, ok := fn.(func(*string, *uint64)); ok {
		return Caller(dyadicVoidPtrStringUint64CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *uint64)); ok {
		return Caller(dyadicVoidPtrBytesUint64CallFn(to))
	}
	if to, ok := fn.(func(*bool, *uint64)); ok {
		return Caller(dyadicVoidPtrBoolUint64CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *uint64)); ok {
		return Caller(dyadicVoidPtrTimeUint64CallFn(to))
	}
	if to, ok := fn.(func(*string, *uint32)); ok {
		return Caller(dyadicVoidPtrStringUint32CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *uint32)); ok {
		return Caller(dyadicVoidPtrBytesUint32CallFn(to))
	}
	if to, ok := fn.(func(*bool, *uint32)); ok {
		return Caller(dyadicVoidPtrBoolUint32CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *uint32)); ok {
		return Caller(dyadicVoidPtrTimeUint32CallFn(to))
	}
	if to, ok := fn.(func(*string, *uint8)); ok {
		return Caller(dyadicVoidPtrStringUint8CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *uint8)); ok {
		return Caller(dyadicVoidPtrBytesUint8CallFn(to))
	}
	if to, ok := fn.(func(*bool, *uint8)); ok {
		return Caller(dyadicVoidPtrBoolUint8CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *uint8)); ok {
		return Caller(dyadicVoidPtrTimeUint8CallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr3x2(fn interface{}) Caller {
	if to, ok := fn.(func(*string, *float64)); ok {
		return Caller(dyadicVoidPtrStringFloat64CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *float64)); ok {
		return Caller(dyadicVoidPtrBytesFloat64CallFn(to))
	}
	if to, ok := fn.(func(*bool, *float64)); ok {
		return Caller(dyadicVoidPtrBoolFloat64CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *float64)); ok {
		return Caller(dyadicVoidPtrTimeFloat64CallFn(to))
	}
	if to, ok := fn.(func(*string, *float32)); ok {
		return Caller(dyadicVoidPtrStringFloat32CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *float32)); ok {
		return Caller(dyadicVoidPtrBytesFloat32CallFn(to))
	}
	if to, ok := fn.(func(*bool, *float32)); ok {
		return Caller(dyadicVoidPtrBoolFloat32CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *float32)); ok {
		return Caller(dyadicVoidPtrTimeFloat32CallFn(to))
	}
	if to, ok := fn.(func(*string, *complex128)); ok {
		return Caller(dyadicVoidPtrStringComplex128CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *complex128)); ok {
		return Caller(dyadicVoidPtrBytesComplex128CallFn(to))
	}
	if to, ok := fn.(func(*bool, *complex128)); ok {
		return Caller(dyadicVoidPtrBoolComplex128CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *complex128)); ok {
		return Caller(dyadicVoidPtrTimeComplex128CallFn(to))
	}
	return nil
}

func lutDyadicVoidPtr3x3(fn interface{}) Caller {
	if to, ok := fn.(func(*string, *string)); ok {
		return Caller(dyadicVoidPtrStringStringCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *string)); ok {
		return Caller(dyadicVoidPtrBytesStringCallFn(to))
	}
	if to, ok := fn.(func(*bool, *string)); ok {
		return Caller(dyadicVoidPtrBoolStringCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *string)); ok {
		return Caller(dyadicVoidPtrTimeStringCallFn(to))
	}
	if to, ok := fn.(func(*string, *[]byte)); ok {
		return Caller(dyadicVoidPtrStringBytesCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *[]byte)); ok {
		return Caller(dyadicVoidPtrBytesBytesCallFn(to))
	}
	if to, ok := fn.(func(*bool, *[]byte)); ok {
		return Caller(dyadicVoidPtrBoolBytesCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *[]byte)); ok {
		return Caller(dyadicVoidPtrTimeBytesCallFn(to))
	}
	if to, ok := fn.(func(*string, *bool)); ok {
		return Caller(dyadicVoidPtrStringBoolCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *bool)); ok {
		return Caller(dyadicVoidPtrBytesBoolCallFn(to))
	}
	if to, ok := fn.(func(*bool, *bool)); ok {
		return Caller(dyadicVoidPtrBoolBoolCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *bool)); ok {
		return Caller(dyadicVoidPtrTimeBoolCallFn(to))
	}
	if to, ok := fn.(func(*string, *time.Time)); ok {
		return Caller(dyadicVoidPtrStringTimeCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *time.Time)); ok {
		return Caller(dyadicVoidPtrBytesTimeCallFn(to))
	}
	if to, ok := fn.(func(*bool, *time.Time)); ok {
		return Caller(dyadicVoidPtrBoolTimeCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *time.Time)); ok {
		return Caller(dyadicVoidPtrTimeTimeCallFn(to))
	}
	return nil
}

func lutMonadicErrorPtrIdx0(fn interface{}) Caller {
	if to, ok := fn.(func(*int) error); ok {
		return Caller(monadicErrorPtrIntCallFn(to))
	}
	if to, ok := fn.(func(*int64) error); ok {
		return Caller(monadicErrorPtrInt64CallFn(to))
	}
	if to, ok := fn.(func(*int32) error); ok {
		return Caller(monadicErrorPtrInt32CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration) error); ok {
		return Caller(monadicErrorPtrDurationCallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr0x0(fn interface{}) Caller {
	if to, ok := fn.(func(*int, *int) error); ok {
		return Caller(dyadicErrorPtrIntIntCallFn(to))
	}
	if to, ok := fn.(func(*int64, *int) error); ok {
		return Caller(dyadicErrorPtrInt64IntCallFn(to))
	}
	if to, ok := fn.(func(*int32, *int) error); ok {
		return Caller(dyadicErrorPtrInt32IntCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *int) error); ok {
		return Caller(dyadicErrorPtrDurationIntCallFn(to))
	}
	if to, ok := fn.(func(*int, *int64) error); ok {
		return Caller(dyadicErrorPtrIntInt64CallFn(to))
	}
	if to, ok := fn.(func(*int64, *int64) error); ok {
		return Caller(dyadicErrorPtrInt64Int64CallFn(to))
	}
	if to, ok := fn.(func(*int32, *int64) error); ok {
		return Caller(dyadicErrorPtrInt32Int64CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *int64) error); ok {
		return Caller(dyadicErrorPtrDurationInt64CallFn(to))
	}
	if to, ok := fn.(func(*int, *int32) error); ok {
		return Caller(dyadicErrorPtrIntInt32CallFn(to))
	}
	if to, ok := fn.(func(*int64, *int32) error); ok {
		return Caller(dyadicErrorPtrInt64Int32CallFn(to))
	}
	if to, ok := fn.(func(*int32, *int32) error); ok {
		return Caller(dyadicErrorPtrInt32Int32CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *int32) error); ok {
		return Caller(dyadicErrorPtrDurationInt32CallFn(to))
	}
	if to, ok := fn.(func(*int, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrIntDurationCallFn(to))
	}
	if to, ok := fn.(func(*int64, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrInt64DurationCallFn(to))
	}
	if to, ok := fn.(func(*int32, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrInt32DurationCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrDurationDurationCallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr0x1(fn interface{}) Caller {
	if to, ok := fn.(func(*int, *uint) error); ok {
		return Caller(dyadicErrorPtrIntUintCallFn(to))
	}
	if to, ok := fn.(func(*int64, *uint) error); ok {
		return Caller(dyadicErrorPtrInt64UintCallFn(to))
	}
	if to, ok := fn.(func(*int32, *uint) error); ok {
		return Caller(dyadicErrorPtrInt32UintCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *uint) error); ok {
		return Caller(dyadicErrorPtrDurationUintCallFn(to))
	}
	if to, ok := fn.(func(*int, *uint64) error); ok {
		return Caller(dyadicErrorPtrIntUint64CallFn(to))
	}
	if to, ok := fn.(func(*int64, *uint64) error); ok {
		return Caller(dyadicErrorPtrInt64Uint64CallFn(to))
	}
	if to, ok := fn.(func(*int32, *uint64) error); ok {
		return Caller(dyadicErrorPtrInt32Uint64CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *uint64) error); ok {
		return Caller(dyadicErrorPtrDurationUint64CallFn(to))
	}
	if to, ok := fn.(func(*int, *uint32) error); ok {
		return Caller(dyadicErrorPtrIntUint32CallFn(to))
	}
	if to, ok := fn.(func(*int64, *uint32) error); ok {
		return Caller(dyadicErrorPtrInt64Uint32CallFn(to))
	}
	if to, ok := fn.(func(*int32, *uint32) error); ok {
		return Caller(dyadicErrorPtrInt32Uint32CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *uint32) error); ok {
		return Caller(dyadicErrorPtrDurationUint32CallFn(to))
	}
	if to, ok := fn.(func(*int, *uint8) error); ok {
		return Caller(dyadicErrorPtrIntUint8CallFn(to))
	}
	if to, ok := fn.(func(*int64, *uint8) error); ok {
		return Caller(dyadicErrorPtrInt64Uint8CallFn(to))
	}
	if to, ok := fn.(func(*int32, *uint8) error); ok {
		return Caller(dyadicErrorPtrInt32Uint8CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *uint8) error); ok {
		return Caller(dyadicErrorPtrDurationUint8CallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr0x2(fn interface{}) Caller {
	if to, ok := fn.(func(*int, *float64) error); ok {
		return Caller(dyadicErrorPtrIntFloat64CallFn(to))
	}
	if to, ok := fn.(func(*int64, *float64) error); ok {
		return Caller(dyadicErrorPtrInt64Float64CallFn(to))
	}
	if to, ok := fn.(func(*int32, *float64) error); ok {
		return Caller(dyadicErrorPtrInt32Float64CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *float64) error); ok {
		return Caller(dyadicErrorPtrDurationFloat64CallFn(to))
	}
	if to, ok := fn.(func(*int, *float32) error); ok {
		return Caller(dyadicErrorPtrIntFloat32CallFn(to))
	}
	if to, ok := fn.(func(*int64, *float32) error); ok {
		return Caller(dyadicErrorPtrInt64Float32CallFn(to))
	}
	if to, ok := fn.(func(*int32, *float32) error); ok {
		return Caller(dyadicErrorPtrInt32Float32CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *float32) error); ok {
		return Caller(dyadicErrorPtrDurationFloat32CallFn(to))
	}
	if to, ok := fn.(func(*int, *complex128) error); ok {
		return Caller(dyadicErrorPtrIntComplex128CallFn(to))
	}
	if to, ok := fn.(func(*int64, *complex128) error); ok {
		return Caller(dyadicErrorPtrInt64Complex128CallFn(to))
	}
	if to, ok := fn.(func(*int32, *complex128) error); ok {
		return Caller(dyadicErrorPtrInt32Complex128CallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *complex128) error); ok {
		return Caller(dyadicErrorPtrDurationComplex128CallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr0x3(fn interface{}) Caller {
	if to, ok := fn.(func(*int, *string) error); ok {
		return Caller(dyadicErrorPtrIntStringCallFn(to))
	}
	if to, ok := fn.(func(*int64, *string) error); ok {
		return Caller(dyadicErrorPtrInt64StringCallFn(to))
	}
	if to, ok := fn.(func(*int32, *string) error); ok {
		return Caller(dyadicErrorPtrInt32StringCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *string) error); ok {
		return Caller(dyadicErrorPtrDurationStringCallFn(to))
	}
	if to, ok := fn.(func(*int, *[]byte) error); ok {
		return Caller(dyadicErrorPtrIntBytesCallFn(to))
	}
	if to, ok := fn.(func(*int64, *[]byte) error); ok {
		return Caller(dyadicErrorPtrInt64BytesCallFn(to))
	}
	if to, ok := fn.(func(*int32, *[]byte) error); ok {
		return Caller(dyadicErrorPtrInt32BytesCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *[]byte) error); ok {
		return Caller(dyadicErrorPtrDurationBytesCallFn(to))
	}
	if to, ok := fn.(func(*int, *bool) error); ok {
		return Caller(dyadicErrorPtrIntBoolCallFn(to))
	}
	if to, ok := fn.(func(*int64, *bool) error); ok {
		return Caller(dyadicErrorPtrInt64BoolCallFn(to))
	}
	if to, ok := fn.(func(*int32, *bool) error); ok {
		return Caller(dyadicErrorPtrInt32BoolCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *bool) error); ok {
		return Caller(dyadicErrorPtrDurationBoolCallFn(to))
	}
	if to, ok := fn.(func(*int, *time.Time) error); ok {
		return Caller(dyadicErrorPtrIntTimeCallFn(to))
	}
	if to, ok := fn.(func(*int64, *time.Time) error); ok {
		return Caller(dyadicErrorPtrInt64TimeCallFn(to))
	}
	if to, ok := fn.(func(*int32, *time.Time) error); ok {
		return Caller(dyadicErrorPtrInt32TimeCallFn(to))
	}
	if to, ok := fn.(func(*time.Duration, *time.Time) error); ok {
		return Caller(dyadicErrorPtrDurationTimeCallFn(to))
	}
	return nil
}

func lutMonadicErrorPtrIdx1(fn interface{}) Caller {
	if to, ok := fn.(func(*uint) error); ok {
		return Caller(monadicErrorPtrUintCallFn(to))
	}
	if to, ok := fn.(func(*uint64) error); ok {
		return Caller(monadicErrorPtrUint64CallFn(to))
	}
	if to, ok := fn.(func(*uint32) error); ok {
		return Caller(monadicErrorPtrUint32CallFn(to))
	}
	if to, ok := fn.(func(*uint8) error); ok {
		return Caller(monadicErrorPtrUint8CallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr1x0(fn interface{}) Caller {
	if to, ok := fn.(func(*uint, *int) error); ok {
		return Caller(dyadicErrorPtrUintIntCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *int) error); ok {
		return Caller(dyadicErrorPtrUint64IntCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *int) error); ok {
		return Caller(dyadicErrorPtrUint32IntCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *int) error); ok {
		return Caller(dyadicErrorPtrUint8IntCallFn(to))
	}
	if to, ok := fn.(func(*uint, *int64) error); ok {
		return Caller(dyadicErrorPtrUintInt64CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *int64) error); ok {
		return Caller(dyadicErrorPtrUint64Int64CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *int64) error); ok {
		return Caller(dyadicErrorPtrUint32Int64CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *int64) error); ok {
		return Caller(dyadicErrorPtrUint8Int64CallFn(to))
	}
	if to, ok := fn.(func(*uint, *int32) error); ok {
		return Caller(dyadicErrorPtrUintInt32CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *int32) error); ok {
		return Caller(dyadicErrorPtrUint64Int32CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *int32) error); ok {
		return Caller(dyadicErrorPtrUint32Int32CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *int32) error); ok {
		return Caller(dyadicErrorPtrUint8Int32CallFn(to))
	}
	if to, ok := fn.(func(*uint, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrUintDurationCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrUint64DurationCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrUint32DurationCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrUint8DurationCallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr1x1(fn interface{}) Caller {
	if to, ok := fn.(func(*uint, *uint) error); ok {
		return Caller(dyadicErrorPtrUintUintCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *uint) error); ok {
		return Caller(dyadicErrorPtrUint64UintCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *uint) error); ok {
		return Caller(dyadicErrorPtrUint32UintCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *uint) error); ok {
		return Caller(dyadicErrorPtrUint8UintCallFn(to))
	}
	if to, ok := fn.(func(*uint, *uint64) error); ok {
		return Caller(dyadicErrorPtrUintUint64CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *uint64) error); ok {
		return Caller(dyadicErrorPtrUint64Uint64CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *uint64) error); ok {
		return Caller(dyadicErrorPtrUint32Uint64CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *uint64) error); ok {
		return Caller(dyadicErrorPtrUint8Uint64CallFn(to))
	}
	if to, ok := fn.(func(*uint, *uint32) error); ok {
		return Caller(dyadicErrorPtrUintUint32CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *uint32) error); ok {
		return Caller(dyadicErrorPtrUint64Uint32CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *uint32) error); ok {
		return Caller(dyadicErrorPtrUint32Uint32CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *uint32) error); ok {
		return Caller(dyadicErrorPtrUint8Uint32CallFn(to))
	}
	if to, ok := fn.(func(*uint, *uint8) error); ok {
		return Caller(dyadicErrorPtrUintUint8CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *uint8) error); ok {
		return Caller(dyadicErrorPtrUint64Uint8CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *uint8) error); ok {
		return Caller(dyadicErrorPtrUint32Uint8CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *uint8) error); ok {
		return Caller(dyadicErrorPtrUint8Uint8CallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr1x2(fn interface{}) Caller {
	if to, ok := fn.(func(*uint, *float64) error); ok {
		return Caller(dyadicErrorPtrUintFloat64CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *float64) error); ok {
		return Caller(dyadicErrorPtrUint64Float64CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *float64) error); ok {
		return Caller(dyadicErrorPtrUint32Float64CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *float64) error); ok {
		return Caller(dyadicErrorPtrUint8Float64CallFn(to))
	}
	if to, ok := fn.(func(*uint, *float32) error); ok {
		return Caller(dyadicErrorPtrUintFloat32CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *float32) error); ok {
		return Caller(dyadicErrorPtrUint64Float32CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *float32) error); ok {
		return Caller(dyadicErrorPtrUint32Float32CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *float32) error); ok {
		return Caller(dyadicErrorPtrUint8Float32CallFn(to))
	}
	if to, ok := fn.(func(*uint, *complex128) error); ok {
		return Caller(dyadicErrorPtrUintComplex128CallFn(to))
	}
	if to, ok := fn.(func(*uint64, *complex128) error); ok {
		return Caller(dyadicErrorPtrUint64Complex128CallFn(to))
	}
	if to, ok := fn.(func(*uint32, *complex128) error); ok {
		return Caller(dyadicErrorPtrUint32Complex128CallFn(to))
	}
	if to, ok := fn.(func(*uint8, *complex128) error); ok {
		return Caller(dyadicErrorPtrUint8Complex128CallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr1x3(fn interface{}) Caller {
	if to, ok := fn.(func(*uint, *string) error); ok {
		return Caller(dyadicErrorPtrUintStringCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *string) error); ok {
		return Caller(dyadicErrorPtrUint64StringCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *string) error); ok {
		return Caller(dyadicErrorPtrUint32StringCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *string) error); ok {
		return Caller(dyadicErrorPtrUint8StringCallFn(to))
	}
	if to, ok := fn.(func(*uint, *[]byte) error); ok {
		return Caller(dyadicErrorPtrUintBytesCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *[]byte) error); ok {
		return Caller(dyadicErrorPtrUint64BytesCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *[]byte) error); ok {
		return Caller(dyadicErrorPtrUint32BytesCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *[]byte) error); ok {
		return Caller(dyadicErrorPtrUint8BytesCallFn(to))
	}
	if to, ok := fn.(func(*uint, *bool) error); ok {
		return Caller(dyadicErrorPtrUintBoolCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *bool) error); ok {
		return Caller(dyadicErrorPtrUint64BoolCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *bool) error); ok {
		return Caller(dyadicErrorPtrUint32BoolCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *bool) error); ok {
		return Caller(dyadicErrorPtrUint8BoolCallFn(to))
	}
	if to, ok := fn.(func(*uint, *time.Time) error); ok {
		return Caller(dyadicErrorPtrUintTimeCallFn(to))
	}
	if to, ok := fn.(func(*uint64, *time.Time) error); ok {
		return Caller(dyadicErrorPtrUint64TimeCallFn(to))
	}
	if to, ok := fn.(func(*uint32, *time.Time) error); ok {
		return Caller(dyadicErrorPtrUint32TimeCallFn(to))
	}
	if to, ok := fn.(func(*uint8, *time.Time) error); ok {
		return Caller(dyadicErrorPtrUint8TimeCallFn(to))
	}
	return nil
}

func lutMonadicErrorPtrIdx2(fn interface{}) Caller {
	if to, ok := fn.(func(*float64) error); ok {
		return Caller(monadicErrorPtrFloat64CallFn(to))
	}
	if to, ok := fn.(func(*float32) error); ok {
		return Caller(monadicErrorPtrFloat32CallFn(to))
	}
	if to, ok := fn.(func(*complex128) error); ok {
		return Caller(monadicErrorPtrComplex128CallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr2x0(fn interface{}) Caller {
	if to, ok := fn.(func(*float64, *int) error); ok {
		return Caller(dyadicErrorPtrFloat64IntCallFn(to))
	}
	if to, ok := fn.(func(*float32, *int) error); ok {
		return Caller(dyadicErrorPtrFloat32IntCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *int) error); ok {
		return Caller(dyadicErrorPtrComplex128IntCallFn(to))
	}
	if to, ok := fn.(func(*float64, *int64) error); ok {
		return Caller(dyadicErrorPtrFloat64Int64CallFn(to))
	}
	if to, ok := fn.(func(*float32, *int64) error); ok {
		return Caller(dyadicErrorPtrFloat32Int64CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *int64) error); ok {
		return Caller(dyadicErrorPtrComplex128Int64CallFn(to))
	}
	if to, ok := fn.(func(*float64, *int32) error); ok {
		return Caller(dyadicErrorPtrFloat64Int32CallFn(to))
	}
	if to, ok := fn.(func(*float32, *int32) error); ok {
		return Caller(dyadicErrorPtrFloat32Int32CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *int32) error); ok {
		return Caller(dyadicErrorPtrComplex128Int32CallFn(to))
	}
	if to, ok := fn.(func(*float64, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrFloat64DurationCallFn(to))
	}
	if to, ok := fn.(func(*float32, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrFloat32DurationCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrComplex128DurationCallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr2x1(fn interface{}) Caller {
	if to, ok := fn.(func(*float64, *uint) error); ok {
		return Caller(dyadicErrorPtrFloat64UintCallFn(to))
	}
	if to, ok := fn.(func(*float32, *uint) error); ok {
		return Caller(dyadicErrorPtrFloat32UintCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *uint) error); ok {
		return Caller(dyadicErrorPtrComplex128UintCallFn(to))
	}
	if to, ok := fn.(func(*float64, *uint64) error); ok {
		return Caller(dyadicErrorPtrFloat64Uint64CallFn(to))
	}
	if to, ok := fn.(func(*float32, *uint64) error); ok {
		return Caller(dyadicErrorPtrFloat32Uint64CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *uint64) error); ok {
		return Caller(dyadicErrorPtrComplex128Uint64CallFn(to))
	}
	if to, ok := fn.(func(*float64, *uint32) error); ok {
		return Caller(dyadicErrorPtrFloat64Uint32CallFn(to))
	}
	if to, ok := fn.(func(*float32, *uint32) error); ok {
		return Caller(dyadicErrorPtrFloat32Uint32CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *uint32) error); ok {
		return Caller(dyadicErrorPtrComplex128Uint32CallFn(to))
	}
	if to, ok := fn.(func(*float64, *uint8) error); ok {
		return Caller(dyadicErrorPtrFloat64Uint8CallFn(to))
	}
	if to, ok := fn.(func(*float32, *uint8) error); ok {
		return Caller(dyadicErrorPtrFloat32Uint8CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *uint8) error); ok {
		return Caller(dyadicErrorPtrComplex128Uint8CallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr2x2(fn interface{}) Caller {
	if to, ok := fn.(func(*float64, *float64) error); ok {
		return Caller(dyadicErrorPtrFloat64Float64CallFn(to))
	}
	if to, ok := fn.(func(*float32, *float64) error); ok {
		return Caller(dyadicErrorPtrFloat32Float64CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *float64) error); ok {
		return Caller(dyadicErrorPtrComplex128Float64CallFn(to))
	}
	if to, ok := fn.(func(*float64, *float32) error); ok {
		return Caller(dyadicErrorPtrFloat64Float32CallFn(to))
	}
	if to, ok := fn.(func(*float32, *float32) error); ok {
		return Caller(dyadicErrorPtrFloat32Float32CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *float32) error); ok {
		return Caller(dyadicErrorPtrComplex128Float32CallFn(to))
	}
	if to, ok := fn.(func(*float64, *complex128) error); ok {
		return Caller(dyadicErrorPtrFloat64Complex128CallFn(to))
	}
	if to, ok := fn.(func(*float32, *complex128) error); ok {
		return Caller(dyadicErrorPtrFloat32Complex128CallFn(to))
	}
	if to, ok := fn.(func(*complex128, *complex128) error); ok {
		return Caller(dyadicErrorPtrComplex128Complex128CallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr2x3(fn interface{}) Caller {
	if to, ok := fn.(func(*float64, *string) error); ok {
		return Caller(dyadicErrorPtrFloat64StringCallFn(to))
	}
	if to, ok := fn.(func(*float32, *string) error); ok {
		return Caller(dyadicErrorPtrFloat32StringCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *string) error); ok {
		return Caller(dyadicErrorPtrComplex128StringCallFn(to))
	}
	if to, ok := fn.(func(*float64, *[]byte) error); ok {
		return Caller(dyadicErrorPtrFloat64BytesCallFn(to))
	}
	if to, ok := fn.(func(*float32, *[]byte) error); ok {
		return Caller(dyadicErrorPtrFloat32BytesCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *[]byte) error); ok {
		return Caller(dyadicErrorPtrComplex128BytesCallFn(to))
	}
	if to, ok := fn.(func(*float64, *bool) error); ok {
		return Caller(dyadicErrorPtrFloat64BoolCallFn(to))
	}
	if to, ok := fn.(func(*float32, *bool) error); ok {
		return Caller(dyadicErrorPtrFloat32BoolCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *bool) error); ok {
		return Caller(dyadicErrorPtrComplex128BoolCallFn(to))
	}
	if to, ok := fn.(func(*float64, *time.Time) error); ok {
		return Caller(dyadicErrorPtrFloat64TimeCallFn(to))
	}
	if to, ok := fn.(func(*float32, *time.Time) error); ok {
		return Caller(dyadicErrorPtrFloat32TimeCallFn(to))
	}
	if to, ok := fn.(func(*complex128, *time.Time) error); ok {
		return Caller(dyadicErrorPtrComplex128TimeCallFn(to))
	}
	return nil
}

func lutMonadicErrorPtrIdx3(fn interface{}) Caller {
	if to, ok := fn.(func(*string) error); ok {
		return Caller(monadicErrorPtrStringCallFn(to))
	}
	if to, ok := fn.(func(*[]byte) error); ok {
		return Caller(monadicErrorPtrBytesCallFn(to))
	}
	if to, ok := fn.(func(*bool) error); ok {
		return Caller(monadicErrorPtrBoolCallFn(to))
	}
	if to, ok := fn.(func(*time.Time) error); ok {
		return Caller(monadicErrorPtrTimeCallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr3x0(fn interface{}) Caller {
	if to, ok := fn.(func(*string, *int) error); ok {
		return Caller(dyadicErrorPtrStringIntCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *int) error); ok {
		return Caller(dyadicErrorPtrBytesIntCallFn(to))
	}
	if to, ok := fn.(func(*bool, *int) error); ok {
		return Caller(dyadicErrorPtrBoolIntCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *int) error); ok {
		return Caller(dyadicErrorPtrTimeIntCallFn(to))
	}
	if to, ok := fn.(func(*string, *int64) error); ok {
		return Caller(dyadicErrorPtrStringInt64CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *int64) error); ok {
		return Caller(dyadicErrorPtrBytesInt64CallFn(to))
	}
	if to, ok := fn.(func(*bool, *int64) error); ok {
		return Caller(dyadicErrorPtrBoolInt64CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *int64) error); ok {
		return Caller(dyadicErrorPtrTimeInt64CallFn(to))
	}
	if to, ok := fn.(func(*string, *int32) error); ok {
		return Caller(dyadicErrorPtrStringInt32CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *int32) error); ok {
		return Caller(dyadicErrorPtrBytesInt32CallFn(to))
	}
	if to, ok := fn.(func(*bool, *int32) error); ok {
		return Caller(dyadicErrorPtrBoolInt32CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *int32) error); ok {
		return Caller(dyadicErrorPtrTimeInt32CallFn(to))
	}
	if to, ok := fn.(func(*string, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrStringDurationCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrBytesDurationCallFn(to))
	}
	if to, ok := fn.(func(*bool, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrBoolDurationCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *time.Duration) error); ok {
		return Caller(dyadicErrorPtrTimeDurationCallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr3x1(fn interface{}) Caller {
	if to, ok := fn.(func(*string, *uint) error); ok {
		return Caller(dyadicErrorPtrStringUintCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *uint) error); ok {
		return Caller(dyadicErrorPtrBytesUintCallFn(to))
	}
	if to, ok := fn.(func(*bool, *uint) error); ok {
		return Caller(dyadicErrorPtrBoolUintCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *uint) error); ok {
		return Caller(dyadicErrorPtrTimeUintCallFn(to))
	}
	if to, ok := fn.(func(*string, *uint64) error); ok {
		return Caller(dyadicErrorPtrStringUint64CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *uint64) error); ok {
		return Caller(dyadicErrorPtrBytesUint64CallFn(to))
	}
	if to, ok := fn.(func(*bool, *uint64) error); ok {
		return Caller(dyadicErrorPtrBoolUint64CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *uint64) error); ok {
		return Caller(dyadicErrorPtrTimeUint64CallFn(to))
	}
	if to, ok := fn.(func(*string, *uint32) error); ok {
		return Caller(dyadicErrorPtrStringUint32CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *uint32) error); ok {
		return Caller(dyadicErrorPtrBytesUint32CallFn(to))
	}
	if to, ok := fn.(func(*bool, *uint32) error); ok {
		return Caller(dyadicErrorPtrBoolUint32CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *uint32) error); ok {
		return Caller(dyadicErrorPtrTimeUint32CallFn(to))
	}
	if to, ok := fn.(func(*string, *uint8) error); ok {
		return Caller(dyadicErrorPtrStringUint8CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *uint8) error); ok {
		return Caller(dyadicErrorPtrBytesUint8CallFn(to))
	}
	if to, ok := fn.(func(*bool, *uint8) error); ok {
		return Caller(dyadicErrorPtrBoolUint8CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *uint8) error); ok {
		return Caller(dyadicErrorPtrTimeUint8CallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr3x2(fn interface{}) Caller {
	if to, ok := fn.(func(*string, *float64) error); ok {
		return Caller(dyadicErrorPtrStringFloat64CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *float64) error); ok {
		return Caller(dyadicErrorPtrBytesFloat64CallFn(to))
	}
	if to, ok := fn.(func(*bool, *float64) error); ok {
		return Caller(dyadicErrorPtrBoolFloat64CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *float64) error); ok {
		return Caller(dyadicErrorPtrTimeFloat64CallFn(to))
	}
	if to, ok := fn.(func(*string, *float32) error); ok {
		return Caller(dyadicErrorPtrStringFloat32CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *float32) error); ok {
		return Caller(dyadicErrorPtrBytesFloat32CallFn(to))
	}
	if to, ok := fn.(func(*bool, *float32) error); ok {
		return Caller(dyadicErrorPtrBoolFloat32CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *float32) error); ok {
		return Caller(dyadicErrorPtrTimeFloat32CallFn(to))
	}
	if to, ok := fn.(func(*string, *complex128) error); ok {
		return Caller(dyadicErrorPtrStringComplex128CallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *complex128) error); ok {
		return Caller(dyadicErrorPtrBytesComplex128CallFn(to))
	}
	if to, ok := fn.(func(*bool, *complex128) error); ok {
		return Caller(dyadicErrorPtrBoolComplex128CallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *complex128) error); ok {
		return Caller(dyadicErrorPtrTimeComplex128CallFn(to))
	}
	return nil
}

func lutDyadicErrorPtr3x3(fn interface{}) Caller {
	if to, ok := fn.(func(*string, *string) error); ok {
		return Caller(dyadicErrorPtrStringStringCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *string) error); ok {
		return Caller(dyadicErrorPtrBytesStringCallFn(to))
	}
	if to, ok := fn.(func(*bool, *string) error); ok {
		return Caller(dyadicErrorPtrBoolStringCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *string) error); ok {
		return Caller(dyadicErrorPtrTimeStringCallFn(to))
	}
	if to, ok := fn.(func(*string, *[]byte) error); ok {
		return Caller(dyadicErrorPtrStringBytesCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *[]byte) error); ok {
		return Caller(dyadicErrorPtrBytesBytesCallFn(to))
	}
	if to, ok := fn.(func(*bool, *[]byte) error); ok {
		return Caller(dyadicErrorPtrBoolBytesCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *[]byte) error); ok {
		return Caller(dyadicErrorPtrTimeBytesCallFn(to))
	}
	if to, ok := fn.(func(*string, *bool) error); ok {
		return Caller(dyadicErrorPtrStringBoolCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *bool) error); ok {
		return Caller(dyadicErrorPtrBytesBoolCallFn(to))
	}
	if to, ok := fn.(func(*bool, *bool) error); ok {
		return Caller(dyadicErrorPtrBoolBoolCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *bool) error); ok {
		return Caller(dyadicErrorPtrTimeBoolCallFn(to))
	}
	if to, ok := fn.(func(*string, *time.Time) error); ok {
		return Caller(dyadicErrorPtrStringTimeCallFn(to))
	}
	if to, ok := fn.(func(*[]byte, *time.Time) error); ok {
		return Caller(dyadicErrorPtrBytesTimeCallFn(to))
	}
	if to, ok := fn.(func(*bool, *time.Time) error); ok {
		return Caller(dyadicErrorPtrBoolTimeCallFn(to))
	}
	if to, ok := fn.(func(*time.Time, *time.Time) error); ok {
		return Caller(dyadicErrorPtrTimeTimeCallFn(to))
	}
	return nil
}
