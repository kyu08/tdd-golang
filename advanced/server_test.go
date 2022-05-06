package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type spyPlayerStore struct {
	score map[string]string
}

func (s *spyPlayerStore) GetPlayerScore(name string) string {
	return s.score[name]
}

func TestGETPlayers(t *testing.T) {
	store := spyPlayerStore{score: map[string]string{"14": "10", "94": "20"}}
	server := &PlayerServer{store: &store}

	t.Run("returns 94's score", func(t *testing.T) {
		request := newGetScoreRequest("94")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertScore(t, got, want)
	})

	t.Run("returns 14's score", func(t *testing.T) {
		request := newGetScoreRequest("14")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertScore(t, got, want)
	})
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)

	return request
}

func assertScore(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("👺 got %q, want %q", got, want)
	}
}
