package graph

func PacificAtlanticWaterFlowBackTracking(heights [][]int) [][]int {
	row, col := len(heights), len(heights[0])
	resIndex := make([][]int, 0)

	pacific := false
	atlantic := false
	directions := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	var dfs func(r, c, curHeight int)
	dfs = func(r, c, curHeight int) {
		if r < 0 || c < 0 {
			pacific = true
			return
		}

		if r >= row || c >= col {
			atlantic = true
			return
		}

		tempHeight := heights[r][c]
		if curHeight < tempHeight {
			return
		}

		heights[r][c] = int(^uint(0) >> 1)
		for _, d := range directions {
			dfs(r+d[0], c+d[1], tempHeight)
			if pacific && atlantic {
				break
			}
		}
		heights[r][c] = tempHeight
	}

	for r := range row {
		for c := range col {
			atlantic = false
			pacific = false
			dfs(r, c, heights[r][c])
			if atlantic && pacific {
				resIndex = append(resIndex, []int{r, c})
			}
		}
	}

	return resIndex
}

func PacificAtlanticWaterFlowSet(heights [][]int) [][]int {
	row, col := len(heights), len(heights[0])
	pacificSet := make(map[[2]int]struct{})
	atlanticSet := make(map[[2]int]struct{})
	directions := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

	var dfs func(r, c int, visitedMap map[[2]int]struct{}, prevHeight int)
	dfs = func(r, c int, visitedMap map[[2]int]struct{}, prevHeight int) {
		if r < 0 || c < 0 || c >= col || r >= row || prevHeight > heights[r][c] {
			return
		}
		if _, ok := visitedMap[[2]int{r, c}]; ok {
			return
		}

		visitedMap[[2]int{r, c}] = struct{}{}

		prevHeight = heights[r][c]
		for _, d := range directions {
			dfs(r+d[0], c+d[1], visitedMap, prevHeight)
		}
	}
	for c := range col {
		dfs(0, c, pacificSet, heights[0][c])
		dfs(row-1, c, atlanticSet, heights[row-1][c])
	}

	for r := range row {
		dfs(r, 0, pacificSet, heights[r][0])
		dfs(r, col-1, atlanticSet, heights[r][col-1])
	}

	res := make([][]int, 0)
	for r := range row {
		for c := range col {
			coord := [2]int{r, c}
			if _, ok := pacificSet[coord]; ok {
				if _, ok := atlanticSet[coord]; ok {
					res = append(res, []int{r, c})
				}
			}
		}
	}

	return res
}
