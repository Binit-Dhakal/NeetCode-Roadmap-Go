package twopointers_test

import (
	"testing"

	twopointers "github.com/Binit-Dhakal/LeetCode-Go/TwoPointers"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestThreeSumSort(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output [][]int
	}{
		{"with negative numbers", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"with only three numbers", []int{0, 1, 1}, [][]int{}},
		{"with all zero", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	t.Run("By sorting and two pointers", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, twopointers.ThreeSumSort, test.input)
		}
	})
}
