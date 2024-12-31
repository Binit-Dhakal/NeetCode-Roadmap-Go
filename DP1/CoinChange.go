package dp1

func CoinChangeDFS(coins []int, amount int) int {
	var dfs func(idx int, amt int) int
	dfs = func(idx, amt int) int {
		if idx == 0 {
			if amt%coins[0] == 0 {
				return amt / coins[0]
			}
			return 1e9
		}

		notPick := 0 + dfs(idx-1, amt)
		var pick int = 1e9
		if coins[idx] <= amt {
			pick = 1 + dfs(idx, amt-coins[idx])
		}
		return min(pick, notPick)
	}

	ans := dfs(len(coins)-1, amount)
	if ans == 1e9 {
		return -1
	}
	return ans
}

func CoinChangeDFSMemoization(coins []int, amount int) int {
	dp := make([][]int, len(coins))
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(idx int, amt int) int
	dfs = func(idx, amt int) int {
		if idx == 0 {
			if amt%coins[0] == 0 {
				return amt / coins[0]
			}
			return 1e9
		}

		if dp[idx][amt] != -1 {
			return dp[idx][amt]
		}
		notPick := 0 + dfs(idx-1, amt)
		var pick int = 1e9
		if coins[idx] <= amt {
			pick = 1 + dfs(idx, amt-coins[idx])
		}
		dp[idx][amt] = min(pick, notPick)
		return dp[idx][amt]
	}

	ans := dfs(len(coins)-1, amount)
	if ans == 1e9 {
		return -1
	}
	return ans
}

func CoinChangeOptimal(coins []int, amount int) int {
	prev := make([]int, amount+1) // i-1
	cur := make([]int, amount+1)  // i

	for amt := 0; amt <= amount; amt++ {
		if amt%coins[0] == 0 {
			prev[amt] = amt / coins[0]
		} else {
			prev[amt] = 1e9
		}
	}

	for i := 1; i < len(coins); i++ {
		for amt := 0; amt <= amount; amt++ {
			notPick := prev[amt]
			var pick int = 1e9
			if coins[i] <= amt {
				pick = 1 + cur[amt-coins[i]]
			}
			cur[amt] = min(notPick, pick)
		}
		prev = cur
	}

	ans := prev[amount]
	if ans >= 1e9 {
		return -1
	}
	return ans
}
