package lut

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

type tbCase struct {
	tb testing.TB
}

func (c *tbCase) bind(tb testing.TB) *tbCase {
	c.tb = tb
	return c
}

func (c *tbCase) must(err error) *tbCase {
	if err != nil {
		c.tb.Fatal(err)
	}
	return c
}

func (c *tbCase) check(err error) *tbCase {
	if err != nil {
		c.tb.Error(err)
	}
	return c
}

func (c *tbCase) bail() *tbCase {
	c.tb.FailNow()
	return c
}

func (c *tbCase) expect(expect, got interface{}) *tbCase {
	if !reflect.DeepEqual(expect, got) {
		c.tb.Fatalf("DeepEquals() %#v != %#v", expect, got)
	}
	return c
}

type tCase struct {
	tbCase
	t *testing.T
}

func (c *tCase) bind(t *testing.T) *tCase {
	c.tbCase.bind(t)
	c.t = t
	return c
}

type bCase struct {
	tbCase
	b *testing.B
}

func (c *bCase) bind(b *testing.B) *bCase {
	c.tbCase.bind(b)
	c.b = b
	return c
}

type testData struct {
	expect   string
	subjects []string
}

var methods = []string{
	http.MethodGet,
	http.MethodPut,
	http.MethodHead,
	http.MethodPost,
	http.MethodTrace,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
}

var tSingleColumnLeftLookupCases = []testData{
	testData{
		expect: "*lut.SingleColumnLeftLookup",
		subjects: []string{
			// Lookup should be col 1, all uniq vals
			"abc",
			"def",
			"ghi",
		},
	},
	testData{
		expect: "*lut.SingleColumnLeftLookup",
		subjects: []string{
			// Lookup should be 2nd col
			"abc",
			"aef",
			"ghi",
		},
	},
	testData{
		expect: "*lut.SingleColumnLeftLookup",
		subjects: []string{
			// Lookup should be third col
			"abc",
			"abf",
			"ghi",
		},
	},
}

var tSingleColumnRightLookupCases = []testData{
	testData{
		expect: "*lut.SingleColumnRightLookup",
		subjects: []string{
			// Lookup should be col 0, all uniq vals
			"a",
			"ab",
			"aac",
		},
	},
	testData{
		expect: "*lut.SingleColumnRightLookup",
		subjects: []string{
			// Lookup should be 2nd col
			"aa",
			"aab",
			"aaac",
		},
	},
	testData{
		expect: "*lut.SingleColumnRightLookup",
		subjects: []string{
			// Lookup should be 3rd col
			"aaa",
			"aab",
			"aabc",
		},
	},
}

var tMapLookupCases = []testData{
	testData{
		expect: "*lut.MapLookup",
		subjects: []string{
			// no columns are unique
			"abc", "abz", "azc", "jkl", "mno",
		},
	},
}

func TestAnalyze(t *testing.T) {
	var tcases []testData
	tcases = append(tcases, tSingleColumnLeftLookupCases...)
	tcases = append(tcases, tSingleColumnRightLookupCases...)
	tcases = append(tcases, tMapLookupCases...)

	for _, tcase := range tcases {

		t.Run(fmt.Sprintf(`%s`, tcase.expect), func(t *testing.T) {
			l := Analyze(tcase.subjects)

			if got := fmt.Sprintf(`%T`, l); got != tcase.expect {
				t.Logf(`given subjects: %s`, tcase.subjects)
				t.Fatalf(`got %s, wanted %s`, got, tcase.expect)
			}
			for i, s := range tcase.subjects {
				i += 10
				l.Associate(s, i)

				if got := l.Lookup(s); i != got {
					t.Logf(`when looking up %s`, s)
					t.Logf(`given subjects: %s`, tcase.subjects)
					t.Fatalf(`got %d, wanted %d`, got, i)
				}
				if got := l.Lookup("miss"); -1 != got {
					t.Logf(`when looking up miss`)
					t.Logf(`given subjects: %s`, tcase.subjects)
					t.Fatalf(`got %d, wanted %d`, got, -1)
				}
			}
		})
	}
}

