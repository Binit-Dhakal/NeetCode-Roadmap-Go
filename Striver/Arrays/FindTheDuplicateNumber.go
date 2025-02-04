package arrays

func DuplicateNumber(nums []int) int {
	cur := nums[0]
	for {
		// might change infinite loop
		next := nums[cur]
		if nums[cur] == -1 {
			return cur
		}
		nums[cur] = -1
		cur = next
	}
}

func DuplicateNumberBinarySearch(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) / 2
		lessOrEqualCnt := 0
		for _, num := range nums {
			if num <= mid {
				lessOrEqualCnt++
			}
		}

		if lessOrEqualCnt <= mid {
			low = mid + 1
		} else {
			high = mid
		}

	}
	return low
}

func DuplicateNumberCycleDetection(nums []int) int {
	slowPtr := 0
	fastPtr := 0

	// always cyclic
	for {
		slowPtr = nums[slowPtr]
		fastPtr = nums[nums[fastPtr]]

		if slowPtr == fastPtr {
			break
		}
	}

	slowPtr = 0

	for {
		slowPtr = nums[slowPtr]
		fastPtr = nums[fastPtr]
		if slowPtr == fastPtr {
			return slowPtr
		}
	}
}

func DuplicateNumberXOR(nums []int) int {
	// Doesn't work for more than one duplicates number like [3,3,3,3,3]
	xor_nums := 0
	for _, num := range nums {
		xor_nums ^= num
	}

	xor_range := 0
	for i := range len(nums) {
		xor_range ^= i
	}

	return xor_nums ^ xor_range
}
