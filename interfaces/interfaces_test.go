package interfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 20.0}
	got := Perimeter(rect)
	want := 60.0

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestArea(t *testing.T) {
	assertArea := func(t testing.TB, s Shape, want float64) {
		t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 20.0}
		assertArea(t, rect, 200.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		assertArea(t, circle, 314.1592653589793)
	})
}