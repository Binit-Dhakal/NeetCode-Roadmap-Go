package trees_test

import (
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestSubTreeOfAnotherTree(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name    string
		root    []int
		subroot []int
		output  bool
	}{
		{"is subroot", []int{3, 4, 5, 1, 2}, []int{4, 1, 2}, true},
		{"is not subroot", []int{3, 4, 5, 1, 2, none, none, none, none, 0}, []int{4, 1, 2}, false},
		{"single subroot", []int{1, 1}, []int{1}, true},
	}

	t.Run("Test Subtree", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input1 := trees.BuildBinaryTree(test.root)
				input2 := trees.BuildBinaryTree(test.subroot)
				output := trees.SubTreeOfAnotherTreeDFS(input1, input2)

				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Test Serialize", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input1 := trees.BuildBinaryTree(test.root)
				input2 := trees.BuildBinaryTree(test.subroot)
				output := trees.SubTreeOfAnotherTreeSerialize(input1, input2)

				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
