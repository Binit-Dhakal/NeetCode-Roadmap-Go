package arrays

func LongestConsecutiveSequence(nums []int) int {
	hash := make(map[int]int)
	maxSequence := 0

	for _, num := range nums {
		hash[num] = 1
	}

	for num := range hash {
		temp := 1
		if _, ok := hash[num-1]; !ok {
			n := num + 1
			for {
				if _, ok2 := hash[n]; !ok2 {
					break
				}

				temp += hash[n]
				n++
			}
			maxSequence = max(maxSequence, temp)
		}
	}

	return maxSequence
}
