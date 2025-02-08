package arrays

import (
	"slices"
)

func FourSum(nums []int, target int) [][]int {
	slices.Sort(nums)

	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			t := target - nums[i] - nums[j]
			leftPtr, rightPtr := j+1, len(nums)-1
			for leftPtr < rightPtr {
				if nums[leftPtr]+nums[rightPtr] < t {
					leftPtr++
				} else if nums[leftPtr]+nums[rightPtr] > t {
					rightPtr--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[leftPtr], nums[rightPtr]})
					leftPtr++
					rightPtr--

					for leftPtr < rightPtr {
						if nums[leftPtr] == nums[leftPtr-1] {
							leftPtr++
						} else {
							break
						}
					}

					for leftPtr < rightPtr {
						if nums[rightPtr] == nums[rightPtr+1] {
							rightPtr--
						} else {
							break
						}
					}
				}
			}
		}
	}
	return result
}
