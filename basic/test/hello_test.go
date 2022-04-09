package test

import (
	"hello/src"
	"testing"
)

func TestHello(t *testing.T) {
	got := src.Hello("ken")
	want := "hello, ken"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
