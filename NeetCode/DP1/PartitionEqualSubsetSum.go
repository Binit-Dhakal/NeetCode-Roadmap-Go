package dp1

func PartitionEqualSubsetSumTopDown(nums []int) bool {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	if totalSum%2 != 0 {
		// odd sum cannot have two equal partition
		return false
	}
	target := totalSum / 2
	n := len(nums)

	// dp[i][sum] will store if we can acheive 'sum' using elements from index i onwards
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(idx int, currentSum int) bool
	dfs = func(idx int, currentSum int) bool {
		if currentSum == target {
			return true
		}

		if idx == len(nums) || currentSum > target {
			return false
		}

		if dp[idx][currentSum] != -1 {
			return dp[idx][currentSum] == 1
		}

		// not pick
		if dfs(idx+1, currentSum) {
			dp[idx][currentSum] = 1
			return true
		}

		// pick
		if dfs(idx+1, currentSum+nums[idx]) {
			dp[idx][currentSum] = 1
			return true
		}
		dp[idx][currentSum] = 0
		return false
	}

	return dfs(0, 0)
}

func PartitionEqualSubsetSumBottomUp(nums []int) bool {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	if totalSum%2 != 0 {
		// odd sum cannot have two equal partition
		return false
	}

	target := totalSum / 2
	n := len(nums)

	// dp[i][sum] will store if we can acheive 'sum' using elements from index i onwards
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
}
