package arrays

import (
	"slices"
)

func NextPermutation(nums []int) {
	// finding pivot
	pivot := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			pivot = i
			break
		}
	}

	// we might be dealing with reverse array in this case
	if pivot == -1 {
		slices.Sort(nums)
		return
	}

	// setting initial successor which is greater than pivot
	for i := len(nums) - 1; i > pivot; i-- {
		if nums[i] > nums[pivot] {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			break
		}
	}

	// Reverse from pivot+1
	slices.Reverse(nums[pivot+1:])
}
