package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
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

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

func TestCountdown(t *testing.T) {
	t.Run(`prints countdown from 3`, func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleepPrinter := CountdownOperationsSpy{}

		Countdown(&buffer, &spySleepPrinter)

		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

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
