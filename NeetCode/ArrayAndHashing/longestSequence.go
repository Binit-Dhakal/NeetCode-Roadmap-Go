package arrayandhashing

func SetlongestConsecutive(nums []int) int {
	longest := 0

	seqArr := make(map[int]struct{})
	for _, num := range nums {
		seqArr[num] = struct{}{}
	}

	for key := range seqArr {
		if _, ok := seqArr[key-1]; ok {
			continue // not the start of chain
		}
		chainLen := 1

		for {
			if _, ok := seqArr[key+1]; ok {
				key += 1
				chainLen++
			} else {
				break
			}
		}

		if chainLen > longest {
			longest = chainLen
		}
	}
	return longest
}
