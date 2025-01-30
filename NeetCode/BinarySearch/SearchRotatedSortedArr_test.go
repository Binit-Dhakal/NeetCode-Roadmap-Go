package binarysearch_test

import (
	"testing"

	binarysearch "github.com/Binit-Dhakal/LeetCode-Go/BinarySearch"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestSearchRotatedSortedArr(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		target int
		output int
	}{
		{"number is present", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"number is not present", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"number in original sorted array", []int{0, 1, 2, 4, 5, 6, 7}, 6, 5},
		{"only one number in arr and target preset", []int{1}, 1, 0},
		{"only one number in arr and target not preset", []int{1}, 0, -1},
	}

	t.Run("Test Search in Rotated Sorted Arr", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, binarysearch.SearchRotatedSortedBS, test.input, test.target)
		}
	})
}
