package arrays

func SortColorsCounting(nums []int) {
	countColors := [3]int{}

	for _, num := range nums {
		countColors[num]++
	}

	i := 0
	for idx, color := range countColors {
		for color != 0 {
			nums[i] = idx
			color--
			i++
		}
	}
}

func SortColorsThreePtrs(nums []int) {
	low, cur, high := 0, 0, len(nums)-1
	for cur <= high {
		switch nums[cur] {
		case 0:
			nums[low], nums[cur] = nums[cur], nums[low]
			low++
			cur++
		case 1:
			cur++
		case 2:
			nums[high], nums[cur] = nums[cur], nums[high]
			high--
		}
	}
}
