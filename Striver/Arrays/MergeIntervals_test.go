package arrays_test

import (
	"reflect"
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestMergeIntervals(t *testing.T) {
	testCases := []struct {
		name      string
		intervals [][]int
		output    [][]int
	}{
		{"Ex-1 Sorted intervals", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"Ex-2", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{"Ex-3 Unsorted intervals", [][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{"Ex-4", [][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
	}

	t.Run("Naive approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.MergeIntervals(test.intervals)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
