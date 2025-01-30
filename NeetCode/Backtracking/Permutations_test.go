package backtracking_test

import (
	"testing"

	backtracking "github.com/Binit-Dhakal/LeetCode-Go/Backtracking"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestPermutations(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output [][]int
	}{
		{"Ex-1", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{"Ex-2", []int{7}, [][]int{{7}}},
		{"Ex-3", []int{0, 1}, [][]int{{0, 1}, {1, 0}}},
	}

	t.Run("Recursion solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.PermutationsRecursion(test.nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, Want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Iterative Solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.PermutationsIteration(test.nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, Want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Backtracking solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.PermutationsBackTracking(test.nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, Want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Backtracking Optimal space solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.PermutationsBackOptimal(test.nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, Want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Backtracking Bitmask", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.PermutationsBackBitOp(test.nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, Want: %v", output, test.output)
				}
			})
		}
	})
}
