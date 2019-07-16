package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run(`search test`, func(t *testing.T) {
		dictionary := map[string]string{
			"test": "this is just a test",
		}
		got := Search(dictionary, "test")
		want := "this is just a test"
		assertString(t, got, want)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
