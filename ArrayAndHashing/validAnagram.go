package arrayandhashing

func AnagramHash(s string, t string) bool {
	// Space - O(n) and time - O(n)
	// It evens works for unicode

	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, r := range s {
		count[r]++
	}

	for _, r := range t {
		if count[r] == 0 {
			return false
		}
		count[r]--

	}

	return true
}

func AnagramArray(s string, t string) bool {
	// Space: O(1), Time: O(n)
	// It uses the fact that both string have only lower case letter
	if len(s) != len(t) {
		return false
	}

	arr := [26]int{}

	for _, char := range s {
		arr[int(char-'a')]++
	}

	for _, char := range t {
		arr[int(char-'a')]--
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}
