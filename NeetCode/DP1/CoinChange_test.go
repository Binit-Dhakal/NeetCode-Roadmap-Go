package dp1_test

import (
	"testing"

	dp1 "github.com/Binit-Dhakal/LeetCode-Go/DP1"
)

func TestCoinChange(t *testing.T) {
	testCases := []struct {
		name   string
		coins  []int
		amount int
		output int
	}{
		{"Ex-1", []int{1, 2, 5}, 7, 2},
		{"Ex-2", []int{2}, 3, -1},

		{"Ex-3", []int{1}, 0, 0},
		{"Ex-4", []int{1}, 1, 1},
		{"Ex-5", []int{1, 3, 5}, 8, 2},
		{"Ex-6", []int{1, 3, 5}, 19, 5},
		{"Ex-7", []int{2, 5, 10, 1}, 27, 4},
	}

	t.Run("DFS approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.CoinChangeDFS(test.coins, test.amount)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("DFS Memoization approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.CoinChangeDFSMemoization(test.coins, test.amount)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("DFS Optimal time", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.CoinChangeOptimal(test.coins, test.amount)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
