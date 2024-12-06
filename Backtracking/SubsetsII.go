package backtracking

import "slices"

func SubsetsII(nums []int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	subset := make([]int, 0)

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}

		subset = append(subset, nums[idx])
		dfs(idx + 1)
		subset = subset[:len(subset)-1]
		for idx+1 < len(nums) && nums[idx] == nums[idx+1] {
			idx++
		}
		dfs(idx + 1)
	}

	dfs(0)
	return res
}

func SubsetsIIIterative(nums []int) [][]int {
	res := [][]int{{}}
	prevIdx, idx := 0, 0

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			idx = prevIdx
		} else {
			idx = 0
		}

		prevIdx = len(res)
		for j := idx; j < prevIdx; j++ {
			temp := make([]int, len(res[j]))
			copy(temp, res[j])
			temp = append(temp, nums[i])
			res = append(res, temp)
		}
	}
	return res
}
