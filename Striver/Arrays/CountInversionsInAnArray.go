package arrays

func CountInversions(nums []int) int {
	inversions := 0

	merge := func(low, mid, high int) {
		temp := make([]int, 0)

		leftPtr, rightPtr := low, mid+1
		for leftPtr <= mid && rightPtr <= high {
			if nums[leftPtr] <= nums[rightPtr] {
				temp = append(temp, nums[leftPtr])
				leftPtr++
			} else {
				inversions += (mid - leftPtr + 1)
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

	var mergeSort func(low, high int)
	mergeSort = func(low, high int) {
		if low >= high {
			return
		}
		mid := (low + high) / 2
		mergeSort(low, mid)
		mergeSort(mid+1, high)
		merge(low, mid, high)
	}

	mergeSort(0, len(nums)-1)
	return inversions
}
