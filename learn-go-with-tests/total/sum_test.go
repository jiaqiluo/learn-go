package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	//t.Run("collection of any size", func(t *testing.T) {
	//	numbers := []int{1, 2, 3}
	//
	//	got := Sum(numbers)
	//	want := 6
	//
	//	if got != want {
	//		t.Errorf("got %d want %d given, %v", got, want, numbers)
	//	}
	//})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("simple slides", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("complex slides", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 4}, []int{0, 9, 8, 7})
		want := []int{9, 24}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("empty slide", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 8, 7})
		want := []int{0, 24}
		checkSums(t, got, want)
	})

	t.Run("one item", func(t *testing.T) {
		got := SumAllTails([]int{2}, []int{0, 9, 8, 7})
		want := []int{0, 24}
		checkSums(t, got, want)
	})
}
