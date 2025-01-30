package linkedlist

import "container/heap"

func MergeKSortedLists(lists []*ListNode) *ListNode {
	// O(n*k) -- O(1)  where n is total number of nodes across all lists and k is total number of list
	dummyNode := &ListNode{}
	head := dummyNode

	for {
		minVal := int(^uint(0) >> 1)
		var minNodeIdx int = -1
		for i := range lists {
			if lists[i] == nil {
				continue
			}
			if lists[i].Val < int(minVal) {
				minVal = lists[i].Val
				minNodeIdx = i
			}
		}

		if minNodeIdx == -1 {
			// we have reached end of all listnode
			break
		}
		// add min val to the resultNode and current.next the minNode
		dummyNode.Next = lists[minNodeIdx]
		dummyNode = dummyNode.Next
		lists[minNodeIdx] = lists[minNodeIdx].Next
	}

	dummyNode.Next = nil

	return head.Next
}

type NodeHeap []*ListNode

func (n NodeHeap) Len() int { return len(n) }

func (n NodeHeap) Less(i, j int) bool { return n[i].Val < n[j].Val }

func (n NodeHeap) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

func (n *NodeHeap) Push(x any) {
	*n = append(*n, x.(*ListNode))
}

func (n *NodeHeap) Pop() any {
	old := *n
	l := len(old)
	x := old[l-1]
	*n = old[0 : l-1]
	return x
}

func MergeKSortedListsHeap(lists []*ListNode) *ListNode {
	nodes := &NodeHeap{}
	heap.Init(nodes)

	for _, list := range lists {
		if list != nil {
			heap.Push(nodes, list)
		}
	}
	res := &ListNode{}
	current := res
	for len(*nodes) != 0 {
		node := heap.Pop(nodes).(*ListNode)
		if node.Next != nil {
			heap.Push(nodes, node.Next)
		}
		current.Next = node
		current = current.Next
	}
	current.Next = nil
	return res.Next
}

func MergeKSortedListsDAC(lists []*ListNode) *ListNode {
	if len(lists) < 2 {
		return lists[0]
	}
	mid := len(lists) / 2

	first := MergeKSortedListsDAC(lists[:mid])
	second := MergeKSortedListsDAC(lists[mid:])

	var merge func(*ListNode, *ListNode) *ListNode
	merge = func(ln1, ln2 *ListNode) *ListNode {
		dummyNode := &ListNode{}
		res := dummyNode
		for ln1 != nil && ln2 != nil {
			if ln1.Val < ln2.Val {
				dummyNode.Next = ln1
				ln1 = ln1.Next
			} else {
				dummyNode.Next = ln2
				ln2 = ln2.Next
			}
			dummyNode = dummyNode.Next
		}

		if ln1 != nil {
			dummyNode.Next = ln1
		}

		if ln2 != nil {
			dummyNode.Next = ln2
		}
		return res.Next
	}

	return merge(first, second)
}
