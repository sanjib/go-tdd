package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	assert := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %#f, want %#f", got, want)
		}
	}
	t.Run(`get perimeter from length and width`, func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0
		assert(t, got, want)
	})
}

func TestArea(t *testing.T) {
	assert := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %#v, want %#v", got, want)
		}
	}
	t.Run(`get area from width and height`, func(t *testing.T) {
		got := Area(12.0, 6.0)
		want := 72.0
		assert(t, got, want)
	})
}
