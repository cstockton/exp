package lut

// Analyzer todo
type Analyzer interface {
	Analyze([]string) Lookable
}

// LookableAnalyzer todo
type LookableAnalyzer struct {
	f func() []Lookable
}

// NewLookableAnalyzer todo
func NewLookableAnalyzer() *LookableAnalyzer {
	return &LookableAnalyzer{}
}

// Analyze todo
func (a LookableAnalyzer) Analyze(subjects []string) Lookable {
	defaultLookup := &MapLookup{}
	lookables := []Lookable{
		&SingleColumnLeftLookup{},
		&SingleColumnRightLookup{},
		&DoubleColumnLeftLookup{},
		defaultLookup,
	}
	for _, lookable := range lookables {
		if lookable.Satisfied(subjects) {
			return lookable
		}
	}
	return defaultLookup
}

/*
// LookableAnalyzer todo
type LookableAnalyzer struct {
	lookables []Lookable
}

// NewLookableAnalyzer todo
func NewLookableAnalyzer(lookables ...Lookable) *LookableAnalyzer {
	return &LookableAnalyzer{
		lookables: lookables,
	}
}

// Analyze todo
func (a LookableAnalyzer) Analyze(subjects []string) Lookable {
	for _, lookable := range a.lookables {
		if lookable.Satisfied(subjects) {
			return lookable
		}
	}
	return &LookupMap{}
}


*/

//

/*

// Analyzer todo
type Analyzer interface {
	Heuristic
	Lookable() Lookable
}
// Analyzer todo
type Analyzer interface {
	Analyze([]string) (lookable Lookable, viability float64)
}

// Analyzers todo
type Analyzers []Analyzer

// Analyze todo
func (a Analyzers) Analyze(subjects []string) (Lookable, float64) {
	var lookable Lookable
	mostViable := math.SmallestNonzeroFloat64

	for _, analyzer := range a {
		candidate, viability := analyzer.Analyze(subjects)

		if viability >= mostViable {
			mostViable = viability
			lookable = candidate
		}
	}

	return lookable, mostViable
}

// HeuristicsAnalyzer todo
type HeuristicsAnalyzer struct {
	heuristics Heuristics
}

/*

// Analyze todo
func (a *HeuristicsAnalyzer) Analyze(subjects []string) (lookable Lookable, viability float64) {
	var lookable Lookable
	mostViable := math.SmallestNonzeroFloat64

	for _, analyzer := range a {
		candidate, viability := analyzer.Analyze(subjects)

		if viability >= mostViable {
			mostViable = viability
			lookable = candidate
		}
	}

	return lookable, mostViable

}

// HeuristicAnalyzer todo
type HeuristicAnalyzer struct {
	heuristics Heuristics
	lookable   func() Lookable
}

// Analyze todo
func (a *HeuristicAnalyzer) Analyze(subjects []string) (lookable Lookable, viability float64) {
	return
}

// Lookup1Analyzer todo
var Lookup1Analyzer = &HeuristicAnalyzer{
	heuristics: Heuristics{
		[]Heuristics{
			&columnarHeuristic{},
		},
	},
	lookable: func() Lookable {
		return &Lookup1{}
	},
}
*/
