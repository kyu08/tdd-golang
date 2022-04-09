package test

import (
	"hello/cmd/iterator"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("empty string can be passed as arg", func(t *testing.T) {
		repeated := iterator.Repeat("", 4)
		expected := ""

		if repeated != expected {
			t.Errorf("👺 expected %s, got %s", repeated, expected)
		}
	})
	t.Run("repeat 5 times supplied string", func(t *testing.T) {
		repeated := iterator.Repeat("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("👺 expected %s, got %s", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iterator.Repeat("a", 5)
	}
}
