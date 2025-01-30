package dp1

func HouseRobberII(nums []int) int {
	// -- recursion formula -- max(nums[i]+dfs(i-2), dfs(i-1))
	// max(nums[n]+dfs(n-2),dfs(n-1)) -- using last house  [1....n]
	// max(nums[n-1]+dfs(n-3),dfs(n-2)) -- skipping last house [0...n-1]
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	nums1 := nums[:n-1]
	nums2 := nums[1:n]
	cache1 := make([]int, n)
	cache2 := make([]int, n)

	for i := 0; i < n; i++ {
		cache1[i] = -1
		cache2[i] = -1
	}

	var dfs func(idx int, nums []int, cache []int) int
	dfs = func(idx int, nums []int, cache []int) int {
		if idx == 0 {
			return nums[idx]
		}
		if idx < 0 {
			return 0
		}

		if cache[idx] != -1 {
			return cache[idx]
		}

		cache[idx] = max(nums[idx]+dfs(idx-2, nums, cache), dfs(idx-1, nums, cache))
		return cache[idx]
	}

	return max(dfs(n-2, nums1, cache1), dfs(n-2, nums2, cache2))
}

func HouseRobberIIBottomUp(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var maxRobbery func(nums []int) int
	maxRobbery = func(nums []int) int {
		a, b := 0, 0
		for _, n := range nums {
			temp := max(a+n, b)
			a = b
			b = temp
		}
		return b
	}

	return max(maxRobbery(nums[1:]), maxRobbery(nums[:len(nums)-1]))
}
