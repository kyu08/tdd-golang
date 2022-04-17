package test

import (
	"hello/cmd/maps"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "test dayo"}
	got := maps.Search(dictionary, "test")
	want := "test dayo"
	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("👺 got %q but want %q", got, want)
	}
}
