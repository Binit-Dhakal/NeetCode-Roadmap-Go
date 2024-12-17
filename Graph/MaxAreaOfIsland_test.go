package graph_test

import (
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
)

func TestMaxAreaOfIsland(t *testing.T) {
	testCases := []struct {
		name  string
		input [][]int
		ouput int
	}{
		{"Ex-1", [][]int{
			{0, 1, 1, 0, 1},
			{1, 0, 1, 0, 1},
			{0, 1, 1, 0, 1},
			{0, 1, 0, 0, 1},
		}, 6},
		{"Ex-2", [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		}, 6},
		{"Ex-3", [][]int{
			{0, 1, 0, 1, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 1, 0},
		}, 6},
	}

	t.Run("DFS method", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				inp := copy2DSlice(test.input)
				output := graph.MaxAreaOfIslandDFS(inp)
				if output != test.ouput {
					t.Errorf("Got: %v, want: %v", output, test.ouput)
				}
			})
		}
	})
	t.Run("BFS method", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				inp := copy2DSlice(test.input)
				output := graph.MaxAreaOfIslandBFS(inp)
				if output != test.ouput {
					t.Errorf("Got: %v, want: %v", output, test.ouput)
				}
			})
		}
	})
}
