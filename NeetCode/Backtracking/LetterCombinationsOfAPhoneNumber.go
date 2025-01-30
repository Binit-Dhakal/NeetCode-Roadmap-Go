package backtracking

func LetterCombinations(digits string) []string {
	n := len(digits)

	res := make([]string, 0)
	if n == 0 {
		return res
	}
	atoiMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var dfs func(idx int, formedStr string)

	dfs = func(idx int, formedStr string) {
		if len(formedStr) == n {
			res = append(res, formedStr)
			return
		}

		for _, ch := range atoiMap[digits[idx]] {
			dfs(idx+1, formedStr+string(ch))
		}
	}

	dfs(0, "")
	return res
}

func LetterCombinationIteration(digits string) []string {
	n := len(digits)

	if n == 0 {
		return []string{}
	}

	res := []string{""}
	atoiMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	for i := 0; i < n; i++ {
		digit := digits[i]
		tmp := []string{}
		for _, curStr := range res {
			for _, c := range atoiMap[digit] {
				tmp = append(tmp, curStr+string(c))
			}
		}
		res = tmp
	}
	return res
}
