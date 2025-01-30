package slidingwindow

import (
	"container/heap"
	"fmt"
	"slices"
)

func SlidingWindowMaxBF(nums []int, k int) []int {
	leftPtr := 0
	res := make([]int, 0)

	for rightPtr := 0; rightPtr < len(nums); rightPtr++ {
		if rightPtr-k+1 >= 0 {
			res = append(res, slices.Max(nums[leftPtr:rightPtr+1]))
			leftPtr++
		}
	}
	return res
}

func SlidingWindowMaxDeque(nums []int, k int) []int {
	deque := make([]int, 0)
	res := make([]int, 0)

	for idx, num := range nums {
		for len(deque) > 0 && num > deque[len(deque)-1] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, num)

		if idx-k+1 >= 0 {
			res = append(res, deque[0])
			// pop first element if largest is the first element in deque
			if nums[idx-k+1] == deque[0] {
				deque = deque[1:]
			}
		}

	}
	return res
}

type Pair struct {
	value, idx int
}

type PairHeap []*Pair

func (h PairHeap) Len() int { return len(h) }

func (h PairHeap) Less(i, j int) bool {
	return h[i].value > h[j].value || (h[i].value == h[j].value && h[i].idx < h[j].idx)
}

func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x any) {
	item := x.(*Pair)
	*h = append(*h, item)
}

func (h *PairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h PairHeap) Peek() any {
	return h[0]
}

func SlidingWindowMaxPQ(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)

	resIdx := 0
	pq := &PairHeap{}
	heap.Init(pq)

	for idx, num := range nums {
		heap.Push(pq, &Pair{idx: idx, value: num})

		if idx >= k-1 {
			for pq.Peek().(*Pair).idx <= idx-k {
				heap.Pop(pq)
			}
			res[resIdx] = pq.Peek().(*Pair).value
			resIdx++
		}

	}
	return res
}
