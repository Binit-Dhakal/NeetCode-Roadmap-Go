package binarysearch

func SearchRotatedSortedBS(nums []int, target int) int {
	leftPtr, rightPtr := 0, len(nums)-1

	// find deflection point
	for leftPtr < rightPtr {
		midPtr := (leftPtr + rightPtr) / 2

		if nums[midPtr] > nums[rightPtr] {
			leftPtr = midPtr + 1
		} else {
			rightPtr = midPtr
		}
	}

	pivot := leftPtr

	binarySearch := func(l, r int) int {
		for l <= r {
			m := (l + r) / 2
			if nums[m] == target {
				return m
			} else if nums[m] > target {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		return -1
	}

	res := binarySearch(0, pivot-1)
	if res != -1 {
		return res
	}

	return binarySearch(pivot, len(nums)-1)
}
