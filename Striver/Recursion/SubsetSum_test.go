package recursion_test

import (
	"testing"

	recursion "github.com/Binit-Dhakal/LeetCode-Go/Striver/Recursion"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestSubsetSum(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		sums []int
	}{
		{"Ex-1", []int{2, 3}, []int{0, 2, 3, 5}},
		{"Ex-2", []int{1, 2, 1}, []int{0, 1, 1, 2, 2, 3, 3, 4}},
		{"Ex-3", []int{5, 6, 7}, []int{0, 5, 6, 7, 11, 12, 13, 18}},
	}

	t.Run("Iteration", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := recursion.SubsetSum(test.nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.sums); !ok {
					t.Errorf("Got: %v, want: %v", output, test.sums)
				}
			})
		}
	})

	t.Run("Recursion", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := recursion.SubsetSumRecursion(test.nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.sums); !ok {
					t.Errorf("Got: %v, want: %v", output, test.sums)
				}
			})
		}
	})
}
