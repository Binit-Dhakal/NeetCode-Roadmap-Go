package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/Striver/LinkedList"
)

func TestMiddleOfLinkedList(t *testing.T) {
	testCases := []struct {
		name   string
		input  *linkedlist.ListNode
		output []int
	}{
		{"Ex-1", linkedlist.CreateLinkedList([]int{1, 2, 3, 4, 5}), []int{3, 4, 5}},
		{"Ex-2", linkedlist.CreateLinkedList([]int{1, 2, 3, 4, 5, 6}), []int{4, 5, 6}},
	}

	t.Run("Fast and Slow Pointer", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				midNode := linkedlist.MiddleOfLinkedListFastSlowPtr(test.input)
				mid := linkedlist.TraverseLinkedList(midNode)
				if !reflect.DeepEqual(test.output, mid) {
					t.Errorf("Got: %v, want: %v", mid, test.output)
				}
			})
		}
	})
}
