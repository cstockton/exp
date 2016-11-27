package changes

import (
	"time"
)

// Elapsed todo
func Elapsed(elapsed time.Duration) Changes {
	changes := make(Changes)
	after := time.After(elapsed)

	go func() {
		defer close(changes)
		changes <- Change{
			when: <-after,
		}
	}()

	return changes
}

// Interval todo
func Interval(interval time.Duration) Changes {
	changes := make(Changes)
	ticker := time.Tick(interval)

	go func() {
		for now := range ticker {
			changes <- Change{
				when: now,
			}
		}
	}()

	return changes
}

// Until todo
func Until(until time.Duration, changes Changes) Changes {
	changesOut := make(Changes)
	after := time.After(until)

	go func() {
		for {
			select {
			case change := <-changes:
				changesOut <- change
			case <-after:
				close(changesOut)
				return
			}
		}
	}()

	return changesOut
}
