package linkedlist_test

import (
	"fmt"
	"reflect"
	"testing"

	linkedlist "github.com/Binit-Dhakal/LeetCode-Go/Striver/LinkedList"
)

func TestReverseLinkedList(t *testing.T) {
	testCases := []struct {
		name   string
		input  *linkedlist.ListNode
		output []int
	}{
		{"Ex-1", linkedlist.CreateLinkedList([]int{1, 2, 3, 4, 5}), []int{5, 4, 3, 2, 1}},
		{"Ex-2", linkedlist.CreateLinkedList([]int{4, 3, 2, 1, 8, 9}), []int{9, 8, 1, 2, 3, 4}},
	}

	t.Run("StraightForward Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				fmt.Println(linkedlist.TraverseLinkedList(test.input))
				out := linkedlist.ReverseList(test.input)
				output := linkedlist.TraverseLinkedList(out)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
