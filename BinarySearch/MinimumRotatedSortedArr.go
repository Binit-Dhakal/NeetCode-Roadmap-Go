package binarysearch

func MinRotatedSortedArrBS(nums []int) int {
	leftPtr, rightPtr := 0, len(nums)-1
	res := nums[0]

	for leftPtr <= rightPtr {
		if nums[leftPtr] < res {
			res = nums[leftPtr]
		}

		midPtr := (rightPtr + leftPtr) / 2
		if nums[midPtr] >= nums[leftPtr] {
			leftPtr = midPtr + 1
		} else {
			rightPtr = midPtr
		}
	}
	return res
}

func MinRotatedSortedArrBSRight(nums []int) int {
	leftPtr, rightPtr := 0, len(nums)-1
	for leftPtr < rightPtr {
		midPtr := (rightPtr + leftPtr) / 2
		if nums[midPtr] > nums[rightPtr] {
			leftPtr = midPtr + 1
		} else {
			rightPtr = midPtr
		}
	}
	return nums[leftPtr]
}
