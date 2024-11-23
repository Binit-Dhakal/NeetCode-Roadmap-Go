package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/Binit-Dhakal/LeetCode-Go/SlidingWindow"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestMinimumWindowSubstring(t *testing.T) {
	testCases := []struct {
		name   string
		input1 string
		input2 string
		output string
	}{
		{"Ex-1", "OUZODYXAZV", "XYZ", "YXAZ"},
		{"Ex-2 same length", "xyz", "xyz", "xyz"},
		{"Ex-3 s is less than t", "x", "xy", ""},
		{"Ex-4", "ADOBECODEBANC", "ABC", "BANC"},
		{"Ex-5", "aa", "aa", "aa"},
	}

	t.Run("using brute force", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, slidingwindow.MinimumWindowSubstringBF, test.input1, test.input2)
		}
	})

	t.Run("using sliding window", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, slidingwindow.MinimumWindowSubstringSW, test.input1, test.input2)
		}
	})
}
