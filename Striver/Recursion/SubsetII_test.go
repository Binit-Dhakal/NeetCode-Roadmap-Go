package recursion_test

import (
	"testing"

	recursion "github.com/Binit-Dhakal/LeetCode-Go/Striver/Recursion"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestSubsetII(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		result [][]int
	}{
		{"Ex-1", []int{1, 2, 2}, [][]int{{}, {1}, {2}, {1, 2}, {1, 2, 2}, {2, 2}}},
		{"Ex-2", []int{0}, [][]int{{}, {0}}},
		{"Ex-3", []int{1, 2, 1}, [][]int{{}, {1}, {1, 2}, {1, 1}, {1, 2, 1}, {2}}},
	}

	t.Run("Cascading", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := recursion.SubsetsIICascading(test.nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.result); !ok {
					t.Errorf("Got: %v, want: %v", output, test.result)
				}
			})
		}
	})

	t.Run("Recursion", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := recursion.SubsetsII(test.nums)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.result); !ok {
					t.Errorf("Got: %v, want: %v", output, test.result)
				}
			})
		}
	})
}
