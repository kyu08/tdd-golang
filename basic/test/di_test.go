package test

import (
	"bytes"
	"hello/cmd/di"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	di.Greet(&buffer, "ken")
	got := buffer.String()
	want := "Hello, ken"
	if got != want {
		t.Errorf("👺 got %s but want %s", got, want)
	}
}
