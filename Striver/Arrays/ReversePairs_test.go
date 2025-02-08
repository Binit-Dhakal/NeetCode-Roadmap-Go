package arrays_test

import (
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestReversePairs(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output int
	}{
		{"Ex-1", []int{1, 3, 2, 3, 1}, 2},
		{"Ex-2", []int{2, 4, 3, 5, 1}, 3},
		{"Ex-3", []int{-5, -5}, 1},
	}

	t.Run("Merge Sort", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.ReversePairs(test.nums)
				if output != test.output {
					t.Errorf("Got %v, want: %v", output, test.output)
				}
			})
		}
	})
}
