package graph

func ValidTreeGraph(n int, edges [][]int) bool {
	adjList := make([][]int, n)
	for i := range n {
		adjList[i] = make([]int, 0)
	}

	for _, edge := range edges {
		// as this is undirected graph
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visitedNode := make([]bool, n)
	var dfs func(node int, parent int) bool
	dfs = func(node int, parent int) bool {
		if visitedNode[node] == true {
			return false
		}

		visitedNode[node] = true
		for _, child := range adjList[node] {
			if child != parent && !dfs(child, node) {
				// child is not parent
				return false
			}
		}

		return true
	}

	if !dfs(0, -1) {
		return false
	}

	for _, node := range visitedNode {
		if !node {
			return false
		}
	}

	return true
}

func ValidTreeGraphBFS(n int, edges [][]int) bool {
	adjList := make([][]int, n)
	for i := range n {
		adjList[i] = make([]int, 0)
	}

	for _, edge := range edges {
		// as this is undirected graph
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	queue := [][2]int{{0, -1}} // (node, parent)

	for len(queue) > 0 {
		node, parentNode := queue[0][0], queue[0][1]
		queue = queue[1:]
		visited[node] = true

		for _, childNode := range adjList[node] {
			if childNode == parentNode {
				continue
			}
			if visited[childNode] {
				// cyclic case
				return false
			}
			queue = append(queue, [2]int{childNode, node})

		}
	}

	for _, node := range visited {
		if !node {
			return false
		}
	}
	return true
}

type DSUValid struct {
	parent []int
	size   []int
}

func NewDSUValid(n int) *DSUValid {
	dsu := DSUValid{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := range dsu.parent {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return &dsu
}

func (dsu *DSUValid) find(node int) int {
	// path compression
	if dsu.parent[node] != node {
		dsu.parent[node] = dsu.find(dsu.parent[node])
	}
	return dsu.parent[node]
}

func (dsu *DSUValid) union(u int, v int) bool {
	du := dsu.find(u)
	dv := dsu.find(v)

	if du == dv {
		// cycle exists
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

func ValidTreeGraphDSU(n int, edges [][]int) bool {
	dsu := NewDSUValid(n)

	for _, edge := range edges {
		if !dsu.union(edge[0], edge[1]) {
			return false
		}
	}

	for _, par := range dsu.parent {
		if par != dsu.parent[0] {
			return false
		}
	}
	return true
}
