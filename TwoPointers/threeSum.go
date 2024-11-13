package twopointers

import "slices"

func ThreeSumSort(nums []int) [][]int {
	slices.Sort(nums)
	results := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		target := -nums[i]
		leftPtr, rightPtr := i+1, len(nums)-1

		for leftPtr < rightPtr {
			s := nums[leftPtr] + nums[rightPtr]
			if s < target {
				leftPtr++
			} else if s > target {
				rightPtr--
			} else {
				results = append(results, []int{-target, nums[leftPtr], nums[rightPtr]})
				leftPtr++
				for leftPtr < rightPtr && nums[leftPtr] == nums[leftPtr-1] {
					leftPtr++
					continue
				}

				rightPtr--
				if leftPtr < rightPtr && nums[rightPtr] == nums[rightPtr+1] {
					rightPtr--
					continue
				}
			}
		}
	}
	return results
}
