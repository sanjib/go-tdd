package selectpkg

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacerUsingConcurrency(t *testing.T) {
	t.Run(`test fastest server`, func(t *testing.T) {
		slowServer := makeServer(10 * time.Millisecond)
		fastServer := makeServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := RacerUsingConcurrency(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run(`return error if server doesn't respond in 5s'`, func(t *testing.T) {
		serverA := makeServer(6 * time.Second)
		serverB := makeServer(7 * time.Second)
		defer serverA.Close()
		defer serverB.Close()

		_, err := RacerUsingConcurrencyWithTimeout(serverA.URL, serverB.URL, 5*time.Second)

		if err == nil {
			t.Error("was expecting an error but didn't get any")
		}
	})
}

func TestRacer(t *testing.T) {
	slowServer := makeServer(10 * time.Millisecond)
	fastServer := makeServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()

	want := fastServer.URL
	got := Racer(slowServer.URL, fastServer.URL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func makeServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
