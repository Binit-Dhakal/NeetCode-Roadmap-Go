package slidingwindow

func LongestSubstringWithoutRepeatingTwoPointers(s string) int {
	if len(s) == 0 {
		return 0
	}

	countMap := make(map[byte]int)

	leftPtr, rightPtr := 0, 1
	longestSub := 0

	countMap[s[leftPtr]] = leftPtr

	for rightPtr < len(s) {
		if idx, ok := countMap[s[rightPtr]]; ok && idx >= leftPtr {
			longestSub = max(longestSub, rightPtr-leftPtr)
			leftPtr = idx + 1
		}
		countMap[s[rightPtr]] = rightPtr
		rightPtr++
	}

	longestSub = max(longestSub, rightPtr-leftPtr)
	return longestSub
}

func LongestSubstringWithoutRepeatingOptimized(s string) int {
	countMap := make(map[rune]int)

	longestSub := 0
	start := 0

	for i, val := range s {
		if idx, ok := countMap[val]; ok && idx >= start {
			start = idx + 1
		}

		countMap[val] = i
		longestSub = max(longestSub, i-start+1)
	}

	return longestSub
}
