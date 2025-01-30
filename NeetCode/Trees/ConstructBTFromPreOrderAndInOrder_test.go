package trees_test

import (
	"reflect"
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestConstructBTFromPreOrderAndInOrder(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name     string
		preorder []int
		inorder  []int
		output   []int
	}{
		{"Ex-1", []int{1, 2, 3, 4}, []int{2, 1, 3, 4}, []int{1, 2, 3, none, none, none, 4}},
		{"Ex-2", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []int{3, 9, 20, none, none, 15, 7}},
		{"Ex-3", []int{-1}, []int{-1}, []int{-1}},
	}

	t.Run("Using hashmap", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := trees.ConstructTreeFromPreOrderAndInOrder(test.preorder, test.inorder)
				got := trees.TraverseBinaryTree(output)
				if !reflect.DeepEqual(got, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
