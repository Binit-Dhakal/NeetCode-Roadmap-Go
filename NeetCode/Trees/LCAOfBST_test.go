package trees_test

import (
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestLCAOfBST(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name   string
		root   []int
		p, q   int
		output int
	}{
		{"node is not descandant of itself", []int{6, 2, 8, 0, 4, 7, 9, none, none, 3, 5}, 2, 8, 6},
		{"node is descandant of itself", []int{6, 2, 8, 0, 4, 7, 9, none, none, 3, 5}, 2, 4, 2},
		{"simple tree case", []int{2, 1}, 2, 1, 2},
	}

	t.Run("Test LCA of BST Recursion", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input1 := trees.BuildBinaryTree(test.root)
				inputP := trees.BuildBinaryTree([]int{test.p})
				inputQ := trees.BuildBinaryTree([]int{test.q})
				output := trees.LCAOfBSTRecursion(input1, inputP, inputQ)
				if output.Val != test.output {
					t.Errorf("Got: %v, want: %v", output.Val, test.output)
				}
			})
		}
	})

	t.Run("Test LCA of BST Iteration", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input1 := trees.BuildBinaryTree(test.root)
				inputP := trees.BuildBinaryTree([]int{test.p})
				inputQ := trees.BuildBinaryTree([]int{test.q})
				output := trees.LCAOfBSTIteration(input1, inputP, inputQ)

				if output.Val != test.output {
					t.Errorf("Got: %v, want: %v", output.Val, test.output)
				}
			})
		}
	})
}
