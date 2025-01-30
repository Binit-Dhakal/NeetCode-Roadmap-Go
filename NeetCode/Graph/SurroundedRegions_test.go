package graph_test

import (
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestSurroundedRegions(t *testing.T) {
	testCases := []struct {
		name   string
		board  [][]byte
		output [][]byte
	}{
		{"Ex-1", [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}, [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		}},
		{"Ex-2", [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}, [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}},
		{"Ex-3", [][]byte{
			{'O', 'X', 'O', 'O', 'O', 'X'},
			{'O', 'O', 'X', 'X', 'X', 'O'},
			{'X', 'X', 'X', 'X', 'X', 'O'},
			{'O', 'O', 'O', 'O', 'X', 'X'},
			{'X', 'X', 'O', 'O', 'X', 'O'},
			{'O', 'O', 'X', 'X', 'X', 'X'},
		}, [][]byte{
			{'O', 'X', 'O', 'O', 'O', 'X'},
			{'O', 'O', 'X', 'X', 'X', 'O'},
			{'X', 'X', 'X', 'X', 'X', 'O'},
			{'O', 'O', 'O', 'O', 'X', 'X'},
			{'X', 'X', 'O', 'O', 'X', 'O'},
			{'O', 'O', 'X', 'X', 'X', 'X'},
		}},
		{"Ex-4", [][]byte{
			{'O', 'X', 'O'},
			{'X', 'O', 'X'},
			{'O', 'X', 'O'},
		}, [][]byte{
			{'O', 'X', 'O'},
			{'X', 'X', 'X'},
			{'O', 'X', 'O'},
		}},
	}

	t.Run("DFS approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				graph.SurroundedRegionsDFS(test.board)
				if ok, _ := utils.CheckUnorderedSliceEquality(test.board, test.output); !ok {
					t.Errorf("Got: %v, want: %v", test.board, test.output)
				}
			})
		}
	})
}
