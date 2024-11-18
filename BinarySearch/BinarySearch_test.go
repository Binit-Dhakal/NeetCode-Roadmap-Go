package binarysearch_test

import (
	"testing"

	binarysearch "github.com/Binit-Dhakal/LeetCode-Go/BinarySearch"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestBinarySeach(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		target int
		output int
	}{
		{"Target exists", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"Target doesnot exist", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	t.Run("Iterative Binary Search", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, binarysearch.IterativeSearch, test.input, test.target)
		}
	})

	t.Run("Recursive Binary Search", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, binarysearch.RecursiveSearch, test.input, test.target)
		}
	})

	t.Run("Built-in Golang search", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, binarysearch.BuiltInSearch, test.input, test.target)
		}
	})
}
