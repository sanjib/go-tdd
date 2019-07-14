package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	assertByte := func(t *testing.T, got, want, given []byte) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v, given %#v", got, want, given)
		}
	}
	assertRune := func(t *testing.T, got, want, given []rune) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v, given %#v", got, want, given)
		}
	}
	assertString := func(t *testing.T, got, want, given []string) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v, given %#v", got, want, given)
		}
	}
	t.Run(`slice a byte slice`, func(t *testing.T) {
		given := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
		got := given[1:4]
		want := []byte{'o', 'l', 'a'}
		assertByte(t, got, want, given)
	})
	t.Run(`slice a rune slice`, func(t *testing.T) {
		given := []rune{'g', 'o', 'l', 'a', 'n', 'g'}
		got := given[:2]
		want := []rune{'g', 'o'}
		assertRune(t, got, want, given)
	})
	t.Run(`slice a string slice`, func(t *testing.T) {
		given := []string{"g", "o", "l", "a", "n", "g"}
		got := given[:2]
		want := []string{"g", "o"}
		assertString(t, got, want, given)
	})
	t.Run(`array copy via ref`, func(t *testing.T) {
		given := [3]string{"Лайка", "Белка", "Стрелка"}
		got := given[:]
		got[1] = "Belka"
		want := []string{"Лайка", "Belka", "Стрелка"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v, given %#v", got, want, given)
		}
	})
	t.Run(`array copy via value`, func(t *testing.T) {
		given := [3]string{"Лайка", "Белка", "Стрелка"}
		got := make([]string, len(given))
		copy(got, given[:])
		got[1] = "Belka"
		want := []string{"Лайка", "Belka", "Стрелка"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v, given %#v", got, want, given)
		}
	})
}
