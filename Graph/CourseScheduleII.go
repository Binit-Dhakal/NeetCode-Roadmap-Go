package graph

func CourseScheduleIIReverseNodes(numCourses int, prerequisites [][]int) []int {
	adjList := make([][]int, numCourses)
	for i := range numCourses {
		adjList[i] = make([]int, 0)
	}

	for _, req := range prerequisites {
		adjList[req[0]] = append(adjList[req[0]], req[1])
	}

	visitedNodes := make(map[int]struct{})
	pathNodes := make(map[int]struct{})
	res := make([]int, 0)

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if _, ok := pathNodes[node]; ok {
			// this is cyclic
			return false
		}

		if _, ok := visitedNodes[node]; ok {
			return true
		}

		visitedNodes[node] = struct{}{}
		pathNodes[node] = struct{}{}

		for _, childNode := range adjList[node] {
			if !dfs(childNode) {
				return false
			}
		}
		res = append(res, node)
		delete(pathNodes, node)
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			// cyclic
			return []int{}
		}
	}
	return res
}

func CourseScheduleIIKahnAlgo(numCourses int, prerequisites [][]int) []int {
	indegree := make([]int, numCourses)
	adjList := make([][]int, numCourses)
	for i := range numCourses {
		adjList[i] = make([]int, 0)
	}

	for _, req := range prerequisites {
		adjList[req[1]] = append(adjList[req[1]], req[0])
		indegree[req[0]]++
	}

	queue := []int{}
	for idx := range indegree {
		if indegree[idx] == 0 {
			queue = append(queue, idx)
		}
	}
	res := []int{}
	visited := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)
		visited++

		for _, child := range adjList[node] {
			indegree[child]--
			if indegree[child] == 0 {
				queue = append(queue, child)
			}
		}

	}

	if visited != numCourses {
		return []int{}
	}

	return res
}
