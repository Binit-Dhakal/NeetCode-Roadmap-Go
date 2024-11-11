package arrayandhashing_test

import (
	"testing"

	arrayandhashing "github.com/Binit-Dhakal/LeetCode-Go/ArrayAndHashing"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestEncodeString(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
	}{
		{"Encode string", []string{"Neet", "code", "loves", "you"}},
		{"Encode string with more than 10 word length", []string{"Neeeeeeetttt", "coooooooode", "loves", "you"}},
	}

	for _, test := range testCases {
		utils.HelperTest(t, test.name, test.input, arrayandhashing.Decode, arrayandhashing.Encode(test.input))
	}
}
