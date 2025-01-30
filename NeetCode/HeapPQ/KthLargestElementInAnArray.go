package heappq

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func KthLargestElementInAnArray(nums []int, k int) int {
	h := IntHeap(nums[:k])
	heap.Init(&h)

	for _, num := range nums[k:] {
		if num > h[0] {
			heap.Pop(&h)
			heap.Push(&h, num)

		}
	}
	return h[0]
}

func KthLargestElementInAnArrayQuickSelect(nums []int, k int) int {
	// Timeout
	l, r := 0, len(nums)-1

	kIdx := len(nums) - k
	var i int = len(nums)
	for i != kIdx {
		pivotIdx := r
		i = l
		for j := l; j < r; j++ {
			if nums[j] <= nums[pivotIdx] {
				nums[j], nums[i] = nums[i], nums[j]
				i++
			}
		}

		nums[pivotIdx], nums[i] = nums[i], nums[pivotIdx]

		if i < kIdx {
			l = i + 1
		} else if i > kIdx {
			r = i - 1
		}
	}
	return nums[kIdx]
}

func KthLargestElementInAnArrayBucketSort(nums []int, k int) int {
	count := make([]int, 2*10000+1)
	for _, num := range nums {
		count[num+10000]++
	}

	c := 0
	for i := 2 * 10000; i >= 0; i-- {
		c += count[i]
		if c >= k {
			return i - 10000
		}
	}

	return 0
}
