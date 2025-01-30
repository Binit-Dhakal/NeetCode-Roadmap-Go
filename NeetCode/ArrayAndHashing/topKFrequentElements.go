package arrayandhashing

import (
	"container/heap"
	"sort"
)

// Normal Sorting Algorithm
func TopKSorting(nums []int, k int) []int {
	// O(nlogn) -> O(n)
	freqMap := make(map[int]int, 0)

	for _, num := range nums {
		freqMap[num]++
	}

	count := make([][2]int, 0)
	for ele, freq := range freqMap {
		count = append(count, [2]int{ele, freq})
	}

	sort.Slice(count, func(i, j int) bool {
		return count[i][1] > count[j][1]
	})

	result := make([]int, k)
	for idx := range k {
		result[idx] = count[idx][0]
	}

	return result
}

func TopKBucketSort(nums []int, k int) []int {
	// O(n) - O(n)
	type Freq int
	freqMap := make(map[int]Freq, 0)

	for _, num := range nums {
		freqMap[num]++
	}

	countFreq := make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		countFreq[freq] = append(countFreq[freq], num)
	}

	var result []int
	for i := len(countFreq) - 1; i > 0; i-- {
		for _, res := range countFreq[i] {
			result = append(result, res)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}

type FreqStruct struct {
	Elem int
	Freq int
}

type Tuple [2]int

type IntHeap []Tuple

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(Tuple))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TopKHeap(nums []int, k int) []int {
	freqMap := make(map[int]int, 0)

	for _, num := range nums {
		freqMap[num]++
	}

	var intHeap IntHeap
	for elem, freq := range freqMap {
		var elemTuple [2]int
		elemTuple[0] = elem
		elemTuple[1] = freq
		intHeap = append(intHeap, elemTuple)
	}

	heap.Init(&intHeap)

	result := make([]int, k)
	for idx := range k {
		result[idx] = heap.Pop(&intHeap).(Tuple)[0]
	}

	return result
}
