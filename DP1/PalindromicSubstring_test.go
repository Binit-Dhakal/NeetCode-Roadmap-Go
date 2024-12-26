package dp1_test

import (
	"testing"

	dp1 "github.com/Binit-Dhakal/LeetCode-Go/DP1"
)

func TestPalindromicSubstring(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output int
	}{
		{"Ex-1", "abc", 3},
		{"Ex-2", "aaa", 6},
		{"Ex-3", "aaaa", 10},
	}

	t.Run("DP approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.PalindromicSubstringDP(test.input)
				if output != test.output {
					t.Errorf("Got %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Two pointers approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.PalindromicSubstringTwoPointers(test.input)
				if output != test.output {
					t.Errorf("Got %v, want: %v", output, test.output)
				}
			})
		}
	})
}
