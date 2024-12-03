package trees_test

import (
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestBTMaxPathSum(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name   string
		root   []int
		output int
	}{
		{"Ex-1", []int{1, 2, 3}, 6},
		{"Ex-2", []int{-10, 9, 20, none, none, 15, 7}, 42},
		{"Ex-3", []int{-15, 10, 20, none, none, 15, 5, -5}, 40},
		{"Ex-4", []int{2, -1}, 2},
		{"Ex-5 negative sum", []int{-2, -1}, -1},
		{"Ex-6", []int{1, -2, 3}, 4},
		{"Ex-7 linear", []int{1, 2, none, 3, none, none, none, 4, none, none, none, none, none, none, none, 5}, 15},
	}

	t.Run("Binary Tree Maximum Path BF", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.root)
				output := trees.BTMaxPathSum(input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
