package test

import (
	"hello/cmd/structs"
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape structs.Shape
		want  float64
	}{
		{"rectangle", structs.Rectangle{12, 6}, 72.0},
		{"circle", structs.Circle{10}, 314.1592653589793},
		{"triangle", structs.Triangle{3, 2}, 3},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("👺 got %v, but want %v", got, tt.want)
			}
		})
	}

	checkArea := func(t *testing.T, shape structs.Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangle area", func(t *testing.T) {
		rectangle := structs.Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := structs.Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
