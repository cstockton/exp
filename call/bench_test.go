package call

import "testing"

var (
	universeBenchAnalyzer Analyzer
	universeBenchCaller   Caller
)

func BenchmarkCalling(b *testing.B) {
	b.Run("Monadic", func(b *testing.B) {
		var (
			a Analyzer
			c Caller
		)
		tc := testLutLookupsCases[len(testLutLookupsCases)/2]
		caller, err := a.Analyze(tc.fn)
		if err != nil {
			b.Fatal(err)
		}
		if caller == nil {
			b.Fatal("caller was nil")
		}
		c = caller
		universeBenchCaller = c
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := c.Call(tc.arg1); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("Dyadic", func(b *testing.B) {
		var (
			a Analyzer
			c Caller
		)
		tc := testLutLookupsCases2x[len(testLutLookupsCases2x)/2]
		caller, err := a.Analyze(tc.fn)
		if err != nil {
			b.Fatal(err)
		}
		if caller == nil {
			b.Fatal("caller was nil")
		}
		c = caller
		universeBenchCaller = c
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := c.Call(tc.arg1, tc.arg2); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkAnalyzerCaller(b *testing.B) {
	b.Run("Monadic", func(b *testing.B) {
		var (
			a Analyzer
			c Caller
		)
		tc := testLutLookupsCases[len(testLutLookupsCases)/2]
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			caller, err := a.Analyze(tc.fn)
			if err != nil {
				b.Fatal(err)
			}
			if caller == nil {
				b.Fatal("caller was nil")
			}
			c = caller
		}
		universeBenchCaller = c
	})
	b.Run("Monadics", func(b *testing.B) {
		var (
			a Analyzer
			c Caller
		)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, tc := range testLutLookupsCases {
				caller, err := a.Analyze(tc.fn)
				if err != nil {
					b.Fatal(err)
				}
				if caller == nil {
					b.Fatal("caller was nil")
				}
				c = caller
			}
		}
		universeBenchCaller = c
	})
	b.Run("Dyadic", func(b *testing.B) {
		var (
			a Analyzer
			c Caller
		)
		tc := testLutLookupsCases2x[len(testLutLookupsCases2x)/2]
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {

			caller, err := a.Analyze(tc.fn)
			if err != nil {
				b.Fatal(err)
			}
			if caller == nil {
				b.Fatal("caller was nil")
			}
			c = caller
		}
		universeBenchCaller = c
	})
	b.Run("Dyadics", func(b *testing.B) {
		var (
			a Analyzer
			c Caller
		)
		for i := 0; i < b.N; i++ {
			for _, tc := range testLutLookupsCases2x {
				caller, err := a.Analyze(tc.fn)
				if err != nil {
					b.Fatal(err)
				}
				if caller == nil {
					b.Fatal("caller was nil")
				}
				c = caller
			}
		}
		universeBenchCaller = c
	})
}

func BenchmarkAnalyzerAnalyze(b *testing.B) {
	b.Run("Monadics", func(b *testing.B) {
		var (
			a Analyzer
			c Caller
		)
		for i := 0; i < b.N; i++ {
			for _, tc := range testLutLookupsCases {
				caller, err := a.Analyze(tc.fn)
				if err != nil {
					b.Fatal(err)
				}
				if caller == nil {
					b.Fatal("caller was nil")
				}
				c = caller
			}
		}
		universeBenchCaller = c
	})
	b.Run("Dyadics", func(b *testing.B) {
		var (
			a Analyzer
			c Caller
		)
		for i := 0; i < b.N; i++ {
			for _, tc := range testLutLookupsCases2x {
				caller, err := a.Analyze(tc.fn)
				if err != nil {
					b.Fatal(err)
				}
				if caller == nil {
					b.Fatal("caller was nil")
				}
				c = caller
			}
		}
		universeBenchCaller = c
	})
}
