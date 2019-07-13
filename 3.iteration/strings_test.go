package iteration

import (
	"reflect"
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	assert := func(t *testing.T, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	}
	t.Run("seafood contains foo", func(t *testing.T) {
		got := strings.Contains("seafood", "foo")
		want := true
		assert(t, got, want)
	})
	t.Run("seafood contains bar", func(t *testing.T) {
		got := strings.Contains("seafood", "bar")
		want := false
		assert(t, got, want)
	})
	t.Run(`seafood contains ""`, func(t *testing.T) {
		got := strings.Contains("seafood", "")
		want := true
		assert(t, got, want)
	})
}

func TestContainsRune(t *testing.T) {
	assert := func(t *testing.T, got, want bool) {
		t.Helper()
		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	}
	t.Run("aardvark contains rune 97 'a'", func(t *testing.T) {
		got := strings.ContainsRune("aardvark", 97)
		want := true
		assert(t, got, want)
	})
	t.Run("timeout contains rune 97 'a'", func(t *testing.T) {
		got := strings.ContainsRune("timeout", 97)
		want := false
		assert(t, got, want)
	})
	t.Run("সালাম contains rune 97 'স'", func(t *testing.T) {
		got := strings.ContainsRune("সালাম", 'স')
		want := true
		assert(t, got, want)
	})
}

func TestCount(t *testing.T) {
	assert := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("wanted %v, got %v", want, got)
		}
	}
	t.Run("cheese count e", func(t *testing.T) {
		got := strings.Count("cheese", "e")
		want := 3
		assert(t, got, want)
	})
	t.Run(`five count ""`, func(t *testing.T) {
		got := strings.Count("five", "")
		want := 5
		assert(t, got, want)
	})
}

func TestIndex(t *testing.T) {
	assert := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	}
	t.Run("chicken index ken", func(t *testing.T) {
		got := strings.Index("chicken", "ken")
		want := 4
		assert(t, got, want)
	})
	t.Run("chicken index dmr", func(t *testing.T) {
		got := strings.Index("chicken", "dmr")
		want := -1
		assert(t, got, want)
	})
}

func TestJoin(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
	t.Run("chicken index ken", func(t *testing.T) {
		s := []string{"foo", "bar", "baz"}
		got := strings.Join(s, ", ")
		want := "foo, bar, baz"
		assert(t, got, want)
	})
}

func TestSplit(t *testing.T) {
	assert := func(t *testing.T, got, want []string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, wanted %#v", got, want)
		}
	}
	t.Run("split abc with ,", func(t *testing.T) {
		got := strings.Split("a,b,c", ",")
		want := []string{"a", "b", "c"}
		assert(t, got, want)
	})
	t.Run(`split "a man a plan a canal panama" with "a "`, func(t *testing.T) {
		got := strings.Split("a man a plan a canal panama", "a ")
		want := []string{"", "man ", "plan ", "canal panama"}
		assert(t, got, want)
	})
	t.Run(`split " xyz " with ""`, func(t *testing.T) {
		got := strings.Split(" xyz ", "")
		want := []string{" ", "x", "y", "z", " "}
		assert(t, got, want)
	})
	t.Run(`split "" with "Bernardo O'Higgins"`, func(t *testing.T) {
		got := strings.Split("", "Bernardo O'Higgins")
		want := []string{""}
		assert(t, got, want)
	})
	t.Run(`split "moon" with "sun"`, func(t *testing.T) {
		got := strings.Split("moon", "sun")
		want := []string{"moon"}
		assert(t, got, want)
	})
	t.Run(`split "moon" with "jupitor"`, func(t *testing.T) {
		got := strings.Split("moon", "x")
		want := []string{"moon"}
		assert(t, got, want)
	})
}
