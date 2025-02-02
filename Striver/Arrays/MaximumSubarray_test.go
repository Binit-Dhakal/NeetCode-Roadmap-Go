package arrays_test

import (
	"reflect"
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestMaxSubArray(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output int
	}{
		{"Ex-1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"Ex-2", []int{1}, 1},
		{"Ex-3", []int{5, 4, -1, 7, 8}, 23},
		{"Ex-4", []int{-10, -2, -3, -1}, -1},
	}

	t.Run("Brute-Force Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.MaximumSubArrayBrute(test.nums)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Kadenes Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.MaximumSubArrayKadenes(test.nums)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
