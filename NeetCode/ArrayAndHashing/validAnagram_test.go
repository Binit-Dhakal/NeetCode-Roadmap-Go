package arrayandhashing_test

import (
	"testing"

	arrayandhashing "github.com/Binit-Dhakal/LeetCode-Go/ArrayAndHashing"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestAnagram(t *testing.T) {
	tests := []struct {
		input1 string
		input2 string
		output bool
		name   string
	}{
		{"anagram", "nagaram", true, "base true case"},
		{"rat", "cat", false, "base false case"},
	}

	t.Run("Test Hashmap solution", func(t *testing.T) {
		for _, test := range tests {
			utils.HelperTest(t, test.name, test.output, arrayandhashing.AnagramHash, test.input1, test.input2)
		}
	})

	t.Run("Test Array solution", func(t *testing.T) {
		for _, test := range tests {
			utils.HelperTest(t, test.name, test.output, arrayandhashing.AnagramArray, test.input1, test.input2)
		}
	})
}
