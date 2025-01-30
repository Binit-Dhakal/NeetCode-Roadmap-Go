package trees_test

import (
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{"Normal case through root", []int{1, 2, 3, 4, 5}, 3},
		{"Normal case without root", []int{1, none, 2, none, none, 3, 4, none, none, none, none, 5}, 3},
		{"with only three element", []int{1, 2, 3}, 2},
		{"with only two node", []int{1, 2}, 1},
	}

	t.Run("With recursion", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.input)
				output := trees.DiameterOfBinaryTree(input)
				if output != test.output {
					t.Errorf("got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
