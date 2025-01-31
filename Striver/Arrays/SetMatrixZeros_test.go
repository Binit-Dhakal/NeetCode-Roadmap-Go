package arrays_test

import (
	"reflect"
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestSetMatrixZeros(t *testing.T) {
	testCases := []struct {
		name   string
		matrix [][]int
		output [][]int
	}{
		{"Ex-1", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{"Ex-2", [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{"Ex-3", [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}, [][]int{{0, 0, 3, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}},
		{"Ex-4", [][]int{{1, 0, 3}}, [][]int{{0, 0, 0}}},
	}

	t.Run("Using hashmap", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				matrix := utils.Copy2DSlice(test.matrix)
				arrays.SimpleMatrixZerosSimple(matrix)
				if !reflect.DeepEqual(matrix, test.output) {
					t.Errorf("Got: %v, want: %v", matrix, test.output)
				}
			})
		}
	})

	t.Run("Using Space Optimized", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				matrix := utils.Copy2DSlice(test.matrix)
				arrays.SimpleMatrixZerosSpaceOptimized(matrix)
				if !reflect.DeepEqual(matrix, test.output) {
					t.Errorf("Got: %v, want: %v", matrix, test.output)
				}
			})
		}
	})
}
