package heappq

type KthLargest struct {
	pq []int
	k  int
}

func Constructor(k int, nums []int) KthLargest {
	kLargest := KthLargest{
		pq: make([]int, 0, k),
		k:  k,
	}
	for _, val := range nums {
		kLargest.Add(val)
	}
	return kLargest
}

func (this *KthLargest) Add(val int) int {
	if len(this.pq) < this.k {
		this.pq = append(this.pq, val)
		this.heapifyUp(len(this.pq) - 1)
	} else if this.pq[0] < val {
		this.pq[0] = val
		this.heapifyDown(0)
	}

	return this.pq[0]
}

func (this *KthLargest) lessFunc(a, b int) bool {
	return this.pq[a] < this.pq[b]
}

func (this *KthLargest) heapifyUp(childIdx int) {
	for childIdx > 0 {
		parentIdx := (childIdx - 1) / 2
		if !this.lessFunc(childIdx, parentIdx) {
			return
		}

		this.swap(parentIdx, childIdx)
		childIdx = parentIdx
	}
}

func (this *KthLargest) swap(aIdx, bIdx int) {
	this.pq[aIdx], this.pq[bIdx] = this.pq[bIdx], this.pq[aIdx]
}

func (this *KthLargest) heapifyDown(parentIdx int) {
	for {
		l, r := 2*parentIdx+1, 2*parentIdx+2
		lessIdx := parentIdx
		n := len(this.pq)

		if l < n && this.lessFunc(l, lessIdx) {
			lessIdx = l
		}

		if r < n && this.lessFunc(r, lessIdx) {
			lessIdx = r
		}

		if lessIdx == parentIdx {
			return
		}

		this.swap(lessIdx, parentIdx)
		parentIdx = lessIdx
	}
}
