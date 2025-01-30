package graph

func CourseSchedule(numCourses int, prerequisites [][]int) bool {
	// create adjency list representation of prerequisites for numCourses
	adjList := make([][]int, numCourses)
	for _, req := range prerequisites {
		adjList[req[1]] = append(adjList[req[1]], req[0])
	}

	// now detect if cycle is present in directed graph using dfs
	visitedNode := make([]bool, numCourses)
	recStack := make([]bool, numCourses)

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if !visitedNode[node] {
			visitedNode[node] = true
			recStack[node] = true

			for _, node := range adjList[node] {
				if !visitedNode[node] && dfs(node) {
					return true
				} else if recStack[node] {
					return true
				}
			}

		}

		recStack[node] = false
		return false
	}

	for i := range numCourses {
		if !visitedNode[i] && len(adjList) != 0 && dfs(i) {
			// cyclic
			return false
		}
	}

	return true
}

func CourseScheduleKahnAlgo(numCourses int, prerequisites [][]int) bool {
	adjList := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for i := range numCourses {
		adjList[i] = make([]int, 0)
	}

	for _, req := range prerequisites {
		adjList[req[1]] = append(adjList[req[1]], req[0])
		inDegree[req[0]]++
	}

	queue := []int{}
	for i, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, i)
		}
	}

	finished := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		finished++
		for _, child := range adjList[node] {
			inDegree[child]--
			if inDegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	return finished == numCourses
}
