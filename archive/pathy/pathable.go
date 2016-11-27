package pathy

import "fmt"

// Pathable is anything which is able to return Path.
// @TODO - maybe make a better name
type Pathable interface {
	Path() Path
	String() string
}

// Pathify will call PathifyChk and simply discards the error if one occurred.
func Pathify(p interface{}) Pathable {
	pathable, _ := PathifyChk(p)
	return pathable
}

// PathifyChk will return a single Pathable from an interface type, if no
// conversion was possible it will return an empty Path and an error.
func PathifyChk(p interface{}) (Pathable, error) {
	switch v := p.(type) {
	case string:
		return Path(v), nil
	case []byte:
		return Path(v), nil
	case Path, Pathy, Pathable:
		return v.(Pathable), nil
	case *Path, *Pathy, *Pathable:
		return *v.(*Pathable), nil
	case int, int8, int16, int32, int64:
		return Path(fmt.Sprint(v)), nil
	case uint, uint8, uint16, uint32, uint64:
		return Path(fmt.Sprint(v)), nil
	case float32, float64, complex64, complex128:
		return Path(fmt.Sprint(v)), nil
	default:
		return Path(``), fmt.Errorf("could not convert %v because type %T is not pathable", v, v)
	}
}

// Pathifier will return a channel which will return a channel of Pathable and
// produce an item for each interface that was able to be converted by PathifyChk.
func Pathifier(ps ...interface{}) chan Pathable {
	ch := make(chan Pathable)

	sender := func(p interface{}) {
		pathable, err := PathifyChk(p)

		if err == nil {
			ch <- pathable
		}
	}

	go func() {

		for _, p := range ps {
			// I could not think of a better way to write this.. although I udnerstand
			// []T{} -> []interface{} conversion to a degree.. this pattern strikes
			// me as bizarre and really cumbersome to write. Probably is a more
			// idiomatic way to do this- or what I am doing is odd to begin with...
			switch v := p.(type) {
			case []string:
				for _, s := range v {
					sender(s)
				}
			case []Path:
				for _, s := range v {
					sender(s)
				}
			case Paths:
				for _, s := range v {
					sender(s)
				}
			case []Pathable:
				for _, s := range v {
					sender(s)
				}
			case Pathables:
				for _, s := range v {
					sender(s)
				}
			default:
				sender(p)
			}
		}
		close(ch)
	}()

	return ch
}
