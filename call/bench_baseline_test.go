package call_test

import "testing"

func funcMonadicInt(v int) error {
	return nil
}

func funcMonadicInterface(v interface{}) error {
	return nil
}

func funcVariadicInt(v ...int) error {
	return nil
}

func funcVariadicInterface(v ...interface{}) error {
	return nil
}

type monadicIntCaller interface {
	Call(in int) error
}

type monadicInterfaceCaller interface {
	Call(in interface{}) error
}

type variadicIntCaller interface {
	Call(in ...int) error
}

type variadicInterfaceCaller interface {
	Call(in ...interface{}) error
}

type monadicIntCallerStruct struct{}

func (c monadicIntCallerStruct) Call(in int) error {
	return nil
}

type monadicInterfaceCallerStruct struct{}

func (c monadicInterfaceCallerStruct) Call(in interface{}) error {
	return nil
}

type variadicIntCallerStruct struct{}

func (c variadicIntCallerStruct) Call(in ...int) error {
	return nil
}

type variadicInterfaceCallerStruct struct{}

func (c variadicInterfaceCallerStruct) Call(in ...interface{}) error {
	return nil
}

var (
	mMonadicInt                                      = monadicIntCallerStruct{}
	mMonadicIntCaller        monadicIntCaller        = monadicIntCallerStruct{}
	mMonadicInterface                                = monadicInterfaceCallerStruct{}
	mMonadicInterfaceCaller  monadicInterfaceCaller  = monadicInterfaceCallerStruct{}
	mVariadicInt                                     = variadicIntCallerStruct{}
	mVariadicIntCaller       variadicIntCaller       = variadicIntCallerStruct{}
	mVariadicInterface                               = variadicInterfaceCallerStruct{}
	mVariadicInterfaceCaller variadicInterfaceCaller = variadicInterfaceCallerStruct{}
)

// Monadic/Int/FnUniverse-24              2000000000               0.42 ns/op
// Monadic/Int/MethodUniverse-24          2000000000               0.41 ns/op
// Monadic/Int/MethodCallerUniverse-24            300000000                4.23 ns/op
// Monadic/Int/FnParent-24                        500000000                2.58 ns/op
// Monadic/Int/FnLocal-24                         500000000                2.61 ns/op
// Monadic/Interface/FnUniverse-24                200000000                8.40 ns/op
// Monadic/Interface/MethodUniverse-24            100000000               10.0 ns/op
// Monadic/Interface/MethodCallerUniverse-24      20000000                58.2 ns/op
// Monadic/Interface/FnParent-24                  20000000                58.0 ns/op
// Monadic/Interface/FnLocal-24                   20000000                59.9 ns/op
// Variadic/Int/FnUniverse-24                     500000000                2.93 ns/op
// Variadic/Int/MethodUniverse-24                 300000000                3.62 ns/op
// Variadic/Int/MethodCallerUniverse-24           20000000                51.8 ns/op
// Variadic/Int/FnParent-24                       30000000                43.4 ns/op
// Variadic/Int/FnLocal-24                        50000000                42.9 ns/op
// Variadic/Interface/FnUniverse-24               100000000               10.8 ns/op
// Variadic/Interface/MethodUniverse-24           100000000               10.6 ns/op
// Variadic/Interface/MethodCallerUniverse-24     10000000               135 ns/op
// Variadic/Interface/FnParent-24                 10000000               131 ns/op
// Variadic/Interface/FnLocal-24                  10000000               131 ns/op
func BenchmarkBaselineMonadic(b *testing.B) {
	b.SkipNow()

	fnParam1 := 123
	fnMonadicInt := func(v int) error {
		return nil
	}
	fnMonadicInterface := func(v interface{}) error {
		return nil
	}

	b.Run("Int", func(b *testing.B) {
		b.Run("FnUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				funcMonadicInt(fnParam1)
			}
		})
		b.Run("MethodUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mMonadicInt.Call(fnParam1)
			}
		})
		b.Run("MethodCallerUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mMonadicIntCaller.Call(fnParam1)
			}
		})
		b.Run("FnParent", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fnMonadicInt(fnParam1)
			}
		})
		b.Run("FnLocal", func(b *testing.B) {
			localMonadicInt := func(v int) error {
				return nil
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				localMonadicInt(fnParam1)
			}
		})
	})

	b.Run("Interface", func(b *testing.B) {
		b.Run("FnUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				funcMonadicInterface(fnParam1)
			}
		})
		b.Run("MethodUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mMonadicInterface.Call(fnParam1)
			}
		})
		b.Run("MethodCallerUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mMonadicInterfaceCaller.Call(fnParam1)
			}
		})
		b.Run("FnParent", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fnMonadicInterface(fnParam1)
			}
		})
		b.Run("FnLocal", func(b *testing.B) {
			localMonadicInterface := func(v interface{}) error {
				return nil
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				localMonadicInterface(fnParam1)
			}
		})
	})
}

func BenchmarkBaselineVariadic(b *testing.B) {
	b.SkipNow()

	fnParam1 := 123
	fnVariadicInt := func(v ...int) error {
		return nil
	}
	fnVariadicInterface := func(in ...interface{}) error {
		return nil
	}

	b.Run("Int", func(b *testing.B) {
		b.Run("FnUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				funcVariadicInt(fnParam1)
			}
		})
		b.Run("MethodUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mVariadicInt.Call(fnParam1)
			}
		})
		b.Run("MethodCallerUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mVariadicIntCaller.Call(fnParam1)
			}
		})
		b.Run("FnParent", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fnVariadicInt(fnParam1)
			}
		})
		b.Run("FnLocal", func(b *testing.B) {
			localVariadicInt := func(v ...int) error {
				return nil
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				localVariadicInt(fnParam1)
			}
		})
	})

	b.Run("Interface", func(b *testing.B) {
		b.Run("FnUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				funcVariadicInterface(fnParam1)
			}
		})
		b.Run("MethodUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mVariadicInterface.Call(fnParam1)
			}
		})
		b.Run("MethodCallerUniverse", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mVariadicInterfaceCaller.Call(fnParam1)
			}
		})
		b.Run("FnParent", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fnVariadicInterface(fnParam1)
			}
		})
		b.Run("FnLocal", func(b *testing.B) {
			localVariadicInterface := func(v ...interface{}) error {
				return nil
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				localVariadicInterface(fnParam1)
			}
		})
	})
}
