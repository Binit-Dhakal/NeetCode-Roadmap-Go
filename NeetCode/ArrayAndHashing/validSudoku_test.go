package arrayandhashing_test

import (
	"testing"

	arrayandhashing "github.com/Binit-Dhakal/LeetCode-Go/ArrayAndHashing"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestValidSudoku(t *testing.T) {
	testCases := []struct {
		board  [][]byte
		result bool
		name   string
	}{
		{
			[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}, true, "Valid sudoku",
		},
		{
			[][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}, false, "Invalid sudoku - column number repeat",
		},
		{
			[][]byte{
				{'6', '2', '.', '.', '3', '.', '.', '.', '.'},
				{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
				{'.', '9', '1', '.', '.', '.', '.', '.', '3'},
				{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
				{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '.', '.', '.', '.', '.', '8', '.', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			}, false, "Invalid sudoku - number repeat in 3*3 subbox",
		},
	}

	t.Run("Naive loop solution", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.result, arrayandhashing.ValidSudoku, test.board)
		}
	})

	t.Run("Solution with only one loop", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.result, arrayandhashing.ValidSudokuBest, test.board)
		}
	})
}
