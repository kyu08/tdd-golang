package test

import (
	"fmt"
	"hello/cmd/integers"
	"testing"
)

func ExampleAdd() {
	sum := integers.Add(1, 5)
	fmt.Println(sum)
}

func TestAdder(t *testing.T) {
	t.Run("adder", func(t *testing.T) {
		sum := integers.Add(2, 2)

		expected := 4
		if sum != expected {
			t.Errorf("👺 expected '%d', got '%d'", expected, sum)
		}
	})
}
