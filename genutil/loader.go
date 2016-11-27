package genutil

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"sync"
	"text/template"
)

type funcLoader func() ([]byte, error)

func (l funcLoader) Load() ([]byte, error) {
	return l()
}

type fileLoader struct {
	once sync.Once
	path string
	buf  []byte
	err  error
}

func (l *fileLoader) Load() ([]byte, error) {
	l.once.Do(func() {
		l.buf, l.err = ioutil.ReadFile(l.path)
	})
	return l.buf, l.err
}

type readerLoader struct {
	once sync.Once
	r    io.Reader
	buf  []byte
	err  error
}

func (l *readerLoader) Load() ([]byte, error) {
	l.once.Do(func() {
		if l.r == nil {
			l.err = errors.New("genutil: io.Reader was nil")
			return
		}
		l.buf, l.err = ioutil.ReadAll(l.r)
	})
	return l.buf, l.err
}

type tplLoader struct {
	once sync.Once
	path string
	buf  []byte
	err  error
	tpl  *template.Template
	data interface{}
}

func (l *tplLoader) Load() ([]byte, error) {
	l.once.Do(func() {
		if l.tpl == nil {
			if 0 == len(l.path) {
				l.err = errors.New("genutil: template was nil and path was empty")
				return
			}
			if l.tpl, l.err = template.ParseFiles(l.path); l.err != nil {
				return
			}
		}

		var buf bytes.Buffer
		if err := l.tpl.Execute(&buf, l.data); err != nil {
			l.err = fmt.Errorf(`loader: error "%s" executing tpl "%s"`, err, l.tpl.Name())
		}
		l.buf = buf.Bytes()
	})
	return l.buf, l.err
}

type cmdCommitLoader struct {
	once sync.Once
	cmd  *exec.Cmd
	buf  []byte
	err  error
}

func (c *cmdCommitLoader) Stale(l Loader) bool {
	if l == nil {
		return false
	}
	stale := c.cmd == nil || c.cmd.ProcessState == nil || !c.cmd.ProcessState.Exited()
	return stale && c.err == nil
}

var (
	errCmdNil    = errors.New("genutil: command must not be nil")
	errCmdExited = errors.New("genutil: command has already exited")
	errCmdStdin  = errors.New("genutil: command stdin already set")
	errCmdStdout = errors.New("genutil: command stdout already set")
	errCmdEmpty  = errors.New("genutil: command was empty")
)

func (c *cmdCommitLoader) Load() ([]byte, error) {
	c.once.Do(func() {
		if c.cmd == nil {
			c.err = errCmdNil
			return
		}
		if c.cmd.ProcessState != nil && c.cmd.ProcessState.Exited() {
			c.err = errCmdExited
			return
		}
		if c.cmd.Stdout != nil {
			c.err = errCmdStdout
			return
		}

		var stdout bytes.Buffer
		c.cmd.Stdout = &stdout
		if c.err = c.cmd.Run(); c.err != nil {
			return
		}
		c.buf = stdout.Bytes()
		return
	})
	return c.buf, c.err
}

func (c *cmdCommitLoader) Commit(l Loader) error {
	if l == nil {
		return errNilLoader
	}
	c.once.Do(func() {
		if c.cmd == nil {
			c.err = errCmdNil
			return
		}
		if c.cmd.ProcessState != nil && c.cmd.ProcessState.Exited() {
			c.err = errCmdExited
			return
		}
		if c.cmd.Stdin != nil {
			c.err = errCmdStdin
			return
		}
		b, err := l.Load()
		if err != nil {
			c.err = err
			return
		}
		c.cmd.Stdin = bytes.NewReader(b)
		c.err = c.cmd.Run()
	})
	return c.err
}
