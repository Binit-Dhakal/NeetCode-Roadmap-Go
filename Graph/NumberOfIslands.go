package graph

func NumIslandsDFS(grid [][]byte) int {
	numIsland := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'

		dfs(r-1, c)
		dfs(r, c-1)
		dfs(r+1, c)
		dfs(r, c+1)
	}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '1' {
				dfs(row, col)
				numIsland++
			}
		}
	}
	return numIsland
}

func NumIslandsBFS(grid [][]byte) int {
	numIsland := 0

	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	var bfs func(r, c int)

	bfs = func(r, c int) {
		q := [][]int{{r, c}}
		grid[r][c] = '0'

		for len(q) > 0 {
			front := q[0]
			q = q[1:]

			row, col := front[0], front[1]
			for _, dir := range directions {
				dx, dy := row+dir[0], col+dir[1]
				if dx < 0 || dx >= len(grid) || dy < 0 || dy >= len(grid[row]) || grid[dx][dy] == '0' {
					continue
				}
				grid[dx][dy] = '0'
				q = append(q, []int{dx, dy})
			}
		}
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '1' {
				numIsland++
				bfs(r, c)
			}
		}
	}

	return numIsland
}
