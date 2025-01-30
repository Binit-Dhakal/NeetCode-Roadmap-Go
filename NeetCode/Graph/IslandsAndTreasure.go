package graph

func IslandsAndTreasureDFS(grids [][]int) {
	// O((m*n)^2)
	directions := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	var dfs func(r, c, dist int)
	dfs = func(r, c, dist int) {
		if r < 0 || r >= len(grids) || c < 0 || c >= len(grids[0]) || grids[r][c] == -1 {
			return
		}

		if grids[r][c] < dist {
			return
		}

		grids[r][c] = dist
		dist++

		for _, direction := range directions {
			dr, dc := r+direction[0], c+direction[1]
			dfs(dr, dc, dist)
		}
	}
	for r := range grids {
		for c := range grids[0] {
			if grids[r][c] == 0 {
				// treasure place
				dfs(r, c, 0)
			}
		}
	}
}

func IslandsAndTreasureMultiBFS(grids [][]int) {
	queue := [][2]int{}
	for row := range grids {
		for col := range grids[0] {
			if grids[row][col] == 0 {
				queue = append(queue, [2]int{row, col})
			}
		}
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		r, c := point[0], point[1]

		for _, d := range directions {
			dr, dc := r+d[0], c+d[1]
			if dr < 0 || dr >= len(grids) || dc < 0 || dc >= len(grids[0]) || grids[dr][dc] <= grids[r][c] {
				continue
			}
			grids[dr][dc] = grids[r][c] + 1
			queue = append(queue, [2]int{dr, dc})
		}
	}
}
