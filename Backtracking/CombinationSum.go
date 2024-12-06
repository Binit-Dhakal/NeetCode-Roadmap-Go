package backtracking

func CombinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	subset := make([]int, 0)
	sum := 0

	var dfs func(idx int)
	dfs = func(idx int) {
		if sum > target || idx >= len(candidates) {
			return
		}
		if sum == target {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}

		sum += candidates[idx]
		subset = append(subset, candidates[idx])
		dfs(idx)
		sum -= candidates[idx]
		subset = subset[:len(subset)-1]
		dfs(idx + 1)
	}

	dfs(0)

	return res
}
