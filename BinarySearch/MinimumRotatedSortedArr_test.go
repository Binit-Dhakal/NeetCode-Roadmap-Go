package binarysearch_test

import (
	"testing"

	binarysearch "github.com/Binit-Dhakal/LeetCode-Go/BinarySearch"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestMinRotatedSortedArr(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{"Rotated 3 times", []int{3, 4, 5, 1, 2}, 1},
		{"Rotated 4 times", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"Rotated till original array", []int{11, 13, 15, 17}, 11},
	}

	t.Run("Binary Search method", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, binarysearch.MinRotatedSortedArrBS, test.input)
		}
	})
	t.Run("Binary Search with no variable", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, binarysearch.MinRotatedSortedArrBSRight, test.input)
		}
	})
}
