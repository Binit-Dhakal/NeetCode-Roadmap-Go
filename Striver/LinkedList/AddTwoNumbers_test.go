package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/Striver/LinkedList"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		name   string
		list1  []int
		list2  []int
		output []int
	}{
		{"Ex-1", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"Ex-2", []int{0}, []int{0}, []int{0}},
		{"Ex-3", []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	t.Run("BF approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				l1 := linkedlist.CreateLinkedList(test.list1)
				l2 := linkedlist.CreateLinkedList(test.list2)
				outll := linkedlist.AddTwoNumbers(l1, l2)
				out := linkedlist.TraverseLinkedList(outll)
				if !reflect.DeepEqual(out, test.output) {
					t.Errorf("Got: %v, want: %v", out, test.output)
				}
			})
		}
	})
}
