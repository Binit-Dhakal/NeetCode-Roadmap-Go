package arrays

func MajorityElement(nums []int) int {
	// Boyer-Moore Majority Vote Algorithm
	maxElement := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			maxElement = num
			count++
			continue
		}

		if maxElement == num {
			count++
		} else {
			count--
		}
	}

	return maxElement
}

func MajorityElementBitwise(nums []int) int {
	majority := 0
	n := len(nums)

	for i := 0; i < 32; i++ {
		count := 0
		mask := 1 << i

		for _, num := range nums {
			if mask&num != 0 {
				count++
			}
		}

		if count > n/2 {
			majority |= mask
		}
	}

	// handling negative number
	if (1<<31)&majority != 0 {
		majority -= (1 << 32)
	}

	return majority
}
