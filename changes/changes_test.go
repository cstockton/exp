package changes

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	seed := int64(1)
	fmt.Println(seed)
	rand.Seed(seed)
}

func assertChange(t *testing.T, change Interface) {
	now := time.Now()
	when := change.When()

	if nil == change {
		t.Fatalf("change was nil")
	}
	if now.Before(when) {
		t.Fatalf("change.When() was before Now()")
	}
}

func TestChanges(t *testing.T) {

}
