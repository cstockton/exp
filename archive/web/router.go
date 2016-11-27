package web

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"

	"code.google.com/p/go.net/context"
)

// Consts
const (
	RequestTimeout time.Duration = time.Second * 2
	RouteParam                   = 1 << iota
	QueryParam
)

// HandleFunc is a handler function for our router
type HandleFunc func(Context, http.ResponseWriter, *http.Request)

// PanicHandleFunc todo
type PanicHandleFunc func(Context, http.ResponseWriter, *http.Request, interface{})

// Pattern interface
type Pattern interface {
	IsMatch(Context, string) bool
}

// LiteralPattern struct
type LiteralPattern struct {
	Literal string
}

// RegexpPattern struct
type RegexpPattern struct {
	Regexp *regexp.Regexp
}

// AnyPattern struct
type AnyPattern struct{}

// Route to a handle func
type Route struct {
	Method  string
	Pattern Pattern
	Handler HandleFunc
}

// IsMatch bool
func (s LiteralPattern) IsMatch(ctx Context, path string) bool {
	return path == s.Literal
}

// IsMatch bool
func (s RegexpPattern) IsMatch(ctx Context, path string) bool {
	fmt.Printf("  The regexp are  : %v\n", s.Regexp)
	fmt.Printf("  The path are  : %v\n", path)

	names := s.Regexp.SubexpNames()
	matches := s.Regexp.FindAllStringSubmatch(path, -1)

	if nil == matches || 1 > len(matches) {
		log.Printf("RegexpPattern :: IsMatch :: no matches")
		return false
	}

	for index, value := range matches[0][1:] {
		fmt.Printf("  params :: Index(%d) Name(%s) Value(%s)\n", index, names[index+1], value)

		(*ctx.Params) = (*ctx.Params)[:len(*ctx.Params)+1]
		(*ctx.Params)[index].Name = names[index+1]
		(*ctx.Params)[index].Value = value
	}

	return true
}

// IsMatch bool
func (s AnyPattern) IsMatch(ctx Context, path string) bool {
	return true
}

// IsMatch todo
func (r *Route) IsMatch(ctx Context, path string) bool {
	log.Printf("Route :: IsMatch :: for Method(%s) path(%s)", r.Method, path)

	return r.Pattern.IsMatch(ctx, path)
}

// Router is the http.Handler
type Router struct {
	Routes         []Route
	PanicHandler   PanicHandleFunc
	StatusHandlers map[int]HandleFunc
}

// DefaultPanicHandler todo
func DefaultPanicHandler(ctx Context, w http.ResponseWriter, req *http.Request, recover interface{}) {
	fmt.Fprintln(w, "DefaultPanicHandler")
}

// DefaultNotFoundHandler todo
func DefaultNotFoundHandler(ctx Context, w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "DefaultNotFoundHandler")
}

// DefaultInternalServerErrorHandler todo
func DefaultInternalServerErrorHandler(ctx Context, w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "DefaultInternalServerErrorHandler")
}

// New returns a initialized router
func New() *Router {
	var statusHandlers = map[int]HandleFunc{
		http.StatusNotFound:            DefaultNotFoundHandler,
		http.StatusInternalServerError: DefaultInternalServerErrorHandler,
	}

	return &Router{Routes: make([]Route, 0), PanicHandler: DefaultPanicHandler, StatusHandlers: statusHandlers}
}

// PathWalker iterable that walks segments of a url
func (r *Router) PathWalker(path string, f func(segment string)) {
	pathLen := len(path)
	position := 0
	width := 0
	segment := ""

	// Peek
	p := func() (string, int) {
		runeValue, runeLength := utf8.DecodeRuneInString(path[position:])

		return string(runeValue), runeLength
	}

	// Peek behind
	pb := func() string {
		if width > 0 {
			runeValue, _ := utf8.DecodeRuneInString(path[position-width:])

			return string(runeValue)
		}

		return ""
	}

	// Yield a path segment to caller
	yield := func() {
		if 0 < len(segment) {
			f(segment)
			segment = ""
		}
	}

	for ; position < pathLen; position += width {
		p, runeLength := p()
		pb := pb()
		width = runeLength

		// Escape sequence, not used yet, may be removed
		if "/" == p && "\\" == pb {
			continue
		} else if "/" == p {
			yield()
		} else {
			segment += p
		}
	}

	yield()
}

