package test

import (
	"hello/cmd/maps"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("search existing value", func(t *testing.T) {
		dictionary := maps.Dictionary{"test": "test dayo"}
		got, _ := dictionary.Search("test")
		want := "test dayo"
		assertStrings(t, got, want)
	})
	t.Run("search non-existing value", func(t *testing.T) {
		dictionary := maps.Dictionary{"test": "test dayo"}
		_, err := dictionary.Search("nekonekoneko")
		assertError(t, err)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("👺 got %q but want %q", got, want)
	}
}

func assertError(t *testing.T, got error) {
	t.Helper()
	if got == nil {
		t.Errorf("👺 got nil but want error")
	}
}
