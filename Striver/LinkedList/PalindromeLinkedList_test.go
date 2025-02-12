package linkedlist_test

import (
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/Striver/LinkedList"
)

func TestPalindromeLinkedList(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output bool
	}{
		{"Ex-1", []int{1, 2, 2, 1}, true},
		{"Ex-2", []int{1, 2}, false},
		{"Ex-3", []int{1, 2, 3, 2, 1}, true},
		{"Ex-4", []int{1, 2, 3, 4, 5, 6}, false},
	}

	t.Run("O(1) space approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				head := linkedlist.CreateLinkedList(test.input)
				output := linkedlist.PalindromeLinkedList(head)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Recursion solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				head := linkedlist.CreateLinkedList(test.input)
				output := linkedlist.PalindromeLinkedListRecursion(head)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
