package call

import (
	"fmt"
	"reflect"
)

type niladicCallerFunc func()

func (c niladicCallerFunc) Call(in ...interface{}) error {
	c()
	return nil
}

type niladicErrorCallerFunc func() error

func (c niladicErrorCallerFunc) Call(in ...interface{}) error {
	return c()
}

type monadicReflectCaller struct {
	fnVal reflect.Value
	inTyp reflect.Type
}

func (c *monadicReflectCaller) Call(in ...interface{}) error {
	inLen := len(in)
	if inLen != 1 {
		return fmt.Errorf("call: expected 1 arguments, got %d", inLen)
	}

	argVal := reflect.ValueOf(in[0])
	if !argVal.IsValid() {
		return fmt.Errorf("call: argument 1 (%v) was invalid", argVal)
	}

	argTyp := argVal.Type()
	if err := paramOf(argTyp, c.inTyp); err != nil {
		return err
	}

	out := c.fnVal.Call([]reflect.Value{argVal})
	if len(out) == 0 {
		return nil
	}

	errVal := out[0]
	if errVal.IsValid() && errVal.CanInterface() {
		errIfc := errVal.Interface()
		if err, ok := errIfc.(error); ok {
			return err
		}
	}
	return nil
}

type dyadicReflectCaller struct {
	fnVal  reflect.Value
	inTyp1 reflect.Type
	inTyp2 reflect.Type
}

func (c *dyadicReflectCaller) Call(in ...interface{}) error {
	inLen := len(in)
	if inLen != 2 {
		return fmt.Errorf("call: expected 2 arguments, got %d", inLen)
	}

	argVal1 := reflect.ValueOf(in[0])
	if !argVal1.IsValid() {
		return fmt.Errorf("call: argument 1 (%v) was invalid", argVal1)
	}
	argVal2 := reflect.ValueOf(in[0])
	if !argVal2.IsValid() {
		return fmt.Errorf("call: argument 2 (%v) was invalid", argVal2)
	}

	argTyp1 := argVal1.Type()
	if err := paramOf(argTyp1, c.inTyp1); err != nil {
		return err
	}
	argTyp2 := argVal2.Type()
	if err := paramOf(argTyp2, c.inTyp2); err != nil {
		return err
	}

	out := c.fnVal.Call([]reflect.Value{argVal1, argVal2})
	if len(out) == 0 {
		return nil
	}

	errVal := out[0]
	if errVal.IsValid() && errVal.CanInterface() {
		errIfc := errVal.Interface()
		if err, ok := errIfc.(error); ok {
			return err
		}
	}
	return nil

}

// variadicReflectCaller covers polyadic functions too, it's really slow and
// only used if no other heuristics match.
type variadicReflectCaller struct {
	fnVar  bool
	fnVal  reflect.Value
	vrTyp  reflect.Type
	inTyps []reflect.Type
}

func (c *variadicReflectCaller) Call(in ...interface{}) error {
	inLen := len(in)
	expLen := len(c.inTyps)
	if !c.fnVar && inLen != expLen {
		return fmt.Errorf("call: expected %d arguments, got %d", expLen, inLen)
	} else if c.fnVar {
		expLen--
		if inLen < expLen {
			return fmt.Errorf("call: expected at least %d arguments, got %d", expLen-1, inLen)
		}
	}

	var callArgs []reflect.Value
	for i, arg := range in {
		argVal := reflect.ValueOf(arg)
		if !argVal.IsValid() {
			return fmt.Errorf("call: argument %d (%v) was invalid", i, argVal)
		}

		argTyp := argVal.Type()
		var expTyp reflect.Type
		if i >= expLen {
			expTyp = c.vrTyp
		} else {
			expTyp = c.inTyps[i]
		}
		if err := paramOf(argTyp, expTyp); err != nil {
			return err
		}
		callArgs = append(callArgs, argVal)
	}

	var out []reflect.Value
	if c.fnVar {
		variadics := callArgs[expLen:]
		variadicVal := reflect.Zero(c.inTyps[len(c.inTyps)-1])

		if len(variadics) > 0 {
			variadicVal = reflect.Append(variadicVal, variadics...)
		}
		callArgs = append(callArgs[:expLen], variadicVal)

		out = c.fnVal.CallSlice(append(callArgs[:expLen], variadicVal))
	} else {
		out = c.fnVal.Call(callArgs)
	}
	if len(out) == 0 {
		return nil
	}

	errVal := out[0]
	if errVal.IsValid() && errVal.CanInterface() {
		errIfc := errVal.Interface()
		if err, ok := errIfc.(error); ok {
			return err
		}
	}
	return nil
}

func paramOf(arg, param reflect.Type) error {
	if arg == nil {
		return fmt.Errorf("call: arg was nil")
	}
	if param == nil {
		return fmt.Errorf("call: param was nil")
	}
	if arg.AssignableTo(param) {
		return nil
	}
	return fmt.Errorf("call: argument type %v did not match param type %v", arg, param)
}

func paramsOf(args, sig []reflect.Type) error {
	if args == nil {
		return fmt.Errorf("call: args was nil")
	}
	if sig == nil {
		return fmt.Errorf("call: sig was nil")
	}
	if len(args) != len(sig) {
		return fmt.Errorf("call: signature mismatch")
	}
	for i := range args {
		if err := paramOf(args[i], sig[i]); err != nil {
			return fmt.Errorf("call: faield at index %d, %v", i, err)
		}
	}
	return nil
}

// indirect will perform recursive indirection on the given value. It should
// never panic and will return a value unless indirection is impossible due to
// infinite recursion in cases like `type Element *Element`.
func indirect(value interface{}) interface{} {

	// Just to be safe, recursion should not be possible but I may be
	// missing an edge case.
	for {

		val := reflect.ValueOf(value)
		if !val.IsValid() || val.Kind() != reflect.Ptr {
			// Value is not a pointer.
			return value
		}

		res := reflect.Indirect(val)
		if !res.IsValid() || !res.CanInterface() {
			// Invalid value or can't be returned as interface{}.
			return value
		}

		// Test for a circular type.
		if res.Kind() == reflect.Ptr && val.Pointer() == res.Pointer() {
			return value
		}

		// Next round.
		value = res.Interface()
	}
}
