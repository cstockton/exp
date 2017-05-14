package call

import "time"

var testLutLookupsCases = []testLutLookups{
	{0, *new(time.Duration), testLutMonadicErrorIdx0, "func(time.Duration) error"},
	{1, *new(uint), testLutMonadicErrorIdx1, "func(uint) error"},
	{2, *new(complex128), testLutMonadicErrorIdx2, "func(complex128) error"},
	{3, *new(string), testLutMonadicErrorIdx3, "func(string) error"},
	{0, new(int64), testLutMonadicErrorPtrIdx0, "func(*int64) error"},
	{1, new(uint), testLutMonadicErrorPtrIdx1, "func(*uint) error"},
	{2, new(float64), testLutMonadicErrorPtrIdx2, "func(*float64) error"},
	{3, new(bool), testLutMonadicErrorPtrIdx3, "func(*bool) error"},
	{0, *new(int), testLutMonadicVoidIdx0, "func(int)"},
	{1, *new(uint64), testLutMonadicVoidIdx1, "func(uint64)"},
	{2, *new(float64), testLutMonadicVoidIdx2, "func(float64)"},
	{3, *new(string), testLutMonadicVoidIdx3, "func(string)"},
	{0, new(int32), testLutMonadicVoidPtrIdx0, "func(*int32)"},
	{1, new(uint8), testLutMonadicVoidPtrIdx1, "func(*uint8)"},
	{2, new(float64), testLutMonadicVoidPtrIdx2, "func(*float64)"},
	{3, new([]byte), testLutMonadicVoidPtrIdx3, "func(*[]byte)"},
}

