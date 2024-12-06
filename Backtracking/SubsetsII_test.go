package backtracking_test

import (
	"testing"

	backtracking "github.com/Binit-Dhakal/LeetCode-Go/Backtracking"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestSubsetII(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output [][]int
	}{
		{"Ex-1", []int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		{"Ex-2", []int{0}, [][]int{{}, {0}}},
		{"Ex-3", []int{1, 2, 1}, [][]int{{}, {1}, {2}, {1, 2}, {1, 1}, {1, 2, 1}}},
	}

	t.Run("BackTracking solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.SubsetsII(test.nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Iterative solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.SubsetsIIIterative(test.nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
