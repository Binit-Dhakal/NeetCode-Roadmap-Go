package arrayandhashing

func TwoSumBrute(nums []int, target int) []int {
	// Time-O(n^2) and Space- O(1)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSumHash(nums []int, target int) []int {
	// Time- O(n) and Space- O(n)

	hashmap := make(map[int]int)

	for idx, num := range nums {
		if idy, ok := hashmap[num]; ok {
			return []int{idy, idx}
		}
		hashmap[target-num] = idx

	}
	return nil
}
