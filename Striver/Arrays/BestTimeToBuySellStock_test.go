package arrays_test

import (
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestBestTimeToBuySell(t *testing.T) {
	testCases := []struct {
		name   string
		prices []int
		output int
	}{
		{"Ex-1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"Ex-2", []int{7, 6, 4, 3, 1}, 0},
	}

	t.Run("O(n) approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.BestTime(test.prices)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
