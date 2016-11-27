package changes

import (
	"sync"
	"time"
)

// When the first change happens, return the Changes channel
func When(when <-chan Interface, changes ...<-chan Interface) (Interface, <-chan Interface) {
	return Fetch(when), Merge(changes...)
}

// Fetch a single change from a chan
func Fetch(changes <-chan Interface) (out Interface) {
	select {
	case change, ok := <-changes:
		if !ok {
			return
		}
		out = change
	}

	return out
}

// Filter todo
func Filter(filter FilterFunc, changes ...<-chan Interface) <-chan Interface {
	out := make(chan Interface)

	go func() {
		for change := range Merge(changes...) {
			if filter(change) {
				out <- change
			}
		}
	}()

	return out
}

// Reject todo
func Reject(filter FilterFunc, changes ...<-chan Interface) <-chan Interface {
	out := make(chan Interface)

	go func() {
		for change := range Merge(changes...) {
			if !filter(change) {
				out <- change
			}
		}
	}()

	return out
}

// Merge todo
func Merge(changes ...<-chan Interface) <-chan Interface {
	var wg sync.WaitGroup
	out := make(chan Interface)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan Interface) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(changes))
	for _, c := range changes {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Throttle returns every change no faster than the supplied rate
func Throttle(rate time.Duration, changes ...<-chan Interface) <-chan Interface {
	out := make(chan Interface)
	throttle := time.Tick(rate)
	merged := Merge(changes...)

	go func() {
		for {
			select {
			case change, ok := <-merged:
				if !ok {
					return
				}
				<-throttle
				out <- change
			}
		}
	}()

	return out
}

// Discard should discard anything outside of rate
func Discard(rate time.Duration, changes ...<-chan Interface) <-chan Interface {
	out := make(chan Interface)
	throttle := time.Tick(rate)
	merged := Merge(changes...)

	// Is there a better pattern for this?
	go func() {
		out <- <-merged

		for {
			select {
			case change, ok := <-merged:
				if !ok {
					return
				}
				select {
				case <-throttle:
					out <- change
				default:

				}
			}
		}
	}()

	return out
}
