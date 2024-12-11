package backtracking_test

import (
	"testing"

	backtracking "github.com/Binit-Dhakal/LeetCode-Go/Backtracking"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestNQueens(t *testing.T) {
	testCases := []struct {
		name   string
		input  int
		output [][]string
	}{
		{"For n=4", 4, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
		{"For n=1", 1, [][]string{{"Q"}}},
		{"For n=3", 3, [][]string{}},
		{"For n=5", 5, [][]string{
			{"Q....", "..Q..", "....Q", ".Q...", "...Q."},
			{"Q....", "...Q.", ".Q...", "....Q", "..Q.."},
			{".Q...", "...Q.", "Q....", "..Q..", "....Q"},
			{".Q...", "....Q", "..Q..", "Q....", "...Q."},
			{"..Q..", "Q....", "...Q.", ".Q...", "....Q"},
			{"..Q..", "....Q", ".Q...", "...Q.", "Q...."},
			{"...Q.", "Q....", "..Q..", "....Q", ".Q..."},
			{"...Q.", ".Q...", "....Q", "..Q..", "Q...."},
			{"....Q", ".Q...", "...Q.", "Q....", "..Q.."},
			{"....Q", "..Q..", "Q....", "...Q.", ".Q..."},
		}},
	}

	t.Run("Brute force approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.SolveNQueens(test.input)
				if ok, _ := utils.CheckUnorderedSliceEquality(test.output, output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Using hashmap", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.SolveNQueensMap(test.input)
				if ok, _ := utils.CheckUnorderedSliceEquality(test.output, output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Bitwise", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := backtracking.SolveNQueensBitwise(test.input)
				if ok, _ := utils.CheckUnorderedSliceEquality(test.output, output); !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
