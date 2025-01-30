package dp1

func PalindromicSubstringDP(s string) int {
	// O(n^2) time and O(n^2) space
	n := len(s)
	cache := make([][]bool, n)
	for i := range n {
		cache[i] = make([]bool, n)
	}
	count := 0

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || cache[i+1][j-1]) {
				cache[i][j] = true
				count++
			}
		}
	}
	return count
}

func PalindromicSubstringTwoPointers(s string) int {
	n := len(s)
	count := 0
	var countPali func(l, r int)
	countPali = func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			count++
			l--
			r++
		}
	}
	for i := 0; i < n; i++ {
		count++

		countPali(i-1, i+1)
		countPali(i, i+1)

	}
	return count
}
