package trees_test

import (
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestBalancedBinaryTree(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name   string
		input  []int
		output bool
	}{
		{"Balanced Tree", []int{3, 9, 20, none, none, 15, 7}, true},
		{"Not Balanced Tree", []int{1, 2, 2, 3, 3, none, none, 4, 4}, false},
	}

	t.Run("Recursive solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.input)
				output := trees.BalancedBinaryTreeDFS(input)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
