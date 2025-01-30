package graph_test

import (
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestPacificAtlanticWaterFlow(t *testing.T) {
	testCases := []struct {
		name   string
		input  [][]int
		output [][]int
	}{
		{
			"Ex-1",
			[][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			[][]int{
				{0, 4},
				{1, 3},
				{1, 4},
				{2, 2},
				{3, 0},
				{3, 1},
				{4, 0},
			},
		},
		{"Ex-2", [][]int{{1}}, [][]int{{0, 0}}},
		{"Ex-3", [][]int{{4, 2, 7, 3, 4}, {7, 4, 6, 4, 7}, {6, 3, 5, 3, 6}}, [][]int{{0, 2}, {0, 4}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 0}}},
		{"Ex-4", [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, [][]int{{0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
	}

	t.Run("Test Back tracking", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.PacificAtlanticWaterFlowBackTracking(test.input)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Test DFS", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.PacificAtlanticWaterFlowSet(test.input)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
