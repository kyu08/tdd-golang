package test

import (
	"hello/cmd/integers"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("adder", func(t *testing.T) {
		sum := integers.Add(2, 2)

		expected := 4
		if sum != expected {
			t.Errorf("expected '%d', got '%d'", expected, sum)
		}
	})
}
