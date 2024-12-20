package graph_test

import (
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
)

func TestConnectedComponent(t *testing.T) {
	testCases := []struct {
		name   string
		n      int
		edges  [][]int
		output int
	}{
		{"Ex-1", 3, [][]int{{0, 1}, {0, 2}}, 1},
		{"Ex-2", 6, [][]int{{0, 1}, {1, 2}, {2, 3}, {4, 5}}, 2},
		{"Ex-3 cyclic", 6, [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 3}, {4, 5}}, 2},
	}
	t.Run("Using DSU", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.ConnectedComponents(test.n, test.edges)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Using DFS", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.ConnectedComponentsDFS(test.n, test.edges)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
