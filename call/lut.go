package call

import "reflect"

//go:generate genutil *.tpl -- -arity 2 -pointers
//go:generate gofmt -w *.go

type lutInfo int

const (
	lutInts reflect.Kind = iota
	lutUints
	lutFloats
	lutStrings
	lutSize
)

const (
	lutVoid lutInfo = iota
	lutVoidPtr
	lutError
	lutErrorPtr
)

func (l lutInfo) String() string {
	switch l {
	case lutError:
		return "LutError"
	case lutErrorPtr:
		return "LutVoidPtr"
	case lutVoid:
		return "LutVoid"
	default:
		return "LutErrorPtr"
	}
}

func lutIndexFromKind(k reflect.Kind) reflect.Kind {
	switch {
	case k > reflect.Ptr:
		return lutStrings
	case k > reflect.Uint64:
		return lutFloats
	case k > reflect.Int64:
		return lutUints
	case k > reflect.Bool:
		return lutInts
	default:
		return lutStrings
	}
}

func lutLookupMonadic(info lutInfo, idx1 reflect.Kind, fn interface{}) Caller {
	if lutSize < idx1 {
		return nil
	}
	switch info {
	case lutError:
		return lutMonadicError[idx1](fn)
	case lutErrorPtr:
		return lutMonadicErrorPtr[idx1](fn)
	case lutVoid:
		return lutMonadicVoid[idx1](fn)
	default:
		return lutMonadicVoidPtr[idx1](fn)
	}
}

func lutLookupDyadic(info lutInfo, idx1, idx2 reflect.Kind, fn interface{}) Caller {
	if lutSize < idx1 || idx2 > lutSize {
		return nil
	}
	switch info {
	case lutError:
		return lutDyadicError[idx1][idx2](fn)
	case lutErrorPtr:
		return lutDyadicErrorPtr[idx1][idx2](fn)
	case lutVoid:
		return lutDyadicVoid[idx1][idx2](fn)
	default:
		return lutDyadicVoidPtr[idx1][idx2](fn)
	}
}

type lutFunc func(fn interface{}) Caller

var (
	lutMonadicError = [lutSize]lutFunc{
		lutMonadicErrorIdx0,
		lutMonadicErrorIdx1,
		lutMonadicErrorIdx2,
		lutMonadicErrorIdx3,
	}
	lutMonadicErrorPtr = [lutSize]lutFunc{
		lutMonadicErrorPtrIdx0,
		lutMonadicErrorPtrIdx1,
		lutMonadicErrorPtrIdx2,
		lutMonadicErrorPtrIdx3,
	}
	lutMonadicVoid = [lutSize]lutFunc{
		lutMonadicVoidIdx0,
		lutMonadicVoidIdx1,
		lutMonadicVoidIdx2,
		lutMonadicVoidIdx3,
	}
	lutMonadicVoidPtr = [lutSize]lutFunc{
		lutMonadicVoidPtrIdx0,
		lutMonadicVoidPtrIdx1,
		lutMonadicVoidPtrIdx2,
		lutMonadicVoidPtrIdx3,
	}
	lutDyadicVoid = [lutSize][lutSize]lutFunc{
		{
			lutDyadicVoid0x0,
			lutDyadicVoid0x1,
			lutDyadicVoid0x2,
			lutDyadicVoid0x3,
		},
		{
			lutDyadicVoid1x0,
			lutDyadicVoid1x1,
			lutDyadicVoid1x2,
			lutDyadicVoid1x3,
		},
		{
			lutDyadicVoid2x0,
			lutDyadicVoid2x1,
			lutDyadicVoid2x2,
			lutDyadicVoid2x3,
		},
		{
			lutDyadicVoid3x0,
			lutDyadicVoid3x1,
			lutDyadicVoid3x2,
			lutDyadicVoid3x3,
		},
	}
	lutDyadicVoidPtr = [lutSize][lutSize]lutFunc{
		{
			lutDyadicVoidPtr0x0,
			lutDyadicVoidPtr0x1,
			lutDyadicVoidPtr0x2,
			lutDyadicVoidPtr0x3,
		},
		{
			lutDyadicVoidPtr1x0,
			lutDyadicVoidPtr1x1,
			lutDyadicVoidPtr1x2,
			lutDyadicVoidPtr1x3,
		},
		{
			lutDyadicVoidPtr2x0,
			lutDyadicVoidPtr2x1,
			lutDyadicVoidPtr2x2,
			lutDyadicVoidPtr2x3,
		},
		{
			lutDyadicVoidPtr3x0,
			lutDyadicVoidPtr3x1,
			lutDyadicVoidPtr3x2,
			lutDyadicVoidPtr3x3,
		},
	}
	lutDyadicError = [lutSize][lutSize]lutFunc{
		{
			lutDyadicError0x0,
			lutDyadicError0x1,
			lutDyadicError0x2,
			lutDyadicError0x3,
		},
		{
			lutDyadicError1x0,
			lutDyadicError1x1,
			lutDyadicError1x2,
			lutDyadicError1x3,
		},
		{
			lutDyadicError2x0,
			lutDyadicError2x1,
			lutDyadicError2x2,
			lutDyadicError2x3,
		},
		{
			lutDyadicError3x0,
			lutDyadicError3x1,
			lutDyadicError3x2,
			lutDyadicError3x3,
		},
	}
	lutDyadicErrorPtr = [lutSize][lutSize]lutFunc{
		{
			lutDyadicErrorPtr0x0,
			lutDyadicErrorPtr0x1,
			lutDyadicErrorPtr0x2,
			lutDyadicErrorPtr0x3,
		},
		{
			lutDyadicErrorPtr1x0,
			lutDyadicErrorPtr1x1,
			lutDyadicErrorPtr1x2,
			lutDyadicErrorPtr1x3,
		},
		{
			lutDyadicErrorPtr2x0,
			lutDyadicErrorPtr2x1,
			lutDyadicErrorPtr2x2,
			lutDyadicErrorPtr2x3,
		},
		{
			lutDyadicErrorPtr3x0,
			lutDyadicErrorPtr3x1,
			lutDyadicErrorPtr3x2,
			lutDyadicErrorPtr3x3,
		},
	}
)
