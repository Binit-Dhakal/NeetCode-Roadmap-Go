package graph_test

import (
	"reflect"
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
)

func TestRedundantConnection(t *testing.T) {
	testCases := []struct {
		name   string
		edges  [][]int
		output []int
	}{
		{"Ex-1", [][]int{{1, 2}, {1, 3}, {2, 3}}, []int{2, 3}},
		{"Ex-2", [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, []int{1, 4}},
		{"Ex-3", [][]int{{1, 2}, {2, 5}, {3, 5}, {3, 4}, {1, 3}}, []int{1, 3}},
	}

	t.Run("DSU solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.RedundantConnectionDSU(test.edges)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Kahn Algorithm solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.RedundantConnectionKahnAlgo(test.edges)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
