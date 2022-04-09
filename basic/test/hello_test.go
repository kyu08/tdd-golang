package test

import (
	"hello/src"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying Hello to people", func(t *testing.T) {
		got := src.Hello("ken", "Spanish")
		want := "Hola, ken"
		assertCorrectMessage(t, got, want)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := src.Hello("", "JavaScript")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("greet french ver", func(t *testing.T) {
		got := src.Hello("", "French")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})
}
