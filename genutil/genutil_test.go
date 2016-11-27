package genutil

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
)

const (
	testNotFound          = `_not_found_for_testing`
	testTextFile          = "testdata/test_from_file.txt"
	testTextFileData      = "test_from_file\n"
	testTplFile           = "testdata/test_from_tpl.tpl"
	testTplFileVar        = "tpldata"
	testTplFileData       = "test_from_tpl tpldata\n"
	testTplFileParseError = "testdata/test_from_tpl_parse_error.tpl"
)

var testNoEnt = filepath.Join(testNotFound, testNotFound)

func TestFrom(t *testing.T) {
	errorLoaders := []Loader{
		FromCmd(nil),
		FromCmd(exec.Command("example_command_not_found_error")),
		FromCommand("example_command_not_found_error"),
		FromFile("testdata/example_file_not_found_error"),
		FromReader(nil),
		FromTpl(nil, nil),
		FromTplPath("", nil),
		FromTplPath(testTplFileParseError, nil),
		FromTpl(template.New("empty"), nil),
	}

	t.Run("FromCmd", func(t *testing.T) {
		loader := FromCmd(exec.Command("go", "version"))

		b, err := loader.Load()
		if err != nil {
			t.Fatal(err)
		}
		if got := string(b); !strings.Contains(got, "go version") {
			t.Fatalf("wanted go version; got %q", got)
		}

		// add neg loaders
		errStdoutSet := exec.Command("go")
		errStdoutSet.Stdout = new(bytes.Buffer)
		errorLoaders = append(errorLoaders, FromCmd(errStdoutSet))

		errStdinSet := exec.Command("go")
		errStdinSet.Stdin = new(bytes.Buffer)
		errorLoaders = append(errorLoaders, FromCmd(errStdinSet))

		errExited := exec.Command("go")
		errExited.Run()
		errorLoaders = append(errorLoaders, FromCmd(errExited))
	})
	t.Run("FromCommand", func(t *testing.T) {
		loader := FromCommand("go", "version")

		b, err := loader.Load()
		if err != nil {
			t.Fatal(err)
		}
		if got := string(b); !strings.Contains(got, "go version") {
			t.Fatalf("wanted go version; got %q", got)
		}
	})
	t.Run("FromFile", func(t *testing.T) {
		loader := FromFile(testTextFile)

		b, err := loader.Load()
		if err != nil {
			t.Fatal(err)
		}
		if got := string(b); got != testTextFileData {
			t.Fatalf("wanted loaded data: %q; got %q", testTextFileData, got)
		}
	})
	t.Run("FromFunc", func(t *testing.T) {
		data := `from_func_test`
		loader := FromFunc(func() ([]byte, error) {
			return []byte(data), nil
		})

		b, err := loader.Load()
		if err != nil {
			t.Fatal(err)
		}
		if got := string(b); got != data {
			t.Fatalf("wanted loaded data: %q; got %q", data, got)
		}
	})
	t.Run("FromNil", func(t *testing.T) {
		loader := FromNil()

		b, err := loader.Load()
		if err != nil {
			t.Fatal(err)
		}
		if got := string(b); got != "" {
			t.Fatalf("wanted loaded data: %q; got %q", "", got)
		}

		// neg
		errorLoaders = append(errorLoaders, &errGenerator{errors.New("err"), true})
	})
	t.Run("FromReader", func(t *testing.T) {
		const readerStr = `FromReader`
		loader := FromReader(strings.NewReader(readerStr))

		b, err := loader.Load()
		if err != nil {
			t.Fatal(err)
		}
		if got := string(b); got != readerStr {
			t.Fatalf("wanted loaded data: %q; got %q", readerStr, got)
		}
	})
	t.Run("FromTplPath", func(t *testing.T) {
		loader := FromTplPath(testTplFile, "tpldata")

		b, err := loader.Load()
		if err != nil {
			t.Fatal(err)
		}
		if got := string(b); got != testTplFileData {
			t.Fatalf("wanted loaded data: %q; got %q", testTplFileData, got)
		}

		t.Run("FromTpl", func(t *testing.T) {
			loader = FromTpl(loader.(*tplLoader).tpl, "tpldata")

			b, err := loader.Load()
			if err != nil {
				t.Fatal(err)
			}
			if got := string(b); got != testTplFileData {
				t.Fatalf("wanted loaded data: %q; got %q", testTplFileData, got)
			}
		})
	})
	t.Run("FromErrors", func(t *testing.T) {
		for _, loader := range errorLoaders {
			b, err := loader.Load()
			if err == nil {
				t.Fatalf("expected error for loader `%#v` but got nil", loader)
			}
			if 0 != len(b) {
				t.Fatalf("expected empty bytes for loader `%#v` but got %v", loader, b)
			}
		}
	})
}

