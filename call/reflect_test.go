package call

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

var (
	nilValues = []interface{}{
		(*interface{})(nil), (**interface{})(nil), (***interface{})(nil),
		(func())(nil), (*func())(nil), (**func())(nil), (***func())(nil),
		(chan int)(nil), (*chan int)(nil), (**chan int)(nil), (***chan int)(nil),
		([]int)(nil), (*[]int)(nil), (**[]int)(nil), (***[]int)(nil),
		(map[int]int)(nil), (*map[int]int)(nil), (**map[int]int)(nil),
		(***map[int]int)(nil),
	}
)

func TestIndirect(t *testing.T) {
	type testIndirectCircular *testIndirectCircular
	teq := func(t testing.TB, exp, got interface{}) {
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("DeepEqual failed:\n  exp: %#v\n  got: %#v", exp, got)
		}
	}

	t.Run("Basic", func(t *testing.T) {
		int64v := int64(123)
		int64vp := &int64v
		int64vpp := &int64vp
		int64vppp := &int64vpp
		int64vpppp := &int64vppp
		teq(t, indirect(int64v), int64v)
		teq(t, indirect(int64vp), int64v)
		teq(t, indirect(int64vpp), int64v)
		teq(t, indirect(int64vppp), int64v)
		teq(t, indirect(int64vpppp), int64v)
	})
	t.Run("Nils", func(t *testing.T) {
		for _, n := range nilValues {
			indirect(n)
		}
	})
	t.Run("Circular", func(t *testing.T) {
		var circular testIndirectCircular
		circular = &circular
		teq(t, indirect(circular), circular)
	})
}

func TestOf(t *testing.T) {
	var (
		stringReader = strings.NewReader("foo")
		bytesBuffer  = bytes.NewBuffer(nil)

		ioReader     = (*io.Reader)(nil)
		ioWriter     = (*io.Reader)(nil)
		ioReadWriter = (*io.ReadWriter)(nil)
	)
	type testOf struct {
		arg, param reflect.Type
		errStr     string
	}
	tests := make(map[string][]testOf)
	tests["Positive"] = []testOf{
		{reflect.TypeOf(""), reflect.TypeOf(""), ""},
		{reflect.TypeOf(1), reflect.TypeOf(1), ""},
		{reflect.TypeOf(stringReader), reflect.TypeOf(stringReader), ""},
		{reflect.TypeOf(&stringReader), reflect.TypeOf(&stringReader), ""},
		{reflect.TypeOf(bytesBuffer), reflect.TypeOf(bytesBuffer), ""},
	}
	tests["Invalid"] = []testOf{
		{reflect.TypeOf(nil), reflect.TypeOf(nil), "arg was nil"},
		{reflect.TypeOf(nil), reflect.TypeOf(""), "arg was nil"},
		{reflect.TypeOf(""), reflect.TypeOf(nil), "param was nil"},
		{reflect.TypeOf(1), reflect.TypeOf(""),
			"argument type int did not match param type string"},
		{reflect.TypeOf(ioReader), reflect.TypeOf(bytesBuffer),
			"argument type *io.Reader did not match param type *bytes.Buffer"},
		{reflect.TypeOf(ioReader).Elem(), reflect.TypeOf(bytesBuffer),
			"argument type io.Reader did not match param type *bytes.Buffer"},
		{reflect.TypeOf([]byte("foo")), reflect.TypeOf(`foo`),
			"argument type []uint8 did not match param type string"},
		{reflect.TypeOf(*stringReader), reflect.TypeOf(ioReader).Elem(),
			"argument type strings.Reader did not match param type io.Reader"},
	}
	for _, v := range nilValues {
		tests["Positive"] = append(tests["Positive"], testOf{
			reflect.TypeOf(v), reflect.TypeOf(v), ``})
		tests["Negative"] = append(tests["Negative"], testOf{
			reflect.TypeOf(v), reflect.TypeOf(nil), "param was nil"})
	}
	tests["Interfaces"] = []testOf{
		{reflect.TypeOf(stringReader), reflect.TypeOf(ioReader).Elem(), ""},
		{reflect.TypeOf(bytesBuffer), reflect.TypeOf(ioReader).Elem(), ""},
		{reflect.TypeOf(bytesBuffer), reflect.TypeOf(ioWriter).Elem(), ""},
		{reflect.TypeOf(bytesBuffer), reflect.TypeOf(ioReadWriter).Elem(), ""},
	}

	for name, tcs := range tests {
		chk := func(tc testOf, f func() error) {
			err := f()
			if err != nil {
				if len(tc.errStr) == 0 {
					t.Error(err)
				}
				if !strings.Contains(err.Error(), tc.errStr) {
					t.Errorf("error did not match:\n  exp: %q\n  got: %q", tc.errStr, err)
				}
			} else if len(tc.errStr) > 0 {
				t.Errorf("expected non-nil err of: %v", tc.errStr)
			}
		}

		t.Run(name, func(t *testing.T) {
			t.Run("Param", func(t *testing.T) {
				for _, tc := range tcs {
					chk(tc, func() error {
						return paramOf(tc.arg, tc.param)
					})
				}
			})
			t.Run("Params", func(t *testing.T) {
				t.Run("Single", func(t *testing.T) {
					for _, tc := range tcs {
						chk(tc, func() error {
							return paramsOf([]reflect.Type{tc.arg}, []reflect.Type{tc.param})
						})
					}
				})
				t.Run("Params", func(t *testing.T) {
					var (
						args       []reflect.Type
						params     []reflect.Type
						firstErrTc testOf
					)
					for _, tc := range tcs {
						args = append(args, tc.arg)
						params = append(params, tc.param)
						if len(tc.errStr) > 0 && len(firstErrTc.errStr) == 0 {
							firstErrTc = tc
						}
					}
					chk(firstErrTc, func() error {
						return paramsOf(args, params)
					})
				})
			})
		})
	}
}
