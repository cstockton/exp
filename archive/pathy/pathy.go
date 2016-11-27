package pathy

import (
	"os"
	//"strings"
)

// Pathys is a slice of Pathy
// @TODO maybe a better name would be Pathies? This name matches Paths but it
//       feels a bit off.
type Pathys []*Pathy

// File todo
type File struct {
	*os.File
}

// Pathy todo
type Pathy struct {
	// This is not embedded because I want to have Path() reserved for the pathable
	// interface.. @TODO maybe decide a different way to do this, embedded would
	// probably be better, but then we need a new name.
	path     Path
	file     *File
	children []*Pathy
}

// New todo
func New(pathable interface{}) *Pathy {
	return &Pathy{Pathify(pathable).Path(), nil, nil}
}

// FromString todo
func FromString(path string) *Pathy {
	return &Pathy{Path(path), nil, nil}
}

// Path todo
func (p Pathy) Path() Path {
	return p.path
}

// String todo
func (p Pathy) String() string {
	return string(p.path)
}
