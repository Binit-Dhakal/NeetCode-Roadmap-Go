package graph_test

import (
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
)

func TestValidGraphTree(t *testing.T) {
	testCases := []struct {
		name   string
		n      int
		edges  [][]int
		output bool
	}{
		{"Ex-1", 5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}, true},
		{"Ex-2 - cyclic undirected edges", 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}, false},
		{"Ex-3 - some node not connected", 6, [][]int{{0, 1}, {1, 2}, {2, 3}, {5, 4}}, false},
	}

	t.Run("DFS approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.ValidTreeGraph(test.n, test.edges)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("BFS approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.ValidTreeGraphBFS(test.n, test.edges)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("DSU approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.ValidTreeGraphDSU(test.n, test.edges)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
