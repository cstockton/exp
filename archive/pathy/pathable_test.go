package pathy

import (
	"fmt"
	"testing"
)

type UserPathableStruct struct {
	path string
}

func (p UserPathableStruct) String() string {
	return p.path
}

func (p UserPathableStruct) Path() Path {
	return Path(p.path)
}

type UserPathableSting string

func (p UserPathableSting) String() string {
	return string(p)
}

func (p UserPathableSting) Path() Path {
	return Path(p)
}

var pathableValidTests = []struct {
	initial interface{}
	expect  Path
}{
	{[]byte(`/a/b/c/byte`), `/a/b/c/byte`},
	{`/a/b/c/string`, `/a/b/c/string`},
	{Path(`/a/b/c/Path`), `/a/b/c/Path`},
	{Pathable(Path(`/a/b/c/Pathable/Path`)), `/a/b/c/Pathable/Path`},
	{*New(`/a/b/c/New`), `/a/b/c/New`},
	{UserPathableSting(`/a/b/c/UserPathableSting`), `/a/b/c/UserPathableSting`},
	{UserPathableStruct{`/a/b/c/UserPathableStruct`}, `/a/b/c/UserPathableStruct`},

	{1234567890, `1234567890`},
	{int(1234567890), `1234567890`},
	{int8(127), `127`},
	{int16(32767), `32767`},
	{int32(2147483647), `2147483647`},
	{int64(9223372036854775807), `9223372036854775807`},
	{uint(1234567890), `1234567890`},
	{uint8(255), `255`},
	{uint16(65535), `65535`},
	{uint32(4294967295), `4294967295`},
	{uint64(18446744073709551615), `18446744073709551615`},
	{123.45, `123.45`},
	{float32(123.45), `123.45`},
	{float64(123.45), `123.45`},
}

var pathableInvalidTests = []struct {
	initial interface{}
	expect  Path
}{
	{true, ``},
	{false, ``},
	{fmt.Println, ``},
}

func ExamplePathify() {
	fmt.Println(Pathify(`/a/b/c`))
	// Output:
	// /a/b/c
}

func TestPathify(t *testing.T) {
	for _, test := range pathableValidTests {
		if got := Pathify(test.initial); got.Path() != test.expect {
			t.Errorf("Given %[1]T to Pathify(%#[1]v) = %v; expected %v", test.initial, got, test.expect)
			if breaker {
				break
			}
		}
	}
	for _, test := range pathableInvalidTests {
		if got := Pathify(test.initial); got.Path() != test.expect {
			t.Errorf("Pathify(%v) = %v; expected %v", test.initial, got, test.expect)
			if breaker {
				break
			}
		}
	}
}

func TestPathifyChk(t *testing.T) {
	for _, test := range pathableValidTests {
		got, gotErr := PathifyChk(test.initial)

		if got.Path() != test.expect {
			t.Errorf("Pathify(%v) = %v; expected %v", test.initial, got, test.expect)
		}
		if gotErr != nil {
			t.Errorf("Pathify(%v) = %v; error was: %v", test.initial, got, gotErr)
		}
	}
	for _, test := range pathableInvalidTests {
		got, gotErr := PathifyChk(test.initial)

		if got.Path() != test.expect {
			t.Errorf("Pathify(%v) = %v; expected %v", test.initial, got, test.expect)
		}
		if gotErr == nil {
			t.Errorf("Pathify(%v) = %v; expected an error", test.initial, got)
		}
	}
}

func TestPathifier(t *testing.T) {
	var pathSlice []interface{}
	for _, test := range pathableValidTests {
		pathSlice = append(pathSlice, test.initial)
	}
	for _, test := range pathableInvalidTests {
		pathSlice = append(pathSlice, test.initial)
	}

	ch := Pathifier(pathSlice...)

	for _, test := range pathableValidTests {
		got := <-ch

		if got.Path() != test.expect {
			t.Errorf("Pathify(%v) = %v; expected %v", test.initial, got, test.expect)
		}
	}
	for pathable := range ch {
		t.Errorf("channel should have been drained, but got %v", pathable)
	}
}

func TestPathifierSlice(t *testing.T) {
	check := func(p ...interface{}) {
		ch := Pathifier(p...)

		for _, test := range pathableValidTests {
			got := <-ch

			if got == nil {
				t.Errorf("Pathifier(%v) = nil; expected %v", test.initial, test.expect)
			}
			if got.Path() != test.expect {
				t.Errorf("Pathify(%v) = %v; expected %v", test.initial, got, test.expect)
			}
		}
		for pathable := range ch {
			t.Errorf("channel should have been drained, but got %v", pathable)
		}
	}

	testStrings := []string{}
	testPaths := Paths{}
	testPathSlice := []Path{}
	testPathables := Pathables{}
	testPathableSlice := []Pathable{}

	for _, test := range pathableValidTests {
		v := Pathify(test.initial)

		testStrings = append(testStrings, string(v.Path()))
		testPaths = append(testPaths, v.Path())
		testPathSlice = append(testPathSlice, v.Path())
		testPathables = append(testPathables, Pathable(v))
		testPathableSlice = append(testPathableSlice, Pathable(v))
	}

	check(testStrings)
	check(testPaths)
	check(testPathSlice)
	check(testPathables)
	check(testPathableSlice)
}
