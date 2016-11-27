package changes

import (
	"testing"
	"time"
)

func TestInterval(t *testing.T) {
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

func TestUntil(t *testing.T) {
	now := time.Now()
	interval := 2 * time.Millisecond
	until := (10 * interval) + time.Millisecond
	changes := Until(until, Interval(interval))
	expect := 10

	for i := 0; i < expect; i++ {
		select {
		case change := <-changes:
			assertChange(t, change)
			if now.After(change.When()) {
				t.Fatalf("change.When() was less than now")
			}
		case <-time.After(2 * interval):
			t.Fatalf("did not get change after %v milliseconds", interval)
		}
	}
}

func TestElapsed(t *testing.T) {
	now := time.Now()
	duration := 2 * time.Millisecond
	changes := Elapsed(duration)

	select {
	case change := <-changes:
		assertChange(t, change)
		if now.After(change.When()) {
			t.Fatalf("change.When() was less than now")
		}
	case <-time.After(2 * duration):
		t.Fatalf("did not get change after %v milliseconds", duration)
	}
}
