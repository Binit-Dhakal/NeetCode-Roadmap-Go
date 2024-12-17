package graph_test

import (
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
)

func TestOrangeRotting(t *testing.T) {
	testCases := []struct {
		name          string
		grid          [][]int
		minuteElapsed int
	}{
		{"Ex-1", [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{"Ex-2", [][]int{{0, 2}}, 0},
		{"Ex-3", [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
	}

	t.Run("Multiple BFS solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.OrangesRotting(test.grid)
				if output != test.minuteElapsed {
					t.Errorf("Got: %v, want: %v", output, test.minuteElapsed)
				}
			})
		}
	})
}
