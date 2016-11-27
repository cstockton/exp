package pathy

import (
	"fmt"
	"reflect"
	"testing"
)

func ExamplePath() {
	fmt.Println(Path(`/a/b/c`))
	fmt.Println(Path(`/a/b/c`).Name())
	// Output:
	// /a/b/c
	// c
}

func TestNewPathList(t *testing.T) {
	var tests = []struct {
		given  Paths
		expect PathList
	}{
		{
			Paths{},
			PathList(FromColon(``)),
		},
		{
			Paths{``},
			PathList(FromColon(``)),
		},
		{
			Paths{``, ``},
			PathList(FromColon(`:`)),
		},
		{
			Paths{` `, ``},
			PathList(FromColon(` :`)),
		},
		{
			Paths{`a`},
			PathList(FromColon(`a`)),
		},
		{
			Paths{`a`, `b`, `c`},
			PathList(FromColon(`a:b:c`)),
		},
		{
			Paths{``, `a`, `b`, `c`},
			PathList(FromColon(`:a:b:c`)),
		},
	}

	for _, test := range tests {
		if got := NewPathList(test.given...); !reflect.DeepEqual(got, test.expect) {
			t.Errorf("NewPathList(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathListPaths(t *testing.T) {
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

func TestPathListUnique(t *testing.T) {
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

func TestNewPaths(t *testing.T) {
	var tests = []struct {
		initial Paths
		expect  Paths
	}{
		{
			Paths{`.`, `./`, `/`, `/a`, `/a/b`, `/a/b/c`},
			NewPaths(`.`, `./`, `/`, `/a`, `/a/b`, `/a/b/c`),
		},
	}

	for _, test := range tests {
		if got := NewPaths(test.initial...); !reflect.DeepEqual(got, test.expect) {
			t.Errorf("NewPaths(%q) = %q; expected %q", test.initial, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestNewPathsFromList(t *testing.T) {
	var tests = []struct {
		given  PathList
		expect Paths
	}{
		{
			NewPathList(`/a`, `/b`, `/c`, `/d`, `./`),
			NewPaths(`/a`, `/b`, `/c`, `/d`, `./`),
		},
	}

	for _, test := range tests {
		if got := NewPathsFromList(test.given); !reflect.DeepEqual(got, test.expect) {
			t.Errorf("NewPathsFromList(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathsSort(t *testing.T) {
	var tests = []struct {
		initial Paths
		expect  Paths
	}{
		{
			Paths{`t`, `q`, `c`, `r`, `e`, `a`},
			Paths{`a`, `c`, `e`, `q`, `r`, `t`},
		},
	}

	for _, test := range tests {
		got := NewPaths(test.initial...)
		got.Sort()

		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("NewPaths(%q).Sort() = %q; expected %q", test.initial, got, test.expect)
			if breaker {
				break
			}
		}
		if reflect.DeepEqual(got, test.initial) {
			t.Errorf("NewPaths(%q).Sort() = %q; expected %q", test.initial, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathsMap(t *testing.T) {
	var tests = []struct {
		initial Paths
		expect  Paths
	}{
		{
			Paths{`.`, `./`, `b`, `./././`, `/a`, `b`, `b/b`},
			Paths{`.`, `.`, `.`, `/a`, `b/b`},
		},
	}

	for _, test := range tests {
		tps := NewPaths(test.initial...)
		got := tps.Map(func(p *Path) bool {
			*p = p.Clean()

			if `b` == *p {
				return false
			}

			return true
		})

		if reflect.DeepEqual(got, tps) {
			t.Errorf("Paths(%q).Map(...) modified the original Paths", test.initial)
			if breaker {
				break
			}
		}
		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("Paths(%q).Filter(...) = %q; expected %q", test.initial, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathsMutate(t *testing.T) {
	var tests = []struct {
		initial Paths
		expect  Paths
	}{
		{
			Paths{`.`, `./`, `b`, `./././`, `/a`, `b`, `b/b`},
			Paths{`.`, `.`, `.`, `/a`, `b/b`},
		},
	}

	for _, test := range tests {
		got := NewPaths(test.initial...)
		got.Mutate(func(p *Path) bool {
			*p = p.Clean()

			if `b` == *p {
				return false
			}

			return true
		})

		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("Paths(%q).Mutate(...) = %q; expected %q", test.initial, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathsSearch(t *testing.T) {
	var tests = []struct {
		initial Paths
		given   Paths
		expect  []int
	}{
		{
			NewPaths(`a`, `c`, `e`, `q`, `r`, `t`),
			NewPaths(`q`, `c`, `e`, `z`),
			[]int{3, 1, 2, 6},
		},
	}

	for _, test := range tests {
		got := NewPaths(test.initial...)

		for expectIndex, search := range test.given {
			index := got.Search(search)

			if index != test.expect[expectIndex] {
				t.Errorf("NewPaths(%q).Search(%q) = %v; expected %v", test.initial, search, index, test.expect[expectIndex])
				if breaker {
					break
				}
			}
		}
	}
}

func TestPathsContains(t *testing.T) {
	var tests = []struct {
		initial Paths
		given   Paths
		expect  []bool
	}{
		{
			NewPaths(`a`, `c`, `e`, `q`, `r`, `t`),
			NewPaths(`q`, `x`, `e`, `z`),
			[]bool{true, false, true, false},
		},
	}

	for _, test := range tests {
		got := NewPaths(test.initial...)

		for expectIndex, search := range test.given {
			found := got.Contains(search)

			if found != test.expect[expectIndex] {
				t.Errorf("NewPaths(%q).Contains(%q) = %v; expected %v", test.initial, search, found, test.expect[expectIndex])
				if breaker {
					break
				}
			}
		}
	}
}

func TestPathsLen(t *testing.T) {
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

func TestPathsJoin(t *testing.T) {
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

func TestPathsJoinList(t *testing.T) {
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

func TestPathString(t *testing.T) {
	var tests = []struct {
		given  string
		expect string
	}{
		{``, ``},
		{`a`, `a`},
		{`a/b`, `a/b`},
		{`a/b/c`, `a/b/c`},
	}

	for _, test := range tests {
		if got := Path(test.given).String(); got != test.expect {
			t.Errorf("String(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathClean(t *testing.T) {
	var tests = []struct {
		given  string
		expect string
	}{
		// Already clean
		{"abc", "abc"},
		{"abc/def", "abc/def"},
		{"a/b/c", "a/b/c"},
		{".", "."},
		{"..", ".."},
		{"../..", "../.."},
		{"../../abc", "../../abc"},
		{"/abc", "/abc"},
		{"/", "/"},

		// Empty is current dir
		{"", "."},

		// Remove trailing slash
		{"abc/", "abc"},
		{"abc/def/", "abc/def"},
		{"a/b/c/", "a/b/c"},
		{"./", "."},
		{"../", ".."},
		{"../../", "../.."},
		{"/abc/", "/abc"},

		// Remove doubled slash
		{"abc//def//ghi", "abc/def/ghi"},
		{"//abc", "/abc"},
		{"///abc", "/abc"},
		{"////abc", "/abc"},
		{"//abc//", "/abc"},
		{"abc//", "abc"},
		{"abc///", "abc"},
		{"abc////", "abc"},

		// Remove . elements
		{"abc/./def", "abc/def"},
		{"/./abc/def", "/abc/def"},
		{"abc/.", "abc"},

		// Remove .. elements
		{"abc/def/ghi/../jkl", "abc/def/jkl"},
		{"abc/def/../ghi/../jkl", "abc/jkl"},
		{"abc/def/..", "abc"},
		{"abc/def/../..", "."},
		{"/abc/def/../..", "/"},
		{"abc/def/../../..", ".."},
		{"/abc/def/../../..", "/"},
		{"abc/def/../../../ghi/jkl/../../../mno", "../../mno"},
		{"/../abc", "/abc"},

		// Combinations
		{"abc/./../def", "def"},
		{"abc//./../def", "def"},
		{"abc/../../././../def", "../../def"},
	}
	for _, test := range tests {
		if got := Path(test.given).Clean(); string(ToSlash(got)) != test.expect {
			t.Errorf("Clean(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathElements(t *testing.T) {
	var tests = []struct {
		given  string
		expect Paths
	}{
		{``, Paths{`.`}}, {`.`, Paths{`.`}}, {`/`, Paths{`/`}},
		{`a`, Paths{`.`, `a`}}, {`./a`, Paths{`.`, `a`}}, {`/a`, Paths{`/`, `a`}},
		{`aa`, Paths{`.`, `aa`}}, {`./aa`, Paths{`.`, `aa`}}, {`/aa`, Paths{`/`, `aa`}},
		{`a/b`, Paths{`.`, `a`, `b`}}, {`./a/b`, Paths{`.`, `a`, `b`}}, {`/a/b`, Paths{`/`, `a`, `b`}},
		{`aa/bb`, Paths{`.`, `aa`, `bb`}}, {`./aa/bb`, Paths{`.`, `aa`, `bb`}}, {`/aa/bb`, Paths{`/`, `aa`, `bb`}},
		{`aa/bb/cc`, Paths{`.`, `aa`, `bb`, `cc`}}, {`./aa/bb`, Paths{`.`, `aa`, `bb`}}, {`/aa/bb`, Paths{`/`, `aa`, `bb`}},
	}

	for _, test := range tests {
		if got := Path(test.given).Elements(); !reflect.DeepEqual(got, test.expect) {
			t.Errorf("Elements(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathParent(t *testing.T) {
	var tests = []struct {
		given  string
		expect Path
	}{
		//*
		{``, Path(`.`)},
		{`.`, Path(`.`)},
		{`./`, Path(`.`)},
		{`.//`, Path(`.`)},
		{`.///`, Path(`.`)},
		{`a`, Path(`.`)},
		{`aa`, Path(`.`)},
		{`aaa`, Path(`.`)},
		{`a/b`, Path(`a`)},
		{`aa/b`, Path(`aa`)},
		{`aaa/b`, Path(`aaa`)},
		{`a/bb`, Path(`a`)},
		{`aa/bb`, Path(`aa`)},
		{`aaa/bb`, Path(`aaa`)},
		{`a/bbb`, Path(`a`)},
		{`aa/bbb`, Path(`aa`)},
		{`aaa/bbb`, Path(`aaa`)},
		{`a/b/ccc`, Path(`a/b`)},
		{`aa/bb/ccc`, Path(`aa/bb`)},
		{`aaa/bbb/ccc`, Path(`aaa/bbb`)},
		{`/`, Path(`/`)},
		{`//`, Path(`/`)},
		{`///`, Path(`/`)},
		{`/.`, Path(`/`)},
		{`/a`, Path(`/`)},
		{`/aa`, Path(`/`)},
		{`/aaa`, Path(`/`)},
		{`//a`, Path(`/`)},
		{`///aa`, Path(`/`)},
		{`////aaa`, Path(`/`)},
		{`/a/b`, Path(`/a`)},
		{`/aa/b`, Path(`/aa`)},
		{`/aaa/b`, Path(`/aaa`)},
		{`/a/bb`, Path(`/a`)},
		{`/aa/bb`, Path(`/aa`)},
		{`/aaa/bb`, Path(`/aaa`)},
		{`/a/bbb`, Path(`/a`)},
		{`/aa/bbb`, Path(`/aa`)},
		{`/aaa/bbb`, Path(`/aaa`)},
		{`/a/b/ccc`, Path(`/a/b`)},
		{`/aa/bb/ccc`, Path(`/aa/bb`)},
		{`/aaa/bbb/ccc`, Path(`/aaa/bbb`)},
	}
	for _, test := range tests {
		//test.expect = FromSlash(test.expect)

		if got := Path(test.given).Parent(); got != test.expect {
			t.Errorf("Parent(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathParents(t *testing.T) {
	var tests = []struct {
		given  string
		expect Paths
	}{
		{`./././`, Paths{`.`}},
		{`./././aa`, Paths{`.`}},
		{`./../../..`, Paths{`../..`, `..`, `.`}},
		{`..`, Paths{`.`}},
		{``, Paths{`.`}},
		{`.`, Paths{`.`}},
		{`a`, Paths{`.`}},
		{`a/b`, Paths{`a`, `.`}},
		{`a/b/c`, Paths{`a/b`, `a`, `.`}},
		{`aa/bb/cc/dd/ee/ff/gg`, Paths{
			`aa/bb/cc/dd/ee/ff`, `aa/bb/cc/dd/ee`,
			`aa/bb/cc/dd`, `aa/bb/cc`, `aa/bb`, `aa`, `.`,
		}},

		{`//a`, Paths{`/`}},
		{`///aa`, Paths{`/`}},
		{`////aaa`, Paths{`/`}},
		{`/`, Paths{`/`}},
		{`/`, Paths{`/`}},
		{`/a`, Paths{`/`}},
		{`/a/b`, Paths{`/a`, `/`}},
		{`/a/b/c`, Paths{`/a/b`, `/a`, `/`}},
		{`/aa/bb/cc/dd/ee/ff/gg`, Paths{
			`/aa/bb/cc/dd/ee/ff`, `/aa/bb/cc/dd/ee`,
			`/aa/bb/cc/dd`, `/aa/bb/cc`, `/aa/bb`, `/aa`, `/`,
		}},
	}
	for _, test := range tests {
		if got := Path(test.given).Parents(); !reflect.DeepEqual(got, test.expect) {
			t.Errorf("Parents(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathJoin(t *testing.T) {
	var tests = []struct {
		initial string
		given   Paths
		expect  string
	}{
		{``, Paths{`a`, `b`, `c`}, `a/b/c`},
		{`.`, Paths{`a`, `b`, `c`}, `a/b/c`},
		{`/`, Paths{`a`, `b`, `c`}, `/a/b/c`},
		{`z`, Paths{`a`, `b`, `c`}, `z/a/b/c`},
		{`./z`, Paths{`a`, `b`, `c`}, `z/a/b/c`},
		{`/z`, Paths{`a`, `b`, `c`}, `/z/a/b/c`},
		{`z/x`, Paths{`a`, `b`, `c`}, `z/x/a/b/c`},
		{`./z/x`, Paths{`a`, `b`, `c`}, `z/x/a/b/c`},
		{`/z/x`, Paths{`a`, `b`, `c`}, `/z/x/a/b/c`},
		{`/z/x`, Paths{`a`}, `/z/x/a`},
		{`/z/x`, Paths{`a`, `..`, `c`, `d`}, `/z/x/c/d`},
		{`/z/x`, Paths{`a`, `..`, `b`, `..`, `d`}, `/z/x/d`},
		{`/z/x`, Paths{`a`, `..`, `..`, `d`}, `/z/d`},
	}
	for _, test := range tests {
		if got := Path(test.initial).Join(test.given...); string(ToSlash(got)) != test.expect {
			t.Errorf("Path(%q).Join(%q) = %q; expected %q", test.initial, test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathSplit(t *testing.T) {
	var tests = []struct {
		given   string
		expect1 string
		expect2 string
	}{
		{``, `.`, ``},
		{`a`, `.`, `a`},
		{`aa`, `.`, `aa`},
		{`/a`, `/`, `a`},
		{`./a`, `.`, `a`},
		{`/a/b`, `/a`, `b`},
		{`a/b/c`, `a/b`, `c`},
		{`/a/b/c`, `/a/b`, `c`},
	}
	for _, test := range tests {
		if got1, got2 := Path(test.given).Split(); string(ToSlash(got1)) != test.expect1 || string(ToSlash(got2)) != test.expect2 {
			t.Errorf("Path(%q).Split() = (%q, %q); expected (%q, %q)",
				test.given, got1, got2, test.expect1, test.expect2)
			if breaker {
				break
			}
		}
	}
}

func TestPathName(t *testing.T) {
	var tests = []struct {
		given  string
		expect string
	}{
		{``, ``},
		{`.`, `.`},
		{`./`, ``},
		{`/a`, `a`}, {`./a`, `a`}, {`a`, `a`},
		{`/aa`, `aa`}, {`./aa`, `aa`}, {`aa`, `aa`},
		{`/b/a`, `a`}, {`./b/a`, `a`}, {`b/a`, `a`},
		{`/bb/aa`, `aa`}, {`./bb/aa`, `aa`}, {`bb/aa`, `aa`},
	}
	for _, test := range tests {
		if got := Path(test.given).Name(); string(got) != test.expect {
			t.Errorf("Name(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathStem(t *testing.T) {
	var tests = []struct {
		given  string
		expect string
	}{
		{`/aa/bb.tar.gz`, `bb`},
		{`/aa/b.tar.gz`, `b`},
		{`/a/bb.tar.gz`, `bb`},
		{`/a/b.tar.gz`, `b`},
		{`./aa/bb.tar.gz`, `bb`},
		{`./a/b.tar.gz`, `b`},
		{`bb.tar.gz`, `bb`},
		{`b.tar.gz`, `b`},
		{`bb.tar`, `bb`},
		{`b.tar`, `b`},
		{`bb`, `bb`},
		{`b`, `b`},
	}
	for _, test := range tests {
		if got := Path(test.given).Stem(); string(got) != test.expect {
			t.Errorf("Stem(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathSuffix(t *testing.T) {
	var tests = []struct {
		given  string
		expect string
	}{
		{`/a/b.tar.gz`, `gz`},
		{`b.tar.gz`, `gz`},
		{`b.tar`, `tar`},
		{`bbb`, ``},
		{`bb`, ``},
		{`./bb`, ``},
		{`./b`, ``},
		{`.b`, `b`},
		{`.`, ``},
	}

	for _, test := range tests {
		if got := Path(test.given).Suffix(); string(got) != test.expect {
			t.Errorf("Suffix(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathSuffixes(t *testing.T) {
	var tests = []struct {
		given  string
		expect Paths
	}{
		{`/bb.tar.gz`, Paths{`tar`, `gz`}},
		{`./bb.tar.gz`, Paths{`tar`, `gz`}},
		{`bb.tar.gz`, Paths{`tar`, `gz`}},

		{`/bb.tar`, Paths{`tar`}},
		{`./bb.tar`, Paths{`tar`}},
		{`bb.tar`, Paths{`tar`}},

		{`/bb`, nil},
		{`./bb`, nil},
		{`bb`, nil},
	}

	for _, test := range tests {
		if got := Path(test.given).Suffixes(); !reflect.DeepEqual(got, test.expect) {
			t.Errorf("Suffixes(%q) = %q; expected %q", test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathWithParent(t *testing.T) {
	var tests = []struct {
		initial string
		given   string
		expect  string
	}{
		{``, ``, ``},
		{``, ``, ``},
		{``, `z`, `z`},
		{``, `/z`, `/z`},
		{``, `/z/x`, `/z/x`},
		{``, `./z`, `z`},

		{`.`, ``, `.`},
		{`.`, `.`, `.`},
		{`.`, `z`, `z`},
		{`.`, `/z`, `/z`},
		{`.`, `/z/x`, `/z/x`},
		{`.`, `./z`, `z`},

		{`c`, ``, `c`},
		{`c`, `.`, `c`},
		{`c`, `z`, `z/c`},
		{`c`, `/z`, `/z/c`},
		{`c`, `/z/x`, `/z/x/c`},
		{`c`, `./z`, `z/c`},

		{`./c`, ``, `c`},
		{`./c`, `.`, `c`},
		{`./c`, `z`, `z/c`},
		{`./c`, `/z`, `/z/c`},
		{`./c`, `/z/x`, `/z/x/c`},
		{`./c`, `./z`, `z/c`},

		{`/c`, ``, `c`},
		{`/c`, `.`, `c`},
		{`/c`, `z`, `z/c`},
		{`/c`, `/z`, `/z/c`},
		{`/c`, `/z/x`, `/z/x/c`},
		{`/c`, `./z`, `z/c`},

		{`a/c`, ``, `c`},
		{`a/c`, `.`, `c`},
		{`a/c`, `z`, `z/c`},
		{`a/c`, `/z`, `/z/c`},
		{`a/c`, `/z/x`, `/z/x/c`},
		{`a/c`, `./z`, `z/c`},

		{`/a/c`, ``, `c`},
		{`/a/c`, `.`, `c`},
		{`/a/c`, `z`, `z/c`},
		{`/a/c`, `/z`, `/z/c`},
		{`/a/c`, `/z/x`, `/z/x/c`},
		{`/a/c`, `./z`, `z/c`},
	}

	for _, test := range tests {
		if got := Path(test.initial).WithParent(Path(test.given)); string(ToSlash(got)) != test.expect {
			t.Errorf("Path(%q).WithParent(%q) = %q; expected %q", test.initial, test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathWithName(t *testing.T) {
	name1 := `z.tar.gz`
	name2 := `zz.tar.gz`

	var tests = []struct {
		initial string
		given   string
		expect  string
	}{
		{`/aa/bb.tar.gz`, name2, `/aa/` + name2},
		{`/aa/b.tar.gz`, name1, `/aa/` + name1},
		{`/a/bb.tar.gz`, name2, `/a/` + name2},
		{`/a/b.tar.gz`, name1, `/a/` + name1},
		{`./aa/bb.tar.gz`, name2, `./aa/` + name2},
		{`./a/b.tar.gz`, name1, `./a/` + name1},
		{`bb.tar.gz`, name2, name2},
		{`b.tar.gz`, name1, name1},
		{`bb.tar`, name2, name2},
		{`b.tar`, name1, name1},
		{`bb`, name2, name2},
		{`b`, name1, name1},
		{``, name2, name2},
	}

	for _, test := range tests {
		if got := Path(test.initial).WithName(Path(test.given)); string(got) != test.expect {
			t.Errorf("Path(%q).WithName(%q) = %q; expected %q", test.initial, test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathWithStem(t *testing.T) {
	var tests = []struct {
		initial string
		given   string
		expect  string
	}{
		{`/aa/bb.tar.gz`, `zz`, `/aa/zz.tar.gz`},
		{`/aa/b.tar.gz`, `z`, `/aa/z.tar.gz`},
		{`/a/bb.tar.gz`, `zz`, `/a/zz.tar.gz`},
		{`/a/b.tar.gz`, `z`, `/a/z.tar.gz`},
		{`./aa/bb.tar.gz`, `zz`, `./aa/zz.tar.gz`},
		{`./a/b.tar.gz`, `z`, `./a/z.tar.gz`},
		{`bb.tar.gz`, `zz`, `zz.tar.gz`},
		{`b.tar.gz`, `z`, `z.tar.gz`},
		{`bb.tar`, `zz`, `zz.tar`},
		{`b.tar`, `z`, `z.tar`},
		{`bb`, `zz`, `zz`},
		{`b`, `z`, `z`},
		{``, `z`, `z`},
	}

	for _, test := range tests {
		if got := Path(test.initial).WithStem(Path(test.given)); string(got) != test.expect {
			t.Errorf("Path(%q).WithStem(%q) = %q; expected %q", test.initial, test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathWithSuffix(t *testing.T) {
	var tests = []struct {
		initial string
		given   string
		expect  string
	}{
		{`/aa/bb.tar.gz`, `gunzip`, `/aa/bb.tar.gunzip`},
		{`/aa/b.tar.gz`, `gunzip`, `/aa/b.tar.gunzip`},
		{`/a/bb.tar.gz`, `gunzip`, `/a/bb.tar.gunzip`},
		{`/a/b.tar.gz`, `gunzip`, `/a/b.tar.gunzip`},
		{`./aa/bb.tar.gz`, `gunzip`, `./aa/bb.tar.gunzip`},
		{`./a/b.tar.gz`, `gunzip`, `./a/b.tar.gunzip`},
		{`bb.tar.gz`, `gunzip`, `bb.tar.gunzip`},
		{`b.tar.gz`, `gunzip`, `b.tar.gunzip`},
		{`bb.tar`, `gunzip`, `bb.gunzip`},
		{`b.tar`, `gunzip`, `b.gunzip`},
		{`bb`, `gunzip`, `bb.gunzip`},
		{`b`, `gunzip`, `b.gunzip`},
		{``, `gunzip`, `.gunzip`},
	}

	for _, test := range tests {
		if got := Path(test.initial).WithSuffix(Path(test.given)); string(got) != test.expect {
			t.Errorf("Path(%q).WithSuffix(%q) = %q; expected %q", test.initial, test.given, got, test.expect)
			if breaker {
				break
			}
		}
	}
}
