package arrays_test

import (
	"reflect"
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"Ex-1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"Ex-2", []int{3, 2, 1}, []int{1, 2, 3}},
		{"Ex-3", []int{1, 4, 3, 2}, []int{2, 1, 3, 4}},
		{"Ex-4", []int{2, 1, 4, 3}, []int{2, 3, 1, 4}},
		{"Ex-5", []int{2, 3, 1}, []int{3, 1, 2}},
	}

	t.Run("Simple Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				arrays.NextPermutation(test.nums)
				if !reflect.DeepEqual(test.nums, test.output) {
					t.Errorf("Got: %v, want: %v", test.nums, test.output)
				}
			})
		}
	})
}
