package arrayandhashing_test

import (
	"testing"

	arrayandhashing "github.com/Binit-Dhakal/LeetCode-Go/ArrayAndHashing"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestLongestSequence(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{"No number repeat case", []int{100, 4, 200, 1, 3, 2}, 4},
		{"Number repeat case", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}

	t.Run("Set solution", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, arrayandhashing.SetlongestConsecutive, test.input)
		}
	})
}
