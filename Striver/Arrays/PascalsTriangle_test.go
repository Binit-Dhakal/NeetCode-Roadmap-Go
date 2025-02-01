package arrays_test

import (
	"reflect"
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestPascalsTriangle(t *testing.T) {
	testCases := []struct {
		name    string
		numRows int
		output  [][]int
	}{
		{"Ex-1", 5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{"Ex-2", 1, [][]int{{1}}},
	}

	t.Run("Brute approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.PascalsTriangle(test.numRows)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Math-Combination Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.PascalsTriangleMath(test.numRows)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
