package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rect1 := Rectangle{10.0, 10.0}

	got := rect1.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got: %.2f, want: %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got: %g, want: %g", got, want)
		}
	}

	t.Run("rectangle", func (t *testing.T) {
		rect2 := Rectangle{12.0, 6.0}
		checkArea(t, rect2, 72.0)
	})

	t.Run("circle", func (t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestTableArea(t *testing.T) {
	testCases := []struct{
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("input: %v, got: %g, want: %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}