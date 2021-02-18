package main
import (
	"testing"
	"math"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{20.0, 30.0}
	got := Perimeter(rectangle)
	want := 100.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{20.0, 30.0}, 600.0},
		{Circle{10}, math.Pi * 100},
	}

	for _, testcase := range areaTests {
		got := testcase.shape.Area()
		if got != testcase.want {
			t.Errorf("got %g, want %g", got, testcase.want)
		}
	}
}
