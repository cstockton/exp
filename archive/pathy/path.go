package pathy

import (
	"fmt"
	"os"
	stdPath "path"
	"sort"
	"strings"
	"unsafe"
)

var _ = fmt.Println

// Paths is just a slice of Path's that has a few convenience methods. The main
// purpose is simply to provide simple entry to other parts of the Pathy api.
// It isn't meant to be an efficient structure to store large numbers of paths.
// You should use a patricia trie or more specialized structure for that.
type Paths []Path

// PathFunc is used to filter and mutate Paths in PathLists and []Paths
type PathFunc func(p *Path) bool

// NewPaths todo
func NewPaths(ps ...Path) Paths {
	// Making a copy seems like a better practice than Paths(ps) ?
	//   How does the method below compare to copy() usage? same?
	return append(Paths(nil), ps...)
}

// NewPathsFromList todo
func NewPathsFromList(pl PathList) Paths {
	return pl.Paths()
}

func (ps Paths) Len() int           { return len(ps) }
func (ps Paths) Less(i, j int) bool { return ps[i] < ps[j] }
func (ps Paths) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

// Sort todo
func (ps Paths) Sort() Paths { sort.Sort(ps); return ps }

// Join todo
func (ps Paths) Join() Path {
	return Path(stdPath.Join(*(*[]string)(unsafe.Pointer(&ps))...))
}

// JoinList todo
func (ps Paths) JoinList() PathList {
	return NewPathList(ps...)
}

// Empty returns true if the list is empty, may remove this.. Len() > 0 is just
// as clear I think.
func (ps Paths) Empty() bool {
	return 0 == len(ps)
}

// Contains todo
func (ps Paths) Contains(path Path) bool {
	_, found := ps.Find(path)
	return found
}

// Find todo
func (ps Paths) Find(path Path) (int, bool) {
	for i, s := range ps {
		if s == path {
			return i, true
		}
	}
	return -1, false
}

// Search performs a binary search on the slice, so it must be sorted prior
// to being called. @TODO Maybe just remove this- it's not the purpose of this
// structure.. just doesn't really fit here.
func (ps Paths) Search(path Path) int {
	index := sort.Search(len(ps), func(i int) bool {
		return ps[i] >= path
	})

	return index
}

// Copy will make a copy of the Paths
func (ps Paths) Copy() Paths {
	return NewPaths(ps...)
}

// Unique will return a slice containing only unique entries
func (ps Paths) Unique() (paths Paths) {
	var empty struct{}
	seen := make(map[Path]struct{})

	for _, path := range ps {
		if _, ok := seen[path]; !ok {
			paths = append(paths, path)
			seen[path] = empty
		}
	}
	return
}

// Map will apply a mutator to each member of the path list and return new list
func (ps Paths) Map(mutator PathFunc) (paths Paths) {
	for _, path := range ps {
		if mutator(&path) {
			paths = append(paths, path)
		}
	}
	return
}

// Mutate will apply a PathFunc to each member and modify it in place
func (ps *Paths) Mutate(mutator PathFunc) {
	var paths Paths

	for _, path := range *ps {
		if mutator(&path) {
			paths = append(paths, path)
		}
	}

	*ps = paths
}

// Insert a Path at the given index
func (ps *Paths) Insert(index int, path Path) {
	*ps = append(*ps, ``)
	if index < len(*ps) {
		copy((*ps)[index+1:], (*ps)[index:])
	}
	(*ps)[index] = path
}

// Push a Path to the end of the list
func (ps *Paths) Push(path Path) {
	*ps = append(*ps, path)
}

// Pop removes and returns the last element in the PathList.
func (ps *Paths) Pop() (popped Path) {
	index := len(*ps) - 1
	popped, *ps = (*ps)[index], (*ps)[:index]
	return
}

// Shift removes and returns the first element in the list
func (ps *Paths) Shift() (shifted Path) {
	shifted, *ps = (*ps)[0], (*ps)[1:]
	return
}

// Unshift adds an element to the front of the list
func (ps *Paths) Unshift(path Path) {
	*ps = append(Paths{path}, *ps...)
}

// PathList is a string which contains one or more paths separated by the
// os.ListSeparator which is ; or : depending on platform. Like PathList it is
// just a simple alias providing mostly convenient entry back into the api.
type PathList string

// NewPathList todo
func NewPathList(ps ...Path) PathList {
	return PathList(strings.Join(*(*[]string)(unsafe.Pointer(&ps)), string(os.PathListSeparator)))
}

// Clean will be called on every Path of this PathList and return it.
// @TODO these are pretty ugly, should iter here with a buffer at some point
func (pl PathList) Clean() PathList {
	//return pl.Paths().Clean().JoinList()
	return pl.Paths().JoinList()
}

// Unique will return a new PathList containing only unique Paths
// @TODO these are pretty ugly, should iter here with a buffer at some point
func (pl PathList) Unique() PathList {
	return pl.Paths().Unique().JoinList()
}

// Paths returns Paths from this PathList
func (pl PathList) Paths() (paths Paths) {
	var cur int

	for i := 0; i < len(pl); i++ {
		if os.PathListSeparator == pl[i] {
			paths = append(paths, Path(pl[cur:i]))
			cur = i + 1
		}
	}

	return append(paths, Path(pl[cur:]))
}

// Path todo
type Path string

// Path satisfies Pathable interface
func (p Path) Path() Path {
	return p
}

// Paths returns Paths with itself as the only member
func (p Path) Paths() Paths {
	return NewPaths(p)
}

// String todo
func (p Path) String() string {
	return string(p)
}

