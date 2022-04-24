package test

import (
	"hello/cmd/reflection"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	tests := []struct {
		name           string
		x              interface{}
		expectedValues []string
	}{
		{
			"one string field",
			struct {
				name string
			}{"john"},
			[]string{"john"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			reflection.Walk(tt.x, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, tt.expectedValues) {
				t.Errorf("got %v, want %v", got, tt.expectedValues)
			}
		})
	}
}
