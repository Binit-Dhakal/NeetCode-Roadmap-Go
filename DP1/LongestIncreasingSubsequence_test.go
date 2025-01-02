package dp1_test

import (
	"testing"

	dp1 "github.com/Binit-Dhakal/LeetCode-Go/DP1"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output int
	}{
		{"Ex-1", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"Ex-2", []int{0, 1, 0, 3, 2, 3}, 4},
		{"Ex-3", []int{7, 7, 7, 7, 7, 7, 7}, 1},
		{"Ex-4", []int{9, 1, 4, 2, 3, 3, 7}, 4},
		{"Ex-5", []int{0, 3, 1, 3, 2, 3}, 4},
		{"Ex-6 strictly increasing sequence", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{"Ex-7", []int{1, 7, 8, 4, 5, 6, -1, 9}, 5},
	}

	t.Run("Top-Bottom Appraoch", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.LongestIncreasingSubsequenceTopDown(test.nums)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Bottom-up Appraoch", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.LongestIncreasingSubsequenceBottomUp(test.nums)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Binary Search Appraoch", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.LongestIncreasingSubsequenceBinarySearch(test.nums)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
