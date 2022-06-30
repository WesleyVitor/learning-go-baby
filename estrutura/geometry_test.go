package estrutura

import "testing"

func TestPerimeter(t *testing.T){
	
	t.Run("Should return the perimeter of rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 2.0, Height: 5.0}
		got := Perimeter(rectangle)
		expected := 14.0

		if got != expected{
			t.Errorf("got %.2f expected %.2f", got, expected)
		}
	})
}	


func TestArea(t *testing.T){

	checkArea := func(t testing.TB, shape Shape, want float64)  {
		t.Helper()
		got := shape.Area()
		if got != want{
			t.Errorf("got %.2f expected %.2f", got, want)
		}
	}

	t.Run("Should return the area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 2.0, Height: 5.0}
		expected := 10.0
		checkArea(t, rectangle,expected)
	})

	t.Run("Should return the area of circle", func(t *testing.T) {
		circle := Circle{Radius:10}
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})

	t.Run("Should return the area of triangle", func(t *testing.T) {
		triangle := Triangle{Base:2.0, Height:5}
		expected := 5.0
		checkArea(t, triangle, expected)
	})
}