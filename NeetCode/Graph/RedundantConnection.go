package graph

type DSURed struct {
	parent []int
	size   []int
}

func (dsu *DSURed) find(node int) int {
	if dsu.parent[node] != node {
		dsu.parent[node] = dsu.find(dsu.parent[node])
	}
	return dsu.parent[node]
}

func (dsu *DSURed) union(u, v int) bool {
	du := dsu.find(u)
	dv := dsu.find(v)
	if du == dv {
		return false
	}

	if dsu.size[du] >= dsu.size[dv] {
		dsu.parent[dv] = du
		dsu.size[du] += dsu.size[dv]
	} else {
		dsu.parent[du] = dv
		dsu.size[dv] += dsu.size[du]
	}
	return true
}

func RedundantConnectionDSU(edges [][]int) []int {
	n := len(edges) + 1
	dsu := &DSURed{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := range n {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}

	res := []int{}
	for _, edge := range edges {
		if !dsu.union(edge[0], edge[1]) {
			res = edge
		}
	}
	return res
}

func RedundantConnectionKahnAlgo(edges [][]int) []int {
	n := len(edges) + 1
	indegree := make([]int, n)
	adjList := make([][]int, n)
	for i := range adjList {
		adjList[i] = make([]int, 0)
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
		indegree[u]++
		indegree[v]++
	}

	queue := make([]int, 0)
	for i, deg := range indegree {
		if deg == 1 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		indegree[node]--

		for _, child := range adjList[node] {
			indegree[child]--
			if indegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	for i := len(edges) - 1; i >= 0; i-- {
		u, v := edges[i][0], edges[i][1]
		if indegree[u] == 2 && indegree[v] == 2 {
			return []int{u, v}
		}
	}

	return []int{}
}
