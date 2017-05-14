// Package call provides a simple interface for safely calling functions (quickly)
// at runtime while retaining types for the callee.
package call

import "errors"

var (
	// ErrArgCount is returned by generated functions for arity mismatch.
	ErrArgCount = errors.New("wrong argument count")

	// ErrArgType is returned by generated functions when they are unable to use
	// the given argument.
	ErrArgType = errors.New("wrong argument type")
)

// Caller accepts variadic interface arguments and will call a underlying
// function with them if they match the types appropriately. If the arguments
// are invalid or the call failed for other reasons it will report why.
type Caller interface {
	Call(in ...interface{}) error
}

// New returns a new Caller backed the given function fn. If the given function
// returns a non-nil error value it will be propagated to the Caller interafce.
func New(fn interface{}) (Caller, error) {
	c, err := defaultAnalyzer.Analyze(fn)
	if err != nil {
		return nil, err
	}
	return c, nil
}
