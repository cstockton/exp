package lut

import "sort"

// Lookable todo
type Lookable interface {
	Satisfied([]string) bool
	Lookup(string) int
	Associate(string, int) error
}

// baseHeuristic todo
type baseHeuristic struct {
	analyzed     bool
	minRune      rune
	maxRune      rune
	matrix       [][]rune
	subjects     []string
	subjectsLens []int
}

// Satisfied todo
func (h *baseHeuristic) Satisfied(in []string) bool {
	h.minRune = 0
	h.maxRune = 0
	h.matrix = make([][]rune, 0)
	h.subjects = make([]string, len(in))
	h.subjectsLens = make([]int, 0)

	copy(h.subjects, in)
	sort.Strings(h.subjects)

	for _, subject := range h.subjects {
		h.analyzeLength(subject)
		h.analyzeRunes(subject)
	}
	sort.Ints(h.subjectsLens)

	return true
}

func (h *baseHeuristic) analyzeLength(subject string) {
	h.subjectsLens = append(h.subjectsLens, len(subject))
}

func (h *baseHeuristic) analyzeRunes(subject string) {
	var runes []rune
	for _, c := range subject {
		if h.minRune == 0 {
			h.minRune = c
		}

		runes = append(runes, c)
		if c < h.minRune {
			h.minRune = c
		} else if c > h.maxRune {
			h.maxRune = c
		}
	}
	h.matrix = append(h.matrix, runes)
}

func (h *baseHeuristic) minSubjectLen() int {
	return h.subjectsLens[0]
}

func (h *baseHeuristic) maxSubjectLen() int {
	return h.subjectsLens[len(h.subjectsLens)-1]
}

func (h *baseHeuristic) getUniqColLeft() int {
	for i := 0; i < h.minSubjectLen(); i++ {
		if h.isUniqColLeft(i) {
			return i
		}
	}
	return -1
}

func (h *baseHeuristic) getUniqColRight() int {
	for i := 0; i < h.minSubjectLen(); i++ {
		if h.isUniqColRight(i) {
			return i
		}
	}
	return -1
}

func (h *baseHeuristic) isUniqColLeft(col int) bool {
	seen := make(map[rune]bool)

	for i := 0; i < len(h.matrix); i++ {
		c := h.matrix[i][col]
		if _, ok := seen[c]; ok {
			return false
		}
		seen[c] = true
	}
	return true
}

func (h *baseHeuristic) isUniqColRight(col int) bool {
	seen := make(map[rune]bool)

	for i := 0; i < len(h.matrix); i++ {
		c := h.matrix[i][len(h.matrix[i])-1-col]
		if _, ok := seen[c]; ok {
			return false
		}
		seen[c] = true
	}
	return true
}

// SingleColumnLookup todo
type SingleColumnLookup struct {
	table    []int
	col1     int
	subjects []string
}

// Satisfied todo
func (l *SingleColumnLookup) Satisfied(subjects []string) bool {
	h := &baseHeuristic{}
	if !h.Satisfied(subjects) {
		return false
	}

	l.col1 = h.getUniqColLeft()
	if l.col1 < 0 {
		return false
	}

	l.table = make([]int, 128)
	for i := 0; i < len(l.table); i++ {
		l.table[i] = -1
	}
	for _, subject := range subjects {
		l.Associate(subject, -1)
	}
	return true
}

// Associate todo
func (l *SingleColumnLookup) Associate(s string, i int) error {
	l.table[s[l.col1]] = i
	return nil
}

// Lookup todo
func (l *SingleColumnLookup) Lookup(s string) int {
	return l.table[s[l.col1]]
}

// SingleColumnOffsetLookup todo
type SingleColumnOffsetLookup struct {
	table    []int
	col1     int
	offset   int
	subjects []string
}

// Satisfied todo
func (l *SingleColumnOffsetLookup) Satisfied(subjects []string) bool {
	h := &baseHeuristic{}
	if !h.Satisfied(subjects) {
		return false
	}

	l.col1 = h.getUniqColLeft()
	if l.col1 < 0 {
		return false
	}

	l.table = make([]int, 128)
	for i := 0; i < len(l.table); i++ {
		l.table[i] = -1
	}
	for _, subject := range subjects {
		l.Associate(subject, -1)
	}
	return true
}

// Associate todo
func (l *SingleColumnOffsetLookup) Associate(s string, i int) error {
	l.table[s[l.col1-l.offset]] = i
	return nil
}

// Lookup todo
func (l *SingleColumnOffsetLookup) Lookup(s string) int {
	return l.table[s[l.col1-l.offset]]
}

// SingleColumnLengthOffsetLookup todo
type SingleColumnLengthOffsetLookup struct {
	table    []int
	col1     int
	subjects []string
}

// Satisfied todo
func (l *SingleColumnLengthOffsetLookup) Satisfied(subjects []string) bool {
	h := &baseHeuristic{}
	if !h.Satisfied(subjects) {
		return false
	}

	l.col1 = h.getUniqColLeft()
	if l.col1 < 0 {
		return false
	}
	l.col1--

	l.table = make([]int, 128)
	for i := 0; i < len(l.table); i++ {
		l.table[i] = -1
	}
	for _, subject := range subjects {
		l.Associate(subject, -1)
	}
	return true
}

