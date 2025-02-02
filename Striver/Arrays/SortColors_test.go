package arrays_test

import (
	"reflect"
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"Ex-1", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"Ex-2", []int{2, 0, 1}, []int{0, 1, 2}},
		{"Ex-3", []int{2, 1}, []int{1, 2}},
	}

	t.Run("Counting sort", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := make([]int, len(test.nums))

				copy(input, test.nums)
				arrays.SortColorsCounting(input)
				if !reflect.DeepEqual(input, test.output) {
					t.Errorf("Got: %v, want: %v", input, test.output)
				}
			})
		}
	})

	t.Run("Three Pointers Solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := make([]int, len(test.nums))

				copy(input, test.nums)
				arrays.SortColorsThreePtrs(input)
				if !reflect.DeepEqual(input, test.output) {
					t.Errorf("Got: %v, want: %v", input, test.output)
				}
			})
		}
	})
}
