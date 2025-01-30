package dp1

func DecodeWaysDFS(s string) int {
	cache := make([]int, len(s)+1)
	for i := range len(s) {
		cache[i] = -1
	}
	cache[len(s)] = 1

	var dfs func(idx int) int
	dfs = func(idx int) int {
		if cache[idx] != -1 {
			return cache[idx]
		}

		if s[idx] == '0' {
			return 0
		}

		res := 0
		// left subtree - pointer move by 1
		res += dfs(idx + 1)

		// right subtree
		if idx+1 < len(s) && (s[idx] == '1' || (s[idx] == '2' && s[idx+1] <= '6')) {
			res += dfs(idx + 2)
		}
		cache[idx] = res
		return res
	}

	return dfs(0)
}

func DecodeWaysBottomUp(s string) int {
	cache := make([]int, len(s)+1)
	cache[len(s)] = 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			cache[i] = 0
		} else {
			cache[i] = cache[i+1]
			if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && s[i+1] <= '6')) {
				cache[i] += cache[i+2]
			}
		}
	}

	return cache[0]
}

func DecodeWaysBottomUpOptimized(s string) int {
	a, b, temp := 1, 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			b, a = a, 0
		} else {
			temp = a
			if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && s[i+1] <= '6')) {
				temp += b
			}

			b, a = a, temp
		}
	}

	return a
}
