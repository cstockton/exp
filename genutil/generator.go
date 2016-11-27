package genutil

import (
	"sync"
	"testing"
)

var nilLoader = errGenerator{}

type errGenerator struct {
	err   error
	stale bool
}

func (g errGenerator) Load() ([]byte, error) {
	return nil, g.err
}

func (g errGenerator) Stale(l Loader) bool {
	return g.stale
}

func (g errGenerator) Commit(l Loader) error {
	return g.err
}

func (g errGenerator) Generate() error {
	return g.err
}

type generator struct {
	m sync.Mutex
	c Committer
	l Loader
}

func (g *generator) Generate() error {
	g.m.Lock()
	defer g.m.Unlock()
	if stale := g.c.Stale(g.l); !stale {
		return nil
	}
	return g.c.Commit(g.l)
}

type groupGenerator []Generator

func (g groupGenerator) Generate() error {
	for _, generator := range g {
		if err := generator.Generate(); err != nil {
			return err
		}
	}
	return nil
}

type testingGenerator struct {
	t *testing.T
	g Generator
}

func (g *testingGenerator) Generate() error {
	if err := g.g.Generate(); err != nil {
		g.t.Fatal(err)
	}
	return nil
}