var testLutLookupsCases2x = []testLutLookups2x{
	{0, 0, new(int), new(int64), testLutDyadicErrorPtr0x0, "func(*int, *int64) error"},
	{0, 3, new(int32), new(bool), testLutDyadicVoidPtr0x3, "func(*int32, *bool)"},
	{0, 0, new(time.Duration), new(int32), testLutDyadicVoidPtr0x0, "func(*time.Duration, *int32)"},
	{1, 3, new(uint64), new(string), testLutDyadicVoidPtr1x3, "func(*uint64, *string)"},
	{3, 1, new([]byte), new(uint32), testLutDyadicVoidPtr3x1, "func(*[]byte, *uint32)"},
	{2, 2, *new(float64), *new(complex128), testLutDyadicError2x2, "func(float64, complex128) error"},
	{1, 0, *new(uint64), *new(int), testLutDyadicVoid1x0, "func(uint64, int)"},
	{3, 2, *new([]byte), *new(float64), testLutDyadicVoid3x2, "func([]byte, float64)"},
	{3, 1, *new([]byte), *new(uint8), testLutDyadicVoid3x1, "func([]byte, uint8)"},
	{0, 1, new(time.Duration), new(uint32), testLutDyadicErrorPtr0x1, "func(*time.Duration, *uint32) error"},
	{0, 1, *new(int32), *new(uint), testLutDyadicError0x1, "func(int32, uint) error"},
	{3, 3, new(string), new([]byte), testLutDyadicVoidPtr3x3, "func(*string, *[]byte)"},
	{3, 2, new([]byte), new(float64), testLutDyadicVoidPtr3x2, "func(*[]byte, *float64)"},
	{3, 3, new(bool), new(bool), testLutDyadicErrorPtr3x3, "func(*bool, *bool) error"},
	{2, 0, *new(float64), *new(int), testLutDyadicVoid2x0, "func(float64, int)"},
	{2, 3, *new(float64), *new(string), testLutDyadicVoid2x3, "func(float64, string)"},
	{1, 2, new(uint64), new(float64), testLutDyadicErrorPtr1x2, "func(*uint64, *float64) error"},
	{1, 3, *new(uint64), *new([]byte), testLutDyadicError1x3, "func(uint64, []byte) error"},
	{3, 0, *new(bool), *new(int64), testLutDyadicVoid3x0, "func(bool, int64)"},
	{0, 2, *new(int), *new(float64), testLutDyadicVoid0x2, "func(int, float64)"},
	{3, 0, new(bool), new(time.Duration), testLutDyadicVoidPtr3x0, "func(*bool, *time.Duration)"},
	{3, 1, new(string), new(uint), testLutDyadicErrorPtr3x1, "func(*string, *uint) error"},
	{2, 0, new(complex128), new(int64), testLutDyadicVoidPtr2x0, "func(*complex128, *int64)"},
	{3, 3, *new([]byte), *new(bool), testLutDyadicVoid3x3, "func([]byte, bool)"},
	{0, 1, *new(time.Duration), *new(uint32), testLutDyadicVoid0x1, "func(time.Duration, uint32)"},
	{0, 3, new(time.Duration), new(bool), testLutDyadicErrorPtr0x3, "func(*time.Duration, *bool) error"},
	{1, 3, new(uint), new([]byte), testLutDyadicErrorPtr1x3, "func(*uint, *[]byte) error"},
	{0, 3, *new(int), *new(bool), testLutDyadicVoid0x3, "func(int, bool)"},
	{3, 0, *new(bool), *new(int), testLutDyadicError3x0, "func(bool, int) error"},
	{0, 3, *new(time.Duration), *new(string), testLutDyadicError0x3, "func(time.Duration, string) error"},
	{2, 2, new(float32), new(complex128), testLutDyadicVoidPtr2x2, "func(*float32, *complex128)"},
	{2, 3, new(float64), new(bool), testLutDyadicVoidPtr2x3, "func(*float64, *bool)"},
	{0, 2, new(time.Duration), new(float32), testLutDyadicErrorPtr0x2, "func(*time.Duration, *float32) error"},
	{2, 1, *new(float64), *new(uint8), testLutDyadicError2x1, "func(float64, uint8) error"},
	{2, 0, new(complex128), new(int32), testLutDyadicErrorPtr2x0, "func(*complex128, *int32) error"},
	{0, 2, new(int32), new(float32), testLutDyadicVoidPtr0x2, "func(*int32, *float32)"},
	{2, 3, *new(complex128), *new(bool), testLutDyadicError2x3, "func(complex128, bool) error"},
	{1, 3, *new(uint32), *new(bool), testLutDyadicVoid1x3, "func(uint32, bool)"},
	{3, 2, new(bool), new(float64), testLutDyadicErrorPtr3x2, "func(*bool, *float64) error"},
	{2, 2, new(float32), new(float64), testLutDyadicErrorPtr2x2, "func(*float32, *float64) error"},
	{2, 1, new(float64), new(uint), testLutDyadicVoidPtr2x1, "func(*float64, *uint)"},
	{1, 1, *new(uint), *new(uint8), testLutDyadicVoid1x1, "func(uint, uint8)"},
	{0, 2, *new(int), *new(float32), testLutDyadicError0x2, "func(int, float32) error"},
	{0, 0, *new(int32), *new(time.Duration), testLutDyadicVoid0x0, "func(int32, time.Duration)"},
	{0, 0, *new(int64), *new(int), testLutDyadicError0x0, "func(int64, int) error"},
	{0, 1, new(int64), new(uint32), testLutDyadicVoidPtr0x1, "func(*int64, *uint32)"},
	{3, 0, new(string), new(int), testLutDyadicErrorPtr3x0, "func(*string, *int) error"},
	{1, 1, new(uint64), new(uint32), testLutDyadicErrorPtr1x1, "func(*uint64, *uint32) error"},
	{2, 1, new(float64), new(uint32), testLutDyadicErrorPtr2x1, "func(*float64, *uint32) error"},
	{1, 0, new(uint64), new(int64), testLutDyadicErrorPtr1x0, "func(*uint64, *int64) error"},
	{1, 1, new(uint), new(uint8), testLutDyadicVoidPtr1x1, "func(*uint, *uint8)"},
	{2, 2, *new(float32), *new(complex128), testLutDyadicVoid2x2, "func(float32, complex128)"},
	{1, 1, *new(uint), *new(uint64), testLutDyadicError1x1, "func(uint, uint64) error"},
	{2, 3, new(float32), new([]byte), testLutDyadicErrorPtr2x3, "func(*float32, *[]byte) error"},
	{1, 2, *new(uint8), *new(float32), testLutDyadicError1x2, "func(uint8, float32) error"},
	{1, 2, new(uint8), new(float32), testLutDyadicVoidPtr1x2, "func(*uint8, *float32)"},
	{2, 1, *new(complex128), *new(uint32), testLutDyadicVoid2x1, "func(complex128, uint32)"},
	{2, 0, *new(float64), *new(int64), testLutDyadicError2x0, "func(float64, int64) error"},
	{3, 3, *new(bool), *new(bool), testLutDyadicError3x3, "func(bool, bool) error"},
	{1, 0, new(uint8), new(int), testLutDyadicVoidPtr1x0, "func(*uint8, *int)"},
	{1, 2, *new(uint8), *new(float64), testLutDyadicVoid1x2, "func(uint8, float64)"},
	{3, 1, *new(string), *new(uint32), testLutDyadicError3x1, "func(string, uint32) error"},
	{1, 0, *new(uint32), *new(int), testLutDyadicError1x0, "func(uint32, int) error"},
	{3, 2, *new(bool), *new(complex128), testLutDyadicError3x2, "func(bool, complex128) error"},
}

