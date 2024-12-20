package graph

func MaxAreaOfIslandDFS(grid [][]int) int {
	maxArea := 0

	area := 0
	var dfs func(row, col int)

	dfs = func(row, col int) {
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) || grid[row][col] == 0 {
			return
		}

		area += 1
		grid[row][col] = 0
		maxArea = max(area, maxArea)

		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				area = 0
				dfs(r, c)
			}
		}
	}
	return maxArea
}

func MaxAreaOfIslandBFS(grid [][]int) int {
	maxArea := 0
	directions := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 0 {
				continue
			}

			area := 0
			queue := [][2]int{}
			queue = append(queue, [2]int{r, c})
			for len(queue) > 0 {
				point := queue[0]
				queue = queue[1:]
				row, col := point[0], point[1]
				grid[row][col] = 0

				area++
				for _, d := range directions {
					dr, dc := row+d[0], col+d[1]
					if dr < 0 || dr >= len(grid) || dc < 0 || dc >= len(grid[0]) || grid[dr][dc] == 0 {
						continue
					}
					queue = append(queue, [2]int{row + d[0], col + d[1]})
				}
			}

			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}

func MaxAreaOfIslandDisjointSet(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dsu := NewDSU(rows * cols)
	directions := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	index := func(r, c int) int {
		return r*cols + c
	}
	area := 0

	for r := range rows {
		for c := range cols {
			if grid[r][c] == 1 {
				for _, dir := range directions {
					dr, dc := r+dir[0], c+dir[1]
					if dr < 0 || dc < 0 || dr >= rows || dc >= cols || grid[dr][dc] == 0 {
						continue
					}

					dsu.union(index(r, c), index(dr, dc))
				}

				area = max(area, dsu.getSize(index(r, c)))
			}
		}
	}

	return area
}
