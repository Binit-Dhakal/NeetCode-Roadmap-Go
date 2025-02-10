package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/Striver/LinkedList"
)

func TestMergeTwoSortedLists(t *testing.T) {
	testCases := []struct {
		name   string
		list1  *linkedlist.ListNode
		list2  *linkedlist.ListNode
		output []int
	}{
		{"Ex-1", linkedlist.CreateLinkedList([]int{1, 2, 4}), linkedlist.CreateLinkedList([]int{1, 3, 4}), []int{1, 1, 2, 3, 4, 4}},
		{"Ex-2", linkedlist.CreateLinkedList([]int{}), linkedlist.CreateLinkedList([]int{}), []int{}},
		{"Ex-3", linkedlist.CreateLinkedList([]int{1, 2, 5, 12}), linkedlist.CreateLinkedList([]int{3, 4, 10}), []int{1, 2, 3, 4, 5, 10, 12}},
	}

	t.Run("Merge Sorted List", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				outputNode := linkedlist.MergeTwoSortedLists(test.list1, test.list2)
				output := linkedlist.TraverseLinkedList(outputNode)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
