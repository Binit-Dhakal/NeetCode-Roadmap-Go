package arrays_test

import (
	"reflect"
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestFindTheDuplicateNumber(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{"Ex-1", []int{1, 3, 4, 2, 2}, 2},
		{"Ex-2", []int{3, 1, 3, 4, 2}, 3},
		{"Ex-3", []int{3, 3, 3, 3, 3}, 3},
	}

	t.Run("Test Linked List Approach(modify array)", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := utils.Copy1DSlice(test.input)
				output := arrays.DuplicateNumber(input)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Binary Search", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := utils.Copy1DSlice(test.input)
				output := arrays.DuplicateNumberBinarySearch(input)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Cycle Detection", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := utils.Copy1DSlice(test.input)
				output := arrays.DuplicateNumberCycleDetection(input)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
