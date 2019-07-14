package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assert := func(t *testing.T, got, want int, given []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, want %d, given %#v", got, want, given)
		}
	}
	t.Run(`sum a slice of numbers`, func(t *testing.T) {
		given := []int{1, 2, 3}
		got := Sum(given)
		want := 6
		assert(t, got, want, given)
	})
}

func TestSumAll(t *testing.T) {
	assert := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, wanted %#v", got, want)
		}
	}
	t.Run(`sum slices of ints`, func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assert(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	assert := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, wanted %#v", got, want)
		}
	}
	t.Run(`sum tails of slices`, func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assert(t, got, want)
	})
}
