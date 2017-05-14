package call

import (
	"time"
)

type dyadicVoidPtrIntIntCallFn func(*int, *int)

func (fn dyadicVoidPtrIntIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntInt64CallFn func(*int, *int64)

func (fn dyadicVoidPtrIntInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntInt32CallFn func(*int, *int32)

func (fn dyadicVoidPtrIntInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntDurationCallFn func(*int, *time.Duration)

func (fn dyadicVoidPtrIntDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntUintCallFn func(*int, *uint)

func (fn dyadicVoidPtrIntUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntUint64CallFn func(*int, *uint64)

func (fn dyadicVoidPtrIntUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntUint32CallFn func(*int, *uint32)

func (fn dyadicVoidPtrIntUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntUint8CallFn func(*int, *uint8)

func (fn dyadicVoidPtrIntUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntFloat64CallFn func(*int, *float64)

func (fn dyadicVoidPtrIntFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntFloat32CallFn func(*int, *float32)

func (fn dyadicVoidPtrIntFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntComplex128CallFn func(*int, *complex128)

func (fn dyadicVoidPtrIntComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntStringCallFn func(*int, *string)

func (fn dyadicVoidPtrIntStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntBytesCallFn func(*int, *[]byte)

func (fn dyadicVoidPtrIntBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntBoolCallFn func(*int, *bool)

func (fn dyadicVoidPtrIntBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrIntTimeCallFn func(*int, *time.Time)

func (fn dyadicVoidPtrIntTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64IntCallFn func(*int64, *int)

func (fn dyadicVoidPtrInt64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64Int64CallFn func(*int64, *int64)

func (fn dyadicVoidPtrInt64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64Int32CallFn func(*int64, *int32)

func (fn dyadicVoidPtrInt64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64DurationCallFn func(*int64, *time.Duration)

func (fn dyadicVoidPtrInt64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64UintCallFn func(*int64, *uint)

func (fn dyadicVoidPtrInt64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64Uint64CallFn func(*int64, *uint64)

func (fn dyadicVoidPtrInt64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64Uint32CallFn func(*int64, *uint32)

func (fn dyadicVoidPtrInt64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64Uint8CallFn func(*int64, *uint8)

func (fn dyadicVoidPtrInt64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64Float64CallFn func(*int64, *float64)

func (fn dyadicVoidPtrInt64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64Float32CallFn func(*int64, *float32)

func (fn dyadicVoidPtrInt64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64Complex128CallFn func(*int64, *complex128)

func (fn dyadicVoidPtrInt64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64StringCallFn func(*int64, *string)

func (fn dyadicVoidPtrInt64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64BytesCallFn func(*int64, *[]byte)

func (fn dyadicVoidPtrInt64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64BoolCallFn func(*int64, *bool)

func (fn dyadicVoidPtrInt64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt64TimeCallFn func(*int64, *time.Time)

func (fn dyadicVoidPtrInt64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32IntCallFn func(*int32, *int)

func (fn dyadicVoidPtrInt32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32Int64CallFn func(*int32, *int64)

func (fn dyadicVoidPtrInt32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32Int32CallFn func(*int32, *int32)

func (fn dyadicVoidPtrInt32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32DurationCallFn func(*int32, *time.Duration)

func (fn dyadicVoidPtrInt32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32UintCallFn func(*int32, *uint)

func (fn dyadicVoidPtrInt32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32Uint64CallFn func(*int32, *uint64)

func (fn dyadicVoidPtrInt32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32Uint32CallFn func(*int32, *uint32)

func (fn dyadicVoidPtrInt32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32Uint8CallFn func(*int32, *uint8)

func (fn dyadicVoidPtrInt32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32Float64CallFn func(*int32, *float64)

func (fn dyadicVoidPtrInt32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32Float32CallFn func(*int32, *float32)

func (fn dyadicVoidPtrInt32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32Complex128CallFn func(*int32, *complex128)

func (fn dyadicVoidPtrInt32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32StringCallFn func(*int32, *string)

func (fn dyadicVoidPtrInt32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32BytesCallFn func(*int32, *[]byte)

func (fn dyadicVoidPtrInt32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32BoolCallFn func(*int32, *bool)

func (fn dyadicVoidPtrInt32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrInt32TimeCallFn func(*int32, *time.Time)

func (fn dyadicVoidPtrInt32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationIntCallFn func(*time.Duration, *int)

func (fn dyadicVoidPtrDurationIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationInt64CallFn func(*time.Duration, *int64)

func (fn dyadicVoidPtrDurationInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationInt32CallFn func(*time.Duration, *int32)

func (fn dyadicVoidPtrDurationInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationDurationCallFn func(*time.Duration, *time.Duration)

func (fn dyadicVoidPtrDurationDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationUintCallFn func(*time.Duration, *uint)

func (fn dyadicVoidPtrDurationUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationUint64CallFn func(*time.Duration, *uint64)

func (fn dyadicVoidPtrDurationUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationUint32CallFn func(*time.Duration, *uint32)

func (fn dyadicVoidPtrDurationUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationUint8CallFn func(*time.Duration, *uint8)

func (fn dyadicVoidPtrDurationUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationFloat64CallFn func(*time.Duration, *float64)

func (fn dyadicVoidPtrDurationFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationFloat32CallFn func(*time.Duration, *float32)

func (fn dyadicVoidPtrDurationFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationComplex128CallFn func(*time.Duration, *complex128)

func (fn dyadicVoidPtrDurationComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationStringCallFn func(*time.Duration, *string)

func (fn dyadicVoidPtrDurationStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationBytesCallFn func(*time.Duration, *[]byte)

func (fn dyadicVoidPtrDurationBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationBoolCallFn func(*time.Duration, *bool)

func (fn dyadicVoidPtrDurationBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrDurationTimeCallFn func(*time.Duration, *time.Time)

func (fn dyadicVoidPtrDurationTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintIntCallFn func(*uint, *int)

func (fn dyadicVoidPtrUintIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintInt64CallFn func(*uint, *int64)

func (fn dyadicVoidPtrUintInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintInt32CallFn func(*uint, *int32)

func (fn dyadicVoidPtrUintInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintDurationCallFn func(*uint, *time.Duration)

func (fn dyadicVoidPtrUintDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintUintCallFn func(*uint, *uint)

func (fn dyadicVoidPtrUintUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintUint64CallFn func(*uint, *uint64)

func (fn dyadicVoidPtrUintUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintUint32CallFn func(*uint, *uint32)

func (fn dyadicVoidPtrUintUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintUint8CallFn func(*uint, *uint8)

func (fn dyadicVoidPtrUintUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintFloat64CallFn func(*uint, *float64)

func (fn dyadicVoidPtrUintFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintFloat32CallFn func(*uint, *float32)

func (fn dyadicVoidPtrUintFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintComplex128CallFn func(*uint, *complex128)

func (fn dyadicVoidPtrUintComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintStringCallFn func(*uint, *string)

func (fn dyadicVoidPtrUintStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintBytesCallFn func(*uint, *[]byte)

func (fn dyadicVoidPtrUintBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintBoolCallFn func(*uint, *bool)

func (fn dyadicVoidPtrUintBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUintTimeCallFn func(*uint, *time.Time)

func (fn dyadicVoidPtrUintTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64IntCallFn func(*uint64, *int)

func (fn dyadicVoidPtrUint64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64Int64CallFn func(*uint64, *int64)

func (fn dyadicVoidPtrUint64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64Int32CallFn func(*uint64, *int32)

func (fn dyadicVoidPtrUint64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64DurationCallFn func(*uint64, *time.Duration)

func (fn dyadicVoidPtrUint64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64UintCallFn func(*uint64, *uint)

func (fn dyadicVoidPtrUint64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64Uint64CallFn func(*uint64, *uint64)

func (fn dyadicVoidPtrUint64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64Uint32CallFn func(*uint64, *uint32)

func (fn dyadicVoidPtrUint64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64Uint8CallFn func(*uint64, *uint8)

func (fn dyadicVoidPtrUint64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64Float64CallFn func(*uint64, *float64)

func (fn dyadicVoidPtrUint64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64Float32CallFn func(*uint64, *float32)

func (fn dyadicVoidPtrUint64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64Complex128CallFn func(*uint64, *complex128)

func (fn dyadicVoidPtrUint64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64StringCallFn func(*uint64, *string)

func (fn dyadicVoidPtrUint64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64BytesCallFn func(*uint64, *[]byte)

func (fn dyadicVoidPtrUint64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64BoolCallFn func(*uint64, *bool)

func (fn dyadicVoidPtrUint64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint64TimeCallFn func(*uint64, *time.Time)

func (fn dyadicVoidPtrUint64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32IntCallFn func(*uint32, *int)

func (fn dyadicVoidPtrUint32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32Int64CallFn func(*uint32, *int64)

func (fn dyadicVoidPtrUint32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32Int32CallFn func(*uint32, *int32)

func (fn dyadicVoidPtrUint32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32DurationCallFn func(*uint32, *time.Duration)

func (fn dyadicVoidPtrUint32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32UintCallFn func(*uint32, *uint)

func (fn dyadicVoidPtrUint32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32Uint64CallFn func(*uint32, *uint64)

func (fn dyadicVoidPtrUint32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32Uint32CallFn func(*uint32, *uint32)

func (fn dyadicVoidPtrUint32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32Uint8CallFn func(*uint32, *uint8)

func (fn dyadicVoidPtrUint32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32Float64CallFn func(*uint32, *float64)

func (fn dyadicVoidPtrUint32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32Float32CallFn func(*uint32, *float32)

func (fn dyadicVoidPtrUint32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32Complex128CallFn func(*uint32, *complex128)

func (fn dyadicVoidPtrUint32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32StringCallFn func(*uint32, *string)

func (fn dyadicVoidPtrUint32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32BytesCallFn func(*uint32, *[]byte)

func (fn dyadicVoidPtrUint32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32BoolCallFn func(*uint32, *bool)

func (fn dyadicVoidPtrUint32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint32TimeCallFn func(*uint32, *time.Time)

func (fn dyadicVoidPtrUint32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8IntCallFn func(*uint8, *int)

func (fn dyadicVoidPtrUint8IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8Int64CallFn func(*uint8, *int64)

func (fn dyadicVoidPtrUint8Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8Int32CallFn func(*uint8, *int32)

func (fn dyadicVoidPtrUint8Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8DurationCallFn func(*uint8, *time.Duration)

func (fn dyadicVoidPtrUint8DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8UintCallFn func(*uint8, *uint)

func (fn dyadicVoidPtrUint8UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8Uint64CallFn func(*uint8, *uint64)

func (fn dyadicVoidPtrUint8Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8Uint32CallFn func(*uint8, *uint32)

func (fn dyadicVoidPtrUint8Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8Uint8CallFn func(*uint8, *uint8)

func (fn dyadicVoidPtrUint8Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8Float64CallFn func(*uint8, *float64)

func (fn dyadicVoidPtrUint8Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8Float32CallFn func(*uint8, *float32)

func (fn dyadicVoidPtrUint8Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8Complex128CallFn func(*uint8, *complex128)

func (fn dyadicVoidPtrUint8Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8StringCallFn func(*uint8, *string)

func (fn dyadicVoidPtrUint8StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8BytesCallFn func(*uint8, *[]byte)

func (fn dyadicVoidPtrUint8BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8BoolCallFn func(*uint8, *bool)

func (fn dyadicVoidPtrUint8BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrUint8TimeCallFn func(*uint8, *time.Time)

func (fn dyadicVoidPtrUint8TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64IntCallFn func(*float64, *int)

func (fn dyadicVoidPtrFloat64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64Int64CallFn func(*float64, *int64)

func (fn dyadicVoidPtrFloat64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64Int32CallFn func(*float64, *int32)

func (fn dyadicVoidPtrFloat64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64DurationCallFn func(*float64, *time.Duration)

func (fn dyadicVoidPtrFloat64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64UintCallFn func(*float64, *uint)

func (fn dyadicVoidPtrFloat64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64Uint64CallFn func(*float64, *uint64)

func (fn dyadicVoidPtrFloat64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64Uint32CallFn func(*float64, *uint32)

func (fn dyadicVoidPtrFloat64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64Uint8CallFn func(*float64, *uint8)

func (fn dyadicVoidPtrFloat64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64Float64CallFn func(*float64, *float64)

func (fn dyadicVoidPtrFloat64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64Float32CallFn func(*float64, *float32)

func (fn dyadicVoidPtrFloat64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64Complex128CallFn func(*float64, *complex128)

func (fn dyadicVoidPtrFloat64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64StringCallFn func(*float64, *string)

func (fn dyadicVoidPtrFloat64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64BytesCallFn func(*float64, *[]byte)

func (fn dyadicVoidPtrFloat64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64BoolCallFn func(*float64, *bool)

func (fn dyadicVoidPtrFloat64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat64TimeCallFn func(*float64, *time.Time)

func (fn dyadicVoidPtrFloat64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32IntCallFn func(*float32, *int)

func (fn dyadicVoidPtrFloat32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32Int64CallFn func(*float32, *int64)

func (fn dyadicVoidPtrFloat32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32Int32CallFn func(*float32, *int32)

func (fn dyadicVoidPtrFloat32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32DurationCallFn func(*float32, *time.Duration)

func (fn dyadicVoidPtrFloat32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32UintCallFn func(*float32, *uint)

func (fn dyadicVoidPtrFloat32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32Uint64CallFn func(*float32, *uint64)

func (fn dyadicVoidPtrFloat32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32Uint32CallFn func(*float32, *uint32)

func (fn dyadicVoidPtrFloat32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32Uint8CallFn func(*float32, *uint8)

func (fn dyadicVoidPtrFloat32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32Float64CallFn func(*float32, *float64)

func (fn dyadicVoidPtrFloat32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32Float32CallFn func(*float32, *float32)

func (fn dyadicVoidPtrFloat32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32Complex128CallFn func(*float32, *complex128)

func (fn dyadicVoidPtrFloat32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32StringCallFn func(*float32, *string)

func (fn dyadicVoidPtrFloat32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32BytesCallFn func(*float32, *[]byte)

func (fn dyadicVoidPtrFloat32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32BoolCallFn func(*float32, *bool)

func (fn dyadicVoidPtrFloat32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrFloat32TimeCallFn func(*float32, *time.Time)

func (fn dyadicVoidPtrFloat32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128IntCallFn func(*complex128, *int)

func (fn dyadicVoidPtrComplex128IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128Int64CallFn func(*complex128, *int64)

func (fn dyadicVoidPtrComplex128Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128Int32CallFn func(*complex128, *int32)

func (fn dyadicVoidPtrComplex128Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128DurationCallFn func(*complex128, *time.Duration)

func (fn dyadicVoidPtrComplex128DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128UintCallFn func(*complex128, *uint)

func (fn dyadicVoidPtrComplex128UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128Uint64CallFn func(*complex128, *uint64)

func (fn dyadicVoidPtrComplex128Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128Uint32CallFn func(*complex128, *uint32)

func (fn dyadicVoidPtrComplex128Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128Uint8CallFn func(*complex128, *uint8)

func (fn dyadicVoidPtrComplex128Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128Float64CallFn func(*complex128, *float64)

func (fn dyadicVoidPtrComplex128Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128Float32CallFn func(*complex128, *float32)

func (fn dyadicVoidPtrComplex128Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128Complex128CallFn func(*complex128, *complex128)

func (fn dyadicVoidPtrComplex128Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128StringCallFn func(*complex128, *string)

func (fn dyadicVoidPtrComplex128StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128BytesCallFn func(*complex128, *[]byte)

func (fn dyadicVoidPtrComplex128BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128BoolCallFn func(*complex128, *bool)

func (fn dyadicVoidPtrComplex128BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrComplex128TimeCallFn func(*complex128, *time.Time)

func (fn dyadicVoidPtrComplex128TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringIntCallFn func(*string, *int)

func (fn dyadicVoidPtrStringIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringInt64CallFn func(*string, *int64)

func (fn dyadicVoidPtrStringInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringInt32CallFn func(*string, *int32)

func (fn dyadicVoidPtrStringInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringDurationCallFn func(*string, *time.Duration)

func (fn dyadicVoidPtrStringDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringUintCallFn func(*string, *uint)

func (fn dyadicVoidPtrStringUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringUint64CallFn func(*string, *uint64)

func (fn dyadicVoidPtrStringUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringUint32CallFn func(*string, *uint32)

func (fn dyadicVoidPtrStringUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringUint8CallFn func(*string, *uint8)

func (fn dyadicVoidPtrStringUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringFloat64CallFn func(*string, *float64)

func (fn dyadicVoidPtrStringFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringFloat32CallFn func(*string, *float32)

func (fn dyadicVoidPtrStringFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringComplex128CallFn func(*string, *complex128)

func (fn dyadicVoidPtrStringComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringStringCallFn func(*string, *string)

func (fn dyadicVoidPtrStringStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringBytesCallFn func(*string, *[]byte)

func (fn dyadicVoidPtrStringBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringBoolCallFn func(*string, *bool)

func (fn dyadicVoidPtrStringBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrStringTimeCallFn func(*string, *time.Time)

func (fn dyadicVoidPtrStringTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesIntCallFn func(*[]byte, *int)

func (fn dyadicVoidPtrBytesIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesInt64CallFn func(*[]byte, *int64)

func (fn dyadicVoidPtrBytesInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesInt32CallFn func(*[]byte, *int32)

func (fn dyadicVoidPtrBytesInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesDurationCallFn func(*[]byte, *time.Duration)

func (fn dyadicVoidPtrBytesDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesUintCallFn func(*[]byte, *uint)

func (fn dyadicVoidPtrBytesUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesUint64CallFn func(*[]byte, *uint64)

func (fn dyadicVoidPtrBytesUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesUint32CallFn func(*[]byte, *uint32)

func (fn dyadicVoidPtrBytesUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesUint8CallFn func(*[]byte, *uint8)

func (fn dyadicVoidPtrBytesUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesFloat64CallFn func(*[]byte, *float64)

func (fn dyadicVoidPtrBytesFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesFloat32CallFn func(*[]byte, *float32)

func (fn dyadicVoidPtrBytesFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesComplex128CallFn func(*[]byte, *complex128)

func (fn dyadicVoidPtrBytesComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesStringCallFn func(*[]byte, *string)

func (fn dyadicVoidPtrBytesStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesBytesCallFn func(*[]byte, *[]byte)

func (fn dyadicVoidPtrBytesBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesBoolCallFn func(*[]byte, *bool)

func (fn dyadicVoidPtrBytesBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBytesTimeCallFn func(*[]byte, *time.Time)

func (fn dyadicVoidPtrBytesTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolIntCallFn func(*bool, *int)

func (fn dyadicVoidPtrBoolIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolInt64CallFn func(*bool, *int64)

func (fn dyadicVoidPtrBoolInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolInt32CallFn func(*bool, *int32)

func (fn dyadicVoidPtrBoolInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolDurationCallFn func(*bool, *time.Duration)

func (fn dyadicVoidPtrBoolDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolUintCallFn func(*bool, *uint)

func (fn dyadicVoidPtrBoolUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolUint64CallFn func(*bool, *uint64)

func (fn dyadicVoidPtrBoolUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolUint32CallFn func(*bool, *uint32)

func (fn dyadicVoidPtrBoolUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolUint8CallFn func(*bool, *uint8)

func (fn dyadicVoidPtrBoolUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolFloat64CallFn func(*bool, *float64)

func (fn dyadicVoidPtrBoolFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolFloat32CallFn func(*bool, *float32)

func (fn dyadicVoidPtrBoolFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolComplex128CallFn func(*bool, *complex128)

func (fn dyadicVoidPtrBoolComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolStringCallFn func(*bool, *string)

func (fn dyadicVoidPtrBoolStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolBytesCallFn func(*bool, *[]byte)

func (fn dyadicVoidPtrBoolBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolBoolCallFn func(*bool, *bool)

func (fn dyadicVoidPtrBoolBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrBoolTimeCallFn func(*bool, *time.Time)

func (fn dyadicVoidPtrBoolTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeIntCallFn func(*time.Time, *int)

func (fn dyadicVoidPtrTimeIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeInt64CallFn func(*time.Time, *int64)

func (fn dyadicVoidPtrTimeInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeInt32CallFn func(*time.Time, *int32)

func (fn dyadicVoidPtrTimeInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeDurationCallFn func(*time.Time, *time.Duration)

func (fn dyadicVoidPtrTimeDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeUintCallFn func(*time.Time, *uint)

func (fn dyadicVoidPtrTimeUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeUint64CallFn func(*time.Time, *uint64)

func (fn dyadicVoidPtrTimeUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeUint32CallFn func(*time.Time, *uint32)

func (fn dyadicVoidPtrTimeUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeUint8CallFn func(*time.Time, *uint8)

func (fn dyadicVoidPtrTimeUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeFloat64CallFn func(*time.Time, *float64)

func (fn dyadicVoidPtrTimeFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeFloat32CallFn func(*time.Time, *float32)

func (fn dyadicVoidPtrTimeFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeComplex128CallFn func(*time.Time, *complex128)

func (fn dyadicVoidPtrTimeComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeStringCallFn func(*time.Time, *string)

func (fn dyadicVoidPtrTimeStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeBytesCallFn func(*time.Time, *[]byte)

func (fn dyadicVoidPtrTimeBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeBoolCallFn func(*time.Time, *bool)

func (fn dyadicVoidPtrTimeBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicVoidPtrTimeTimeCallFn func(*time.Time, *time.Time)

func (fn dyadicVoidPtrTimeTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		fn(arg0, arg1)
		return nil
	}
	return nil
}

type dyadicErrorPtrIntIntCallFn func(*int, *int) error

func (fn dyadicErrorPtrIntIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntInt64CallFn func(*int, *int64) error

func (fn dyadicErrorPtrIntInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntInt32CallFn func(*int, *int32) error

func (fn dyadicErrorPtrIntInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntDurationCallFn func(*int, *time.Duration) error

func (fn dyadicErrorPtrIntDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntUintCallFn func(*int, *uint) error

func (fn dyadicErrorPtrIntUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntUint64CallFn func(*int, *uint64) error

func (fn dyadicErrorPtrIntUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntUint32CallFn func(*int, *uint32) error

func (fn dyadicErrorPtrIntUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntUint8CallFn func(*int, *uint8) error

func (fn dyadicErrorPtrIntUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntFloat64CallFn func(*int, *float64) error

func (fn dyadicErrorPtrIntFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntFloat32CallFn func(*int, *float32) error

func (fn dyadicErrorPtrIntFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntComplex128CallFn func(*int, *complex128) error

func (fn dyadicErrorPtrIntComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntStringCallFn func(*int, *string) error

func (fn dyadicErrorPtrIntStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntBytesCallFn func(*int, *[]byte) error

func (fn dyadicErrorPtrIntBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntBoolCallFn func(*int, *bool) error

func (fn dyadicErrorPtrIntBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrIntTimeCallFn func(*int, *time.Time) error

func (fn dyadicErrorPtrIntTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64IntCallFn func(*int64, *int) error

func (fn dyadicErrorPtrInt64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64Int64CallFn func(*int64, *int64) error

func (fn dyadicErrorPtrInt64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64Int32CallFn func(*int64, *int32) error

func (fn dyadicErrorPtrInt64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64DurationCallFn func(*int64, *time.Duration) error

func (fn dyadicErrorPtrInt64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64UintCallFn func(*int64, *uint) error

func (fn dyadicErrorPtrInt64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64Uint64CallFn func(*int64, *uint64) error

func (fn dyadicErrorPtrInt64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64Uint32CallFn func(*int64, *uint32) error

func (fn dyadicErrorPtrInt64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64Uint8CallFn func(*int64, *uint8) error

func (fn dyadicErrorPtrInt64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64Float64CallFn func(*int64, *float64) error

func (fn dyadicErrorPtrInt64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64Float32CallFn func(*int64, *float32) error

func (fn dyadicErrorPtrInt64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64Complex128CallFn func(*int64, *complex128) error

func (fn dyadicErrorPtrInt64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64StringCallFn func(*int64, *string) error

func (fn dyadicErrorPtrInt64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64BytesCallFn func(*int64, *[]byte) error

func (fn dyadicErrorPtrInt64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64BoolCallFn func(*int64, *bool) error

func (fn dyadicErrorPtrInt64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt64TimeCallFn func(*int64, *time.Time) error

func (fn dyadicErrorPtrInt64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32IntCallFn func(*int32, *int) error

func (fn dyadicErrorPtrInt32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32Int64CallFn func(*int32, *int64) error

func (fn dyadicErrorPtrInt32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32Int32CallFn func(*int32, *int32) error

func (fn dyadicErrorPtrInt32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32DurationCallFn func(*int32, *time.Duration) error

func (fn dyadicErrorPtrInt32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32UintCallFn func(*int32, *uint) error

func (fn dyadicErrorPtrInt32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32Uint64CallFn func(*int32, *uint64) error

func (fn dyadicErrorPtrInt32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32Uint32CallFn func(*int32, *uint32) error

func (fn dyadicErrorPtrInt32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32Uint8CallFn func(*int32, *uint8) error

func (fn dyadicErrorPtrInt32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32Float64CallFn func(*int32, *float64) error

func (fn dyadicErrorPtrInt32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32Float32CallFn func(*int32, *float32) error

func (fn dyadicErrorPtrInt32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32Complex128CallFn func(*int32, *complex128) error

func (fn dyadicErrorPtrInt32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32StringCallFn func(*int32, *string) error

func (fn dyadicErrorPtrInt32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32BytesCallFn func(*int32, *[]byte) error

func (fn dyadicErrorPtrInt32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32BoolCallFn func(*int32, *bool) error

func (fn dyadicErrorPtrInt32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrInt32TimeCallFn func(*int32, *time.Time) error

func (fn dyadicErrorPtrInt32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationIntCallFn func(*time.Duration, *int) error

func (fn dyadicErrorPtrDurationIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationInt64CallFn func(*time.Duration, *int64) error

func (fn dyadicErrorPtrDurationInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationInt32CallFn func(*time.Duration, *int32) error

func (fn dyadicErrorPtrDurationInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationDurationCallFn func(*time.Duration, *time.Duration) error

func (fn dyadicErrorPtrDurationDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationUintCallFn func(*time.Duration, *uint) error

func (fn dyadicErrorPtrDurationUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationUint64CallFn func(*time.Duration, *uint64) error

func (fn dyadicErrorPtrDurationUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationUint32CallFn func(*time.Duration, *uint32) error

func (fn dyadicErrorPtrDurationUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationUint8CallFn func(*time.Duration, *uint8) error

func (fn dyadicErrorPtrDurationUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationFloat64CallFn func(*time.Duration, *float64) error

func (fn dyadicErrorPtrDurationFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationFloat32CallFn func(*time.Duration, *float32) error

func (fn dyadicErrorPtrDurationFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationComplex128CallFn func(*time.Duration, *complex128) error

func (fn dyadicErrorPtrDurationComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationStringCallFn func(*time.Duration, *string) error

func (fn dyadicErrorPtrDurationStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationBytesCallFn func(*time.Duration, *[]byte) error

func (fn dyadicErrorPtrDurationBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationBoolCallFn func(*time.Duration, *bool) error

func (fn dyadicErrorPtrDurationBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrDurationTimeCallFn func(*time.Duration, *time.Time) error

func (fn dyadicErrorPtrDurationTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintIntCallFn func(*uint, *int) error

func (fn dyadicErrorPtrUintIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintInt64CallFn func(*uint, *int64) error

func (fn dyadicErrorPtrUintInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintInt32CallFn func(*uint, *int32) error

func (fn dyadicErrorPtrUintInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintDurationCallFn func(*uint, *time.Duration) error

func (fn dyadicErrorPtrUintDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintUintCallFn func(*uint, *uint) error

func (fn dyadicErrorPtrUintUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintUint64CallFn func(*uint, *uint64) error

func (fn dyadicErrorPtrUintUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintUint32CallFn func(*uint, *uint32) error

func (fn dyadicErrorPtrUintUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintUint8CallFn func(*uint, *uint8) error

func (fn dyadicErrorPtrUintUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintFloat64CallFn func(*uint, *float64) error

func (fn dyadicErrorPtrUintFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintFloat32CallFn func(*uint, *float32) error

func (fn dyadicErrorPtrUintFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintComplex128CallFn func(*uint, *complex128) error

func (fn dyadicErrorPtrUintComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintStringCallFn func(*uint, *string) error

func (fn dyadicErrorPtrUintStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintBytesCallFn func(*uint, *[]byte) error

func (fn dyadicErrorPtrUintBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintBoolCallFn func(*uint, *bool) error

func (fn dyadicErrorPtrUintBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUintTimeCallFn func(*uint, *time.Time) error

func (fn dyadicErrorPtrUintTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64IntCallFn func(*uint64, *int) error

func (fn dyadicErrorPtrUint64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64Int64CallFn func(*uint64, *int64) error

func (fn dyadicErrorPtrUint64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64Int32CallFn func(*uint64, *int32) error

func (fn dyadicErrorPtrUint64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64DurationCallFn func(*uint64, *time.Duration) error

func (fn dyadicErrorPtrUint64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64UintCallFn func(*uint64, *uint) error

func (fn dyadicErrorPtrUint64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64Uint64CallFn func(*uint64, *uint64) error

func (fn dyadicErrorPtrUint64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64Uint32CallFn func(*uint64, *uint32) error

func (fn dyadicErrorPtrUint64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64Uint8CallFn func(*uint64, *uint8) error

func (fn dyadicErrorPtrUint64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64Float64CallFn func(*uint64, *float64) error

func (fn dyadicErrorPtrUint64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64Float32CallFn func(*uint64, *float32) error

func (fn dyadicErrorPtrUint64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64Complex128CallFn func(*uint64, *complex128) error

func (fn dyadicErrorPtrUint64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64StringCallFn func(*uint64, *string) error

func (fn dyadicErrorPtrUint64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64BytesCallFn func(*uint64, *[]byte) error

func (fn dyadicErrorPtrUint64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64BoolCallFn func(*uint64, *bool) error

func (fn dyadicErrorPtrUint64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint64TimeCallFn func(*uint64, *time.Time) error

func (fn dyadicErrorPtrUint64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32IntCallFn func(*uint32, *int) error

func (fn dyadicErrorPtrUint32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32Int64CallFn func(*uint32, *int64) error

func (fn dyadicErrorPtrUint32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32Int32CallFn func(*uint32, *int32) error

func (fn dyadicErrorPtrUint32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32DurationCallFn func(*uint32, *time.Duration) error

func (fn dyadicErrorPtrUint32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32UintCallFn func(*uint32, *uint) error

func (fn dyadicErrorPtrUint32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32Uint64CallFn func(*uint32, *uint64) error

func (fn dyadicErrorPtrUint32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32Uint32CallFn func(*uint32, *uint32) error

func (fn dyadicErrorPtrUint32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32Uint8CallFn func(*uint32, *uint8) error

func (fn dyadicErrorPtrUint32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32Float64CallFn func(*uint32, *float64) error

func (fn dyadicErrorPtrUint32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32Float32CallFn func(*uint32, *float32) error

func (fn dyadicErrorPtrUint32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32Complex128CallFn func(*uint32, *complex128) error

func (fn dyadicErrorPtrUint32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32StringCallFn func(*uint32, *string) error

func (fn dyadicErrorPtrUint32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32BytesCallFn func(*uint32, *[]byte) error

func (fn dyadicErrorPtrUint32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32BoolCallFn func(*uint32, *bool) error

func (fn dyadicErrorPtrUint32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint32TimeCallFn func(*uint32, *time.Time) error

func (fn dyadicErrorPtrUint32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8IntCallFn func(*uint8, *int) error

func (fn dyadicErrorPtrUint8IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8Int64CallFn func(*uint8, *int64) error

func (fn dyadicErrorPtrUint8Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8Int32CallFn func(*uint8, *int32) error

func (fn dyadicErrorPtrUint8Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8DurationCallFn func(*uint8, *time.Duration) error

func (fn dyadicErrorPtrUint8DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8UintCallFn func(*uint8, *uint) error

func (fn dyadicErrorPtrUint8UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8Uint64CallFn func(*uint8, *uint64) error

func (fn dyadicErrorPtrUint8Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8Uint32CallFn func(*uint8, *uint32) error

func (fn dyadicErrorPtrUint8Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8Uint8CallFn func(*uint8, *uint8) error

func (fn dyadicErrorPtrUint8Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8Float64CallFn func(*uint8, *float64) error

func (fn dyadicErrorPtrUint8Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8Float32CallFn func(*uint8, *float32) error

func (fn dyadicErrorPtrUint8Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8Complex128CallFn func(*uint8, *complex128) error

func (fn dyadicErrorPtrUint8Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8StringCallFn func(*uint8, *string) error

func (fn dyadicErrorPtrUint8StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8BytesCallFn func(*uint8, *[]byte) error

func (fn dyadicErrorPtrUint8BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8BoolCallFn func(*uint8, *bool) error

func (fn dyadicErrorPtrUint8BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrUint8TimeCallFn func(*uint8, *time.Time) error

func (fn dyadicErrorPtrUint8TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64IntCallFn func(*float64, *int) error

func (fn dyadicErrorPtrFloat64IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64Int64CallFn func(*float64, *int64) error

func (fn dyadicErrorPtrFloat64Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64Int32CallFn func(*float64, *int32) error

func (fn dyadicErrorPtrFloat64Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64DurationCallFn func(*float64, *time.Duration) error

func (fn dyadicErrorPtrFloat64DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64UintCallFn func(*float64, *uint) error

func (fn dyadicErrorPtrFloat64UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64Uint64CallFn func(*float64, *uint64) error

func (fn dyadicErrorPtrFloat64Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64Uint32CallFn func(*float64, *uint32) error

func (fn dyadicErrorPtrFloat64Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64Uint8CallFn func(*float64, *uint8) error

func (fn dyadicErrorPtrFloat64Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64Float64CallFn func(*float64, *float64) error

func (fn dyadicErrorPtrFloat64Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64Float32CallFn func(*float64, *float32) error

func (fn dyadicErrorPtrFloat64Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64Complex128CallFn func(*float64, *complex128) error

func (fn dyadicErrorPtrFloat64Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64StringCallFn func(*float64, *string) error

func (fn dyadicErrorPtrFloat64StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64BytesCallFn func(*float64, *[]byte) error

func (fn dyadicErrorPtrFloat64BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64BoolCallFn func(*float64, *bool) error

func (fn dyadicErrorPtrFloat64BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat64TimeCallFn func(*float64, *time.Time) error

func (fn dyadicErrorPtrFloat64TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32IntCallFn func(*float32, *int) error

func (fn dyadicErrorPtrFloat32IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32Int64CallFn func(*float32, *int64) error

func (fn dyadicErrorPtrFloat32Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32Int32CallFn func(*float32, *int32) error

func (fn dyadicErrorPtrFloat32Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32DurationCallFn func(*float32, *time.Duration) error

func (fn dyadicErrorPtrFloat32DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32UintCallFn func(*float32, *uint) error

func (fn dyadicErrorPtrFloat32UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32Uint64CallFn func(*float32, *uint64) error

func (fn dyadicErrorPtrFloat32Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32Uint32CallFn func(*float32, *uint32) error

func (fn dyadicErrorPtrFloat32Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32Uint8CallFn func(*float32, *uint8) error

func (fn dyadicErrorPtrFloat32Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32Float64CallFn func(*float32, *float64) error

func (fn dyadicErrorPtrFloat32Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32Float32CallFn func(*float32, *float32) error

func (fn dyadicErrorPtrFloat32Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32Complex128CallFn func(*float32, *complex128) error

func (fn dyadicErrorPtrFloat32Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32StringCallFn func(*float32, *string) error

func (fn dyadicErrorPtrFloat32StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32BytesCallFn func(*float32, *[]byte) error

func (fn dyadicErrorPtrFloat32BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32BoolCallFn func(*float32, *bool) error

func (fn dyadicErrorPtrFloat32BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrFloat32TimeCallFn func(*float32, *time.Time) error

func (fn dyadicErrorPtrFloat32TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128IntCallFn func(*complex128, *int) error

func (fn dyadicErrorPtrComplex128IntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128Int64CallFn func(*complex128, *int64) error

func (fn dyadicErrorPtrComplex128Int64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128Int32CallFn func(*complex128, *int32) error

func (fn dyadicErrorPtrComplex128Int32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128DurationCallFn func(*complex128, *time.Duration) error

func (fn dyadicErrorPtrComplex128DurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128UintCallFn func(*complex128, *uint) error

func (fn dyadicErrorPtrComplex128UintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128Uint64CallFn func(*complex128, *uint64) error

func (fn dyadicErrorPtrComplex128Uint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128Uint32CallFn func(*complex128, *uint32) error

func (fn dyadicErrorPtrComplex128Uint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128Uint8CallFn func(*complex128, *uint8) error

func (fn dyadicErrorPtrComplex128Uint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128Float64CallFn func(*complex128, *float64) error

func (fn dyadicErrorPtrComplex128Float64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128Float32CallFn func(*complex128, *float32) error

func (fn dyadicErrorPtrComplex128Float32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128Complex128CallFn func(*complex128, *complex128) error

func (fn dyadicErrorPtrComplex128Complex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128StringCallFn func(*complex128, *string) error

func (fn dyadicErrorPtrComplex128StringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128BytesCallFn func(*complex128, *[]byte) error

func (fn dyadicErrorPtrComplex128BytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128BoolCallFn func(*complex128, *bool) error

func (fn dyadicErrorPtrComplex128BoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrComplex128TimeCallFn func(*complex128, *time.Time) error

func (fn dyadicErrorPtrComplex128TimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringIntCallFn func(*string, *int) error

func (fn dyadicErrorPtrStringIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringInt64CallFn func(*string, *int64) error

func (fn dyadicErrorPtrStringInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringInt32CallFn func(*string, *int32) error

func (fn dyadicErrorPtrStringInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringDurationCallFn func(*string, *time.Duration) error

func (fn dyadicErrorPtrStringDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringUintCallFn func(*string, *uint) error

func (fn dyadicErrorPtrStringUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringUint64CallFn func(*string, *uint64) error

func (fn dyadicErrorPtrStringUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringUint32CallFn func(*string, *uint32) error

func (fn dyadicErrorPtrStringUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringUint8CallFn func(*string, *uint8) error

func (fn dyadicErrorPtrStringUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringFloat64CallFn func(*string, *float64) error

func (fn dyadicErrorPtrStringFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringFloat32CallFn func(*string, *float32) error

func (fn dyadicErrorPtrStringFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringComplex128CallFn func(*string, *complex128) error

func (fn dyadicErrorPtrStringComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringStringCallFn func(*string, *string) error

func (fn dyadicErrorPtrStringStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringBytesCallFn func(*string, *[]byte) error

func (fn dyadicErrorPtrStringBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringBoolCallFn func(*string, *bool) error

func (fn dyadicErrorPtrStringBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrStringTimeCallFn func(*string, *time.Time) error

func (fn dyadicErrorPtrStringTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesIntCallFn func(*[]byte, *int) error

func (fn dyadicErrorPtrBytesIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesInt64CallFn func(*[]byte, *int64) error

func (fn dyadicErrorPtrBytesInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesInt32CallFn func(*[]byte, *int32) error

func (fn dyadicErrorPtrBytesInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesDurationCallFn func(*[]byte, *time.Duration) error

func (fn dyadicErrorPtrBytesDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesUintCallFn func(*[]byte, *uint) error

func (fn dyadicErrorPtrBytesUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesUint64CallFn func(*[]byte, *uint64) error

func (fn dyadicErrorPtrBytesUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesUint32CallFn func(*[]byte, *uint32) error

func (fn dyadicErrorPtrBytesUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesUint8CallFn func(*[]byte, *uint8) error

func (fn dyadicErrorPtrBytesUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesFloat64CallFn func(*[]byte, *float64) error

func (fn dyadicErrorPtrBytesFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesFloat32CallFn func(*[]byte, *float32) error

func (fn dyadicErrorPtrBytesFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesComplex128CallFn func(*[]byte, *complex128) error

func (fn dyadicErrorPtrBytesComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesStringCallFn func(*[]byte, *string) error

func (fn dyadicErrorPtrBytesStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesBytesCallFn func(*[]byte, *[]byte) error

func (fn dyadicErrorPtrBytesBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesBoolCallFn func(*[]byte, *bool) error

func (fn dyadicErrorPtrBytesBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBytesTimeCallFn func(*[]byte, *time.Time) error

func (fn dyadicErrorPtrBytesTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolIntCallFn func(*bool, *int) error

func (fn dyadicErrorPtrBoolIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolInt64CallFn func(*bool, *int64) error

func (fn dyadicErrorPtrBoolInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolInt32CallFn func(*bool, *int32) error

func (fn dyadicErrorPtrBoolInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolDurationCallFn func(*bool, *time.Duration) error

func (fn dyadicErrorPtrBoolDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolUintCallFn func(*bool, *uint) error

func (fn dyadicErrorPtrBoolUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolUint64CallFn func(*bool, *uint64) error

func (fn dyadicErrorPtrBoolUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolUint32CallFn func(*bool, *uint32) error

func (fn dyadicErrorPtrBoolUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolUint8CallFn func(*bool, *uint8) error

func (fn dyadicErrorPtrBoolUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolFloat64CallFn func(*bool, *float64) error

func (fn dyadicErrorPtrBoolFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolFloat32CallFn func(*bool, *float32) error

func (fn dyadicErrorPtrBoolFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolComplex128CallFn func(*bool, *complex128) error

func (fn dyadicErrorPtrBoolComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolStringCallFn func(*bool, *string) error

func (fn dyadicErrorPtrBoolStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolBytesCallFn func(*bool, *[]byte) error

func (fn dyadicErrorPtrBoolBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolBoolCallFn func(*bool, *bool) error

func (fn dyadicErrorPtrBoolBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrBoolTimeCallFn func(*bool, *time.Time) error

func (fn dyadicErrorPtrBoolTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeIntCallFn func(*time.Time, *int) error

func (fn dyadicErrorPtrTimeIntCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*int)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeInt64CallFn func(*time.Time, *int64) error

func (fn dyadicErrorPtrTimeInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*int64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeInt32CallFn func(*time.Time, *int32) error

func (fn dyadicErrorPtrTimeInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*int32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeDurationCallFn func(*time.Time, *time.Duration) error

func (fn dyadicErrorPtrTimeDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*time.Duration)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeUintCallFn func(*time.Time, *uint) error

func (fn dyadicErrorPtrTimeUintCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*uint)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeUint64CallFn func(*time.Time, *uint64) error

func (fn dyadicErrorPtrTimeUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*uint64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeUint32CallFn func(*time.Time, *uint32) error

func (fn dyadicErrorPtrTimeUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*uint32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeUint8CallFn func(*time.Time, *uint8) error

func (fn dyadicErrorPtrTimeUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*uint8)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeFloat64CallFn func(*time.Time, *float64) error

func (fn dyadicErrorPtrTimeFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*float64)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeFloat32CallFn func(*time.Time, *float32) error

func (fn dyadicErrorPtrTimeFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*float32)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeComplex128CallFn func(*time.Time, *complex128) error

func (fn dyadicErrorPtrTimeComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*complex128)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeStringCallFn func(*time.Time, *string) error

func (fn dyadicErrorPtrTimeStringCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*string)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeBytesCallFn func(*time.Time, *[]byte) error

func (fn dyadicErrorPtrTimeBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*[]byte)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeBoolCallFn func(*time.Time, *bool) error

func (fn dyadicErrorPtrTimeBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*bool)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}

type dyadicErrorPtrTimeTimeCallFn func(*time.Time, *time.Time) error

func (fn dyadicErrorPtrTimeTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 2 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	arg1, ok1 := in[1].(*time.Time)
	if ok0 && ok1 {
		return fn(arg0, arg1)
	}
	return nil
}
