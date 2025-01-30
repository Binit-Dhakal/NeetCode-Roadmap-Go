package twopointers

func MaxArea(height []int) int {
	leftPtr := 0
	rightPtr := len(height) - 1
	area, maxArea := 0, 0

	for leftPtr < rightPtr {
		leftVal := height[leftPtr]
		rightVal := height[rightPtr]

		area = min(leftVal, rightVal) * (rightPtr - leftPtr)
		if area > maxArea {
			maxArea = area
		}

		if leftVal <= rightVal {
			leftPtr++
		} else {
			rightPtr--
		}
	}
	return maxArea
}
