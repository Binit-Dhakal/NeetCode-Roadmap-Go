package dp1

func MinCostClimbingStairsBottomUp(cost []int) int {
	n := len(cost)

	for i := n - 3; i >= 0; i-- {
		cost[i] = cost[i] + min(cost[i+1], cost[i+2])
	}

	return min(cost[0], cost[1])
}

func MinCostClimbingStairsTopDown(cost []int) int {
	cache := make([]int, len(cost))
	for i := 0; i < len(cache); i++ {
		cache[i] = -1
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= len(cost) {
			return 0
		}

		if cache[i] != -1 {
			return cache[i]
		}

		cache[i] = cost[i] + min(dfs(i+1), dfs(i+2))
		return cache[i]
	}

	return min(dfs(0), dfs(1))
}
