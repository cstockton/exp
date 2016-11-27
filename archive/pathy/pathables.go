package pathy

import (
	"bytes"
	"fmt"
	//stdPath "path"
	//"sort"
	//"unsafe"
)

var _ = fmt.Println

// Pathables is just a slice of Path's that has a few convienance methods. The main
// purpose is simply to provide simple entry to other parts of the Pathy api.
// It isn't meant to be an efficient structure to store large numbers of Pathables.
// You should use a patricia trie or more specialized structure for that.
type Pathables []Pathable

// PathableFunc is used to filter and mutate Pathables
type PathableFunc func(p *Pathable) bool

// NewPathables todo
func NewPathables(ps ...interface{}) (paths Pathables) {
	ch := Pathifier(ps...)

	for pathable := range ch {
		paths = append(paths, pathable)
	}
	return
}

// String todo
func (ps Pathables) String() string {
	return string(ps.PathList())
}

// Empty returns true if the list is empty, may remove this.. Len() > 0 is just
// as clear I think.
func (ps Pathables) Empty() bool {
	return 0 == len(ps)
}

// Contains todo
func (ps Pathables) Contains(path Pathable) bool {
	_, found := ps.Find(path)
	return found
}

// Find todo
func (ps Pathables) Find(path Pathable) (int, bool) {
	for i, s := range ps {
		if s == path {
			return i, true
		}
	}
	return -1, false
}

// StringSlice todo
func (ps Pathables) StringSlice() []string {
	paths := []string{}

	for _, p := range ps {
		paths = append(paths, p.String())
	}
	return paths
}

// PathList todo
func (ps Pathables) PathList() PathList {
	var b bytes.Buffer
	for _, p := range ps {
		b.WriteString(p.String())
	}
	return PathList(b.String())
}

// Unique will return a slice containing only unique entries
func (ps Pathables) Unique() (paths Pathables) {
	var empty struct{}
	seen := make(map[Pathable]struct{})

	for _, path := range ps {
		if _, ok := seen[path]; !ok {
			paths = append(paths, path)
			seen[path] = empty
		}
	}
	return
}

// Map will apply a mutator to each member of the path list and return new list
func (ps Pathables) Map(mutator PathableFunc) (paths Pathables) {
	for _, path := range ps {
		if mutator(&path) {
			paths = append(paths, path)
		}
	}
	return
}

// Mutate will apply a PathFunc to each member and modify it in place
func (ps *Pathables) Mutate(mutator PathableFunc) {
	var paths Pathables

	for _, path := range *ps {
		if mutator(&path) {
			paths = append(paths, path)
		}
	}

	*ps = paths
}

// Insert a Path at the given index
func (ps *Pathables) Insert(index int, path Pathable) {
	*ps = append(*ps, path)
	if index < len(*ps) {
		copy((*ps)[index+1:], (*ps)[index:])
	}
	(*ps)[index] = path
}

// Push a Path to the end of the list
func (ps *Pathables) Push(path Pathable) {
	*ps = append(*ps, path)
}

// Pop removes and returns the last element in the PathList.
func (ps *Pathables) Pop() (popped Pathable) {
	index := len(*ps) - 1
	popped, *ps = (*ps)[index], (*ps)[:index]
	return
}

// Shift removes and returns the first element in the list
func (ps *Pathables) Shift() (shifted Pathable) {
	shifted, *ps = (*ps)[0], (*ps)[1:]
	return
}

// Unshift adds an element to the front of the list
func (ps *Pathables) Unshift(path Pathable) {
	*ps = append(Pathables{path}, *ps...)
}
