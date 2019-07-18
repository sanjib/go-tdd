package selectpkg

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 5 * time.Second

func RacerUsingConcurrency(a, b string) (string, error) {
	return RacerUsingConcurrencyWithTimeout(a, b, tenSecondTimeout)
}

func RacerUsingConcurrencyWithTimeout(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out after %#v waiting for %s and %s", timeout, a, b)
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
