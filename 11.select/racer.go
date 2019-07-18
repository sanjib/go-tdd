package selectpkg

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration := measureGetResponse(a)
	bDuration := measureGetResponse(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureGetResponse(url string) time.Duration {
	start := time.Now()
	_, _ = http.Get(url)
	return time.Since(start)
}
