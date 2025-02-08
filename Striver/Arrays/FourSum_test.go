package arrays_test

import (
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestFourSum(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		output [][]int
	}{
		{"Ex-1", []int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"Ex-2", []int{2, 2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{"Ex-3", []int{-2, -1, -1, 1, 1, 2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}}},
		{"Ex-4", []int{-3, -2, -1, 0, 0, 1, 2, 3}, 0, [][]int{{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
	}

	t.Run("O(n^3) approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.FourSum(test.nums, test.target)
				if ok, _ := utils.CheckUnorderedSliceEquality(test.output, output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
