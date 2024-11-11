package arrayandhashing

func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	// prefix slice calculation
	num := 1
	for i := range nums {
		res[i] = num
		num *= nums[i]
	}

	// inplace prefix and suffix
	num = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= num
		num *= nums[i]
	}

	return res
}
