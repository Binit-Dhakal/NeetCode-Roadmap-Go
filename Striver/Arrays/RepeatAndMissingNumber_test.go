package arrays_test

import (
	"reflect"
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestRepeatAndMissingNum(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"Ex-1", []int{3, 1, 2, 5, 3}, []int{3, 4}},
		{"Ex-2", []int{1, 2, 3, 4, 5, 6, 7, 8, 8}, []int{8, 9}},
	}

	t.Run("Math approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.RepeatAndMissingNumber(test.nums)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("XOR approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.RepeatAndMissingNumberXOR(test.nums)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
