package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(c context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, char := range s.response {
			select {
			case <-c.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(char)
			}
		}
		data <- result
	}()

	select {
	case <-c.Done():
		return "", c.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestHandler(t *testing.T) {
	t.Run(`simple store test`, func(t *testing.T) {
		data := "Hello, world!"
		spyStore := SpyStore{response: data, t: t}
		server := Server(&spyStore)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf("got %q, want %q", res.Body.String(), data)
		}
	})
	t.Run(`store should cancel work`, func(t *testing.T) {
		store := SpyStore{response: "foo bar", t: t}
		server := Server(&store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := &SpyResponseWriter{}

		cancelContext, cancelFn := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancelFn)
		req = req.WithContext(cancelContext)

		server.ServeHTTP(res, req)
		if res.written {
			t.Error("a response should not have been written")
		}
	})
}
