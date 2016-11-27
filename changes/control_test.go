package changes

import (
	"testing"
	"time"
)

func TestFilter(t *testing.T) {
	now := time.Now()
	interval := time.Millisecond
	changes := Interval(interval)
	expect := 10

	for i := 0; i < expect; i++ {
		select {
		case change := <-changes:
			assertChange(t, change)
			if now.After(change.When()) {
				t.Fatalf("change.When() was less than now")
			}
		case <-time.After(20 * interval):
			t.Fatalf("did not get change after %v milliseconds", interval)
		}
	}
}
