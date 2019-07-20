package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run(`increment counter 3 times`, func(t *testing.T) {
		counter := NewCounter()
		counter.Add()
		counter.Add()
		counter.Add()

		want := 3
		assertCounter(t, counter, want)
	})
	t.Run(`increment counter concurrently`, func(t *testing.T) {
		want := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Add()
				wg.Done()
			}(&wg)
		}
		wg.Wait()
		assertCounter(t, counter, want)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
