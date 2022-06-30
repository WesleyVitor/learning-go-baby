package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){
	

	t.Run("Should return a sum of 5 numbers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}
		got := Sum(numbers)
		expected := 15

		if got != expected{
			t.Errorf("got %d expected %d given %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T){
	slice1 := []int{1,2}
	slice2 := []int{1,2,3}

	got := SumAll(slice1, slice2)
	expected := []int{3,6}

	if !reflect.DeepEqual(got,expected){
			t.Errorf("got %d expected %d ", got, expected)
	}

}

func TestSumAllTail(t *testing.T){
	

	t.Run("Should return the sum without the head of slice with one value", func(t *testing.T) {
		slice1 := []int{1}
		slice2 := []int{1,2,3}

		got := SumAllTail(slice1, slice2)
		expected := []int{0,5}

		if !reflect.DeepEqual(got,expected){
				t.Errorf("got %d expected %d ", got, expected)
		}
	})

	t.Run("Should return the sum without the head of slice", func(t *testing.T) {
		slice1 := []int{1,2}
		slice2 := []int{1,2,3}

		got := SumAllTail(slice1, slice2)
		expected := []int{2,5}

		if !reflect.DeepEqual(got,expected){
				t.Errorf("got %d expected %d ", got, expected)
		}
	})

	t.Run("Should return the sum without the head of slice with zero values", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{1,2,3}

		got := SumAllTail(slice1, slice2)
		expected := []int{0,5}

		if !reflect.DeepEqual(got,expected){
				t.Errorf("got %d expected %d ", got, expected)
		}
	})
}

