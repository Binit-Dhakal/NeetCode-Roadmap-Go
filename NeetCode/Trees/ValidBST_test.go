package trees_test

import (
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestValidBST(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name   string
		root   []int
		output bool
	}{
		{"Ex-1", []int{2, 1, 3}, true},
		{"Ex-2", []int{5, 1, 4, none, none, 3, 6}, false},
		{"Ex-3", []int{1, 2, 3}, false},
	}

	t.Run("DFS solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.root)
				output := trees.ValidateBST(input)
				if output != test.output {
					t.Errorf("Got: %v, Want: %v", output, test.output)
				}
			})
		}
	})
}
