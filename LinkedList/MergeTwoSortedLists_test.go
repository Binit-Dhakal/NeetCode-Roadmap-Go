package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/LinkedList"
)

func TestMergeTwoSortedList(t *testing.T) {
	testCases := []struct {
		name   string
		list1  []int
		list2  []int
		output []int
	}{
		{"non repeating num", []int{1, 3, 5}, []int{2, 4, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"normal list", []int{2, 2, 4}, []int{1, 3, 4}, []int{1, 2, 2, 3, 4, 4}},
		{"empty input list", []int{}, []int{}, []int{}},
		{"one empty and other no", []int{}, []int{0}, []int{0}},
	}

	t.Run("merge linked list Iteration", func(t *testing.T) {
		for _, test := range testCases {
			list1 := linkedlist.BuildLinkedList(test.list1)
			list2 := linkedlist.BuildLinkedList(test.list2)
			got := linkedlist.TraverseLinkedList(linkedlist.MergeTwoSortedLists(list1, list2))
			if !reflect.DeepEqual(got, test.output) {
				t.Errorf("got: %v, want: %v", got, test.output)
			}
		}
	})

	t.Run("merge linked list Recursion", func(t *testing.T) {
		for _, test := range testCases {
			list1 := linkedlist.BuildLinkedList(test.list1)
			list2 := linkedlist.BuildLinkedList(test.list2)
			head := linkedlist.MergeTwoSortedListsRecursion(list1, list2)
			got := linkedlist.TraverseLinkedList(head)

			if !reflect.DeepEqual(got, test.output) {
				t.Errorf("got: %v, want: %v", got, test.output)
			}
		}
	})
}
