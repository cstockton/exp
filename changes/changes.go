package changes

import (
	"fmt"
	"time"
)

// FilterFunc todo
type FilterFunc func(change Interface) bool

// Changes todo
type Changes chan Interface

// Interface todo
type Interface interface {
	When() time.Time
	String() string
}

// Change todo
type Change struct {
	when time.Time
}

// When todo
func (c Change) When() time.Time {
	return c.when
}

// String todo
func (c Change) String() string {
	return fmt.Sprintf("Change(When: %v)", c.When())
}

// NewChange todo
func NewChange() *Change {
	return &Change{
		when: time.Now(),
	}
}
