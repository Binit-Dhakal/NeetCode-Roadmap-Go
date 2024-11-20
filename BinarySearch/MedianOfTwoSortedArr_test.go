package binarysearch_test

import (
	"testing"

	binarysearch "github.com/Binit-Dhakal/LeetCode-Go/BinarySearch"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestMedianOfTwoSortedArr(t *testing.T) {
	testCases := []struct {
		name   string
		input1 []int
		input2 []int
		output float64
	}{
		{"odd length of merged arr", []int{1, 3}, []int{2}, 2.00000},
		{"even length of merged arr", []int{1, 2}, []int{3, 4}, 2.50000},
		{"nums1 length be 0 and other be 1", []int{}, []int{1}, 1},
		{"nums length total be 2", []int{1}, []int{2}, 1.5},
	}

	t.Run("Brute Force Approach", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, binarysearch.MedianBruteForce, test.input1, test.input2)
		}
	})
}