// Pathy todo
func (p Path) Pathy() *Pathy {
	return &Pathy{}
}

// Clean todo
func (p Path) Clean() Path {
	return Path(stdPath.Clean(string(p)))
}

// Join todo
func (p Path) Join(paths ...Path) Path {
	paths = append(Paths{p}, paths...)
	return Path(stdPath.Join(*(*[]string)(unsafe.Pointer(&paths))...))
}

// Split todo
func (p Path) Split() (Path, Path) {
	return p.Parent(), p.Name()
}

// IsAbsolute provides access to Go standard library path.Clean()
func (p Path) IsAbsolute() bool {
	return len(p) > 0 && os.IsPathSeparator(p[0])
}

// IsRelative is !IsAbsolute
func (p Path) IsRelative() bool {
	return !p.IsAbsolute()
}

// Name todo
func (p Path) Name() Path {
	return p[p.NameIndex():]
}

// NameIndex todo
func (p Path) NameIndex() int {
	for i := len(p) - 1; i >= 0; i-- {
		if os.IsPathSeparator(p[i]) {
			return i + 1
		}
	}
	return 0
}

// Directory todo
func (p Path) Directory() Path {
	return p[0:p.DirectoryIndex()]
}

// DirectoryIndex todo
func (p Path) DirectoryIndex() int {
	for i := len(p) - 1; i >= 0; i-- {
		if os.IsPathSeparator(p[i]) {
			return i
		}
	}
	return len(p)
}

// Elements will return all path elements starting with an indicator of relative
// or absolute pathing i.e.:
//   /a/b/c/d -> (/, a, b, c, d)
//   a/b/c/d -> (., a, b, c, d)
func (p Path) Elements() (paths Paths) {
	var cur int
	p = p.Clean()

	if "." == p || (1 == len(p) && os.IsPathSeparator(p[0])) {
		return append(paths, p)
	}

	if !p.IsAbsolute() {
		paths = append(paths, ".")
	} else {
		paths = append(paths, Path(p[0]))
	}

	for i := 0; i < len(p); i++ {
		if os.IsPathSeparator(p[i]) {
			if cur != i {
				paths = append(paths, p[cur:i])
			}
			cur = i + 1
		}
	}

	return append(paths, p[cur:])
}

// Parents will return the logical ancestors of the path with the final entry
// being an indicator of relative or absolute pathing i.e.:
//   /a/b/c/d -> (/a/b/c, /a/b, /a, /)
//   a/b/c/d -> (a/b/c, a/b, a, .)
func (p Path) Parents() (paths Paths) {
	p = p.Clean()

	if 1 == len(p) && (p == "." || os.IsPathSeparator(p[0])) {
		return append(paths, p)
	}
	for i := len(p) - 1; i >= 0; i-- {
		if os.IsPathSeparator(p[i]) {
			if 0 == i {
				return append(paths, Path(p[i]))
			}
			paths = append(paths, p[:i])
		}
	}

	return append(paths, Path(`.`))
}

// Parent todo
func (p Path) Parent() Path {
	p = p.Clean()

	if 1 == len(p) && os.IsPathSeparator(p[0]) {
		return p
	}
	for i := len(p) - 1; i >= 0; i-- {
		if os.IsPathSeparator(p[i]) {
			if i == 0 {
				return Path(p[i])
			}
			return p[:i]
		}
	}

	return "."
}

// Stem todo
func (p Path) Stem() Path {
	begin, end := p.StemIndex()
	return p[begin:end]
}

// StemIndex todo
func (p Path) StemIndex() (begin, end int) {
	end = len(p)

	for i := end - 1; i >= 0; i-- {
		if p[i] == '.' {
			end = i
		}
		if os.IsPathSeparator(p[i]) {
			begin = i + 1
			break
		}
	}
	return begin, end
}

// Suffix todo
func (p Path) Suffix() Path {
	return p[p.SuffixIndex():]
}

// SuffixIndex todo
func (p Path) SuffixIndex() int {
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] == '.' {
			return i + 1
		}
		if os.IsPathSeparator(p[i]) {
			break
		}
	}
	return len(p)
}

// Suffixes todo
func (p Path) Suffixes() (suffixes Paths) {
	for cur, last := len(p)-1, 0; cur >= 0; cur-- {
		last++

		if p[cur] == '.' && last > 0 {
			suffixes = append(suffixes, p[cur+1:cur+last])
			last = 0
		}
		if os.IsPathSeparator(p[cur]) {
			break
		}
	}
	for i, j := 0, len(suffixes)-1; i < j; i, j = i+1, j-1 {
		suffixes[i], suffixes[j] = suffixes[j], suffixes[i]
	}
	return
}

// WithParent the directory with dir, if no slash is present it will add it
// i.e. given dir of `z`:
//   /a/b/c with "" -> c
//   /a/b/c with . -> c
//   /a/b/c with z -> z/c
//   /a/b/c with / -> /c
//   /a/b/c with /z -> /z/c
func (p Path) WithParent(dir Path) Path {
	return dir.Join(p.Name())
}

// WithName replaces or adds a stem
func (p Path) WithName(name Path) Path {
	index := p.NameIndex()
	return p[0:index] + name
}

// WithStem replaces or adds a stem
func (p Path) WithStem(stem Path) Path {
	begin, end := p.StemIndex()
	return p[0:begin] + stem + p[end:]
}

// WithSuffix replaces or adds a suffix
func (p Path) WithSuffix(suffix Path) Path {
	index, pl := p.SuffixIndex(), len(p)

	if index == pl {
		return p[0:index] + "." + suffix
	}
	return p[0:index] + suffix
}
