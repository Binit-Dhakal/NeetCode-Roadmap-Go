package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/LinkedList"
)

func TestReverseNodesInKGroup(t *testing.T) {
	testCases := []struct {
		name   string
		head   []int
		k      int
		output []int
	}{
		{"Normal case with k=2", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{"Normal case with k=3", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
	}

	t.Run("Test Direct Solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := linkedlist.ReverseNodesInKGroupRecursion(linkedlist.BuildLinkedList(test.head), test.k)
				got := linkedlist.TraverseLinkedList(output)
				if !reflect.DeepEqual(got, test.output) {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})
}
