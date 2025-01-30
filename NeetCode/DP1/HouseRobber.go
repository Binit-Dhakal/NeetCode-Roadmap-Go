package dp1

func HouseRobberTopDown(nums []int) int {
	cache := make([]int, len(nums))
	for i := range cache {
		cache[i] = -1
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= len(nums) {
			return 0
		}

		if cache[i] != -1 {
			return cache[i]
		}

		cache[i] = nums[i] + max(dfs(i+2), dfs(i+3))
		return cache[i]
	}

	return max(dfs(0), dfs(1))
}

func HouseRobberBottomUp(nums []int) int {
	rob1, rob2 := 0, 0
	for _, num := range nums {
		temp := max(num+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}
