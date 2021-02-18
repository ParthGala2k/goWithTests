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
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle", shape: Rectangle{20.0, 30.0},want: 600.0},
		{name: "Circle", shape: Circle{10},want: math.Pi * 100},
	}

	for _, testcase := range areaTests {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.shape.Area()
			if got != testcase.want {
				t.Errorf("%#v: got %g, want %g", testcase.shape, got, testcase.want)
			}
		})
	}
}
