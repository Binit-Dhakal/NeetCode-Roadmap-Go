package twopointers_test

import (
	"testing"

	twopointers "github.com/Binit-Dhakal/LeetCode-Go/TwoPointers"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestTwoSumSorted(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		target int
		output []int
	}{
		{"Normal 4 length case", []int{2, 7, 11, 15}, 9, []int{1, 2}},
		{"3 length case", []int{2, 3, 4}, 6, []int{1, 3}},
		{"2 length case", []int{-1, 0}, -1, []int{1, 2}},
	}

	t.Run("Using two pointers", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, twopointers.TwoSumSortedMap, test.input, test.target)
		}
	})
}
