package arrays

func MaximumSubArrayBrute(nums []int) int {
	maxSum := nums[0]
	for i := range nums {
		sum := 0
		for _, n := range nums[i:] {
			sum += n
			maxSum = max(sum, maxSum)
		}
	}
	return maxSum
}

func MaximumSubArrayKadenes(nums []int) int {
	maxSum := nums[0]
	sum := 0
	for _, n := range nums {
		sum += n
		maxSum = max(sum, maxSum)
		sum = max(sum, 0)
	}
	return maxSum
}
