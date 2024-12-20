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

type DSU struct {
	Parent []int
	Size   []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		Parent: make([]int, n+1),
		Size:   make([]int, n+1),
	}

	for i := 0; i <= n; i++ {
		dsu.Parent[i] = i
		dsu.Size[i] = 1
	}
	return dsu
}

func (dsu *DSU) find(node int) int {
	if dsu.Parent[node] != node {
		// path compression
		dsu.Parent[node] = dsu.find(dsu.Parent[node])
	}
	return dsu.Parent[node]
}

func (dsu *DSU) union(u int, v int) bool {
	du := dsu.find(u)
	dv := dsu.find(v)

	if du == dv {
		return false
	}

	if dsu.Size[du] >= dsu.Size[dv] {
		dsu.Size[du] += dsu.Size[dv]
		dsu.Parent[dv] = du
	} else {
		dsu.Size[dv] += dsu.Size[du]
		dsu.Parent[du] = dv
	}
	return true
}

func (dsu *DSU) getSize(node int) int {
	par := dsu.find(node)
	return dsu.Size[par]
}

func NumIslandsDisjointSet(grid [][]byte) int {
	row, col := len(grid), len(grid[0])
	dsu := NewDSU(row * col)
	directions := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	islands := 0

	index := func(r, c int) int {
		return r*col + c
	}

	for r := range row {
		for c := range col {
			if grid[r][c] == '1' {
				islands++
				for _, dir := range directions {
					dr, dc := r+dir[0], c+dir[1]
					if dr < 0 || dc < 0 || dr >= row || dc >= col || grid[dr][dc] == '0' {
						continue
					}

					if dsu.union(index(r, c), index(dr, dc)) {
						islands--
					}
				}
			}
		}
	}
	return islands
}