func testLutMonadicErrorIdx0(time.Duration) error {
	key := "func(time.Duration) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicErrorDurationCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutMonadicErrorIdx1(uint) error {
	key := "func(uint) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicErrorUintCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutMonadicErrorIdx2(complex128) error {
	key := "func(complex128) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicErrorComplex128CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutMonadicErrorIdx3(string) error {
	key := "func(string) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicErrorStringCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutMonadicErrorPtrIdx0(*int64) error {
	key := "func(*int64) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicErrorPtrInt64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutMonadicErrorPtrIdx1(*uint) error {
	key := "func(*uint) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicErrorPtrUintCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutMonadicErrorPtrIdx2(*float64) error {
	key := "func(*float64) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicErrorPtrFloat64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutMonadicErrorPtrIdx3(*bool) error {
	key := "func(*bool) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicErrorPtrBoolCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutMonadicVoidIdx0(int) {
	key := "func(int)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicVoidIntCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutMonadicVoidIdx1(uint64) {
	key := "func(uint64)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicVoidUint64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutMonadicVoidIdx2(float64) {
	key := "func(float64)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicVoidFloat64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutMonadicVoidIdx3(string) {
	key := "func(string)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicVoidStringCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutMonadicVoidPtrIdx0(*int32) {
	key := "func(*int32)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicVoidPtrInt32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutMonadicVoidPtrIdx1(*uint8) {
	key := "func(*uint8)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicVoidPtrUint8CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutMonadicVoidPtrIdx2(*float64) {
	key := "func(*float64)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicVoidPtrFloat64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutMonadicVoidPtrIdx3(*[]byte) {
	key := "func(*[]byte)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(monadicVoidPtrBytesCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicErrorPtr0x0(*int, *int64) error {
	key := "func(*int, *int64) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrIntInt64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoidPtr0x3(*int32, *bool) {
	key := "func(*int32, *bool)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrInt32BoolCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoidPtr0x0(*time.Duration, *int32) {
	key := "func(*time.Duration, *int32)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrDurationInt32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoidPtr1x3(*uint64, *string) {
	key := "func(*uint64, *string)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrUint64StringCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoidPtr3x1(*[]byte, *uint32) {
	key := "func(*[]byte, *uint32)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrBytesUint32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicError2x2(float64, complex128) error {
	key := "func(float64, complex128) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorFloat64Complex128CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoid1x0(uint64, int) {
	key := "func(uint64, int)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidUint64IntCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoid3x2([]byte, float64) {
	key := "func([]byte, float64)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidBytesFloat64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoid3x1([]byte, uint8) {
	key := "func([]byte, uint8)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidBytesUint8CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicErrorPtr0x1(*time.Duration, *uint32) error {
	key := "func(*time.Duration, *uint32) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrDurationUint32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicError0x1(int32, uint) error {
	key := "func(int32, uint) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorInt32UintCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoidPtr3x3(*string, *[]byte) {
	key := "func(*string, *[]byte)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrStringBytesCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoidPtr3x2(*[]byte, *float64) {
	key := "func(*[]byte, *float64)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrBytesFloat64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicErrorPtr3x3(*bool, *bool) error {
	key := "func(*bool, *bool) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrBoolBoolCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoid2x0(float64, int) {
	key := "func(float64, int)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidFloat64IntCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoid2x3(float64, string) {
	key := "func(float64, string)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidFloat64StringCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicErrorPtr1x2(*uint64, *float64) error {
	key := "func(*uint64, *float64) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrUint64Float64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicError1x3(uint64, []byte) error {
	key := "func(uint64, []byte) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorUint64BytesCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoid3x0(bool, int64) {
	key := "func(bool, int64)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidBoolInt64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoid0x2(int, float64) {
	key := "func(int, float64)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidIntFloat64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoidPtr3x0(*bool, *time.Duration) {
	key := "func(*bool, *time.Duration)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrBoolDurationCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicErrorPtr3x1(*string, *uint) error {
	key := "func(*string, *uint) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrStringUintCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoidPtr2x0(*complex128, *int64) {
	key := "func(*complex128, *int64)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrComplex128Int64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoid3x3([]byte, bool) {
	key := "func([]byte, bool)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidBytesBoolCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoid0x1(time.Duration, uint32) {
	key := "func(time.Duration, uint32)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidDurationUint32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicErrorPtr0x3(*time.Duration, *bool) error {
	key := "func(*time.Duration, *bool) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrDurationBoolCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicErrorPtr1x3(*uint, *[]byte) error {
	key := "func(*uint, *[]byte) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrUintBytesCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoid0x3(int, bool) {
	key := "func(int, bool)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidIntBoolCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicError3x0(bool, int) error {
	key := "func(bool, int) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorBoolIntCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicError0x3(time.Duration, string) error {
	key := "func(time.Duration, string) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorDurationStringCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoidPtr2x2(*float32, *complex128) {
	key := "func(*float32, *complex128)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrFloat32Complex128CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoidPtr2x3(*float64, *bool) {
	key := "func(*float64, *bool)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrFloat64BoolCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicErrorPtr0x2(*time.Duration, *float32) error {
	key := "func(*time.Duration, *float32) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrDurationFloat32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicError2x1(float64, uint8) error {
	key := "func(float64, uint8) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorFloat64Uint8CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicErrorPtr2x0(*complex128, *int32) error {
	key := "func(*complex128, *int32) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrComplex128Int32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoidPtr0x2(*int32, *float32) {
	key := "func(*int32, *float32)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrInt32Float32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicError2x3(complex128, bool) error {
	key := "func(complex128, bool) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorComplex128BoolCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoid1x3(uint32, bool) {
	key := "func(uint32, bool)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidUint32BoolCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicErrorPtr3x2(*bool, *float64) error {
	key := "func(*bool, *float64) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrBoolFloat64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicErrorPtr2x2(*float32, *float64) error {
	key := "func(*float32, *float64) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrFloat32Float64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoidPtr2x1(*float64, *uint) {
	key := "func(*float64, *uint)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrFloat64UintCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoid1x1(uint, uint8) {
	key := "func(uint, uint8)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidUintUint8CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicError0x2(int, float32) error {
	key := "func(int, float32) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorIntFloat32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoid0x0(int32, time.Duration) {
	key := "func(int32, time.Duration)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidInt32DurationCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicError0x0(int64, int) error {
	key := "func(int64, int) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorInt64IntCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoidPtr0x1(*int64, *uint32) {
	key := "func(*int64, *uint32)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrInt64Uint32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicErrorPtr3x0(*string, *int) error {
	key := "func(*string, *int) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrStringIntCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicErrorPtr1x1(*uint64, *uint32) error {
	key := "func(*uint64, *uint32) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrUint64Uint32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicErrorPtr2x1(*float64, *uint32) error {
	key := "func(*float64, *uint32) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrFloat64Uint32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicErrorPtr1x0(*uint64, *int64) error {
	key := "func(*uint64, *int64) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrUint64Int64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoidPtr1x1(*uint, *uint8) {
	key := "func(*uint, *uint8)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrUintUint8CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoid2x2(float32, complex128) {
	key := "func(float32, complex128)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidFloat32Complex128CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicError1x1(uint, uint64) error {
	key := "func(uint, uint64) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorUintUint64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicErrorPtr2x3(*float32, *[]byte) error {
	key := "func(*float32, *[]byte) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorPtrFloat32BytesCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicError1x2(uint8, float32) error {
	key := "func(uint8, float32) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorUint8Float32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoidPtr1x2(*uint8, *float32) {
	key := "func(*uint8, *float32)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrUint8Float32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoid2x1(complex128, uint32) {
	key := "func(complex128, uint32)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidComplex128Uint32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicError2x0(float64, int64) error {
	key := "func(float64, int64) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorFloat64Int64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicError3x3(bool, bool) error {
	key := "func(bool, bool) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorBoolBoolCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicVoidPtr1x0(*uint8, *int) {
	key := "func(*uint8, *int)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidPtrUint8IntCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicVoid1x2(uint8, float64) {
	key := "func(uint8, float64)"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicVoidUint8Float64CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
}

func testLutDyadicError3x1(string, uint32) error {
	key := "func(string, uint32) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorStringUint32CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicError1x0(uint32, int) error {
	key := "func(uint32, int) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorUint32IntCallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}

func testLutDyadicError3x2(bool, complex128) error {
	key := "func(bool, complex128) error"
	testLutLookupsCallTable[key] = func(me interface{}) error {
		if _, ok := me.(dyadicErrorBoolComplex128CallFn); !ok {
			return errTestLutLookupFailed
		}
		return nil
	}
	return nil
}
