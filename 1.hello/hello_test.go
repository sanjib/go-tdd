package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("language arg Spanish", func(t *testing.T) {
		got := Hello("Qarnine", "Spanish")
		want := "Hola, Qarnine"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty name arg", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("name arg", func(t *testing.T) {
		got := Hello("Sanjib", "")
		want := "Hello, Sanjib"
		assertCorrectMessage(t, got, want)
	})

}
