package test

import (
	"hello/cmd/racer"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(200 * time.Millisecond)


	プロセスの同期
  から！！！！！！


	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := racer.Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
