package trees_test

import (
	"math"
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestMaxDepthOfBinaryTree(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{"normal tree", []int{3, 9, 20, math.MaxInt, 7, 15, math.MaxInt}, 3},
		{"small tree", []int{1, math.MaxInt, 2}, 2},
	}

	t.Run("Recursive Solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.input)
				output := trees.MaxDepthOfBinaryTreeRecursion(input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("BFS Solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.input)
				output := trees.MaxDepthOfBinaryTreeBFS(input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Stack Iterative Solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.input)
				output := trees.MaxDepthOfBinaryTreeStack(input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
