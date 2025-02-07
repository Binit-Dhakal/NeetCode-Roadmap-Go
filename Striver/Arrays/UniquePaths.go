package arrays

func UniquePathsTopDown(m int, n int) int {
	dp := make([]int, m*n)
	for i := range dp {
		dp[i] = -1 // will need to check this
	}
	dp[m*n-1] = 1

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x >= m || y >= n {
			// out of grid
			return 0
		}

		if dp[x*n+y] != -1 {
			return dp[x*n+y]
		}

		// sum of right + bottom
		dp[x*n+y] = dfs(x+1, y) + dfs(x, y+1)

		return dp[x*n+y]
	}

	dfs(0, 0)
	return dp[0]
}

func UniquePathsBottomUp(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if r > 0 {
				dp[r][c] += dp[r-1][c]
			}
			if c > 0 {
				dp[r][c] += dp[r][c-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func UniquePathsBottomUpOptimized(m, n int) int {
	dp := make([]int, n)

	dp[0] = 1
	for r := 0; r < m; r++ {
		for c := 1; c < n; c++ {
			dp[c] = dp[c] + dp[c-1]
		}
	}
	return dp[n-1]
}

func UniquePathsCombinatorics(m, n int) int {
	// (m+n-2)C(m-1)
	if m > n {
		m, n = n, m
	}

	path := 1
	for r := 0; r < m-1; r++ {
		path = path * (m + n - 2 - r) / (r + 1)
	}
	return path
}