// StringToRegexp converts string path to regexp
func (r *Router) StringToRegexp(path string) *regexp.Regexp {

	// all paths must start with a /
	if 47 != path[0] {
		panic("Bad path, must start with forward slash")
	}

	// start reg exp
	expstr := "^"

	r.PathWalker(path, func(segment string) {
		c := segment[0]

		if 58 == c { // ":"
			expstr += "\\/(?P<" + segment[1:] + ">[^\\/]+?)"

		} else if 42 == c { // "*"
			clen := len(segment[1:])

			if 0 < clen {
				expstr += "\\/.{1," + strconv.Itoa(len(segment[1:])+1) + "}"
			} else {
				expstr += "\\/.+"
			}
		} else {
			expstr += "\\/" + segment
		}
	})

	// close regexpr
	expstr += "$"

	fmt.Printf("      StringToRegexp :: path(%s) regexp(%s)\n", path, expstr)

	regexp := regexp.MustCompile(expstr)
	regexp.Longest()

	return regexp
}

// Register adds a HandleFunc to the router
func (r *Router) Register(method string, pattern interface{}, handle HandleFunc) {

	// Create a route with proper method and handler
	route := &Route{Method: method, Handler: handle}

	// pattern can be a regexp, string, or Pattern interface
	switch pattern.(type) {
	case string:
		route.Pattern = &RegexpPattern{Regexp: r.StringToRegexp(pattern.(string))}
	case regexp.Regexp:
		expr := pattern.(regexp.Regexp)

		route.Pattern = &RegexpPattern{Regexp: &expr}
	case *regexp.Regexp:
		route.Pattern = &RegexpPattern{Regexp: pattern.(*regexp.Regexp)}
	case Pattern:
		route.Pattern = pattern.(Pattern)
	default:
		panic("Bad route, unknown type.")
	}

	// append the route to our Routers known routes
	r.Routes = append(r.Routes, *route)

}

// GET is a Register call for GET requests
func (r *Router) GET(path string, handle HandleFunc) {
	r.Register("GET", path, handle)
}

// FindRoute todo
func (r *Router) FindRoute(ctx Context, method string, path string) (*Route, bool) {
	var best Route

	for _, route := range r.Routes {
		log.Printf("Router :: FindRoute :: Checking route :: %s :: %s ::", route.Method, method)

		if route.Method == method && route.IsMatch(ctx, path) {
			log.Printf("Router :: FindRoute :: IsMatch")

			return &route, true
		}
	}

	return &best, false
}

// Param contains a router path argument
type Param struct {
	Name  string
	Value string
}

// Params is a slice of Param's
type Params []Param

// ByName returns a Param by name
func (p Params) ByName(name string) string {
	for _, param := range p {
		if name == param.Name {
			return param.Value
		}
	}

	return ""
}

// ByIndex returns a Param by index
func (p Params) ByIndex(index int) string {
	if len(p)-1 >= index {
		return p[index].Value
	}

	return ""
}

// Context google context wrapper
type Context struct {
	context.Context
	Params *Params
}

// NewContext returns a new router context
// https://blog.golang.org/context
func NewContext() (*Context, context.CancelFunc) {

	// root context is a google context with a timeout func
	rootCtx, cancel := context.WithTimeout(context.Background(), RequestTimeout)

	// create a web router Context
	params := make(Params, 0, 255)
	outerContext := Context{rootCtx, &params}

	return &outerContext, cancel
}

// http://golang.org/pkg/net/http/#Server
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// Basic request setup
	path := req.URL.Path
	method := req.Method
	ctx, cancel := NewContext()

	// We cancel our context on exit
	defer cancel()

	log.Printf("Router :: ServeHTTP :: path(%s) method(%s)", path, method)

	// Panic handler
	defer func(ctx Context, w http.ResponseWriter, req *http.Request) {
		if rcv := recover(); rcv != nil {
			r.PanicHandler(ctx, w, req, rcv)
			log.Printf("Router :: panic :: recover(%s)", rcv)
		}
	}(*ctx, w, req)

	// Look for a valid route
	route, ok := r.FindRoute(*ctx, method, path)

	if ok {

		// We got a route, process request
		route.Handler(*ctx, w, req)

	} else {

		// No route found, use that notfound handler
		log.Println("Router :: ServeHTTP :: No Route")

		if val, ok := r.StatusHandlers[http.StatusNotFound]; ok {
			val(*ctx, w, req)
		}
	}

	log.Println("Router :: ServeHTTP :: done")

}
