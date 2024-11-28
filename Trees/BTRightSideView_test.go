package trees_test

import (
	"reflect"
	"testing"

	trees "github.com/Binit-Dhakal/LeetCode-Go/Trees"
)

func TestBTRightSideView(t *testing.T) {
	none := trees.NilNode
	testCases := []struct {
		name   string
		root   []int
		output []int
	}{
		{"Ex-1", []int{1, 2, 3, none, 5, none, 4}, []int{1, 3, 4}},
		{"Ex-2", []int{1, 2, 3, 4, none, none, none, 5}, []int{1, 3, 4, 5}},
		{"Ex-3", []int{1, none, 3}, []int{1, 3}},
		{"Ex-4", []int{}, []int{}},
	}

	t.Run("Test Iterative", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.root)
				output := trees.BTRightSideViewBFS(input)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, Want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Test DFS", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := trees.BuildBinaryTree(test.root)
				output := trees.BTRightSideViewDFS(input)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, Want: %v", output, test.output)
				}
			})
		}
	})
}
