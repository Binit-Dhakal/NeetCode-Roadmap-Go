package linkedlist_test

import (
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/Striver/LinkedList"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output [][]int
	}{
		{"Ex-1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"Ex-2", []int{0, 1, 1}, [][]int{}},
		{"Ex-3", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	t.Run("Hashset two sum approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				nums := utils.Copy1DSlice(test.nums)
				output := linkedlist.ThreeSum(nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(test.output, output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Hashset approach without sort", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				nums := utils.Copy1DSlice(test.nums)
				output := linkedlist.ThreeSumWithoutSort(nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(test.output, output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
