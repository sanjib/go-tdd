package sync

import "testing"

func TestCounter(t *testing.T) {
	counter := Counter{}
	counter.Add()
	counter.Add()
	counter.Add()

	got := counter.Value()
	want := 3

	assertCounter(t, got, want)
}

func assertCounter(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
