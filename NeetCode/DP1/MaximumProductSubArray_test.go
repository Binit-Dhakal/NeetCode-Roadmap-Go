package dp1_test

import (
	"testing"

	dp1 "github.com/Binit-Dhakal/LeetCode-Go/DP1"
)

func TestMaxProductSubArray(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output int
	}{
		{"Ex-1", []int{2, 3, -2, 4}, 6},
		{"Ex-2", []int{-2, 0, -1}, 0},
		{"Ex-3", []int{2, 2, -3, 0, 4, -1}, 4},
		{"Ex-4", []int{2, 2, -3, -1, 0, 4}, 12},
		{"Ex-5", []int{2, 0, 3, 0, -1, 5}, 5},
		{"Ex-6", []int{2, 0, 3, -1, 0, 1, 2}, 3},
	}

	t.Run("My Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.MaxProductSubArrayKadenes(test.nums)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Prefix Suffix Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.MaxProductSubArrayPrefixSuffix(test.nums)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
