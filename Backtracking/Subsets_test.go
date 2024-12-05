package backtracking_test

import (
	"testing"

	backtracking "github.com/Binit-Dhakal/LeetCode-Go/Backtracking"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestSubsets(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output [][]int
	}{
		{"Ex-1", []int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{"Ex-2", []int{0}, [][]int{{}, {0}}},
		{"Ex-3", []int{1, 2, 3, 4}, [][]int{{}, {1}, {2}, {3}, {4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}, {1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}, {1, 2, 3, 4}}},
		{"Ex-4", []int{9, 0, 3, 5, 7}, [][]int{{}, {9}, {0}, {3}, {5}, {7}, {9, 0}, {9, 3}, {9, 5}, {9, 7}, {0, 3}, {0, 5}, {0, 7}, {3, 5}, {3, 7}, {5, 7}, {9, 0, 3}, {9, 0, 5}, {9, 0, 7}, {9, 3, 5}, {9, 3, 7}, {9, 5, 7}, {0, 3, 5}, {0, 3, 7}, {0, 5, 7}, {3, 5, 7}, {9, 0, 3, 5}, {9, 0, 3, 7}, {9, 0, 5, 7}, {9, 3, 5, 7}, {0, 3, 5, 7}, {9, 0, 3, 5, 7}}},
	}

	t.Run("Iterative solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.SubsetsIteration(test.input)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, Want: %v", output, test.output)
					t.Errorf("Got :%v, want: %v", len(output), len(test.output))
				}
			})
		}
	})

	t.Run("BackTracking solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.SubsetsBackTracking(test.input)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, Want: %v", output, test.output)
					t.Errorf("Got :%v, want: %v", len(output), len(test.output))
				}
			})
		}
	})

	t.Run("Bit mask solution", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.SubsetsBitMask(test.input)
				if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
					t.Errorf("Got: %v, Want: %v", output, test.output)
					t.Errorf("Got :%v, want: %v", len(output), len(test.output))
				}
			})
		}
	})
}
