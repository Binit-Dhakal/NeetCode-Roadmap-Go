package arrayandhashing_test

import (
	"testing"

	arrayandhashing "github.com/Binit-Dhakal/LeetCode-Go/ArrayAndHashing"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output bool
	}{
		{"containing duplicate", []int{1, 2, 3, 1}, true},
		{"not containing duplicate", []int{1, 2, 3, 4}, false},
		{"containing multiple duplicate", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	t.Run("Test hashset approach", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, arrayandhashing.ContainsDuplicate, test.input)
		}
	})
}
