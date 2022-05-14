package main

import (
	"fmt"
	"net/http"
	"strings"
)

const playerPrefix = "/players/"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.storeScore(w, r)
	case http.MethodGet:
		s.showScore(w, r)
	}
}

func (s *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := getPlayerName(r)
	score := s.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) storeScore(w http.ResponseWriter, r *http.Request) {
	s.processWin(w, getPlayerName(r))
}

func (s *PlayerServer) processWin(w http.ResponseWriter, name string) {
	s.store.RecordWin(name)
	w.WriteHeader(http.StatusAccepted)
}

func getPlayerName(r *http.Request) string {
	return strings.TrimPrefix(r.URL.Path, playerPrefix)
}
