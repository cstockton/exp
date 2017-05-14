package call

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

var (
	testLutLookupsCallTable = make(map[string]func(interface{}) error)
	errTestLutLookupFailed  = errors.New("lut lookup returned wrong func")
)

type testLutLookups struct {
	idx      int
	arg1, fn interface{}
	name     string
}

type testLutLookups2x struct {
	idx1           int
	idx2           int
	arg1, arg2, fn interface{}
	name           string
}

func TestLutLookups(t *testing.T) {
	var a Analyzer
	for _, tc := range testLutLookupsCases {
		t.Run(tc.name, func(t *testing.T) {
			caller, err := a.Analyze(tc.fn)
			if err != nil {
				t.Fatal(err)
			}
			if caller == nil {
				t.Fatal("caller was nil")
			}
			caller.Call(tc.arg1)
			validateFn, ok := testLutLookupsCallTable[tc.name]
			if !ok {
				t.Fatal("caller was not ran")
			}
			err = validateFn(caller)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
	for _, tc := range testLutLookupsCases2x {
		t.Run(tc.name, func(t *testing.T) {
			caller, err := a.Analyze(tc.fn)
			if err != nil {
				t.Fatal(err)
			}
			if caller == nil {
				t.Fatal("caller was nil")
			}
			caller.Call(tc.arg1, tc.arg2)
			validateFn, ok := testLutLookupsCallTable[tc.name]
			if !ok {
				t.Fatal("caller was not ran")
			}
			err = validateFn(caller)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

type callerTest struct {
	t              *testing.T
	factory        func(interface{}) (Caller, error)
	giveIn, expOut []interface{}
	errNewStr      string
	errOutStr      string
	expCalls       int
	tr             *int
}

func (c *callerTest) in(in ...interface{}) *callerTest {
	c.giveIn = in
	return c
}

func (c *callerTest) out(out ...interface{}) *callerTest {
	c.expOut = out
	return c
}

func (c *callerTest) errNew(expStr string) *callerTest {
	c.errNewStr = expStr
	return c
}

func (c *callerTest) err(expStr string) *callerTest {
	c.errOutStr = expStr
	return c
}

func (c *callerTest) inc() *callerTest {
	c.expCalls++
	return c
}

func (c *callerTest) run(f interface{}) {
	c.t.Run(c.String(), func(t *testing.T) {
		caller, err := c.factory(f)
		if err != nil {
			if len(c.errNewStr) == 0 {
				t.Fatalf("factory err: \n%v", err)
			}
			if !strings.Contains(err.Error(), c.errNewStr) {
				t.Fatalf("factory err: \n%q != %q", err, c.errNewStr)
			}
			t.Log("failure was expected:", err)
			return
		}
		if nil == caller {
			t.Fatalf("unexpected nil caller with no error")
		}
		err = caller.Call(c.giveIn...)
		if err != nil {
			if len(c.errOutStr) == 0 {
				t.Fatalf("call err: \n%v", err)
			}
			if !strings.Contains(err.Error(), c.errOutStr) {
				t.Fatalf("call err: \n%q != %q", err, c.errOutStr)
			}
		}
		if *c.tr != c.expCalls {
			t.Fatalf("expected calls mismatch: \n%d != %d", *c.tr, c.expCalls)
		}
	})
}

func (c *callerTest) String() string {
	if len(c.errOutStr) > 0 {
		return fmt.Sprintf("CallerTest(%T -> %T -(err)-> %v)",
			c.giveIn, c.expOut, c.errOutStr)
	}
	return fmt.Sprintf("CallerTest(%T -> %T (ok))", c.giveIn, c.expOut)
}

func TestCaller(t *testing.T) {
	i := new(int)
	tc := func(t *testing.T) *callerTest {
		*i = 0
		return &callerTest{t: t, factory: New, tr: i, expCalls: 1}
	}

	t.Run("Basic", func(t *testing.T) {
		f := func(a string, b int) {
			*i++
		}
		tc(t).in("A", 2).run(f)
	})

	t.Run("Variadic", func(t *testing.T) {
		f := func(a string, b int, y ...bool) {
			*i++
		}
		tc(t).in("A", 2).run(f)
		tc(t).in("A", 2).run(f)
		tc(t).in("A", 2, true).run(f)
		tc(t).in("A", 2, true, false).run(f)
		tc(t).in("A", 2, true, false, true).run(f)
	})
	t.Run("Empty", func(t *testing.T) {
		_, err := New(func() error {
			return nil
		})
		if err == nil {
			t.Fatal(err)
		}
	})
	t.Run("EmptyAllowNiladic", func(t *testing.T) {
		var a Analyzer
		a.AllowNiladic = true
		caller, err := a.Analyze(func() error {
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		err = caller.Call()
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("One", func(t *testing.T) {
		caller, err := New(func(k int) error {
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		err = caller.Call(1234)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Two", func(t *testing.T) {
		caller, err := New(func(k int, v string) error {
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		err = caller.Call(1234, "fooo")
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("VariadicNone", func(t *testing.T) {
		caller, err := New(func(k int, v ...string) error {
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		err = caller.Call(1234)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("VariadicOne", func(t *testing.T) {
		caller, err := New(func(k int, v ...string) error {
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		err = caller.Call(1234, "arg1")
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("VariadicTwo", func(t *testing.T) {
		caller, err := New(func(k int, v ...string) error {
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		err = caller.Call(1234, "arg1", "arg2")
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("VariadicThree", func(t *testing.T) {
		caller, err := New(func(k int, v ...string) error {
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		err = caller.Call(1234, "arg1", "arg2", "arg3")
		if err != nil {
			t.Fatal(err)
		}
	})
}
