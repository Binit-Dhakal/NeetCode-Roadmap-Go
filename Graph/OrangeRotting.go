package graph

func OrangesRotting(grid [][]int) int {
	// using multisource BFS
	row, col := len(grid), len(grid[0])
	queue := [][2]int{}
	for r := range row {
		for c := range col {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	totalMinute := 0

	for len(queue) > 0 {
		n := len(queue)
		for n > 0 {
			point := queue[0]
			queue = queue[1:]
			r, c := point[0], point[1]

			for _, d := range directions {
				dr, dc := r+d[0], c+d[1]
				if dr < 0 || dr >= row || dc < 0 || dc >= col || grid[dr][dc] == 0 || grid[dr][dc] == 2 {
					continue
				}

				grid[dr][dc] = 2
				queue = append(queue, [2]int{dr, dc})
			}
			n--
		}
		if len(queue) != 0 {
			totalMinute++
		}
	}

	for r := range row {
		for c := range col {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}

	return totalMinute
}
