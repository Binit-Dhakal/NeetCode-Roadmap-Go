package graph

func SurroundedRegionsDFS(board [][]byte) {
	row, col := len(board), len(board[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= row || c < 0 || c >= col || board[r][c] != 'O' {
			return
		}

		board[r][c] = 'T'
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r, c-1)
		dfs(r-1, c)
	}

	for c := range col {
		if board[0][c] == 'O' {
			dfs(0, c)
		}
		if board[row-1][c] == 'O' {
			dfs(row-1, c)
		}
	}
	for r := range row {
		if board[r][0] == 'O' {
			dfs(r, 0)
		}
		if board[r][col-1] == 'O' {
			dfs(r, col-1)
		}
	}

	for r := range row {
		for c := range col {
			if board[r][c] == 'T' {
				board[r][c] = 'O'
			} else {
				board[r][c] = 'X'
			}
		}
	}
}
