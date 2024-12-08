package backtracking

func WordSearch(board [][]byte, word string) bool {
	// for more optimal solution some notes
	// - reverse the string if the frequency of word[0] > word[-1] so as we have to call dfs less
	// - decrease frequency from word map with board map, and if any of the word map has count > 0, then we return false
	// - before calling recursion, check all the condition first, this will not have neat code as we have, but will save code from calling condition if it is going to return false anyway like r < 0 or c <0 or board[r][c] == "*";
	var dfs func(ptr int, r int, c int) bool

	dfs = func(ptr int, r int, c int) bool {
		if ptr == len(word) {
			return true
		}

		if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || board[r][c] != word[ptr] {
			return false
		}

		board[r][c] = '*'

		defer func() { board[r][c] = word[ptr] }()

		return dfs(ptr+1, r+1, c) || dfs(ptr+1, r-1, c) || dfs(ptr+1, r, c+1) || dfs(ptr+1, r, c-1)
	}

	for r := range len(board) {
		for c := range len(board[0]) {
			if board[r][c] == word[0] {
				if dfs(0, r, c) {
					return true
				}
			}
		}
	}

	return false
}
