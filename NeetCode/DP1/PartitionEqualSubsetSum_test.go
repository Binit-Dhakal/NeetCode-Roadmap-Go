package dp1_test

import (
	"testing"

	dp1 "github.com/Binit-Dhakal/LeetCode-Go/DP1"
)

func TestPartitionEqualSubset(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output bool
	}{
		{"Ex-1", []int{1, 5, 11, 5}, true},
		{"Ex-2", []int{1, 2, 3, 5}, false},
		{"Ex-3", []int{1, 2, 3, 4}, true},
		{"Ex-4", []int{1, 2, 3, 4, 5}, false},
		{"Ex-5", []int{5, 3, 2, 1}, false},
		{"Ex-6", []int{11, 1, 5, 5}, true},
	}

	t.Run("Brute force approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.PartitionEqualSubsetSumBF(test.nums)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
