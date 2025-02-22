package recursion

func CombinationSum(candidates []int, target int) [][]int {
	sum := 0
	result := make([][]int, 0)
	var rec func(idx int, subset []int)
	rec = func(idx int, subset []int) {
		if sum >= target {
			if sum == target {
				s := make([]int, len(subset))
				copy(s, subset)
				result = append(result, s)
			}
			return
		}

		if idx >= len(candidates) {
			return
		}

		sum += candidates[idx]
		rec(idx, append(subset, candidates[idx]))
		sum -= candidates[idx]
		rec(idx+1, subset)
	}

	rec(0, []int{})
	return result
}

func CombinationSumBackTracking(candidates []int, target int) [][]int {
	sum := 0
	result := make([][]int, 0)
	var backtracking func(idx int, subset []int)
	backtracking = func(idx int, subset []int) {
		if sum >= target {
			if sum == target {
				s := make([]int, len(subset))
				copy(s, subset)
				result = append(result, s)
			}
			return
		}

		for i := idx; i < len(candidates); i++ {
			sum += candidates[i]
			backtracking(i, append(subset, candidates[i]))
			sum -= candidates[i]
		}
	}

	backtracking(0, []int{})
	return result
}
