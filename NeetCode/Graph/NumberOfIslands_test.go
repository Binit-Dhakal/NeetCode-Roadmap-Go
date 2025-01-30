package graph_test

import (
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
)

func TestNumberOfIslands(t *testing.T) {
	testCases := []struct {
		name   string
		input  [][]byte
		output int
	}{
		{"Ex-1", [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, 1},

		{"Ex-2", [][]byte{{'1', '1', '0', '0', '1'}, {'1', '1', '0', '0', '1'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}, {'0', '0', '0', '0', '1'}}, 4},
		{"Ex-3", [][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}, 1},
		{"Ex-4", [][]byte{{'1', '0', '1', '1', '1'}, {'1', '0', '1', '0', '1'}, {'1', '1', '1', '0', '1'}}, 1},
	}

	t.Run("DFS approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				temp := copy2DSlice(test.input)
				output := graph.NumIslandsDFS(temp)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("BFS approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				temp := copy2DSlice(test.input)
				output := graph.NumIslandsBFS(temp)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Using Union Find Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				temp := copy2DSlice(test.input)
				output := graph.NumIslandsDisjointSet(temp)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}

func copy2DSlice[T any](src [][]T) [][]T {
	// Create a new slice with the same length as the source
	dst := make([][]T, len(src))

	for i := range src {
		// Create a new slice for each inner slice and copy the data
		dst[i] = make([]T, len(src[i]))
		copy(dst[i], src[i])
	}

	return dst
}
