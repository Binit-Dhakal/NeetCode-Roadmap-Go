package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/LinkedList"
)

func TestRemoveNthNodeFromList(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		n      int
		output []int
	}{
		{"Normal", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"only one item", []int{1}, 1, []int{}},
		{"removing from end", []int{1, 2}, 1, []int{1}},
		{"removing from start", []int{1, 2}, 2, []int{2}},
		{"n equal to length", []int{1, 2, 3}, 3, []int{2, 3}},
	}

	t.Run("using brute force approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := linkedlist.BuildLinkedList(test.input)
				input = linkedlist.RemoveNthNodeFromEndBF(input, test.n)
				got := linkedlist.TraverseLinkedList(input)
				if !reflect.DeepEqual(got, test.output) {
					t.Errorf("got: %v, want: %v", got, test.output)
				}
			})
		}
	})

	t.Run("using brute force optimized approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := linkedlist.BuildLinkedList(test.input)
				input = linkedlist.RemoveNthNodeFromEndOptimized(input, test.n)
				got := linkedlist.TraverseLinkedList(input)
				if !reflect.DeepEqual(got, test.output) {
					t.Errorf("got: %v, want: %v", got, test.output)
				}
			})
		}
	})
	t.Run("using two pointers approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := linkedlist.BuildLinkedList(test.input)
				input = linkedlist.RemoveNthNodeFromEndTwoPointers(input, test.n)
				got := linkedlist.TraverseLinkedList(input)
				if !reflect.DeepEqual(got, test.output) {
					t.Errorf("got: %v, want: %v", got, test.output)
				}
			})
		}
	})
	t.Run("using recursion", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := linkedlist.BuildLinkedList(test.input)
				input = linkedlist.RemoveNthNodeFromEndRecursion(input, test.n)
				got := linkedlist.TraverseLinkedList(input)
				if !reflect.DeepEqual(got, test.output) {
					t.Errorf("got: %v, want: %v", got, test.output)
				}
			})
		}
	})
}
