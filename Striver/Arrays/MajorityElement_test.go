package arrays_test

import (
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		name       string
		nums       []int
		maxElement int
	}{
		{"Ex-1", []int{3, 2, 3}, 3},
		{"Ex-2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	t.Run("Boyer-Moore Majority Vote Algorithm", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.MajorityElement(test.nums)
				if output != test.maxElement {
					t.Errorf("Got: %v, want: %v", output, test.maxElement)
				}
			})
		}
	})

	t.Run("Bitwise Solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.MajorityElementBitwise(test.nums)
				if output != test.maxElement {
					t.Errorf("Got: %v, want: %v", output, test.maxElement)
				}
			})
		}
	})
}
