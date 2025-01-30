package dp1_test

import (
	"testing"

	dp1 "github.com/Binit-Dhakal/LeetCode-Go/DP1"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		// {"Ex-1", "babad", "bab"},
		{"Ex-2", "cbbd", "bb"},
		{"Ex-3", "a", "a"},
		{"Ex-4", "abcd1dcba", "abcd1dcba"},
		{"Ex-5 even", "hannah", "hannah"},
		{"Ex-6 even two digit", "bb", "bb"},
		{"Ex-7", "aaaa", "aaaa"},
	}

	t.Run("Brute force approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.LongestPalindromicSubstringBF(test.input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Two pointers approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.LongestPalindromicSubstringTwoPointers(test.input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("DP approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.LongestPalindromicSubstringDP(test.input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
