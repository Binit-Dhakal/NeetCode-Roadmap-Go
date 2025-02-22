package recursion

func SubsetSum(nums []int) []int {
	sums := make([]int, 0)
	sums = append(sums, 0)

	for _, num := range nums {
		for _, sum := range sums {
			sums = append(sums, sum+num)
		}
	}
	return sums
}

func SubsetSumRecursion(nums []int) []int {
	sums := make([]int, 0)

	var rec func(idx, sum int)

	rec = func(idx, sum int) {
		if idx == len(nums) {
			sums = append(sums, sum)
			return
		}

		rec(idx+1, sum)
		rec(idx+1, sum+nums[idx])
	}

	rec(0, 0)
	return sums
}
