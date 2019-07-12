package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("lang arg Bengali", func(t *testing.T) {
		got := Hello("Amma", "Bengali")
		want := "Slamalaikum, Amma"
		assertCorrectMessage(t, got, want)
	})
	t.Run("lang arg French", func(t *testing.T) {
		got := Hello("Tasnima", "French")
		want := "Bonjour, Tasnima"
		assertCorrectMessage(t, got, want)
	})
	t.Run("lang arg Spanish", func(t *testing.T) {
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
