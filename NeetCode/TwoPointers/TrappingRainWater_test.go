package twopointers_test

import (
	"testing"

	twopointers "github.com/Binit-Dhakal/LeetCode-Go/TwoPointers"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestTrappingRainWater(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{"with multiple barrier", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"with just one barrier", []int{4, 2, 0, 3, 2, 5}, 9},
	}

	t.Run("Solution with maxLeft and maxRight arrays", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, twopointers.TrapPrefixSuffixArray, test.input)
		}
	})

	t.Run("Solution with two pointers", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, twopointers.TrapTwoPointers, test.input)
		}
	})
}
