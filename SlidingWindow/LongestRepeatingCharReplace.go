package slidingwindow

func LongestRepeatingCharReplacementTwoPointers(s string, k int) int {
	res := 0
	freqCount := make([]int, 26)

	leftPtr := 0
	maxFreqVal := 0

	for rightPtr := 0; rightPtr < len(s); rightPtr++ {
		windowSize := rightPtr - leftPtr + 1

		freqCount[s[rightPtr]-'A']++
		maxFreqVal = max(maxFreqVal, freqCount[s[rightPtr]-'A'])

		if windowSize-maxFreqVal <= k {
			res = max(res, windowSize)
		} else {
			// invalid window
			freqCount[s[leftPtr]-'A']--
			leftPtr++
		}
	}
	return res
}

func LongestRepeatingCharReplacementBruteForce(s string, k int) int {
	freqCount := make(map[byte]int)
	res := 0
	for leftPtr := 0; leftPtr < len(s); leftPtr++ {
		maxfreqVal := 0
		clear(freqCount)

		for rightPtr := leftPtr; rightPtr < len(s); rightPtr++ {
			windowSize := rightPtr - leftPtr + 1
			freqCount[s[rightPtr]]++
			maxfreqVal = max(maxfreqVal, freqCount[s[rightPtr]])

			if windowSize-maxfreqVal <= k {
				res = max(res, windowSize)
			} else {
				// invalid window
				break
			}
		}
	}
	return res
}
