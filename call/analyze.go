package call

import (
	"fmt"
	"reflect"
)

var (
	typeOfError     = reflect.TypeOf((*error)(nil)).Elem()
	defaultAnalyzer = &Analyzer{}
)

// Analyzer will analyze functions to find the best implementation of a Caller.
type Analyzer struct {
	AllowNiladic bool
}

// Analyze will return the best implementation of a Caller for the given func.
func (a Analyzer) Analyze(fn interface{}) (Caller, error) {
	fnVal := reflect.ValueOf(indirect(fn))
	if !fnVal.IsValid() {
		return nil, fmt.Errorf("call: value type %T is invalid", fn)
	}

	fnTyp := fnVal.Type()
	if fnTyp.Kind() != reflect.Func || fnVal.IsNil() {
		return nil, fmt.Errorf("call: value type %T is not a non-nil function", fn)
	}

	var li lutInfo
	outLen := fnTyp.NumOut()
	if outLen > 0 {
		if outLen != 1 || !fnTyp.Out(0).Implements(typeOfError) {
			return nil, fmt.Errorf("call: %T does not return a single error type", fn)
		}
		li = lutError
	}

	arity := fnTyp.NumIn()
	if !fnTyp.IsVariadic() {
		switch arity {
		case 0:
			return a.analyzeNiladic(fn)
		case 1:
			return a.analyzeMonadic(fn, li, fnVal, fnTyp)
		case 2:
			return a.analyzeDyadic(fn, li, fnVal, fnTyp)
		}
	}
	return a.analyzeVariadic(fn, arity, fnVal, fnTyp)
}

func (a Analyzer) analyzeNiladic(fn interface{}) (Caller, error) {
	if !a.AllowNiladic {
		return nil, fmt.Errorf("call: %T is niladic", fn)
	}
	if fn, ok := fn.(func() error); ok {
		return niladicErrorCallerFunc(fn), nil
	}
	if fn, ok := fn.(func()); ok {
		return niladicCallerFunc(fn), nil
	}
	return nil, fmt.Errorf("call: func %v (%[1]T) was not a compatible niladic func", fn)
}

func (a Analyzer) analyzeMonadic(fn interface{}, li lutInfo, fnVal reflect.Value, fnTyp reflect.Type) (Caller, error) {
	in1 := fnTyp.In(0)
	k1 := in1.Kind()
	if k1 == reflect.Ptr {
		k1 = in1.Elem().Kind()
		li++ // (const ordering)
	}
	ix1 := lutIndexFromKind(k1)
	if caller := lutLookupMonadic(li, ix1, fn); nil != caller {
		return caller, nil
	}
	return &monadicReflectCaller{
		fnVal: fnVal,
		inTyp: in1,
	}, nil
}

func (a Analyzer) analyzeDyadic(fn interface{}, li lutInfo, fnVal reflect.Value, fnTyp reflect.Type) (Caller, error) {
	in1, in2 := fnTyp.In(0), fnTyp.In(1)
	k1, k2 := in1.Kind(), in2.Kind()
	if k1 == reflect.Ptr {
		k1 = in1.Elem().Kind()
		li++
	}
	if k2 == reflect.Ptr {
		k2 = in2.Elem().Kind()
		if lutErrorPtr != li && li != lutVoidPtr {
			li++
		}
	}
	ix1, ix2 := lutIndexFromKind(k1), lutIndexFromKind(k2)
	if caller := lutLookupDyadic(li, ix1, ix2, fn); nil != caller {
		return caller, nil
	}
	return &dyadicReflectCaller{
		fnVal:  fnVal,
		inTyp1: in1,
		inTyp2: in2,
	}, nil
}

func (a Analyzer) analyzeVariadic(fn interface{}, arity int, fnVal reflect.Value, fnTyp reflect.Type) (Caller, error) {
	caller := &variadicReflectCaller{
		fnVar:  fnTyp.IsVariadic(),
		fnVal:  fnVal,
		inTyps: make([]reflect.Type, arity),
	}
	for i := 0; i < arity; i++ {
		caller.inTyps[i] = fnTyp.In(i)
	}
	if caller.fnVar {
		caller.vrTyp = caller.inTyps[arity-1].Elem()
	}
	return caller, nil
}