func TestTo(t *testing.T) {
	load := func() Loader {
		return FromReader(strings.NewReader("test_to_text"))
	}
	errorCommitters := []Committer{
		ToCmd(nil),
		ToCmd(exec.Command("example_command_not_found_error")),
		ToCommand("example_command_not_found_error"),
		ToFile("testdata/example_dir_not_found_error/example_file"),
		ToWriter(nil),
		ToReadWriter(nil),
	}

	t.Run("ToCmd", func(t *testing.T) {
		committer := ToCmd(exec.Command("go", "fmt"))
		if stale := committer.Stale(nil); stale {
			t.Fatalf("expected stale for committer `%#v`", committer)
		}
		if err := committer.Commit(nil); err == nil {
			t.Fatalf("expected error for committer `%#v` but got nil", committer)
		}

		err := committer.Commit(load())
		if err != nil {
			t.Fatal(err)
		}

		// add neg loaders
		errStdoutSet := exec.Command("go")
		errStdoutSet.Stdout = new(bytes.Buffer)
		errorCommitters = append(errorCommitters, ToCmd(errStdoutSet))

		errStdinSet := exec.Command("go")
		errStdinSet.Stdin = new(bytes.Buffer)
		errorCommitters = append(errorCommitters, ToCmd(errStdinSet))

		errExited := exec.Command("go")
		errExited.Run()
		errorCommitters = append(errorCommitters, ToCmd(errExited))

		t.Run("ToCmdNeg", func(t *testing.T) {
			committer := ToCmd(exec.Command("go", "fmt"))
			err = committer.Commit(errGenerator{errors.New("err"), true})
			if err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", committer)
			}
		})
	})
	t.Run("ToCommand", func(t *testing.T) {
		committer := ToCommand("go", "fmt")
		if stale := committer.Stale(nil); stale {
			t.Fatalf("expected stale for committer `%#v`", committer)
		}
		if stale := committer.Stale(load()); !stale {
			t.Fatalf("expected stale for committer `%#v`", committer)
		}
		if err := committer.Commit(nil); err == nil {
			t.Fatalf("expected error for committer `%#v` but got nil", committer)
		}

		err := committer.Commit(load())
		if err != nil {
			t.Fatal(err)
		}
		t.Run("ToCommandNeg", func(t *testing.T) {
			committer := ToCommand("go", "fmt")
			err := committer.Commit(errGenerator{errors.New("err"), true})
			if err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", committer)
			}
		})
	})
	t.Run("ToFile", func(t *testing.T) {
		committer := ToFile("testdata/test_to_file.txt")
		if stale := committer.Stale(nil); stale {
			t.Fatalf("expected stale for committer `%#v`", committer)
		}
		if err := committer.Commit(nil); err == nil {
			t.Fatalf("expected error for committer `%#v` but got nil", committer)
		}
		if err := committer.Commit(load()); err != nil {
			t.Fatal(err)
		}

		t.Run("ToFileNeg", func(t *testing.T) {
			committer := ToFile("testdata/test_to_file.txt")
			err := committer.Commit(errGenerator{errors.New("err"), true})
			if err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", committer)
			}
		})
	})
	t.Run("ToReadWriter", func(t *testing.T) {
		committer := ToReadWriter(new(bytes.Buffer))
		if stale := committer.Stale(nil); stale {
			t.Fatalf("expected stale for committer `%#v`", committer)
		}
		if err := committer.Commit(nil); err == nil {
			t.Fatalf("expected error for committer `%#v` but got nil", committer)
		}

		err := committer.Commit(load())
		if err != nil {
			t.Fatal(err)
		}
		t.Run("ToReadWriterNeg", func(t *testing.T) {
			committer := ToReadWriter(new(bytes.Buffer))
			err := committer.Commit(errGenerator{errors.New("err"), true})
			if err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", committer)
			}
		})
	})
	t.Run("ToWriter", func(t *testing.T) {
		committer := ToWriter(new(bytes.Buffer))
		if stale := committer.Stale(nil); stale {
			t.Fatalf("expected stale for committer `%#v`", committer)
		}
		if err := committer.Commit(nil); err == nil {
			t.Fatalf("expected error for committer `%#v` but got nil", committer)
		}

		err := committer.Commit(load())
		if err != nil {
			t.Fatal(err)
		}
		t.Run("ToWriterNeg", func(t *testing.T) {
			committer := ToWriter(new(bytes.Buffer))
			err := committer.Commit(errGenerator{errors.New("err"), true})
			if err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", committer)
			}
		})
	})
	t.Run("ToErrors", func(t *testing.T) {
		for _, committer := range errorCommitters {
			err := committer.Commit(nil)
			if err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", committer)
			}
			err = committer.Commit(errGenerator{errors.New("err"), true})
			if err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", committer)
			}
			err = committer.Commit(load())
			if err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", committer)
			}
		}
	})
}

