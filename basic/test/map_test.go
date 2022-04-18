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
		assertError(t, err, maps.ErrNotFound)
	})
}











`map_test.go:18: 👺 got ERR: value not exist. but want ERR: value not exist.`
ってなるってことはもしかして error 同士の比較ってうまくできないのかな？







func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := maps.Dictionary{}
		key := "editor"
		value := "vim"
		dictionary.Add(key, value)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := maps.Dictionary{"one": "ichi"}
		err := dictionary.Add("one", "uno")
		if err != nil {
			t.Errorf("👺 got %s but want %s", err, maps.ErrWordExists)
		}
	})
}

// helper
func assertDefinition(t *testing.T, dictionary maps.Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("👺 should find added value: ", err)
	}
	if got != definition {
		t.Errorf("👺 got %q but want %q", got, definition)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("👺 got %q but want %q", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got != nil {
		t.Errorf("👺 got %s but want %s", got, want)
		return
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Errorf("👺 got nil but want %s", want)
	}
}
