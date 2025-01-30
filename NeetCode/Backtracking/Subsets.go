package backtracking

func SubsetsIteration(nums []int) [][]int {
	res := [][]int{{}}

	for _, num := range nums {
		n := len(res)
		for i := 0; i < n; i++ {
			subset := make([]int, len(res[i]))
			copy(subset, res[i])
			subset = append(subset, num)
			res = append(res, subset)
		}
	}

	return res
}

func SubsetsBackTracking(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}

	var dfs func(idx int)

	dfs = func(idx int) {
		if idx >= len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}

		subset = append(subset, nums[idx])
		dfs(idx + 1)
		subset = subset[:len(subset)-1]
		dfs(idx + 1)
	}

	dfs(0)
	return res
}

func SubsetsBitMask(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)

	for i := 0; i < (1 << n); i++ {
		subset := []int{}
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				subset = append(subset, nums[j])
			}
		}
		res = append(res, subset)
	}
	return res
}
