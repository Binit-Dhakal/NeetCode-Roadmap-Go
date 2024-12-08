package backtracking_test

import (
	"testing"

	backtracking "github.com/Binit-Dhakal/LeetCode-Go/Backtracking"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestCombinationSumII(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		target int
		output [][]int
	}{
		{"Ex-1", []int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{"Ex-2", []int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
		{"Ex-3", []int{1, 2, 3, 4, 5}, 7, [][]int{{1, 2, 4}, {2, 5}, {3, 4}}},
	}

	t.Run("Backtracking approach", func(t *testing.T) {
		for _, test := range testCases {
			output := backtracking.CombinationSumII(test.input, test.target)
			if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
				t.Errorf("Got: %v, Want: %v", output, test.output)
			}
		}
	})
	t.Run("Backtracking Optimal approach", func(t *testing.T) {
		for _, test := range testCases {
			output := backtracking.CombinationSumIIOptimal(test.input, test.target)
			if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
				t.Errorf("Got: %v, Want: %v", output, test.output)
			}
		}
	})
}
