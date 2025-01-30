package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/Binit-Dhakal/LeetCode-Go/SlidingWindow"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestBestTimeForTransaction(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{"max profit do exist", []int{7, 1, 5, 3, 6, 4}, 5},
		{"max profit donot exist", []int{7, 6, 4, 3, 1}, 0},
	}

	t.Run("kinda two pointer solution", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, slidingwindow.BestTimeForTransaction, test.input)
		}
	})
}
