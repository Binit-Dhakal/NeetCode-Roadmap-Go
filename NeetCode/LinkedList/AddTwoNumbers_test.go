package linkedlist_test

import (
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/LinkedList"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		name   string
		l1     []int
		l2     []int
		output []int
	}{
		{"equal length of l1 and l2", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"both have only 0 as node", []int{0}, []int{0}, []int{0}},
		{"unequal length of l1 and l2(bigger)", []int{9, 9, 9, 9}, []int{9, 9, 9, 9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
		{"unequal length of l1(bigger) and l2", []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	t.Run("Using Iteration", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				l1 := linkedlist.BuildLinkedList(test.l1)
				l2 := linkedlist.BuildLinkedList(test.l2)
				output := linkedlist.AddTwoNumbers(l1, l2)
				got := linkedlist.TraverseLinkedList(output)

				if !reflect.DeepEqual(got, test.output) {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})

	t.Run("Using Recursion", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				l1 := linkedlist.BuildLinkedList(test.l1)
				l2 := linkedlist.BuildLinkedList(test.l2)
				output := linkedlist.AddTwoNumbersRecursion(l1, l2)
				got := linkedlist.TraverseLinkedList(output)

				if !reflect.DeepEqual(got, test.output) {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})
}
