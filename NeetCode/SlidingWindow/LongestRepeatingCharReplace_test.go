package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/Binit-Dhakal/LeetCode-Go/SlidingWindow"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestLongestRepeatingCharReplacement(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		replacement int
		output      int
	}{
		{"Ex-1", "ABCDEFABA", 3, 5},
		{"Ex-2", "ABAB", 2, 4},
		{"Ex-3", "AABABBA", 1, 4},
		{"Ex-4", "ABBBABBA", 3, 8},
		{"Ex-5", "ABBBABBA", 2, 7},
		{"Ex-6; k=0", "AAAA", 0, 4},
		{"Ex-7; k=0", "AABA", 0, 2},
		{"Ex-8", "EOEMQLLQTRQDDCOERARHGAAARRBKCCMFTDAQOLOKARBIJBISTGNKBQGKKTALSQNFSABASNOPBMMGDIOETPTDICRBOMBAAHINTFLH", 7, 11},
		{"Ex-9", "IMNJJTRMJEGMSOLSCCQICIHLQIOGBJAEHQOCRAJQMBIBATGLJDTBNCPIFRDLRIJHRABBJGQAOLIKRLHDRIGERENNMJSDSSMESSTR", 2, 6},
	}

	t.Run("Two pointers", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, slidingwindow.LongestRepeatingCharReplacementTwoPointers, test.input, test.replacement)
		}
	})

	t.Run("Brute Force", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, slidingwindow.LongestRepeatingCharReplacementBruteForce, test.input, test.replacement)
		}
	})
}
