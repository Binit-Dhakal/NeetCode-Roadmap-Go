package trees_test

import (
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestSameTree(t *testing.T) {
	testCases := []struct {
		name   string
		p      []int
		q      []int
		output bool
	}{
		{"Same tree", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"Different tree-1", []int{1, 2}, []int{1, trees.NilNode, 2}, false},
		{"Invert Tree", []int{1, 2, 1}, []int{1, 1, 2}, false},
	}

	t.Run("Test Using BFS", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input1 := trees.BuildBinaryTree(test.p)
				input2 := trees.BuildBinaryTree(test.q)
				output := trees.SameTreeBFS(input1, input2)
				if test.output != output {
					t.Errorf("Got: %v, Want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Test Using DFS", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input1 := trees.BuildBinaryTree(test.p)
				input2 := trees.BuildBinaryTree(test.q)
				output := trees.SameTreeDFS(input1, input2)
				if test.output != output {
					t.Errorf("Got: %v, Want: %v", output, test.output)
				}
			})
		}
	})
}