func BenchmarkSingleColumnLeftLookup(b *testing.B) {
	l := &SingleColumnLeftLookup{}
	tcase := tSingleColumnLeftLookupCases[0]
	lookupVal := 100
	lookupKey := tcase.subjects[0]

	if !l.Satisfied(tcase.subjects) {
		b.Logf(`given subjects: %s`, tcase.subjects)
		b.Fatalf(`was not satisfied`)
	}
	if err := l.Associate(lookupKey, lookupVal); err != nil {
		b.Logf(`when associating %s to %d`, lookupKey, lookupVal)
		b.Logf(`given subjects: %s`, tcase.subjects)
		b.Fatalf(`got err %s`, err)
	}
	b.ResetTimer()

	b.Run("Hit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if got := l.Lookup(lookupKey); lookupVal != got {
				b.Logf(`when looking up %s`, lookupKey)
				b.Logf(`given subjects: %s`, tcase.subjects)
				b.Fatalf(`got %d, wanted %d`, got, lookupVal)
			}
		}
	})
	b.Run("Miss", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if got := l.Lookup(`badkey`); -1 != got {
				b.Logf(`when looking up %s`, lookupKey)
				b.Logf(`given subjects: %s`, tcase.subjects)
				b.Fatalf(`got %d, wanted %d`, got, -1)
			}
		}
	})
}

func BenchmarkSingleColumnLeftUnknownLookup(b *testing.B) {
	l := &SingleColumnLeftUnknownLookup{}
	tcase := tSingleColumnLeftLookupCases[0]
	lookupVal := 100
	lookupKey := tcase.subjects[0]

	if !l.Satisfied(tcase.subjects) {
		b.Logf(`given subjects: %s`, tcase.subjects)
		b.Fatalf(`was not satisfied`)
	}
	if err := l.Associate(lookupKey, lookupVal); err != nil {
		b.Logf(`when associating %s to %d`, lookupKey, lookupVal)
		b.Logf(`given subjects: %s`, tcase.subjects)
		b.Fatalf(`got err %s`, err)
	}
	b.ResetTimer()

	b.Run("Hit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if got := l.Lookup(lookupKey); lookupVal != got {
				b.Logf(`when looking up %s`, lookupKey)
				b.Logf(`given subjects: %s`, tcase.subjects)
				b.Fatalf(`got %d, wanted %d`, got, lookupVal)
			}
		}
	})
	b.Run("HitKnown", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if got := l.LookupKnown(lookupKey); lookupVal != got {
				b.Logf(`when looking up %s`, lookupKey)
				b.Logf(`given subjects: %s`, tcase.subjects)
				b.Fatalf(`got %d, wanted %d`, got, lookupVal)
			}
		}
	})
	b.Run("Miss", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if got := l.Lookup(`badkey`); -1 != got {
				b.Logf(`when looking up %s`, lookupKey)
				b.Logf(`given subjects: %s`, tcase.subjects)
				b.Fatalf(`got %d, wanted %d`, got, -1)
			}
		}
	})
}

//
// func TestSingleColumnRightLookup(t *testing.T) {
// 	l := &SingleColumnLeftLookup{}
// 	for _, tcase := range tSingleColumnLookupLeftCases {
// 		if !l.Satisfied(tcase.subjects) {
// 			t.Logf(`given subjects: %s`, tcase.subjects)
// 			t.Fatalf(`was not satisfied`)
// 		}
// 		lookupVal := 100
// 		lookupKey := tcase.subjects[0]
//
// 		if err := l.Associate(lookupKey, lookupVal); err != nil {
// 			t.Logf(`when associating %s to %d`, lookupKey, lookupVal)
// 			t.Logf(`given subjects: %s`, tcase.subjects)
// 			t.Fatalf(`got err %s`, err)
// 		}
// 		if got := l.Lookup(lookupKey); lookupVal != got {
// 			t.Logf(`when looking up %s`, lookupKey)
// 			t.Logf(`given subjects: %s`, tcase.subjects)
// 			t.Fatalf(`got %d, wanted %d`, got, lookupVal)
// 		}
// 	}
// }

