package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration time.Duration
	SleepFn  func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.SleepFn(s.Duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		_, _ = fmt.Fprintln(writer, i)
	}

	sleeper.Sleep()
	_, _ = fmt.Fprint(writer, finalWord)
}
