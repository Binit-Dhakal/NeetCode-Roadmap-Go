package trees_test

import (
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestCountGoodNodesInBinaryTree(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name   string
		root   []int
		output int
	}{
		{"Ex-1", []int{3, 1, 4, 3, none, 1, 5}, 4},
		{"Ex-2", []int{3, 3, none, 4, 2}, 3},
		{"Ex-3", []int{1}, 1},
		{"Ex-4", []int{5, 3, none, 2, 2, none, none, 4, none, none, none, none, none, none, none, 4}, 1},
	}

	t.Run("Count Good nodes", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.root)
				output := trees.CountGoodNodesInBT(input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Count Good nodes optimal solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.root)
				output := trees.CountGoodNodesInBTOptimal(input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