func BenchmarkSingleColumnRightLookup(b *testing.B) {
	l := &SingleColumnRightLookup{}
	tcase := tSingleColumnRightLookupCases[0]
	lookupVal := 100
	lookupKey := tcase.subjects[0]

	if !l.Satisfied(tcase.subjects) {
		b.Logf(`given subjects: %s`, tcase.subjects)
		b.Fatalf(`was not satisfied`)
	}
	if err := l.Associate(lookupKey, lookupVal); err != nil {
		b.Logf(`when associating %s to %d`, lookupKey, lookupVal)
		b.Logf(`given subjects: %s`, tcase.subjects)
		b.Fatalf(`got err %s`, err)
	}
	b.ResetTimer()

	b.Run("Hit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if got := l.Lookup(lookupKey); lookupVal != got {
				b.Logf(`when looking up %s`, lookupKey)
				b.Logf(`given subjects: %s`, tcase.subjects)
				b.Fatalf(`got %d, wanted %d`, got, lookupVal)
			}
		}
	})
	b.Run("Miss", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if got := l.Lookup(`badkey`); -1 != got {
				b.Logf(`when looking up %s`, lookupKey)
				b.Logf(`given subjects: %s`, tcase.subjects)
				b.Fatalf(`got %d, wanted %d`, got, -1)
			}
		}
	})
}

func BenchmarkMapLookup(b *testing.B) {
	l := &MapLookup{}
	tcase := tMapLookupCases[0]
	lookupVal := 100
	lookupKey := tcase.subjects[0]

	if !l.Satisfied(tcase.subjects) {
		b.Logf(`given subjects: %s`, tcase.subjects)
		b.Fatalf(`was not satisfied`)
	}
	if err := l.Associate(lookupKey, lookupVal); err != nil {
		b.Logf(`when associating %s to %d`, lookupKey, lookupVal)
		b.Logf(`given subjects: %s`, tcase.subjects)
		b.Fatalf(`got err %s`, err)
	}
	b.ResetTimer()

	b.Run("Hit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if got := l.Lookup(lookupKey); lookupVal != got {
				b.Logf(`when looking up %s`, lookupKey)
				b.Logf(`given subjects: %s`, tcase.subjects)
				b.Fatalf(`got %d, wanted %d`, got, lookupVal)
			}
		}
	})
	b.Run("Miss", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if got := l.Lookup(`badkey`); -1 != got {
				b.Logf(`when looking up %s`, lookupKey)
				b.Logf(`given subjects: %s`, tcase.subjects)
				b.Fatalf(`got %d, wanted %d`, got, -1)
			}
		}
	})
}

/*
func BenchmarkLookups(b *testing.B) {
	for _, tcase := range tcases {

		b.Run(fmt.Sprintf(`%s`, tcase.expect), func(b *testing.B) {
			analyzer := &LookupAnalyzer{}
			l, _ := analyzer.Analyze(tcase.subjects[:])

			if got := fmt.Sprintf(`%T`, l); got != tcase.expect {
				b.Fatalf(`given subjects: %s`, tcase.subjects)
				b.Fatalf(`got %s, wanted %s`, got, tcase.expect)
			}
			for i, s := range tcase.subjects {
				i += 10
				l.Associate(s, i)

				if got := l.Lookup(s); i != got {
					b.Fatalf(`when looking up %s`, s)
					b.Fatalf(`given subjects: %s`, tcase.subjects)
					b.Fatalf(`got %d, wanted %d`, got, i)
				}
			}

			subject := tcase.subjects[0]
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if got := l.Lookup(subject); 10 != got {
					b.Fatalf(`when looking up %s`, subject)
					b.Fatalf(`given subjects: %s`, tcase.subjects)
					b.Fatalf(`got %d, wanted %d`, got, 10)
				}
			}
		})
	}
}
*/

