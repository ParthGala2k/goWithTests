package main
import (
	"testing";
	"reflect"
)

func TestSum(t *testing.T) {
	t.Run("Running with fixed length of array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, %v", got, want, numbers)
		}
	})
}

/*
 Function to test for sum of many individual slices passed to
 SumAll() function which returns a slice of sums of those individual
 slices.
 *Note: all passed slices may be of different lengths
 reflect.DeepEqual compares two slices for equality, though it is not
 type safe ie. it will compile if both passed arguments are of different
 types although it may fail at runtime.
*/
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{4, 5, 6})
	want := []int{3, 15}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
