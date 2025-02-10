package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/Striver/LinkedList"
)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	testCases := []struct {
		name   string
		head   []int
		n      int
		output []int
	}{
		{"Ex-1", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"Ex-2", []int{1, 2, 3, 4, 5, 6}, 1, []int{1, 2, 3, 4, 5}},
		{"Ex-3", []int{1, 5, 6}, 2, []int{1, 6}},
	}

	t.Run("Naive Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				head := linkedlist.CreateLinkedList(test.head)
				output := linkedlist.RemoveNthNodeFromEndOfList(head, test.n)
				o := linkedlist.TraverseLinkedList(output)
				if !reflect.DeepEqual(o, test.output) {
					t.Errorf("Got: %v, want: %v", o, test.output)
				}
			})
		}
	})

	t.Run("Optimized One-Pass Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				head := linkedlist.CreateLinkedList(test.head)
				output := linkedlist.RemoveNthNodeFromEndOfListOptimized(head, test.n)
				o := linkedlist.TraverseLinkedList(output)
				if !reflect.DeepEqual(o, test.output) {
					t.Errorf("Got: %v, want: %v", o, test.output)
				}
			})
		}
	})
}
