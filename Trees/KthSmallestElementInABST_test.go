package trees_test

import (
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestKthSmallestElementInABST(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name   string
		input  []int
		k      int
		output int
	}{
		{"Ex-1", []int{3, 1, 4, none, 2}, 1, 1},
		{"Ex-2", []int{5, 3, 6, 2, 4, none, none, 1}, 3, 3},
	}

	t.Run("With DFS optimal", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.input)
				ele := trees.KthSmallestElementBSTDFS(input, test.k)
				if ele != test.output {
					t.Errorf("Got: %v, want: %v", ele, test.output)
				}
			})
		}
	})

	t.Run("With Iterative optimal", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.input)
				ele := trees.KthSmallestElementBSTIterative(input, test.k)
				if ele != test.output {
					t.Errorf("Got: %v, want: %v", ele, test.output)
				}
			})
		}
	})
	t.Run("With Morris Inorder Traversal", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.input)
				ele := trees.KthSmallestElementBSTMorrisTraversal(input, test.k)
				if ele != test.output {
					t.Errorf("Got: %v, want: %v", ele, test.output)
				}
			})
		}
	})
}
