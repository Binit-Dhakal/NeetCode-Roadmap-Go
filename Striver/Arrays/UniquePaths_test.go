package arrays_test

import (
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		name   string
		m      int
		n      int
		output int
	}{
		{"Ex-1", 3, 7, 28},
		{"Ex-2", 3, 2, 3},
	}

	t.Run("DP-Memoization", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.UniquePathsTopDown(test.m, test.n)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("DP-BottomUp", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.UniquePathsBottomUp(test.m, test.n)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("DP-BottomUp-SpaceOptimized", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.UniquePathsBottomUpOptimized(test.m, test.n)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Combination", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.UniquePathsCombinatorics(test.m, test.n)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
