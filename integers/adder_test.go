package integers

import (
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(5, 9)
	want := 14

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}