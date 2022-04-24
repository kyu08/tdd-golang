package test

import (
	"hello/cmd/reflection"
	"reflect"
	"testing"
)

type Person struct {
	Name Name
	City string
}

type Name struct {
	First string
	Last  string
}

func TestWalk(t *testing.T) {
	tests := []struct {
		Name           string
		X              interface{}
		ExpectedValues []string
	}{
		{
			"one string field",
			struct {
				Name string
			}{"john"},
			[]string{"john"},
		},
		{
			"struct with two string fields",
			struct {
				FirstName string
				LastName  string
			}{"john", "tanaka"},
			[]string{"john", "tanaka"},
		},
		{
			"nested struct",
			Person{
				Name: Name{"john", "tanaka"},
				City: "tokyo",
			},
			[]string{"john", "tanaka", "tokyo"},
		},
		{
			"pointers to things",
			&Person{
				Name: Name{"john", "tanaka"},
				City: "tokyo",
			},
			[]string{"john", "tanaka", "tokyo"},
		},
		{
			"slice",
			[]Person{
				{Name: Name{"john", "tanaka"},
					City: "tokyo",
				},
				{Name: Name{"u", "ishiguro"},
					City: "kyoto",
				},
			},
			[]string{"john", "tanaka", "tokyo", "u", "ishiguro", "kyoto"},
		},
		{
			"array",
			[2]Person{
				{Name: Name{"john", "tanaka"},
					City: "tokyo",
				},
				{Name: Name{"u", "ishiguro"},
					City: "kyoto",
				},
			},
			[]string{"john", "tanaka", "tokyo", "u", "ishiguro", "kyoto"},
		},
		{
			"map",
			map[string]string{
				"one": "ichi", "two": "ni",
			},
			[]string{"ichi", "ni"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			reflection.Walk(tt.X, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, tt.ExpectedValues) {
				t.Errorf("👺 got %v, want %v", got, tt.ExpectedValues)
			}
		})
	}
}
