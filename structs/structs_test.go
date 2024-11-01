package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}

	got := Perimeter(r)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		Name    string
		Shape   Shape
		HasArea float64
	}{
		{Name: "Rectangle", Shape: Rectangle{Width: 10.0, Height: 10.0}, HasArea: 100},
		{Name: "Circle", Shape: Circle{Radius: 10}, HasArea: 314.1592653589793},
		{Name: "Triangle", Shape: Triangle{Base: 12, Height: 6}, HasArea: 36.0},
	}

	for _, v := range areaTests {
		t.Run(v.Name, func(t *testing.T) {
			a := v.Shape.Area()

			if a != v.HasArea {
				t.Errorf("%#v got %g want %g", v.Shape, a, v.HasArea)
			}
		})
	}
}
