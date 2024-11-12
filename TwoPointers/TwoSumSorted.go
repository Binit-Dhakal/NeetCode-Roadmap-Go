package twopointers

func TwoSumSortedMap(numbers []int, target int) []int {
	leftPtr, rightPtr := 0, len(numbers)-1
	for leftPtr < rightPtr {
		sum := numbers[leftPtr] + numbers[rightPtr]
		if sum < target {
			leftPtr++
		} else if sum > target {
			rightPtr--
		} else {
			return []int{leftPtr + 1, rightPtr + 1}
		}
	}
	return []int{-1, -1}
}
