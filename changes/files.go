package changes

import (
	"fmt"
	"time"
)

// FileChange todo
type FileChange struct {
	when time.Time
	path string
}

// When todo
func (c FileChange) When() time.Time {
	return c.when
}

// Path todo
func (c FileChange) Path() string {
	return c.path
}

// String todo
func (c FileChange) String() string {
	return fmt.Sprintf("Change(When: %v, Path: %v)", c.When(), c.Path())
}

// NewFileChange todo
func NewFileChange(path string) *FileChange {
	return &FileChange{
		when: time.Now(),
		path: path,
	}
}
