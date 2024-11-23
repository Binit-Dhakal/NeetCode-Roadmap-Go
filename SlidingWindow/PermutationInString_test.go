package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/Binit-Dhakal/LeetCode-Go/SlidingWindow"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestPermutationInString(t *testing.T) {
	testCases := []struct {
		name   string
		input1 string
		input2 string
		output bool
	}{
		{"with s1 length 2 with true", "ab", "eidbaooo", true},
		{"with s1 length 2 with false", "ab", "eidboaooo", false},
		{"with s1 length 3 with true", "abc", "lecabee", true},
		{"with s1 length 3 with false", "abc", "lecaabee", false},
	}

	t.Run("permutation with sliding window", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, slidingwindow.PermutationInString, test.input1, test.input2)
		}
	})
}
