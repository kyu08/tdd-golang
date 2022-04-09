package test

import (
	"hello/cmd/array"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := array.Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("👺 expected '%d', but got '%d'", want, got)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := array.Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("👺 expected '%d', but got '%d'", want, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("a", func(t *testing.T) {
		got := array.SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("👺 expected '%d', but got '%d'", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("a", func(t *testing.T) {
		got := array.SumAllTails([]int{}, []int{3, 4}, []int{1, 2, 3})
		want := []int{0, 4, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("👺 expected '%d', but got '%d'", want, got)
		}
	})
}
