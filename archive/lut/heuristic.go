package lut

// Heuristic todo
type Heuristic interface {
	Satisfied([]string) bool
}

// Heuristics todo
type Heuristics []Heuristic

// NewHeuristics todo
func NewHeuristics(heuristics ...Heuristic) (out Heuristics) {
	return append(out, heuristics...)
}

// Satisfied todo
func (h Heuristics) Satisfied(subjects []string) bool {
	for _, heuristic := range h {
		if !heuristic.Satisfied(subjects) {
			return false
		}
	}
	return true
}

/*

// Analyze todo
func (a *LookupAnalyzer) Analyze(subjects []string) (lookable Lookable, viability float64) {
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


// Analyze todo
func (a *LookupAnalyzer) Analyze(subjects []string) (lookable Lookable, viability float64) {
	sort.Strings(subjects)
	a.analyzeSubjects(subjects)

	for i := 0; i < a.minSubjectLen(); i++ {
		if a.isUniqCol(i) {
			return NewLookup1(subjects, i), 100
		}
	}

	log.Printf("LookupAnalyzer: %#v\n", a)
	return NewLookupMap(subjects), -1
}

func (a *LookupAnalyzer) isUniqCol(col int) bool {
	seen := make(map[rune]bool)

	for i := 0; i < len(a.matrix); i++ {
		c := a.matrix[i][col]
		if _, ok := seen[c]; ok {
			return false
		}
		seen[c] = true
	}
	return true
}

func (a *LookupAnalyzer) minSubjectLen() int {
	return a.subjectLens[0]
}

func (a *LookupAnalyzer) maxSubjectLen() int {
	return a.subjectLens[len(a.subjectLens)-1]
}

func (a *LookupAnalyzer) analyzeSubjects(subjects []string) {
	a.subjectMap = make(map[int]int)

	for _, subject := range subjects {
		a.analyzeSubject(subject)
	}
	sort.Ints(a.subjectLens)
}

func (a *LookupAnalyzer) analyzeSubject(subject string) {
	a.subjectLens = append(a.subjectLens, len(subject))
	var runes []rune
	for _, c := range subject {
		runes = append(runes, c)
	}
	a.matrix = append(a.matrix, runes)
}
*/
