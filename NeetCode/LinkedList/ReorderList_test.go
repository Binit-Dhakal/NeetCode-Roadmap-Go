package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/LinkedList"
)

func TestReorderList(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output []int
	}{
		{"even number of element", []int{0, 1, 2, 3, 4, 5}, []int{0, 5, 1, 4, 2, 3}},
		{"odd number of element", []int{0, 1, 2, 3, 4, 5, 6}, []int{0, 6, 1, 5, 2, 4, 3}},
		{"only one element", []int{1}, []int{1}},
	}

	t.Run("using struct of pointers", func(t *testing.T) {
		for _, test := range testCases {
			input := linkedlist.BuildLinkedList(test.input)
			linkedlist.ReorderList(input)
			got := linkedlist.TraverseLinkedList(input)
			if !reflect.DeepEqual(got, test.output) {
				t.Errorf("got: %v, want: %v", got, test.output)
			}
		}
	})

	t.Run("using reverse on half of linked list", func(t *testing.T) {
		for _, test := range testCases {
			input := linkedlist.BuildLinkedList(test.input)
			linkedlist.ReorderListWithReverse(input)
			got := linkedlist.TraverseLinkedList(input)
			if !reflect.DeepEqual(got, test.output) {
				t.Errorf("got: %v, want: %v", got, test.output)
			}
		}
	})

	t.Run("using recursion", func(t *testing.T) {
		for _, test := range testCases {
			input := linkedlist.BuildLinkedList(test.input)
			linkedlist.ReorderListWithRecursion(input)
			got := linkedlist.TraverseLinkedList(input)
			if !reflect.DeepEqual(got, test.output) {
				t.Errorf("got: %v, want: %v", got, test.output)
			}
		}
	})
}
