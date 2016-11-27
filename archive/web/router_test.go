package web

import (
	"regexp"
	"testing"
)

func TestLiteralPattern(t *testing.T) {
	ctx, _ := NewContext()
	patterns := []string{
		"/",
		"/one",
		"/one/two",
	}

	for _, pattern := range patterns {
		p := &LiteralPattern{Literal: pattern}
		result := p.IsMatch(*ctx, pattern)

		if !result {
			t.Errorf("LiteralPattern.IsMatch(ctx, %s) = %v; want true", pattern, result)
		}
	}
}

func TestAnyPattern(t *testing.T) {
	ctx, _ := NewContext()
	patterns := []string{
		"/",
		"/one",
		"/one/two",
		"/one/two/three",
	}
	p := &AnyPattern{}

	for _, pattern := range patterns {
		result := p.IsMatch(*ctx, pattern)

		if !result {
			t.Errorf("AnyPattern.IsMatch(ctx, %s) = %v; want true", pattern, result)
		}
	}
}

func TestRegexpPattern(t *testing.T) {
	ctx, _ := NewContext()
	patterns := [][]string{
		{"(?P<one>[a-zA-Z]+)", "/theone"},
		{"(?P<one>[a-zA-Z]+)", "/thetwo"},
	}

	for _, pattern := range patterns {
		p := &RegexpPattern{Regexp: regexp.MustCompile(pattern[0])}
		result := p.IsMatch(*ctx, pattern[1])

		if !result {
			t.Errorf("RegexpPattern.IsMatch(ctx, %s) = %v; want true", pattern[1], result)
		}
	}
}