func TestToEdges(t *testing.T) {
	type testTo struct {
		factory func() Committer
	}
	tests := []testTo{
		{func() Committer { return ToCmd(exec.Command("go", "fmt")) }},
		{func() Committer { return ToCommand("go", "fmt") }},
		{func() Committer { return ToFile("testdata/test_to_file.txt") }},
		{func() Committer { return ToReadWriter(new(bytes.Buffer)) }},
		{func() Committer { return ToWriter(new(bytes.Buffer)) }},
	}
	errTests := []testTo{
		{func() Committer { return ToCmd(nil) }},
		{func() Committer { return ToCmd(exec.Command(testNoEnt)) }},
		{func() Committer { return ToCommand(testNoEnt) }},
		{func() Committer { return ToFile(testNoEnt) }},
		{func() Committer { return ToWriter(nil) }},
		{func() Committer { return ToReadWriter(nil) }},
		{func() Committer {
			cmd := exec.Command("go")
			cmd.Stdout = new(bytes.Buffer)
			return ToCmd(cmd)
		}},
		{func() Committer {
			cmd := exec.Command("go")
			cmd.Stdin = new(bytes.Buffer)
			return ToCmd(cmd)
		}},
		{func() Committer {
			cmd := exec.Command("go")
			cmd.Run()
			return ToCmd(cmd)
		}},
	}
	load := func() Loader {
		return FromReader(strings.NewReader("test_to_text"))
	}
	for _, test := range tests {
		c := test.factory()
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			if stale := c.Stale(nil); stale {
				t.Fatalf("expected stale for committer `%#v`", c)
			}
			if err := c.Commit(nil); err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", c)
			}
			if err := c.Commit(load()); err != nil {
				t.Fatal(err)
			}
			if stale := c.Stale(load()); stale != false {
				t.Fatalf("expected stale for committer `%#v`", c)
			}
			if _, ok := c.(*cmdCommitLoader); ok {
				t.Skip("skipping for cmdCommitLoader")
			}
			if ioc, ok := c.(*ioCommitter); ok {
				if ioc.r == nil {
					t.Skip("skipping for ioCommitter with nil reader")
				}
			}
			if stale := test.factory().Stale(errGenerator{errors.New("err"), true}); stale {
				t.Fatalf("expected false for stale committer `%#v` but got true", c)
			}
		})
		t.Run(fmt.Sprintf("%TErrLoader", c), func(t *testing.T) {
			if err := test.factory().Commit(errGenerator{errors.New("err"), true}); err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", c)
			}
			if err := test.factory().Commit(nil); err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", c)
			}
		})
	}
	for _, test := range errTests {
		c := test.factory()
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			if err := c.Commit(nil); err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", c)
			}
			if err := c.Commit(errGenerator{errors.New("err"), true}); err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", c)
			}
			if err := c.Commit(load()); err == nil {
				t.Fatalf("expected error for committer `%#v` but got nil", c)
			}
			if _, ok := c.(*cmdCommitLoader); ok {
				t.Skip("skipping for cmdCommitLoader")
			}
			if ioc, ok := c.(*ioCommitter); ok {
				if ioc.r == nil {
					t.Skip("skipping for ioCommitter with nil reader")
				}
			}
			if stale := c.Stale(load()); stale {
				t.Fatalf("expected stale for committer `%#v`", c)
			}
		})
	}
}

func capture(t *testing.T, fn func()) string {
	stdout := os.Stdout
	defer func() {
		os.Stdout = stdout
	}()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	defer r.Close()
	os.Stdout = w

	fn()
	buf := make([]byte, 4096)
	n, err := r.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	if n == 0 {
		t.Fatal("no output captured")
	}
	return string(buf)
}

func TestWith(t *testing.T) {
	type testWith struct {
		g Generator
	}
	tests := []testWith{
		{With(FromReader(strings.NewReader(`hello`)),
			ToReadWriter(new(bytes.Buffer)))},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%T", test.g), func(t *testing.T) {
			if err := test.g.Generate(); err != nil {
				t.Fatal(err)
			}
			g := WithTesting(test.g, t)
			if err := g.Generate(); err != nil {
				t.Fatal(err)
			}
			g = WithGroup(test.g, g)
			if err := g.Generate(); err != nil {
				t.Fatal(err)
			}

			errGen := errGenerator{errors.New("err"), true}
			g = WithGroup(g, With(errGen, errGen))
			if err := g.Generate(); err == nil {
				t.Fatal("expected non nil err")
			}

			stdout := capture(t, func() {
				MyTest := func(t *testing.T) {
					g := WithTesting(errGenerator{errors.New("err"), true}, t)
					g.Generate()
				}
				tests := []testing.InternalTest{{Name: "TestWithTesting", F: MyTest}}
				matcher := func(pat, str string) (bool, error) { return true, nil }
				if testing.RunTests(matcher, tests) {
					t.Errorf("expected failure")
				}
			})
			if !strings.Contains(stdout, "FAIL") {
				t.Errorf("Expected failure WithTesting, stdout: %v", stdout)
			}
		})
	}
}
