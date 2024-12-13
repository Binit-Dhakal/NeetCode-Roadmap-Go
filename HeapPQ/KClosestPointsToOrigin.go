package heappq

import "container/heap"

type PointInfo [2]int

type DistanceHeap []PointInfo // (index, distance)

func (h DistanceHeap) Len() int { return len(h) }

func (h DistanceHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }

func (h DistanceHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *DistanceHeap) Push(x any) {
	*h = append(*h, x.(PointInfo))
}

func (h *DistanceHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func KClosestPointsToOrigin(points [][]int, k int) [][]int {
	h := &DistanceHeap{}
	for i, point := range points {
		*h = append(*h, [2]int{i, (point[0]*point[0] + point[1]*point[1])})
	}
	heap.Init(h)

	res := make([][]int, 0)
	for k > 0 {
		x := heap.Pop(h).(PointInfo)
		res = append(res, points[x[0]])
		k--
	}
	return res
}

func KClosestPointsToOriginQuickSelect(points [][]int, k int) [][]int {
	l, r := 0, len(points)-1
	partionIdx := len(points)

	euclideanDist := func(x int, y int) int {
		return x*x + y*y
	}
	var partition func(l int, r int) int
	partition = func(l, r int) int {
		p := points[r]
		pivotDistance := euclideanDist(p[0], p[1])

		i := l
		for j := l; j < r; j++ {
			if pivotDistance >= euclideanDist(points[j][0], points[j][1]) {
				points[j], points[i] = points[i], points[j]
				i++
			}
		}
		points[i], points[r] = points[r], points[i]
		return i
	}

	for partionIdx != k {
		partionIdx = partition(l, r)
		if partionIdx < k {
			l = partionIdx + 1
		}
		if partionIdx > k {
			r = partionIdx - 1
		}
	}
	return points[:k]
}
