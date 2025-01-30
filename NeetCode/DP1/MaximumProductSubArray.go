package dp1

func MaxProductSubArrayKadenes(nums []int) int {
	curMax, curMin := 1, 1
	res := nums[0]
	for _, num := range nums {
		tmp := curMax * num
		curMax = max(tmp, num*curMin, num)
		curMin = min(tmp, num*curMin, num)
		res = max(res, curMax)
	}
	return res
}

func MaxProductSubArrayPrefixSuffix(nums []int) int {
	n := len(nums) - 1
	res := nums[0]
	prefix, suffix := 1, 1
	for i := 0; i <= n; i++ {
		if prefix == 0 {
			prefix = 1
		}
		if suffix == 0 {
			suffix = 1
		}

		prefix *= nums[i]
		suffix *= nums[n-i]
		res = max(res, max(prefix, suffix))
	}
	return res
}
