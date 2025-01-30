package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/Binit-Dhakal/LeetCode-Go/SlidingWindow"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestLongestSubstringWithoutRepeating(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output int
	}{
		{"different character", "abcabcbb", 3},
		{"same character", "bbbbb", 1},
		{"single character", " ", 1},
		{"normal case", "aab", 2},
		{"empty string", "", 0},
		{"pointer correction", "dvdfv", 3},
		{"ex", "abba", 2},
	}

	t.Run("Test Longest Substring Two Pointers", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, slidingwindow.LongestSubstringWithoutRepeatingTwoPointers, test.input)
		}
	})
	t.Run("Test Longest Substring Optimized", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, slidingwindow.LongestSubstringWithoutRepeatingOptimized, test.input)
		}
	})
}
