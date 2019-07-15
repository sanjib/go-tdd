package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	assert := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %#f, want %#f", got, want)
		}
	}
	t.Run(`rectangle get perimeter`, func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0
		assert(t, got, want)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %#v, want %#v", got, tt.want)
		}
	}
}
