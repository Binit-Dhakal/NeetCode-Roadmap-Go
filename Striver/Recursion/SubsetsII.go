package recursion

import "sort"

func SubsetsII(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)
	var rec func(idx int, subset []int)
	rec = func(idx int, subset []int) {
		result = append(result, subset)
		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			rec(i+1, append(subset, nums[i]))
		}
	}

	rec(0, []int{})
	return result
}

func SubsetsIICascading(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	result = append(result, []int{})

	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for _, res := range result {
			result = append(result, append(res, num))
		}
	}
	return result
}
