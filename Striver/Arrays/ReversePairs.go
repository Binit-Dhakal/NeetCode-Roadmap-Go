package arrays

func ReversePairs(nums []int) int {
	pairsCnt := 0
	merge := func(low, mid, high int) {
		temp := []int{}
		leftPtr := low
		rightPtr := mid + 1

		for leftPtr <= mid && rightPtr <= high {
			if nums[leftPtr] < nums[rightPtr] {
				temp = append(temp, nums[leftPtr])
				leftPtr++
			} else {
				temp = append(temp, nums[rightPtr])
				rightPtr++
			}
		}

		for leftPtr <= mid {
			temp = append(temp, nums[leftPtr])
			leftPtr++
		}

		for rightPtr <= high {
			temp = append(temp, nums[rightPtr])
			rightPtr++
		}

		for i := low; i <= high; i++ {
			nums[i] = temp[i-low]
		}
	}

	countPairs := func(low, mid, high int) {
		rightPtr := mid + 1

		for leftPtr := low; leftPtr <= mid; leftPtr++ {
			for rightPtr <= high && nums[leftPtr] > 2*nums[rightPtr] {
				rightPtr++
			}

			pairsCnt += (rightPtr - mid - 1)
		}
	}

	var divide func(low, high int)
	divide = func(low, high int) {
		if low >= high {
			return
		}
		mid := (low + high) / 2
		divide(low, mid)
		divide(mid+1, high)
		countPairs(low, mid, high)

		merge(low, mid, high)
	}

	divide(0, len(nums)-1)
	return pairsCnt
}
