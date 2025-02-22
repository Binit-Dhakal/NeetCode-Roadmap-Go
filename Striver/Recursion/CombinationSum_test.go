package recursion_test

import (
	"testing"

	recursion "github.com/Binit-Dhakal/LeetCode-Go/Striver/Recursion"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestCombinationSum(t *testing.T) {
	testCases := []struct {
		name       string
		candidates []int
		target     int
		output     [][]int
	}{
		{"Ex-1", []int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{"Ex-2", []int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},

		{"Ex-3", []int{7, 6, 2, 3}, 7, [][]int{{2, 2, 3}, {7}}},
		{"Ex-4", []int{2}, 1, [][]int{}},
	}

	t.Run("Recursion-0/1", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := recursion.CombinationSum(test.candidates, test.target)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("BackTracking", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := recursion.CombinationSumBackTracking(test.candidates, test.target)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
