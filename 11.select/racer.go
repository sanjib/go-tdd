package selectpkg

import (
	"net/http"
	"time"
)

func RacerUsingConcurrency(a, b string) string {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func Racer(a, b string) string {
	aDuration := measureGetResponse(a)
	bDuration := measureGetResponse(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func ping(url string) chan bool {
	c := make(chan bool)
	go func(url string) {
		_, _ = http.Get(url)
		c <- true
	}(url)
	return c
}

func measureGetResponse(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}
