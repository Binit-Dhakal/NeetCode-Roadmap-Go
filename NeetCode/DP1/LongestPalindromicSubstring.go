package dp1

func LongestPalindromicSubstringBF(s string) string {
	// O(n^3) time
	n := len(s)
	longest := ""
	sizePalindrome := 0

	var isPalindrome func(t string) bool
	isPalindrome = func(t string) bool {
		i, j := 0, len(t)-1
		for i < j {
			if t[i] != t[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j-i+1 <= sizePalindrome {
				continue
			}
			if isPalindrome(s[i : j+1]) {
				longest = s[i : j+1]
				sizePalindrome = j - i + 1
			}
		}
	}

	return longest
}

func LongestPalindromicSubstringTwoPointers(s string) string {
	// O(n^2) and O(1)
	longest := [2]int{0, 1} // l,size
	n := len(s)

	var expandAroundCenter func(l, r int)
	expandAroundCenter = func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > longest[1] {
				longest = [2]int{l, r - l + 1}
			}

			l--
			r++
		}
	}
	for i := 0; i < n; i++ {
		// odd length
		expandAroundCenter(i-1, i+1)
		// even length
		expandAroundCenter(i, i+1)
	}
	return s[longest[0] : longest[0]+longest[1]]
}

func LongestPalindromicSubstringDP(s string) string {
	// O(n^2) and O(n^2)
	maxLength := 0
	idx := 0

	n := len(s)
	cache := make([][]bool, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i+1 <= 2 || cache[i+1][j-1]) {
				cache[i][j] = true
				if j-i+1 > maxLength {
					maxLength = j - i + 1
					idx = i
				}
			}
		}
	}
	return s[idx : idx+maxLength]
}
