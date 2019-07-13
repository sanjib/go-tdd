package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("repeat character 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assert(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}
