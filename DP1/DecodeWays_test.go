package dp1_test

import (
	"testing"

	dp1 "github.com/Binit-Dhakal/LeetCode-Go/DP1"
)

func TestDecodeWays(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output int
	}{
		{"Ex-1", "12", 2},
		{"Ex-2", "226", 3},
		{"Ex-3", "06", 0},
		{"Ex-4", "11106", 2},
	}

	t.Run("DFS approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.DecodeWaysDFS(test.input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Bottom up approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.DecodeWaysBottomUp(test.input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Bottom up approach space optimized", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.DecodeWaysBottomUpOptimized(test.input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
