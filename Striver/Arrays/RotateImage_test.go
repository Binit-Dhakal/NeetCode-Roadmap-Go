package arrays_test

import (
	"reflect"
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestRotateImage(t *testing.T) {
	testCases := []struct {
		name   string
		matrix [][]int
		output [][]int
	}{
		{"Ex-1", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{"Ex-2", [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},
	}

	t.Run("Naive approach", func(t *testing.T) {
		for _, test := range testCases {
			arrays.RotateImage(test.matrix)
			if !reflect.DeepEqual(test.matrix, test.output) {
				t.Errorf("Got: %v, want: %v", test.matrix, test.output)
			}
		}
	})
}
