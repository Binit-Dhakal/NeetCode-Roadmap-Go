package heappq

import (
	"container/heap"
)

type WeightHeap struct {
	stones []int
}

func (w WeightHeap) Len() int { return len(w.stones) }

func (w WeightHeap) Less(i, j int) bool { return w.stones[i] > w.stones[j] }

func (w WeightHeap) Swap(i, j int) { w.stones[i], w.stones[j] = w.stones[j], w.stones[i] }

func (w *WeightHeap) Push(x any) {
	w.stones = append(w.stones, x.(int))
}

func (w *WeightHeap) Pop() any {
	old := w.stones
	n := len(old)
	x := old[n-1]
	w.stones = old[:n-1]
	return x
}

func LastStoneWeight(stones []int) int {
	wh := &WeightHeap{stones: stones}
	heap.Init(wh)
	for wh.Len() > 1 {
		x, y := heap.Pop(wh).(int), heap.Pop(wh).(int)
		if x > y {
			heap.Push(wh, x-y)
		}
	}

	if wh.Len() == 1 {
		return wh.stones[0]
	} else {
		return 0
	}
}

func LastStoneWeightBucketSort(stones []int) int {
	maxStone := 0
	for _, stone := range stones {
		if stone > maxStone {
			maxStone = stone
		}
	}

	bucket := make([]int, maxStone+1)

	for _, stone := range stones {
		bucket[stone]++
	}

	first, second := maxStone, maxStone
	for first > 0 {
		if bucket[first]%2 == 0 {
			first--
			continue
		}

		j := min(first-1, second)
		for j > 0 && bucket[j] == 0 {
			j--
		}
		if j == 0 {
			return first
		}

		second = j
		bucket[first]--
		bucket[second]--
		bucket[first-second]++

		first = max(second, first-second)
	}

	return first
}
