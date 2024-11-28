package trees_test

import (
	"reflect"
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestBTLevelOrderTraversal(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name   string
		root   []int
		output [][]int
	}{
		{"balanced tree", []int{1, 2, 3, 4, 5, 6, 7}, [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
		{"unbalanced tree", []int{3, 9, 20, none, none, 15, 7}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"only one node", []int{1}, [][]int{{1}}},
		{"no node", []int{}, [][]int{}},
	}

	t.Run("BT Level Order Traversal - BFS", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.root)
				output := trees.BTLevelOrderTraversalBFS(input)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("BT Level Order Traversal - DFS", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.root)
				output := trees.BTLevelOrderTraversalDFS(input)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
