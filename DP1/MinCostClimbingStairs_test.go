package dp1_test

import (
	"testing"

	dp1 "github.com/Binit-Dhakal/LeetCode-Go/DP1"
)

func TestMinCostClimbingStairs(t *testing.T) {
	testCases := []struct {
		name   string
		cost   []int
		output int
	}{
		{"Ex-1", []int{10, 15, 20}, 15},
		{"Ex-2", []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{"Ex-3", []int{1, 2}, 1},
	}
	t.Run("Test Top Down approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.MinCostClimbingStairsTopDown(test.cost)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Test Bottom Up approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.MinCostClimbingStairsBottomUp(test.cost)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
