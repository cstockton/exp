package lut

// @TODO add min / max offset to account for high unicode points

// @TODO add a lookup that shifts indexed words over by N to get
// a match. I.E:
// longword
// secondword
// sec
//   ->
//  longword
// secondword
//  sec

//func New(subjects []string) Lookable {
//	analyzers := Analyzers{&LookupMap{}, &LookupAnalyzer{}}
//	lookable, _ := analyzers.Analyze(subjects)
//	return lookable
//}
//
// // New todo
// func New(subjects []string) Lookable {
// 	return &Lookup1{}
// 	// analyzer := &LookupAnalyzer{}
// 	// lookable, _ := analyzer.Analyze(subjects)
// 	// return lookable
// }

// DefaultAnalyzer todo
var DefaultAnalyzer = NewLookableAnalyzer()

// Analyze todo
func Analyze(subjects []string) Lookable {
	return DefaultAnalyzer.Analyze(subjects)
}
