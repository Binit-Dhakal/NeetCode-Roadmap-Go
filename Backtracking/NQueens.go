package backtracking

import (
	"strings"
)

func SolveNQueens(n int) [][]string {
	res := make([][]string, 0)

	board := make([]string, n)

	for k := range n {
		board[k] = strings.Repeat("*", n)
	}
	var queenPlaced int

	var dfs func(row int)
	dfs = func(row int) {
		if row >= len(board) {
			if queenPlaced == n {
				temp := make([]string, n)
				copy(temp, board)
				res = append(res, temp)
			}
			return
		}

		for c := 0; c < n; c++ {
			if board[row][c] == '.' {
				continue
			}

			tempBoard := make([]string, n)
			copy(tempBoard, board)

			b := board[row]
			b = b[:c] + "Q" + b[c+1:]
			queenPlaced++
			b = strings.Replace(b, "*", ".", -1)
			board[row] = b

			// change all diagonal and vertical to "."
			sep := 1
			for r := row + 1; r < len(board); r++ {
				board[r] = board[r][:c] + "." + board[r][c+1:]

				if c-sep >= 0 {
					board[r] = board[r][:c-sep] + "." + board[r][c-sep+1:]
				}
				if c+sep < len(board[r]) {
					board[r] = board[r][:c+sep] + "." + board[r][c+sep+1:]
				}

				sep++
			}
			dfs(row + 1)
			queenPlaced--
			board = tempBoard
		}
	}

	dfs(0)
	return res
}

func SolveNQueensMap(n int) [][]string {
	res := make([][]string, 0)
	colsMap := make(map[int]bool)
	posDiagMap := make(map[int]bool)
	negDiagMap := make(map[int]bool)

	board := make([]string, n)

	for i := range n {
		board[i] = strings.Repeat(".", n)
	}
	var dfs func(row int)

	dfs = func(row int) {
		if row == n {
			temp := make([]string, n)
			copy(temp, board)
			res = append(res, temp)
			return
		}

		for c := 0; c < n; c++ {
			if colsMap[c] || posDiagMap[row+c] || negDiagMap[row-c] {
				continue
			}

			posDiagMap[row+c] = true
			negDiagMap[row-c] = true
			colsMap[c] = true

			board[row] = board[row][:c] + "Q" + board[row][c+1:]

			dfs(row + 1)

			posDiagMap[row+c] = false
			negDiagMap[row-c] = false
			colsMap[c] = false
			board[row] = board[row][:c] + "." + board[row][c+1:]
		}
	}

	dfs(0)
	return res
}

func SolveNQueensBitwise(n int) [][]string {
	res := make([][]string, 0)
	board := make([]string, n)

	for i := range n {
		board[i] = strings.Repeat(".", n)
	}
	var dfs func(row int, colBit, posDiagBit, negDiagBit int)
	dfs = func(row int, colBit int, posDiagBit int, negDiagBit int) {
		if row == n {
			temp := make([]string, n)
			copy(temp, board)
			res = append(res, temp)
			return
		}

		for c := 0; c < n; c++ {
			if colBit&(1<<c) != 0 || posDiagBit&(1<<(row+c)) != 0 || negDiagBit&(1<<(row-c+n)) != 0 {
				continue
			}
			colBit ^= (1 << c)
			posDiagBit ^= (1 << (row + c))
			negDiagBit ^= (1 << (row - c + n))

			board[row] = board[row][:c] + "Q" + board[row][c+1:]

			dfs(row+1, colBit, posDiagBit, negDiagBit)

			board[row] = board[row][:c] + "." + board[row][c+1:]
			colBit ^= (1 << c)
			posDiagBit ^= (1 << (row + c))
			negDiagBit ^= (1 << (row - c + n))
		}
	}

	dfs(0, 0, 0, 0)
	return res
}
