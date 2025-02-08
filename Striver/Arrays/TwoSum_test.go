package arrays_test

import (
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		output []int
	}{
		{"Ex-1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Ex-2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Ex-3", []int{3, 3}, 6, []int{0, 1}},
	}

	t.Run("HashMap approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.TwoSumHashMap(test.nums, test.target)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
