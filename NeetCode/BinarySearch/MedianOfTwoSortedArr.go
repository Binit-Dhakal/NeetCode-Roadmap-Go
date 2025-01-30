package binarysearch

import (
	"math"
)

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

func calculateMedian(A []int, B []int) float64 {
	total := len(A) + len(B)
	half := (total + 1) / 2

	leftPtr, rightPtr := 0, len(A)
	var BPtr, midPtr int
	for leftPtr <= rightPtr {
		midPtr = (rightPtr + leftPtr) / 2
		BPtr = half - midPtr - 2

		Bleft := math.MinInt
		if BPtr > 0 {
			Bleft = B[BPtr]
		}

		Bright := math.MaxInt
		if BPtr+1 < len(B) {
			Bright = B[BPtr+1]
		}

		Aleft := math.MinInt
		if midPtr > 0 {
			Aleft = A[midPtr]
		}

		Aright := math.MaxInt
		if midPtr+1 < len(A) {
			Aright = A[midPtr+1]
		}

		if Bleft > Aright {
			leftPtr = midPtr + 1
		} else if Bright < Aleft {
			rightPtr = midPtr - 1
		} else {
			// we found our BPtr and midPtr
			if total%2 != 0 {
				return float64(min(Aright, Bright))
			}
			return float64(min(Aright, Bright)+max(Aleft, Bleft)) / 2.0
		}
	}
	return 0.0
}

func MedianBT(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return calculateMedian(nums2, nums1)
	} else {
		return calculateMedian(nums1, nums2)
	}
}
