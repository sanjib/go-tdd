package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	write = "write"
	sleep = "sleep"
)

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	spySleeper := SpySleeper{}

	Countdown(&buffer, &spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, got %d, want %d", spySleeper.Calls, 4)
	}

	t.Run(`sleep before every print`, func(t *testing.T) {
		spySleepPrinter := CountdownOperationsSpy{}
		Countdown(&spySleepPrinter, &spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted %v, got %v", want, spySleepPrinter.Calls)
		}
	})
}
