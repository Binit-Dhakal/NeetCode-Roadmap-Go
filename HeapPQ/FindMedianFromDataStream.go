package heappq

import "container/heap"

type LNode struct {
	num  int
	next *LNode
}

type MedianFinder struct {
	n    int
	node *LNode
}

func ConstructMedian() MedianFinder {
	dummy := LNode{}
	return MedianFinder{node: &dummy}
}

func (this *MedianFinder) ShowNums() []int {
	cur := this.node.next
	res := []int{}
	for cur != nil {
		res = append(res, cur.num)
		cur = cur.next
	}

	return res
}

func (this *MedianFinder) AddNum(num int) {
	cur := this.node
	for cur.next != nil {
		if cur.next.num >= num {
			cur.next = &LNode{num: num, next: cur.next}
			this.n++
			break
		}
		cur = cur.next
	}

	if cur.next == nil {
		cur.next = &LNode{num: num}
		this.n++
	}
}

func (this *MedianFinder) FindMedian() float64 {
	medianHelper := func(n int) (*LNode, *LNode) {
		cur := this.node.next
		tmp := cur
		for n > 0 {
			tmp = cur
			cur = cur.next
			n--
		}
		return cur, tmp
	}

	l1, l2 := medianHelper(this.n / 2)
	if this.n%2 != 0 {
		return float64(l1.num)
	}

	return float64(l1.num+l2.num) / 2
}

type MIntHeap struct {
	data []int
	less func(i, j int) bool
}

func (h MIntHeap) Len() int { return len(h.data) }

func (h MIntHeap) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }

func (h MIntHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h MIntHeap) Peek() int { return h.data[0] }

func (h *MIntHeap) Push(x any) {
	h.data = append(h.data, x.(int))
}

func (h *MIntHeap) Pop() any {
	old := h.data
	num := len(old)
	x := old[num-1]
	h.data = old[:num-1]
	return x
}

func NewMinHeap() *MIntHeap {
	return &MIntHeap{
		data: []int{},
		less: func(i, j int) bool { return i < j },
	}
}

func NewMaxHeap() *MIntHeap {
	return &MIntHeap{
		data: []int{},
		less: func(i, j int) bool { return i > j },
	}
}

type MedianFinderH struct {
	maxHeap *MIntHeap
	minHeap *MIntHeap
}

func ConstructMedianHeap() MedianFinderH {
	maxHeap := NewMaxHeap()
	minHeap := NewMinHeap()

	heap.Init(maxHeap)
	heap.Init(minHeap)
	return MedianFinderH{
		maxHeap: maxHeap,
		minHeap: minHeap,
	}
}

func (this *MedianFinderH) ShowHeaps() ([]int, []int) {
	return this.maxHeap.data, this.minHeap.data
}

func (this *MedianFinderH) AddNum(num int) {
	if this.minHeap.Len() > 0 && num > this.minHeap.Peek() {
		heap.Push(this.minHeap, num)
	} else {
		heap.Push(this.maxHeap, num)
	}

	n1, n2 := this.maxHeap.Len(), this.minHeap.Len()
	if n1-n2 > 1 {
		x := heap.Pop(this.maxHeap)
		heap.Push(this.minHeap, x.(int))
	} else if n2-n1 > 1 {
		x := heap.Pop(this.minHeap)
		heap.Push(this.maxHeap, x.(int))
	}
}

func (this *MedianFinderH) FindMedian() float64 {
	n1, n2 := this.maxHeap.Len(), this.minHeap.Len()
	if n1 > n2 {
		return float64(this.maxHeap.Peek())
	} else if n1 < n2 {
		return float64(this.minHeap.Peek())
	}
	return float64(this.minHeap.Peek()+this.maxHeap.Peek()) / 2
}
