package linkedlist_test

import (
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/LinkedList"
)

func TestDuplicateNumber(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{"Ex-1", []int{1, 3, 4, 2, 2}, 2},
		{"Ex-2", []int{3, 1, 3, 4, 2}, 3},
		{"Ex-3 all same", []int{3, 3, 3, 3, 3}, 3},
	}

	t.Run("Modifying array", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := make([]int, len(test.input))
				copy(input, test.input)
				got := linkedlist.FindDuplicateNumberNegativeMarking(input)
				if got != test.output {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})

	t.Run("Using Floyds Hare And Tortoise Algorithm", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				got := linkedlist.FindDuplicateNumberFloydsAlgo(test.input)
				if got != test.output {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})

	t.Run("Using Binary Search", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				got := linkedlist.FindDuplicateNumberBinarySearch(test.input)
				if got != test.output {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})

	t.Run("Using Bit Manipulation", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				got := linkedlist.FindDuplicateNumberXOR(test.input)
				if got != test.output {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})
}
