package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := rectangle.Area()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

//func TestArea(t *testing.T) {
//
//	checkArea := func(t *testing.T, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//
//		if got != want {
//			t.Errorf("got %.2f want %.2f", got, want)
//		}
//	}
//
//	t.Run("rectangles", func(t *testing.T) {
//		rectangle := Rectangle{Width: 12.0, Height: 6.0}
//		checkArea(t, rectangle, 36.0)
//	})
//
//	t.Run("circles", func(t *testing.T) {
//		circle := Circle{10}
//		checkArea(t, circle, 314.1592653589793)
//	})
//}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 36},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %.2f want %.2f#v", tt.shape, got, tt.want)
		}
	}
}
