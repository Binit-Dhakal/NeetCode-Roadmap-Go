package backtracking

import "slices"

func CombinationSumII(candidates []int, target int) [][]int {
	res := [][]int{}
	slices.Sort(candidates)

	var dfs func(idx int, subset []int, sum int)
	dfs = func(idx int, subset []int, sum int) {
		if sum == target {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}
		if idx == len(candidates) || sum > target {
			return
		}

		dfs(idx+1, append(subset, candidates[idx]), sum+candidates[idx])
		for idx+1 < len(candidates) && candidates[idx+1] == candidates[idx] {
			idx++
		}
		dfs(idx+1, subset, sum)
	}

	dfs(0, []int{}, 0)
	return res
}

func CombinationSumIIOptimal(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var dfs func(idx int, path []int, cur int)
	dfs = func(idx int, path []int, cur int) {
		if cur == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}

			if cur+candidates[i] > target {
				break
			}

			dfs(i+1, append(path, candidates[i]), cur+candidates[i])
		}
	}
	dfs(0, []int{}, 0)

	return res
}
