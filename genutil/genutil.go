// Package genutil provides utilities to integrate a file generation phase as
// a testable member of your code base. You may use it to call `go generate` or
// execute any other arbitrary command. In addition you may generate files using
// standard Go templates.
package genutil

import (
	"io"
	"os/exec"
	"path/filepath"
	"testing"
	"text/template"
)

// @TODO
// If I come back to this I'll change the design of the interfaces because it
// ended up being too inflexible. The main issue being I can't easily chain
// N Loaders into N Committers and vice versa. This should be written more like
// go-flow, probably just use it and provide some genutil specific Flows.

// Loader loads data from various sources.
type Loader interface {
	Load() ([]byte, error)
}

// Committer can check if a loaders result is stale and commit the change.
type Committer interface {
	Stale(Loader) bool
	Commit(Loader) error
}

// Generator will attempt to Commit changes from Loaders.
type Generator interface {
	Generate() error
}

// With returns a new generator backed with the given Loader and Committer.
func With(l Loader, c Committer) Generator {
	return &generator{l: l, c: c}
}

// WithGroup returns a generator that will group together several generators
// and run them serially in the order added returning the first failure, if any.
func WithGroup(g ...Generator) Generator {
	return groupGenerator(g)
}

// WithName will return a Generator with a template loader and file Committer
// based on the given name and the conventions below.
//
// Tpl will be read from:
//   ./testdata/<name>.tpl
//
// Committer will write to:
//   ./<name>
func WithName(name string) Generator {
	return &generator{
		l: FromTplPath(filepath.Join(`testdata/`, name+`.tpl`), nil),
		c: ToFile(name),
	}
}

// WithReadme creates a README.md file from the given projects examples.
func WithReadme(t *testing.T) Generator {
	return WithName("README.md")
}

// WithTesting returns a generator that will generate within a sub-test of the
// given testing.T and report any errors using t.Error/Fatal.
func WithTesting(g Generator, t *testing.T) Generator {
	return &testingGenerator{g: g, t: t}
}

// ToCmd will commit the Loader by passing stdin to the given cmd. It will only
// run the command once. It's stale method will always return true if no error
// is set.
func ToCmd(cmd *exec.Cmd) Committer {
	return &cmdCommitLoader{cmd: cmd}
}

// ToCommand is shorthand for ToCmd(exec.Command("cmdName", "arg1", ...))
func ToCommand(name string, args ...string) Committer {
	return ToCmd(exec.Command(name, args...))
}

// ToFile returns a Committer that will write to a file. Multiple calls to
// commit will overwrite the existing file data.
func ToFile(filename string) Committer {
	return &fileCommitter{path: filename}
}

// ToReadWriter returns a Committer that will write to the given Writer. It will
// compare the bytes of Reader to the Loader for calls to Stale. It will read
// from Reader and write to Writer exactly once.
func ToReadWriter(rw io.ReadWriter) Committer {
	return &ioCommitter{r: rw, w: rw}
}

// ToWriter returns a Committer that will write to the given io.Writer. It will
// report the target as Stale until it commits or an error is set. It will write
// to Writer exactly once.
func ToWriter(w io.Writer) Committer {
	return &ioCommitter{w: w}
}

// FromCmd will run the given cmd and return stdout as the bytes. It will only
// run the command once.
func FromCmd(cmd *exec.Cmd) Loader {
	return &cmdCommitLoader{cmd: cmd}
}

// FromCommand is shorthand for FromCmd(exec.Command("cmdName", "arg1", ...))
func FromCommand(name string, args ...string) Loader {
	return FromCmd(exec.Command(name, args...))
}

// FromFile returns a Loader that will read from a file. A file will only be
// loaded once regardless if multiple calls to Load are made. The returned byte
// slice should not be modified.
func FromFile(filename string) Loader {
	return &fileLoader{path: filename}
}

// FromFunc returns a Loader from a given function. It will be called each time
// a Load() is requested.
func FromFunc(fn func() ([]byte, error)) Loader {
	return funcLoader(fn)
}

// FromNil returns a Loader that always produces empty bytes. Convenient when
// combined with a ToCommand that does not want stdin.
func FromNil() Loader {
	return &errGenerator{}
}

// FromReader returns a Loader that will read from the given io.Reader. It will
// read from Reader exactly once.
func FromReader(r io.Reader) Loader {
	return &readerLoader{r: r}
}

// FromTpl returns a Loader that will exec the given template, using data for
// the template data. The template is only executed once, multiple calls to load
// yield the same result.
func FromTpl(tpl *template.Template, data interface{}) Loader {
	return &tplLoader{tpl: tpl, data: data}
}

// FromTplPath returns a Loader that will read from a template, using data for
// the template data. The template is only executed once, multiple calls to load
// yield the same result.
func FromTplPath(tplPath string, data interface{}) Loader {
	return &tplLoader{path: tplPath, data: data}
}
