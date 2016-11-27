package genutil_test

import (
	"fmt"
	"testing"

	genutil "github.com/cstockton/go-exp/genutil"
)

func Example_testing() {
	MyTest := func(t *testing.T) {
		generator := genutil.WithName(`my_generated_source.go`)
		genutil.WithTesting(generator, t)
	}

	tests := []testing.InternalTest{{"TestWithTesting", MyTest}}
	matcher := func(pat, str string) (bool, error) { return true, nil }
	ok := testing.RunTests(matcher, tests)
	fmt.Println("Test Result:", ok)

	// Output:
	// Test Result: true
}
