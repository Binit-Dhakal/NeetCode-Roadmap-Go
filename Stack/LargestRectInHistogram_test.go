package stack_test

import (
	"testing"

	stack "github.com/Binit-Dhakal/LeetCode-Go/Stack"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestLargestRectInHistogram(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{"Example 1 with normal case", []int{7, 1, 7, 2, 2, 4}, 8},
		{"Example 2 with height in array as biggest rect", []int{1, 3, 7}, 7},
		{"Example 3 with normal case", []int{2, 1, 5, 6, 2, 3}, 10},
	}

	t.Run("Largest Rectangle using Brute Force", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, stack.LargestRectInHistogramBruteForce, test.input)
		}
	})

	t.Run("Largest Rectangle using Stack", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, stack.LargestRectInHistogramStack, test.input)
		}
	})
}
