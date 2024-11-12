package arrayandhashing

func ValidSudoku(board [][]byte) bool {
	// check row for unique value
	var countArr [10]bool
	for row := 0; row < len(board); row++ {
		for k := range countArr { // Resetting value of array
			countArr[k] = false
		}
		for col := 0; col < len(board); col++ {
			val := board[row][col]
			if val == '.' {
				continue
			}

			if countArr[val-48] == true {
				return false
			}
			countArr[val-48] = true
		}
	}

	//  check column for unique row
	for row := 0; row < len(board); row++ {
		for k := range countArr { // Resetting value of array
			countArr[k] = false
		}
		for col := 0; col < len(board); col++ {
			val := board[col][row]
			if val == '.' {
				continue
			}

			if countArr[val-48] == true {
				return false
			}
			countArr[val-48] = true
		}
	}

	for square := 0; square < 9; square++ {
		for k := range countArr { // Resetting value of array
			countArr[k] = false
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				row := (square/3)*3 + i
				col := (square%3)*3 + j

				val := board[row][col]
				if val == '.' {
					continue
				}

				if countArr[val-48] == true {
					return false
				}
				countArr[val-48] = true
			}
		}
	}

	return true
}

func ValidSudokuBest(board [][]byte) bool {
	// copied from https://leetcode.com/problems/valid-sudoku/solutions/5911077/simple-o-1-sudoku-validator-using-bitwise-masks-for-efficient-frequency-checking-beats-100/
	// liked this solution

	var row, col, box [9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				mask := 1 << num
				if (row[i]&mask) > 0 || (col[j]&mask) > 0 || (box[i/3*3+j/3]&mask) > 0 {
					return false
				}
				row[i] |= mask
				col[j] |= mask
				box[i/3*3+j/3] |= mask
			}
		}
	}

	return true
}
