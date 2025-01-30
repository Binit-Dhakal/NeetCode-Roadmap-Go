package dp1_test

import (
	"testing"

	dp1 "github.com/Binit-Dhakal/LeetCode-Go/DP1"
)

func TestHouseRobberII(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output int
	}{
		{"Ex-1", []int{2, 3, 2}, 3},
		{"Ex-2", []int{1, 2, 3, 1}, 4},
		{"Ex-3", []int{1, 2, 3}, 3},
		{"Ex-4", []int{1}, 1},
		{"Ex-5", []int{1, 2}, 2},
	}

	t.Run("DFS naive approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.HouseRobberII(test.nums)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Bottom up optimized approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.HouseRobberIIBottomUp(test.nums)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
