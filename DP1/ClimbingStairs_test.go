package dp1_test

import (
	"testing"

	dp1 "github.com/Binit-Dhakal/LeetCode-Go/DP1"
)

func TestClimbingStairs(t *testing.T) {
	testCases := []struct {
		name   string
		n      int
		output int
	}{
		{"Ex-1", 2, 2},
		{"Ex-2", 3, 3},
		{"Ex-3", 4, 5},
		{"Ex-5", 5, 8},
	}

	t.Run("Climbing stairs Bottom up", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.ClimbingStairsBottomUp(test.n)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Climbing stairs Top Down with memoization", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.ClimbingStairsTopDown(test.n)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Climbing stairs math", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.ClimbingStairsMath(test.n)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
