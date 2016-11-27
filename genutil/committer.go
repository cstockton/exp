package genutil

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"sync"
)

var (
	errNilWriter = errors.New("genutil: io.Writer was nil")
	errNilReader = errors.New("genutil: io.Reader was nil")
	errNilLoader = errors.New("genutil: loader was nil")
)

type fileCommitter struct {
	m    sync.Mutex
	path string
}

func (c *fileCommitter) Stale(l Loader) bool {
	c.m.Lock()
	defer c.m.Unlock()
	if l == nil {
		return false
	}

	src, err := l.Load()
	if err != nil {
		return false
	}

	dst, err := ioutil.ReadFile(c.path)
	if err != nil {
		return false
	}
	return !bytes.Equal(dst, src)
}

func (c *fileCommitter) Commit(l Loader) error {
	c.m.Lock()
	defer c.m.Unlock()
	if l == nil {
		return errNilLoader
	}

	b, err := l.Load()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.path, b, 0644)
}

type ioCommitter struct {
	rOnce sync.Once
	wOnce sync.Once
	buf   []byte
	r     io.Reader
	w     io.Writer
	err   error
}

func (c *ioCommitter) Stale(l Loader) bool {
	if l == nil {
		return false
	}
	if c.err != nil || c.r == nil {
		return false
	}
	c.rOnce.Do(func() {
		c.buf, c.err = ioutil.ReadAll(c.r)
	})

	b, err := l.Load()
	if err != nil {
		return false
	}
	return !bytes.Equal(b, c.buf)
}

func (c *ioCommitter) Commit(l Loader) error {
	if l == nil {
		return errNilLoader
	}
	if c.w == nil {
		return errNilWriter
	}
	c.wOnce.Do(func() {
		src, err := l.Load()
		if err != nil {
			c.err = err
			return
		}
		_, c.err = io.Copy(c.w, bytes.NewReader(src))
	})
	return c.err
}
