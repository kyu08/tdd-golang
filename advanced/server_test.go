package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type spyPlayerStore struct {
	score    map[string]int
	winCalls []string
}

func (s *spyPlayerStore) GetPlayerScore(name string) int {
	return s.score[name]
}

func (s *spyPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := spyPlayerStore{
		score: map[string]int{
			"14": 10,
			"94": 20,
		},
	}
	server := &PlayerServer{store: &store}

	t.Run("returns 94's score", func(t *testing.T) {
		request := newGetScoreRequest("94")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertScore(t, got, want)
		assertResponseCode(t, response.Code, http.StatusOK)
	})

	t.Run("returns 14's score", func(t *testing.T) {
		request := newGetScoreRequest("14")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertScore(t, got, want)
		assertResponseCode(t, response.Code, http.StatusOK)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("zeb")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseCode(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := spyPlayerStore{
		score:    map[string]int{},
		winCalls: nil,
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		name := "14"
		request, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseCode(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("👺 got %d but want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != name {
			t.Errorf("👺 got %s but want %s", store.winCalls[0], name)
		}
	})
}

/*
helpers
*/

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return request
}

func assertResponseCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("👺 got %d, want %d", got, want)
	}
}

func assertScore(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("👺 got %q, want %q", got, want)
	}
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return request
}
