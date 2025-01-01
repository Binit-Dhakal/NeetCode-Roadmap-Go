package dp1_test

import (
	"testing"

	dp1 "github.com/Binit-Dhakal/LeetCode-Go/DP1"
)

func TestWordBreak(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		wordDict []string
		output   bool
	}{
		{"Ex-1", "leetcode", []string{"leet", "code"}, true},
		{"Ex-2", "applepenapple", []string{"apple", "pen"}, true},
		{"Ex-3", "catsanddog", []string{"cats", "dogs", "sand", "and", "cat"}, false},
		{"Ex-4 not solvable by greedy approach", "cars", []string{"car", "rs", "ca"}, true},
		{"Ex-5", "bb", []string{"a", "b", "bbb", "bbbb"}, true},
		{"Ex-6", "cars", []string{"car", "ca", "s"}, true},
	}

	t.Run("Brute force approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.WordBreakNaive(test.s, test.wordDict)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Trie approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.WordBreakTrie(test.s, test.wordDict)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("DP Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.WordBreakTopDown(test.s, test.wordDict)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("DP approach with hashset", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := dp1.WordBreakTopDownHashSet(test.s, test.wordDict)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
