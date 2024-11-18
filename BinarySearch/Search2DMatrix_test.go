package binarysearch_test

import (
	"testing"

	binarysearch "github.com/Binit-Dhakal/LeetCode-Go/BinarySearch"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestSearch2DMatrix(t *testing.T) {
	testCases := []struct {
		name   string
		input  [][]int
		target int
		output bool
	}{
		{"Target exists", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{"Target not exists", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
	}

	t.Run("Iterative Search in 2D Matrix", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, binarysearch.IterativeSearchMatrix, test.input, test.target)
		}
	})

	t.Run("Recursive Search in 2D Matrix", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, binarysearch.RecursiveSearchMatrix, test.input, test.target)
		}
	})
}
