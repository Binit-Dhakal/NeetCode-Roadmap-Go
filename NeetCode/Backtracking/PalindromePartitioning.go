package backtracking

func PalindromePartitioning(s string) [][]string {
	res := [][]string{}

	isPalindrome := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	var dfs func(int, []string)
	dfs = func(ptr int, path []string) {
		if ptr == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := ptr; i < len(s); i++ {
			if isPalindrome(ptr, i) {
				dfs(i+1, append(path, string(s[ptr:i+1])))
			}
		}
	}

	dfs(0, []string{})
	return res
}

func PalindromePartitioningDP(s string) [][]string {
	res := [][]string{}
	n := len(s)
	dp := make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// pre calculating palindrome
	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			dp[i][i+l-1] = (s[i] == s[i+l-1]) && (i+1 > (i+l-2) || dp[i+1][i+l-2])
		}
	}

	var dfs func(int, []string)
	dfs = func(ptr int, path []string) {
		if ptr == n {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := ptr; i < n; i++ {
			if dp[ptr][i] {
				dfs(i+1, append(path, string(s[ptr:i+1])))
			}
		}
	}

	dfs(0, []string{})
	return res
}
