package graph_test

import (
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestIslandsAndTreasure(t *testing.T) {
	testCases := []struct {
		name   string
		input  [][]int
		output [][]int
	}{
		{"Ex-1", [][]int{
			{2147483647, -1, 0, 2147483647},
			{2147483647, 2147483647, 2147483647, -1},
			{2147483647, -1, 2147483647, -1},
			{0, -1, 2147483647, 2147483647},
		}, [][]int{
			{3, -1, 0, 1},
			{2, 2, 1, -1},
			{1, -1, 2, -1},
			{0, -1, 3, 4},
		}},
		{"Ex-2", [][]int{
			{0, -1},
			{2147483647, 2147483647},
		}, [][]int{{0, -1}, {1, 2}}},
	}

	t.Run("DFS solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := copy2DSlice(test.input)

				graph.IslandsAndTreasureDFS(input)
				if ok, _ := utils.CheckUnorderedSliceEquality(input, test.output); !ok {
					t.Errorf("Got: %v, want: %v", input, test.output)
				}
			})
		}
	})

	t.Run("Multi-source BFS solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := copy2DSlice(test.input)

				graph.IslandsAndTreasureMultiBFS(input)
				if ok, _ := utils.CheckUnorderedSliceEquality(input, test.output); !ok {
					t.Errorf("Got: %v, want: %v", input, test.output)
				}
			})
		}
	})
}
