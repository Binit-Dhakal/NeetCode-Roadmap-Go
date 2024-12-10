package backtracking_test

import (
	"testing"

	backtracking "github.com/Binit-Dhakal/LeetCode-Go/Backtracking"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestLetterCombinationPhoneNo(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output []string
	}{
		{"Two digit input", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		// 	{"Empty digit", "", []string{}},
		// 	{"One digit", "2", []string{"a", "b", "c"}},
		// 	{"Three digit input", "234", []string{
		// 		"adg", "adh", "adi",
		// 		"aeg", "aeh", "aei",
		// 		"afg", "afh", "afi",
		// 		"bdg", "bdh", "bdi",
		// 		"beg", "beh", "bei",
		// 		"bfg", "bfh", "bfi",
		// 		"cdg", "cdh", "cdi",
		// 		"ceg", "ceh", "cei",
		// 		"cfg", "cfh", "cfi",
		// 	}},
	}

	t.Run("dfs solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.LetterCombinations(test.input)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
