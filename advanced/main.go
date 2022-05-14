package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	score map[string]int
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.score[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
}

func main() {
	s := &PlayerServer{store: &InMemoryPlayerStore{
		score: map[string]int{
			"14": 10,
			"94": 20,
		},
	}}

	handler := http.HandlerFunc(s.ServeHTTP)
	log.Fatal(http.ListenAndServe("127.0.0.1:5001", handler))
}
