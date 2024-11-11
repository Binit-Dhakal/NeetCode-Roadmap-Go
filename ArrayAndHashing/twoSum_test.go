package arrayandhashing_test

import (
	"testing"

	arrayandhashing "github.com/Binit-Dhakal/LeetCode-Go/ArrayAndHashing"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		output []int
		name   string
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 0}, "length 4 input"},
		{[]int{3, 2, 4}, 6, []int{1, 2}, "length 3 input"},
		{[]int{3, 3}, 6, []int{0, 1}, "length 2 input"},
	}

	t.Run("Test Brute force", func(t *testing.T) {
		for _, test := range tests {
			utils.HelperTest(t, test.name, test.output, arrayandhashing.TwoSumBrute, test.input, test.target)
		}
	})

	t.Run("Test hash map", func(t *testing.T) {
		for _, test := range tests {
			utils.HelperTest(t, test.name, test.output, arrayandhashing.TwoSumBrute, test.input, test.target)
		}
	})
}
