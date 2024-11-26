package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/LinkedList"
)

func TestMergeKSortedLists(t *testing.T) {
	testCases := []struct {
		name   string
		input  [][]int
		output []int
	}{
		{"Normal case", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"Empty Linked list inside list", [][]int{{}}, []int{}},
		{"One item each case", [][]int{{1}, {0}}, []int{0, 1}},
	}

	t.Run("Using Normal merge", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := []*linkedlist.ListNode{}
				for _, inp := range test.input {
					input = append(input, linkedlist.BuildLinkedList(inp))
				}
				output := linkedlist.MergeKSortedLists(input)
				got := linkedlist.TraverseLinkedList(output)
				if !reflect.DeepEqual(test.output, got) {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})

	t.Run("Using Heap", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := []*linkedlist.ListNode{}
				for _, inp := range test.input {
					input = append(input, linkedlist.BuildLinkedList(inp))
				}
				output := linkedlist.MergeKSortedListsHeap(input)
				got := linkedlist.TraverseLinkedList(output)
				if !reflect.DeepEqual(test.output, got) {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})
	t.Run("Using Merge Sort", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := []*linkedlist.ListNode{}
				for _, inp := range test.input {
					input = append(input, linkedlist.BuildLinkedList(inp))
				}
				output := linkedlist.MergeKSortedListsDAC(input)
				got := linkedlist.TraverseLinkedList(output)
				if !reflect.DeepEqual(test.output, got) {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})
}
