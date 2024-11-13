package twopointers_test

import (
	"testing"

	twopointers "github.com/Binit-Dhakal/LeetCode-Go/TwoPointers"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{"Normal case", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"only two heights give", []int{1, 1}, 1},
	}

	t.Run("Using Two pointers", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, twopointers.MaxArea, test.input)
		}
	})
}
