package arrays

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	// O(n^2) approach
	total := m + n
	leftPtr := 0
	rightPtr := 0

	for leftPtr < m && rightPtr < n {
		if nums1[leftPtr] > nums2[rightPtr] {
			// make space for nums2[rightPtr] and insert nums2[rightPtr]
			for i := total - 1; i > leftPtr; i-- {
				nums1[i], nums1[i-1] = nums1[i-1], nums1[i]
			}
			nums1[leftPtr] = nums2[rightPtr]
			rightPtr++
			m++
		}
		leftPtr++
	}

	if rightPtr < n {
		nums1 = append(nums1[:m], nums2[rightPtr:]...)
	}
}

func MergeSortedArrayOptimal(nums1 []int, m int, nums2 []int, n int) {
	// O(m+n) key intution is to start from end so as to not overwrite
	leftPtr, rightPtr := m-1, n-1
	curPtr := m + n - 1

	for rightPtr >= 0 {
		if leftPtr >= 0 && nums1[leftPtr] > nums2[rightPtr] {
			nums1[curPtr] = nums1[leftPtr]
			leftPtr--
		} else {
			nums1[curPtr] = nums2[rightPtr]
			rightPtr--
		}
		curPtr--
	}
}
