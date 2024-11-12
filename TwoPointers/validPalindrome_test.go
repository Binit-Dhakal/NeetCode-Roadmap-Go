package twopointers_test

import (
	"testing"

	twopointers "github.com/Binit-Dhakal/LeetCode-Go/TwoPointers"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output bool
	}{
		{"with a lot of non-alphanumeric char", "A man, a plan, a canal: Panama", true},
		{"with only space as non-alphanumeric", "race a car", false},
		{"empty string", " ", true},
	}

	t.Run("Run normal function", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, twopointers.IsPalindromeNaive, test.input)
		}
	})

	t.Run("Function that removes non-alphanumeric at start", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, twopointers.IsPalindromeRemoveNonAlphanumeric, test.input)
		}
	})
}
