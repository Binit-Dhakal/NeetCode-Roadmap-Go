package arrays_test

import (
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestLongestConsecutiveSequence(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output int
	}{
		{"Ex-1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"Ex-2", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}

	t.Run("Hashmap Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.LongestConsecutiveSequence(test.nums)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
