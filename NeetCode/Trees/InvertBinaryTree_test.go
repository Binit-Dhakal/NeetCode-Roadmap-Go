package trees_test

import (
	"reflect"
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestInvertBinaryTree(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name   string
		input  []int
		output []int
	}{
		{"tree with one depth", []int{2, 1, 3}, []int{2, 3, 1}},
		{"normal tree", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 3, 2, 7, 6, 5, 4}},
		{"with nil", []int{3, 9, 20, none, 7, 15}, []int{3, 20, 9, none, 15, 7}},
	}

	t.Run("Using Recursion", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.input)

				output := trees.InvertBinaryTree(input)
				got := trees.TraverseBinaryTree(output)
				if !reflect.DeepEqual(got, test.output) {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})

	t.Run("Using BFS", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.input)
				output := trees.InvertBinaryTreeBFS(input)
				got := trees.TraverseBinaryTree(output)
				if !reflect.DeepEqual(got, test.output) {
					t.Errorf("Got: %v, want: %v", got, test.output)
				}
			})
		}
	})
}
