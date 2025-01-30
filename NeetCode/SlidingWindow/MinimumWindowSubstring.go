package slidingwindow

func MinimumWindowSubstringBF(s string, t string) string {
	if t == "" {
		return ""
	}
	tMap := make(map[byte]int)

	for i := range t {
		tMap[t[i]]++
	}

	res := [2]int{-1, -1}
	resLen := int(^uint(0) >> 1) // maxint value

	for i := 0; i < len(s); i++ {
		if _, ok := tMap[s[i]]; !ok {
			continue
		}
		sMap := make(map[byte]int)
		for j := i; j < len(s); j++ {
			sMap[s[j]]++

			flag := true
			for k, v := range tMap {
				if v > sMap[k] {
					flag = false
					break
				}
			}

			if flag && (j-i+1) < resLen {
				res = [2]int{i, j}
				resLen = j - i + 1
				// break
			}
		}
	}

	if res == [2]int{-1, -1} {
		return ""
	}

	return s[res[0] : res[1]+1]
}

func MinimumWindowSubstringSW(s string, t string) string {
	if t == "" {
		return ""
	}
	tMap := make(map[byte]int)
	for i := range t {
		tMap[t[i]]++
	}

	have, need := 0, len(tMap)
	leftPtr := 0
	res := [2]int{-1, -1}
	resLen := int(^uint(0) >> 1)

	window := make(map[byte]int)
	for rightPtr := 0; rightPtr < len(s); rightPtr++ {
		if _, ok := tMap[s[rightPtr]]; !ok {
			continue
		}

		window[s[rightPtr]]++

		if window[s[rightPtr]] == tMap[s[rightPtr]] {
			have++
		}

		for have == need {
			if rightPtr-leftPtr+1 < resLen {
				resLen = rightPtr - leftPtr + 1
				res = [2]int{leftPtr, rightPtr}
			}

			if window[s[leftPtr]] > 0 && window[s[leftPtr]] == tMap[s[leftPtr]] {
				have--
			}
			window[s[leftPtr]]--
			leftPtr++

		}

	}

	if res[0] == -1 {
		return ""
	}
	return s[res[0] : res[1]+1]
}
