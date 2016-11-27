package pathy

import (
	"reflect"
	"testing"
)

func TestNewPathables(t *testing.T) {
	exp := Pathables{
		Path(`.`), Path(`./`), Path(`/`), Path(`/a`),
		Path(`/a/b`), Path(`/a/b/c`)}
	var tests = []struct {
		initial []string
		expect  Pathables
	}{
		{[]string{`.`, `./`, `/`, `/a`, `/a/b`, `/a/b/c`}, exp},
	}

	for _, test := range tests {
		if got := NewPathables(test.initial); !reflect.DeepEqual(got, test.expect) {
			t.Errorf("NewPathables(%q) = %q; expected %q", test.initial, got, test.expect)
			if breaker {
				break
			}
		}
	}
	if !NewPathables([]string{}).Empty() {
		t.Fatal("expected empty")
	}
}

func TestPathables(t *testing.T) {
	var tests = []struct {
		given  PathList
		expect Paths
	}{
		{
			PathList(FromColon(``)),
			Paths{``},
		},
		{
			PathList(FromColon(`:`)),
			Paths{``, ``},
		},
		{
			PathList(FromColon(`::`)),
			Paths{``, ``, ``},
		},
		{
			PathList(FromColon(`a`)),
			Paths{`a`},
		},
		{
			PathList(FromColon(`a:b:c`)),
			Paths{`a`, `b`, `c`},
		},
		{
			PathList(FromColon(`:a:b:c`)),
			Paths{``, `a`, `b`, `c`},
		},
	}

	for _, test := range tests {
		got := test.given.Paths()

		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("NewPathList(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathablesUnique(t *testing.T) {
	var tests = []struct {
		given  PathList
		expect PathList
	}{
		{
			PathList(FromColon(``)),
			PathList(FromColon(``)),
		},
		{
			PathList(FromColon(`:`)),
			PathList(FromColon(``)),
		},
		{
			PathList(FromColon(`::`)),
			PathList(FromColon(``)),
		},
		{
			PathList(FromColon(`a::a:b:c:a:d:e`)),
			PathList(FromColon(`a::b:c:d:e`)),
		},
		{
			PathList(FromColon(`a:b:c:a::a`)),
			PathList(FromColon(`a:b:c:`)),
		},
		{
			PathList(FromColon(`:a:b:c::b`)),
			PathList(FromColon(`:a:b:c`)),
		},
	}

	for _, test := range tests {
		got := test.given.Unique()

		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("NewPathList(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathablesContains(t *testing.T) {
	var tests = []struct {
		initial   Paths
		given     Paths
		expect    []bool
		expectIdx []int
	}{
		{
			NewPaths(`a`, `c`, `e`, `q`, `r`, `t`),
			NewPaths(`q`, `x`, `e`, `z`),
			[]bool{true, false, true, false},
			[]int{3, -1, 2, -1},
		},
	}

	for _, test := range tests {
		got := NewPaths(test.initial...)

		for expectIndex, search := range test.given {
			found := got.Contains(search)
			idx, f := got.Find(search)
			if f != found {
				t.Errorf("NewPaths(%q).Find(%q) = %v, %v;\nexpected %v, %v", test.initial, search, idx, found, test.expectIdx[expectIndex], test.expect[expectIndex])
			}
			if idx != test.expectIdx[expectIndex] {
				t.Errorf("NewPaths(%q).Find(%q) = %v, %v;\nexpected %v, %v", test.initial, search, idx, found, test.expectIdx[expectIndex], test.expect[expectIndex])
			}
			if found != test.expect[expectIndex] {
				t.Errorf("NewPaths(%q).Contains(%q) = %v;\nexpected %v", test.initial, search, found, test.expect[expectIndex])
				if breaker {
					break
				}
			}
		}
	}
}

func TestPathablesLen(t *testing.T) {
	var tests = []struct {
		given  []Paths
		expect []int
	}{
		{
			[]Paths{
				NewPaths(`a`, `c`, `e`, `q`, `r`, `t`),
				NewPaths(`a`, `b`),
				NewPaths(`a`),
				NewPaths(),
			},

			[]int{6, 2, 1, 0},
		},
	}

	for _, test := range tests {
		for index, paths := range test.given {
			got := NewPaths(paths...).Len()

			if got != test.expect[index] {
				t.Errorf("NewPaths(%q).Len() = %v; expected %v", test.given, got, test.expect[index])
				if breaker {
					break
				}
			}
		}
	}
}

func TestPathablesPaths(t *testing.T) {
	var tests = []struct {
		given  Paths
		expect string
	}{
		{Paths{``, `a`, `b`, `c`}, `a/b/c`},
		{Paths{`.`, `a`, `b`, `c`}, `a/b/c`},
		{Paths{`/`, `a`, `b`, `c`}, `/a/b/c`},
		{Paths{`z`, `a`, `b`, `c`}, `z/a/b/c`},
		{Paths{`./z`, `a`, `b`, `c`}, `z/a/b/c`},
		{Paths{`/z`, `a`, `b`, `c`}, `/z/a/b/c`},
		{Paths{`z/x`, `a`, `b`, `c`}, `z/x/a/b/c`},
		{Paths{`./z/x`, `a`, `b`, `c`}, `z/x/a/b/c`},
		{Paths{`/z/x`, `a`, `b`, `c`}, `/z/x/a/b/c`},
		{Paths{`/z/x`, `a`}, `/z/x/a`},
		{Paths{`/z/x`, `a`, `..`, `c`, `d`}, `/z/x/c/d`},
		{Paths{`/z/x`, `a`, `..`, `b`, `..`, `d`}, `/z/x/d`},
		{Paths{`/z/x`, `a`, `..`, `..`, `d`}, `/z/d`},
	}

	for _, test := range tests {
		if got := test.given.Join(); string(ToSlash(got)) != test.expect {
			t.Errorf("Paths(%q).Join() = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathablesPathList(t *testing.T) {
	var tests = []struct {
		given  Paths
		expect PathList
	}{
		{
			Paths{``, `a`, `b`, `c`},
			PathList(FromColon(`:a:b:c`)),
		},
		{
			Paths{},
			PathList(FromColon(``)),
		},
	}

	for _, test := range tests {
		if got := test.given.JoinList(); !reflect.DeepEqual(got, test.expect) {
			t.Errorf("Paths(%v).JoinList() = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}
