package linkedlist

func FindDuplicateNumberNegativeMarking(nums []int) int {
	// not valid as we are modifying input array
	abs := func(num int) int {
		if num > 0 {
			return num
		}
		return -num
	}

	for i := range nums {
		if nums[abs(nums[i])-1] < 0 {
			return abs(nums[i])
		}
		nums[abs(nums[i])-1] *= -1
	}

	return 0
}

func FindDuplicateNumberFloydsAlgo(nums []int) int {
	// we take each num in nums as pointer to next num
	// and we try to find start of cycle
	// this is problem of linked list solved using Floyds Algorithm

	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		slow2 = nums[slow2]
		slow = nums[slow]
		if slow == slow2 {
			// we found our cycle first index
			return slow
		}
	}
}

func FindDuplicateNumberBinarySearch(nums []int) int {
	n := len(nums)
	low, high := 1, n-1

	for low < high {
		mid := low + (high-low)/2
		lessOrEqual := 0

		for _, num := range nums {
			if num <= mid {
				lessOrEqual++
			}
		}

		if lessOrEqual <= mid {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

func FindDuplicateNumberXOR(nums []int) int {
	duplicate := 0
	n := len(nums)

	for bit := range 32 {
		count_nums := 0
		count_range := 0
		mask := 1 << bit

		for _, num := range nums {
			if num&mask > 0 {
				count_nums++
			}
		}

		for i := 1; i < n; i++ {
			if i&mask > 0 {
				count_range++
			}
		}

		if count_nums > count_range {
			duplicate |= mask
		}
	}

	return duplicate
}
