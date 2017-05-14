package call

import "testing"

var (
	universeBenchAnalyzer Analyzer
	universeBenchCaller   Caller
)

func BenchmarkAnalyzerCaller(b *testing.B) {
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
