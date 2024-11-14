package twopointers

func TrapPrefixSuffixArray(height []int) int {
	prefix := make([]int, len(height))
	suffix := make([]int, len(height))

	for i := 0; i < len(height)-1; i++ {
		prefix[i+1] = max(height[i], prefix[i])
	}

	for i := len(height) - 1; i > 0; i-- {
		suffix[i-1] = max(height[i], suffix[i])
	}

	totalHeight := 0

	for i := 0; i < len(height); i++ {
		h := min(prefix[i], suffix[i]) - height[i]
		if h > 0 {
			totalHeight += h
		}
	}

	return totalHeight
}

func TrapTwoPointers(height []int) int {
	leftPtr, rightPtr := 0, len(height)-1
	maxLeftValue, maxRightValue := height[leftPtr], height[rightPtr]
	totalHeight := 0

	for leftPtr < rightPtr {
		if maxLeftValue <= maxRightValue {
			leftPtr++
			maxLeftValue = max(maxLeftValue, height[leftPtr])
			totalHeight += maxLeftValue - height[leftPtr]
		} else {
			rightPtr--
			maxRightValue = max(maxRightValue, height[rightPtr])
			totalHeight += maxRightValue - height[rightPtr]
		}
	}
	return totalHeight
}
