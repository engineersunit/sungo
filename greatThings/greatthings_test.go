package greatthings_test

import (
	"testing"

	"org.sun.ghosh.golang/go-dev-1/greatthings"
)

func TestGreatThings(t *testing.T) {
	if greatthings.PopGreatThings() != "Learning" {
		t.Fatal("Not a great thing")
	}
}
