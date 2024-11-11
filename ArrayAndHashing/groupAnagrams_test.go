package arrayandhashing_test

import (
	"testing"

	arrayandhashing "github.com/Binit-Dhakal/LeetCode-Go/ArrayAndHashing"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name   string
		input  []string
		output [][]string
	}{
		{
			"normal case", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			"empty string", []string{""}, [][]string{{""}},
		},
		{
			"one element", []string{"a"}, [][]string{{"a"}},
		},
	}

	t.Run("Test Sorting version", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, arrayandhashing.GroupAnagramsSort, test.input)
		}
	})

	t.Run("Test Array version", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, arrayandhashing.GroupAnagramsArray, test.input)
		}
	})
}
