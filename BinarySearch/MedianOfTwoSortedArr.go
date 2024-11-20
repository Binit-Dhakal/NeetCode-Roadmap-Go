package binarysearch

func MedianBruteForce(nums1 []int, nums2 []int) float64 {
	nums := make([]int, len(nums1)+len(nums2))

	leftNumPtr, rightNumPtr := 0, 0
	ptr := 0
	for leftNumPtr < len(nums1) && rightNumPtr < len(nums2) {
		if nums1[leftNumPtr] <= nums2[rightNumPtr] {
			nums[ptr] = nums1[leftNumPtr]
			leftNumPtr++
		} else {
			nums[ptr] = nums2[rightNumPtr]
			rightNumPtr++
		}
		ptr++
	}

	for leftNumPtr < len(nums1) {
		nums[ptr] = nums1[leftNumPtr]
		leftNumPtr++
		ptr++
	}

	for rightNumPtr < len(nums2) {
		nums[ptr] = nums2[rightNumPtr]
		rightNumPtr++
		ptr++
	}

	if len(nums)%2 != 0 {
		medianPtr := len(nums) / 2
		return float64(nums[medianPtr])
	} else {
		medianPtr := len(nums) / 2
		return float64(nums[medianPtr]+nums[medianPtr-1]) / 2
	}
}
