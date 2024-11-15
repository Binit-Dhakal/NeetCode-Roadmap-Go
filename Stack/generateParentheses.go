package stack

func GenerateParenthesesBruteForce(n int) []string {
	result := make([]string, 0)

	isValid := func(val string) bool {
		open := 0
		for _, v := range val {
			if v == '(' {
				open++
			} else {
				open--
			}
			if open < 0 {
				return false
			}
		}
		return open == 0
	}

	// this is like a binary tree
	var dfs func(string)
	dfs = func(s string) {
		if len(s) == n*2 {
			if isValid(s) {
				result = append(result, s)
			}
			return
		}

		dfs(s + "(")
		dfs(s + ")")
	}

	dfs("")
	return result
}

func GenerateParenthesesBackTracking(n int) []string {
	results := []string{}
	var backTracking func(string, int, int)
	backTracking = func(s string, openN int, closeN int) {
		if (closeN == openN) && (len(s) == 2*n) {
			results = append(results, s)
			return
		}

		if openN < n {
			backTracking(s+"(", openN+1, closeN)
		}

		if closeN < openN {
			backTracking(s+")", openN, closeN+1)
		}
	}

	backTracking("", 0, 0)
	return results
}
