package call

import (
	"time"
)

type monadicVoidIntCallFn func(int)

func (fn monadicVoidIntCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidInt64CallFn func(int64)

func (fn monadicVoidInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidInt32CallFn func(int32)

func (fn monadicVoidInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidDurationCallFn func(time.Duration)

func (fn monadicVoidDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidUintCallFn func(uint)

func (fn monadicVoidUintCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidUint64CallFn func(uint64)

func (fn monadicVoidUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidUint32CallFn func(uint32)

func (fn monadicVoidUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidUint8CallFn func(uint8)

func (fn monadicVoidUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidFloat64CallFn func(float64)

func (fn monadicVoidFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidFloat32CallFn func(float32)

func (fn monadicVoidFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidComplex128CallFn func(complex128)

func (fn monadicVoidComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidStringCallFn func(string)

func (fn monadicVoidStringCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidBytesCallFn func([]byte)

func (fn monadicVoidBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidBoolCallFn func(bool)

func (fn monadicVoidBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidTimeCallFn func(time.Time)

func (fn monadicVoidTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicErrorIntCallFn func(int) error

func (fn monadicErrorIntCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorInt64CallFn func(int64) error

func (fn monadicErrorInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int64)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorInt32CallFn func(int32) error

func (fn monadicErrorInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(int32)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorDurationCallFn func(time.Duration) error

func (fn monadicErrorDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Duration)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorUintCallFn func(uint) error

func (fn monadicErrorUintCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorUint64CallFn func(uint64) error

func (fn monadicErrorUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint64)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorUint32CallFn func(uint32) error

func (fn monadicErrorUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint32)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorUint8CallFn func(uint8) error

func (fn monadicErrorUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(uint8)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorFloat64CallFn func(float64) error

func (fn monadicErrorFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float64)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorFloat32CallFn func(float32) error

func (fn monadicErrorFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(float32)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorComplex128CallFn func(complex128) error

func (fn monadicErrorComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(complex128)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorStringCallFn func(string) error

func (fn monadicErrorStringCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(string)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorBytesCallFn func([]byte) error

func (fn monadicErrorBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].([]byte)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorBoolCallFn func(bool) error

func (fn monadicErrorBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(bool)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorTimeCallFn func(time.Time) error

func (fn monadicErrorTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(time.Time)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicVoidPtrIntCallFn func(*int)

func (fn monadicVoidPtrIntCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrInt64CallFn func(*int64)

func (fn monadicVoidPtrInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrInt32CallFn func(*int32)

func (fn monadicVoidPtrInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrDurationCallFn func(*time.Duration)

func (fn monadicVoidPtrDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrUintCallFn func(*uint)

func (fn monadicVoidPtrUintCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrUint64CallFn func(*uint64)

func (fn monadicVoidPtrUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrUint32CallFn func(*uint32)

func (fn monadicVoidPtrUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrUint8CallFn func(*uint8)

func (fn monadicVoidPtrUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrFloat64CallFn func(*float64)

func (fn monadicVoidPtrFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrFloat32CallFn func(*float32)

func (fn monadicVoidPtrFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrComplex128CallFn func(*complex128)

func (fn monadicVoidPtrComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrStringCallFn func(*string)

func (fn monadicVoidPtrStringCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrBytesCallFn func(*[]byte)

func (fn monadicVoidPtrBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrBoolCallFn func(*bool)

func (fn monadicVoidPtrBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicVoidPtrTimeCallFn func(*time.Time)

func (fn monadicVoidPtrTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	if ok0 {
		fn(arg0)
		return nil
	}
	return nil
}

type monadicErrorPtrIntCallFn func(*int) error

func (fn monadicErrorPtrIntCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrInt64CallFn func(*int64) error

func (fn monadicErrorPtrInt64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int64)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrInt32CallFn func(*int32) error

func (fn monadicErrorPtrInt32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*int32)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrDurationCallFn func(*time.Duration) error

func (fn monadicErrorPtrDurationCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Duration)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrUintCallFn func(*uint) error

func (fn monadicErrorPtrUintCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrUint64CallFn func(*uint64) error

func (fn monadicErrorPtrUint64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint64)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrUint32CallFn func(*uint32) error

func (fn monadicErrorPtrUint32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint32)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrUint8CallFn func(*uint8) error

func (fn monadicErrorPtrUint8CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*uint8)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrFloat64CallFn func(*float64) error

func (fn monadicErrorPtrFloat64CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float64)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrFloat32CallFn func(*float32) error

func (fn monadicErrorPtrFloat32CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*float32)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrComplex128CallFn func(*complex128) error

func (fn monadicErrorPtrComplex128CallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*complex128)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrStringCallFn func(*string) error

func (fn monadicErrorPtrStringCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*string)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrBytesCallFn func(*[]byte) error

func (fn monadicErrorPtrBytesCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*[]byte)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrBoolCallFn func(*bool) error

func (fn monadicErrorPtrBoolCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*bool)
	if ok0 {
		return fn(arg0)
	}
	return nil
}

type monadicErrorPtrTimeCallFn func(*time.Time) error

func (fn monadicErrorPtrTimeCallFn) Call(in ...interface{}) error {
	if len(in) != 1 {
		return ErrArgCount
	}
	arg0, ok0 := in[0].(*time.Time)
	if ok0 {
		return fn(arg0)
	}
	return nil
}
