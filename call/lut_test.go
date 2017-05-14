package call

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLutIndexFromKind(t *testing.T) {
	type testLutIndexFromKind struct {
		exp  reflect.Kind
		give []reflect.Kind
	}
	tests := []testLutIndexFromKind{
		{0, []reflect.Kind{
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64}},
		{1, []reflect.Kind{
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64}},
		{2, []reflect.Kind{
			reflect.Complex128, reflect.Complex64, reflect.Float64, reflect.Float32,
			reflect.Uintptr}},
		{2, []reflect.Kind{
			reflect.Array, reflect.Chan, reflect.Func, reflect.Interface,
			reflect.Map, reflect.Ptr}},
		{3, []reflect.Kind{
			reflect.Bool, reflect.Invalid}},
		{3, []reflect.Kind{
			reflect.Slice, reflect.String, reflect.Struct, reflect.UnsafePointer}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%v==%v", tc.give, tc.exp), func(t *testing.T) {
			for _, k := range tc.give {
				if got := lutIndexFromKind(k); tc.exp != got {
					t.Fatalf("expected %v; got: %v", uint(tc.exp), uint(got))
				}
			}
		})
	}
}
