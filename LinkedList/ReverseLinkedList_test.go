package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/LinkedList"
)

func TestReverseLinkedList(t *testing.T) {
	testCases := []struct {
		name  string
		input *linkedlist.ListNode
		ouput []int
	}{
		{"Normal case", linkedlist.BuildLinkedList([]int{1, 2, 3, 4, 5}), []int{5, 4, 3, 2, 1}},
		{"only two numbers", linkedlist.BuildLinkedList([]int{1, 2}), []int{2, 1}},
		{"no input", linkedlist.BuildLinkedList([]int{}), []int{}},
	}

	t.Run("Reverse Linked List", func(t *testing.T) {
		for _, test := range testCases {
			got := linkedlist.TraverseLinkedList(linkedlist.ReverseLinkedList(test.input))
			if !reflect.DeepEqual(test.ouput, got) {
				t.Errorf("Got %v, want %v", got, test.ouput)
			}
		}
	})
}