/*
type testAnalyzerCase struct {
	tbCase
	t     *testing.T
	cases []testData
}

type testAnalyzerCases []testAnalyzerCase

func (tc *testAnalyzerCase) test(t *testing.T) *testAnalyzerCase {
	tc.bindT(t)

	t.Run(fmt.Sprintf(`%s`, tc.expect), func(t *testing.T) {
		analyzer := &LookupAnalyzer{}
		l, _ := analyzer.Analyze(tc.subjects)

		if got := fmt.Sprintf(`%T`, l); got != tc.expect {
			t.Logf(`given subjects: %s`, tc.subjects)
			t.Logf(`got %s, wanted %s`, got, tc.expect)
		}

		for i, s := range tc.subjects {
			i += 10
			l.Associate(s, i)

			if got := l.Lookup(s); i != got {
				t.Logf(`when looking up %s`, s)
				t.Logf(`given subjects: %s`, tc.subjects)
				t.Logf(`got %d, wanted %d`, got, i)
			}
		}
	})
	return tc
}

func (tc *testAnalyzerCase) bench(b *testing.B) *testAnalyzerCase {
	tc.bindB(b)

	b.Run(fmt.Sprintf(`%s`, tc.expect), func(t *testing.T) {
		analyzer := &LookupAnalyzer{}
		l, _ := analyzer.Analyze(tc.subjects)

		if got := fmt.Sprintf(`%T`, l); got != tc.expect {
			t.Logf(`given subjects: %s`, tc.subjects)
			t.Logf(`got %s, wanted %s`, got, tc.expect)
		}

		for i, s := range tc.subjects {
			i += 10
			l.Associate(s, i)

			if got := l.Lookup(s); i != got {
				t.Logf(`when looking up %s`, s)
				t.Logf(`given subjects: %s`, tc.subjects)
				t.Logf(`got %d, wanted %d`, got, i)
			}
		}
	})
	return tc
}

var cases = []testData{
	testData{
		expect: "*lut.Lookup1",
		subjects: []string{
			// Lookup should be col 1, all uniq vals
			"abc", "def", "ghi", "jkl", "mno",
		},
	},
	testData{
		expect: "*lut.Lookup1",
		subjects: []string{
			// Lookup should be 2nd col
			"abc", "zbc", "def", "ghi", "jkl", "mno",
		},
	},
	testData{
		expect: "*lut.Lookup1",
		subjects: []string{
			// Lookup should be third col
			"abc", "abz", "def", "ghi", "jkl", "mno",
		},
	},
	testData{
		expect: "*lut.LookupMap",
		subjects: []string{
			// no columns are unique
			"abc", "abz", "azc", "jkl", "mno",
		},
	},
}

func TestAnalyzers(t *testing.T) {
	for _, tac := range testAnalyzerCases {
		tac.run(t)
	}
}

func TestNew(t *testing.T) {

		uniqVals := []string{"abc", "def", "ghi", "jkl", "mno"}

		lut := New(uniqVals)
		log.Printf("lookable: %#v\n", lut)
		log.Printf("lookable: %#v\n", lut.Associate(`def`, 100))
		log.Printf("lookable: %#v\n", lut.Lookup(`def`))
		log.Printf("lookable: %#v\n", lut.Associate(`def`, 100))

		//log.Printf("lookable: %#v\n", AnalyzeStrings(methods))

	t.Run(`Given`, func(t *testing.T) {
		t.Run(`SameSize`, func(t *testing.T) {
			t.Run(`Uniq Col expect Lookup1`, func(t *testing.T) {
				uniqVals := []string{"abc", "def", "ghi", "jkl", "mno"}

				lut := New(uniqVals)
				log.Printf("lookable: %#v\n", lut.Lookup("abc"))
				lut.Associate("abc", 123)
				log.Printf("lookable: %#v\n", lut.Lookup("abc"))
			})
		})
	})
}

func BenchmarkLookupTable1(b *testing.B) {
	uniqVals := []string{"abc", "def", "ghi", "jkl", "mno"}
	lut := New(uniqVals)
	lut.Associate("abc", 123)
	log.Printf("lookable: %#v\n", lut)
	b.ResetTimer()

	b.Run("Hit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if 123 != lut.Lookup(`abc`) {
				b.Fatalf("fail")
			}
		}
	})
	b.Run("Miss", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := lut.Lookup(`ZZZ`)
			if -1 != v {
				b.Fatalf("fail %d", v)
			}
		}
	})
}

func BenchmarkLookupMap(b *testing.B) {
	uniqVals := []string{"abc", "abcdd", "GGGG", "GGZ", "SDA"}
	lut := New(uniqVals)
	lut.Associate("abc", 123)
	b.ResetTimer()

	b.Run("Hit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if 123 != lut.Lookup(`abc`) {
				b.Fatalf("fail")
			}
		}
	})
	b.Run("Miss", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if -1 != lut.Lookup(`ZZZ`) {
				b.Fatalf("fail")
			}
		}
	})
}

/*

func BenchmarkLookupMap(b *testing.B) {
	rand.Seed(1)
	uniqVals := []string{
		http.MethodGet,
		http.MethodPut,
		http.MethodHead,
		http.MethodPost,
		http.MethodTrace,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
	}
	subjects := uniqVals
	subjectMember := subjects[0]
	subjectExists := "GET"
	lut := NewLookupMap()
	lut.Analyze(subjects)

	b.ResetTimer()

	b.Run(`lut.Lookup("def")`, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup("GET")
		}
	})
	b.Run("lut.Lookup(`def`)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup(`GET`)
		}
	})
	b.Run("lut.Lookup(http.MethodPost)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup(http.MethodPost)
		}
	})
	b.Run("lut.Lookup(subjectExists)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup(subjectExists)
		}
	})
	b.Run("lut.Lookup(subjectMember)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup(subjectMember)
		}
	})
}

func BenchmarkLookup2(b *testing.B) {
	rand.Seed(1)
	uniqVals := []string{
		http.MethodGet,
		http.MethodPut,
		http.MethodHead,
		http.MethodPost,
		http.MethodTrace,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
	}
	subjects := uniqVals
	subjectMember := subjects[0]
	subjectExists := "GET"
	lut := &Lookup2{}
	lut.Analyze(subjects)

	b.ResetTimer()

	b.Run(`lut.Lookup("def")`, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup("GET")
		}
	})
	b.Run("lut.Lookup(`def`)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup(`GET`)
		}
	})
	b.Run("lut.Lookup(http.MethodPost)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup(http.MethodPost)
		}
	})
	b.Run("lut.Lookup(subjectExists)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup(subjectExists)
		}
	})
	b.Run("lut.Lookup(subjectMember)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup(subjectMember)
		}
	})
}

func BenchmarkInterfaceMap(b *testing.B) {
	rand.Seed(1)
	uniqVals := []string{"abc", "def", "ghi", "jkl", "mno"}
	subjects := uniqVals
	subject := subjects[0]
	table := make(map[string]int)

	for _, subject := range subjects {
		table[subject] = 0
	}

	b.ResetTimer()

	b.Run("access1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = table[subject]
		}
	})
}

func BenchmarkActualMap(b *testing.B) {
	rand.Seed(1)
	subjects := []string{"abc", "def", "ghi", "jkl", "mno"}
	subject := subjects[0]
	table := make(map[string]int)

	for _, subject := range subjects {
		table[subject] = 0
	}

	b.ResetTimer()

	b.Run("access1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = table[subject]
		}
	})
}

func BenchmarkAsInterface(b *testing.B) {
	rand.Seed(1)
	uniqVals := []string{"abc", "def", "ghi", "jkl", "mno"}
	subjects := uniqVals
	subject := subjects[0]
	lut := New(subjects)

	b.ResetTimer()

	b.Run("lut.Lookup(StringSubject(`def`))", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup(`def`)
		}
	})
	b.Run("lut.Lookup(subject)", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lut.Lookup(subject)
		}
	})
}

/*
func New() *Lookup2 {
	var foo := [bar]
	h := &Lookup{
		table: [8][90]int{},
	}
	for _, m := range methods {
		h.table[len(m)][m[0]] = 1
	}
	return h
}

func (l *Lookup) SetFunc(method string, index int) {
	l.table[len(method)][method[0]] = index
}

func (l *Lookup) RunFunc(method string) int {
	return l.table[len(method)][method[0]]
}
*/

