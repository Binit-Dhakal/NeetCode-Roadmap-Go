package binarysearch

import "sort"

func IterativeSearch(nums []int, target int) int {
	leftPtr, rightPtr := 0, len(nums)-1

	for leftPtr <= rightPtr {
		midPtr := (rightPtr + leftPtr) / 2
		if nums[midPtr] < target {
			leftPtr = midPtr + 1
		} else if nums[midPtr] > target {
			rightPtr = midPtr - 1
		} else {
			return midPtr
		}
	}

	return -1
}

func RecursiveSearch(nums []int, target int) int {
	var searchFn func(int, int) int
	searchFn = func(leftPtr, rightPtr int) int {
		if leftPtr > rightPtr {
			return -1
		}

		midPtr := (leftPtr + rightPtr) / 2
		if nums[midPtr] < target {
			return searchFn(midPtr+1, rightPtr)
		} else if nums[midPtr] > target {
			return searchFn(leftPtr, midPtr-1)
		} else {
			return midPtr
		}
	}

	return searchFn(0, len(nums)-1)
}

func BuiltInSearch(nums []int, target int) int {
	idx := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})

	if idx < len(nums) && nums[idx] == target {
		return idx
	} else {
		return -1
	}
}
