package call

import (
	"time"
)

type dyadicVoidIntIntCallFn func(int, int)

func (fn dyadicVoidIntIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntInt64CallFn func(int, int64)

func (fn dyadicVoidIntInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntInt32CallFn func(int, int32)

func (fn dyadicVoidIntInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntDurationCallFn func(int, time.Duration)

func (fn dyadicVoidIntDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntUintCallFn func(int, uint)

func (fn dyadicVoidIntUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntUint64CallFn func(int, uint64)

func (fn dyadicVoidIntUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntUint32CallFn func(int, uint32)

func (fn dyadicVoidIntUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntUint8CallFn func(int, uint8)

func (fn dyadicVoidIntUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntFloat64CallFn func(int, float64)

func (fn dyadicVoidIntFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntFloat32CallFn func(int, float32)

func (fn dyadicVoidIntFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntComplex128CallFn func(int, complex128)

func (fn dyadicVoidIntComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntStringCallFn func(int, string)

func (fn dyadicVoidIntStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntBytesCallFn func(int, []byte)

func (fn dyadicVoidIntBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntBoolCallFn func(int, bool)

func (fn dyadicVoidIntBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidIntTimeCallFn func(int, time.Time)

func (fn dyadicVoidIntTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64IntCallFn func(int64, int)

func (fn dyadicVoidInt64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64Int64CallFn func(int64, int64)

func (fn dyadicVoidInt64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64Int32CallFn func(int64, int32)

func (fn dyadicVoidInt64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64DurationCallFn func(int64, time.Duration)

func (fn dyadicVoidInt64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64UintCallFn func(int64, uint)

func (fn dyadicVoidInt64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64Uint64CallFn func(int64, uint64)

func (fn dyadicVoidInt64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64Uint32CallFn func(int64, uint32)

func (fn dyadicVoidInt64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64Uint8CallFn func(int64, uint8)

func (fn dyadicVoidInt64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64Float64CallFn func(int64, float64)

func (fn dyadicVoidInt64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64Float32CallFn func(int64, float32)

func (fn dyadicVoidInt64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64Complex128CallFn func(int64, complex128)

func (fn dyadicVoidInt64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64StringCallFn func(int64, string)

func (fn dyadicVoidInt64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64BytesCallFn func(int64, []byte)

func (fn dyadicVoidInt64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64BoolCallFn func(int64, bool)

func (fn dyadicVoidInt64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt64TimeCallFn func(int64, time.Time)

func (fn dyadicVoidInt64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32IntCallFn func(int32, int)

func (fn dyadicVoidInt32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32Int64CallFn func(int32, int64)

func (fn dyadicVoidInt32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32Int32CallFn func(int32, int32)

func (fn dyadicVoidInt32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32DurationCallFn func(int32, time.Duration)

func (fn dyadicVoidInt32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32UintCallFn func(int32, uint)

func (fn dyadicVoidInt32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32Uint64CallFn func(int32, uint64)

func (fn dyadicVoidInt32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32Uint32CallFn func(int32, uint32)

func (fn dyadicVoidInt32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32Uint8CallFn func(int32, uint8)

func (fn dyadicVoidInt32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32Float64CallFn func(int32, float64)

func (fn dyadicVoidInt32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32Float32CallFn func(int32, float32)

func (fn dyadicVoidInt32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32Complex128CallFn func(int32, complex128)

func (fn dyadicVoidInt32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32StringCallFn func(int32, string)

func (fn dyadicVoidInt32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32BytesCallFn func(int32, []byte)

func (fn dyadicVoidInt32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32BoolCallFn func(int32, bool)

func (fn dyadicVoidInt32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidInt32TimeCallFn func(int32, time.Time)

func (fn dyadicVoidInt32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationIntCallFn func(time.Duration, int)

func (fn dyadicVoidDurationIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationInt64CallFn func(time.Duration, int64)

func (fn dyadicVoidDurationInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationInt32CallFn func(time.Duration, int32)

func (fn dyadicVoidDurationInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationDurationCallFn func(time.Duration, time.Duration)

func (fn dyadicVoidDurationDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationUintCallFn func(time.Duration, uint)

func (fn dyadicVoidDurationUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationUint64CallFn func(time.Duration, uint64)

func (fn dyadicVoidDurationUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationUint32CallFn func(time.Duration, uint32)

func (fn dyadicVoidDurationUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationUint8CallFn func(time.Duration, uint8)

func (fn dyadicVoidDurationUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationFloat64CallFn func(time.Duration, float64)

func (fn dyadicVoidDurationFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationFloat32CallFn func(time.Duration, float32)

func (fn dyadicVoidDurationFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationComplex128CallFn func(time.Duration, complex128)

func (fn dyadicVoidDurationComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationStringCallFn func(time.Duration, string)

func (fn dyadicVoidDurationStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationBytesCallFn func(time.Duration, []byte)

func (fn dyadicVoidDurationBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationBoolCallFn func(time.Duration, bool)

func (fn dyadicVoidDurationBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidDurationTimeCallFn func(time.Duration, time.Time)

func (fn dyadicVoidDurationTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintIntCallFn func(uint, int)

func (fn dyadicVoidUintIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintInt64CallFn func(uint, int64)

func (fn dyadicVoidUintInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintInt32CallFn func(uint, int32)

func (fn dyadicVoidUintInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintDurationCallFn func(uint, time.Duration)

func (fn dyadicVoidUintDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintUintCallFn func(uint, uint)

func (fn dyadicVoidUintUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintUint64CallFn func(uint, uint64)

func (fn dyadicVoidUintUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintUint32CallFn func(uint, uint32)

func (fn dyadicVoidUintUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintUint8CallFn func(uint, uint8)

func (fn dyadicVoidUintUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintFloat64CallFn func(uint, float64)

func (fn dyadicVoidUintFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintFloat32CallFn func(uint, float32)

func (fn dyadicVoidUintFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintComplex128CallFn func(uint, complex128)

func (fn dyadicVoidUintComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintStringCallFn func(uint, string)

func (fn dyadicVoidUintStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintBytesCallFn func(uint, []byte)

func (fn dyadicVoidUintBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintBoolCallFn func(uint, bool)

func (fn dyadicVoidUintBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUintTimeCallFn func(uint, time.Time)

func (fn dyadicVoidUintTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64IntCallFn func(uint64, int)

func (fn dyadicVoidUint64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64Int64CallFn func(uint64, int64)

func (fn dyadicVoidUint64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64Int32CallFn func(uint64, int32)

func (fn dyadicVoidUint64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64DurationCallFn func(uint64, time.Duration)

func (fn dyadicVoidUint64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64UintCallFn func(uint64, uint)

func (fn dyadicVoidUint64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64Uint64CallFn func(uint64, uint64)

func (fn dyadicVoidUint64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64Uint32CallFn func(uint64, uint32)

func (fn dyadicVoidUint64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64Uint8CallFn func(uint64, uint8)

func (fn dyadicVoidUint64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64Float64CallFn func(uint64, float64)

func (fn dyadicVoidUint64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64Float32CallFn func(uint64, float32)

func (fn dyadicVoidUint64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64Complex128CallFn func(uint64, complex128)

func (fn dyadicVoidUint64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64StringCallFn func(uint64, string)

func (fn dyadicVoidUint64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64BytesCallFn func(uint64, []byte)

func (fn dyadicVoidUint64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64BoolCallFn func(uint64, bool)

func (fn dyadicVoidUint64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint64TimeCallFn func(uint64, time.Time)

func (fn dyadicVoidUint64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32IntCallFn func(uint32, int)

func (fn dyadicVoidUint32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32Int64CallFn func(uint32, int64)

func (fn dyadicVoidUint32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32Int32CallFn func(uint32, int32)

func (fn dyadicVoidUint32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32DurationCallFn func(uint32, time.Duration)

func (fn dyadicVoidUint32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32UintCallFn func(uint32, uint)

func (fn dyadicVoidUint32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32Uint64CallFn func(uint32, uint64)

func (fn dyadicVoidUint32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32Uint32CallFn func(uint32, uint32)

func (fn dyadicVoidUint32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32Uint8CallFn func(uint32, uint8)

func (fn dyadicVoidUint32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32Float64CallFn func(uint32, float64)

func (fn dyadicVoidUint32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32Float32CallFn func(uint32, float32)

func (fn dyadicVoidUint32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32Complex128CallFn func(uint32, complex128)

func (fn dyadicVoidUint32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32StringCallFn func(uint32, string)

func (fn dyadicVoidUint32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32BytesCallFn func(uint32, []byte)

func (fn dyadicVoidUint32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32BoolCallFn func(uint32, bool)

func (fn dyadicVoidUint32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint32TimeCallFn func(uint32, time.Time)

func (fn dyadicVoidUint32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8IntCallFn func(uint8, int)

func (fn dyadicVoidUint8IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8Int64CallFn func(uint8, int64)

func (fn dyadicVoidUint8Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8Int32CallFn func(uint8, int32)

func (fn dyadicVoidUint8Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8DurationCallFn func(uint8, time.Duration)

func (fn dyadicVoidUint8DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8UintCallFn func(uint8, uint)

func (fn dyadicVoidUint8UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8Uint64CallFn func(uint8, uint64)

func (fn dyadicVoidUint8Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8Uint32CallFn func(uint8, uint32)

func (fn dyadicVoidUint8Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8Uint8CallFn func(uint8, uint8)

func (fn dyadicVoidUint8Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8Float64CallFn func(uint8, float64)

func (fn dyadicVoidUint8Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8Float32CallFn func(uint8, float32)

func (fn dyadicVoidUint8Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8Complex128CallFn func(uint8, complex128)

func (fn dyadicVoidUint8Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8StringCallFn func(uint8, string)

func (fn dyadicVoidUint8StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8BytesCallFn func(uint8, []byte)

func (fn dyadicVoidUint8BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8BoolCallFn func(uint8, bool)

func (fn dyadicVoidUint8BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidUint8TimeCallFn func(uint8, time.Time)

func (fn dyadicVoidUint8TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64IntCallFn func(float64, int)

func (fn dyadicVoidFloat64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64Int64CallFn func(float64, int64)

func (fn dyadicVoidFloat64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64Int32CallFn func(float64, int32)

func (fn dyadicVoidFloat64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64DurationCallFn func(float64, time.Duration)

func (fn dyadicVoidFloat64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64UintCallFn func(float64, uint)

func (fn dyadicVoidFloat64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64Uint64CallFn func(float64, uint64)

func (fn dyadicVoidFloat64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64Uint32CallFn func(float64, uint32)

func (fn dyadicVoidFloat64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64Uint8CallFn func(float64, uint8)

func (fn dyadicVoidFloat64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64Float64CallFn func(float64, float64)

func (fn dyadicVoidFloat64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64Float32CallFn func(float64, float32)

func (fn dyadicVoidFloat64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64Complex128CallFn func(float64, complex128)

func (fn dyadicVoidFloat64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64StringCallFn func(float64, string)

func (fn dyadicVoidFloat64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64BytesCallFn func(float64, []byte)

func (fn dyadicVoidFloat64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64BoolCallFn func(float64, bool)

func (fn dyadicVoidFloat64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat64TimeCallFn func(float64, time.Time)

func (fn dyadicVoidFloat64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32IntCallFn func(float32, int)

func (fn dyadicVoidFloat32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32Int64CallFn func(float32, int64)

func (fn dyadicVoidFloat32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32Int32CallFn func(float32, int32)

func (fn dyadicVoidFloat32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32DurationCallFn func(float32, time.Duration)

func (fn dyadicVoidFloat32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32UintCallFn func(float32, uint)

func (fn dyadicVoidFloat32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32Uint64CallFn func(float32, uint64)

func (fn dyadicVoidFloat32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32Uint32CallFn func(float32, uint32)

func (fn dyadicVoidFloat32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32Uint8CallFn func(float32, uint8)

func (fn dyadicVoidFloat32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32Float64CallFn func(float32, float64)

func (fn dyadicVoidFloat32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32Float32CallFn func(float32, float32)

func (fn dyadicVoidFloat32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32Complex128CallFn func(float32, complex128)

func (fn dyadicVoidFloat32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32StringCallFn func(float32, string)

func (fn dyadicVoidFloat32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32BytesCallFn func(float32, []byte)

func (fn dyadicVoidFloat32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32BoolCallFn func(float32, bool)

func (fn dyadicVoidFloat32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidFloat32TimeCallFn func(float32, time.Time)

func (fn dyadicVoidFloat32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128IntCallFn func(complex128, int)

func (fn dyadicVoidComplex128IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128Int64CallFn func(complex128, int64)

func (fn dyadicVoidComplex128Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128Int32CallFn func(complex128, int32)

func (fn dyadicVoidComplex128Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128DurationCallFn func(complex128, time.Duration)

func (fn dyadicVoidComplex128DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128UintCallFn func(complex128, uint)

func (fn dyadicVoidComplex128UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128Uint64CallFn func(complex128, uint64)

func (fn dyadicVoidComplex128Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128Uint32CallFn func(complex128, uint32)

func (fn dyadicVoidComplex128Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128Uint8CallFn func(complex128, uint8)

func (fn dyadicVoidComplex128Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128Float64CallFn func(complex128, float64)

func (fn dyadicVoidComplex128Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128Float32CallFn func(complex128, float32)

func (fn dyadicVoidComplex128Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128Complex128CallFn func(complex128, complex128)

func (fn dyadicVoidComplex128Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128StringCallFn func(complex128, string)

func (fn dyadicVoidComplex128StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128BytesCallFn func(complex128, []byte)

func (fn dyadicVoidComplex128BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128BoolCallFn func(complex128, bool)

func (fn dyadicVoidComplex128BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidComplex128TimeCallFn func(complex128, time.Time)

func (fn dyadicVoidComplex128TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringIntCallFn func(string, int)

func (fn dyadicVoidStringIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringInt64CallFn func(string, int64)

func (fn dyadicVoidStringInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringInt32CallFn func(string, int32)

func (fn dyadicVoidStringInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringDurationCallFn func(string, time.Duration)

func (fn dyadicVoidStringDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringUintCallFn func(string, uint)

func (fn dyadicVoidStringUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringUint64CallFn func(string, uint64)

func (fn dyadicVoidStringUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringUint32CallFn func(string, uint32)

func (fn dyadicVoidStringUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringUint8CallFn func(string, uint8)

func (fn dyadicVoidStringUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringFloat64CallFn func(string, float64)

func (fn dyadicVoidStringFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringFloat32CallFn func(string, float32)

func (fn dyadicVoidStringFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringComplex128CallFn func(string, complex128)

func (fn dyadicVoidStringComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringStringCallFn func(string, string)

func (fn dyadicVoidStringStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringBytesCallFn func(string, []byte)

func (fn dyadicVoidStringBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringBoolCallFn func(string, bool)

func (fn dyadicVoidStringBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidStringTimeCallFn func(string, time.Time)

func (fn dyadicVoidStringTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesIntCallFn func([]byte, int)

func (fn dyadicVoidBytesIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesInt64CallFn func([]byte, int64)

func (fn dyadicVoidBytesInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesInt32CallFn func([]byte, int32)

func (fn dyadicVoidBytesInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesDurationCallFn func([]byte, time.Duration)

func (fn dyadicVoidBytesDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesUintCallFn func([]byte, uint)

func (fn dyadicVoidBytesUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesUint64CallFn func([]byte, uint64)

func (fn dyadicVoidBytesUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesUint32CallFn func([]byte, uint32)

func (fn dyadicVoidBytesUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesUint8CallFn func([]byte, uint8)

func (fn dyadicVoidBytesUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesFloat64CallFn func([]byte, float64)

func (fn dyadicVoidBytesFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesFloat32CallFn func([]byte, float32)

func (fn dyadicVoidBytesFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesComplex128CallFn func([]byte, complex128)

func (fn dyadicVoidBytesComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesStringCallFn func([]byte, string)

func (fn dyadicVoidBytesStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesBytesCallFn func([]byte, []byte)

func (fn dyadicVoidBytesBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesBoolCallFn func([]byte, bool)

func (fn dyadicVoidBytesBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBytesTimeCallFn func([]byte, time.Time)

func (fn dyadicVoidBytesTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolIntCallFn func(bool, int)

func (fn dyadicVoidBoolIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolInt64CallFn func(bool, int64)

func (fn dyadicVoidBoolInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolInt32CallFn func(bool, int32)

func (fn dyadicVoidBoolInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolDurationCallFn func(bool, time.Duration)

func (fn dyadicVoidBoolDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolUintCallFn func(bool, uint)

func (fn dyadicVoidBoolUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolUint64CallFn func(bool, uint64)

func (fn dyadicVoidBoolUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolUint32CallFn func(bool, uint32)

func (fn dyadicVoidBoolUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolUint8CallFn func(bool, uint8)

func (fn dyadicVoidBoolUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolFloat64CallFn func(bool, float64)

func (fn dyadicVoidBoolFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolFloat32CallFn func(bool, float32)

func (fn dyadicVoidBoolFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolComplex128CallFn func(bool, complex128)

func (fn dyadicVoidBoolComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolStringCallFn func(bool, string)

func (fn dyadicVoidBoolStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolBytesCallFn func(bool, []byte)

func (fn dyadicVoidBoolBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolBoolCallFn func(bool, bool)

func (fn dyadicVoidBoolBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidBoolTimeCallFn func(bool, time.Time)

func (fn dyadicVoidBoolTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeIntCallFn func(time.Time, int)

func (fn dyadicVoidTimeIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeInt64CallFn func(time.Time, int64)

func (fn dyadicVoidTimeInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeInt32CallFn func(time.Time, int32)

func (fn dyadicVoidTimeInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeDurationCallFn func(time.Time, time.Duration)

func (fn dyadicVoidTimeDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeUintCallFn func(time.Time, uint)

func (fn dyadicVoidTimeUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeUint64CallFn func(time.Time, uint64)

func (fn dyadicVoidTimeUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeUint32CallFn func(time.Time, uint32)

func (fn dyadicVoidTimeUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeUint8CallFn func(time.Time, uint8)

func (fn dyadicVoidTimeUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeFloat64CallFn func(time.Time, float64)

func (fn dyadicVoidTimeFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeFloat32CallFn func(time.Time, float32)

func (fn dyadicVoidTimeFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeComplex128CallFn func(time.Time, complex128)

func (fn dyadicVoidTimeComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeStringCallFn func(time.Time, string)

func (fn dyadicVoidTimeStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeBytesCallFn func(time.Time, []byte)

func (fn dyadicVoidTimeBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeBoolCallFn func(time.Time, bool)

func (fn dyadicVoidTimeBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidTimeTimeCallFn func(time.Time, time.Time)

func (fn dyadicVoidTimeTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicErrorIntIntCallFn func(int, int) error

func (fn dyadicErrorIntIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntInt64CallFn func(int, int64) error

func (fn dyadicErrorIntInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntInt32CallFn func(int, int32) error

func (fn dyadicErrorIntInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntDurationCallFn func(int, time.Duration) error

func (fn dyadicErrorIntDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntUintCallFn func(int, uint) error

func (fn dyadicErrorIntUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntUint64CallFn func(int, uint64) error

func (fn dyadicErrorIntUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntUint32CallFn func(int, uint32) error

func (fn dyadicErrorIntUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntUint8CallFn func(int, uint8) error

func (fn dyadicErrorIntUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntFloat64CallFn func(int, float64) error

func (fn dyadicErrorIntFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntFloat32CallFn func(int, float32) error

func (fn dyadicErrorIntFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntComplex128CallFn func(int, complex128) error

func (fn dyadicErrorIntComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntStringCallFn func(int, string) error

func (fn dyadicErrorIntStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntBytesCallFn func(int, []byte) error

func (fn dyadicErrorIntBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntBoolCallFn func(int, bool) error

func (fn dyadicErrorIntBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorIntTimeCallFn func(int, time.Time) error

func (fn dyadicErrorIntTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64IntCallFn func(int64, int) error

func (fn dyadicErrorInt64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64Int64CallFn func(int64, int64) error

func (fn dyadicErrorInt64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64Int32CallFn func(int64, int32) error

func (fn dyadicErrorInt64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64DurationCallFn func(int64, time.Duration) error

func (fn dyadicErrorInt64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64UintCallFn func(int64, uint) error

func (fn dyadicErrorInt64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64Uint64CallFn func(int64, uint64) error

func (fn dyadicErrorInt64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64Uint32CallFn func(int64, uint32) error

func (fn dyadicErrorInt64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64Uint8CallFn func(int64, uint8) error

func (fn dyadicErrorInt64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64Float64CallFn func(int64, float64) error

func (fn dyadicErrorInt64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64Float32CallFn func(int64, float32) error

func (fn dyadicErrorInt64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64Complex128CallFn func(int64, complex128) error

func (fn dyadicErrorInt64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64StringCallFn func(int64, string) error

func (fn dyadicErrorInt64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64BytesCallFn func(int64, []byte) error

func (fn dyadicErrorInt64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64BoolCallFn func(int64, bool) error

func (fn dyadicErrorInt64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt64TimeCallFn func(int64, time.Time) error

func (fn dyadicErrorInt64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32IntCallFn func(int32, int) error

func (fn dyadicErrorInt32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32Int64CallFn func(int32, int64) error

func (fn dyadicErrorInt32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32Int32CallFn func(int32, int32) error

func (fn dyadicErrorInt32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32DurationCallFn func(int32, time.Duration) error

func (fn dyadicErrorInt32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32UintCallFn func(int32, uint) error

func (fn dyadicErrorInt32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32Uint64CallFn func(int32, uint64) error

func (fn dyadicErrorInt32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32Uint32CallFn func(int32, uint32) error

func (fn dyadicErrorInt32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32Uint8CallFn func(int32, uint8) error

func (fn dyadicErrorInt32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32Float64CallFn func(int32, float64) error

func (fn dyadicErrorInt32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32Float32CallFn func(int32, float32) error

func (fn dyadicErrorInt32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32Complex128CallFn func(int32, complex128) error

func (fn dyadicErrorInt32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32StringCallFn func(int32, string) error

func (fn dyadicErrorInt32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32BytesCallFn func(int32, []byte) error

func (fn dyadicErrorInt32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32BoolCallFn func(int32, bool) error

func (fn dyadicErrorInt32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorInt32TimeCallFn func(int32, time.Time) error

func (fn dyadicErrorInt32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationIntCallFn func(time.Duration, int) error

func (fn dyadicErrorDurationIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationInt64CallFn func(time.Duration, int64) error

func (fn dyadicErrorDurationInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationInt32CallFn func(time.Duration, int32) error

func (fn dyadicErrorDurationInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationDurationCallFn func(time.Duration, time.Duration) error

func (fn dyadicErrorDurationDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationUintCallFn func(time.Duration, uint) error

func (fn dyadicErrorDurationUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationUint64CallFn func(time.Duration, uint64) error

func (fn dyadicErrorDurationUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationUint32CallFn func(time.Duration, uint32) error

func (fn dyadicErrorDurationUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationUint8CallFn func(time.Duration, uint8) error

func (fn dyadicErrorDurationUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationFloat64CallFn func(time.Duration, float64) error

func (fn dyadicErrorDurationFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationFloat32CallFn func(time.Duration, float32) error

func (fn dyadicErrorDurationFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationComplex128CallFn func(time.Duration, complex128) error

func (fn dyadicErrorDurationComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationStringCallFn func(time.Duration, string) error

func (fn dyadicErrorDurationStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationBytesCallFn func(time.Duration, []byte) error

func (fn dyadicErrorDurationBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationBoolCallFn func(time.Duration, bool) error

func (fn dyadicErrorDurationBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorDurationTimeCallFn func(time.Duration, time.Time) error

func (fn dyadicErrorDurationTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintIntCallFn func(uint, int) error

func (fn dyadicErrorUintIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintInt64CallFn func(uint, int64) error

func (fn dyadicErrorUintInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintInt32CallFn func(uint, int32) error

func (fn dyadicErrorUintInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintDurationCallFn func(uint, time.Duration) error

func (fn dyadicErrorUintDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintUintCallFn func(uint, uint) error

func (fn dyadicErrorUintUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintUint64CallFn func(uint, uint64) error

func (fn dyadicErrorUintUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintUint32CallFn func(uint, uint32) error

func (fn dyadicErrorUintUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintUint8CallFn func(uint, uint8) error

func (fn dyadicErrorUintUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintFloat64CallFn func(uint, float64) error

func (fn dyadicErrorUintFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintFloat32CallFn func(uint, float32) error

func (fn dyadicErrorUintFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintComplex128CallFn func(uint, complex128) error

func (fn dyadicErrorUintComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintStringCallFn func(uint, string) error

func (fn dyadicErrorUintStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintBytesCallFn func(uint, []byte) error

func (fn dyadicErrorUintBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintBoolCallFn func(uint, bool) error

func (fn dyadicErrorUintBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUintTimeCallFn func(uint, time.Time) error

func (fn dyadicErrorUintTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64IntCallFn func(uint64, int) error

func (fn dyadicErrorUint64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64Int64CallFn func(uint64, int64) error

func (fn dyadicErrorUint64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64Int32CallFn func(uint64, int32) error

func (fn dyadicErrorUint64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64DurationCallFn func(uint64, time.Duration) error

func (fn dyadicErrorUint64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64UintCallFn func(uint64, uint) error

func (fn dyadicErrorUint64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64Uint64CallFn func(uint64, uint64) error

func (fn dyadicErrorUint64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64Uint32CallFn func(uint64, uint32) error

func (fn dyadicErrorUint64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64Uint8CallFn func(uint64, uint8) error

func (fn dyadicErrorUint64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64Float64CallFn func(uint64, float64) error

func (fn dyadicErrorUint64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64Float32CallFn func(uint64, float32) error

func (fn dyadicErrorUint64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64Complex128CallFn func(uint64, complex128) error

func (fn dyadicErrorUint64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64StringCallFn func(uint64, string) error

func (fn dyadicErrorUint64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64BytesCallFn func(uint64, []byte) error

func (fn dyadicErrorUint64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64BoolCallFn func(uint64, bool) error

func (fn dyadicErrorUint64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint64TimeCallFn func(uint64, time.Time) error

func (fn dyadicErrorUint64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32IntCallFn func(uint32, int) error

func (fn dyadicErrorUint32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32Int64CallFn func(uint32, int64) error

func (fn dyadicErrorUint32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32Int32CallFn func(uint32, int32) error

func (fn dyadicErrorUint32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32DurationCallFn func(uint32, time.Duration) error

func (fn dyadicErrorUint32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32UintCallFn func(uint32, uint) error

func (fn dyadicErrorUint32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32Uint64CallFn func(uint32, uint64) error

func (fn dyadicErrorUint32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32Uint32CallFn func(uint32, uint32) error

func (fn dyadicErrorUint32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32Uint8CallFn func(uint32, uint8) error

func (fn dyadicErrorUint32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32Float64CallFn func(uint32, float64) error

func (fn dyadicErrorUint32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32Float32CallFn func(uint32, float32) error

func (fn dyadicErrorUint32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32Complex128CallFn func(uint32, complex128) error

func (fn dyadicErrorUint32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32StringCallFn func(uint32, string) error

func (fn dyadicErrorUint32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32BytesCallFn func(uint32, []byte) error

func (fn dyadicErrorUint32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32BoolCallFn func(uint32, bool) error

func (fn dyadicErrorUint32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint32TimeCallFn func(uint32, time.Time) error

func (fn dyadicErrorUint32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8IntCallFn func(uint8, int) error

func (fn dyadicErrorUint8IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8Int64CallFn func(uint8, int64) error

func (fn dyadicErrorUint8Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8Int32CallFn func(uint8, int32) error

func (fn dyadicErrorUint8Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8DurationCallFn func(uint8, time.Duration) error

func (fn dyadicErrorUint8DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8UintCallFn func(uint8, uint) error

func (fn dyadicErrorUint8UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8Uint64CallFn func(uint8, uint64) error

func (fn dyadicErrorUint8Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8Uint32CallFn func(uint8, uint32) error

func (fn dyadicErrorUint8Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8Uint8CallFn func(uint8, uint8) error

func (fn dyadicErrorUint8Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8Float64CallFn func(uint8, float64) error

func (fn dyadicErrorUint8Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8Float32CallFn func(uint8, float32) error

func (fn dyadicErrorUint8Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8Complex128CallFn func(uint8, complex128) error

func (fn dyadicErrorUint8Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8StringCallFn func(uint8, string) error

func (fn dyadicErrorUint8StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8BytesCallFn func(uint8, []byte) error

func (fn dyadicErrorUint8BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8BoolCallFn func(uint8, bool) error

func (fn dyadicErrorUint8BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorUint8TimeCallFn func(uint8, time.Time) error

func (fn dyadicErrorUint8TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64IntCallFn func(float64, int) error

func (fn dyadicErrorFloat64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64Int64CallFn func(float64, int64) error

func (fn dyadicErrorFloat64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64Int32CallFn func(float64, int32) error

func (fn dyadicErrorFloat64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64DurationCallFn func(float64, time.Duration) error

func (fn dyadicErrorFloat64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64UintCallFn func(float64, uint) error

func (fn dyadicErrorFloat64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64Uint64CallFn func(float64, uint64) error

func (fn dyadicErrorFloat64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64Uint32CallFn func(float64, uint32) error

func (fn dyadicErrorFloat64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64Uint8CallFn func(float64, uint8) error

func (fn dyadicErrorFloat64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64Float64CallFn func(float64, float64) error

func (fn dyadicErrorFloat64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64Float32CallFn func(float64, float32) error

func (fn dyadicErrorFloat64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64Complex128CallFn func(float64, complex128) error

func (fn dyadicErrorFloat64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64StringCallFn func(float64, string) error

func (fn dyadicErrorFloat64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64BytesCallFn func(float64, []byte) error

func (fn dyadicErrorFloat64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64BoolCallFn func(float64, bool) error

func (fn dyadicErrorFloat64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat64TimeCallFn func(float64, time.Time) error

func (fn dyadicErrorFloat64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32IntCallFn func(float32, int) error

func (fn dyadicErrorFloat32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32Int64CallFn func(float32, int64) error

func (fn dyadicErrorFloat32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32Int32CallFn func(float32, int32) error

func (fn dyadicErrorFloat32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32DurationCallFn func(float32, time.Duration) error

func (fn dyadicErrorFloat32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32UintCallFn func(float32, uint) error

func (fn dyadicErrorFloat32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32Uint64CallFn func(float32, uint64) error

func (fn dyadicErrorFloat32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32Uint32CallFn func(float32, uint32) error

func (fn dyadicErrorFloat32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32Uint8CallFn func(float32, uint8) error

func (fn dyadicErrorFloat32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32Float64CallFn func(float32, float64) error

func (fn dyadicErrorFloat32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32Float32CallFn func(float32, float32) error

func (fn dyadicErrorFloat32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32Complex128CallFn func(float32, complex128) error

func (fn dyadicErrorFloat32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32StringCallFn func(float32, string) error

func (fn dyadicErrorFloat32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32BytesCallFn func(float32, []byte) error

func (fn dyadicErrorFloat32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32BoolCallFn func(float32, bool) error

func (fn dyadicErrorFloat32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorFloat32TimeCallFn func(float32, time.Time) error

func (fn dyadicErrorFloat32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128IntCallFn func(complex128, int) error

func (fn dyadicErrorComplex128IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128Int64CallFn func(complex128, int64) error

func (fn dyadicErrorComplex128Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128Int32CallFn func(complex128, int32) error

func (fn dyadicErrorComplex128Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128DurationCallFn func(complex128, time.Duration) error

func (fn dyadicErrorComplex128DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128UintCallFn func(complex128, uint) error

func (fn dyadicErrorComplex128UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128Uint64CallFn func(complex128, uint64) error

func (fn dyadicErrorComplex128Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128Uint32CallFn func(complex128, uint32) error

func (fn dyadicErrorComplex128Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128Uint8CallFn func(complex128, uint8) error

func (fn dyadicErrorComplex128Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128Float64CallFn func(complex128, float64) error

func (fn dyadicErrorComplex128Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128Float32CallFn func(complex128, float32) error

func (fn dyadicErrorComplex128Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128Complex128CallFn func(complex128, complex128) error

func (fn dyadicErrorComplex128Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128StringCallFn func(complex128, string) error

func (fn dyadicErrorComplex128StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128BytesCallFn func(complex128, []byte) error

func (fn dyadicErrorComplex128BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128BoolCallFn func(complex128, bool) error

func (fn dyadicErrorComplex128BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorComplex128TimeCallFn func(complex128, time.Time) error

func (fn dyadicErrorComplex128TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringIntCallFn func(string, int) error

func (fn dyadicErrorStringIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringInt64CallFn func(string, int64) error

func (fn dyadicErrorStringInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringInt32CallFn func(string, int32) error

func (fn dyadicErrorStringInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringDurationCallFn func(string, time.Duration) error

func (fn dyadicErrorStringDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringUintCallFn func(string, uint) error

func (fn dyadicErrorStringUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringUint64CallFn func(string, uint64) error

func (fn dyadicErrorStringUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringUint32CallFn func(string, uint32) error

func (fn dyadicErrorStringUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringUint8CallFn func(string, uint8) error

func (fn dyadicErrorStringUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringFloat64CallFn func(string, float64) error

func (fn dyadicErrorStringFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringFloat32CallFn func(string, float32) error

func (fn dyadicErrorStringFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringComplex128CallFn func(string, complex128) error

func (fn dyadicErrorStringComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringStringCallFn func(string, string) error

func (fn dyadicErrorStringStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringBytesCallFn func(string, []byte) error

func (fn dyadicErrorStringBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringBoolCallFn func(string, bool) error

func (fn dyadicErrorStringBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorStringTimeCallFn func(string, time.Time) error

func (fn dyadicErrorStringTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesIntCallFn func([]byte, int) error

func (fn dyadicErrorBytesIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesInt64CallFn func([]byte, int64) error

func (fn dyadicErrorBytesInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesInt32CallFn func([]byte, int32) error

func (fn dyadicErrorBytesInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesDurationCallFn func([]byte, time.Duration) error

func (fn dyadicErrorBytesDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesUintCallFn func([]byte, uint) error

func (fn dyadicErrorBytesUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesUint64CallFn func([]byte, uint64) error

func (fn dyadicErrorBytesUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesUint32CallFn func([]byte, uint32) error

func (fn dyadicErrorBytesUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesUint8CallFn func([]byte, uint8) error

func (fn dyadicErrorBytesUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesFloat64CallFn func([]byte, float64) error

func (fn dyadicErrorBytesFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesFloat32CallFn func([]byte, float32) error

func (fn dyadicErrorBytesFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesComplex128CallFn func([]byte, complex128) error

func (fn dyadicErrorBytesComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesStringCallFn func([]byte, string) error

func (fn dyadicErrorBytesStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesBytesCallFn func([]byte, []byte) error

func (fn dyadicErrorBytesBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesBoolCallFn func([]byte, bool) error

func (fn dyadicErrorBytesBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBytesTimeCallFn func([]byte, time.Time) error

func (fn dyadicErrorBytesTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolIntCallFn func(bool, int) error

func (fn dyadicErrorBoolIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolInt64CallFn func(bool, int64) error

func (fn dyadicErrorBoolInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolInt32CallFn func(bool, int32) error

func (fn dyadicErrorBoolInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolDurationCallFn func(bool, time.Duration) error

func (fn dyadicErrorBoolDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolUintCallFn func(bool, uint) error

func (fn dyadicErrorBoolUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolUint64CallFn func(bool, uint64) error

func (fn dyadicErrorBoolUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolUint32CallFn func(bool, uint32) error

func (fn dyadicErrorBoolUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolUint8CallFn func(bool, uint8) error

func (fn dyadicErrorBoolUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolFloat64CallFn func(bool, float64) error

func (fn dyadicErrorBoolFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolFloat32CallFn func(bool, float32) error

func (fn dyadicErrorBoolFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolComplex128CallFn func(bool, complex128) error

func (fn dyadicErrorBoolComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolStringCallFn func(bool, string) error

func (fn dyadicErrorBoolStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolBytesCallFn func(bool, []byte) error

func (fn dyadicErrorBoolBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolBoolCallFn func(bool, bool) error

func (fn dyadicErrorBoolBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorBoolTimeCallFn func(bool, time.Time) error

func (fn dyadicErrorBoolTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeIntCallFn func(time.Time, int) error

func (fn dyadicErrorTimeIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeInt64CallFn func(time.Time, int64) error

func (fn dyadicErrorTimeInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeInt32CallFn func(time.Time, int32) error

func (fn dyadicErrorTimeInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeDurationCallFn func(time.Time, time.Duration) error

func (fn dyadicErrorTimeDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeUintCallFn func(time.Time, uint) error

func (fn dyadicErrorTimeUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeUint64CallFn func(time.Time, uint64) error

func (fn dyadicErrorTimeUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeUint32CallFn func(time.Time, uint32) error

func (fn dyadicErrorTimeUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeUint8CallFn func(time.Time, uint8) error

func (fn dyadicErrorTimeUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeFloat64CallFn func(time.Time, float64) error

func (fn dyadicErrorTimeFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeFloat32CallFn func(time.Time, float32) error

func (fn dyadicErrorTimeFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeComplex128CallFn func(time.Time, complex128) error

func (fn dyadicErrorTimeComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeStringCallFn func(time.Time, string) error

func (fn dyadicErrorTimeStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeBytesCallFn func(time.Time, []byte) error

func (fn dyadicErrorTimeBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].([]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeBoolCallFn func(time.Time, bool) error

func (fn dyadicErrorTimeBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorTimeTimeCallFn func(time.Time, time.Time) error

func (fn dyadicErrorTimeTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	arg1, ok1 := in[1].(time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}