/*

var methods = []string{
	http.MethodGet,
	http.MethodPut,
	http.MethodHead,
	http.MethodPost,
	http.MethodTrace,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
}

var table = [8][90]LookupFunc{}

func init() {
	for _, m := range methods {
		table[len(m)][m[0]] = func() {}
	}
}

type LookupFunc func()

type MethodLookup struct {
	table [8][90]LookupFunc
}

func NewFunc() *MethodLookup {
	h := &MethodLookup{
		table: [8][90]LookupFunc{},
	}
	for _, m := range methods {
		h.table[len(m)][m[0]] = func() {}
	}
	return h
}

func (l *MethodLookup) SetFunc(method string, f LookupFunc) {
	l.table[len(method)][method[0]] = f
}

func (l *MethodLookup) RunFunc(method string) {
	l.table[len(method)][method[0]]()
}

type Lookup struct {
	table [8][90]int
}

type Heuristic interface {
}

type Analyzer interface {
	Analyze([]string)
}

type Lookable interface {
	Lookup(string) int
}

func New() *Lookup {
	h := &Lookup{
		table: [8][90]int{},
	}
	for _, m := range methods {
		h.table[len(m)][m[0]] = 1
	}
	return h
}

func (l *Lookup) SetFunc(method string, index int) {
	l.table[len(method)][method[0]] = index
}

func (l *Lookup) RunFunc(method string) int {
	return l.table[len(method)][method[0]]
}

func TestReal(t *testing.T) {
	for _, m := range methods {
		log.Printf("%s: %d\n", m, len(m)+int(m[1]))
	}
}

/*
type HttpLookupOne struct {
	table [5 * 25]HttpLookupFunc
}

func NewOne() *HttpLookupOne {
	h := &HttpLookupOne{
		table: [5 * 25]HttpLookupFunc{},
	}
	for _, m := range methods {
		offset := (len(m) - 3) * 25
		pos := int(m[0] - 65)
		h.table[offset+pos] = func() {}
	}
	return h
}
func Index() int {
	return 0
}
func (l *HttpLookupOne) SetFunc(method string, f HttpLookupFunc) {
	//l.table[len(method)-3][method[0]-65] = f
	//pos := int(method[0] - 65)
	l.table[((len(method)-3)*25)+int(method[0]-65)] = f
}

func (l *HttpLookupOne) RunFunc(method string) {
	//pos := int(method[0] - 65)
	l.table[((len(method)-3)*25)+int(method[0]-65)]()
}

type HttpLookupTable struct {
	table [5][25]int
}

func NewTable() *HttpLookupTable {
	h := &HttpLookupTable{
		table: [5][25]int{},
	}
	for _, m := range methods {
		h.table[len(m)-3][m[0]-65] = 0
	}
	return h
}

func (l *HttpLookupTable) Set(method string, i int) {
	l.table[len(method)-3][method[0]-65] = i
}

func (l *HttpLookupTable) Run(method string) int {
	return l.table[len(method)-3][method[0]-65]
}

func TestReal(t *testing.T) {
	for _, m := range methods {
		log.Printf("%s: %d\n", m, len(m)+int(m[1]))
	}
}

func BenchmarkCost(b *testing.B) {
	rand.Seed(1)
	method := methods[rand.Intn(len(methods))]
	method = "GET"
	l1 := [89]HttpLookupFunc{}
	l2 := [8][90]HttpLookupFunc{}
	for _, m := range methods {
		log.Printf("%s: %d\n", m, len(m)+int(m[1]))
		l1[len(m)+int(m[1])] = func() {}
		l2[len(m)][int(m[1])] = func() {}
	}

	b.ResetTimer()

	b.Run("runflat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l1[len(method)+int(method[1])]()
		}
	})

	b.Run("run dbl", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			l2[len(method)][int(method[1])]()
		}
	})
	b.Run("cost1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = len(method) + int(method[1])
		}
	})

	b.Run("cost2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = len(method) + int(method[1])
		}
	})

}

func BenchmarkAccessReal(b *testing.B) {
	rand.Seed(1)
	method := methods[rand.Intn(len(methods))]
	method = "GET"
	methodLen := len(method)
	log.Println("http.Method:", method, methodLen, 25)

	lookup := [5][25]int{
		[25]int{},
	}

	mapl := make(map[string]HttpLookupFunc)
	log.Println("lookup-pre", lookup)
	for _, m := range methods {
		lookup[len(m)-3][m[0]-65] = 1
		ml := len(m)
		log.Println("m:", m, ml, m[0]-65)
		mapl[m] = func() {}
	}

	log.Println("lookup-pos", lookup)
	log.Println("char", method[0])
	f := New()
	f.SetFunc(http.MethodGet, func() {
	})
	ft := NewTable()
	ft.Set(http.MethodGet, 10)
	fo := NewOne()
	fo.SetFunc(http.MethodGet, func() {
	})

	d1 := [5]HttpLookupFunc{func() {}}
	d2 := [5][5]HttpLookupFunc{[5]HttpLookupFunc{func() {}}}
	d3 := [5][5][5]HttpLookupFunc{[5][5]HttpLookupFunc{[5]HttpLookupFunc{func() {}}}}
	b.ResetTimer()

	b.Run("accessDepth1c", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = d1[0]
		}
	})

	b.Run("accessDepth1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d1[0]()
		}
	})

	b.Run("accessDepth2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d2[0][0]()
		}
	})

	b.Run("accessDepth3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d3[0][0][0]()
		}
	})

	b.Run("httplookup", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f.RunFunc(method)
		}
	})
	b.Run("httplookuptable", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if 1 == ft.Run(method) {

			}
		}
	})
	b.Run("httplookupone", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f.RunFunc(method)
		}
	})

	b.Run("Switch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			switch method {
			case http.MethodGet:
				_ = true
			case http.MethodPut:
				_ = true
			case http.MethodHead:
				_ = true
			case http.MethodPost:
				_ = true
			case http.MethodTrace:
				_ = true
			case http.MethodPatch:
				_ = true
			case http.MethodDelete:
				_ = true
			case http.MethodConnect:
				_ = true
			case http.MethodOptions:
				_ = true
			}
		}
	})
	b.Run("Lookup", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if 1 == lookup[len(method)-3][method[0]-65] {

			}
		}
	})
	b.Run("Map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mapl[method]()
		}
	})
}

func BenchmarkAccess(b *testing.B) {
	accessStatic := [6]int{0, 1, 2, 3, 4, 5}
	val := 0
	b.ResetTimer()
	b.Run("Equal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if val == 0 {

			}
		}
	})
	b.Run("EqualAccess", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if accessStatic[0] == 0 {

			}
		}
	})
}

func BenchmarkAccessSlice(b *testing.B) {
	accessStatic := []int{0, 1, 2, 3, 4, 5}
	b.ResetTimer()
	b.Run("Equal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if accessStatic[0] == 0 {

			}
		}
	})
}

func BenchmarkCompare(b *testing.B) {
	cmpStackEq1 := "test/string"
	cmpStackEq2 := "test/string"
	cmpStackNeq1 := "test/string"
	cmpStackNeq2 := "test/strinZ"
	b.ResetTimer()

	b.Run("Equal", func(b *testing.B) {

		b.Run("Const -> ", func(b *testing.B) {
			b.Run("Const", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if cmpConstEq1 == cmpConstEq2 {

					}
				}
			})
			b.Run("Stack", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if cmpConstEq1 == cmpStackEq2 {

					}
				}
			})
			b.Run("Heap", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if cmpConstEq1 == cmpHeapEq2 {

					}
				}
			})
		})

		b.Run("Stack -> ", func(b *testing.B) {
			b.Run("Stack", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if cmpStackEq1 == cmpStackEq2 {

					}
				}
			})
		})

		b.Run("Heap -> ", func(b *testing.B) {
			b.Run("Heap", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if cmpHeapEq1 == cmpHeapEq2 {

					}
				}
			})
		})

	})

	b.Run("Not Equal", func(b *testing.B) {

		b.Run("Const -> ", func(b *testing.B) {
			b.Run("Const", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if cmpConstNeq1 == cmpConstNeq2 {

					}
				}
			})
			b.Run("Stack", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if cmpConstNeq1 == cmpStackNeq2 {

					}
				}
			})
			b.Run("Heap", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if cmpConstNeq1 == cmpHeapNeq2 {

					}
				}
			})
		})

		b.Run("Stack -> ", func(b *testing.B) {
			b.Run("Stack", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if cmpStackNeq1 == cmpStackNeq2 {

					}
				}
			})
		})

		b.Run("Heap -> ", func(b *testing.B) {
			b.Run("Heap", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					if cmpHeapNeq1 == cmpHeapNeq2 {

					}
				}
			})
		})

	})
}
*/
