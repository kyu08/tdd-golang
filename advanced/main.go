package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) string {
	return "123"
}
func main() {
	s := &PlayerServer{store: &InMemoryPlayerStore{}}

	handler := http.HandlerFunc(s.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5001", handler))
}
