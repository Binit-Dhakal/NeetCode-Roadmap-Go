package arrays

func TwoSumHashMap(nums []int, target int) []int {
	hash := make(map[int]int)

	for idx, num := range nums {
		if jdx, ok := hash[num]; ok {
			return []int{jdx, idx}
		}

		hash[target-num] = idx
	}
	return []int{-1, -1}
}
