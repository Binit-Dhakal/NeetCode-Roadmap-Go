package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/Striver/LinkedList"
)

func TestReverseNodesInKGroup(t *testing.T) {
	testCases := []struct {
		name   string
		head   []int
		k      int
		output []int
	}{
		{"Ex-1", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{"Ex-2", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
	}

	t.Run("Reverse nodes ", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				head := linkedlist.CreateLinkedList(test.head)
				output := linkedlist.ReverseNodesInKGroup(head, test.k)
				outputList := linkedlist.TraverseLinkedList(output)
				if !reflect.DeepEqual(outputList, test.output) {
					t.Errorf("Got: %v, want: %v", outputList, test.output)
				}
			})
		}
	})

	t.Run("Optimized-Temporary", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				head := linkedlist.CreateLinkedList(test.head)
				output := linkedlist.ReverseNodesInKGroupRecursion(head, test.k)
				outputList := linkedlist.TraverseLinkedList(output)
				if !reflect.DeepEqual(outputList, test.output) {
					t.Errorf("Got: %v, want: %v", outputList, test.output)
				}
			})
		}
	})
}
