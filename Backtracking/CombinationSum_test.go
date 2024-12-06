package backtracking_test

import (
	"testing"

	backtracking "github.com/Binit-Dhakal/LeetCode-Go/Backtracking"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestCombinationSum(t *testing.T) {
	testCases := []struct {
		name       string
		candidates []int
		target     int
		result     [][]int
	}{
		{"Ex-1", []int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{"Ex-2", []int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"Ex-3", []int{2}, 1, [][]int{}},
	}

	t.Run("Naive solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.CombinationSum(test.candidates, test.target)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.result); !ok {
					t.Errorf("Got: %v, want: %v", output, test.result)
				}
			})
		}
	})
}
