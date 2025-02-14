package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/Striver/LinkedList"
)

func TestRotateList(t *testing.T) {
	testCases := []struct {
		name   string
		head   []int
		k      int
		output []int
	}{
		{"Ex-1", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"Ex-2", []int{0, 1, 2}, 4, []int{2, 0, 1}},
		{"Ex-3", []int{1}, 10, []int{1}},
		{"Ex-4", []int{1}, 1, []int{1}},
	}

	t.Run("O(2n) approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := linkedlist.CreateLinkedList(test.head)
				output := linkedlist.RotateList(input, test.k)
				outputList := linkedlist.TraverseLinkedList(output)
				if !reflect.DeepEqual(outputList, test.output) {
					t.Errorf("Got: %v, want: %v", outputList, test.output)
				}
			})
		}
	})
}
