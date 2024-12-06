package backtracking

import "slices"

func PermutationsRecursion(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	permute := PermutationsRecursion(nums[1:])

	var res [][]int
	for _, perm := range permute {
		for i := 0; i <= len(perm); i++ {
			temp := make([]int, len(perm))
			copy(temp, perm)
			temp = slices.Insert(temp, i, nums[0])
			res = append(res, temp)
		}
	}
	return res
}

func PermutationsIteration(nums []int) [][]int {
	res := [][]int{{}}

	for _, n := range nums {
		perms := make([][]int, 0)
		for _, perm := range res {
			for i := 0; i <= len(perm); i++ {
				temp := make([]int, len(perm))
				copy(temp, perm)
				temp = slices.Insert(temp, i, n)
				perms = append(perms, temp)
			}
		}
		res = perms
	}
	return res
}

func PermutationsBackTracking(nums []int) [][]int {
	res := [][]int{}

	var dfs func([]bool, []int)
	dfs = func(boolMask []bool, perm []int) {
		if len(perm) == len(nums) {
			temp := make([]int, len(perm))
			copy(temp, perm)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if boolMask[i] == true {
				continue
			}

			perm = append(perm, nums[i])
			boolMask[i] = true
			dfs(boolMask, perm)

			boolMask[i] = false
			perm = perm[:len(perm)-1]
		}
	}

	dfs(make([]bool, len(nums)), []int{})
	return res
}

func PermutationsBackOptimal(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func(int, []int)
	dfs = func(idx int, nums []int) {
		if idx == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			res = append(res, temp)
			return
		}

		for i := idx; i < len(nums); i++ {
			nums[i], nums[idx] = nums[idx], nums[i]
			dfs(idx+1, nums)
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	dfs(0, nums)
	return res
}

func PermutationsBackBitOp(nums []int) [][]int {
	res := make([][]int, 0)
	var dfs func([]int, int)
	dfs = func(subset []int, bitmask int) {
		if len(subset) == len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if bitmask&(1<<i) != 0 {
				continue
			}
			subset = append(subset, nums[i])
			dfs(subset, bitmask|(1<<i))
			subset = subset[:len(subset)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}
