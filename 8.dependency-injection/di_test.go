package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sanjib")
	got := buffer.String()
	want := "Hello, Sanjib"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
