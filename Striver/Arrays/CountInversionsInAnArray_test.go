package arrays_test

import (
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestCountInversions(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output int
	}{
		{"Ex-1", []int{1, 2, 3, 4, 5}, 0},
		{"Ex-2", []int{5, 4, 3, 2, 1}, 10},
		{"Ex-3", []int{5, 3, 2, 1, 4}, 7},
	}

	t.Run("Merge Sort Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.CountInversions(test.nums)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
