package graph

type DSUConnect struct {
	parent []int
	size   []int
}

func (dsu *DSUConnect) find(node int) int {
	if dsu.parent[node] != node {
		dsu.parent[node] = dsu.find(dsu.parent[node])
	}
	return dsu.parent[node]
}

func (dsu *DSUConnect) union(u, v int) {
	du := dsu.find(u)
	dv := dsu.find(v)
	if du == dv {
		return
	}

	if dsu.size[du] <= dsu.size[dv] {
		dsu.parent[dv] = du
		dsu.size[du] += dsu.size[dv]
		dsu.size[dv] = 0
	} else {
		dsu.parent[du] = dv
		dsu.size[dv] += dsu.size[du]
		dsu.size[du] = 0
	}

	return
}

func ConnectedComponents(n int, edges [][]int) int {
	dsu := &DSUConnect{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := range dsu.parent {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}

	for _, edge := range edges {
		dsu.union(edge[0], edge[1])
	}

	count := 0
	for _, size := range dsu.size {
		if size != 0 {
			count++
		}
	}
	return count
}

func ConnectedComponentsDFS(n int, edges [][]int) int {
	// create adjacency list
	adjList := make([][]int, n)
	visit := make([]bool, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	var dfs func(node int)
	dfs = func(node int) {
		visit[node] = true

		for _, child := range adjList[node] {
			if visit[child] == true {
				continue
			}

			dfs(child)
		}
	}

	res := 0
	for i := range n {
		if !visit[i] {
			dfs(i)
			res++
		}
	}
	return res
}
