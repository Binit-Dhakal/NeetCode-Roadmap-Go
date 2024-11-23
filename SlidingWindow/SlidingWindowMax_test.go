package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/Binit-Dhakal/LeetCode-Go/SlidingWindow"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestSlidingWindowMax(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		k      int
		output []int
	}{
		{"Ex-1", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"Ex-2", []int{1}, 1, []int{1}},
		{"Ex-3", []int{1, 2, 1, 0, 4, 2, 6}, 3, []int{2, 2, 4, 4, 6}},
		{"Ex-4", []int{10, 9, 7, 6, 5, 4, 3, 2, 1}, 2, []int{10, 9, 7, 6, 5, 4, 3, 2}},
	}

	t.Run("Brute force approach", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, slidingwindow.SlidingWindowMaxBF, test.input, test.k)
		}
	})

	t.Run("Deque approach", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, slidingwindow.SlidingWindowMaxDeque, test.input, test.k)
		}
	})

	t.Run("Heap approach", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, slidingwindow.SlidingWindowMaxPQ, test.input, test.k)
		}
	})
}