// Associate todo
func (l *SingleColumnLengthOffsetLookup) Associate(s string, i int) error {
	l.table[s[len(s)-l.col1]] = i
	return nil
}

// Lookup todo
func (l *SingleColumnLengthOffsetLookup) Lookup(s string) int {
	return l.table[s[len(s)-l.col1]]
}

// SingleColumnLeftLookup todo
type SingleColumnLeftLookup struct {
	table    []int
	col1     int
	subjects []string
}

// Satisfied todo
func (l *SingleColumnLeftLookup) Satisfied(subjects []string) bool {
	h := &baseHeuristic{}
	if !h.Satisfied(subjects) {
		return false
	}

	l.col1 = h.getUniqColLeft()
	if l.col1 < 0 {
		return false
	}

	l.table = make([]int, 128)
	for i := 0; i < len(l.table); i++ {
		l.table[i] = -1
	}
	for _, subject := range subjects {
		l.Associate(subject, -1)
	}
	return true
}

// Associate todo
func (l *SingleColumnLeftLookup) Associate(s string, i int) error {
	l.table[s[l.col1]] = i
	return nil
}

// Lookup todo
func (l *SingleColumnLeftLookup) Lookup(s string) int {
	return l.table[s[l.col1]]
}

type tableEntry struct {
	index int
	value string
}

// SingleColumnLeftUnknownLookup todo
type SingleColumnLeftUnknownLookup struct {
	table    []tableEntry
	col1     int
	subjects []string
}

// Satisfied todo
func (l *SingleColumnLeftUnknownLookup) Satisfied(subjects []string) bool {
	h := &baseHeuristic{}
	if !h.Satisfied(subjects) {
		return false
	}

	l.col1 = h.getUniqColLeft()
	if l.col1 < 0 {
		return false
	}

	l.table = make([]tableEntry, 128)
	for i := 0; i < len(l.table); i++ {
		l.table[i] = tableEntry{-1, ""}
	}
	for _, subject := range subjects {
		l.Associate(subject, -1)
	}
	return true
}

// Associate todo
func (l *SingleColumnLeftUnknownLookup) Associate(s string, i int) error {
	l.table[s[l.col1]] = tableEntry{i, s}
	return nil
}

// Lookup todo
func (l *SingleColumnLeftUnknownLookup) Lookup(s string) int {
	if res := l.table[s[l.col1]]; res.value == s {
		return res.index
	}
	return -1
}

// LookupKnown todo
func (l *SingleColumnLeftUnknownLookup) LookupKnown(s string) int {
	return l.table[s[l.col1]].index
}

// SingleColumnRightLookup todo
type SingleColumnRightLookup struct {
	table    []int
	col1     int
	subjects []string
}

// Satisfied todo
func (l *SingleColumnRightLookup) Satisfied(subjects []string) bool {
	h := &baseHeuristic{}
	if !h.Satisfied(subjects) {
		return false
	}

	l.col1 = h.getUniqColRight()
	if l.col1 < 0 {
		return false
	}

	l.table = make([]int, 128)
	for i := 0; i < len(l.table); i++ {
		l.table[i] = -1
	}
	for _, subject := range subjects {
		l.Associate(subject, -1)
	}
	return true
}

// Associate todo
func (l *SingleColumnRightLookup) Associate(s string, i int) error {
	l.table[s[len(s)-1-l.col1]] = i
	return nil
}

// Lookup todo
func (l *SingleColumnRightLookup) Lookup(s string) int {
	return l.table[s[len(s)-1-l.col1]]
}

// DoubleColumnLeftLookup todo
type DoubleColumnLeftLookup struct {
	table    []int
	col1     int
	subjects []string
}

// Satisfied todo
func (l *DoubleColumnLeftLookup) Satisfied(subjects []string) bool {
	h := &baseHeuristic{}
	if !h.Satisfied(subjects) {
		return false
	}

	l.col1 = h.getUniqColLeft()
	if l.col1 < 0 {
		return false
	}

	l.table = make([]int, 128)
	for i := 0; i < len(l.table); i++ {
		l.table[i] = -1
	}
	for _, subject := range subjects {
		l.Associate(subject, -1)
	}
	return true
}

// Associate todo
func (l *DoubleColumnLeftLookup) Associate(s string, i int) error {
	l.table[s[l.col1]] = i
	return nil
}

// Lookup todo
func (l *DoubleColumnLeftLookup) Lookup(s string) int {
	return l.table[s[l.col1]]
}

// MapLookup todo
type MapLookup struct {
	table map[string]int
}

// Satisfied todo
func (l *MapLookup) Satisfied(subjects []string) bool {
	l.table = make(map[string]int)

	for _, subject := range subjects {
		l.table[subject] = -1
	}
	return true
}

// Associate todo
func (l *MapLookup) Associate(subject string, i int) error {
	l.table[subject] = i
	return nil
}

// Lookup todo
func (l *MapLookup) Lookup(subject string) int {
	if val, ok := l.table[subject]; ok {
		return val
	}
	return -1
}
